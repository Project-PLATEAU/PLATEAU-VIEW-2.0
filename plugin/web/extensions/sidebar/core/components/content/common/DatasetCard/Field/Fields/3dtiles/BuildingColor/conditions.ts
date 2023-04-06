import {
  compareGreaterThan,
  compareRange,
  defaultConditionalNumber,
  equalNumber,
  equalString,
  variable,
  stringOrNumber,
} from "@web/extensions/sidebar/utils";

import { BuildingColor } from "../types";

import { INDEPENDENT_COLOR_TYPE, LAND_SLIDE_RISK_FIELD } from "./constants";

type Condition = {
  condition: string;
  color: BuildingColor;
  label: string;
  default?: boolean;
};

const DEFAULT_CONDITION: Condition = {
  condition: "true",
  color: "rgba(255, 255, 255, 1)",
  label: "",
  default: true,
};

const conditionalHeight = defaultConditionalNumber("計測高さ");
const HEIGHT_CONDITIONS: Condition[] = [
  {
    condition: compareGreaterThan(conditionalHeight, 180),
    color: "rgba(247, 255, 0, 1)",
    label: "180",
  },
  {
    condition: compareRange(conditionalHeight, [120, 180]),
    color: "rgba(255, 205, 0, 1)",
    label: "120",
  },
  {
    condition: compareRange(conditionalHeight, [60, 120]),
    color: "rgba(240, 211, 123, 1)",
    label: "60",
  },
  {
    condition: compareRange(conditionalHeight, [31, 60]),
    color: "rgba(166, 117, 190, 1)",
    label: "31",
  },
  {
    condition: compareRange(conditionalHeight, [12, 31]),
    color: "rgba(90, 34, 200, 1)",
    label: "12",
  },
  { condition: compareRange(conditionalHeight, [0, 12]), color: "rgba(56, 42, 84, 1)", label: "0" },
  DEFAULT_CONDITION,
];

const conditionalPurpose = "用途";
const PURPOSE_CONDITIONS: Condition[] = [
  {
    condition: equalString(conditionalPurpose, "業務施設"),
    color: "rgba(255, 127, 80, 1)",
    label: "業務施設",
  },
  {
    condition: equalString(conditionalPurpose, "商業施設"),
    color: "rgba(255, 69, 0, 1)",
    label: "商業施設",
  },
  {
    condition: equalString(conditionalPurpose, "宿泊施設"),
    color: "rgba(255, 255, 0, 1)",
    label: "宿泊施設",
  },
  {
    condition: equalString(conditionalPurpose, "商業系複合施設"),
    color: "rgba(255, 69, 0, 1)",
    label: "商業系複合施設",
  },
  {
    condition: equalString(conditionalPurpose, "住宅"),
    color: "rgba(50, 205, 50, 1)",
    label: "住宅",
  },
  {
    condition: equalString(conditionalPurpose, "共同住宅"),
    color: "rgba(0, 255, 127, 1)",
    label: "共同住宅",
  },
  {
    condition: equalString(conditionalPurpose, "店舗等併用住宅"),
    color: "rgba(0, 255, 255, 1)",
    label: "店舗等併用住宅",
  },
  {
    condition: equalString(conditionalPurpose, "店舗等併用共同住宅"),
    color: "rgba(0, 255, 255, 1)",
    label: "店舗等併用共同住宅",
  },
  {
    condition: equalString(conditionalPurpose, "作業所併用住宅"),
    color: "rgba(0, 255, 255, 1)",
    label: "作業所併用住宅",
  },
  {
    condition: equalString(conditionalPurpose, "官公庁施設"),
    color: "rgba(65, 105, 225, 1)",
    label: "官公庁施設",
  },
  {
    condition: equalString(conditionalPurpose, "文教厚生施設"),
    color: "rgba(0, 0, 255, 1)",
    label: "文教厚生施設",
  },
  {
    condition: equalString(conditionalPurpose, "運輸倉庫施設"),
    color: "rgba(147, 112, 219, 1)",
    label: "運輸倉庫施設",
  },
  {
    condition: equalString(conditionalPurpose, "工場"),
    color: "rgba(135, 206, 250, 1)",
    label: "工場",
  },
  {
    condition: equalString(conditionalPurpose, "農林漁業用施設"),
    color: "rgba(0, 128, 0, 1)",
    label: "農林漁業用施設",
  },
  {
    condition: equalString(conditionalPurpose, "併給処理施設"),
    color: "rgba(139, 69, 19, 1)",
    label: "併給処理施設",
  },
  {
    condition: equalString(conditionalPurpose, "防衛施設"),
    color: "rgba(178, 34, 34, 1)",
    label: "防衛施設",
  },
  {
    condition: equalString(conditionalPurpose, "その他"),
    color: "rgba(216, 191, 216, 1)",
    label: "その他",
  },
  {
    condition: equalString(conditionalPurpose, "不明"),
    color: "rgba(230, 230, 250, 1)",
    label: "不明",
  },
  DEFAULT_CONDITION,
];

