import { useEffect, useMemo } from "react";

import Checkbox from "@reearth-cms/components/atoms/Checkbox";
import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import { useT } from "@reearth-cms/i18n";

type Props = {
  uploadUrl: { url: string; autoUnzip: boolean };
  setUploadUrl: (uploadUrl: { url: string; autoUnzip: boolean }) => void;
};

export type FormValues = {
  url: string;
};

const UrlTab: React.FC<Props> = ({ uploadUrl, setUploadUrl }) => {
  const isCompressedFile = useMemo(() => uploadUrl.url.match(/\.zip|\.7z$/), [uploadUrl]);
  const t = useT();
  const [form] = Form.useForm();

  const initialValues: FormValues = {
    url: "",
  };

  useEffect(() => {
    form.setFieldValue("url", uploadUrl.url);
  }, [form, uploadUrl.url]);

  return (
    <Form form={form} layout="vertical" initialValues={initialValues}>
      <Form.Item
        name="url"
        label={t("URL")}
        rules={[
          { required: true },
          { message: t("Please input the URL of the asset!") },
          { type: "url", warningOnly: true },
        ]}>
        <Input
          placeholder={t("Please input a valid URL")}
          onChange={e =>
            setUploadUrl({
              ...uploadUrl,
              url: e.target.value,
            })
          }
        />
      </Form.Item>
      {isCompressedFile && (
        <Checkbox
          checked={uploadUrl.autoUnzip}
          onChange={() => {
            setUploadUrl({
              ...uploadUrl,
              autoUnzip: !uploadUrl.autoUnzip,
            });
          }}>
          {t("Auto Unzip")}
        </Checkbox>
      )}
    </Form>
  );
};

export default UrlTab;
