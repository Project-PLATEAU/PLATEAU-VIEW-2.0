import { cloneDeep, mergeWith } from "lodash";

import { ActionType } from "../../types";

export function postMsg(action: ActionType, payload?: any) {
  if (parent === window) return;
  parent.postMessage({
    action,
    payload,
  });
}

export function mergeProperty(a: any, b: any) {
  const a2 = cloneDeep(a);
  return mergeWith(
    a2,
    b,
    (s: any, v: any, _k: string | number | symbol, _obj: any, _src: any, stack: { size: number }) =>
      stack.size > 0 || Array.isArray(v) ? v ?? s : undefined,
  );
}

export function generateId() {
  return "xxxxxxxxxxxxxxxxxxxxxxxxxx".replace(/[x]/g, function () {
    return ((Math.random() * 16) | 0).toString(16);
  });
}
