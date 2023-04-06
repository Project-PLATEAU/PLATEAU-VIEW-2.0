import { expect, test } from "vitest";

import { getExtension, createFileName, normalizeExtension } from "./file";

test("getExtension", () => {
  expect(getExtension("test.geojson")).toBe("geojson");
  expect(getExtension("test.")).toBe("");
  expect(getExtension("test")).toBe("");
});

test("createFileName", () => {
  expect(createFileName("test", "geojson")).toBe("test.geojson");
  expect(createFileName("", ".czml")).toBe("");
  expect(createFileName("test", "")).toBe("");
  expect(createFileName("", "")).toBe("");
});

test("normalizeExtension", () => {
  expect(normalizeExtension(" CZML ")).toBe("czml");
  expect(normalizeExtension("CZML")).toBe("czml");
  expect(normalizeExtension("")).toBe("");
  expect(normalizeExtension()).toBe("");
});
