import { DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import { PostMessageProps, Project, PluginMessage } from "@web/extensions/sidebar/types";
import { isObject, mergeWith, omit, cloneDeep, merge as lodashMerge } from "lodash";

import html from "../dist/web/sidebar/core/index.html?raw";
import clipVideoHtml from "../dist/web/sidebar/modals/clipVideo/index.html?raw";
import dataCatalogHtml from "../dist/web/sidebar/modals/datacatalog/index.html?raw";
import mapVideoHtml from "../dist/web/sidebar/modals/mapVideo/index.html?raw";
import welcomeScreenHtml from "../dist/web/sidebar/modals/welcomescreen/index.html?raw";
import buildingSearchHtml from "../dist/web/sidebar/popups/buildingSearch/index.html?raw";
import groupSelectPopupHtml from "../dist/web/sidebar/popups/groupSelect/index.html?raw";
import helpPopupHtml from "../dist/web/sidebar/popups/help/index.html?raw";
import mobileDropdownHtml from "../dist/web/sidebar/popups/mobileDropdown/index.html?raw";

import { inEditor } from "./utils/ineditor";
import { proxyGTFS } from "./utils/proxy";

const defaultProject: Project = {
  sceneOverrides: {
    default: {
      camera: {
        lat: 35.65075152248653,
        lng: 139.7617718208305,
        altitude: 2219.7187259974316,
        heading: 6.132702058010316,
        pitch: -0.5672459184621266,
        roll: 0.00019776785897196447,
        fov: 1.0471975511965976,
        height: 2219.7187259974316,
      },
      sceneMode: "3d",
      depthTestAgainstTerrain: false,
    },
    terrain: {
      terrain: true,
      terrainType: "cesiumion",
      terrainCesiumIonAccessToken:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI5N2UyMjcwOS00MDY1LTQxYjEtYjZjMy00YTU0ZTg5MmViYWQiLCJpZCI6ODAzMDYsImlhdCI6MTY0Mjc0ODI2MX0.dkwAL1CcljUV7NA7fDbhXXnmyZQU_c-G5zRx8PtEcxE",
      terrainCesiumIonAsset: "770371",
    },
    tiles: [
      {
        id: "tokyo_1",
        tile_url: "https://cyberjapandata.gsi.go.jp/xyz/seamlessphoto/{z}/{x}/{y}.jpg",
        tile_type: "url",
      },
      {
        id: "tokyo_2",
        tile_url:
          "https://gic-plateau.s3.ap-northeast-1.amazonaws.com/2020/ortho/tiles/{z}/{x}/{y}.png",
        tile_type: "url",
      },
    ],
    atmosphere: { shadows: true },
    light: {
      lightType: "directionalLight",
      lightColor: "#ffffffff",
      lightIntensity: 2,
      lightDirectionX: 0.7650124487710819,
      lightDirectionY: -0.6418383470612292,
      lightDirectionZ: -0.05291020191779678,
    },
  },
  datasets: [],
  userStory: undefined,
};

type PluginExtensionInstance = {
  id: string;
  name: string;
  pluginId: string;
  extensionId: string;
  runTimes?: number;
};

const reearth = (globalThis as any).reearth;

let welcomePageIsOpen = false;
let mobileDropdownIsOpen = false;
let openedpendingBuildingSearchDataID: string | null = null;
let pendingBuildingSearchData: { viewport: any; data: any } | undefined = undefined;

// this is used for infobox
let currentSelected: string | undefined = undefined;

const defaultLocation = { zone: "outer", section: "left", area: "middle" };
const mobileLocation = { zone: "outer", section: "center", area: "top" };

let addedDatasets: [dataID: string, status: "showing" | "hidden", layerID?: string][] = [];
let templates: Template[] | undefined = undefined;

let searchTerm = "";
let expandedFolders: { id?: string; name?: string }[] = [];
let dataset: DataCatalogItem | undefined = undefined;
let filter = "city";

const sidebarInstance: PluginExtensionInstance = reearth.plugins.instances.find(
  (i: PluginExtensionInstance) => i.id === reearth.widget.id,
);

// ************************************************
// initializations

reearth.ui.show(html, { extended: true });

if (
  sidebarInstance.runTimes === 1 ||
  (sidebarInstance.runTimes === 2 && reearth.viewport.isMobile)
) {
  reearth.visualizer.overrideProperty(defaultProject.sceneOverrides);
  reearth.clientStorage.setAsync("draftProject", defaultProject);

  if (reearth.viewport.isMobile) {
    reearth.clientStorage.setAsync("isMobile", true);
    reearth.widget.moveTo(mobileLocation);
  } else {
    reearth.clientStorage.setAsync("isMobile", false);
  }
  reearth.clientStorage.getAsync("doNotShowWelcome").then((value: any) => {
    if (!value && !inEditor()) {
      reearth.modal.show(welcomeScreenHtml, {
        width: reearth.viewport.width,
        height: reearth.viewport.height,
      });
      welcomePageIsOpen = true;
    }
  });
} else {
  reearth.clientStorage.getAsync("isMobile").then((value: any) => {
    if (reearth.viewport.isMobile) {
      if (!value) {
        reearth.widget.moveTo(mobileLocation);
        reearth.clientStorage.setAsync("isMobile", true);
      }
    } else {
      if (value) {
        reearth.widget.moveTo(defaultLocation);
        reearth.clientStorage.setAsync("isMobile", false);
      }
    }
  });
}
// ************************************************

reearth.on("message", ({ action, payload }: PostMessageProps) => {
  // Mobile specific
  if (action === "mobileDropdownOpen") {
    reearth.popup.show(mobileDropdownHtml, {
      position: "bottom",
      width: reearth.viewport.width - 12,
    });
    mobileDropdownIsOpen = true;
  } else if (action === "msgToMobileDropdown") {
    reearth.popup.postMessage({ action: "msgToPopup", payload });
  }

  // Sidebar
  if (action === "init") {
    reearth.clientStorage.getAsync("draftProject").then((draftProject: Project) => {
      const outBoundPayload = {
        projectID: reearth.viewport.query.share || reearth.viewport.query.projectID,
        inEditor: inEditor(),
        catalogURL: reearth.widget.property.default?.catalogURL ?? "",
        catalogProjectName: reearth.widget.property.default?.catalogProjectName ?? "",
        reearthURL: reearth.widget.property.default?.reearthURL ?? "",
        backendURL: reearth.widget.property.default?.plateauURL ?? "",
        backendProjectName: reearth.widget.property.default?.projectName ?? "",
        backendAccessToken: reearth.widget.property.default?.plateauAccessToken ?? "",
        enableGeoPub: reearth.widget.property.default?.enableGeoPub ?? false,
        draftProject,
        searchTerm,
      };
      reearth.ui.postMessage({ action, payload: outBoundPayload });
    });
  } else if (action === "storageSave") {
    reearth.clientStorage.setAsync(payload.key, payload.value);
  } else if (action === "storageFetch") {
    reearth.clientStorage.getAsync(payload.key).then((value: any) => {
      reearth.ui.postMessage({
        type: "getAsync",
        payload: value,
      });
    });
  } else if (action === "storageKeys") {
    reearth.clientStorage.keysAsync().then((value: any) => {
      reearth.ui.postMessage({
        type: "keysAsync",
        payload: value,
      });
    });
  } else if (action === "storageDelete") {
    reearth.clientStorage.deleteAsync(payload.key);
  } else if (action === "updateProject") {
    reearth.visualizer.overrideProperty(payload.sceneOverrides);
    reearth.clientStorage.setAsync("draftProject", payload);
  } else if (action === "addDatasetToScene") {
    if (addedDatasets.find(d => d[0] === payload.dataset.dataID)) {
      const idx = addedDatasets.findIndex(ad => ad[0] === payload.dataset.dataID);
      addedDatasets[idx][1] = "showing";
      reearth.layers.show(addedDatasets[idx][2]);
    } else {
      const data = createLayer(payload.dataset, payload.overrides);
      console.log("DATA to add", data);
      const layerID = reearth.layers.add(data);
      const idx = addedDatasets.push([payload.dataset.dataID, "showing", layerID]);
      if (!payload.dataset.visible) {
        reearth.layers.hide(addedDatasets[idx][2]);
        addedDatasets[idx][1] = "hidden";
      }
      reearth.modal.postMessage({
        action: "updateDataCatalog",
        payload: { updatedDatasetDataIDs: addedDatasets.map(d => d[0]) },
      });
    }
  } else if (action === "updateDatasetInScene") {
    const layerId = addedDatasets.find(ad => ad[0] === payload.dataID)?.[2];
    const layer = reearth.layers.findById(layerId);
    reearth.layers.override(
      layerId,
      layer.data.type === "gtfs" ? proxyGTFS(payload.overrides) : payload.overrides,
    );
  } else if (action === "updateDatasetVisibility") {
    const idx = addedDatasets.findIndex(ad => ad[0] === payload.dataID);
    if (payload.hide) {
      reearth.layers.hide(addedDatasets[idx][2]);
      addedDatasets[idx][1] = "hidden";
    } else {
      reearth.layers.show(addedDatasets[idx][2]);
      addedDatasets[idx][1] = "showing";
    }
  } else if (action === "removeDatasetFromScene") {
    reearth.layers.delete(addedDatasets.find(ad => ad[0] === payload)?.[2]);
    const idx = addedDatasets.findIndex(ad => ad[0] === payload);
    addedDatasets.splice(idx, 1);
    if (openedpendingBuildingSearchDataID === payload) {
      reearth.popup.close();
    }
  } else if (action === "removeAllDatasetsFromScene") {
    reearth.layers.delete(...addedDatasets.map(ad => ad[2]));
    addedDatasets = [];
    if (openedpendingBuildingSearchDataID) {
      reearth.popup.close();
    }
  } else if (action === "updateDataset") {
    reearth.ui.postMessage({ action, payload });
  } else if (
    action === "screenshot" ||
    action === "screenshotPreview" ||
    action === "screenshotSave"
  ) {
    reearth.ui.postMessage({
      action,
      payload: reearth.scene.captureScreen(undefined, 0.01),
    });
    reearth.popup.postMessage({
      action,
      payload: reearth.scene.captureScreen(undefined, 0.01),
    });
  } else if (action === "msgFromModal") {
    reearth.ui.postMessage({ action, payload });
  } else if (action === "minimize") {
    if (payload) {
      reearth.ui.resize(undefined, undefined, false);
    } else {
      reearth.ui.resize(350, undefined, true);
    }
  } else if (action === "catalogModalOpen") {
    reearth.modal.show(dataCatalogHtml, { background: "transparent" });
    templates = payload.templates;
  } else if (action === "triggerCatalogOpen") {
    reearth.ui.postMessage({ action });
  } else if (action === "saveSearchTerm") {
    searchTerm = payload.searchTerm;
  } else if (action === "saveExpandedFolders") {
    expandedFolders = [...payload.expandedFolders];
  } else if (action === "saveDataset") {
    dataset = { ...payload.dataset };
  } else if (action === "saveFilter") {
    filter = payload.filter;
  } else if (action === "triggerHelpOpen") {
    reearth.ui.postMessage({ action });
  } else if (action === "modalClose") {
    reearth.modal.close();
    welcomePageIsOpen = false;
  } else if (action === "initDataCatalog") {
    if (reearth.viewport.isMobile) {
      reearth.popup.postMessage({
        action,
        payload: {
          searchTerm,
          expandedFolders,
          dataset,
          filter,
        },
      });
    } else {
      reearth.modal.postMessage({
        action,
        payload: {
          addedDatasets: addedDatasets.map(d => d[0]),
          inEditor: inEditor(),
          backendProjectName: reearth.widget.property.default?.projectName ?? "",
          backendAccessToken: reearth.widget.property.default?.plateauAccessToken ?? "",
          backendURL: reearth.widget.property.default?.plateauURL ?? "",
          catalogURL: reearth.widget.property.default?.catalogURL ?? "",
          catalogProjectName: reearth.widget.property.default?.catalogProjectName ?? "",
          enableGeoPub: reearth.widget.property.default?.enableGeoPub ?? false,
          templates,
          searchTerm,
          expandedFolders,
          dataset,
          filter,
        },
      });
    }
  } else if (action === "helpPopupOpen") {
    reearth.popup.show(helpPopupHtml, { position: "right-start", offset: 4 });
  } else if (action === "groupSelectOpen") {
    reearth.popup.show(groupSelectPopupHtml, { position: "right", offset: 4 });
    reearth.popup.postMessage({ action: "groupSelectInit", payload });
  } else if (action === "saveGroups") {
    reearth.ui.postMessage({ action, payload });
    reearth.popup.close();
  } else if (action === "initPopup") {
    if (payload?.type === "buildingSearch") {
      reearth.popup.postMessage({
        type: "msgToPopup",
        payload: { ...pendingBuildingSearchData, type: "buildingSearchData" },
      });
      pendingBuildingSearchData = undefined;
    } else {
      reearth.ui.postMessage({ action });
    }
  } else if (action === "initWelcome") {
    reearth.modal.postMessage({ type: "msgToModal", message: reearth.viewport.isMobile });
  } else if (action === "msgToPopup") {
    reearth.popup.postMessage({ action: "msgToPopup", payload });
  } else if (action === "msgFromPopup") {
    if (payload.height) {
      reearth.popup.update({ height: payload.height, width: reearth.viewport.width - 12 });
    } else if (payload.currentTab) {
      reearth.ui.postMessage({ action: "msgFromPopup", payload: payload.currentTab });
    }
  } else if (action === "popupClose") {
    reearth.popup.close();
    reearth.ui.postMessage({ action });
    mobileDropdownIsOpen = false;
  } else if (action === "mapModalOpen") {
    reearth.modal.show(mapVideoHtml, { background: "transparent" });
  } else if (action === "clipModalOpen") {
    reearth.modal.show(clipVideoHtml, { background: "transparent" });
  } else if (action === "cameraFlyTo") {
    if (Array.isArray(payload)) {
      reearth.camera.flyTo(...payload);
    } else {
      const layerID = addedDatasets.find(ad => ad[0] === payload)?.[2];
      reearth.camera.flyTo(layerID);
    }
  } else if (action === "cameraLookAt") {
    if (reearth.scene?.property?.terrain?.terrain) {
      reearth.scene
        .sampleTerrainHeight(payload[0].lng, payload[0].lat)
        .then((terrainHeight: number | undefined) => {
          reearth.camera.lookAt(
            {
              ...payload[0],
              height: (payload[0].height ?? 0) + (terrainHeight ?? 0),
            },
            payload[1],
          );
        });
    } else {
      reearth.camera.lookAt(...payload);
    }
  } else if (action === "getCurrentCamera") {
    reearth.ui.postMessage({ action, payload: reearth.camera.position });
    reearth.popup.postMessage({ action, payload: reearth.camera.position });
  } else if (action === "checkIfMobile") {
    reearth.ui.postMessage({ action, payload: reearth.viewport.isMobile });
  } else if (action === "extendPopup") {
    reearth.popup.update({
      height: reearth.viewport.height - 68,
      width: reearth.viewport.width - 12,
    });
  } else if (action === "findLayerByDataID") {
    const { dataID } = payload;
    const layerID = addedDatasets.find(a => a[0] === dataID)?.[2];
    const layer = reearth.layers.findById(layerID);
    const overriddenLayer = reearth.layers.overridden.find((l: any) => l.id === layerID);
    if (reearth.viewport.isMobile) {
      reearth.popup.postMessage({
        action,
        payload: {
          dataID,
          layer: {
            id: layer?.id,
            data: { ...layer?.data, ...overriddenLayer?.data },
          },
        },
      });
    } else {
      reearth.ui.postMessage({
        action,
        payload: {
          dataID,
          layer: {
            id: layer?.id,
            data: { ...layer?.data, ...overriddenLayer?.data },
          },
        },
      });
    }
  } else if (action === "getOverriddenLayerByDataID") {
    const { dataID } = payload;
    const layerID = addedDatasets.find(l => l[0] === dataID)?.[2];
    const overriddenLayer = reearth.layers.overridden.find((l: any) => l.id === layerID);
    if (reearth.viewport.isMobile) {
      reearth.popup.postMessage({
        action,
        payload: {
          overriddenLayer,
        },
      });
    } else {
      reearth.ui.postMessage({
        action,
        payload: {
          overriddenLayer,
        },
      });
    }
  } else if (action === "updateMVTRaster") {
    const { layerId, maxzoom } = payload;
    reearth.layers.override(layerId, {
      raster: {
        maximumLevel: maxzoom,
      },
    });
  } else if (action === "unselect") {
    reearth.layers.select();
  } else if (action === "getCenterOnScreen") {
    const viewport = reearth.viewport;
    const centerOnScreen = reearth.scene.getLocationFromScreen(
      viewport.width / 2,
      viewport.height / 2,
    );
    if (viewport.isMobile) {
      reearth.popup.postMessage({
        action,
        payload: {
          centerOnScreen,
        },
      });
    } else {
      reearth.ui.postMessage({
        action,
        payload: {
          centerOnScreen,
        },
      });
    }
  }

  // ************************************************
  // Mobile actions
  else if (
    action === "mobileDatasetAdd" ||
    action === "mobileDatasetUpdate" ||
    action === "mobileDatasetRemove" ||
    action === "mobileDatasetRemoveAll" ||
    action === "mobileProjectDatasetsUpdate" ||
    action === "mobileProjectSceneUpdate" ||
    action === "mobileBuildingSearch"
  ) {
    reearth.ui.postMessage({ action, payload });
  } else if (action === "initMobileCatalog") {
    if (payload) {
      reearth.popup.postMessage({ action, payload });
    } else {
      reearth.ui.postMessage({ action });
    }
  } else if (action === "mobileCatalogOpen") {
    reearth.popup.postMessage({ action, payload });
  }

  // ************************************************
  // Building Search
  else if (action === "buildingSearchOverride") {
    reearth.ui.postMessage({
      action,
      payload,
    });
  } else if (action === "buildingSearchOpen") {
    if (openedpendingBuildingSearchDataID) {
      reearth.ui.postMessage({
        action: "buildingSearchClose",
        payload: {
          dataID: openedpendingBuildingSearchDataID,
        },
      });
      reearth.popup.postMessage({
        type: "msgToPopup",
        payload: { viewport: reearth.viewport, data: payload, type: "buildingSearchData" },
      });
      pendingBuildingSearchData = undefined;
    } else {
      pendingBuildingSearchData = {
        viewport: reearth.viewport,
        data: payload,
      };
      reearth.popup.show(buildingSearchHtml, {
        position: reearth.viewport.isMobile ? "bottom-start" : "right-start",
        offset: {
          mainAxis: 4,
          crossAxis: reearth.viewport.isMobile ? reearth.viewport.width * 0.05 : 0,
        },
      });
    }
    openedpendingBuildingSearchDataID = payload.dataID;
  }

  // ************************************************
  // Story
  else if (
    action === "storyPlay" ||
    action === "storyEdit" ||
    action === "storyEditFinish" ||
    action === "storyDelete"
  ) {
    const storyTellingWidgetId = reearth.plugins.instances.find(
      (instance: PluginExtensionInstance) => instance.extensionId === "storytelling",
    )?.id;
    if (!storyTellingWidgetId) return;
    reearth.plugins.postMessage(storyTellingWidgetId, {
      action,
      payload,
    });
  }

  // ************************************************
  // for infobox
  else if (action === "infoboxFieldsFetch") {
    const infoboxInstanceId = reearth.plugins.instances.find(
      (instance: PluginExtensionInstance) => instance.extensionId === "infobox",
    )?.id;
    if (!infoboxInstanceId) return;
    reearth.plugins.postMessage(infoboxInstanceId, {
      action: "infoboxFieldsFetch",
      payload,
    });
  } else if (action === "infoboxFieldsSaved") {
    const infoboxInstanceId = reearth.plugins.instances.find(
      (instance: PluginExtensionInstance) => instance.extensionId === "infobox",
    )?.id;
    if (!infoboxInstanceId) return;
    reearth.plugins.postMessage(infoboxInstanceId, {
      action: "infoboxFieldsSaved",
    });
  }

  // ************************************************
  // For 3dtiles
  if (action === "findTileset") {
    const { dataID } = payload;
    const tilesetLayerID = addedDatasets.find(a => a[0] === dataID)?.[2];
    const tilesetLayer = reearth.layers.findById(tilesetLayerID);
    const overriddenTilesetLayer = reearth.layers.overridden.find(
      (l: any) => l.id === tilesetLayerID,
    );
    const postMsgResp = {
      action,
      payload: {
        dataID,
        layer: {
          id: tilesetLayer.id,
          data: {
            ...(tilesetLayer?.data ?? {}),
            ...(overriddenTilesetLayer?.data ?? {}),
          },
          ["3dtiles"]: {
            ...(tilesetLayer?.["3dtiles"] ?? {}),
            ...(overriddenTilesetLayer?.["3dtiles"] ?? {}),
          },
        },
      },
    };
    reearth.ui.postMessage(postMsgResp);
    reearth.popup.postMessage(postMsgResp);
  }
});

reearth.on("update", () => {
  reearth.ui.postMessage({
    type: "extended",
    payload: reearth.widget.extended,
  });
});

reearth.on("resize", () => {
  // Modals
  if (welcomePageIsOpen) {
    reearth.modal.update({
      width: reearth.viewport.width,
      height: reearth.viewport.height,
    });
    reearth.modal.postMessage({ type: "msgToModal", payload: reearth.viewport.isMobile });
  }
  // Popups
  if (mobileDropdownIsOpen) {
    reearth.popup.update({
      width: reearth.viewport.width - 10,
    });
  }

  if (openedpendingBuildingSearchDataID) {
    reearth.popup.postMessage({
      type: "msgToPopup",
      payload: { ...reearth.viewport, type: "resize" },
    });
    if (reearth.viewport.isMobile) {
      reearth.popup.update({
        offset: {
          mainAxis: 4,
          crossAxis: reearth.viewport.isMobile ? reearth.viewport.width * 0.05 : 0,
        },
      });
    }
  }
});

reearth.on("pluginmessage", (pluginMessage: PluginMessage) => {
  if (pluginMessage.data.action === "storyShare") {
    reearth.ui.postMessage(pluginMessage.data);
  } else if (pluginMessage.data.action === "storySaveData") {
    reearth.ui.postMessage(pluginMessage.data);
  } else if (pluginMessage.data.action === "infoboxFieldsFetch") {
    reearth.ui.postMessage({
      action: "infoboxFieldsFetch",
      payload: addedDatasets.find(ad => ad[2] === currentSelected)?.[0],
    });
  } else if (pluginMessage.data.action === "infoboxFieldsSave") {
    reearth.ui.postMessage(pluginMessage.data);
  }
});

let currentSelectedFeatureId: string;

const findMergedLayer = (id: string | undefined) => {
  const l = reearth.layers.findById(id);
  const overridden = reearth.layers.overridden.find((l: any) => l.id === id);
  return {
    ...l,
    ...overridden,
    data: {
      ...(l?.data ?? {}),
      ...(overridden?.data ?? {}),
    },
    "3dtiles": {
      ...(l?.["3dtiles"] ?? {}),
      ...(overridden?.["3dtiles"] ?? {}),
    },
  };
};

reearth.on("select", (selected: string | undefined) => {
  const prevSelected = currentSelected ?? selected;
  // this is used for infobox
  currentSelected = selected;

  const isSameLayer = currentSelected === prevSelected;

  const featureId = reearth.layers.selectedFeature?.id; // For 3dtiles
  const prevSelectedFeatureId = currentSelectedFeatureId ?? featureId;
  currentSelectedFeatureId = featureId;

  let nextConditions: any[] | undefined;
  let shouldUpdateTilesetColor = false;

  // Reset previous select color for 3dtiles
  const prevOverriddenLayer = findMergedLayer(prevSelected);
  const prevCondition = prevOverriddenLayer?.["3dtiles"]?.color?.expression?.conditions;
  if (prevOverriddenLayer && prevOverriddenLayer.data.type === "3dtiles") {
    shouldUpdateTilesetColor =
      !!currentSelectedFeatureId &&
      !prevCondition?.find(
        (c: [string, string]) => c[0] === `\${id} === "${currentSelectedFeatureId}"`,
      );
    nextConditions = prevCondition?.filter(
      (c: [string, string]) => !c[0].startsWith('${id} === "'),
    );
  }

  if (
    ((!currentSelected && !currentSelectedFeatureId) || !isSameLayer) &&
    prevOverriddenLayer.data.type === "3dtiles" &&
    prevSelectedFeatureId &&
    prevCondition?.find((c: [string, string]) => c[0] === `\${id} === "${prevSelectedFeatureId}"`)
  ) {
    reearth.layers.override(prevSelected, {
      "3dtiles": {
        color: {
          expression: {
            conditions: nextConditions ? nextConditions : [["true", "color('white')"]],
          },
        },
      },
    });
    if (isSameLayer) {
      return;
    }
  }

  // Handle select color for 3dtiles
  const overriddenLayer = findMergedLayer(currentSelected);
  if (
    overriddenLayer?.data?.type === "3dtiles" &&
    currentSelectedFeatureId &&
    shouldUpdateTilesetColor
  ) {
    reearth.layers.override(currentSelected, {
      "3dtiles": {
        color: {
          expression: {
            conditions: [
              [`\${id} === "${currentSelectedFeatureId}"`, "color('red')"],
              ...(isSameLayer && nextConditions
                ? nextConditions
                : overriddenLayer?.["3dtiles"]?.color?.expression?.conditions ?? [
                    ["true", "color('white')"],
                  ]),
            ],
          },
        },
      },
    });
  }
});

reearth.on("popupclose", () => {
  if (
    openedpendingBuildingSearchDataID &&
    !(reearth.viewport.isMobile && pendingBuildingSearchData)
  ) {
    reearth.ui.postMessage({
      action: "buildingSearchClose",
      payload: {
        dataID: openedpendingBuildingSearchDataID,
      },
    });
    openedpendingBuildingSearchDataID = null;
  }
});

function createLayer(dataset: DataCatalogItem, overrides?: any) {
  const format = dataset.format?.toLowerCase().replace(/\s/g, "");
  const merge = (obj1: any, obj2: any): any => {
    const merged = cloneDeep(obj1);
    mergeWith(merged, obj2, (mergedValue, obj2Value) => {
      if (isObject(mergedValue)) {
        return merge(mergedValue, obj2Value);
      }
      return obj2Value !== undefined ? obj2Value : mergedValue;
    });
    return merged;
  };
  return {
    type: "simple",
    title: dataset.name,
    data: {
      type: dataset.config?.data?.[0].type.toLowerCase().replace(/\s/g, "") ?? format,
      url: dataset.config?.data?.[0].url ?? dataset.url,
      layers: dataset.config?.data?.[0].layer ?? dataset.layers,
      ...(format === "wms" ? { parameters: { transparent: "true", format: "image/png" } } : {}),
      ...(["luse", "lsld", "urf", "rail", "tran"].includes(dataset.type_en) ||
      (dataset.type_en === "tran" && format === "mvt")
        ? { jsonProperties: ["attributes"] }
        : {}),
      ...(overrides?.data || {}),
    },
    visible: true,
    infobox: lodashMerge(
      [
        "bldg",
        "tran",
        "frn",
        "veg",
        "luse",
        "lsld",
        "urf",
        "fld",
        "htd",
        "tnm",
        "ifld",
        "rail",
      ].includes(dataset.type_en)
        ? {
            blocks: [
              {
                pluginId: reearth.plugins.instances.find(
                  (i: PluginExtensionInstance) => i.name === "plateau-plugin",
                ).pluginId,
                extensionId: "infobox",
              },
            ],
            property: {
              default: {
                bgcolor: "#d9d9d9ff",
                heightType: "auto",
                showTitle: false,
                size: "medium",
              },
            },
          }
        : {},
      overrides?.infobox ?? {},
      infoboxGlobal,
    ),
    ...(overrides !== undefined
      ? merge(defaultOverrides, omit(overrides, ["data", "infobox"]))
      : format === "geojson" || format === "czml"
      ? defaultOverrides
      : format === "gtfs"
      ? proxyGTFS(overrides)
      : format === "mvt"
      ? {
          polygon: {},
        }
      : format === "wms"
      ? {
          raster: {
            alpha: 0.8,
          },
        }
      : { ...(omit(overrides, ["data", "infobox"]) ?? {}) }),
  };
}

const infoboxGlobal = {
  property: {
    default: {
      unselectOnClose: true,
    },
  },
};

const defaultOverrides = {
  resource: {
    clampToGround: true,
  },
  marker: {
    heightReference: "clamp",
  },
  polygon: {
    heightReference: "clamp",
  },
  polyline: {
    clampToGround: true,
  },
};
