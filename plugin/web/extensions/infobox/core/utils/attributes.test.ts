import { expect, test, vi } from "vitest";

import {
  attributesMap,
  getAttributes,
  getRootFields,
  name,
  fldName,
  constructionYear,
} from "./attributes";
import type { Json } from "./json";

test("getAttributes", () => {
  const src: Json = {
    bbb: {},
    aaa: {
      bbb: "ccc",
      ddd: [{ c: "b" }, { ddd_code: "a", a_code: "" }],
    },
  };
  expect(flatKeys(src)).toEqual([
    "",
    "bbb",
    "aaa",
    "aaa.bbb",
    "aaa.ddd",
    "aaa.ddd.0",
    "aaa.ddd.0.c",
    "aaa.ddd.1",
    "aaa.ddd.1.ddd_code",
    "aaa.ddd.1.a_code",
  ]);

  const actual = getAttributes(src);
  expect(flatKeys(actual)).toEqual([
    "",
    "aaa",
    "aaa.bbb",
    "aaa.ddd",
    "aaa.ddd.0",
    "aaa.ddd.0.c",
    "aaa.ddd.1",
    "aaa.ddd.1.a_code",
    "aaa.ddd.1.ddd_code",
    "bbb",
  ]);

  const actual1 = getAttributes(src, "label");
  expect(flatKeys(actual1)).toEqual([
    "",
    "AAA",
    "AAA.bbb",
    "AAA.DDD",
    "AAA.DDD.0",
    "AAA.DDD.0.c",
    "AAA.DDD.1",
    "AAA.DDD.1.a_code",
    "AAA.DDD.1.DDDコード",
    "bbb",
  ]);

  const actual2 = getAttributes(src, "both");
  expect(flatKeys(actual2)).toEqual([
    "",
    "AAA（aaa）",
    "AAA（aaa）.bbb",
    "AAA（aaa）.DDD（ddd）",
    "AAA（aaa）.DDD（ddd）.0",
    "AAA（aaa）.DDD（ddd）.0.c",
    "AAA（aaa）.DDD（ddd）.1",
    "AAA（aaa）.DDD（ddd）.1.a_code",
    "AAA（aaa）.DDD（ddd）.1.DDDコード（ddd_code）",
    "bbb",
  ]);
});

test("getRootFields bldg", () => {
  const res = getRootFields({
    attributes: {
      "gml:id": "id",
      "bldg:class": "堅ろう建物",
      "bldg:yearOfConstruction": 2008,
      "bldg:usage": ["共同住宅"],
      "bldg:measuredHeight": 20,
      "bldg:measuredHeightUmo": "m",
      "bldg:storeysAboveGround": 12,
      "bldg:storeysBelowGround": 0,
      "uro:majorUsage": "建物利用現況（大分類）",
      "uro:orgUsage": "建物利用現況（中分類）",
      "uro:orgUsage2": "建物利用現況（小分類）",
      "uro:detailedUsage": "建物利用現況（詳細分類）",
      "uro:BuildingDetailAttribute": [
        {
          "uro:siteArea": 100,
          "uro:buildingStructureType": "鉄筋コンクリート造",
          "uro:buildingStructureOrgType": "鉄筋コンクリート造",
          "uro:fireproofStructureType": "不明",
          "uro:surveyYear": 2018,
          "uro:urbanPlanType": "都市計画区域",
          "uro:areaClassificationType": "区域区分",
          "uro:districtsAndZonesType": ["地域地区"],
          "uro:buildingRoofEdgeArea": 6399.9406,
          "uro:totalFloorArea": 1000,
        },
      ],
      "uro:BuildingRiverFloodingRiskAttribute": [
        {
          "uro:description": "六角川水系武雄川",
          "uro:depth": 0.618,
          "uro:depth_uom": "m",
          "uro:adminType": "国",
          "uro:scale": "L2（想定最大規模）",
          "uro:rank_code": "1",
          "uro:duration": "継続時間",
        },
        {
          "uro:description": "六角川水系武雄川",
          "uro:depth": 0,
          "uro:depth_uom": "m",
          "uro:adminType": "都道府県",
          "uro:scale": "L1（計画規模）",
          "uro:duration": "継続時間",
          "uro:rankOrg": "2",
        },
      ],
      "uro:BuildingLandSlideRiskAttribute": [
        {
          "uro:description": "急傾斜地の崩落",
          "uro:description_code": "1",
          "uro:areaType": "土砂災害警戒区域（指定済）",
          "uro:areaType_code": "1",
        },
      ],
      "uro:BuildingIDAttribute": [{ "uro:buildingID": "22213-bldg-26889" }],
      "uro:BuildingDataQualityAttribute": [{ "uro:lod1HeightType": "航空写真図化_最高高さ" }],
    },
  });

  expect(res).toEqual({
    gml_id: "id",
    分類: "堅ろう建物",
    用途: "共同住宅",
    建築年: 2008,
    計測高さ: 20,
    地上階数: 12,
    地下階数: 0,
    敷地面積: 100,
    延床面積: 1000,
    構造種別: "鉄筋コンクリート造",
    "構造種別（独自）": "鉄筋コンクリート造",
    耐火構造種別: "不明",
    都市計画区域: "都市計画区域",
    区域区分: "区域区分",
    地域地区: "地域地区",
    調査年: 2018,
    "建物利用現況（大分類）": "建物利用現況（大分類）",
    "建物利用現況（中分類）": "建物利用現況（中分類）",
    "建物利用現況（小分類）": "建物利用現況（小分類）",
    "建物利用現況（詳細分類）": "建物利用現況（詳細分類）",
    建物ID: "22213-bldg-26889",
    図上面積: 6399.9406,
    LOD1立ち上げに使用する高さ: "航空写真図化_最高高さ",
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_浸水ランク": "1",
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_浸水深": 0.618,
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_継続時間": "継続時間",
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_浸水ランク": "2",
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_浸水深": 0,
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_継続時間": "継続時間",
    土砂災害警戒区域: "急傾斜地の崩落",
  });

  expect(flatKeys(res)).toEqual([
    "",
    "gml_id",
    "分類",
    "用途",
    "建築年",
    "計測高さ",
    "地上階数",
    "地下階数",
    "敷地面積",
    "延床面積",
    "構造種別",
    "構造種別（独自）",
    "耐火構造種別",
    "都市計画区域",
    "区域区分",
    "地域地区",
    "調査年",
    "建物利用現況（大分類）",
    "建物利用現況（中分類）",
    "建物利用現況（小分類）",
    "建物利用現況（詳細分類）",
    "建物ID",
    "図上面積",
    "LOD1立ち上げに使用する高さ",
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_浸水ランク",
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_浸水深",
    "六角川水系武雄川（国管理区間）_L2（想定最大規模）_継続時間",
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_浸水ランク",
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_浸水深",
    "六角川水系武雄川（都道府県管理区間）_L1（計画規模）_継続時間",
    "土砂災害警戒区域",
  ]);
});

