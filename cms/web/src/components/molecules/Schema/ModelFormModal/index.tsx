import { useCallback, useState, useEffect } from "react";

import Form, { FieldError } from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import Modal from "@reearth-cms/components/atoms/Modal";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { Model } from "@reearth-cms/components/molecules/ProjectOverview";
import { useT } from "@reearth-cms/i18n";
import { validateKey } from "@reearth-cms/utils/regex";

export interface FormValues {
  modelId?: string;
  name: string;
  description: string;
  key: string;
}

export interface Props {
  model?: Model;
  open?: boolean;
  isKeyAvailable: boolean;
  onClose: () => void;
  onCreate?: (values: FormValues) => Promise<void> | void;
  OnUpdate?: (values: FormValues) => Promise<void> | void;
  onModelKeyCheck: (key: string, ignoredKey?: string) => Promise<boolean>;
}

const ModelFormModal: React.FC<Props> = ({
  model,
  open,
  onClose,
  onCreate,
  OnUpdate,
  onModelKeyCheck,
}) => {
  const t = useT();
  const [form] = Form.useForm();
  const [buttonDisabled, setButtonDisabled] = useState(true);

  useEffect(() => {
    if (!model) {
      form.resetFields();
    } else {
      form.setFieldsValue(model);
    }
  }, [form, model]);

  const handleSubmit = useCallback(async () => {
    const values = await form.validateFields();
    await onModelKeyCheck(values.key, model?.key);
    if (!model?.id) {
      await onCreate?.(values);
    } else {
      await OnUpdate?.({ modelId: model.id, ...values });
    }
    onClose();
    form.resetFields();
  }, [onModelKeyCheck, model, form, onClose, onCreate, OnUpdate]);

  const handleClose = useCallback(() => {
    form.resetFields();
    onClose();
  }, [form, onClose]);

  return (
    <Modal
      visible={open}
      onCancel={handleClose}
      onOk={handleSubmit}
      okButtonProps={{ disabled: buttonDisabled }}
      title={!model?.id ? t("New Model") : t("Update Model")}>
      <Form
        form={form}
        layout="vertical"
        onValuesChange={() => {
          form
            .validateFields()
            .then(() => {
              setButtonDisabled(false);
            })
            .catch(fieldsError => {
              setButtonDisabled(
                fieldsError.errorFields.some((item: FieldError) => item.errors.length > 0),
              );
            });
        }}>
        <Form.Item
          name="name"
          label={t("Model name")}
          rules={[{ required: true, message: t("Please input the name of the model!") }]}>
          <Input />
        </Form.Item>
        <Form.Item name="description" label={t("Model description")}>
          <TextArea rows={4} />
        </Form.Item>
        <Form.Item
          name="key"
          label={t("Model key")}
          extra={t(
            "Model key must be unique and at least 1 character long. It can only contain letters, numbers, underscores and dashes.",
          )}
          rules={[
            {
              message: t("Key is not valid"),
              required: true,
              validator: async (_, value) => {
                if (!validateKey(value)) return Promise.reject();
                const isKeyAvailable = await onModelKeyCheck(value, model?.key);
                if (isKeyAvailable) {
                  return Promise.resolve();
                } else {
                  return Promise.reject();
                }
              },
            },
          ]}>
          <Input />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default ModelFormModal;
