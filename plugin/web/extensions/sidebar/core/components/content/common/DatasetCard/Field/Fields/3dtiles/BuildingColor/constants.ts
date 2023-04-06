import floodsImage from "./floods.webp";

export const INDEPENDENT_COLOR_TYPE: Record<
  string,
  {
    id: string;
    label: string;
    featurePropertyName: string;
    order: number;
    always?: boolean;
  }
> = {
  height: {
    id: "height",
    label: "高さによる塗分け",
    featurePropertyName: "計測高さ",
    order: 1,
  },
  purpose: {
    id: "purpose",
    label: "用途による塗分け",
    featurePropertyName: "用途",
    order: 2,
  },
  structure: {
    id: "structure",
    label: "建物構造による塗分け",
    featurePropertyName: "建物構造",
    always: true,
    order: 3,
  },
  structureType: {
    id: "structureType",
    label: "構造種別による塗分け",
    featurePropertyName: "構造種別",
    order: 4,
  },
  fireproof: {
    id: "fireproof",
    label: "耐火構造種別による塗分け",
    featurePropertyName: "耐火構造種別",
    order: 5,
  },
};

export const LAND_SLIDE_RISK_FIELD = {
  steepSlope: {
    id: "steepSlope",
    label: "急傾斜による塗分け",
  },
  mudflow: {
    id: "mudflow",
    label: "土石流による塗分け",
  },
  landslide: {
    id: "landslide",
    label: "地すべりによる塗分け",
  },
};

export const LEGEND_IMAGES: Record<"floods", string> = {
  floods: floodsImage,
};
