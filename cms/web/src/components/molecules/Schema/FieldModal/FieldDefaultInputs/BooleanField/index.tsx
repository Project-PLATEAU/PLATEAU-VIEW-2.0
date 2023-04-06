import React from "react";

import Form from "@reearth-cms/components/atoms/Form";
import Switch from "@reearth-cms/components/atoms/Switch";
import MultiValueSwitch from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueSwitch";
import { useT } from "@reearth-cms/i18n";

type Props = {
  multiple?: boolean;
};

const BooleanField: React.FC<Props> = ({ multiple }) => {
  const t = useT();

  return (
    <Form.Item name="defaultValue" valuePropName="checked" label={t("Set default value")}>
      {multiple ? <MultiValueSwitch /> : <Switch />}
    </Form.Item>
  );
};

export default BooleanField;
