import { useCallback, useEffect } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Form from "@reearth-cms/components/atoms/Form";
import Modal from "@reearth-cms/components/atoms/Modal";
import Select from "@reearth-cms/components/atoms/Select";
import { IntegrationMember, Role } from "@reearth-cms/components/molecules/Integration/types";
import { useT } from "@reearth-cms/i18n";

export type FormValues = {
  role: Role;
};

export type Props = {
  selectedIntegrationMember?: IntegrationMember;
  open?: boolean;
  onClose?: () => void;
  onSubmit?: (role: Role) => Promise<void> | void;
};

const IntegrationSettingsModal: React.FC<Props> = ({
  open,
  onClose,
  onSubmit,
  selectedIntegrationMember,
}) => {
  const t = useT();
  const { Option } = Select;
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue({
      role: selectedIntegrationMember?.integrationRole,
    });
  }, [form, selectedIntegrationMember]);

  const handleSubmit = useCallback(async () => {
    try {
      const values = await form.validateFields();
      await onSubmit?.(values.role);
      onClose?.();
      form.resetFields();
    } catch (info) {
      console.log("Validate Failed:", info);
    }
  }, [form, onClose, onSubmit]);

  return (
    <Modal
      title={t("Integration Setting") + "  " + selectedIntegrationMember?.integration?.name}
      visible={open}
      onCancel={() => onClose?.()}
      footer={[
        <Button key="back" onClick={() => onClose?.()}>
          {t("Cancel")}
        </Button>,
        <Button key="submit" type="primary" onClick={handleSubmit}>
          {t("Save")}
        </Button>,
      ]}>
      <Form
        form={form}
        layout="vertical"
        initialValues={{ role: selectedIntegrationMember?.integrationRole }}>
        <Form.Item
          name="role"
          label="Role"
          rules={[
            {
              required: true,
              message: t("Please input the appropriate role for this integration!"),
            },
          ]}>
          <Select placeholder={t("select role")}>
            <Option value="READER">{t("Reader")}</Option>
            <Option value="WRITER">{t("Writer")}</Option>
            <Option value="MAINTAINER">{t("Maintainer")}</Option>
            <Option value="OWNER">{t("Owner")}</Option>
          </Select>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default IntegrationSettingsModal;
