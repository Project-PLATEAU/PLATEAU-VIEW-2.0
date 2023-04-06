import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Wrapper } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { Select } from "@web/sharedComponents";
import { useCallback, useState } from "react";

import { BaseFieldProps } from "../types";

const displayStyleOptions = [
  { value: null, label: "Default" },
  { value: "attributes", label: "Properties" },
  { value: "description", label: "Description" },
];

const InfoboxStyle: React.FC<BaseFieldProps<"infoboxStyle">> = ({ value, editMode, onUpdate }) => {
  const [displayStyleValue, setDisplayStyleValue] = useState<"attributes" | "description" | null>(
    value.displayStyle,
  );

  const handleEventTypeChange = useCallback(
    (selectedProperty: "attributes" | "description" | null) => {
      setDisplayStyleValue(selectedProperty);
      onUpdate({
        ...value,
        displayStyle: selectedProperty,
        override: {
          infobox: {
            property: {
              default: {
                defaultContent: selectedProperty,
              },
            },
          },
        },
      });
    },
    [onUpdate, value],
  );

  return editMode ? (
    <Wrapper>
      <Field
        title="Display Style"
        titleWidth={88}
        noBorder
        value={
          <Select
            defaultValue={null}
            options={displayStyleOptions}
            style={{ width: "100%" }}
            value={displayStyleValue}
            onChange={handleEventTypeChange}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}
          />
        }
      />
    </Wrapper>
  ) : null;
};

export default InfoboxStyle;
