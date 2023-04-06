import Feedback from "@web/extensions/sidebar/core/components/content/Feedback";
import Help from "@web/extensions/sidebar/core/components/content/Help";
import MapSettings from "@web/extensions/sidebar/core/components/content/MapSettings";
import Selection from "@web/extensions/sidebar/core/components/content/Selection";
import Share from "@web/extensions/sidebar/core/components/content/Share";
import Templates from "@web/extensions/sidebar/core/components/content/Templates";
import Header from "@web/extensions/sidebar/core/components/Header";
import { Content } from "@web/sharedComponents";
import { styled, commonStyles } from "@web/theme";
import { memo, useCallback, useEffect, useState } from "react";

import { postMsg } from "../../../utils";

import useHooks from "./hooks";

export type Props = {
  className?: string;
};

const DesktopSidebar: React.FC<Props> = ({ className }) => {
  const {
    project,
    inEditor,
    reearthURL,
    backendURL,
    backendProjectName,
    templates,
    currentPage,
    loading,
    buildingSearch,
    handlePageChange,
    handleTemplateAdd,
    handleTemplateSave,
    handleTemplateRemove,
    handleDatasetSave,
    handleProjectDatasetRemove,
    handleDatasetUpdate,
    handleProjectDatasetRemoveAll,
    handleProjectDatasetsUpdate,
    handleProjectSceneUpdate,
    handleModalOpen,
    handleBuildingSearch,
    handleOverride,
  } = useHooks();

  const [minimized, setMinimize] = useState(false);

  useEffect(() => {
    setTimeout(() => {
      postMsg({ action: "minimize", payload: minimized });
    }, 250);
  }, [minimized]);

  const handleMinimize = useCallback(() => {
    const html = document.querySelector("html");
    const body = document.querySelector("body");
    const root = document.getElementById("root");
    if (!minimized) {
      html?.classList.add("minimized");
      body?.classList.add("minimized");
      root?.classList.add("minimized");
    } else {
      html?.classList.remove("minimized");
      body?.classList.remove("minimized");
      root?.classList.remove("minimized");
    }
    setMinimize(!minimized);
  }, [minimized, setMinimize]);

  return (
    <Wrapper className={className} minimized={minimized}>
      <Header
        current={currentPage}
        inEditor={inEditor}
        minimized={minimized}
        onMinimize={handleMinimize}
        onClick={handlePageChange}
      />
      {!minimized && (
        <ContentWrapper className={className}>
          {
            {
              data: (
                <Selection
                  inEditor={inEditor}
                  selectedDatasets={project.datasets}
                  templates={templates}
                  buildingSearch={buildingSearch}
                  savingDataset={loading}
                  onDatasetSave={handleDatasetSave}
                  onDatasetUpdate={handleDatasetUpdate}
                  onDatasetRemove={handleProjectDatasetRemove}
                  onDatasetRemoveAll={handleProjectDatasetRemoveAll}
                  onProjectDatasetsUpdate={handleProjectDatasetsUpdate}
                  onModalOpen={handleModalOpen}
                  onBuildingSearch={handleBuildingSearch}
                  onOverride={handleOverride}
                  onSceneUpdate={handleProjectSceneUpdate}
                />
              ),
              map: (
                <MapSettings
                  overrides={project.sceneOverrides}
                  onOverridesUpdate={handleProjectSceneUpdate}
                />
              ),
              share: (
                <Share
                  project={project}
                  reearthURL={reearthURL}
                  backendURL={backendURL}
                  backendProjectName={backendProjectName}
                />
              ),
              help: <Help />,
              feedback: <Feedback backendURL={backendURL} />,
              template: (
                <Templates
                  templates={templates}
                  savingTemplate={loading}
                  onTemplateAdd={handleTemplateAdd}
                  onTemplateSave={handleTemplateSave}
                  onTemplateRemove={handleTemplateRemove}
                />
              ),
            }[currentPage]
          }
        </ContentWrapper>
      )}
    </Wrapper>
  );
};

export default memo(DesktopSidebar);

const Wrapper = styled.div<{ minimized?: boolean }>`
  display: flex;
  flex-direction: column;
  ${commonStyles.mainWrapper}
  transition: height 0.5s, width 0.5s, border-radius 0.5s;
  ${({ minimized }) => minimized && commonStyles.minimizedWrapper}
`;

const ContentWrapper = styled(Content)`
  flex: 1;
  background: #e7e7e7;
  box-sizing: border-box;
  overflow: auto;
`;
