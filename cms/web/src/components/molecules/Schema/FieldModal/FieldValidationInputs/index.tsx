import React from "react";

import Form from "@reearth-cms/components/atoms/Form";
import InputNumber from "@reearth-cms/components/atoms/InputNumber";
import { useT } from "@reearth-cms/i18n";

import { FieldType } from "../../types";

export interface Props {
  selectedType: FieldType;
}

const FieldValidationInputs: React.FC<Props> = ({ selectedType }) => {
  const t = useT();
  return selectedType ? (
    selectedType === "Text" || selectedType === "TextArea" || selectedType === "MarkdownText" ? (
      <Form.Item name="maxLength" label={t("Set maximum length")}>
        <InputNumber type="number" />
      </Form.Item>
    ) : selectedType === "Integer" ? (
      <>
        <Form.Item name="min" label={t("Set minimum value")}>
          <InputNumber type="number" />
        </Form.Item>
        <Form.Item name="max" label={t("Set maximum value")}>
          <InputNumber type="number" />
        </Form.Item>
      </>
    ) : null
  ) : null;
};

export default FieldValidationInputs;
