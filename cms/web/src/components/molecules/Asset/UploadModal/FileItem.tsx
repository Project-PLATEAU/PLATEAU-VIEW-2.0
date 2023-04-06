import { Button } from "antd";
import { useMemo, useState } from "react";

import Checkbox from "@reearth-cms/components/atoms/Checkbox";
import Icon from "@reearth-cms/components/atoms/Icon";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { useT } from "@reearth-cms/i18n";

import { isImageUrl } from "./util";

type Props = {
  file: UploadFile<any>;
  remove: () => void;
};

const FileItem: React.FC<Props> = ({ file, remove }) => {
  const t = useT();
  const [checked, setChecked] = useState<boolean>(!file?.skipDecompression);
  const isCompressedFile = useMemo(() => file.name?.match(/\.zip|\.7z$/), [file.name]);

  return (
    <div className="ant-upload-list-item">
      <span className="ant-upload-span">
        <div className="ant-upload-list-item-thumbnail ant-upload-list-item-file">
          <Icon icon={isImageUrl(file) ? "pictureTwoTone" : "fileTwoTone"} />
        </div>
        <span className="ant-upload-list-item-name" title={file.name}>
          {file.name}
        </span>

        {isCompressedFile && (
          <Checkbox
            checked={checked}
            onChange={() => {
              setChecked(checked => !checked);
              if (!file?.skipDecompression) file.skipDecompression = true;
              else file.skipDecompression = !file.skipDecompression;
            }}>
            {t("Auto Unzip")}
          </Checkbox>
        )}
        <Button
          type="text"
          title={t("Remove file")}
          icon={<Icon icon="delete" />}
          onClick={remove}
        />
      </span>
    </div>
  );
};

export default FileItem;