const createStructureCondition = (value: string) => {
  return `${variable(
    `$.attributes["uro:KeyValuePairAttribute"][?(@["uro:key"]=="建物構造コード"&&@["uro:codeValue"]==${stringOrNumber(
      value,
    )})]["uro:codeValue"]`,
  )} === ${stringOrNumber(value)}`;
};

const STRUCTURE_CONDITIONS: Condition[] = [
  {
    condition: createStructureCondition("耐火構造"),
    color: "rgba(124, 123, 135, 1)",
    label: "耐火構造",
  },
  {
    condition: createStructureCondition("防火造"),
    color: "rgba(188, 143, 143, 1)",
    label: "防火造",
  },
  {
    condition: createStructureCondition("準防火造"),
    color: "rgba(214, 202, 174, 1)",
    label: "準防火造",
  },
  {
    condition: createStructureCondition("木造"),
    color: "rgba(210, 180, 140, 1)",
    label: "木造",
  },
  DEFAULT_CONDITION,
];

const conditionalStructureType = "構造種別";
const STRUCTURE_TYPE_CONDITIONS: Condition[] = [
  {
    condition: equalString(conditionalStructureType, "木造・土蔵造"),
    color: "rgba(178, 180, 140, 1)",
    label: "木造・土蔵造",
  },
  {
    condition: equalString(conditionalStructureType, "鉄骨鉄筋コンクリート造"),
    color: "rgba(229, 225, 64, 1)",
    label: "鉄骨鉄筋コンクリート造",
  },
  {
    condition: equalString(conditionalStructureType, "鉄筋コンクリート造"),
    color: "rgba(234, 164, 37, 1)",
    label: "鉄筋コンクリート造",
  },
  {
    condition: equalString(conditionalStructureType, "鉄骨造"),
    color: "rgba(153, 99, 50, 1)",
    label: "鉄骨造",
  },
  {
    condition: equalString(conditionalStructureType, "軽量鉄骨造"),
    color: "rgba(160, 79, 146, 1)",
    label: "軽量鉄骨造",
  },
  {
    condition: equalString(conditionalStructureType, "レンガ造・コンクリートブロック造・石造"),
    color: "rgba(119, 23, 28, 1)",
    label: "レンガ造・コンクリートブロック造・石造",
  },
  {
    condition: equalString(conditionalStructureType, "非木造"),
    color: "rgba(137, 182, 220, 1)",
    label: "非木造",
  },
  {
    condition: equalString(conditionalStructureType, "耐火"),
    color: "rgba(127, 123, 133, 1)",
    label: "耐火",
  },
  {
    condition: equalString(conditionalStructureType, "簡易耐火"),
    color: "rgba(140, 155, 177, 1)",
    label: "簡易耐火",
  },
  {
    condition: equalString(conditionalStructureType, "不明"),
    color: "rgba(34, 34, 34, 1)",
    label: "不明",
  },
  DEFAULT_CONDITION,
];