test("getRootFields luse", () => {
  const res = getRootFields({
    attributes: {
      "gml:id": "id",
      "luse:class": "山林（樹林地）",
      "uro:LandUseDetailAttribute": [
        {
          "uro:orgLandUse": "山林",
          "uro:city": "広島県福山市",
          "uro:surfeyYear": 2000, // sufey is an invalid attribute
        },
      ],
    },
  });

  expect(res).toEqual({
    gml_id: "id",
    分類: "山林（樹林地）",
    都市名: "広島県福山市",
    調査年: 2000,
  });

  expect(flatKeys(res)).toEqual(["", "gml_id", "分類", "都市名", "調査年"]);
});

test("attributesMap", () => {
  expect(attributesMap.get("ddd")).toBe("DDD");
});

test("name", () => {
  expect(name({ attributes: { "gml:name": "aaaaaa" } }, "htd", "bbbbbb")).toEqual({
    name: "aaaaaa",
  });
  expect(name({ attributes: { "gml:name": "01" } }, "bldg", "bbbbbb")).toEqual({ 名称: "01" });
  expect(
    name(
      { attributes: { "gml:name": "01" } },
      "htd",
      "高潮浸水想定区域モデル 有明海沿岸（小城市）",
    ),
  ).toEqual({ name: "有明海沿岸高潮浸水想定区域図" });
  expect(
    name(
      { attributes: { "gml:name": "" } },
      "tnm",
      "津波浸水想定区域モデル 宮城県津波浸水想定図（仙台市）",
    ),
  ).toBe(undefined);
});

test("fldName", () => {
  expect(
    fldName(
      "洪水浸水想定区域モデル 芦田川水系芦田川（国管理区間）（福山市）",
      "fld",
      "想定最大規模",
    ),
  ).toBe("芦田川水系芦田川洪水浸水想定区域図【想定最大規模】");

  expect(fldName("高潮浸水想定区域モデル 有明海沿岸（小城市）", "htd")).toBe(
    "有明海沿岸高潮浸水想定区域図",
  );

  expect(fldName("津波浸水想定区域モデル 宮城県津波浸水想定図（仙台市）", "tnm")).toBe(
    "宮城県津波浸水想定図",
  );

  expect(fldName("津波浸水想定区域モデル 津波浸水想定（福山市）", "tnm")).toBe("津波浸水想定");
});

function flatKeys(obj: Json, parentKey?: string): string[] {
  if (typeof obj !== "object" || !obj) return [parentKey || ""];
  return [
    parentKey || "",
    ...Object.entries(obj).flatMap(([k, v]) =>
      flatKeys(v, `${parentKey ? `${parentKey}.` : ""}${k}`),
    ),
  ];
}

test("constructionYear", () => {
  expect(constructionYear("")).toBe(undefined);
  expect(constructionYear(null)).toBe(undefined);
  expect(constructionYear(undefined)).toBe(undefined);
  expect(constructionYear(-1)).toBe(-1);
  expect(constructionYear(0)).toBe(0);
  expect(constructionYear(1)).toBe(1);
  expect(constructionYear("0")).toBe("0");
  expect(constructionYear("1")).toBe("1");
  expect(constructionYear("0000")).toBe("0000");
  expect(constructionYear("0001")).toBe("0001");
  expect(constructionYear(2001)).toBe(2001);
  expect(constructionYear("2001")).toBe("2001");
  expect(constructionYear(2)).toBe(2);
  expect(constructionYear("", "bldg")).toBe("不明");
  expect(constructionYear(null, "bldg")).toBe("不明");
  expect(constructionYear(undefined, "bldg")).toBe("不明");
  expect(constructionYear(-1, "bldg")).toBe("不明");
  expect(constructionYear(0, "bldg")).toBe("不明");
  expect(constructionYear(1, "bldg")).toBe("不明");
  expect(constructionYear("0", "bldg")).toBe("不明");
  expect(constructionYear("1", "bldg")).toBe("不明");
  expect(constructionYear("0000", "bldg")).toBe("不明");
  expect(constructionYear("0001", "bldg")).toBe("不明");
  expect(constructionYear(2001, "bldg")).toBe(2001);
  expect(constructionYear("2001", "bldg")).toBe("2001");
  expect(constructionYear(2, "bldg")).toBe(2);
});

vi.mock("./attributes.csv?raw", () => ({
  default: "ddd,DDD\naaa,AAA\n_code,コード\n",
}));
