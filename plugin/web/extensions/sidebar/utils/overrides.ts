import { cleanseOverrides } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/fieldConstants";
import { FieldComponent } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/Fields/types";
import { BuildingSearch, DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import { getTransparencyExpression } from "@web/extensions/sidebar/utils/color";
import { flattenComponents } from "@web/extensions/sidebar/utils/dataset";
import { cloneDeep, isEqual, merge } from "lodash";

export const prepareComponentsForOverride = (
  activeIDs: string[],
  dataset: DataCatalogItem,
  templates?: Template[],
  buildingSearch?: BuildingSearch,
) => {
  const flattenedComponents = flattenComponents(dataset.components, templates);
  const inactiveComponents = flattenedComponents
    ?.filter(c => !activeIDs.find(id => id === c.id))
    .map(c => {
      if (c.type === "switchDataset" && !c.cleanseOverride) {
        c.cleanseOverride = {
          data: {
            url: dataset.config?.data?.[0].url,
            time: {
              updateClockOnLoad: false,
            },
          },
        };
      }
      return c;
    });

  const activeComponents = flattenedComponents
    ?.filter(c => !!activeIDs.find(id => id === c.id))
    .map(c => {
      if (c.type === "switchDataset" && !c.cleanseOverride) {
        c.cleanseOverride = {
          data: {
            url: dataset.config?.data?.[0].url,
            time: {
              updateClockOnLoad: false,
            },
          },
        };
      }
      return c;
    });

  const buildingSearchField = buildingSearch?.find(b => b.dataID === dataset.dataID);
  if (buildingSearchField) {
    if (buildingSearchField.active) {
      activeComponents?.push(buildingSearchField.field as FieldComponent);
    } else {
      inactiveComponents?.push(buildingSearchField.cleanseField as FieldComponent);
    }
  }
  return {
    activeComponents,
    inactiveComponents,
  };
};

export const mergeOverrides = (
  action: "update" | "cleanse",
  components?: FieldComponent[],
  startingOverride?: any,
) => {
  if (!components || !components.length) {
    if (startingOverride) {
      return startingOverride;
    }
    return;
  }

  const overrides = cloneDeep(startingOverride ?? {});

  const needOrderComponents = components
    .filter(c => (c as any).userSettings?.updatedAt)
    .sort(
      (a, b) =>
        ((a as any).userSettings?.updatedAt?.getTime?.() ?? 0) -
        ((b as any).userSettings?.updatedAt?.getTime?.() ?? 0),
    );
  for (const component of needOrderComponents) {
    merge(
      overrides,
      action === "cleanse"
        ? cleanseOverrides[component.type]
        : (component as any).userSettings?.override ?? component.override,
    );
  }

  let transparency = 100;
  let switchGroupExist = false;

  for (let i = 0; i < components.length; i++) {
    if ((components[i] as any).userSettings?.updatedAt) {
      transparency = (components[i] as any)?.userSettings?.transparency ?? 100;
      continue;
    }
    if (components[i].type === "switchDataset") {
      const switchDatasetOverride =
        (components[i] as any).userSettings?.override ??
        (action === "cleanse"
          ? components[i].cleanseOverride
          : {
              data: {
                ...(components[i].cleanseOverride.data.url
                  ? { url: components[i].cleanseOverride.data.url }
                  : {}),
                time: {
                  updateClockOnLoad: true,
                },
              },
            });
      merge(overrides, switchDatasetOverride);
      continue;
    }

    if (components[i].type === "switchGroup") {
      switchGroupExist = true;
    }

    merge(
      overrides,
      action === "cleanse"
        ? cleanseOverrides[components[i].type]
        : (components[i] as any).userSettings?.override ?? components[i].override,
    );
  }

  // This is a temporary solution for switch groups and transparency to work together: @pyshx
  if (switchGroupExist && transparency != 100) {
    const { expression } = getTransparencyExpression(overrides, transparency / 100, false);
    merge(overrides, {
      "3dtiles": {
        color: expression,
      },
    });
  }

  return isEqual(overrides, {}) ? undefined : overrides;
};
