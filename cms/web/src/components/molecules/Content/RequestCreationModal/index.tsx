import { useCallback } from "react";

import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import Modal from "@reearth-cms/components/atoms/Modal";
import Select, { SelectProps } from "@reearth-cms/components/atoms/Select";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { RequestState } from "@reearth-cms/components/molecules/Request/types";
import { Member } from "@reearth-cms/components/molecules/Workspace/types";
import { useT } from "@reearth-cms/i18n";

export type FormValues = {
  title: string;
  description: string;
  state: RequestState;
  reviewersId: string[];
  items: {
    itemId: string;
  }[];
};

export type Props = {
  open?: boolean;
  itemId: string;
  workspaceUserMembers: Member[];
  onClose?: (refetch?: boolean) => void;
  onSubmit?: (data: FormValues) => Promise<void>;
};

const initialValues: FormValues = {
  title: "",
  description: "",
  state: "WAITING",
  reviewersId: [],
  items: [
    {
      itemId: "",
    },
  ],
};

const RequestCreationModal: React.FC<Props> = ({
  open,
  itemId,
  workspaceUserMembers,
  onClose,
  onSubmit,
}) => {
  const t = useT();
  const [form] = Form.useForm();

  const reviewers: SelectProps["options"] = [];
  for (const member of workspaceUserMembers) {
    reviewers.push({
      label: member.user.name,
      value: member.userId,
    });
  }

  const handleSubmit = useCallback(async () => {
    try {
      const values = await form.validateFields();
      values.items = [{ itemId }];
      values.state = "WAITING";
      await onSubmit?.(values);
      onClose?.(true);
      form.resetFields();
    } catch (info) {
      console.log("Validate Failed:", info);
    }
  }, [itemId, form, onClose, onSubmit]);

  const handleClose = useCallback(() => {
    onClose?.(true);
  }, [onClose]);
  return (
    <Modal visible={open} onCancel={handleClose} onOk={handleSubmit} title={t("New Request")}>
      <Form form={form} layout="vertical" initialValues={initialValues}>
        <Form.Item
          name="title"
          label={t("Title")}
          rules={[{ required: true, message: t("Please input the title of your request!") }]}>
          <Input />
        </Form.Item>
        <Form.Item name="description" label={t("Description")}>
          <TextArea rows={4} showCount maxLength={100} />
        </Form.Item>
        <Form.Item
          name="reviewersId"
          label="Reviewer"
          rules={[
            {
              required: true,
              message: t("Please select a reviewer!"),
            },
          ]}>
          <Select
            filterOption={(input, option) =>
              (option?.label?.toString().toLowerCase() ?? "").includes(input.toLowerCase())
            }
            placeholder={t("Reviewer")}
            mode="multiple"
            options={reviewers}
            allowClear
          />
        </Form.Item>
      </Form>
    </Modal>
  );
};

export default RequestCreationModal;
