export type MouseEvent = {
  lat?: number;
  lng?: number;
};

export type DistanceLegend = {
  label?: string;
  unitLine?: number;
};
type actionType = "initLocation" | "googleModalOpen" | "terrainModalOpen" | "modalClose";

export type PostMessageProps = { action: actionType; payload?: any };
