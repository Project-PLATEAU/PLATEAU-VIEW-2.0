import { expect, test } from "vitest";

import { getAttributes } from ".";

test("getAttributes fld", () => {
  expect(
    getAttributes(
      {
        "gml:name": "xxx川水系yyy川（計画規模）",
        "wtr:class": "Flooded land",
        "wtr:class_code": "1140",
        "wtr:class_codeSpace": "../../../../codelists/WaterBody_class.xml",
        "wtr:function": "洪水浸水想定区域",
        "wtr:function_code": "1",
        "wtr:function_codeSpace": "../../../../codelists/WaterBody_function.xml",
        "uro:WaterBodyRiverFloodingRiskAttribute": {
          "uro:description": "xxx水系yyy川",
          "uro:description_code": "1",
          "uro:description_codeSpace":
            "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_description.xml",
          "uro:rank": "0.5m以上3m未満",
          "uro:rank_code": "2",
          "uro:rank_codeSpace":
            "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_rank.xml",
          "uro:adminType": "国",
          "uro:adminType_code": "1",
          "uro:adminType_codeSpace":
            "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_adminType.xml",
          "uro:scale": "L1（計画規模）",
          "uro:scale_code": "L1",
          "uro:scale_codeSpace":
            "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_scale.xml",
        },
      },
      "label",
    ),
  ).toEqual({
    名称: "xxx川水系yyy川（計画規模）",
    分類: "Flooded land",
    分類コード: "1140",
    // "wtr:class_codeSpace": "../../../../codelists/WaterBody_class.xml",
    機能: "洪水浸水想定区域",
    機能コード: "1",
    // "wtr:function_codeSpace": "../../../../codelists/WaterBody_function.xml",
    洪水浸水想定区域: {
      設定等名称: "xxx水系yyy川",
      設定等名称コード: "1",
      // "uro:description_codeSpace": "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_description.xml",
      浸水ランク: "0.5m以上3m未満",
      浸水ランクコード: "2",
      // "uro:rank_codeSpace": "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_rank.xml",
      指定機関: "国",
      指定機関コード: "1",
      // "uro:adminType_codeSpace":  "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_adminType.xml",
      規模: "L1（計画規模）",
      規模コード: "L1",
      // "uro:scale_codeSpace":  "../../../../codelists/WaterBodyRiverFloodingRiskAttribute_scale.xml",
    },
  });
});
