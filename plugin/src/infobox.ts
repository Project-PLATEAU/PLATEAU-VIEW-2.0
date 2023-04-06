import {
  PostMessageProps,
  PluginExtensionInstance,
  PluginMessage,
} from "@web/extensions/infobox/types";

import html from "../dist/web/infobox/core/index.html?raw";

import { inEditor } from "./utils/ineditor";

const reearth = (globalThis as any).reearth;

reearth.ui.show(html);

let sidebarId: string;
const getSidebarId = () => {
  if (sidebarId) return;
  sidebarId = reearth.plugins.instances.find(
    (instance: PluginExtensionInstance) => instance.extensionId === "sidebar",
  )?.id;
};
getSidebarId();

const infoboxFieldsFetch = () => {
  getSidebarId();
  if (!sidebarId) return;
  reearth.plugins.postMessage(sidebarId, {
    action: "infoboxFieldsFetch",
  });
};

reearth.on("message", ({ action, payload }: PostMessageProps) => {
  if (action === "init") {
    reearth.ui.postMessage({
      action: "getInEditor",
      payload: inEditor(),
    });
    infoboxFieldsFetch();
  } else if (action === "saveTemplate") {
    getSidebarId();
    if (!sidebarId) return;
    reearth.plugins.postMessage(sidebarId, {
      action: "infoboxFieldsSave",
      payload,
    });
  }
});

reearth.on("pluginmessage", (pluginMessage: PluginMessage) => {
  if (pluginMessage.data.action === "infoboxFieldsFetch") {
    if (reearth.layers.selectedFeature && pluginMessage.data.payload) {
      reearth.ui.postMessage({
        action: "fillData",
        payload: {
          properties: reearth.layers.selectedFeature.properties,
          template: pluginMessage.data.payload,
          fldInfo: pluginMessage.data.payload?.fldInfo,
        },
      });
    } else {
      reearth.ui.postMessage({
        action: "setEmpty",
      });
    }
  } else if (pluginMessage.data.action === "infoboxFieldsSaved") {
    reearth.ui.postMessage({
      action: "saveFinish",
    });
  }
});

let lastSelectedFeatureId: string | undefined;

reearth.on("select", () => {
  const featureId = reearth.layers.selectedFeature?.id;
  if (featureId !== lastSelectedFeatureId) {
    lastSelectedFeatureId = featureId;
    infoboxFieldsFetch();

    reearth.ui.postMessage({
      action: "setLoading",
    });
  }
});
