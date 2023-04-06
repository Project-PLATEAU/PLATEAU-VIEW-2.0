import { BuildingSearch, DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import {
  generateID,
  mergeOverrides,
  moveItemDown,
  moveItemUp,
  postMsg,
} from "@web/extensions/sidebar/utils";
import { getActiveFieldIDs } from "@web/extensions/sidebar/utils/dataset";
import { useCallback, useEffect, useState } from "react";

import { cleanseOverrides } from "./Field/fieldConstants";
import generateFieldComponentsList from "./Field/fieldHooks";
import { ConfigData } from "./Field/Fields/types";

export default ({
  dataset,
  inEditor,
  templates,
  buildingSearch,
  onDatasetUpdate,
  onOverride,
}: {
  dataset: DataCatalogItem;
  inEditor?: boolean;
  templates?: Template[];
  buildingSearch?: BuildingSearch;
  onDatasetUpdate: (dataset: DataCatalogItem, cleanseOverride?: any) => void;
  onOverride?: (dataID: string, activeIDs?: string[]) => void;
}) => {
  const [activeComponentIDs, setActiveIDs] = useState<string[] | undefined>();

  useEffect(() => {
    const newActiveIDs = getActiveFieldIDs(
      dataset.components,
      dataset.selectedGroup,
      dataset.config?.data,
      templates,
    );

    if (newActiveIDs !== activeComponentIDs) {
      setActiveIDs(newActiveIDs);
    }
  }, [dataset.components, dataset.selectedGroup, templates]); // eslint-disable-line react-hooks/exhaustive-deps

  useEffect(() => {
    const buildingSearchActive = buildingSearch?.find(b => b.dataID === dataset.dataID)?.active;
    if (activeComponentIDs || buildingSearchActive) {
      onOverride?.(dataset.dataID, [...(activeComponentIDs ?? [])]);
    }
  }, [activeComponentIDs, buildingSearch]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleCurrentGroupUpdate = useCallback(
    (fieldGroupID?: string) => {
      if (fieldGroupID === dataset.selectedGroup) return;
      postMsg({ action: "unselect" });
      onDatasetUpdate?.({
        ...dataset,
        selectedGroup: fieldGroupID,
      });
    },
    [dataset, onDatasetUpdate],
  );

  const handleCurrentDatasetUpdate = useCallback(
    (selectedDataset?: ConfigData) => {
      if (selectedDataset === dataset.selectedDataset) return;
      postMsg({ action: "unselect" });
      onDatasetUpdate?.({
        ...dataset,
        selectedDataset,
      });
    },
    [dataset, onDatasetUpdate],
  );

  const handleFieldAdd =
    (property: any) =>
    ({ key }: { key: string }) => {
      if (!inEditor) return;
      const newField = {
        id: generateID(),
        type: key,
        ...property,
      };

      onDatasetUpdate?.({
        ...dataset,
        components: [...(dataset.components ?? []), newField],
      });
    };

  const handleFieldUpdate = useCallback(
    (id: string) => (property: any) => {
      const newDatasetComponents = dataset.components ? [...dataset.components] : [];
      const componentIndex = newDatasetComponents?.findIndex(c => c.id === id);

      if (!newDatasetComponents || componentIndex === undefined) return;

      newDatasetComponents[componentIndex] = { ...property };

      onDatasetUpdate?.({
        ...dataset,
        components: newDatasetComponents,
      });
    },
    [dataset, onDatasetUpdate],
  );

  const handleFieldRemove = useCallback(
    (id: string) => {
      if (!inEditor) return;
      const newDatasetComponents = dataset.components ? [...dataset.components] : [];
      const componentIndex = newDatasetComponents?.findIndex(c => c.id === id);

      if (!newDatasetComponents || componentIndex === undefined) return;

      const removedComponent = newDatasetComponents.splice(componentIndex, 1)[0];

      if (removedComponent.type === "switchGroup") {
        handleCurrentGroupUpdate(undefined);
      }

      let cleanseOverride: any = undefined;
      if (removedComponent.type === "template") {
        cleanseOverride = mergeOverrides(
          "cleanse",
          templates?.find(t => t.id === removedComponent.templateID)?.components,
        );
      } else {
        cleanseOverride = cleanseOverrides[removedComponent.type] ?? undefined;
      }

      onDatasetUpdate?.(
        {
          ...dataset,
          components: newDatasetComponents,
        },
        cleanseOverride,
      );
    },
    [dataset, inEditor, templates, onDatasetUpdate, handleCurrentGroupUpdate],
  );

  const handleGroupsUpdate = useCallback(
    (fieldID: string) => (selectedGroupID?: string) => {
      if (!inEditor) return;

      const newDatasetComponents = dataset.components ? [...dataset.components] : [];
      const componentIndex = newDatasetComponents.findIndex(c => c.id === fieldID);

      if (newDatasetComponents.length > 0 && componentIndex !== undefined) {
        newDatasetComponents[componentIndex].group = selectedGroupID;
      }

      onDatasetUpdate?.({
        ...dataset,
        components: newDatasetComponents,
      });
    },
    [dataset, inEditor, onDatasetUpdate],
  );

  const fieldComponentsList = generateFieldComponentsList({
    dataset,
    onFieldAdd: handleFieldAdd,
  });

  const handleMoveUp = useCallback(
    (idx: number) => {
      const newComponents = moveItemUp(idx, dataset.components);
      if (newComponents) onDatasetUpdate({ ...dataset, components: newComponents });
    },
    [dataset, onDatasetUpdate],
  );

  const handleMoveDown = useCallback(
    (idx: number) => {
      const newComponents = moveItemDown(idx, dataset.components);
      if (newComponents) onDatasetUpdate({ ...dataset, components: newComponents });
    },
    [dataset, onDatasetUpdate],
  );

  return {
    activeComponentIDs,
    fieldComponentsList,
    handleFieldUpdate,
    handleFieldRemove,
    handleMoveUp,
    handleMoveDown,
    handleCurrentGroupUpdate,
    handleCurrentDatasetUpdate,
    handleGroupsUpdate,
  };
};
