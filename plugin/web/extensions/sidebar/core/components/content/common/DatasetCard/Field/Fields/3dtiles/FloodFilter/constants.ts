export type FilteringField = {
  value?: [from: number, to: number];
  min?: number;
  max?: number;
  isOrg?: boolean;
};

export const FEATURE_PROPERTY_NAME_RANK_CODE = "rank_code";
export const FEATURE_PROPERTY_NAME_RANK_ORG_CODE = "rank_org_code";
