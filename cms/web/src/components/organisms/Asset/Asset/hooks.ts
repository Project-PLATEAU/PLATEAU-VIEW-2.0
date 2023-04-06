import { useCallback, useEffect, useMemo, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import {
  Asset,
  AssetItem,
  PreviewType,
  ViewerType,
  AssetFile,
} from "@reearth-cms/components/molecules/Asset/asset.type";
import { viewerRef } from "@reearth-cms/components/molecules/Asset/Asset/AssetBody/Asset";
import {
  geoFormats,
  geo3dFormats,
  geoMvtFormat,
  model3dFormats,
  imageFormats,
  imageSVGFormat,
  compressedFileFormats,
} from "@reearth-cms/components/molecules/Common/Asset";
import {
  Asset as GQLAsset,
  PreviewType as GQLPreviewType,
  useDecompressAssetMutation,
  useGetAssetFileQuery,
  useGetAssetItemQuery,
  useGetAssetQuery,
  useUpdateAssetMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { getExtension } from "@reearth-cms/utils/file";

import { convertAsset } from "../convertAsset";

export default (assetId?: string) => {
  const t = useT();
  const navigate = useNavigate();
  const { workspaceId, projectId } = useParams();
  const [selectedPreviewType, setSelectedPreviewType] = useState<PreviewType>("IMAGE");
  const [decompressing, setDecompressing] = useState(false);
  const [isModalVisible, setIsModalVisible] = useState<boolean>(false);
  const [collapsed, setCollapsed] = useState(true);

  const { data: rawAsset, loading } = useGetAssetQuery({
    variables: {
      assetId: assetId ?? "",
    },
    fetchPolicy: "network-only",
  });

  const { data: rawFile, loading: loading2 } = useGetAssetFileQuery({
    variables: {
      assetId: assetId ?? "",
    },
    fetchPolicy: "network-only",
  });

  const { data: rawAssetItem } = useGetAssetItemQuery({
    variables: {
      assetId: assetId ?? "",
    },
  });

  const convertedAsset: Asset | undefined = useMemo(() => {
    return rawAsset?.node?.__typename === "Asset"
      ? convertAsset(rawAsset.node as GQLAsset)
      : undefined;
  }, [rawAsset]);

  const asset = useMemo(() => {
    return convertedAsset
      ? {
          ...convertedAsset,
          ...(convertedAsset && rawAssetItem?.node?.__typename === "Asset"
            ? {
                items: rawAssetItem?.node.items ?? [],
              }
            : {}),
          ...(convertedAsset && rawFile?.assetFile
            ? {
                file: rawFile?.assetFile as AssetFile,
              }
            : {}),
        }
      : undefined;
  }, [convertedAsset, rawAssetItem?.node, rawFile?.assetFile]);

  const [updateAssetMutation] = useUpdateAssetMutation();
  const handleAssetUpdate = useCallback(
    (assetId: string, previewType?: PreviewType) =>
      (async () => {
        if (!assetId) return;
        const result = await updateAssetMutation({
          variables: { id: assetId, previewType: previewType as GQLPreviewType },
          refetchQueries: ["GetAsset"],
        });
        if (result.errors || !result.data?.updateAsset) {
          Notification.error({ message: t("Failed to update asset.") });
        }
        if (result) {
          Notification.success({ message: t("Asset was successfully updated!") });
        }
      })(),
    [t, updateAssetMutation],
  );

  const [decompressAssetMutation] = useDecompressAssetMutation();
  const handleAssetDecompress = useCallback(
    (assetId: string) =>
      (async () => {
        if (!assetId) return;
        setDecompressing(true);
        const result = await decompressAssetMutation({
          variables: { assetId },
          refetchQueries: ["GetAsset"],
        });
        setDecompressing(false);
        if (result.errors || !result.data?.decompressAsset) {
          Notification.error({ message: t("Failed to decompress asset.") });
        }
        if (result) {
          Notification.success({ message: t("Asset is being decompressed!") });
        }
      })(),
    [t, decompressAssetMutation, setDecompressing],
  );

  useEffect(() => {
    if (convertedAsset?.previewType) {
      setSelectedPreviewType(convertedAsset.previewType);
    }
  }, [convertedAsset?.previewType]);

  const handleTypeChange = useCallback((value: PreviewType) => {
    setSelectedPreviewType(value);
  }, []);

  const [viewerType, setViewerType] = useState<ViewerType>("unknown");
  const assetFileExt = getExtension(convertedAsset?.fileName);

  useEffect(() => {
    switch (true) {
      case selectedPreviewType === "GEO" &&
        (geoFormats.includes(assetFileExt) || compressedFileFormats.includes(assetFileExt)):
        setViewerType("geo");
        break;
      case selectedPreviewType === "GEO_3D_TILES" &&
        (geo3dFormats.includes(assetFileExt) || compressedFileFormats.includes(assetFileExt)):
        setViewerType("geo_3d_tiles");
        break;
      case selectedPreviewType === "GEO_MVT" &&
        (geoMvtFormat.includes(assetFileExt) || compressedFileFormats.includes(assetFileExt)):
        setViewerType("geo_mvt");
        break;
      case selectedPreviewType === "MODEL_3D" &&
        (model3dFormats.includes(assetFileExt) || compressedFileFormats.includes(assetFileExt)):
        setViewerType("model_3d");
        break;
      case selectedPreviewType === "IMAGE" && imageFormats.includes(assetFileExt):
        setViewerType("image");
        break;
      case selectedPreviewType === "IMAGE_SVG" && imageSVGFormat.includes(assetFileExt):
        setViewerType("image_svg");
        break;
      default:
        setViewerType("unknown");
        break;
    }
  }, [convertedAsset?.previewType, assetFileExt, selectedPreviewType]);

  const displayUnzipFileList = useMemo(
    () => compressedFileFormats.includes(assetFileExt),
    [assetFileExt],
  );

  const handleFullScreen = useCallback(() => {
    if (
      viewerType === "geo" ||
      viewerType === "geo_3d_tiles" ||
      viewerType === "model_3d" ||
      viewerType === "geo_mvt"
    ) {
      viewerRef?.canvas.requestFullscreen();
    } else if (viewerType === "image" || viewerType === "image_svg") {
      setIsModalVisible(true);
    }
  }, [viewerType]);

  const handleAssetItemSelect = useCallback(
    (assetItem: AssetItem) => {
      navigate(
        `/workspace/${workspaceId}/project/${projectId}/content/${assetItem.modelId}/details/${assetItem.itemId}`,
      );
    },
    [navigate, projectId, workspaceId],
  );

  const handleModalCancel = useCallback(() => {
    setIsModalVisible(false);
  }, []);

  const handleToggleCommentMenu = useCallback(
    (value: boolean) => {
      setCollapsed(value);
    },
    [setCollapsed],
  );

  return {
    asset,
    assetFileExt,
    isLoading: loading || loading2,
    selectedPreviewType,
    isModalVisible,
    collapsed,
    viewerType,
    displayUnzipFileList,
    decompressing,
    handleAssetItemSelect,
    handleAssetDecompress,
    handleToggleCommentMenu,
    handleAssetUpdate,
    handleTypeChange,
    handleModalCancel,
    handleFullScreen,
  };
};
