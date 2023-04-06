import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import { Project, ReearthApi } from "@web/extensions/sidebar/types";
import {
  mergeProperty,
  postMsg,
  mergeOverrides,
  prepareComponentsForOverride,
} from "@web/extensions/sidebar/utils";
import { getActiveFieldIDs, processDatasetToAdd } from "@web/extensions/sidebar/utils/dataset";
import { useCallback, useEffect, useRef, useState } from "react";

import { BuildingSearch, Data, DataCatalogItem, Template } from "../../../types";
import {
  StoryItem,
  Story as FieldStory,
} from "../../content/common/DatasetCard/Field/Fields/types";

export const defaultProject: Project = {
  sceneOverrides: {
    default: {
      camera: {
        lat: 35.65075152248653,
        lng: 139.7617718208305,
        altitude: 2219.7187259974316,
        heading: 6.132702058010316,
        pitch: -0.5672459184621266,
        roll: 0.00019776785897196447,
        fov: 1.0471975511965976,
        height: 2219.7187259974316,
      },
      sceneMode: "3d",
      depthTestAgainstTerrain: false,
    },
    terrain: {
      terrain: true,
      terrainType: "cesiumion",
      terrainCesiumIonAccessToken:
        "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJqdGkiOiI5N2UyMjcwOS00MDY1LTQxYjEtYjZjMy00YTU0ZTg5MmViYWQiLCJpZCI6ODAzMDYsImlhdCI6MTY0Mjc0ODI2MX0.dkwAL1CcljUV7NA7fDbhXXnmyZQU_c-G5zRx8PtEcxE",
      terrainCesiumIonAsset: "770371",
    },
    tiles: [
      {
        id: "tokyo_1",
        tile_url: "https://cyberjapandata.gsi.go.jp/xyz/seamlessphoto/{z}/{x}/{y}.jpg",
        tile_type: "url",
      },
      {
        id: "tokyo_2",
        tile_url:
          "https://gic-plateau.s3.ap-northeast-1.amazonaws.com/2020/ortho/tiles/{z}/{x}/{y}.png",
        tile_type: "url",
      },
    ],
    atmosphere: { shadows: true },
    light: {
      lightType: "directionalLight",
      lightColor: "#ffffffff",
      lightIntensity: 2,
      lightDirectionX: 0.7650124487710819,
      lightDirectionY: -0.6418383470612292,
      lightDirectionZ: -0.05291020191779678,
    },
  },
  datasets: [],
  userStory: undefined,
};

