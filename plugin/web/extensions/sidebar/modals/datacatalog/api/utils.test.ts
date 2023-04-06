import { expect, test } from "vitest";

import { makeTree } from "./utils";

test("makeTree", () => {
  expect(
    makeTree([
      { path: ["a", "b", "c"], id: "c", type_en: "tran" },
      { path: ["d"], id: "d", type_en: "tran" },
      { path: ["a", "e"], id: "e", type_en: "tran" },
    ]),
  ).toEqual([
    {
      id: "node-0",
      desc: "",
      name: "a",
      children: [
        {
          id: "node-1",
          desc: "",
          name: "b",
          children: [
            {
              id: "node-2",
              desc: "",
              name: "c",
              item: { path: ["a", "b", "c"], id: "c", type_en: "tran" },
            },
          ],
        },
        {
          id: "node-4",
          desc: "",
          name: "e",
          item: { path: ["a", "e"], id: "e", type_en: "tran" },
        },
      ],
    },
    {
      id: "node-3",
      desc: "",
      name: "d",
      item: { path: ["d"], id: "d", type_en: "tran" },
    },
  ]);
});
