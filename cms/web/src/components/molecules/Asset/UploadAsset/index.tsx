import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import { UploadFile, UploadProps } from "@reearth-cms/components/atoms/Upload";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import { useT } from "@reearth-cms/i18n";

import UploadModal from "../UploadModal/UploadModal";

type Props = {
  alsoLink?: boolean;
  uploadProps: UploadProps;
  fileList: UploadFile<File>[];
  uploading: boolean;
  uploadModalVisibility: boolean;
  uploadUrl: { url: string; autoUnzip: boolean };
  uploadType: UploadType;
  setUploadUrl: (uploadUrl: { url: string; autoUnzip: boolean }) => void;
  setUploadType: (type: UploadType) => void;
  onUploadModalClose?: () => void;
  displayUploadModal: () => void;
  onUploadModalCancel: () => void;
  onUpload: () => void;
};

const UploadAsset: React.FC<Props> = ({
  alsoLink,
  uploadProps,
  fileList,
  uploading,
  uploadModalVisibility,
  onUploadModalClose,
  displayUploadModal,
  onUploadModalCancel,
  uploadUrl,
  uploadType,
  setUploadUrl,
  setUploadType,
  onUpload,
}) => {
  const t = useT();
  return (
    <>
      <Button type="primary" icon={<Icon icon="upload" />} onClick={displayUploadModal}>
        {t("Upload Asset")}
      </Button>
      <UploadModal
        alsoLink={alsoLink}
        uploadProps={uploadProps}
        fileList={fileList}
        uploading={uploading}
        uploadUrl={uploadUrl}
        uploadType={uploadType}
        setUploadUrl={setUploadUrl}
        setUploadType={setUploadType}
        onUploadModalClose={onUploadModalClose}
        onUpload={onUpload}
        visible={uploadModalVisibility}
        onCancel={onUploadModalCancel}
      />
    </>
  );
};

export default UploadAsset;
