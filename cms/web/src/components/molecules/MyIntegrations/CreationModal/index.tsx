import { useCallback } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import Modal from "@reearth-cms/components/atoms/Modal";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { IntegrationType } from "@reearth-cms/components/molecules/MyIntegrations/types";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  open?: boolean;
  onClose: () => void;
  onSubmit?: (values: FormValues) => Promise<void> | void;
};

export type FormValues = {
  name: string;
  description: string;
  logoUrl: string;
  type: IntegrationType;
};

const initialValues: FormValues = {
  name: "",
  description: "",
  logoUrl: "",
  type: IntegrationType.Private,
};

const IntegrationCreationModal: React.FC<Props> = ({ open, onClose, onSubmit }) => {
  const t = useT();
  const [form] = Form.useForm();

  const handleSubmit = useCallback(() => {
    form
      .validateFields()
      .then(async (values: FormValues) => {
        // TODO: when assets upload is ready to use
        values.logoUrl = "_";
        values.type = IntegrationType.Private;
        await onSubmit?.(values);
        onClose();
        form.resetFields();
      })
      .catch(info => {
        console.log("Validate Failed:", info);
      });
  }, [form, onClose, onSubmit]);

  const handleClose = useCallback(() => {
    form.resetFields();
    onClose();
  }, [onClose, form]);

  return (
    <Modal
      visible={open}
      onCancel={handleClose}
      onOk={handleSubmit}
      title={t("New Integration")}
      footer={[
        <Button key="back" onClick={handleClose}>
          {t("Cancel")}
        </Button>,
        <Button key="submit" type="primary" onClick={handleSubmit}>
          {t("Create")}
        </Button>,
      ]}>
      <Form form={form} layout="vertical" initialValues={initialValues}>
        <Form.Item
          name="name"
          label={t("Integration Name")}
          rules={[
            {
              required: true,
              message: t("Please input the title of the integration!"),
            },
          ]}>
          <Input />
        </Form.Item>
        <Form.Item name="description" label={t("Description")}>
          <TextArea rows={3} showCount maxLength={100} />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default IntegrationCreationModal;
