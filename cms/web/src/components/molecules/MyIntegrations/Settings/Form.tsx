import styled from "@emotion/styled";
import { useCallback } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Col from "@reearth-cms/components/atoms/Col";
import Divider from "@reearth-cms/components/atoms/Divider";
import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import Row from "@reearth-cms/components/atoms/Row";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { Integration } from "@reearth-cms/components/molecules/MyIntegrations/types";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  integration: Integration;
  onIntegrationUpdate: (data: { name: string; description: string; logoUrl: string }) => void;
};

const MyIntegrationForm: React.FC<Props> = ({ integration, onIntegrationUpdate }) => {
  const t = useT();
  const [form] = Form.useForm();

  const handleSubmit = useCallback(async () => {
    try {
      const values = await form.validateFields();
      // TODO: when assets upload is ready to use
      values.logoUrl = "_";
      await onIntegrationUpdate?.(values);
    } catch (info) {
      console.log("Validate Failed:", info);
    }
  }, [form, onIntegrationUpdate]);

  return (
    <Form form={form} layout="vertical" initialValues={integration}>
      <Row gutter={32}>
        <Col span={11}>
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
          <Form.Item label={t("Integration Token")}>
            <Input.Password value={integration.config.token} contentEditable={false} />
          </Form.Item>
          <Form.Item>
            <Button type="primary" htmlType="submit" onClick={handleSubmit}>
              {t("Save")}
            </Button>
          </Form.Item>
        </Col>
        <Col>
          <Divider type="vertical" style={{ height: "100%" }} />
        </Col>
        <Col span={11}>
          <CodeExampleTitle>{t("Code Example")}</CodeExampleTitle>
          <CodeExample>
            curl --location --request POST <br />
            &apos;{window.REEARTH_CONFIG?.api}/models/
            <CodeImportant>“{t("your model id here")}”</CodeImportant>/items&apos;&nbsp;\
            <br />
            --header &apos;Authorization: Bearer&nbsp;
            <CodeImportant>“your Integration Token here”</CodeImportant>&apos;
          </CodeExample>
        </Col>
      </Row>
    </Form>
  );
};

const CodeExampleTitle = styled.h2`
  font-weight: 400;
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.85);
`;

const CodeExample = styled.div`
  border: 1px solid #d9d9d9;
  padding: 5px 12px;
  font-weight: 400;
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.85);
`;

const CodeImportant = styled.span`
  color: #ff4d4f;
`;

export default MyIntegrationForm;
