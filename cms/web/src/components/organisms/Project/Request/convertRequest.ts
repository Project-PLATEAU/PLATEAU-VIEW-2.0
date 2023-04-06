import { Request, Comment } from "@reearth-cms/components/molecules/Request/types";
import {
  Request as GQLRequest,
  Comment as GQLComment,
  ItemField,
} from "@reearth-cms/gql/graphql-client-api";

export const convertRequest = (request: GQLRequest | undefined): Request | undefined => {
  if (!request) return;
  return {
    id: request.id,
    threadId: request.thread?.id ?? "",
    title: request.title,
    description: request.description ?? "",
    comments: request.thread?.comments?.map(comment => convertComment(comment)) ?? [],
    createdAt: request.createdAt,
    reviewers: request.reviewers,
    state: request.state,
    createdBy: request.createdBy ?? undefined,
    updatedAt: request.updatedAt,
    approvedAt: request.approvedAt ?? undefined,
    closedAt: request.closedAt ?? undefined,
    items: request.items?.map(item => ({
      id: item.itemId,
      modelName: item?.item?.value.model.name,
      initialValues: getInitialFormValues(item.item?.value.fields),
      schema: item.item?.value.schema ? item.item?.value.schema : undefined,
    })),
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

export const getInitialFormValues = (fields?: ItemField[]): { [key: string]: any } => {
  const initialValues: { [key: string]: any } = {};
  fields?.forEach(field => {
    initialValues[field.schemaFieldId] = field.value;
  });
  return initialValues;
};
