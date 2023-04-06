import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";

import {
  ConfigData,
  FieldComponent,
  SwitchGroup,
} from "../core/components/content/common/DatasetCard/Field/Fields/types";
import { Data, DataCatalogItem, Template } from "../core/types";
import { RawDataCatalogItem } from "../modals/datacatalog/api/api";

import { generateID } from ".";

export const getActiveFieldIDs = (
  components?: FieldComponent[],
  selectedGroup?: string,
  config?: ConfigData[],
  templates?: Template[],
) =>
  flattenComponents(components, templates)
    ?.filter(
      c =>
        !selectedGroup ||
        !c.group ||
        c.type === "switchGroup" ||
        (c.group && c.group === selectedGroup),
    )
    ?.filter(c => !(!config && c.type === "switchDataset"))
    ?.map(c => c.id);

export const flattenComponents = (components?: FieldComponent[], baseTemplates?: Template[]) =>
  components?.reduce((a: FieldComponent[], c?: FieldComponent) => {
    if (!c) return a;
    if (c.type === "template") {
      return [
        ...a,
        c,
        ...((c as any).userSettings?.components ??
          baseTemplates?.find(t => t.id === c.templateID)?.components ??
          []),
      ];
    } else {
      return [...a, c];
    }
  }, []);

export const getDefaultGroup = (components?: FieldComponent[], baseTemplates?: Template[]) => {
  if (!components) return;

  const switchGroupComponents = flattenComponents(components, baseTemplates)?.filter(
    c => c.type === "switchGroup",
  ) as SwitchGroup[] | undefined;

  if (switchGroupComponents && switchGroupComponents.length > 0) {
    return switchGroupComponents[switchGroupComponents.length - 1].groups[0].fieldGroupID;
  }
};

export const getDefaultDataset = (dataset?: DataCatalogItem | UserDataItem) => {
  if (!dataset) return;

  if (dataset.config?.data && dataset.config?.data?.length > 0) {
    return dataset.config.data[0];
  }
};

export const processComponentsToSave = (components?: FieldComponent[], templates?: Template[]) => {
  if (!components) return;
  return components?.map((c: any) => {
    const newComp = Object.assign({}, c);
    if (newComp.type === "template" && newComp.components) {
      newComp.components =
        templates
          ?.find(t => t.id === newComp.templateID)
          ?.components?.map(c => {
            return { ...c, userSettings: undefined };
          }) ?? [];
    }
    return { ...newComp, userSettings: undefined };
  });
};

export const convertDatasetToData = (dataset: DataCatalogItem, templates?: Template[]): Data => {
  return {
    dataID: dataset.dataID,
    public: dataset.public,
    components: processComponentsToSave(dataset.components, templates),
  };
};

export const newItem = (ri: RawDataCatalogItem): DataCatalogItem => {
  return {
    ...ri,
    dataID: ri.id,
    public: false,
    visible: true,
  };
};

export const handleDataCatalogProcessing = (
  catalog: (DataCatalogItem | RawDataCatalogItem)[],
  savedData?: Data[],
): DataCatalogItem[] =>
  catalog.map(item => {
    if (!savedData) return newItem(item);

    const savedData2 = savedData.find(d => d.dataID === ("dataID" in item ? item.dataID : item.id));
    if (savedData2) {
      return {
        ...item,
        ...savedData2,
        visible: true,
        selectedGroup: getDefaultGroup(savedData2.components),
      };
    } else {
      return newItem(item);
    }
  });

export const processDatasetToAdd = (
  dataset: DataCatalogItem | UserDataItem,
  templates?: Template[],
): DataCatalogItem => {
  const datasetToAdd = { ...dataset };

  datasetToAdd.selectedGroup = getDefaultGroup(datasetToAdd.components, templates);
  datasetToAdd.selectedDataset = getDefaultDataset(datasetToAdd);

  if (!dataset.components?.length) {
    const defaultTemplate = templates?.find(t =>
      dataset.type2
        ? t.name.includes(dataset.type2)
        : dataset.type
        ? t.name.includes(dataset.type)
        : undefined,
    );
    if (defaultTemplate && !datasetToAdd.components) {
      datasetToAdd.components = [
        {
          id: generateID(),
          type: "template",
          templateID: defaultTemplate.id,
          userSettings: {
            components: defaultTemplate.components,
          },
        },
      ];
      datasetToAdd.selectedGroup = getDefaultGroup(defaultTemplate.components);
    }
  }
  return datasetToAdd as DataCatalogItem;
};
