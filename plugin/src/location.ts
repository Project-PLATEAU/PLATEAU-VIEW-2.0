import { PostMessageProps, MouseEvent } from "@web/extensions/location/types";

import html from "../dist/web/location/core/index.html?raw";
import googleAnalyticsHtml from "../dist/web/location/modals/googleAnalytics/index.html?raw";
import terrainHtml from "../dist/web/location/modals/terrain/index.html?raw";

const reearth = (globalThis as any).reearth;

reearth.ui.show(html, { width: 350, height: 40 });

if (reearth.viewport.isMobile) {
  reearth.ui.close();
}

reearth.on("mousemove", (mousedata: MouseEvent) => {
  reearth.ui.postMessage({ type: "mousedata", payload: mousedata });
});

reearth.on("cameramove", () => {
  reearth.ui.postMessage({
    type: "getScreenLocation",
    payload: getScreenLocation(),
  });
});

reearth.on("message", ({ action }: PostMessageProps) => {
  if (action === "initLocation") {
    reearth.ui.postMessage({
      type: "getScreenLocation",
      payload: getScreenLocation(),
    });
  } else if (action === "googleModalOpen") {
    reearth.modal.show(googleAnalyticsHtml, { background: "transparent", width: 572, height: 670 });
  } else if (action === "terrainModalOpen") {
    reearth.modal.show(terrainHtml, {
      background: "transparent",
      width: 572,
      height: 222,
    });
  } else if (action === "modalClose") {
    reearth.modal.close();
  }
});

function getScreenLocation() {
  return {
    point1: reearth.scene.getLocationFromScreen(
      reearth.viewport.width / 2,
      reearth.viewport.height - 1,
    ),
    point2: reearth.scene.getLocationFromScreen(
      reearth.viewport.width / 2 + 1,
      reearth.viewport.height - 1,
    ),
  };
}
