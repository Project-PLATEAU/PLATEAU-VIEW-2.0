import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import { ViewerType } from "@reearth-cms/components/molecules/Asset/asset.type";
import PreviewModal from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/previewModal";
import { useT } from "@reearth-cms/i18n";

type Props = {
  url: string;
  isModalVisible: boolean;
  viewerType?: ViewerType;
  handleCodeSourceClick: () => void;
  handleRenderClick: () => void;
  handleFullScreen: () => void;
  handleModalCancel: () => void;
};

const PreviewToolbar: React.FC<Props> = ({
  url,
  isModalVisible,
  viewerType,
  handleCodeSourceClick,
  handleRenderClick,
  handleFullScreen,
  handleModalCancel,
}) => {
  const t = useT();
  const isSVGButtonVisible = viewerType === "image_svg";
  const isFullScreenButtonVisible = viewerType !== "unknown";

  return (
    <>
      {isSVGButtonVisible && (
        <>
          <Button onClick={handleCodeSourceClick}>{t("Source Code")}</Button>
          <Button onClick={handleRenderClick}>{t("Render")}</Button>
        </>
      )}
      {isFullScreenButtonVisible && (
        <Button
          type="link"
          icon={<Icon icon="fullscreen" />}
          size="large"
          onClick={handleFullScreen}
        />
      )}
      <PreviewModal url={url} visible={isModalVisible} handleCancel={handleModalCancel} />
    </>
  );
};

export default PreviewToolbar;
