import styled from "@emotion/styled";

import Form from "@reearth-cms/components/atoms/Form";
import Input from "@reearth-cms/components/atoms/Input";
import InputNumber from "@reearth-cms/components/atoms/InputNumber";
import MarkdownInput from "@reearth-cms/components/atoms/Markdown";
import Select from "@reearth-cms/components/atoms/Select";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { Asset } from "@reearth-cms/components/molecules/Asset/asset.type";
import { UploadType } from "@reearth-cms/components/molecules/Asset/AssetList";
import AssetItem from "@reearth-cms/components/molecules/Common/Form/AssetItem";
import MultiValueField from "@reearth-cms/components/molecules/Common/MultiValueField";
import MultiValueAsset from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueAsset";
import MultiValueSelect from "@reearth-cms/components/molecules/Common/MultiValueField/MultiValueSelect";
import FieldTitle from "@reearth-cms/components/molecules/Content/Form/FieldTitle";
import {
  AssetSortType,
  SortDirection,
} from "@reearth-cms/components/organisms/Asset/AssetList/hooks";

export interface Props {
  initialFormValues: any;
  schema?: any;
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
}
const RequestItemForm: React.FC<Props> = ({
  schema,
  initialFormValues,
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
  const { Option } = Select;
  const [form] = Form.useForm();
  return (
    <StyledForm form={form} layout="vertical" initialValues={initialFormValues}>
      <FormItemsWrapper>
        {schema?.fields.map((field: any) =>
          field.type === "TextArea" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  disabled={true}
                  rows={3}
                  showCount
                  maxLength={field.typeProperty.maxLength ?? false}
                  FieldInput={TextArea}
                />
              ) : (
                <TextArea
                  disabled={true}
                  rows={3}
                  showCount
                  maxLength={field.typeProperty.maxLength ?? false}
                />
              )}
            </Form.Item>
          ) : field.type === "MarkdownText" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  disabled={true}
                  maxLength={field.typeProperty.maxLength ?? false}
                  FieldInput={MarkdownInput}
                />
              ) : (
                <MarkdownInput disabled={true} maxLength={field.typeProperty.maxLength ?? false} />
              )}
            </Form.Item>
          ) : field.type === "Integer" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  disabled={true}
                  type="number"
                  min={field.typeProperty.min}
                  max={field.typeProperty.max}
                  FieldInput={InputNumber}
                />
              ) : (
                <InputNumber
                  disabled={true}
                  type="number"
                  min={field.typeProperty.min}
                  max={field.typeProperty.max}
                />
              )}
            </Form.Item>
          ) : field.type === "Asset" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueAsset
                  disabled={true}
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
              ) : (
                <AssetItem
                  disabled={true}
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
              )}
            </Form.Item>
          ) : field.type === "Select" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueSelect disabled={true} selectedValues={field.typeProperty?.values} />
              ) : (
                <Select disabled={true} allowClear>
                  {field.typeProperty?.values?.map((value: string) => (
                    <Option key={value} value={value}>
                      {value}
                    </Option>
                  ))}
                </Select>
              )}
            </Form.Item>
          ) : field.type === "URL" ? (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  disabled={true}
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                  FieldInput={Input}
                />
              ) : (
                <Input
                  disabled={true}
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                />
              )}
            </Form.Item>
          ) : (
            <Form.Item
              extra={field.description}
              name={field.id}
              label={<FieldTitle title={field.title} isUnique={field.unique} />}>
              {field.multiple ? (
                <MultiValueField
                  disabled={true}
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                  FieldInput={Input}
                />
              ) : (
                <Input
                  disabled={true}
                  showCount={true}
                  maxLength={field.typeProperty.maxLength ?? 500}
                />
              )}
            </Form.Item>
          ),
        )}
      </FormItemsWrapper>
    </StyledForm>
  );
};

export default RequestItemForm;

const StyledForm = styled(Form)`
  padding: 16px;
  width: 100%;
  height: 100%;
  overflow-y: auto;
  background: #fff;
`;

const FormItemsWrapper = styled.div`
  width: 50%;
  @media (max-width: 1200px) {
    width: 100%;
  }
`;
