import { expect, test } from "vitest";

import { fileName } from "./utils";

test("fileName", () => {
  expect(fileName("https://example.com/aa/bb.zip")).toBe("bb.zip");
  expect(fileName("/bb.zip")).toBe("bb.zip");
  expect(fileName("")).toBe("");
});
