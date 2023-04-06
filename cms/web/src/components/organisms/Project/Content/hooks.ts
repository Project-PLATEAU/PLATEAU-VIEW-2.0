import { useCallback, useEffect, useMemo, useState } from "react";

import Notification from "@reearth-cms/components/atoms/Notification";
import { Request } from "@reearth-cms/components/molecules/Request/types";
import { convertRequest } from "@reearth-cms/components/organisms/Project/Request/convertRequest";
import {
  useUpdateRequestMutation,
  RequestState as GQLRequestState,
  Request as GQLRequest,
  useGetModalRequestsQuery,
  useUnpublishItemMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useModel, useProject, useWorkspace } from "@reearth-cms/state";

export default () => {
  const [currentModel] = useModel();
  const [currentWorkspace] = useWorkspace();
  const [currentProject] = useProject();
  const [addItemToRequestModalShown, setAddItemToRequestModalShown] = useState(false);
  const t = useT();

  const [page, setPage] = useState<number>(1);
  const [pageSize, setPageSize] = useState<number>(10);

  useEffect(() => {
    setPage(+page);
    setPageSize(+pageSize);
  }, [setPage, setPageSize, page, pageSize]);

  const { data, loading } = useGetModalRequestsQuery({
    fetchPolicy: "no-cache",
    variables: {
      projectId: currentProject?.id ?? "",
      pagination: { first: pageSize, offset: (page - 1) * pageSize },
      sort: { key: "createdAt", reverted: true },
      state: ["WAITING"] as GQLRequestState[],
    },
    skip: !currentProject?.id,
  });

  const requests: Request[] = useMemo(
    () =>
      (data?.requests.nodes
        .map(request => request as GQLRequest)
        .map(convertRequest)
        .filter(request => !!request) as Request[]) ?? [],
    [data?.requests.nodes],
  );

  const [updateRequest] = useUpdateRequestMutation();

  const handleAddItemToRequest = useCallback(
    async (request: Request, itemIds: string[]) => {
      const item = await updateRequest({
        variables: {
          requestId: request.id,
          description: request.description,
          items: [
            ...new Set([...request.items.map(item => item.id), ...itemIds.map(itemId => itemId)]),
          ].map(itemId => ({ itemId })),
          reviewersId: request.reviewers.map(reviewer => reviewer.id),
          title: request.title,
          state: request.state as GQLRequestState,
        },
      });
      if (item.errors || !item.data?.updateRequest) {
        Notification.error({ message: t("Failed to update request.") });
        return;
      }

      Notification.success({ message: t("Successfully updated Request!") });
    },
    [updateRequest, t],
  );

  const [unpublishItem] = useUnpublishItemMutation();

  const handleUnpublish = useCallback(
    async (itemIds: string[]) => {
      const item = await unpublishItem({
        variables: {
          itemId: itemIds,
        },
        refetchQueries: ["SearchItem", "GetItem"],
      });
      if (item.errors || !item.data?.unpublishItem) {
        Notification.error({ message: t("Failed to unpublish items.") });
        return;
      }

      Notification.success({ message: t("Successfully unpublished items!") });
    },
    [unpublishItem, t],
  );

  const handleAddItemToRequestModalClose = useCallback(
    () => setAddItemToRequestModalShown(false),
    [],
  );

  const handleAddItemToRequestModalOpen = useCallback(() => {
    setPage(1);
    setPageSize(10);
    setAddItemToRequestModalShown(true);
  }, []);

  const handleRequestTableChange = useCallback((page: number, pageSize: number) => {
    setPage(page);
    setPageSize(pageSize);
  }, []);

  return {
    currentWorkspace,
    currentModel,
    currentProject,
    requests,
    addItemToRequestModalShown,
    handleUnpublish,
    handleRequestTableChange,
    handleAddItemToRequest,
    handleAddItemToRequestModalClose,
    handleAddItemToRequestModalOpen,
    loading,
    totalCount: data?.requests.totalCount ?? 0,
    page,
    pageSize,
  };
};
