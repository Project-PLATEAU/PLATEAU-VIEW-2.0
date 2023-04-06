import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import { getExtension } from "@web/extensions/sidebar/utils/file";
import { Input, Form, Button } from "@web/sharedComponents";
import { useCallback, useState } from "react";

import WebFileTypeSelect, { FileType, getSupportedType } from "./WebFileTypeSelect";

type Props = {
  onOpenDetails?: (data?: UserDataItem, needLayerName?: boolean) => void;
  setSelectedWebItem?: (data?: UserDataItem) => void;
};

const WebDataTab: React.FC<Props> = ({ onOpenDetails, setSelectedWebItem }) => {
  const [dataUrl, setDataUrl] = useState("");
  const [fileType, setFileType] = useState<FileType>("auto");

  const setDataFormat = useCallback((type: FileType, filename: string) => {
    if (type === "auto") {
      let extension = getSupportedType(filename);
      // Remove this in future
      if (!extension) extension = getExtension(filename);
      return extension;
    }
    return type;
  }, []);

  const needsLayerName = useCallback((url: string): boolean => {
    const serviceTypes = ["mvt", "wms", "wmts"];
    for (const serviceType of serviceTypes) {
      if (url.includes(serviceType)) {
        return true;
      }
    }
    return false;
  }, []);

  const handleClick = useCallback(async () => {
    // Catalog Item
    const filename = dataUrl.substring(dataUrl.lastIndexOf("/") + 1);
    const id = "id" + Math.random().toString(16).slice(2);
    const item: UserDataItem = {
      type: "item",
      id: id,
      dataID: id,
      description:
        "著作権や制約に関する情報などの詳細については、このデータの提供者にお問い合わせください。",
      name: filename,
      url: dataUrl,
      visible: true,
      format: setDataFormat(fileType, filename),
    };
    const requireLayerName = needsLayerName(dataUrl);
    if (onOpenDetails) onOpenDetails(item, requireLayerName);
    if (setSelectedWebItem) setSelectedWebItem(item);
  }, [dataUrl, fileType, needsLayerName, onOpenDetails, setDataFormat, setSelectedWebItem]);

  const handleFileTypeSelect = useCallback((type: string) => {
    setFileType(type as FileType);
  }, []);

  return (
    <Form layout="vertical">
      <Form.Item name="file-type" label="ファイルタイプを選択">
        <WebFileTypeSelect onFileTypeSelect={handleFileTypeSelect} />
      </Form.Item>
      <Form.Item
        name="url"
        label="データのURLを入力"
        rules={[
          { required: true },
          { message: "データファイルまたはWebサービスのURLを入力してください。" },
          { type: "url", warningOnly: true },
        ]}>
        <Input
          placeholder={"正しいURLを入力してください。"}
          onChange={e => setDataUrl(e.target.value)}
        />
      </Form.Item>
      <Form.Item style={{ textAlign: "right" }}>
        <Button type="primary" htmlType="submit" onClick={handleClick}>
          データの閲覧
        </Button>
      </Form.Item>
    </Form>
  );
};

export default WebDataTab;
