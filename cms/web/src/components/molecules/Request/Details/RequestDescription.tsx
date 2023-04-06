import styled from "@emotion/styled";
import moment from "moment";
import { useMemo } from "react";

import Collapse from "@reearth-cms/components/atoms/Collapse";
import AntDComment from "@reearth-cms/components/atoms/Comment";
import Tooltip from "@reearth-cms/components/atoms/Tooltip";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import { Request } from "@reearth-cms/components/molecules/Request/types";
import {
  AssetSortType,
  SortDirection,
} from "@reearth-cms/components/organisms/Asset/AssetList/hooks";

import RequestItemForm from "./ItemForm";

const { Panel } = Collapse;

type Props = {
  currentRequest: Request;
  assetList: Asset[];
  fileList: UploadFile[];
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
};

export const RequestDescription: React.FC<Props> = ({
  currentRequest,
  assetList,
  fileList,
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
}) => {
  const fromNow = useMemo(
    () => moment(currentRequest.createdAt?.toString()).fromNow(),
    [currentRequest.createdAt],
  );

  return (
    <StyledAntDComment
      author={<a>{currentRequest.createdBy?.name}</a>}
      avatar={<UserAvatar username={currentRequest.createdBy?.name} />}
      content={
        <>
          <RequestTextWrapper>
            <RequestTitle>{currentRequest.title}</RequestTitle>
            <RequestText>{currentRequest.description}</RequestText>
          </RequestTextWrapper>
          <RequestItemsWrapper>
            <Collapse>
              {currentRequest.items
                .filter(item => item.schema)
                .map((item, index) => (
                  <Panel header={item.modelName} key={index}>
                    <RequestItemForm
                      key={index}
                      schema={item.schema}
                      initialFormValues={item.initialValues}
                      assetList={assetList}
                      fileList={fileList}
                      loadingAssets={loadingAssets}
                      uploading={uploading}
                      uploadModalVisibility={uploadModalVisibility}
                      uploadUrl={uploadUrl}
                      uploadType={uploadType}
                      totalCount={totalCount}
                      page={page}
                      pageSize={pageSize}
                      onAssetTableChange={onAssetTableChange}
                      onUploadModalCancel={onUploadModalCancel}
                      setUploadUrl={setUploadUrl}
                      setUploadType={setUploadType}
                      onAssetsCreate={onAssetsCreate}
                      onAssetCreateFromUrl={onAssetCreateFromUrl}
                      onAssetsReload={onAssetsReload}
                      onAssetSearchTerm={onAssetSearchTerm}
                      setFileList={setFileList}
                      setUploadModalVisibility={setUploadModalVisibility}
                    />
                  </Panel>
                ))}
            </Collapse>
          </RequestItemsWrapper>
        </>
      }
      datetime={
        currentRequest.createdAt && (
          <Tooltip title={currentRequest.createdAt.toString()}>
            <span>{fromNow}</span>
          </Tooltip>
        )
      }
    />
  );
};

const StyledAntDComment = styled(AntDComment)`
  .ant-comment-content-author {
    padding: 16px 24px;
    margin: 0;
    border-bottom: 1px solid #f0f0f0;
    .ant-comment-content-author-name {
      font-weight: 400;
      font-size: 14px;
      color: #00000073;
    }
  }
  .ant-comment-inner {
    padding-top: 0;
  }
  .ant-comment-content {
    background-color: #fff;
  }
`;

const RequestTitle = styled.h1`
  border-bottom: 1px solid #f0f0f0;
  padding: 8px 0;
  color: #000000d9;
`;

const RequestTextWrapper = styled.div`
  padding: 24px;
  border-bottom: 1px solid #f0f0f0;
`;

const RequestText = styled.p`
  padding-top: 8px;
`;

const RequestItemsWrapper = styled.div`
  padding: 12px;
  .ant-pro-card-body {
    padding: 0;
  }
`;
