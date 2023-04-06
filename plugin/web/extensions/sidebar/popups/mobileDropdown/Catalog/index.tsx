import { Data } from "@web/extensions/sidebar/core/types";
import {
  DataCatalogGroup,
  DataCatalogItem,
  GroupBy,
  RawDataCatalogItem,
  getDataCatalog,
} from "@web/extensions/sidebar/modals/datacatalog/api/api";
import DatasetTree from "@web/extensions/sidebar/modals/datacatalog/components/content/DatasetsPage/DatasetTree";
import DatasetDetails from "@web/extensions/sidebar/modals/datacatalog/components/content/DatasetsPage/Details";
import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import { handleDataCatalogProcessing, postMsg } from "@web/extensions/sidebar/utils";
import { styled } from "@web/theme";
import { useCallback, useEffect, useMemo, useState } from "react";

import PopupItem from "../sharedComponents/PopupItem";

type Props = {
  addedDatasetDataIDs?: string[];
  isMobile?: boolean;
  searchTerm: string;
  expandedFolders?: {
    id?: string | undefined;
    name?: string | undefined;
  }[];
  selectedDataset?: DataCatalogItem;
  inEditor?: boolean;
  catalogProjectName?: string;
  catalogURL?: string;
  backendURL?: string;
  backendProjectName?: string;
  setSelectedDataset: React.Dispatch<React.SetStateAction<DataCatalogItem | undefined>>;
  setExpandedFolders?: React.Dispatch<
    React.SetStateAction<
      {
        id?: string | undefined;
        name?: string | undefined;
      }[]
    >
  >;
  onSearch: (e: React.ChangeEvent<HTMLInputElement>) => void;
  setSearchTerm: (searchTerm: string) => void;
  onDatasetAdd: (dataset: DataCatalogItem | UserDataItem, keepModalOpen?: boolean) => void;
};

const Catalog: React.FC<Props> = ({
  addedDatasetDataIDs,
  isMobile,
  searchTerm,
  expandedFolders,
  selectedDataset,
  inEditor,
  catalogProjectName,
  catalogURL,
  backendURL,
  backendProjectName,
  setSelectedDataset,
  setExpandedFolders,
  onSearch,
  setSearchTerm,
  onDatasetAdd,
}) => {
  const [filter, setFilter] = useState<GroupBy>("city");
  const [page, setPage] = useState<"catalog" | "details">("catalog");

  const [catalogData, setCatalog] = useState<RawDataCatalogItem[]>([]);
  const [data, setData] = useState<Data[]>();

  const processedCatalog = useMemo(() => {
    const c = handleDataCatalogProcessing(catalogData, data);
    return inEditor ? c : c.filter(c => !!c.public);
  }, [catalogData, inEditor, data]);

  useEffect(() => {
    const catalogBaseUrl = catalogURL || backendURL;
    if (catalogBaseUrl) {
      getDataCatalog(catalogBaseUrl, catalogProjectName).then(res => {
        setCatalog(res);
      });
    }
  }, [backendURL, catalogProjectName, catalogURL]);

  useEffect(() => {
    if (!backendURL) return;
    handleDataFetch();
  }, [backendURL]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleDataFetch = useCallback(async () => {
    if (!backendURL || !backendProjectName) return;
    const res = await fetch(`${backendURL}/sidebar/${backendProjectName}/data`);
    if (res.status !== 200) return;
    const resData = await res.json();

    setData(resData);
  }, [backendURL, backendProjectName]);

  const handleOpenDetails = useCallback(
    (data?: DataCatalogItem | DataCatalogGroup) => {
      if (data && "dataID" in data) {
        setSelectedDataset(data);
        setPage("details");
      }
    },
    [setSelectedDataset],
  );

  const handleFilter = useCallback(
    (filter: GroupBy) => {
      setFilter(filter);
      postMsg({ action: "saveFilter", payload: { filter } });
      setExpandedFolders?.([]);
      postMsg({ action: "saveExpandedFolders", payload: { expandedFolders: [] } });
    },
    [setExpandedFolders],
  );

  const addDisabled = useCallback(
    (dataID: string) => {
      return !!addedDatasetDataIDs?.find(dataID2 => dataID2 === dataID);
    },
    [addedDatasetDataIDs],
  );

  const handleBack = useCallback(() => {
    setPage("catalog");
    setSelectedDataset(undefined);
  }, [setPage, setSelectedDataset]);

  useEffect(() => {
    postMsg({ action: "extendPopup" });
  }, []);

  useEffect(() => {
    if (selectedDataset && page !== "details") {
      setPage("details");
    }
  }, [selectedDataset, page, setPage, setSelectedDataset]);

  useEffect(() => {
    postMsg({ action: "initDataCatalog" });
  }, []);

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "initDataCatalog") {
        if (e.data.payload.filter) setFilter(e.data.payload.filter);
        if (e.data.payload.searchTerm) setSearchTerm(e.data.payload.searchTerm);
        if (e.data.payload.expandedFolders) setExpandedFolders?.(e.data.payload.expandedFolders);
      }
    };
    addEventListener("message", eventListenerCallback);
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, [handleFilter, onSearch, setExpandedFolders, setSearchTerm]);

  return (
    <Wrapper>
      {page === "catalog" && (
        <>
          <PopupItem>
            <Title>データカタログ</Title>
          </PopupItem>
          <DatasetTree
            addedDatasetDataIDs={addedDatasetDataIDs}
            selectedItem={selectedDataset}
            isMobile={isMobile}
            catalog={processedCatalog}
            filter={filter}
            searchTerm={searchTerm}
            expandedFolders={expandedFolders}
            setExpandedFolders={setExpandedFolders}
            addDisabled={addDisabled}
            onSearch={onSearch}
            onFilter={handleFilter}
            onSelect={handleOpenDetails}
            onDatasetAdd={onDatasetAdd}
          />
        </>
      )}
      {page === "details" && (
        <>
          <PopupItem onBack={handleBack}>
            <Title>データ詳細</Title>
          </PopupItem>
          <DatasetDetails
            dataset={selectedDataset}
            isMobile={isMobile}
            addDisabled={addDisabled}
            onDatasetAdd={onDatasetAdd}
          />
        </>
      )}
    </Wrapper>
  );
};

export default Catalog;

const Wrapper = styled.div`
  border-top: 1px solid #d9d9d9;
`;

const Title = styled.p`
  margin: 0;
`;
