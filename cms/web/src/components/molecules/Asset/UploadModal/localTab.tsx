import { ReactElement } from "react";

import Icon from "@reearth-cms/components/atoms/Icon";
import Upload, { UploadFile, UploadProps } from "@reearth-cms/components/atoms/Upload";
import FileItem from "@reearth-cms/components/molecules/Asset/UploadModal/FileItem";
import { useT } from "@reearth-cms/i18n";

const { Dragger } = Upload;

type Props = {
  uploadProps: UploadProps;
};

const LocalTab: React.FC<Props> = ({ uploadProps }) => {
  const t = useT();
  return (
    <div>
      <Dragger
        itemRender={(
          _originNode: ReactElement,
          file: UploadFile<any>,
          _fileList: UploadFile<any>[],
          { remove },
        ) => <FileItem file={file} remove={remove} />}
        {...uploadProps}>
        <p className="ant-upload-drag-icon">
          <Icon icon="inbox" />
        </p>
        <p className="ant-upload-text">{t("Click or drag files to this area to upload")}</p>
        <p className="ant-upload-hint">{t("Single or multiple file upload is supported")}</p>
      </Dragger>
    </div>
  );
};

export default LocalTab;