const conditionalFireproof = "耐火構造種別";
const FIREPROOF_CONDITIONS: Condition[] = [
  {
    condition: equalString(conditionalFireproof, "耐火"),
    color: "rgba(127, 123, 133, 1)",
    label: "耐火",
  },
  {
    condition: equalString(conditionalFireproof, "準耐火造"),
    color: "rgba(140, 155, 177, 1)",
    label: "準耐火造",
  },
  {
    condition: equalString(conditionalFireproof, "その他"),
    color: "rgba(250, 131, 158, 1)",
    label: "その他",
  },
  {
    condition: equalString(conditionalFireproof, "不明"),
    color: "rgba(120, 194, 243, 1)",
    label: "不明",
  },
  DEFAULT_CONDITION,
];

const LAND_SLIDE_RISK_CODES = [
  "1", // 急傾斜地の崩壊
  "2", // 土石流
  "3", // 地すべり
];
const LAND_SLIDE_RISK_TYPE_CODES = [
  "1", // 警戒区域
  "2", // 特別警戒区域
  "3", // 警戒区域(指定前)
  "4", // 特別警戒区域(指定前)
];

// The code of land slide risk has 3 types, and `BuildingLandSlideRiskAttribute` property has array, so we need to check 3 items.
const makeLandSlideRiskCondition = (
  propertyKey: string,
  riskCode: string,
  riskTypeCodes: string[],
) =>
  LAND_SLIDE_RISK_CODES.reduce((res, _code, i) => {
    const next = res ? `(${res}) || ` : "";
    return `${next}${variable(propertyKey)} !== undefined && ${variable(
      `${propertyKey}[${i}]`,
    )} !== undefined && ${equalString(
      `${propertyKey}[${i}]["uro:description_code"]`,
      riskCode,
    )} && (${riskTypeCodes
      .map(code => equalString(`${propertyKey}[${i}]["uro:areaType_code"]`, code))
      .join("||")})`;
  }, "");

const conditionalLandSlideRisk = "attributes['uro:BuildingLandSlideRiskAttribute']";
const STEEP_SLOPE_RISK_CONDITIONS: Condition[] = [
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[0], [
      LAND_SLIDE_RISK_TYPE_CODES[0],
      // LAND_SLIDE_RISK_TYPE_CODES[2],
    ]),
    color: "rgba(255, 237, 76, 1)",
    label: "急傾斜地の崩落: 警戒区域",
  },
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[0], [
      LAND_SLIDE_RISK_TYPE_CODES[1],
      // LAND_SLIDE_RISK_TYPE_CODES[3],
    ]),
    color: "rgba(251, 104, 76, 1)",
    label: "急傾斜地の崩落: 特別警戒区域",
  },
  DEFAULT_CONDITION,
];
const MUDFLOW_RISK_CONDITIONS: Condition[] = [
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[1], [
      LAND_SLIDE_RISK_TYPE_CODES[0],
      // LAND_SLIDE_RISK_TYPE_CODES[2],
    ]),
    color: "rgba(237, 216, 111, 1)",
    label: "土石流: 警戒区域",
  },
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[1], [
      LAND_SLIDE_RISK_TYPE_CODES[1],
      // LAND_SLIDE_RISK_TYPE_CODES[3],
    ]),
    color: "rgba(192, 76, 99, 1)",
    label: "土石流: 特別警戒区域",
  },
  DEFAULT_CONDITION,
];
const LANDSLIDE_RISK_CONDITIONS: Condition[] = [
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[2], [
      LAND_SLIDE_RISK_TYPE_CODES[0],
      // LAND_SLIDE_RISK_TYPE_CODES[2],
    ]),
    color: "rgba(255, 183, 76, 1)",
    label: "地すべり: 警戒区域",
  },
  {
    condition: makeLandSlideRiskCondition(conditionalLandSlideRisk, LAND_SLIDE_RISK_CODES[2], [
      LAND_SLIDE_RISK_TYPE_CODES[1],
      // LAND_SLIDE_RISK_TYPE_CODES[3],
    ]),
    color: "rgba(202, 76, 149, 1)",
    label: "地すべり: 特別警戒区域",
  },
  DEFAULT_CONDITION,
];

