import AssetListBody from "@reearth-cms/components/molecules/Asset/AssetList";
import CommentsPanel from "@reearth-cms/components/organisms/Common/CommentsPanel";
import { useT } from "@reearth-cms/i18n";

import useHooks from "./hooks";

const AssetList: React.FC = () => {
  const t = useT();
  const {
    assetList,
    selection,
    fileList,
    uploading,
    uploadModalVisibility,
    loading,
    uploadUrl,
    uploadType,
    selectedAsset,
    collapsed,
    totalCount,
    searchTerm,
    page,
    pageSize,
    sort,
    handleToggleCommentMenu,
    handleAssetItemSelect,
    handleAssetSelect,
    handleUploadModalCancel,
    setUploadUrl,
    setUploadType,
    setSelection,
    setFileList,
    setUploadModalVisibility,
    handleAssetsCreate,
    handleAssetCreateFromUrl,
    handleAssetDelete,
    handleSearchTerm,
    handleAssetsReload,
    handleNavigateToAsset,
    handleAssetTableChange,
  } = useHooks();

  return (
    <AssetListBody
      commentsPanel={
        <CommentsPanel
          collapsed={collapsed}
          onCollapse={handleToggleCommentMenu}
          emptyText={
            selectedAsset
              ? t("No comments.")
              : t("Please click the comment bubble in the table to check comments.")
          }
          comments={assetList.find(asset => asset.id === selectedAsset?.id)?.comments}
          threadId={assetList.find(asset => asset.id === selectedAsset?.id)?.threadId}
        />
      }
      assetList={assetList}
      onAssetTableChange={handleAssetTableChange}
      totalCount={totalCount}
      searchTerm={searchTerm}
      sort={sort}
      page={page}
      pageSize={pageSize}
      fileList={fileList}
      selection={selection}
      uploading={uploading}
      uploadModalVisibility={uploadModalVisibility}
      loading={loading}
      uploadUrl={uploadUrl}
      uploadType={uploadType}
      onAssetItemSelect={handleAssetItemSelect}
      onAssetSelect={handleAssetSelect}
      onUploadModalCancel={handleUploadModalCancel}
      setUploadUrl={setUploadUrl}
      setUploadType={setUploadType}
      selectedAsset={selectedAsset}
      onAssetsCreate={handleAssetsCreate}
      onAssetCreateFromUrl={handleAssetCreateFromUrl}
      onAssetDelete={handleAssetDelete}
      onAssetsReload={handleAssetsReload}
      onSearchTerm={handleSearchTerm}
      onEdit={handleNavigateToAsset}
      setSelection={setSelection}
      setFileList={setFileList}
      setUploadModalVisibility={setUploadModalVisibility}
    />
  );
};

export default AssetList;
