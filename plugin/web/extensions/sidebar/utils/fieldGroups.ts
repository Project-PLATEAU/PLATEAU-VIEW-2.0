import { Group } from "../core/types";

export const fieldGroupMax = 24; // Lowering this from 24 will break some saved Data. Change only with utmost caution.

const generateFieldGroups: () => Group[] = () => {
  const groups = [];
  for (let i = 1; i <= fieldGroupMax; i++) {
    groups.push({ id: `group${i}`, name: `グループ${i}` });
  }
  return groups;
};

export const fieldGroups: Group[] = generateFieldGroups();
