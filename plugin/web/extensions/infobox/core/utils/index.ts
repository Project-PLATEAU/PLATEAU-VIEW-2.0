import { ActionType } from "../../types";

export { getAttributes, getRootFields, commonPropertiesMap } from "./attributes";

export function postMsg(action: ActionType, payload?: any) {
  if (parent === window) return;
  parent.postMessage({
    action,
    payload,
  });
}

export const cesium3DTilesAppearanceKeys: string[] = [
  "tileset",
  "show",
  "color",
  "pointSize",
  "styleUrl",
  "shadows",
  "colorBlendMode",
  "edgeWidth",
  "edgeColor",
  "experimental_clipping",
];
