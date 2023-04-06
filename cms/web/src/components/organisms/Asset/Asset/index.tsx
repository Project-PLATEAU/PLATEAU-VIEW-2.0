import { useNavigate, useParams } from "react-router-dom";

import Loading from "@reearth-cms/components/atoms/Loading";
import AssetWrapper from "@reearth-cms/components/molecules/Asset/Asset/AssetBody";
import CommentsPanel from "@reearth-cms/components/organisms/Common/CommentsPanel";

import useHooks from "./hooks";

const Asset: React.FC = () => {
  const navigate = useNavigate();
  const { workspaceId, projectId, assetId } = useParams();
  const {
    asset,
    assetFileExt,
    isLoading,
    selectedPreviewType,
    isModalVisible,
    collapsed,
    viewerType,
    displayUnzipFileList,
    decompressing,
    handleAssetDecompress,
    handleAssetItemSelect,
    handleToggleCommentMenu,
    handleAssetUpdate,
    handleTypeChange,
    handleModalCancel,
    handleFullScreen,
  } = useHooks(assetId);

  const handleSave = async () => {
    if (assetId) {
      await handleAssetUpdate(assetId, selectedPreviewType);
    }
  };

  const handleBack = () => {
    navigate(`/workspace/${workspaceId}/project/${projectId}/asset/`);
  };

  return isLoading ? (
    <Loading spinnerSize="large" minHeight="100vh" />
  ) : (
    <AssetWrapper
      commentsPanel={
        <CommentsPanel
          comments={asset?.comments}
          threadId={asset?.threadId}
          collapsed={collapsed}
          onCollapse={handleToggleCommentMenu}
        />
      }
      asset={asset}
      assetFileExt={assetFileExt}
      selectedPreviewType={selectedPreviewType}
      isModalVisible={isModalVisible}
      viewerType={viewerType}
      displayUnzipFileList={displayUnzipFileList}
      decompressing={decompressing}
      onAssetItemSelect={handleAssetItemSelect}
      onAssetDecompress={handleAssetDecompress}
      onTypeChange={handleTypeChange}
      onModalCancel={handleModalCancel}
      onChangeToFullScreen={handleFullScreen}
      onBack={handleBack}
      onSave={handleSave}
    />
  );
};

export default Asset;
