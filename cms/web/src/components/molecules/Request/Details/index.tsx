import RequestMolecule from "@reearth-cms//components/molecules/Request/Details/Request";
import Loading from "@reearth-cms/components/atoms/Loading";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { User } from "@reearth-cms/components/molecules/AccountSettings/types";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import { Request, RequestUpdatePayload } from "@reearth-cms/components/molecules/Request/types";
import { Member } from "@reearth-cms/components/molecules/Workspace/types";
import {
  AssetSortType,
  SortDirection,
} from "@reearth-cms/components/organisms/Asset/AssetList/hooks";

export type Props = {
  me?: User;
  isCloseActionEnabled: boolean;
  isApproveActionEnabled: boolean;
  currentRequest?: Request;
  workspaceUserMembers: Member[];
  onRequestApprove: (requestId: string) => Promise<void>;
  onRequestUpdate: (data: RequestUpdatePayload) => Promise<void>;
  onRequestDelete: (requestsId: string[]) => Promise<void>;
  onCommentCreate: (content: string) => Promise<void>;
  onCommentUpdate: (commentId: string, content: string) => Promise<void>;
  onCommentDelete: (commentId: string) => Promise<void>;
  onBack: () => void;
  assetList: Asset[];
  fileList: UploadFile[];
  loadingAssets: boolean;
  loading: boolean;
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

const RequestDetailsMolecule: React.FC<Props> = ({
  me,
  isCloseActionEnabled,
  isApproveActionEnabled,
  currentRequest,
  workspaceUserMembers,
  onRequestApprove,
  onRequestUpdate,
  onRequestDelete,
  onCommentCreate,
  onCommentUpdate,
  onCommentDelete,
  onBack,
  assetList,
  fileList,
  loadingAssets,
  loading,
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
  return currentRequest ? (
    loading ? (
      <Loading spinnerSize="large" minHeight="100vh" />
    ) : (
      <RequestMolecule
        me={me}
        isCloseActionEnabled={isCloseActionEnabled}
        isApproveActionEnabled={isApproveActionEnabled}
        currentRequest={currentRequest}
        workspaceUserMembers={workspaceUserMembers}
        onRequestApprove={onRequestApprove}
        onRequestUpdate={onRequestUpdate}
        onRequestDelete={onRequestDelete}
        onCommentCreate={onCommentCreate}
        onCommentUpdate={onCommentUpdate}
        onCommentDelete={onCommentDelete}
        onBack={onBack}
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
    )
  ) : null;
};
export default RequestDetailsMolecule;
