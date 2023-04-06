import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Wrapper } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { Select } from "@web/sharedComponents";
import { useState, useCallback } from "react";

import { BaseFieldProps, ClassificationType } from "../../types";

const classificationTypesOptions = [
  {
    value: "both",
    label: "両方",
  },
  {
    value: "3dtiles",
    label: "3D Tilesのみ",
  },
  {
    value: "terrain",
    label: "テラインのみ",
  },
];

const PolylineClassificationType: React.FC<BaseFieldProps<"polylineClassificationType">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  const [selectedType, setSelectedType] = useState<ClassificationType>(
    value.classificationType ?? "both",
  );

  const handleTypeChange = useCallback(
    (selected: ClassificationType) => {
      setSelectedType(selected);

      onUpdate({
        ...value,
        classificationType: selected,
        override: {
          polyline: { classificationType: selected },
        },
      });
    },
    [onUpdate, value],
  );

  return editMode ? (
    <Wrapper>
      <Field
        title="種類"
        titleWidth={88}
        noBorder
        value={
          <Select
            defaultValue={null}
            options={classificationTypesOptions}
            style={{ width: "100%" }}
            value={selectedType}
            onChange={handleTypeChange}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}
          />
        }
      />
    </Wrapper>
  ) : null;
};

export default PolylineClassificationType;
