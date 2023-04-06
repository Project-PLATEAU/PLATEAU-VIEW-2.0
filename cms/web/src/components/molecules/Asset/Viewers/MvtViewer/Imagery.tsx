import styled from "@emotion/styled";
import { VectorTileFeature } from "@mapbox/vector-tile";
import {
  Cartesian3,
  ImageryLayer,
  ImageryLayerCollection,
  Math,
  BoundingSphere,
  HeadingPitchRange,
  type Viewer,
} from "cesium";
import { MVTImageryProvider } from "cesium-mvt-imagery-provider";
import md5 from "js-md5";
import { useCallback, useEffect, useMemo, useState } from "react";
import { useCesium } from "resium";

import AutoComplete from "@reearth-cms/components/atoms/AutoComplete";

const defaultCameraPosition: [number, number, number] = [139.767052, 35.681167, 100];
const defaultOffset = new HeadingPitchRange(0, Math.toRadians(-90.0), 3000000);
const normalOffset = new HeadingPitchRange(0, Math.toRadians(-90.0), 200000);

type Props = {
  url: string;
  handleProperties: (prop: Property) => void;
  selectFeature?: (selected: boolean) => void;
};

export type Property = { [k: string]: string | number | boolean };

// TODO: these two types should be imported from cesium-mvt-imagery-provider library instead
type URLTemplate = `http${"s" | ""}://${string}/{z}/{x}/{y}${string}`;
type TileCoordinates = {
  x: number;
  y: number;
  level: number;
};

export const Imagery: React.FC<Props> = ({ url, handleProperties, selectFeature }) => {
  const { viewer } = useCesium() as { viewer: Viewer | undefined };
  const [selectedFeature, setSelectFeature] = useState<string>();
  const [urlTemplate, setUrlTemplate] = useState<URLTemplate>(url as URLTemplate);
  const [currentLayer, setCurrentLayer] = useState("");
  const [layers, setLayers] = useState<string[]>([]);

  const zoomTo = useCallback(
    ([lng, lat, height]: [lng: number, lat: number, height: number], useDefaultRange?: boolean) => {
      viewer?.camera.flyToBoundingSphere(
        new BoundingSphere(Cartesian3.fromDegrees(lng, lat, height)),
        {
          duration: 0,
          offset: useDefaultRange ? defaultOffset : normalOffset,
        },
      );
    },
    [viewer],
  );

  const loadData = useCallback(
    async (url: string) => {
      try {
        const data = await fetchLayers(url);
        if (data) {
          setUrlTemplate(`${data.base}/{z}/{x}/{y}.mvt` as URLTemplate);
          setLayers(data.layers ?? []);
          setCurrentLayer(data.layers?.[0] || "");
        }
        zoomTo(data?.center || defaultCameraPosition, !data?.center);
      } catch (error) {
        console.error(error);
      }
    },
    [zoomTo],
  );

  const style = useCallback(
    (f: VectorTileFeature, tile: TileCoordinates) => {
      const fid = idFromGeometry(f.loadGeometry(), tile);
      return {
        strokeStyle: "white",
        fillStyle: selectedFeature === fid ? "orange" : "red",
        lineWidth: 1,
      };
    },
    [selectedFeature],
  );

  const onSelectFeature = useCallback(
    (feature: VectorTileFeature, tileCoords: TileCoordinates) => {
      const id = idFromGeometry(feature.loadGeometry(), tileCoords);
      selectFeature?.(true);
      setSelectFeature(id);
      handleProperties(feature.properties);
    },
    [handleProperties, selectFeature],
  );

  useEffect(() => {
    loadData(url);
  }, [loadData, url]);

  useEffect(() => {
    const imageryProvider = new MVTImageryProvider({
      urlTemplate,
      layerName: currentLayer,
      style,
      onSelectFeature,
    });

    if (viewer) {
      const layers: ImageryLayerCollection = viewer.scene.imageryLayers;
      const currentLayer: ImageryLayer = layers.addImageryProvider(imageryProvider);
      currentLayer.alpha = 0.5;

      return () => {
        layers.remove(currentLayer);
      };
    }
  }, [
    viewer,
    selectedFeature,
    url,
    urlTemplate,
    currentLayer,
    layers,
    handleProperties,
    selectFeature,
    onSelectFeature,
    style,
  ]);

  const handleChange = useCallback((value: unknown) => {
    if (typeof value !== "string") return;
    setCurrentLayer(value);
  }, []);

  const options = useMemo(() => layers.map(l => ({ label: l, value: l })), [layers]);

  return (
    <StyledInput
      placeholder="Layer name"
      value={currentLayer}
      options={options}
      onChange={handleChange}
      onSelect={handleChange}
    />
  );
};

const StyledInput = styled(AutoComplete)`
  position: absolute;
  top: 10px;
  left: 10px;
  width: 147px;
`;

const getMvtBaseUrl = (url: string) => {
  const templateRegex = /\/\d{1,5}\/\d{1,5}\/\d{1,5}\.\w+$/;
  const compressedExtRegex = /\.zip|\.7z$/;
  const nameRegex = /\/\w+\.\w+$/;
  const base = url.match(templateRegex)
    ? url.replace(templateRegex, "")
    : url.match(compressedExtRegex)
    ? url.replace(compressedExtRegex, "")
    : url.replace(nameRegex, "");
  return base;
};

const fetchLayers = async (url: string) => {
  const base = getMvtBaseUrl(url);
  const res = await fetch(`${base}/metadata.json`);
  if (!res.ok) return;
  return { ...parseMetadata(await res.json()), base };
};

type TileCoords = { x: number; y: number; level: number };

const idFromGeometry = (
  geometry: ReturnType<VectorTileFeature["loadGeometry"]>,
  tile: TileCoords,
) => {
  const id = [tile.x, tile.y, tile.level, ...geometry.flatMap(i => i.map(j => [j.x, j.y]))].join(
    ":",
  );

  const hash = md5.create();
  hash.update(id);

  return hash.hex();
};

export function parseMetadata(
  json: any,
):
  | { layers: string[]; center: [lng: number, lat: number, height: number] | undefined }
  | undefined {
  if (!json) return;

  let layers: string[] = [];
  if (typeof json.json === "string") {
    try {
      layers = JSON.parse(json.json)?.vector_layers?.map((l: any): string => l.id);
    } catch {
      // ignore
    }
  }

  let center: [lng: number, lat: number, height: number] | undefined = undefined;
  try {
    if (typeof json.center === "string") {
      const c = (json.center as string).split(",", 3).map(s => parseFloat(s));
      center = [c[0], c[1], c[2]];
    }
  } catch {
    // ignore
  }

  return { layers, center };
}
