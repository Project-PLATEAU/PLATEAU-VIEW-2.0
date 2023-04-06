import { equalNumber, equalString } from "@web/extensions/sidebar/utils";

import { BaseFieldProps } from "../../types";
import { BuildingColor } from "../types";

type Condition = [condition: string, color: BuildingColor];

const DEFAULT_CONDITION: Condition = ["true", "rgba(209, 82, 209, 1)"];
export const DEFAULT_TRANSPARENCY = 1;

const conditionalFloodRank = "rank_code";
const conditionalFloodRankOrgCode = "rank_org_code";
const conditionalFloodRankOrg = "rank_org";
const FLOOD_RANK_CONDITIONS: Condition[] = [
  [
    `${equalNumber(conditionalFloodRankOrgCode, 1)} || ${equalString(
      conditionalFloodRankOrg,
      "1",
    )} || ${equalNumber(conditionalFloodRank, 1)}`,
    `rgba(255, 247, 75, ${DEFAULT_TRANSPARENCY})`,
  ],
  [
    `${equalNumber(conditionalFloodRankOrgCode, 2)} || ${equalString(
      conditionalFloodRankOrg,
      "2",
    )} || ${equalNumber(conditionalFloodRank, 2)}`,
    `rgba(255, 157, 84, ${DEFAULT_TRANSPARENCY})`,
  ],
  [
    `${equalNumber(conditionalFloodRankOrgCode, 3)} || ${equalString(
      conditionalFloodRankOrg,
      "3",
    )} || ${equalNumber(conditionalFloodRank, 3)}`,
    `rgba(255, 95, 89, ${DEFAULT_TRANSPARENCY})`,
  ],
  [
    `${equalNumber(conditionalFloodRankOrgCode, 4)} || ${equalString(
      conditionalFloodRankOrg,
      "4",
    )} || ${equalNumber(conditionalFloodRank, 4)}`,
    `rgba(255, 61, 50, ${DEFAULT_TRANSPARENCY})`,
  ],
  [
    `${equalNumber(conditionalFloodRankOrgCode, 5)} || ${equalString(
      conditionalFloodRankOrg,
      "5",
    )} || ${equalNumber(conditionalFloodRank, 5)}`,
    `rgba(252, 39, 182, ${DEFAULT_TRANSPARENCY})`,
  ],
  [
    `${equalNumber(conditionalFloodRankOrgCode, 6)} || ${equalString(
      conditionalFloodRankOrg,
      "6",
    )} || ${equalNumber(conditionalFloodRank, 6)}`,
    `rgba(228, 30, 255, ${DEFAULT_TRANSPARENCY})`,
  ],
  DEFAULT_CONDITION,
];

export const CONDITIONS: Record<
  BaseFieldProps<"floodColor">["value"]["userSettings"]["colorType"],
  Condition[]
> = {
  water: [["true", `rgba(255, 255, 255, 1)`]],
  rank: FLOOD_RANK_CONDITIONS,
};
