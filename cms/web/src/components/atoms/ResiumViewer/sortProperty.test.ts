import { test, expect, describe } from "vitest";

import { sortProperties } from "./sortProperty";

describe("sort Properties", () => {
  test("sort depth1: simple", () => {
    const obj = {
      a: "a",
      c: "c",
      b: "b",
    };
    const expected = {
      a: "a",
      b: "b",
      c: "c",
    };
    expect(sortProperties(obj)).toEqual(expected);
  });
  test("sort depth1: other data types", () => {
    const obj = {
      a: ["hoge", "fuga"],
      c: "c",
      b: 1000,
    };
    const expected = {
      a: ["hoge", "fuga"],
      b: 1000,
      c: "c",
    };
    expect(sortProperties(obj)).toEqual(expected);
  });

  test("sort depth2: simple", () => {
    const obj = {
      a: {
        a: "a",
        c: "c",
        b: "b",
      },
      c: ["1", "2", "3"],
      d: 100,
      b: {
        a: "a",
        c: "c",
        b: "b",
      },
    };
    const expected = {
      a: {
        a: "a",
        b: "b",
        c: "c",
      },
      b: {
        a: "a",
        b: "b",
        c: "c",
      },
      d: 100,
      c: ["1", "2", "3"],
    };
    expect(sortProperties(obj)).toStrictEqual(expected);
  });
});