export default ({
  fieldTemplates,
  backendURL,
  backendProjectName,
  buildingSearch,
}: {
  fieldTemplates?: Template[];
  backendURL?: string;
  backendProjectName?: string;
  buildingSearch?: BuildingSearch;
}) => {
  const [projectID, setProjectID] = useState<string>();
  const [project, updateProject] = useState<Project>(defaultProject);
  const [cleanseOverride, setCleanseOverride] = useState<any>();

  const processOverrides = useCallback(
    (dataset: DataCatalogItem, activeIDs?: string[]) => {
      if (!activeIDs) return undefined;
      let overrides = undefined;

      const { activeComponents, inactiveComponents } = prepareComponentsForOverride(
        activeIDs,
        dataset,
        fieldTemplates,
        buildingSearch,
      );

      const cleanseOverrides = mergeOverrides("cleanse", inactiveComponents, cleanseOverride);
      overrides = mergeOverrides("update", activeComponents, cleanseOverrides);

      setCleanseOverride(undefined);

      return overrides;
    },
    [cleanseOverride, fieldTemplates, buildingSearch],
  );

  const handleProjectSceneUpdate = useCallback(
    (updatedProperties: Partial<ReearthApi>) => {
      updateProject(({ sceneOverrides, datasets }) => {
        const updatedProject: Project = {
          sceneOverrides: [sceneOverrides, updatedProperties].reduce((p, v) => mergeProperty(p, v)),
          datasets,
        };
        postMsg({ action: "updateProject", payload: updatedProject });
        postMsg({ action: "msgToPopup", payload: { project: updatedProject } });
        return updatedProject;
      });
    },
    [updateProject],
  );

  const handleProjectDatasetAdd = useCallback(
    (dataset: DataCatalogItem | UserDataItem) => {
      const datasetToAdd = processDatasetToAdd(dataset, fieldTemplates);

      updateProject(project => {
        const datasets = [...project.datasets];
        datasets.unshift(datasetToAdd);
        const updatedProject: Project = {
          ...project,
          datasets,
        };

        postMsg({ action: "updateProject", payload: updatedProject });
        postMsg({ action: "msgToPopup", payload: { project: updatedProject } });

        return updatedProject;
      });

      const activeIDs = getActiveFieldIDs(
        datasetToAdd.components,
        datasetToAdd.selectedGroup,
        datasetToAdd.config?.data,
        fieldTemplates,
      );

      const overrides = processOverrides(datasetToAdd, activeIDs);

      postMsg({
        action: "addDatasetToScene",
        payload: {
          dataset: datasetToAdd,
          overrides,
        },
      });
    },
    [fieldTemplates, processOverrides],
  );

  const handleProjectDatasetRemove = useCallback((dataID: string) => {
    updateProject(({ sceneOverrides, datasets }) => {
      const updatedProject = {
        sceneOverrides,
        datasets: datasets.filter(d => d.dataID !== dataID),
      };
      postMsg({ action: "updateProject", payload: updatedProject });
      postMsg({ action: "msgToPopup", payload: { project: updatedProject } });
      return updatedProject;
    });
    postMsg({ action: "removeDatasetFromScene", payload: dataID });
  }, []);

  const handleProjectDatasetRemoveAll = useCallback(() => {
    updateProject(({ sceneOverrides }) => {
      const updatedProject = {
        sceneOverrides,
        datasets: [],
      };
      postMsg({ action: "updateProject", payload: updatedProject });
      postMsg({ action: "msgToPopup", payload: { project: updatedProject } });
      return updatedProject;
    });
    postMsg({ action: "removeAllDatasetsFromScene" });
  }, []);

  const handleProjectDatasetsUpdate = useCallback((datasets: DataCatalogItem[]) => {
    updateProject(({ sceneOverrides }) => {
      const updatedProject = {
        sceneOverrides,
        datasets,
      };
      postMsg({ action: "updateProject", payload: updatedProject });
      postMsg({ action: "msgToPopup", payload: { project: updatedProject } });
      return updatedProject;
    });
  }, []);

  const handleOverride = useCallback(
    (dataset: DataCatalogItem, activeIDs?: string[]) => {
      if (dataset) {
        const overrides = processOverrides(dataset, activeIDs);

        postMsg({
          action: "updateDatasetInScene",
          payload: { dataID: dataset.dataID, overrides },
        });
      }
    },
    [processOverrides],
  );

  const handleStorySaveData = useCallback(
    (story: StoryItem & { dataID?: string }) => {
      if (story.id && story.dataID) {
        // save database story
        updateProject(project => {
          const tarStory = (
            project.datasets
              .find(d => d.dataID === story.dataID)
              ?.components?.find(c => c.type === "story") as FieldStory
          )?.stories?.find((st: StoryItem) => st.id === story.id);
          if (tarStory) {
            tarStory.scenes = story.scenes;
          }
          return project;
        });
      }

      // save user story
      updateProject(project => {
        const updatedProject: Project = {
          ...project,
          userStory: {
            scenes: story.scenes,
          },
        };
        postMsg({ action: "updateProject", payload: updatedProject });
        return updatedProject;
      });
    },
    [updateProject],
  );

  const handleInitUserStory = useCallback((story: StoryItem) => {
    postMsg({ action: "storyPlay", payload: story });
  }, []);

  const fetchedSharedProject = useRef(false);

  useEffect(() => {
    if (!backendURL || !backendProjectName || fetchedSharedProject.current) return;
    if (projectID) {
      (async () => {
        const res = await fetch(`${backendURL}/share/${backendProjectName}/${projectID}`);
        if (res.status !== 200) return;
        const data = await res.json();
        if (data) {
          (data.datasets as Data[]).forEach(d => {
            handleProjectDatasetAdd(d);
          });
          if (data.userStory && data.userStory.length > 0) {
            handleInitUserStory(data.userStory);
          }
          handleProjectSceneUpdate(data.sceneOverrides);
        }
        fetchedSharedProject.current = true;
      })();
    }
  }, [
    projectID,
    backendURL,
    backendProjectName,
    handleProjectDatasetAdd,
    handleInitUserStory,
    handleProjectSceneUpdate,
  ]);

  return {
    project,
    updateProject,
    setProjectID,
    setCleanseOverride,
    handleOverride,
    handleProjectSceneUpdate,
    handleProjectDatasetAdd,
    handleProjectDatasetRemove,
    handleProjectDatasetRemoveAll,
    handleProjectDatasetsUpdate,
    handleStorySaveData,
  };
};
