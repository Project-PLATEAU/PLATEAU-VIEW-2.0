import { useCallback, useEffect, useMemo, useState } from "react";
import { useNavigate, useParams, useSearchParams } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { ProColumns } from "@reearth-cms/components/atoms/ProTable";
import { ContentTableField, ItemStatus } from "@reearth-cms/components/molecules/Content/types";
import { Request } from "@reearth-cms/components/molecules/Request/types";
import {
  convertItem,
  convertComment,
} from "@reearth-cms/components/organisms/Project/Content/convertItem";
import useContentHooks from "@reearth-cms/components/organisms/Project/Content/hooks";
import {
  Item as GQLItem,
  useDeleteItemMutation,
  Comment as GQLComment,
  SortDirection as GQLSortDirection,
  ItemSortType as GQLItemSortType,
  useSearchItemQuery,
  Asset as GQLAsset,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";

import { fileName } from "./utils";

export type ItemSortType = "CREATION_DATE" | "MODIFICATION_DATE";
export type SortDirection = "ASC" | "DESC";

export default () => {
  const {
    currentModel,
    currentWorkspace,
    currentProject,
    requests,
    addItemToRequestModalShown,
    handleUnpublish,
    handleAddItemToRequest,
    handleAddItemToRequestModalClose,
    handleAddItemToRequestModalOpen,
    handleRequestTableChange,
    loading: requestModalLoading,
    totalCount: requestModalTotalCount,
    page: requestModalPage,
    pageSize: requestModalPageSize,
  } = useContentHooks();
  const t = useT();
  const [searchParams, setSearchParams] = useSearchParams();

  const pageParam = useMemo(() => searchParams.get("page"), [searchParams]);
  const pageSizeParam = useMemo(() => searchParams.get("pageSize"), [searchParams]);
  const sortType = useMemo(() => searchParams.get("sortType"), [searchParams]);
  const direction = useMemo(() => searchParams.get("direction"), [searchParams]);
  const searchTermParam = useMemo(() => searchParams.get("searchTerm"), [searchParams]);
  const navigate = useNavigate();
  const { modelId } = useParams();
  const [searchTerm, setSearchTerm] = useState<string>(searchTermParam ?? "");
  const [page, setPage] = useState<number>(pageParam ? +pageParam : 1);
  const [pageSize, setPageSize] = useState<number>(pageSizeParam ? +pageSizeParam : 10);
  const [sort, setSort] = useState<{ type?: ItemSortType; direction?: SortDirection } | undefined>({
    type: sortType ? (sortType as ItemSortType) : "MODIFICATION_DATE",
    direction: direction ? (direction as SortDirection) : "DESC",
  });

  useEffect(() => {
    setPage(pageParam ? +pageParam : 1);
    setPageSize(pageSizeParam ? +pageSizeParam : 10);
    setSort({
      type: sortType ? (sortType as ItemSortType) : "MODIFICATION_DATE",
      direction: direction ? (direction as SortDirection) : "DESC",
    });
    setSearchTerm(searchTermParam ?? "");
  }, [pageParam, pageSizeParam, sortType, direction, searchTermParam]);

  const { data, refetch, loading } = useSearchItemQuery({
    fetchPolicy: "no-cache",
    variables: {
      query: {
        project: currentProject?.id as string,
        schema: currentModel?.schema.id,
        q: searchTerm,
      },
      pagination: { first: pageSize, offset: (page - 1) * pageSize },
      sort: sort
        ? { sortBy: sort.type as GQLItemSortType, direction: sort.direction as GQLSortDirection }
        : undefined,
    },
    skip: !currentModel?.schema.id,
  });

  const handleItemsReload = useCallback(() => {
    refetch();
  }, [refetch]);

  const [collapsedModelMenu, collapseModelMenu] = useState(false);
  const [collapsedCommentsPanel, collapseCommentsPanel] = useState(true);
  const [selectedItemId, setSelectedItemId] = useState<string>();
  const [selection, setSelection] = useState<{ selectedRowKeys: string[] }>({
    selectedRowKeys: [],
  });

  const contentTableFields: ContentTableField[] | undefined = useMemo(() => {
    return data?.searchItem.nodes
      ?.map(item =>
        item
          ? {
              id: item.id,
              schemaId: item.schemaId,
              status: item.status as ItemStatus,
              author: item.user?.name ?? item.integration?.name,
              fields: item?.fields?.reduce(
                (obj, field) =>
                  Object.assign(obj, {
                    [field.schemaFieldId]:
                      field.type === "Asset"
                        ? Array.isArray(field.value)
                          ? field.value
                              .map(value =>
                                fileName(
                                  (item?.assets as GQLAsset[])?.find(asset => asset?.id === value)
                                    ?.url,
                                ),
                              )
                              .join(", ")
                          : fileName(
                              (item?.assets as GQLAsset[])?.find(asset => asset?.id === field.value)
                                ?.url,
                            )
                        : Array.isArray(field.value)
                        ? field.value.join(", ")
                        : field.value
                        ? "" + field.value
                        : field.value,
                  }),
                {},
              ),
              comments: item.thread.comments.map(comment => convertComment(comment as GQLComment)),
              createdAt: item.createdAt,
              updatedAt: item.updatedAt,
            }
          : undefined,
      )
      .filter((contentTableField): contentTableField is ContentTableField => !!contentTableField);
  }, [data?.searchItem.nodes]);

  const contentTableColumns: ProColumns<ContentTableField>[] | undefined = useMemo(() => {
    if (!currentModel) return;
    return [
      {
        title: t("Created By"),
        dataIndex: "author",
        key: "author",
        width: 128,
        minWidth: 128,
        ellipsis: true,
      },
      ...currentModel.schema.fields.map(field => ({
        title: field.title,
        dataIndex: ["fields", field.id],
        key: field.id,
        width: 128,
        minWidth: 128,
        ellipsis: true,
      })),
    ];
  }, [currentModel, t]);

  useEffect(() => {
    if (!modelId && currentModel?.id) {
      navigate(
        `/workspace/${currentWorkspace?.id}/project/${currentProject?.id}/content/${currentModel?.id}`,
      );
    }
  }, [modelId, currentWorkspace?.id, currentProject?.id, currentModel?.id, navigate]);

  const handleModelSelect = useCallback(
    (modelId: string) => {
      navigate(
        `/workspace/${currentWorkspace?.id}/project/${currentProject?.id}/content/${modelId}`,
      );
    },
    [currentWorkspace?.id, currentProject?.id, navigate],
  );

  const handleNavigateToItemForm = useCallback(() => {
    navigate(
      `/workspace/${currentWorkspace?.id}/project/${currentProject?.id}/content/${currentModel?.id}/details`,
    );
  }, [currentWorkspace?.id, currentProject?.id, currentModel?.id, navigate]);

  const handleNavigateToItemEditForm = useCallback(
    (itemId: string) => {
      navigate(
        `/workspace/${currentWorkspace?.id}/project/${currentProject?.id}/content/${currentModel?.id}/details/${itemId}`,
        { state: location.search },
      );
    },
    [currentWorkspace?.id, currentProject?.id, currentModel?.id, navigate],
  );

  const [deleteItemMutation] = useDeleteItemMutation();
  const handleItemDelete = useCallback(
    (itemIds: string[]) =>
      (async () => {
        const results = await Promise.all(
          itemIds.map(async itemId => {
            const result = await deleteItemMutation({
              variables: { itemId },
              refetchQueries: ["SearchItem"],
            });
            if (result.errors) {
              Notification.error({ message: t("Failed to delete one or more items.") });
            }
          }),
        );
        if (results) {
          Notification.success({ message: t("One or more items were successfully deleted!") });
          setSelection({ selectedRowKeys: [] });
        }
      })(),
    [t, deleteItemMutation],
  );

  const handleItemSelect = useCallback(
    (id: string) => {
      setSelectedItemId(id);
      collapseCommentsPanel(false);
    },
    [setSelectedItemId],
  );

  const selectedItem = useMemo(
    () => convertItem(data?.searchItem.nodes.find(item => item?.id === selectedItemId) as GQLItem),
    [data?.searchItem.nodes, selectedItemId],
  );

  const handleContentTableChange = useCallback(
    (
      page: number,
      pageSize: number,
      sorter?: { type?: ItemSortType; direction?: SortDirection },
    ) => {
      searchParams.set("page", page.toString());
      searchParams.set("pageSize", pageSize.toString());
      searchParams.set("sortType", sorter?.type ? sorter.type : "");
      searchParams.set("direction", sorter?.direction ? sorter.direction : "");
      setSearchParams(searchParams);
    },
    [setSearchParams, searchParams],
  );

  const handleSearchTerm = useCallback(
    (term?: string) => {
      searchParams.set("searchTerm", term ?? "");
      setSearchParams(searchParams);
    },
    [setSearchParams, searchParams],
  );

  const handleBulkAddItemToRequest = useCallback(
    async (request: Request, itemIds: string[]) => {
      await handleAddItemToRequest(request, itemIds);
      refetch();
      setSelection({ selectedRowKeys: [] });
    },
    [handleAddItemToRequest, refetch],
  );

  return {
    currentModel,
    loading,
    contentTableFields,
    contentTableColumns,
    collapsedModelMenu,
    collapsedCommentsPanel,
    selectedItem,
    selection,
    totalCount: data?.searchItem.totalCount ?? 0,
    sort,
    searchTerm,
    page,
    pageSize,
    requests,
    addItemToRequestModalShown,
    handleRequestTableChange,
    requestModalLoading,
    requestModalTotalCount,
    requestModalPage,
    requestModalPageSize,
    handleUnpublish,
    handleBulkAddItemToRequest,
    handleAddItemToRequestModalClose,
    handleAddItemToRequestModalOpen,
    handleSearchTerm,
    setSelection,
    handleItemSelect,
    collapseCommentsPanel,
    collapseModelMenu,
    handleModelSelect,
    handleNavigateToItemForm,
    handleNavigateToItemEditForm,
    handleItemsReload,
    handleItemDelete,
    handleContentTableChange,
  };
};
