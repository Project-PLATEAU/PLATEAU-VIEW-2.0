import {
  PostMessageProps,
  Viewport,
  PluginMessage,
  PluginExtensionInstance,
  StoryShare,
  StorySaveData,
  StoryCancelPlay,
} from "@web/extensions/storytelling/types";

import html from "../dist/web/storytelling/core/index.html?raw";
import storyeditorHtml from "../dist/web/storytelling/modals/sceneEditor/index.html?raw";

const reearth = (globalThis as any).reearth;

reearth.ui.show(html, { width: 122, height: 40, extended: false });

let sidebarId: string;
const getSidebarId = () => {
  if (sidebarId) return;
  sidebarId = reearth.plugins.instances.find(
    (instance: PluginExtensionInstance) => instance.extensionId === "sidebar",
  )?.id;
};
getSidebarId();

reearth.on("pluginmessage", (pluginMessage: PluginMessage) => {
  reearth.ui.postMessage(pluginMessage.data);
});

reearth.on("message", ({ action, payload }: PostMessageProps) => {
  if (action === "resize") {
    reearth.ui.resize(...payload);
  } else if (action === "getViewport") {
    reearth.ui.postMessage({
      action: "getViewport",
      payload: reearth.viewport,
    });
  } else if (action === "sceneCapture") {
    reearth.ui.postMessage({
      action: "sceneCapture",
      payload: reearth.camera.position,
    });
  } else if (action === "sceneView") {
    reearth.camera.flyTo(payload, { duration: 1.5 });
  } else if (action === "sceneRecapture") {
    reearth.ui.postMessage({
      action: "sceneRecapture",
      payload: { camera: reearth.camera.position, id: payload },
    });
  } else if (action === "sceneEdit") {
    reearth.modal.show(storyeditorHtml, { background: "transparent", width: 580, height: 320 });
    reearth.modal.postMessage({
      action: "sceneEdit",
      payload,
    });
  } else if (action === "sceneEditorClose") {
    reearth.modal.close();
  } else if (action === "sceneSave") {
    reearth.ui.postMessage({
      action: "sceneSave",
      payload,
    });
    reearth.modal.close();
  } else if (action === "storyShare") {
    getSidebarId();
    if (!sidebarId) return;
    reearth.plugins.postMessage(sidebarId, {
      action: "storyShare",
      payload,
    } as StoryShare);
  } else if (action === "storySaveData") {
    getSidebarId();
    if (!sidebarId) return;
    reearth.plugins.postMessage(sidebarId, {
      action: "storySaveData",
      payload,
    } as StorySaveData);
  } else if (action === "storyCancelPlay") {
    getSidebarId();
    if (!sidebarId) return;
    reearth.plugins.postMessage(sidebarId, {
      action: "storyCancelPlay",
      payload,
    } as StoryCancelPlay);
  }
});

reearth.on("resize", (viewport: Viewport) => {
  reearth.ui.postMessage({
    action: "viewport",
    payload: viewport,
  });
});
