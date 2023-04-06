import styled from "@emotion/styled";
import { useEffect } from "react";
import { Link } from "react-router-dom";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import Tooltip from "@reearth-cms/components/atoms/Tooltip";
import { UploadProps, UploadFile } from "@reearth-cms/components/atoms/Upload";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import LinkAssetModal from "@reearth-cms/components/molecules/Common/LinkAssetModal/LinkAssetModal";
import {
  AssetSortType,
  SortDirection,
} from "@reearth-cms/components/organisms/Asset/AssetList/hooks";
import { useT } from "@reearth-cms/i18n";

import useHooks from "./hooks";

type Props = {
  assetList: Asset[];
  fileList: UploadFile[];
  value?: string;
  loadingAssets: boolean;
  uploading: boolean;
  uploadModalVisibility: boolean;
  uploadUrl: { url: string; autoUnzip: boolean };
  uploadType: UploadType;
  totalCount: number;
  page: number;
  pageSize: number;
  onAssetTableChange: (
    page: number,
    pageSize: number,
    sorter?: { type?: AssetSortType; direction?: SortDirection },
  ) => void;
  onUploadModalCancel: () => void;
  setUploadUrl: (uploadUrl: { url: string; autoUnzip: boolean }) => void;
  setUploadType: (type: UploadType) => void;
  onAssetsCreate: (files: UploadFile[]) => Promise<(Asset | undefined)[]>;
  onAssetCreateFromUrl: (url: string, autoUnzip: boolean) => Promise<Asset | undefined>;
  onAssetsReload: () => void;
  onAssetSearchTerm: (term?: string | undefined) => void;
  setFileList: (fileList: UploadFile<File>[]) => void;
  setUploadModalVisibility: (visible: boolean) => void;
  onChange?: (value: string) => void;
  disabled?: boolean;
};

const AssetItem: React.FC<Props> = ({
  assetList,
  fileList,
  value,
  loadingAssets,
  uploading,
  uploadModalVisibility,
  uploadUrl,
  uploadType,
  totalCount,
  page,
  pageSize,
  onAssetTableChange,
  onUploadModalCancel,
  setUploadUrl,
  setUploadType,
  onAssetsCreate,
  onAssetCreateFromUrl,
  onAssetsReload,
  onAssetSearchTerm,
  setFileList,
  setUploadModalVisibility,
  onChange,
  disabled,
}) => {
  const t = useT();
  const {
    asset,
    loading,
    visible,
    workspaceId,
    projectId,
    handleClick,
    handleLinkAssetModalCancel,
    displayUploadModal,
    handleUploadAndLink,
  } = useHooks(
    fileList,
    uploadUrl,
    uploadType,
    onAssetsCreate,
    onAssetCreateFromUrl,
    setUploadModalVisibility,
    onChange,
    value,
  );

  const uploadProps: UploadProps = {
    name: "file",
    multiple: false,
    maxCount: 1,
    directory: false,
    showUploadList: true,
    accept: "*",
    listType: "picture",
    onRemove: _file => {
      setFileList([]);
    },
    beforeUpload: file => {
      setFileList([file]);
      return false;
    },
    fileList,
  };

  useEffect(() => {
    if (Array.isArray(value)) onChange?.("");
  }, [onChange, value]);

  return (
    <AssetWrapper>
      {value ? (
        <>
          <AssetDetailsWrapper>
            <AssetButton enabled={!!asset} disabled={disabled} onClick={handleClick}>
              <div>
                <Icon icon="folder" size={24} />
                <div style={{ marginTop: 8, overflow: "hidden" }}>{asset?.fileName ?? value}</div>
              </div>
            </AssetButton>
            <Tooltip title={asset?.fileName}>
              <Link
                to={`/workspace/${workspaceId}/project/${projectId}/asset/${value}`}
                target="_blank">
                <AssetLinkedName enabled={!!asset} type="link">
                  {asset?.fileName ?? value + " (removed)"}
                </AssetLinkedName>
              </Link>
            </Tooltip>
          </AssetDetailsWrapper>
          <Space />
          {asset && (
            <Link
              to={`/workspace/${workspaceId}/project/${projectId}/asset/${value}`}
              target="_blank">
              <AssetLink type="link" icon={<Icon icon="arrowSquareOut" size={20} />} />
            </Link>
          )}
          {value && (
            <AssetLink
              type="link"
              icon={<Icon icon={"unlinkSolid"} size={16} />}
              onClick={() => onChange?.("")}
            />
          )}
        </>
      ) : (
        <AssetButton disabled={disabled} onClick={handleClick}>
          <div>
            {loading ? <Icon icon="loading" size={24} /> : <Icon icon="linkSolid" size={14} />}
            <div style={{ marginTop: 4 }}>{t("Asset")}</div>
          </div>
        </AssetButton>
      )}
      <LinkAssetModal
        visible={visible}
        onLinkAssetModalCancel={handleLinkAssetModalCancel}
        linkedAsset={asset}
        assetList={assetList}
        fileList={fileList}
        loading={loadingAssets}
        uploading={uploading}
        uploadProps={uploadProps}
        uploadModalVisibility={uploadModalVisibility}
        uploadUrl={uploadUrl}
        uploadType={uploadType}
        totalCount={totalCount}
        page={page}
        pageSize={pageSize}
        onAssetTableChange={onAssetTableChange}
        setUploadUrl={setUploadUrl}
        setUploadType={setUploadType}
        onChange={onChange}
        onAssetsReload={onAssetsReload}
        onSearchTerm={onAssetSearchTerm}
        displayUploadModal={displayUploadModal}
        onUploadModalCancel={onUploadModalCancel}
        onUploadAndLink={handleUploadAndLink}
      />
    </AssetWrapper>
  );
};

const AssetButton = styled(Button)<{ enabled?: boolean }>`
  width: 100px;
  height: 100px;
  border: 1px dashed;
  border-color: ${({ enabled }) => (enabled ? "#d9d9d9" : "#00000040")};
  color: ${({ enabled }) => (enabled ? "#000000D9" : "#00000040")};
`;

const Space = styled.div`
  flex: 1;
`;

const AssetWrapper = styled.div`
  display: flex;
  align-items: center;
  width: 100%;
`;

const AssetLink = styled(Button)`
  color: #000000d9;
  margin-top: 4px;
  top: 3px;
  &:disabled {
    cursor: pointer;
    pointer-events: auto;
  }
`;

const AssetLinkedName = styled(Button)<{ enabled?: boolean }>`
  color: #1890ff;
  color: ${({ enabled }) => (enabled ? "#1890ff" : "#00000040")};
  margin-left: 12px;
  span {
    text-align: start;
    white-space: normal;
    display: -webkit-box;
    -webkit-line-clamp: 1;
    -webkit-box-orient: vertical;
    overflow: hidden;
    word-break: break-all;
  }
`;

const AssetDetailsWrapper = styled.div`
  display: flex;
  align-items: center;
`;

export default AssetItem;
