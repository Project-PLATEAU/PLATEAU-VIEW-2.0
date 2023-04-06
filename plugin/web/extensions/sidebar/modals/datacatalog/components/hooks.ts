import { Data, DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import {
  convertDatasetToData,
  handleDataCatalogProcessing,
  postMsg,
} from "@web/extensions/sidebar/utils";
import { debounce } from "lodash";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";

import { RawDataCatalogItem, getDataCatalog, GroupBy, DataCatalogGroup } from "../api/api";

export type Tab = "dataset" | "your-data";

export default () => {
  const [currentTab, changeTabs] = useState<Tab>("dataset");
  const [addedDatasetDataIDs, setAddedDatasetDataIDs] = useState<string[]>();
  const [catalogData, setCatalog] = useState<RawDataCatalogItem[]>([]);
  const [inEditor, setEditorState] = useState(false);
  const [selectedItem, selectItem] = useState<DataCatalogItem | DataCatalogGroup>();
  const [expandedFolders, setExpandedFolders] = useState<{ id?: string; name?: string }[]>([]);
  const [searchTerm, setSearchTerm] = useState("");

  const [templates, setTemplates] = useState<Template[]>([]);

  const [backendProjectName, setBackendProjectName] = useState<string>();
  const [backendAccessToken, setBackendAccessToken] = useState<string>();
  const [backendURL, setBackendURL] = useState<string>();
  const [catalogURL, setCatalogURL] = useState<string>();
  const [publishToGeospatial, setPublishToGeospatial] = useState(false);

  const [catalogProjectName, setCatalogProjectName] = useState<string>();

  const [data, setData] = useState<Data[]>();

  const processedCatalog = useMemo(() => {
    if (catalogData.length < 1 || data === undefined) return;
    const c = handleDataCatalogProcessing(catalogData, data);
    return inEditor ? c : c.filter(c => !!c.public || c.type_en === "folder");
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

    setData(resData ?? []);
  }, [backendURL, backendProjectName]);

  const handleDataRequest = useCallback(
    async (dataset?: DataCatalogItem) => {
      if (!backendURL || !backendAccessToken || !dataset) return;
      const datasetToSave = convertDatasetToData(dataset, templates);

      const isNew = !data?.find(d => d.dataID === dataset.dataID);

      const fetchURL = !isNew
        ? `${backendURL}/sidebar/${backendProjectName}/data/${dataset.id}` // should be id and not dataID because id here is the CMS item's id
        : `${backendURL}/sidebar/${backendProjectName}/data`;

      const method = !isNew ? "PATCH" : "POST";

      const res = await fetch(fetchURL, {
        headers: {
          authorization: `Bearer ${backendAccessToken}`,
        },
        method,
        body: JSON.stringify(datasetToSave),
      });
      if (res.status === 200) {
        const resData = await res.json();
        setData(prevData => {
          if (!prevData) {
            return [resData];
          }
          const index = prevData?.findIndex(d => d.dataID === resData.dataID);
          if (index) {
            const updatedData = [...prevData];
            updatedData[index] = resData;
            return updatedData;
          }
        });
      }
    },
    [data, templates, backendAccessToken, backendURL, backendProjectName],
  );

  const handleDatasetPublish = useCallback(
    (dataID: string, publish: boolean) => {
      if (!inEditor || !processedCatalog) return;
      const dataset = processedCatalog.find(item => item.dataID === dataID);

      if (!dataset) return;

      dataset.public = publish;

      postMsg({ action: "updateDataset", payload: dataset });
      handleDataRequest(dataset);

      if (publish && publishToGeospatial && dataset.itemId && backendURL && backendAccessToken) {
        fetch(`${backendURL}/publish_to_geospatialjp`, {
          headers: {
            authorization: `Bearer ${backendAccessToken}`,
            "Content-Type": "application/json",
          },
          method: "POST",
          body: JSON.stringify({ id: dataset.itemId }),
        })
          .then(r => {
            if (!r.ok)
              throw `failed to publish the data on gspatial.jp: status code is ${r.statusText}`;
          })
          .catch(console.error);
      }
    },
    [
      processedCatalog,
      inEditor,
      backendAccessToken,
      backendURL,
      publishToGeospatial,
      handleDataRequest,
    ],
  );
  const [filter, setFilter] = useState<GroupBy>("city");

  const debouncedSearchRef = useRef(
    debounce((value: string) => {
      postMsg({ action: "saveSearchTerm", payload: { searchTerm: value } });
      setExpandedFolders([]);
      postMsg({ action: "saveExpandedFolders", payload: { expandedFolders: [] } });
    }, 300),
  );

  const handleSearch = useCallback(
    ({ target: { value } }: React.ChangeEvent<HTMLInputElement>) => {
      setSearchTerm(value);
      debouncedSearchRef.current(value);
    },
    [debouncedSearchRef],
  );

  const handleSelect = useCallback((item?: DataCatalogItem | DataCatalogGroup) => {
    selectItem(item);
  }, []);

  const handleClose = useCallback(() => {
    postMsg({ action: "modalClose" });
  }, []);

  const handleFilter = useCallback((filter: GroupBy) => {
    setFilter(filter);
    postMsg({ action: "saveFilter", payload: { filter } });
    setExpandedFolders([]);
    postMsg({ action: "saveExpandedFolders", payload: { expandedFolders: [] } });
  }, []);

  const handleDatasetAdd = useCallback(
    (dataset: DataCatalogItem | UserDataItem, keepModalOpen?: boolean) => {
      postMsg({
        action: "msgFromModal",
        payload: {
          dataset,
        },
      });
      if (!keepModalOpen) handleClose();
    },
    [handleClose],
  );

  useEffect(() => {
    postMsg({ action: "initDataCatalog" }); // Needed to trigger sending selected dataset ids from Sidebar
  }, []);

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "initDataCatalog") {
        setAddedDatasetDataIDs(e.data.payload.addedDatasets);
        setEditorState(e.data.payload.inEditor);
        setBackendProjectName(e.data.payload.backendProjectName);
        setBackendAccessToken(e.data.payload.backendAccessToken);
        setBackendURL(e.data.payload.backendURL);
        setCatalogURL(e.data.payload.catalogURL);
        setCatalogProjectName(e.data.payload.catalogProjectName);
        setPublishToGeospatial(e.data.payload.enableGeoPub);
        setTemplates(e.data.payload.templates);
        if (e.data.payload.filter) setFilter(e.data.payload.filter);
        if (e.data.payload.searchTerm) setSearchTerm(e.data.payload.searchTerm);
        if (e.data.payload.expandedFolders) setExpandedFolders(e.data.payload.expandedFolders);
        if (e.data.payload.dataset) {
          const item = e.data.payload.dataset;
          handleSelect(item);
          if (item.path) {
            setExpandedFolders(
              item.path
                .map((item: string) => ({ name: item }))
                .filter((folder: { name?: string }) => folder.name !== item.name),
            );
          }
          postMsg({
            action: "saveDataset",
            payload: { dataset: undefined },
          });
        }
      } else if (e.data.action === "updateDataCatalog") {
        if (e.data.payload.updatedDatasetDataIDs) {
          setAddedDatasetDataIDs(e.data.payload.updatedDatasetDataIDs);
        }
      }
    };
    addEventListener("message", eventListenerCallback);
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, [handleFilter, handleSelect]);

  return {
    currentTab,
    catalog: processedCatalog,
    addedDatasetDataIDs,
    inEditor,
    selectedItem,
    expandedFolders,
    searchTerm,
    filter,
    setExpandedFolders,
    handleSearch,
    handleSelect,
    handleFilter,
    handleClose,
    handleTabChange: changeTabs,
    handleDatasetAdd,
    handleDatasetPublish,
  };
};
