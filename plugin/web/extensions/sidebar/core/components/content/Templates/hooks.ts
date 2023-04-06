import { Template } from "@web/extensions/sidebar/core/types";
import { generateID, moveItemDown, moveItemUp } from "@web/extensions/sidebar/utils";
import { useCallback } from "react";

import generateFieldComponentsList from "../common/DatasetCard/Field/fieldHooks";

export default ({
  template,
  onTemplateUpdate,
}: {
  template: Template;
  onTemplateUpdate?: (template: Template) => void;
}) => {
  const handleFieldUpdate = useCallback(
    (id: string) => (property: any) => {
      const newDatasetComponents = template.components ? [...template.components] : [];
      const componentIndex = newDatasetComponents?.findIndex(c => c.id === id);

      if (!newDatasetComponents || componentIndex === undefined) return;

      newDatasetComponents[componentIndex] = { ...property };

      onTemplateUpdate?.({
        ...template,
        components: newDatasetComponents,
      });
    },
    [template, onTemplateUpdate],
  );

  const handleFieldRemove = useCallback(
    (id: string) => {
      const newDatasetComponents = template.components ? [...template.components] : [];
      const componentIndex = newDatasetComponents?.findIndex(c => c.id === id);

      if (!newDatasetComponents || componentIndex === undefined) return;

      newDatasetComponents.splice(componentIndex, 1);

      onTemplateUpdate?.({
        ...template,
        components: newDatasetComponents,
      });
    },
    [template, onTemplateUpdate],
  );

  const handleMoveFieldUp = useCallback(
    (idx: number) => {
      const newComponents = moveItemUp(idx, template.components);
      if (newComponents) onTemplateUpdate?.({ ...template, components: newComponents });
    },
    [onTemplateUpdate, template],
  );

  const handleMoveFieldDown = useCallback(
    (idx: number) => {
      const newComponents = moveItemDown(idx, template.components);
      if (newComponents) onTemplateUpdate?.({ ...template, components: newComponents });
    },
    [template, onTemplateUpdate],
  );

  const handleGroupsUpdate = useCallback(
    (fieldID: string) => (selectedGroupID?: string) => {
      const newDatasetComponents = template.components ? [...template.components] : [];
      const componentIndex = newDatasetComponents.findIndex(c => c.id === fieldID);

      if (newDatasetComponents.length > 0 && componentIndex !== undefined) {
        newDatasetComponents[componentIndex].group = selectedGroupID;
      }

      onTemplateUpdate?.({
        ...template,
        components: newDatasetComponents,
      });
    },
    [template, onTemplateUpdate],
  );

  const fieldComponentsList = generateFieldComponentsList({
    onFieldAdd:
      (property: any) =>
      ({ key }: { key: string }) => {
        onTemplateUpdate?.({
          ...template,
          components: [
            ...(template.components ?? []),
            {
              id: generateID(),
              type: key,
              ...property,
            },
          ],
        });
      },
  });

  return {
    fieldComponentsList,
    handleFieldUpdate,
    handleFieldRemove,
    handleMoveFieldUp,
    handleMoveFieldDown,
    handleGroupsUpdate,
  };
};
