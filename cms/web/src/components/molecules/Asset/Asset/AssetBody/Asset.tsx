import styled from "@emotion/styled";
import { Viewer as CesiumViewer } from "cesium";
import { useCallback, useState } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import DownloadButton from "@reearth-cms/components/atoms/DownloadButton";
import Icon from "@reearth-cms/components/atoms/Icon";
import { DefaultOptionType } from "@reearth-cms/components/atoms/Select";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";
import { Asset, AssetItem, ViewerType } from "@reearth-cms/components/molecules/Asset/asset.type";
import Card from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/card";
import PreviewToolbar from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/previewToolbar";
import {
  PreviewType,
  PreviewTypeSelect,
} from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/previewTypeSelect";
import SideBarCard from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/sideBarCard";
import UnzipFileList from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/UnzipFileList";
import ViewerNotSupported from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/viewerNotSupported";
import ArchiveExtractionStatus from "@reearth-cms/components/molecules/Asset/AssetListTable/ArchiveExtractionStatus";
import {
  GeoViewer,
  Geo3dViewer,
  SvgViewer,
  ImageViewer,
  GltfViewer,
  MvtViewer,
} from "@reearth-cms/components/molecules/Asset/Viewers";
import { useT } from "@reearth-cms/i18n";
import { dateTimeFormat } from "@reearth-cms/utils/format";

import useHooks from "./hooks";

type Props = {
  asset: Asset;
  assetFileExt?: string;
  selectedPreviewType: PreviewType;
  isModalVisible: boolean;
  viewerType: ViewerType;
  displayUnzipFileList: boolean;
  decompressing: boolean;
  onAssetItemSelect: (item: AssetItem) => void;
  onAssetDecompress: (assetId: string) => void;
  onModalCancel: () => void;
  onTypeChange: (
    value: PreviewType,
    option: DefaultOptionType | DefaultOptionType[],
  ) => void | undefined;
  onChangeToFullScreen: () => void;
};

export let viewerRef: CesiumViewer | undefined;

const AssetMolecule: React.FC<Props> = ({
  asset,
  assetFileExt,
  selectedPreviewType,
  isModalVisible,
  viewerType,
  displayUnzipFileList,
  decompressing,
  onAssetItemSelect,
  onAssetDecompress,
  onTypeChange,
  onModalCancel,
  onChangeToFullScreen,
}) => {
  const t = useT();
  const { svgRender, handleCodeSourceClick, handleRenderClick } = useHooks();
  const [assetUrl, setAssetUrl] = useState(asset.url);
  const assetBaseUrl = asset.url.slice(0, asset.url.lastIndexOf("/"));
  const formattedCreatedAt = dateTimeFormat(asset.createdAt);

  const getViewer = (viewer: CesiumViewer | undefined) => {
    viewerRef = viewer;
  };

  const renderPreview = useCallback(() => {
    switch (true) {
      case viewerType === "geo":
        return <GeoViewer url={assetUrl} assetFileExt={assetFileExt} onGetViewer={getViewer} />;
      case viewerType === "geo_3d_tiles":
        return <Geo3dViewer url={assetUrl} setAssetUrl={setAssetUrl} onGetViewer={getViewer} />;
      case viewerType === "geo_mvt":
        return <MvtViewer url={assetUrl} onGetViewer={getViewer} />;
      case viewerType === "image":
        return <ImageViewer url={assetUrl} />;
      case viewerType === "image_svg":
        return <SvgViewer url={assetUrl} svgRender={svgRender} />;
      case viewerType === "model_3d":
        return <GltfViewer url={assetUrl} onGetViewer={getViewer} />;
      case viewerType === "unknown":
      default:
        return <ViewerNotSupported />;
    }
  }, [assetFileExt, assetUrl, svgRender, viewerType]);

  return (
    <BodyContainer>
      <BodyWrapper>
        <Card
          title={asset.fileName}
          toolbar={
            <PreviewToolbar
              url={assetUrl}
              isModalVisible={isModalVisible}
              viewerType={viewerType}
              handleCodeSourceClick={handleCodeSourceClick}
              handleRenderClick={handleRenderClick}
              handleFullScreen={onChangeToFullScreen}
              handleModalCancel={onModalCancel}
            />
          }>
          {renderPreview()}
        </Card>
        {displayUnzipFileList && asset.file && (
          <Card
            title={
              <>
                {asset.fileName}{" "}
                <CopyIcon
                  icon="copy"
                  onClick={() => {
                    navigator.clipboard.writeText(asset.url);
                  }}
                />
              </>
            }
            toolbar={
              <>
                <ArchiveExtractionStatus archiveExtractionStatus={asset.archiveExtractionStatus} />
                {asset.archiveExtractionStatus === "SKIPPED" && (
                  <UnzipButton
                    onClick={() => {
                      onAssetDecompress(asset.id);
                    }}
                    loading={decompressing}
                    icon={<Icon icon="unzip" />}>
                    {t("Unzip")}
                  </UnzipButton>
                )}
              </>
            }>
            <UnzipFileList
              file={asset.file}
              assetBaseUrl={assetBaseUrl}
              archiveExtractionStatus={asset.archiveExtractionStatus}
              setAssetUrl={setAssetUrl}
            />
          </Card>
        )}
        <DownloadButton type="ghost" selected={asset ? [asset] : undefined} displayDefaultIcon />
      </BodyWrapper>
      <SideBarWrapper>
        <SideBarCard title={t("Asset Type")}>
          <PreviewTypeSelect
            style={{ width: "75%" }}
            value={selectedPreviewType}
            onTypeChange={onTypeChange}
          />
        </SideBarCard>
        <SideBarCard title={t("Created Time")}>{formattedCreatedAt}</SideBarCard>
        <SideBarCard title={t("Created By")}>
          <UserAvatar username={asset.createdBy} shadow />
        </SideBarCard>
        <SideBarCard title={t("Linked to")}>
          {asset.items.map(item => (
            <div key={item.itemId}>
              <Button style={{ padding: 0 }} type="link" onClick={() => onAssetItemSelect(item)}>
                {item.itemId}
              </Button>
            </div>
          ))}
        </SideBarCard>
      </SideBarWrapper>
    </BodyContainer>
  );
};

const CopyIcon = styled(Icon)<{ selected?: boolean }>`
  margin-left: 16px;
  &:active {
    color: #096dd9;
  }
`;

const UnzipButton = styled(Button)`
  margin-left: 24px;
`;

const BodyContainer = styled.div`
  display: flex;
  flex-direction: row;
  width: 100%;
  height: calc(100% - 72px);
  .ant-tree-show-line .ant-tree-switcher {
    background-color: transparent;
  }
`;

const BodyWrapper = styled.div`
  padding: 16px 24px;
  width: 70%;
  height: 100%;
  overflow-y: auto;
  flex: 1;
`;

const SideBarWrapper = styled.div`
  padding: 8px;
  width: 272px;
`;

export default AssetMolecule;
