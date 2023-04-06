import {
  ArchiveExtractionStatus,
  Asset,
  Comment,
} from "@reearth-cms/components/molecules/Asset/asset.type";
import { Asset as GQLAsset, Comment as GQLComment } from "@reearth-cms/gql/graphql-client-api";

import { fileName } from "./utils";

export const convertAsset = (asset: GQLAsset | undefined): Asset | undefined => {
  if (!asset) return;
  return {
    id: asset.id,
    fileName: fileName(asset.url),
    createdAt: asset.createdAt.toString(),
    createdBy: asset.createdBy?.name ?? "",
    createdByType: asset.createdByType,
    previewType: asset.previewType || "UNKNOWN",
    projectId: asset.projectId,
    size: asset.size,
    url: asset.url,
    threadId: asset.thread?.id ?? "",
    comments: asset.thread?.comments?.map(comment => convertComment(comment)) ?? [],
    archiveExtractionStatus: asset.archiveExtractionStatus as ArchiveExtractionStatus,
    items: asset.items ?? [],
  };
};

export const convertComment = (GQLComment: GQLComment): Comment => {
  return {
    id: GQLComment.id,
    author: {
      id: GQLComment.author?.id,
      name: GQLComment.author?.name ?? "Anonymous",
      type: GQLComment.author
        ? GQLComment.author.__typename === "User"
          ? "User"
          : "Integration"
        : null,
    },
    content: GQLComment.content,
    createdAt: GQLComment.createdAt.toString(),
  };
};
