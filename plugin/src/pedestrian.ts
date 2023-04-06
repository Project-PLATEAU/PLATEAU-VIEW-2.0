import { PostMessageProps, MouseEvent, Camera } from "@web/extensions/pedestrian/types";

import html from "../dist/web/pedestrian/core/index.html?raw";
import pedestrianControllerHtml from "../dist/web/pedestrian/popups/pedestrianController/index.html?raw";

const reearth = (globalThis as any).reearth;

if (reearth.viewport.isMobile) {
  reearth.ui.close();
  reearth.camera.enableScreenSpaceController(true);
} else {
  reearth.ui.show(html);

  // status
  let mode: "ready" | "picking" | "pedestrian" = "ready";
  let initCamera: Camera | undefined = undefined;
  let controllerShown = false;

  const initControllerOptions = {
    width: 208,
    height: 356,
    position: "bottom-end",
    offset: 4,
  };

  const fullControllerOptions = {
    width: 208,
    height: 568,
    position: "bottom-end",
    offset: 4,
  };

  const flags = {
    looking: false,
    moveForward: false,
    moveBackward: false,
    moveUp: false,
    moveDown: false,
    moveLeft: false,
    moveRight: false,
  };

  const startPos: { x: number | undefined; y: number | undefined } = {
    x: 0,
    y: 0,
  };
  const lookFactor = 0.00005;
  const lookAmount = {
    x: 0,
    y: 0,
  };

  const oppositeMove = new Map<keyof typeof flags, keyof typeof flags>([
    ["moveForward", "moveBackward"],
    ["moveBackward", "moveForward"],
    ["moveUp", "moveDown"],
    ["moveDown", "moveUp"],
    ["moveLeft", "moveRight"],
    ["moveRight", "moveLeft"],
  ]);

  const updateCamera = () => {
    let moveRate = reearth.camera.position.height / 100.0;
    if (moveRate < 1) moveRate = 1;

    if (flags.moveForward) {
      reearth.camera.moveForward(moveRate);
    }
    if (flags.moveBackward) {
      reearth.camera.moveBackward(moveRate);
    }
    if (flags.moveUp) {
      reearth.camera.moveUp(moveRate);
    }
    if (flags.moveDown) {
      if (reearth.camera.position.height > 3) {
        reearth.camera.moveDown(moveRate);
      }
    }
    if (flags.moveLeft) {
      reearth.camera.moveLeft(moveRate);
    }
    if (flags.moveRight) {
      reearth.camera.moveRight(moveRate);
    }
    if (flags.looking) {
      reearth.camera.lookHorizontal(lookAmount.x);
      reearth.camera.lookVertical(lookAmount.y);
    }

    if (
      flags.moveForward ||
      flags.moveBackward ||
      flags.moveUp ||
      flags.moveDown ||
      flags.moveRight ||
      flags.moveLeft
    ) {
      reearth.camera.moveOverTerrain(1.8);
    }
  };

  const handleCameraMove = ({ moveType, on }: { moveType: keyof typeof flags; on: boolean }) => {
    flags[moveType] = on;
    if (on) {
      const op = oppositeMove.get(moveType);
      if (op) {
        flags[op] = false;
      }
    }
  };

  const handlePedestrianExit = () => {
    const curCamera = reearth.camera.position;
    if (initCamera) {
      reearth.camera.flyTo(
        {
          lng: curCamera.lng,
          lat: curCamera.lat,
          height: initCamera?.height,
          heading: initCamera?.heading,
          pitch: initCamera?.pitch,
          roll: initCamera?.roll,
          fov: initCamera?.fov,
        },
        { duration: 2 },
      );
    }
    mode = "ready";
    initCamera = undefined;
    reearth.camera.enableScreenSpaceController(true);

    if (controllerShown) {
      reearth.popup.update(initControllerOptions);
    }
  };

  reearth.on("message", ({ action, payload }: PostMessageProps) => {
    if (action === "pedestrianShow") {
      controllerShown = true;
      reearth.popup.show(pedestrianControllerHtml, initControllerOptions);
    } else if (action === "pedestrianClose") {
      controllerShown = false;
      reearth.popup.close();
      handlePedestrianExit();
    } else if (action === "pickingStart") {
      mode = "picking";
    } else if (action === "pedestrianExit") {
      handlePedestrianExit();
    } else if (action === "cameraMove") {
      handleCameraMove(payload);
    } else if (action === "controllerReady") {
      if (controllerShown) {
        reearth.popup.update(initControllerOptions);
        reearth.popup.postMessage({
          type: "controllerReady",
        });
      }
    }
  });

  reearth.on("click", (mouseData: MouseEvent) => {
    if (mode === "picking" && mouseData.lat !== undefined && mouseData.lng !== undefined) {
      initCamera = reearth.camera.position;
      reearth.camera.enableScreenSpaceController(false);
      reearth.camera.flyToGround(
        {
          lng: mouseData.lng,
          lat: mouseData.lat,
          height: 100,
          heading: initCamera?.heading ?? 0,
          pitch: 0,
          roll: 0,
          fov: initCamera?.fov ?? 45,
        },
        {
          duration: 2,
        },
        20,
      );
      reearth.popup.postMessage({
        type: "pickingDone",
        payload: mouseData,
      });
      mode = "pedestrian";
      reearth.popup.update(fullControllerOptions);
    }
  });

  reearth.on("mousedown", (mousedata: MouseEvent) => {
    if (mode !== "pedestrian") return;
    if (mousedata.x !== undefined && mousedata.y !== undefined) {
      startPos.x = mousedata.x;
      startPos.y = mousedata.y;
      flags.looking = true;
    }
  });

  reearth.on("mousemove", (mousedata: MouseEvent) => {
    if (mode !== "pedestrian") return;
    if (
      flags.looking &&
      mousedata.x !== undefined &&
      mousedata.y !== undefined &&
      startPos.x !== undefined &&
      startPos.y !== undefined
    ) {
      lookAmount.x = (mousedata.x - startPos.x) * lookFactor;
      lookAmount.y = (mousedata.y - startPos.y) * lookFactor;
    }
  });

  reearth.on("mouseup", () => {
    if (mode !== "pedestrian") return;
    startPos.x = undefined;
    startPos.y = undefined;
    flags.looking = false;
  });

  reearth.on("popupclose", () => {
    if (controllerShown) {
      controllerShown = false;
      handlePedestrianExit();
    }
  });

  reearth.on("cameramove", () => {
    if (mode !== "pedestrian") return;
    reearth.popup.postMessage({
      type: "updateMiniMap",
      payload: reearth.camera.position,
    });
  });

  reearth.on("tick", updateCamera);
}
