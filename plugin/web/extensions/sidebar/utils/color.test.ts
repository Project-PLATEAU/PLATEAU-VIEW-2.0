import { expect, test } from "vitest";

import { getRGBAFromString, generateColorGradient } from "./color";

test("getRGBAFromString", () => {
  expect(getRGBAFromString("rgba(100, 24, 255, 1)")).toEqual([100, 24, 255, 1]);
  expect(getRGBAFromString("rgba(100,24,255,0.5)")).toEqual([100, 24, 255, 0.5]);
});

test("generateColorGradient", () => {
  expect(generateColorGradient("#ff0000", "#0000ff", 5)).toEqual([
    "#ff0000",
    "#bf0040",
    "#800080",
    "#4000bf",
    "#0000ff",
  ]);
});
