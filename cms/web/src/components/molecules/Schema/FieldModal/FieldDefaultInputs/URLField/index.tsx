import React from "react";

import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import MultiValueField from "@reearth-cms/components/molecules/Common/MultiValueField";
import { useT } from "@reearth-cms/i18n";
import { validateURL } from "@reearth-cms/utils/regex";

type Props = {
  multiple?: boolean;
};

const URLField: React.FC<Props> = ({ multiple }) => {
  const t = useT();
  return (
    <Form.Item
      name="defaultValue"
      label="Set default value"
      extra={t("Default value must be a valid URL and start with 'http://' or 'https://'.")}
      rules={[
        {
          message: "URL is not valid",
          validator: async (_, value) => {
            if (value) {
              if (
                multiple &&
                value.some((valueItem: string) => !validateURL(valueItem) && valueItem.length > 0)
              )
                return Promise.reject();
              else if (!multiple && !validateURL(value) && value?.length > 0)
                return Promise.reject();
            }
            return Promise.resolve();
          },
        },
      ]}>
      {multiple ? <MultiValueField FieldInput={Input} /> : <Input />}
    </Form.Item>
  );
};

export default URLField;
