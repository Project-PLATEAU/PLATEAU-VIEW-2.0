import CommonPage from "@web/extensions/sidebar/core/components/content/CommonPage";
import { Button, Form, Icon, Input, Checkbox, message } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { memo } from "react";

import useHooks from "./hooks";

const plateauWebsiteUrl = "https://www.mlit.go.jp/plateau/";

type Props = {
  backendURL?: string;
  isMobile?: boolean;
};

const Feedback: React.FC<Props> = ({ backendURL, isMobile }) => {
  const [form] = Form.useForm();
  const addScreenshot: boolean = Form.useWatch("screenshot", form);
  const [messageApi, contextHolder] = message.useMessage();

  const { validateMessages, handleSend, handleCancel } = useHooks({
    form,
    addScreenshot,
    backendURL,
    messageApi,
  });

  return (
    <CommonPage isMobile={isMobile}>
      <>
        {contextHolder}
        <Paragraph>
          PLATEAU は、国土交通省が進める 3D都市モデル整備・活用・オープンデータ化
          のリーディングプロジェクトである。都市活動のプラットフォームデータとして
          3D都市モデルを整備し、
          そのユースケースを創出。さらにこれをオープンデータとして公開することで、誰もが自由に都市のデータを引き出し、活用できるようになる。
        </Paragraph>
        <PlateauButton onClick={() => window.open(plateauWebsiteUrl, "_blank", "noopener")}>
          <Icon icon="plateauLogoPart" />
          PLATEAUプロジェクトサイト
        </PlateauButton>
      </>
      <>
        <Subtitle>ご意見をお聞かせください。</Subtitle>
        <Form
          form={form}
          name="feedback"
          onFinish={handleSend}
          layout="vertical"
          validateMessages={validateMessages}>
          <FormItems name="name" label="お名前（任意）">
            <Input />
          </FormItems>
          <FormItems
            name="email"
            rules={[{ type: "email", required: true, validateTrigger: "onSubmit" }]}
            label="メールアドレス">
            <Input />
          </FormItems>
          <FormItems name="comment" label="コメントまたは質問" rules={[{ required: true }]}>
            <Input.TextArea />
          </FormItems>
          <FormItems name="screenshot" valuePropName="checked">
            <Checkbox>
              <Text>マッププレビューを添付する</Text>
            </Checkbox>
          </FormItems>
          <FormButtons>
            <Button htmlType="button" onClick={handleCancel}>
              クリア
            </Button>
            <SendButton type="primary" htmlType="submit">
              送信
            </SendButton>
          </FormButtons>
        </Form>
      </>
    </CommonPage>
  );
};
export default memo(Feedback);

const Subtitle = styled.p`
  margin: 0;
  font-size: 16px;
  color: inherit;
`;

const Text = styled.p`
  margin: 0;
`;

const Paragraph = styled.p`
  font-size: 12px;
  line-height: 18px;
`;

const PlateauButton = styled.button`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  gap: 8px;
  height: 48px;
  width: 100%;
  background: transparent;
  border: 1px solid #c7c5c5;
  margin: 24px 0;
  border-radius: 4px;
  cursor: pointer;
  transition: background 0.3s;

  :hover {
    background: #d1d1d1;
  }
`;

const FormItems = styled(Form.Item)`
  margin-bottom: 8px;
  color: red;
`;

const FormButtons = styled(Form.Item)`
  display: flex;
  justify-content: flex-end;
`;

const SendButton = styled(Button)`
  margin-left: 12px;
`;
