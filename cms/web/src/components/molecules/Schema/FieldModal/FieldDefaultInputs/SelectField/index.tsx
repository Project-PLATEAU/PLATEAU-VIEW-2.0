import React from "react";

import Form from "@reearth-cms/components/atoms/Form";
import Select from "@reearth-cms/components/atoms/Select";
import MultiValueSelect from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueSelect";
import { useT } from "@reearth-cms/i18n";

export interface Props {
  selectedValues: string[];
  multiple?: boolean;
}

const SelectField: React.FC<Props> = ({ selectedValues, multiple }) => {
  const t = useT();
  const { Option } = Select;

  return (
    <Form.Item name="defaultValue" label={t("Set default value")}>
      {multiple ? (
        <MultiValueSelect selectedValues={selectedValues} />
      ) : (
        <Select allowClear>
          {selectedValues?.map((value: string) => (
            <Option key={value} value={value}>
              {value}
            </Option>
          ))}
        </Select>
      )}
    </Form.Item>
  );
};

export default SelectField;
