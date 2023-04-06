import mapBing from "@web/extensions/sidebar/core/assets/bgmap_bing.png";
import bgmap_darkmatter from "@web/extensions/sidebar/core/assets/bgmap_darkmatter.png";
import bgmap_gsi from "@web/extensions/sidebar/core/assets/bgmap_gsi.png";
import bgmap_tokyo from "@web/extensions/sidebar/core/assets/bgmap_tokyo.png";
import { ReearthApi } from "@web/extensions/sidebar/types";
import { useCallback, useMemo } from "react";

import { MapViewData, BaseMapData, ViewSelection } from "./types";

const mapViewData: MapViewData[] = [
  {
    key: "3d-terrain",
    title: "3D地形",
  },
  { key: "3d-smooth", title: "3D地形なし" },
  { key: "2d", title: "2D" },
];

const baseMapData: BaseMapData[] = [
  {
    key: "tokyo",
    title: "全国最新写真 (シームレス)",
    icon: bgmap_tokyo,
    url: [
      "https://cyberjapandata.gsi.go.jp/xyz/seamlessphoto/{z}/{x}/{y}.jpg",
      "https://gic-plateau.s3.ap-northeast-1.amazonaws.com/2020/ortho/tiles/{z}/{x}/{y}.png",
    ],
  },
  {
    key: "bing",
    title: "空中写真 (Bing)",
    icon: mapBing,
    tile_type: "default_label",
  },
  {
    key: "gsi",
    title: "地理院地図 (淡色)",
    icon: bgmap_gsi,
    url: ["https://cyberjapandata.gsi.go.jp/xyz/pale/{z}/{x}/{y}.png"],
  },
  {
    key: "dark-matter",
    title: "Dark Matter",
    icon: bgmap_darkmatter,
    url: ["https://basemaps.cartocdn.com/dark_all/{z}/{x}/{y}.jpg"],
  },
];

export type Props = {
  overrides: ReearthApi;
  onOverridesUpdate: (updatedProperties: Partial<ReearthApi>) => void;
};

export default ({ overrides, onOverridesUpdate }: Props) => {
  const {
    default: {
      sceneMode: currentSceneMode,
      depthTestAgainstTerrain: currentHideUnderground,
      allowEnterGround: currentAllowEnterGround,
    } = {},
    terrain: { terrain: currentTerrain } = {},
    tiles: currentTiles,
    light: { lightType: currentLightType } = {},
  } = overrides;

  const currentView: ViewSelection = useMemo(
    () => (currentSceneMode === "2d" ? "2d" : !currentTerrain ? "3d-smooth" : "3d-terrain"),
    [currentSceneMode, currentTerrain],
  );

  const handleViewChange = useCallback(
    (view: ViewSelection) => {
      let newView: Partial<ReearthApi> = {};
      if (view === "3d-terrain") {
        newView = {
          default: {
            sceneMode: "3d",
          },
          terrain: {
            terrain: true,
            terrainType: "cesiumion",
            terrainCesiumIonAccessToken:
              "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI5N2UyMjcwOS00MDY1LTQxYjEtYjZjMy00YTU0ZTg5MmViYWQiLCJpZCI6ODAzMDYsImlhdCI6MTY0Mjc0ODI2MX0.dkwAL1CcljUV7NA7fDbhXXnmyZQU_c-G5zRx8PtEcxE",
            terrainCesiumIonAsset: "770371",
          },
        };
      } else if (view === "3d-smooth") {
        newView = {
          default: {
            sceneMode: "3d",
          },
          terrain: {
            terrain: false,
          },
        };
      } else if (view === "2d") {
        newView = {
          default: {
            sceneMode: "2d",
          },
          terrain: {
            terrain: false,
          },
        };
      }
      onOverridesUpdate(newView);
    },
    [onOverridesUpdate],
  );

  const handleTileChange = useCallback(
    (tile: BaseMapData) => {
      if (tile.url) {
        onOverridesUpdate({
          tiles: tile.url.map((t, i, a) => ({
            id: tile.key + (a.length > 1 ? "_" + (i + 1) : ""),
            tile_url: t,
            tile_type: "url",
          })),
        });
      } else if (tile.tile_type) {
        onOverridesUpdate({ tiles: [{ id: tile.key, tile_type: tile.tile_type }] });
      }
    },
    [onOverridesUpdate],
  );

  const handleHideUnderGround = useCallback(() => {
    onOverridesUpdate({ default: { depthTestAgainstTerrain: !currentHideUnderground } });
  }, [currentHideUnderground, onOverridesUpdate]);

  const handleAllowEnterGround = useCallback(() => {
    onOverridesUpdate({ default: { allowEnterGround: !currentAllowEnterGround } });
  }, [currentAllowEnterGround, onOverridesUpdate]);

  const handleUseSunLight = useCallback(() => {
    onOverridesUpdate(
      currentLightType === "sunLight"
        ? {
            light: {
              lightType: "directionalLight",
              lightColor: "#ffffffff",
              lightIntensity: 2,
              lightDirectionX: 0.7650124487710819,
              lightDirectionY: -0.6418383470612292,
              lightDirectionZ: -0.05291020191779678,
            },
          }
        : {
            light: {
              lightType: "sunLight",
              lightIntensity: 2,
            },
          },
    );
  }, [currentLightType, onOverridesUpdate]);

  return {
    mapViewData,
    baseMapData,
    currentView,
    currentTiles,
    currentHideUnderground,
    currentAllowEnterGround,
    currentLightType,
    handleViewChange,
    handleTileChange,
    handleHideUnderGround,
    handleAllowEnterGround,
    handleUseSunLight,
  };
};