const createFloodCondition = (
  featurePropertyName: string,
  { rank, scale }: { rank: number; scale: number | undefined },
  useOwnData: boolean | undefined,
  color: string,
): [condition: string, color: string][] => {
  if (!useOwnData && featurePropertyName) {
    return [[equalNumber(featurePropertyName, rank), color]];
  }
  const createJSONPath = (useCode: boolean, rankAsNumber: boolean) => {
    const rankProperty = useCode ? `uro:rankOrg_code` : "uro:rankOrg";
    const scaleProperty = "uro:scale_code";
    const convertedRank = rankAsNumber ? rank : rank.toString();
    return `${variable(
      `$.attributes["uro:BuildingRiverFloodingRiskAttribute"][?(@["uro:description"]==${stringOrNumber(
        featurePropertyName,
      )}&&@["${rankProperty}"]==${stringOrNumber(convertedRank)}${
        scale ? `&&@["${scaleProperty}"]==${stringOrNumber(scale.toString())}` : ""
      })]["${rankProperty}"]`,
    )} === ${stringOrNumber(convertedRank)}`;
  };
  return [
    [createJSONPath(true, false), color],
    [createJSONPath(true, true), color],
    [createJSONPath(false, false), color],
  ];
};

export const makeSelectedFloodCondition = ({
  featurePropertyName,
  useOwnData,
  floodScale,
}: {
  featurePropertyName?: string;
  useOwnData?: boolean;
  floodScale?: number;
}): [condition: string, color: string][] | undefined =>
  featurePropertyName
    ? [
        ...createFloodCondition(
          featurePropertyName,
          { rank: 0, scale: floodScale },
          useOwnData,
          "rgba(243, 240, 122, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 1, scale: floodScale },
          useOwnData,
          "rgba(243, 240, 122, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 2, scale: floodScale },
          useOwnData,
          "rgba(255, 184, 141, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 3, scale: floodScale },
          useOwnData,
          "rgba(255, 132, 132, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 4, scale: floodScale },
          useOwnData,
          "rgba(255, 94, 94, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 5, scale: floodScale },
          useOwnData,
          "rgba(237, 87, 181, 1)",
        ),
        ...createFloodCondition(
          featurePropertyName,
          { rank: 6, scale: floodScale },
          useOwnData,
          "rgba(209, 82, 209, 1)",
        ),
        [DEFAULT_CONDITION.condition, DEFAULT_CONDITION.color],
      ]
    : [[DEFAULT_CONDITION.condition, DEFAULT_CONDITION.color]];

export const COLOR_TYPE_CONDITIONS: {
  [K in
    | keyof typeof INDEPENDENT_COLOR_TYPE
    | "none"
    | keyof typeof LAND_SLIDE_RISK_FIELD]: Condition[];
} = {
  none: [DEFAULT_CONDITION],
  height: HEIGHT_CONDITIONS,
  purpose: PURPOSE_CONDITIONS,
  structure: STRUCTURE_CONDITIONS,
  structureType: STRUCTURE_TYPE_CONDITIONS,
  fireproof: FIREPROOF_CONDITIONS,
  steepSlope: STEEP_SLOPE_RISK_CONDITIONS,
  mudflow: MUDFLOW_RISK_CONDITIONS,
  landslide: LANDSLIDE_RISK_CONDITIONS,
};
