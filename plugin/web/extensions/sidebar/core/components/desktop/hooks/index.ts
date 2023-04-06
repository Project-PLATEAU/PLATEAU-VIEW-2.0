import { postMsg, generateID, updateExtended } from "@web/extensions/sidebar/utils";
import { useCallback, useEffect, useState } from "react";

import { BuildingSearch, Template } from "../../../types";
import { Pages } from "../../Header";

import useDatasetHooks from "./datasetHooks";
import useProjectHooks from "./projectHooks";
import useTemplateHooks from "./templateHooks";

export default () => {
  const [inEditor, setInEditor] = useState(true);

  const [reearthURL, setReearthURL] = useState<string>();
  const [backendURL, setBackendURL] = useState<string>();
  const [backendProjectName, setBackendProjectName] = useState<string>();
  const [backendAccessToken, setBackendAccessToken] = useState<string>();
  const [buildingSearch, setBuildingSearch] = useState<BuildingSearch>([]);

  const [loading, setLoading] = useState<boolean>(false);

  const [searchTerm, setSearchTerm] = useState("");

  const handleSearch = useCallback(({ target: { value } }: React.ChangeEvent<HTMLInputElement>) => {
    setSearchTerm(value);
    postMsg({ action: "saveSearchTerm", payload: { searchTerm: value } });
  }, []);

  const {
    fieldTemplates,
    infoboxTemplates,
    updatedTemplateIDs,
    setUpdatedTemplateIDs,
    handleInfoboxFieldsSave,
    setFieldTemplates,
    setInfoboxTemplates,
    handleTemplateAdd,
    handleTemplateSave,
    handleTemplateRemove,
  } = useTemplateHooks({
    backendURL,
    backendProjectName,
    backendAccessToken,
    setLoading,
  });

  const {
    project,
    updateProject,
    setProjectID,
    setCleanseOverride,
    handleProjectSceneUpdate,
    handleProjectDatasetAdd,
    handleProjectDatasetRemove,
    handleProjectDatasetRemoveAll,
    handleProjectDatasetsUpdate,
    handleStorySaveData,
    handleInfoboxFieldsFetch,
    handleOverride,
  } = useProjectHooks({
    fieldTemplates,
    infoboxTemplates,
    updatedTemplateIDs,
    backendURL,
    backendProjectName,
    buildingSearch,
    setUpdatedTemplateIDs,
  });

  const { handleDatasetUpdate, handleDatasetSave } = useDatasetHooks({
    templates: fieldTemplates,
    project,
    backendURL,
    backendProjectName,
    backendAccessToken,
    inEditor,
    setCleanseOverride,
    setLoading,
    updateProject,
  });

  // ****************************************
  // Init

  useEffect(() => {
    postMsg({ action: "init" }); // Needed to trigger sending initialization data to sidebar
  }, []);

  useEffect(() => {
    (async () => {
      if (!backendURL || !backendProjectName) return;
      const res = await fetch(`${backendURL}/sidebar/${backendProjectName}/templates`);
      if (res.status !== 200) return;
      const resData = await res.json();

      if (resData) {
        setFieldTemplates(resData.filter((t: Template) => t.type === "field"));
        setInfoboxTemplates(resData.filter((t: Template) => t.type === "infobox"));
      }
    })();
  }, [backendURL, backendProjectName]); // eslint-disable-line react-hooks/exhaustive-deps

  // ****************************************

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "msgFromModal") {
        if (e.data.payload.dataset) {
          handleProjectDatasetAdd(e.data.payload.dataset);
        }
      } else if (e.data.action === "init" && e.data.payload) {
        setProjectID(e.data.payload.projectID);
        setInEditor(e.data.payload.inEditor);
        setReearthURL(`${e.data.payload.reearthURL}`);
        setBackendURL(e.data.payload.backendURL);
        setBackendProjectName(e.data.payload.backendProjectName);
        setBackendAccessToken(e.data.payload.backendAccessToken);

        if (e.data.payload.searchTerm) setSearchTerm(e.data.payload.searchTerm);
        if (e.data.payload.draftProject) {
          updateProject(e.data.payload.draftProject);
        }
      } else if (e.data.action === "triggerCatalogOpen") {
        handleModalOpen();
      } else if (e.data.action === "triggerHelpOpen") {
        handlePageChange("help");
      } else if (e.data.action === "storyShare") {
        setCurrentPage("share");
      } else if (e.data.action === "storySaveData") {
        handleStorySaveData(e.data.payload);
      } else if (e.data.action === "infoboxFieldsFetch") {
        handleInfoboxFieldsFetch(e.data.payload);
      } else if (e.data.action === "infoboxFieldsSave") {
        handleInfoboxFieldsSave(e.data.payload);
      } else if (e.data.action === "buildingSearchOverride") {
        handleBuildingSearchOverride(e.data.payload);
      } else if (e.data.action === "buildingSearchClose") {
        handleBuildingSearchClose(e.data.payload);
      }
    };
    addEventListener("message", eventListenerCallback);
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, [handleInfoboxFieldsFetch]); // eslint-disable-line react-hooks/exhaustive-deps

  const [currentPage, setCurrentPage] = useState<Pages>("data");

  const handlePageChange = useCallback((p: Pages) => {
    setCurrentPage(p);
  }, []);

  // ****************************************
  // Building Search
  const handleBuildingSearch = useCallback(
    (dataID: string) => {
      const plateauItem = project.datasets.find(pd => pd.dataID === dataID);
      const searchIndex = plateauItem?.["search_index"];
      postMsg({
        action: "buildingSearchOpen",
        payload: {
          title: plateauItem?.["name"] ?? "",
          dataID,
          searchIndex,
        },
      });
    },
    [project.datasets],
  );

  const handleBuildingSearchOverride = useCallback(
    ({ dataID, overrides }: { dataID: string; overrides: any }) => {
      setBuildingSearch(bs => {
        const id = generateID();
        const fieldItem = {
          dataID,
          active: true,
          field: {
            id,
            type: "search",
            updatedAt: new Date(),
            override: overrides,
          },
          cleanseField: {
            id,
            type: "search",
            updatedAt: new Date(),
          },
        };
        const target = bs.find(b => b.dataID === dataID);
        if (target) {
          target.active = true;
          target.field = fieldItem.field;
          target.cleanseField = fieldItem.cleanseField;
        } else {
          bs.push(fieldItem);
        }
        return [...bs];
      });
    },
    [],
  );

  const handleBuildingSearchClose = useCallback(({ dataID }: { dataID: string }) => {
    setBuildingSearch(bs => {
      const target = bs.find(b => b.dataID === dataID);
      if (target) {
        target.active = false;
      }
      return [...bs];
    });
  }, []);

  const handleModalOpen = useCallback(() => {
    postMsg({
      action: "catalogModalOpen",
      payload: {
        templates: fieldTemplates,
      },
    });
  }, [fieldTemplates]);

  return {
    project,
    inEditor,
    reearthURL,
    backendURL,
    backendProjectName,
    templates: fieldTemplates,
    currentPage,
    loading,
    buildingSearch,
    searchTerm,
    handleSearch,
    handlePageChange,
    handleTemplateAdd,
    handleTemplateSave,
    handleTemplateRemove,
    handleDatasetSave,
    handleDatasetUpdate,
    handleProjectDatasetAdd,
    handleProjectDatasetRemove,
    handleProjectDatasetRemoveAll,
    handleProjectDatasetsUpdate,
    handleProjectSceneUpdate,
    handleModalOpen,
    handleBuildingSearch,
    handleOverride,
  };
};

addEventListener("message", e => {
  if (e.source !== parent) return;
  if (e.data.type) {
    if (e.data.type === "extended") {
      updateExtended(e.data.payload);
    }
  }
});
