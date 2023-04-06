import { Project } from "@web/extensions/sidebar/types";
import { postMsg, convertDatasetToData } from "@web/extensions/sidebar/utils";
import { useCallback, useEffect } from "react";

import { Data, DataCatalogItem, Template } from "../../../types";

export default ({
  templates,
  project,
  backendURL,
  backendProjectName,
  backendAccessToken,
  inEditor,
  setCleanseOverride,
  setLoading,
  updateProject,
}: {
  templates?: Template[];
  project?: Project;
  backendURL?: string;
  backendProjectName?: string;
  backendAccessToken?: string;
  inEditor?: boolean;
  setCleanseOverride?: React.Dispatch<React.SetStateAction<string | undefined>>;
  setLoading?: React.Dispatch<React.SetStateAction<boolean>>;
  updateProject?: React.Dispatch<React.SetStateAction<Project>>;
}) => {
  const handleDataFetch = useCallback(async () => {
    if (!backendURL) return;
    const res = await fetch(`${backendURL}/sidebar/${backendProjectName}/data`);
    if (res.status !== 200) return;
    const resData = await res.json();

    return resData;
  }, [backendURL, backendProjectName]);

  const handleDataRequest = useCallback(
    async (dataset?: DataCatalogItem) => {
      if (!backendURL || !backendAccessToken || !dataset) return;
      const datasetToSave = convertDatasetToData(dataset, templates);

      const data = await handleDataFetch();
      const isNew = data ? !data.find((d: Data) => d.dataID === dataset.dataID) : undefined;

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
      if (res.status !== 200) {
        console.log("A problem occured accessing the server:", res.statusText);
        return;
      }
    },
    [templates, backendAccessToken, backendURL, backendProjectName, handleDataFetch],
  );

  const handleDatasetUpdate = useCallback(
    (updatedDataset: DataCatalogItem, cleanseOverride?: any) => {
      updateProject?.(project => {
        const updatedDatasets = [...project.datasets];
        const datasetIndex = updatedDatasets.findIndex(d2 => d2.dataID === updatedDataset.dataID);
        if (datasetIndex >= 0) {
          if (updatedDatasets[datasetIndex].visible !== updatedDataset.visible) {
            postMsg({
              action: "updateDatasetVisibility",
              payload: { dataID: updatedDataset.dataID, hide: !updatedDataset.visible },
            });
          }
          if (cleanseOverride) {
            setCleanseOverride?.(cleanseOverride);
          }
          updatedDatasets[datasetIndex] = updatedDataset;
        }
        const updatedProject = {
          ...project,
          datasets: updatedDatasets,
        };
        postMsg({ action: "updateProject", payload: updatedProject });
        return updatedProject;
      });
    },
    [updateProject, setCleanseOverride],
  );

  const handleDatasetSave = useCallback(
    (dataID: string) => {
      (async () => {
        if (!inEditor) return;
        setLoading?.(true);
        const selectedDataset = project?.datasets.find(d => d.dataID === dataID);

        await handleDataRequest(selectedDataset);
        setLoading?.(false);
      })();
    },
    [inEditor, project?.datasets, setLoading, handleDataRequest],
  );

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "updateDataset" && e.data.payload) {
        handleDatasetUpdate(e.data.payload);
      }
    };
    addEventListener("message", eventListenerCallback);
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, [handleDatasetUpdate]);

  return {
    handleDatasetUpdate,
    handleDatasetSave,
  };
};
