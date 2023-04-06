import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import {
  AssetSortType,
  SortDirection,
} from "@reearth-cms/components/organisms/Asset/AssetList/hooks";

import { FieldType } from "../../types";

import AssetField from "./AssetField";
import BooleanField from "./BooleanField";
import IntegerField from "./IntegerField";
import MarkdownField from "./Markdown";
import SelectField from "./SelectField";
import TextAreaField from "./TextArea";
import TextField from "./TextField";
import URLField from "./URLField";

export interface Props {
  selectedType: FieldType;
  multiple?: boolean;
  selectedValues: string[];
  assetList: Asset[];
  fileList: UploadFile[];
  loadingAssets: boolean;
  uploading: boolean;
  defaultValue?: string;
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
  onAssetSearchTerm: (term?: string | undefined) => void;
  onAssetsReload: () => void;
  setFileList: (fileList: UploadFile<File>[]) => void;
  setUploadModalVisibility: (visible: boolean) => void;
}

const FieldDefaultInputs: React.FC<Props> = ({
  selectedType,
  selectedValues,
  multiple,
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
  onAssetSearchTerm,
  onAssetsReload,
  onAssetsCreate,
  onAssetCreateFromUrl,
  setFileList,
  setUploadModalVisibility,
}) => {
  return selectedType ? (
    selectedType === "TextArea" ? (
      <TextAreaField multiple={multiple} />
    ) : selectedType === "MarkdownText" ? (
      <MarkdownField multiple={multiple} />
    ) : selectedType === "Integer" ? (
      <IntegerField multiple={multiple} />
    ) : selectedType === "Bool" ? (
      <BooleanField multiple={multiple} />
    ) : selectedType === "Asset" ? (
      <AssetField
        multiple={multiple}
        assetList={assetList}
        fileList={fileList}
        loadingAssets={loadingAssets}
        uploading={uploading}
        uploadModalVisibility={uploadModalVisibility}
        uploadUrl={uploadUrl}
        uploadType={uploadType}
        onAssetTableChange={onAssetTableChange}
        totalCount={totalCount}
        page={page}
        pageSize={pageSize}
        onUploadModalCancel={onUploadModalCancel}
        setUploadUrl={setUploadUrl}
        setUploadType={setUploadType}
        onAssetsCreate={onAssetsCreate}
        onAssetCreateFromUrl={onAssetCreateFromUrl}
        onAssetSearchTerm={onAssetSearchTerm}
        onAssetsReload={onAssetsReload}
        setFileList={setFileList}
        setUploadModalVisibility={setUploadModalVisibility}
      />
    ) : selectedType === "Select" ? (
      <SelectField selectedValues={selectedValues} multiple={multiple} />
    ) : selectedType === "URL" ? (
      <URLField multiple={multiple} />
    ) : (
      <TextField multiple={multiple} />
    )
  ) : null;
};

export default FieldDefaultInputs;
