import { cloneDeep, mergeWith } from "lodash";

import { Cond } from "../core/components/content/common/DatasetCard/Field/Fields/types";

export * from "./array";
export * from "./color";
export * from "./dataset";
export * from "./overrides";
export * from "./postMessage";

export function mergeProperty(a: any, b: any) {
  const a2 = cloneDeep(a);
  return mergeWith(
    a2,
    b,
    (s: any, v: any, _k: string | number | symbol, _obj: any, _src: any, stack: { size: number }) =>
      stack.size > 0 || Array.isArray(v) ? v ?? s : undefined,
  );
}

export function generateID() {
  return Date.now().toString(36) + Math.random().toString(16).slice(2);
}

export const checkKeyPress = (e: React.MouseEvent<HTMLButtonElement>, keys: string[]) => {
  let keyPressed = false;
  keys.forEach(k => {
    if (e[`${k}Key` as keyof typeof e]) {
      keyPressed = true;
    }
  });
  return keyPressed;
};

export const updateExtended = (e: { vertically: boolean }) => {
  const html = document.querySelector("html");
  const body = document.querySelector("body");
  const root = document.getElementById("root");

  if (e?.vertically) {
    html?.classList.add("extended");
    body?.classList.add("extended");
    root?.classList.add("extended");
  } else {
    html?.classList.remove("extended");
    body?.classList.remove("extended");
    root?.classList.remove("extended");
  }
};

export function stringifyCondition(condition: Cond<any>): string {
  return String(condition.operand) + String(condition.operator) + String(condition.value);
}

export const defaultConditionalNumber = (prop: string, defaultValue?: number) =>
  `((${variable(prop)} === "" || ${variable(prop)} === null || isNaN(Number(${variable(
    prop,
  )}))) ? ${defaultValue || 1} : Number(${variable(prop)}))`;

export const compareRange = (conditionalValue: string, range: [from: number, to: number]) =>
  `(${conditionalValue} >= ${range?.[0]} && ${conditionalValue} <= ${range?.[1]})`;

export const compareGreaterThan = (conditionalValue: string, num: number) =>
  `(${conditionalValue} >= ${num})`;

export const equalString = (prop: string, value: string) => `(${variable(prop)} === "${value}")`;

export const equalNumber = (prop: string, value: number) => `(${variable(prop)} === ${value})`;

export const stringOrNumber = (v: string | number) =>
  typeof v === "number" ? v.toString() : `"${v}"`;

export const variable = (prop: string) => `\${${prop}}`;
