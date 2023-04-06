import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Wrapper } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { Select } from "@web/sharedComponents";
import { useCallback, useState } from "react";

import { BaseFieldProps } from "../types";

const heightReferenceOptions = [
  { value: "clamp", label: "地表に固定" },
  { value: "relative", label: "地表からの高度" },
  { value: "none", label: "なし" },
];

const HeightReference: React.FC<BaseFieldProps<"heightReference">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  const [heightReferenceType, setHeightReferenceType] = useState<"clamp" | "relative" | "none">(
    value.heightReferenceType ?? "clamp",
  );

  const handleEventTypeChange = useCallback(
    (selectedProperty: string) => {
      const newHeightReferenceType = selectedProperty as "clamp" | "relative" | "none";

      setHeightReferenceType(newHeightReferenceType);

      onUpdate({
        ...value,
        heightReferenceType: newHeightReferenceType,
        override: {
          resource: {
            clampToGround: newHeightReferenceType === "clamp",
          },
          marker: {
            heightReference: newHeightReferenceType,
          },
          polygon: {
            heightReference: newHeightReferenceType,
          },
          polyline: {
            clampToGround: newHeightReferenceType === "clamp",
          },
        },
      });
    },
    [onUpdate, value],
  );

  return editMode ? (
    <Wrapper>
      <Field
        title="UIスタイル"
        titleWidth={88}
        noBorder
        value={
          <Select
            defaultValue={null}
            options={heightReferenceOptions}
            style={{ width: "100%" }}
            value={heightReferenceType}
            onChange={handleEventTypeChange}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}
          />
        }
      />
    </Wrapper>
  ) : null;
};

export default HeightReference;
