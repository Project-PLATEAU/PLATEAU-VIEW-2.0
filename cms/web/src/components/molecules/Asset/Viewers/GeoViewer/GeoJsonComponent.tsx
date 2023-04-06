import { GeoJsonDataSource } from "cesium";
import { ComponentProps, useCallback } from "react";
import { GeoJsonDataSource as ResiumGeoJsonDataSource, useCesium } from "resium";

type Props = ComponentProps<typeof ResiumGeoJsonDataSource>;

const GeoJsonComponent: React.FC<Props> = ({ data }) => {
  const { viewer } = useCesium();

  const handleLoad = useCallback(
    async (ds: GeoJsonDataSource) => {
      try {
        await viewer?.zoomTo(ds);
        ds.show = true;
      } catch (error) {
        console.error(error);
      }
    },
    [viewer],
  );

  return <ResiumGeoJsonDataSource data={data} clampToGround={true} onLoad={handleLoad} />;
};

export default GeoJsonComponent;
