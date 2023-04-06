export type ActionType =
  | "pedestrianShow"
  | "pedestrianClose"
  | "pickingStart"
  | "pedestrianExit"
  | "cameraMove"
  | "controllerReady";

export type PostMessageProps = { action: ActionType; payload?: any };

// reearth API
export type MouseEvent = {
  x?: number;
  y?: number;
  lat?: number;
  lng?: number;
  height?: number;
  layerId?: string;
  delta?: number;
};

export type Camera = {
  lat: number;
  lng: number;
  height: number;
  heading: number;
  pitch: number;
  roll: number;
  fov: number;
};
