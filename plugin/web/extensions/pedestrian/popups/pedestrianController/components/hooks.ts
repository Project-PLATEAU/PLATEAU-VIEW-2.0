import type { Camera } from "@web/extensions/pedestrian/types";
import { postMsg } from "@web/extensions/pedestrian/utils";
import L, { Map as LeafletMap } from "leaflet";
import { useCallback, useEffect, useState, useRef } from "react";

export default () => {
  const [mainButtonText, setMainButtonText] = useState<"始める" | "終わる">("始める");
  const [mode, setMode] = useState<"ready" | "picking" | "pedestrian">("ready");
  const [visible, setVisible] = useState(false);

  const [moveForwardOn, setMoveForwardOn] = useState(false);
  const [moveBackwardOn, setMoveBackwardOn] = useState(false);
  const [moveLeftOn, setMoveLeftOn] = useState(false);
  const [moveRightOn, setMoveRightOn] = useState(false);
  const [moveUpOn, setMoveUpOn] = useState(false);
  const [moveDownOn, setMoveDownOn] = useState(false);

  const handleMoveForwardClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveForwardOn : enable;
      setMoveForwardOn(on);
      postMsg("cameraMove", { moveType: "moveForward", on });
      if (on && moveBackwardOn) {
        setMoveBackwardOn(false);
      }
    },
    [moveForwardOn, moveBackwardOn],
  );

  const handleMoveBackwardClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveBackwardOn : enable;
      setMoveBackwardOn(on);
      postMsg("cameraMove", { moveType: "moveBackward", on });
      if (on && moveForwardOn) {
        setMoveForwardOn(false);
      }
    },
    [moveBackwardOn, moveForwardOn],
  );

  const handleMoveLeftClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveLeftOn : enable;
      setMoveLeftOn(on);
      postMsg("cameraMove", { moveType: "moveLeft", on });
      if (on && moveRightOn) {
        setMoveRightOn(false);
      }
    },
    [moveLeftOn, moveRightOn],
  );

  const handleMoveRightClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveRightOn : enable;
      setMoveRightOn(on);
      postMsg("cameraMove", { moveType: "moveRight", on });
      if (on && moveLeftOn) {
        setMoveLeftOn(false);
      }
    },
    [moveRightOn, moveLeftOn],
  );

  const handleMoveUpClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveUpOn : enable;
      setMoveUpOn(on);
      postMsg("cameraMove", { moveType: "moveUp", on });
      if (on && moveDownOn) {
        setMoveDownOn(false);
      }
    },
    [moveUpOn, moveDownOn],
  );

  const handleMoveDownClick = useCallback(
    (enable?: boolean) => {
      const on = typeof enable !== "boolean" ? !moveDownOn : enable;
      setMoveDownOn(on);
      postMsg("cameraMove", { moveType: "moveDown", on });
      if (on && moveUpOn) {
        setMoveUpOn(false);
      }
    },
    [moveDownOn, moveUpOn],
  );

  const onExit = useCallback(() => {
    setMode("ready");
    setMainButtonText("始める");
    postMsg("pedestrianExit");
  }, []);

  const onPicking = useCallback(() => {
    setMode("picking");
    setMainButtonText("終わる");
    postMsg("pickingStart");
  }, []);

  const handlePickingDone = useCallback(() => {
    setMode("pedestrian");
    setTimeout(() => {
      miniMap.current?.invalidateSize();
    }, 500);
  }, []);

  const onMainButtonClick = useCallback(() => {
    if (mode === "ready") {
      onPicking();
    } else {
      onExit();
    }
  }, [mode, onPicking, onExit]);

  const onKeyDown = useCallback(
    (e: KeyboardEvent) => {
      if (mode !== "pedestrian") return;
      switch (e.code) {
        case "KeyW":
          handleMoveForwardClick(true);
          break;
        case "KeyA":
          handleMoveLeftClick(true);
          break;
        case "KeyS":
          handleMoveBackwardClick(true);
          break;
        case "KeyD":
          handleMoveRightClick(true);
          break;
        case "Space":
          handleMoveUpClick(true);
          break;
        case "ShiftLeft":
        case "ShiftRight":
          handleMoveDownClick(true);
          break;
        default:
          return undefined;
      }
    },
    [
      mode,
      handleMoveForwardClick,
      handleMoveBackwardClick,
      handleMoveLeftClick,
      handleMoveRightClick,
      handleMoveUpClick,
      handleMoveDownClick,
    ],
  );

  const onKeyUp = useCallback(
    (e: KeyboardEvent) => {
      if (mode !== "pedestrian") return;
      switch (e.code) {
        case "KeyW":
          handleMoveForwardClick(false);
          break;
        case "KeyA":
          handleMoveLeftClick(false);
          break;
        case "KeyS":
          handleMoveBackwardClick(false);
          break;
        case "KeyD":
          handleMoveRightClick(false);
          break;
        case "Space":
          handleMoveUpClick(false);
          break;
        case "ShiftLeft":
        case "ShiftRight":
          handleMoveDownClick(false);
          break;
        default:
          return undefined;
      }
    },
    [
      mode,
      handleMoveForwardClick,
      handleMoveBackwardClick,
      handleMoveLeftClick,
      handleMoveRightClick,
      handleMoveUpClick,
      handleMoveDownClick,
    ],
  );

  const onClose = useCallback(() => {
    if (mode !== "ready") {
      onExit();
    }
    postMsg("pedestrianClose");
  }, [mode, onExit]);

  const miniMap = useRef<LeafletMap>();
  const [miniMapViewRotate, setMiniMapViewRotate] = useState<number>();

  const initMiniMap = useCallback(() => {
    miniMap.current = L.map("minimap", {
      zoomControl: false,
      dragging: false,
      boxZoom: false,
      doubleClickZoom: false,
      keyboard: false,
      scrollWheelZoom: false,
      touchZoom: false,
      easeLinearity: 1,
    })
      .setView([0, 0], 18)
      .whenReady(() => {
        miniMap.current?.invalidateSize();
      });

    L.tileLayer("https://{s}.tile.openstreetmap.org/{z}/{x}/{y}.png", {
      attribution: '&copy; <a href="https://www.openstreetmap.org/copyright">OpenStreetMap</a>',
    }).addTo(miniMap.current);
  }, []);

  const handleUpdateMiniMap = useCallback((camera: Camera) => {
    if (miniMap.current) {
      miniMap.current.stop();
      miniMap.current.panTo([camera.lat, camera.lng], {
        animate: true,
        duration: 0.5,
        easeLinearity: 1,
      });
      setMiniMapViewRotate((camera.heading / Math.PI) * 180);
    }
  }, []);

  const handleControllerReady = useCallback(() => {
    setVisible(true);
  }, []);

  const onMessage = useCallback(
    (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      switch (e.data.type) {
        case "pickingDone":
          handlePickingDone();
          break;
        case "updateMiniMap":
          handleUpdateMiniMap(e.data.payload);
          break;
        case "controllerReady":
          handleControllerReady();
          break;
        default:
          break;
      }
    },
    [handlePickingDone, handleUpdateMiniMap, handleControllerReady],
  );

  useEffect(() => {
    document.documentElement.style.setProperty("--theme-color", "#00BEBE");
    (globalThis as any).parent.document.body.setAttribute("tabindex", "0");

    if (!miniMap.current) {
      initMiniMap();
    }

    postMsg("controllerReady");
  }, [initMiniMap]);

  useEffect(() => {
    document.addEventListener("keydown", onKeyDown, false);
    document.addEventListener("keyup", onKeyUp, false);
    (globalThis as any).parent.document.addEventListener("keydown", onKeyDown, false);
    (globalThis as any).parent.document.addEventListener("keyup", onKeyUp, false);

    return () => {
      document.removeEventListener("keydown", onKeyDown, false);
      document.removeEventListener("keyup", onKeyUp, false);
      (globalThis as any).parent.document.removeEventListener("keydown", onKeyDown);
      (globalThis as any).parent.document.removeEventListener("keyup", onKeyUp);
    };
  }, [onKeyDown, onKeyUp]);

  useEffect(() => {
    addEventListener("message", onMessage);
    return () => {
      removeEventListener("message", onMessage);
    };
  }, [onMessage]);

  return {
    mode,
    mainButtonText,
    moveForwardOn,
    moveBackwardOn,
    moveLeftOn,
    moveRightOn,
    moveUpOn,
    moveDownOn,
    miniMapViewRotate,
    visible,
    handleMoveForwardClick,
    handleMoveBackwardClick,
    handleMoveLeftClick,
    handleMoveRightClick,
    handleMoveUpClick,
    handleMoveDownClick,
    onClose,
    onMainButtonClick,
  };
};
