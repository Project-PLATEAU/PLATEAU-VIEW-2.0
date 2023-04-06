import { expect, test } from "vitest";

import { getLocationNamesFromFeatureProperties } from "./csv";

test("getLocationNamesFromFeatureProperties", () => {
  expect(
    getLocationNamesFromFeatureProperties({ properties: { lng: 1, lat: 2, height: 3 } }),
  ).toEqual({
    lng: "lng",
    lat: "lat",
    height: "height",
  });

  expect(getLocationNamesFromFeatureProperties({ properties: { lng: 1, lat: 2 } })).toEqual({
    lng: "lng",
    lat: "lat",
  });

  expect(
    getLocationNamesFromFeatureProperties({ properties: { lng: 1, latitude: 2, lat: 3 } }),
  ).toEqual({
    lng: "lng",
    lat: "latitude",
  });
});
