import { postMsg } from "@web/extensions/sidebar/utils";
import { useCallback, useEffect, useState } from "react";

const validateMessages = {
  required: "このフィルドがない場合は返信できません",
  types: {
    email: "メールアドレスが正しくないです",
  },
};

export default ({
  form,
  addScreenshot,
  backendURL,
  messageApi,
}: {
  form: any;
  addScreenshot: boolean;
  backendURL?: string;
  messageApi: any;
}) => {
  const [screenshot, setScreenshot] = useState<any>();

  const handleSend = useCallback(
    async (values: { name: string; email: string; comment: string }) => {
      if (!backendURL) return;
      const formData = new FormData();
      formData.append("name", values.name);
      formData.append("email", values.email);
      formData.append("content", values.comment);
      if (screenshot) {
        const file = dataURItoBlob(screenshot);
        formData.append("file", file);
      }

      const resp = await fetch(`${backendURL}/opinion`, {
        method: "POST",
        body: formData,
      });
      if (resp.status !== 200) {
        messageApi.open({
          type: "error",
          content:
            "フィードバックの送信に失敗しました。しばらく時間を空けてからもう一度お試しください。",
        });
      } else {
        messageApi.open({
          type: "success",
          content: "フィードバックを送りました。ありがとうございます。",
        });
        form.resetFields();
      }
    },
    [form, backendURL, messageApi, screenshot],
  );

  const handleCancel = useCallback(() => {
    form.resetFields();
  }, [form]);

  useEffect(() => {
    if (addScreenshot) {
      postMsg({ action: "screenshot" });
    } else {
      setScreenshot(undefined);
    }
  }, [addScreenshot]);

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "screenshot") {
        setScreenshot(e.data.payload);
      }
    };
    addEventListener("message", eventListenerCallback);
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, []);

  return {
    validateMessages,
    handleSend,
    handleCancel,
  };
};

export const dataURItoBlob = (dataURI: string) => {
  const byteString = atob(dataURI.split(",")[1]);
  const mimeString = dataURI.split(",")[0].split(":")[1].split(";")[0];
  const ab = new ArrayBuffer(byteString.length);
  const ia = new Uint8Array(ab);
  for (let i = 0; i < byteString.length; i++) {
    ia[i] = byteString.charCodeAt(i);
  }
  return new Blob([ab], { type: mimeString });
};
