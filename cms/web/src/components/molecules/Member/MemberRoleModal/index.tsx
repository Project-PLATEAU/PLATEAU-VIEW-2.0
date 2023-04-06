import React, { useCallback, useEffect } from "react";

import Form from "@reearth-cms/components/atoms/Form";
import Modal from "@reearth-cms/components/atoms/Modal";
import Select from "@reearth-cms/components/atoms/Select";
import { RoleUnion } from "@reearth-cms/components/organisms/Settings/Members/hooks";
import { useT } from "@reearth-cms/i18n";

export interface FormValues {
  userId: string;
  role: string;
}

export interface Props {
  open?: boolean;
  member?: any;
  onClose?: (refetch?: boolean) => void;
  onSubmit?: (userId: string, role: RoleUnion) => Promise<void>;
}

const MemberRoleModal: React.FC<Props> = ({ open, onClose, onSubmit, member }) => {
  const t = useT();
  const { Option } = Select;
  const [form] = Form.useForm();

  useEffect(() => {
    form.setFieldsValue({
      userId: member?.userId,
      role: member?.role,
    });
  }, [form, member]);

  const handleSubmit = useCallback(() => {
    form
      .validateFields()
      .then(async values => {
        await onSubmit?.(member?.userId, values?.role);
        onClose?.(true);
        form.resetFields();
      })
      .catch(info => {
        console.log("Validate Failed:", info);
      });
  }, [form, onClose, onSubmit, member?.userId]);

  const handleClose = useCallback(() => {
    form.resetFields();
    onClose?.(true);
  }, [form, onClose]);

  return (
    <Modal title={t("Role Settings")} visible={open} onCancel={handleClose} onOk={handleSubmit}>
      <Form
        form={form}
        layout="vertical"
        initialValues={{
          userId: member?.userId,
          role: member?.role,
        }}>
        <Form.Item
          name="role"
          label="Role"
          rules={[
            {
              required: true,
              message: t("Please input the appropriate role for this member!"),
            },
          ]}>
          <Select placeholder={t("select role")}>
            <Option value="OWNER">{t("Owner")}</Option>
            <Option value="WRITER">{t("Writer")}</Option>
            <Option value="MAINTAINER">{t("Maintainer")}</Option>
            <Option value="READER">{t("Reader")}</Option>
          </Select>
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default MemberRoleModal;
