import { useCallback, useMemo } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Form from "@reearth-cms/components/atoms/Form";
import Select from "@reearth-cms/components/atoms/Select";
import { User } from "@reearth-cms/components/molecules/AccountSettings/types";
import { localesWithLabel, useT } from "@reearth-cms/i18n";

export type Props = {
  user?: User;
  onLanguageUpdate: (lang?: string | undefined) => Promise<void>;
};

const AccountServiceForm: React.FC<Props> = ({ user, onLanguageUpdate }) => {
  const [form] = Form.useForm();
  const { Option } = Select;
  const t = useT();

  const langItems = useMemo(
    () => [
      { key: "und", label: t("Auto") },
      ...Object.keys(localesWithLabel).map(l => ({
        key: l as keyof typeof localesWithLabel,
        label: localesWithLabel[l as keyof typeof localesWithLabel],
      })),
    ],
    [t],
  );

  const handleSubmit = useCallback(async () => {
    const values = await form.validateFields();
    await onLanguageUpdate?.(values.lang);
  }, [form, onLanguageUpdate]);

  return (
    <Form
      style={{ maxWidth: 400 }}
      form={form}
      initialValues={user}
      layout="vertical"
      autoComplete="off">
      <Form.Item
        name="lang"
        label={t("Service Language")}
        extra={t("This will change the UI language")}>
        <Select placeholder={t("Language")}>
          {langItems?.map(langItem => (
            <Option key={langItem.key} value={langItem.key}>
              {langItem.label}
            </Option>
          ))}
        </Select>
      </Form.Item>
      <Button onClick={handleSubmit} type="primary" htmlType="submit">
        {t("Save")}
      </Button>
    </Form>
  );
};

export default AccountServiceForm;
