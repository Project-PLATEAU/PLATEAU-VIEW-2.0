import { expect, test } from "vitest";

import { proxyGTFS } from "./proxy";

test("proxyGTFS", () => {
  expect(
    proxyGTFS({
      data: { url: "data.example.com" },
      marker: { imageColor: "red" },
      model: { url: "example.com" },
    }),
  ).toEqual({
    data: { url: "data.example.com" },
    marker: { imageColor: "red", near: 1000, clampToGround: true },
    model: { url: "example.com" },
  });
  expect(
    proxyGTFS({
      data: { url: "data.example.com" },
    }),
  ).toEqual({
    data: { url: "data.example.com" },
  });
  expect(proxyGTFS({})).toEqual({});
});
