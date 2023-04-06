import { useEffect, useState, useCallback, Key, useMemo } from "react";
import { matchPath, useLocation, useNavigate, useParams, useSearchParams } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { UploadFile } from "@reearth-cms/components/atoms/Upload";
import { Asset, AssetItem } from "@reearth-cms/components/molecules/Asset/asset.type";
import { convertAsset } from "@reearth-cms/components/organisms/Asset/convertAsset";
import {
  useGetAssetsQuery,
  useCreateAssetMutation,
  useDeleteAssetMutation,
  Asset as GQLAsset,
  SortDirection as GQLSortDirection,
  AssetSortType as GQLSortType,
  useGetAssetsItemsQuery,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";

export type AssetSortType = "DATE" | "NAME" | "SIZE";
export type SortDirection = "ASC" | "DESC";
type UploadType = "local" | "url";

export default () => {
  const t = useT();
  const { pathname } = useLocation();
  const isAssetList = useMemo(
    () => matchPath(":workspaceId/project/:projectId/asset", pathname),
    [pathname],
  );
  const [searchParams, setSearchParams] = useSearchParams();
  const [uploadModalVisibility, setUploadModalVisibility] = useState(false);
  const pageParam = useMemo(
    () => (isAssetList ? searchParams.get("page") : 1),
    [searchParams, isAssetList],
  );
  const pageSizeParam = useMemo(
    () => (isAssetList ? searchParams.get("pageSize") : 10),
    [searchParams, isAssetList],
  );
  const sortType = useMemo(
    () => (isAssetList ? searchParams.get("sortType") : "DATE"),
    [searchParams, isAssetList],
  );
  const direction = useMemo(
    () => (isAssetList ? searchParams.get("direction") : "DESC"),
    [searchParams, isAssetList],
  );
  const searchTermParam = useMemo(
    () => (isAssetList ? searchParams.get("searchTerm") : ""),
    [searchParams, isAssetList],
  );

  const { workspaceId, projectId } = useParams();
  const navigate = useNavigate();
  const [assetList, setAssetList] = useState<Asset[]>([]);
  const [selection, setSelection] = useState<{ selectedRowKeys: Key[] }>({
    selectedRowKeys: [],
  });
  const [selectedAssetId, setSelectedAssetId] = useState<string>();
  const [fileList, setFileList] = useState<UploadFile<File>[]>([]);
  const [uploadUrl, setUploadUrl] = useState<{ url: string; autoUnzip: boolean }>({
    url: "",
    autoUnzip: true,
  });
  const [uploadType, setUploadType] = useState<UploadType>("local");
  const [uploading, setUploading] = useState(false);
  const [collapsed, setCollapsed] = useState(true);
  const [page, setPage] = useState<number>(pageParam ? +pageParam : 1);
  const [pageSize, setPageSize] = useState<number>(pageSizeParam ? +pageSizeParam : 10);
  const [searchTerm, setSearchTerm] = useState<string>(searchTermParam ?? "");

  const [createAssetMutation] = useCreateAssetMutation();

  const [sort, setSort] = useState<{ type?: AssetSortType; direction?: SortDirection } | undefined>(
    {
      type: sortType ? (sortType as AssetSortType) : "DATE",
      direction: direction ? (direction as SortDirection) : "DESC",
    },
  );

  useEffect(() => {
    setPage(pageParam ? +pageParam : 1);
    setPageSize(pageSizeParam ? +pageSizeParam : 10);
    setSort({
      type: sortType ? (sortType as AssetSortType) : "DATE",
      direction: direction ? (direction as SortDirection) : "DESC",
    });
    setSearchTerm(searchTermParam ?? "");
  }, [pageParam, pageSizeParam, sortType, direction, searchTermParam]);

  const { data, refetch, loading, networkStatus } = useGetAssetsQuery({
    fetchPolicy: "no-cache",
    variables: {
      projectId: projectId ?? "",
      pagination: { first: pageSize, offset: (page - 1) * pageSize },
      sort: sort
        ? { sortBy: sort.type as GQLSortType, direction: sort.direction as GQLSortDirection }
        : undefined,
      keyword: searchTerm,
    },
    notifyOnNetworkStatusChange: true,
    skip: !projectId,
  });

  const { data: assetsItems, refetch: refetchAssetsItems } = useGetAssetsItemsQuery({
    fetchPolicy: "no-cache",
    variables: {
      projectId: projectId ?? "",
      pagination: { first: pageSize, offset: (page - 1) * pageSize },
      sort: sort
        ? { sortBy: sort.type as GQLSortType, direction: sort.direction as GQLSortDirection }
        : undefined,
      keyword: searchTerm,
    },
    notifyOnNetworkStatusChange: true,
    skip: !projectId,
  });

  const isRefetching = networkStatus === 3;

  const handleUploadModalCancel = useCallback(() => {
    setUploadModalVisibility(false);
    setUploading(false);
    setFileList([]);
    setUploadUrl({ url: "", autoUnzip: true });
    setUploadType("local");
  }, [setUploadModalVisibility, setUploading, setFileList, setUploadUrl, setUploadType]);

  const handleAssetsCreate = useCallback(
    (files: UploadFile<File>[]) =>
      (async () => {
        if (!projectId) return [];
        setUploading(true);
        const results = (
          await Promise.all(
            files.map(async file => {
              const result = await createAssetMutation({
                variables: {
                  projectId,
                  file,
                  skipDecompression: !!file.skipDecompression,
                },
              });
              if (result.errors || !result.data?.createAsset) {
                Notification.error({ message: t("Failed to add one or more assets.") });
                handleUploadModalCancel();
                return undefined;
              }
              return convertAsset(result.data.createAsset.asset as GQLAsset);
            }),
          )
        ).filter(Boolean);
        if (results?.length > 0) {
          Notification.success({ message: t("Successfully added one or more assets!") });
          await refetch();
          refetchAssetsItems();
        }
        handleUploadModalCancel();
        return results;
      })(),
    [projectId, handleUploadModalCancel, createAssetMutation, t, refetch, refetchAssetsItems],
  );

  const handleAssetCreateFromUrl = useCallback(
    async (url: string, autoUnzip: boolean) => {
      if (!projectId) return undefined;
      setUploading(true);
      try {
        const result = await createAssetMutation({
          variables: {
            projectId,
            file: null,
            url,
            skipDecompression: !autoUnzip,
          },
        });
        if (result.data?.createAsset) {
          Notification.success({ message: t("Successfully added asset!") });
          await refetch();
          refetchAssetsItems();
          return convertAsset(result.data.createAsset.asset as GQLAsset);
        }
        return undefined;
      } catch {
        Notification.error({ message: t("Failed to add asset.") });
      } finally {
        handleUploadModalCancel();
      }
    },
    [projectId, createAssetMutation, t, refetch, refetchAssetsItems, handleUploadModalCancel],
  );

  const [deleteAssetMutation] = useDeleteAssetMutation();
  const handleAssetDelete = useCallback(
    (assetIds: string[]) =>
      (async () => {
        if (!projectId) return;
        const results = await Promise.all(
          assetIds.map(async assetId => {
            const result = await deleteAssetMutation({
              variables: { assetId },
            });
            if (result.errors) {
              Notification.error({ message: t("Failed to delete one or more assets.") });
            }
          }),
        );
        if (results) {
          await refetch();
          refetchAssetsItems();
          Notification.success({ message: t("One or more assets were successfully deleted!") });
          setSelection({ selectedRowKeys: [] });
        }
      })(),
    [t, deleteAssetMutation, refetch, refetchAssetsItems, projectId],
  );

  const handleSearchTerm = useCallback(
    (term?: string) => {
      if (isAssetList) {
        searchParams.set("searchTerm", term ?? "");
        setSearchParams(searchParams);
      } else {
        setSearchTerm(term ?? "");
      }
    },
    [setSearchParams, searchParams, isAssetList],
  );

  const handleAssetsReload = useCallback(() => {
    refetch();
    refetchAssetsItems();
  }, [refetch, refetchAssetsItems]);

  const handleNavigateToAsset = (asset: Asset) => {
    navigate(`/workspace/${workspaceId}/project/${projectId}/asset/${asset.id}`);
  };

  useEffect(() => {
    const assets =
      (data?.assets.nodes
        .map(asset => asset as GQLAsset)
        .map(convertAsset)
        .filter(asset => !!asset) as Asset[]) ?? [];
    setAssetList(
      assets.map(asset => ({
        ...asset,
        items: assetsItems?.assets.nodes.find(assetItem => assetItem?.id === asset.id)?.items ?? [],
      })),
    );
  }, [data?.assets.nodes, assetsItems?.assets.nodes, setAssetList]);

  const handleAssetSelect = useCallback(
    (id: string) => {
      setSelectedAssetId(id);
      setCollapsed(false);
    },
    [setCollapsed, setSelectedAssetId],
  );

  const handleAssetItemSelect = useCallback(
    (assetItem: AssetItem) => {
      navigate(
        `/workspace/${workspaceId}/project/${projectId}/content/${assetItem.modelId}/details/${assetItem.itemId}`,
      );
    },
    [navigate, projectId, workspaceId],
  );

  const handleToggleCommentMenu = useCallback(
    (value: boolean) => {
      setCollapsed(value);
    },
    [setCollapsed],
  );

  const selectedAsset = useMemo(
    () => assetList.find(asset => asset.id === selectedAssetId),
    [assetList, selectedAssetId],
  );

  const handleAssetTableChange = useCallback(
    (
      page: number,
      pageSize: number,
      sorter?: { type?: AssetSortType; direction?: SortDirection },
    ) => {
      if (isAssetList) {
        searchParams.set("page", page.toString());
        searchParams.set("pageSize", pageSize.toString());
        searchParams.set("sortType", sorter?.type ? sorter.type : "");
        searchParams.set("direction", sorter?.direction ? sorter.direction : "");
      } else {
        setSearchParams(searchParams);
        setPage(page);
        setPageSize(pageSize);
        setSort({
          type: sorter?.type ? (sorter.type as AssetSortType) : "DATE",
          direction: sorter?.direction ? (sorter.direction as SortDirection) : "DESC",
        });
      }
    },
    [setSearchParams, searchParams, isAssetList],
  );

  return {
    assetList,
    selection,
    fileList,
    uploading,
    isLoading: loading ?? isRefetching,
    uploadModalVisibility,
    loading,
    uploadUrl,
    uploadType,
    selectedAsset,
    collapsed,
    totalCount: data?.assets.totalCount ?? 0,
    searchTerm,
    page,
    pageSize,
    sort,
    handleToggleCommentMenu,
    handleAssetItemSelect,
    handleAssetSelect,
    handleUploadModalCancel,
    setUploadUrl,
    setUploadType,
    setSelection,
    setFileList,
    setUploading,
    setUploadModalVisibility,
    handleAssetsCreate,
    handleAssetCreateFromUrl,
    handleAssetTableChange,
    handleAssetDelete,
    handleSearchTerm,
    handleAssetsReload,
    handleNavigateToAsset,
  };
};
