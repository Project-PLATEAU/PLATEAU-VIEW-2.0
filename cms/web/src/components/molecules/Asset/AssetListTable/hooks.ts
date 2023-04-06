import { useEffect, useState } from "react";

import { ArchiveExtractionStatus } from "@reearth-cms/components/molecules/Asset/asset.type";
import { useT } from "@reearth-cms/i18n";

export default (archiveExtractionStatus: ArchiveExtractionStatus) => {
  const t = useT();

  const [status, setStatus] = useState<string>("");
  const [statusColor, setStatusColor] = useState<string>("");

  const DECOMPRESSED = t("Decompressed");
  const FAILED = t("Failed");
  const DECOMPRESSING = t("Decompressing");
  const SKIPPED = t("Skipped");
  const PENDING = t("Pending");

  useEffect(() => {
    switch (archiveExtractionStatus) {
      case "DONE":
        setStatusColor("#52C41A");
        setStatus(DECOMPRESSED);
        break;
      case "FAILED":
        setStatusColor("#F5222D");
        setStatus(FAILED);
        break;
      case "IN_PROGRESS":
        setStatusColor("#FA8C16");
        setStatus(DECOMPRESSING);
        break;
      case "SKIPPED":
        setStatusColor("#BFBFBF");
        setStatus(SKIPPED);
        break;
      case "PENDING":
        setStatusColor("#BFBFBF");
        setStatus(PENDING);
        break;
      default:
        setStatusColor("");
        setStatus("");
        break;
    }
  }, [DECOMPRESSED, DECOMPRESSING, FAILED, SKIPPED, PENDING, archiveExtractionStatus, t]);

  return { status, statusColor };
};
