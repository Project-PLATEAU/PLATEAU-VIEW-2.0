import { StoryItem } from "./core/components/content/common/DatasetCard/Field/Fields/types";
import { DataCatalogItem } from "./core/types";

type ActionType =
  | "init"
  | "initDataCatalog"
  | "initPopup"
  | "initWelcome"
  | "storageSave"
  | "storageFetch"
  | "storageKeys"
  | "storageDelete"
  | "updateDataCatalog"
  | "updateProject"
  | "screenshot"
  | "screenshotPreview"
  | "screenshotSave"
  | "addDatasetToScene"
  | "updateDatasetInScene"
  | "updateDatasetVisibility"
  | "removeDatasetFromScene"
  | "removeAllDatasetsFromScene"
  | "updateDataset"
  | "catalogModalOpen"
  | "triggerCatalogOpen"
  | "saveSearchTerm"
  | "saveExpandedFolders"
  | "saveDataset"
  | "saveFilter"
  | "triggerHelpOpen"
  | "mapModalOpen"
  | "clipModalOpen"
  | "modalClose"
  | "msgFromModal"
  | "helpPopupOpen"
  | "popupClose"
  | "msgToPopup"
  | "msgFromPopup"
  | "mobileDropdownOpen"
  | "msgToMobileDropdown"
  | "checkIfMobile"
  | "extendPopup"
  | "minimize"
  | "buildingSearchOpen"
  | "buildingSearchOverride"
  | "buildingSearchInit"
  | "groupSelectOpen"
  | "saveGroups"
  | "cameraFlyTo"
  | "cameraLookAt"
  | "getCurrentCamera"
  | "storyPlay"
  | "storyEdit"
  | "storyEditFinish"
  | "storyDelete"
  | "updateClippingBox"
  | "removeClippingBox"
  | "update3dtilesShow"
  | "reset3dtilesShow"
  | "update3dtilesTransparency"
  | "reset3dtilesTransparency"
  | "update3dtilesColor"
  | "reset3dtilesColor"
  | "findTileset"
  | "update3dtilesShadow"
  | "reset3dtilesShadow"
  | "infoboxFieldsFetch"
  | "infoboxFieldsSaved"
  | "findLayerByDataID"
  | "getOverriddenLayerByDataID"
  | "unselect"
  | "mobileDatasetAdd"
  | "mobileDatasetRemove"
  | "mobileDatasetRemoveAll"
  | "mobileProjectDatasetsUpdate"
  | "mobileBuildingSearch"
  | "mobileProjectSceneUpdate"
  | "mobileDatasetUpdate"
  | "mobileCatalogOpen"
  | "initMobileCatalog"
  | "updateMVTRaster"
  | "getCenterOnScreen";
// FIXME(@keiya01): support auto csv field complement
// | "getLocationNamesFromCSVFeatureProperty"

export type PostMessageProps = { action: ActionType; payload?: any };

export type Project = {
  sceneOverrides: ReearthApi;
  datasets: DataCatalogItem[];
  userStory?: Omit<StoryItem, "id">;
};

export type ReearthApi = {
  default?: {
    camera?: Camera;
    sceneMode?: SceneMode;
    depthTestAgainstTerrain?: boolean;
    allowEnterGround?: boolean;
  };
  terrain?: {
    terrain?: boolean;
    terrainType?: "cesiumion";
    terrainCesiumIonAsset?: string;
    terrainCesiumIonAccessToken?: string;
    terrainCesiumIonUrl?: string;
    terrainExaggeration?: number;
    terrainExaggerationRelativeHeight?: number;
    depthTestAgainstTerrain?: boolean;
  };
  tiles?: Tile[];
  atmosphere?: {
    enable_sun?: boolean;
    enable_lighting?: boolean;
    ground_atmosphere?: boolean;
    sky_atmosphere?: boolean;
    shadows?: boolean;
    fog?: boolean;
    fog_density?: number;
    brightness_shift?: number;
    hue_shift?: number;
    surturation_shift?: number;
  };
  timeline?: {
    current?: string;
    start?: string;
    stop?: string;
  };
  light?: {
    lightType?: "sunLight" | "directionalLight";
    lightDirectionX?: number;
    lightDirectionY?: number;
    lightDirectionZ?: number;
    lightColor?: string;
    lightIntensity?: number;
  };
};

type SceneMode = "3d" | "2d";

type Tile = {
  id: string;
  tile_url?: string;
  tile_type: string;
};

export type Camera = {
  lat?: number;
  lng?: number;
  altitude?: number;
  height?: number;
  heading?: number;
  pitch?: number;
  roll?: number;
  fov?: number;
};

type PluginActionType = "storyShare";

export type PluginMessage = {
  data: {
    action: PluginActionType;
    payload: any;
  };
  sender: string;
};
