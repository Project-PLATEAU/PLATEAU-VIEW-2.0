import fileDownload from "js-file-download";
import React, { MouseEventHandler, useCallback } from "react";

import Button, { ButtonProps } from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import { useT } from "@reearth-cms/i18n";

export type PreviewType =
  | "GEO"
  | "GEO_3D_TILES"
  | "GEO_MVT"
  | "IMAGE"
  | "IMAGE_SVG"
  | "MODEL_3D"
  | "UNKNOWN";

export type AssetFile = {
  children?: AssetFile[];
  contentType?: string;
  name: string;
  path: string;
  size: number;
};

export type Asset = {
  id: string;
  fileName: string;
  url: string;
};

export type Comment = {
  id: string;
  author: { id?: string; name: string; type: "User" | "Integration" | null };
  content: string;
  createdAt: string;
};

type DownloadButtonProps = {
  title?: string;
  selected?: Asset[];
  displayDefaultIcon?: boolean;
} & ButtonProps;

const DownloadButton: React.FC<DownloadButtonProps> = ({
  title,
  selected,
  displayDefaultIcon,
  ...props
}) => {
  const t = useT();
  const handleDownload: MouseEventHandler<HTMLElement> | undefined = useCallback(async () => {
    selected?.map(async s => {
      const res = await fetch(s.url, {
        method: "GET",
      });
      const blob = await res.blob();
      fileDownload(blob, s.fileName);
    });
  }, [selected]);

  return (
    <Button
      icon={displayDefaultIcon && <Icon icon="download" />}
      onClick={handleDownload}
      disabled={!selected || selected.length <= 0}
      {...props}>
      {title ?? t("Download")}
    </Button>
  );
};

export default DownloadButton;
