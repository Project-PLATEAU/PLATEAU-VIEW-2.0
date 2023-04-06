import Footer from "@web/extensions/sidebar/core/components/Footer";
import { ReearthApi } from "@web/extensions/sidebar/types";
import { swap } from "@web/extensions/sidebar/utils";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import { BuildingSearch, DataCatalogItem, Template } from "../../../types";
import DatasetCard from "../common/DatasetCard";

export type Props = {
  className?: string;
  inEditor?: boolean;
  selectedDatasets?: DataCatalogItem[];
  templates?: Template[];
  buildingSearch?: BuildingSearch;
  savingDataset?: boolean;
  isMobile?: boolean;
  onDatasetSave?: (dataID: string) => void;
  onDatasetUpdate: (dataset: DataCatalogItem, cleanseOverride?: any) => void;
  onDatasetRemove: (dataID: string) => void;
  onDatasetRemoveAll: () => void;
  onProjectDatasetsUpdate: (datasets: DataCatalogItem[]) => void;
  onModalOpen?: () => void;
  onBuildingSearch: (id: string) => void;
  onOverride?: (dataID: string, activeIDs?: string[]) => void;
  onSceneUpdate: (updatedProperties: Partial<ReearthApi>) => void;
};

const Selection: React.FC<Props> = ({
  className,
  inEditor,
  selectedDatasets,
  templates,
  buildingSearch,
  savingDataset,
  isMobile,
  onDatasetSave,
  onDatasetUpdate,
  onDatasetRemove,
  onDatasetRemoveAll,
  onProjectDatasetsUpdate,
  onModalOpen,
  onBuildingSearch,
  onOverride,
  onSceneUpdate,
}) => {
  const moveCard = useCallback(
    (dragIndex: number, hoverIndex: number) => {
      if (selectedDatasets) {
        const updatedDatasets = [...selectedDatasets];
        swap(updatedDatasets, dragIndex, hoverIndex);
        onProjectDatasetsUpdate(updatedDatasets);
      }
    },
    [onProjectDatasetsUpdate, selectedDatasets],
  );

  return (
    <Wrapper className={className}>
      <InnerWrapper>
        {onModalOpen && (
          <StyledButton onClick={onModalOpen}>
            <StyledIcon icon="plusCircle" size={20} />
            <ButtonText>カタログから検索する</ButtonText>
          </StyledButton>
        )}
        <DndProvider backend={HTML5Backend}>
          {selectedDatasets?.map((d, i) => (
            <DatasetCard
              key={d.id}
              index={i}
              id={d.id}
              dataset={d}
              templates={templates}
              buildingSearch={buildingSearch}
              savingDataset={savingDataset}
              inEditor={inEditor}
              isMobile={isMobile}
              moveCard={moveCard}
              onDatasetSave={onDatasetSave}
              onDatasetUpdate={onDatasetUpdate}
              onDatasetRemove={onDatasetRemove}
              onBuildingSearch={onBuildingSearch}
              onOverride={onOverride}
              onSceneUpdate={onSceneUpdate}
            />
          ))}
        </DndProvider>
      </InnerWrapper>
      <Footer datasetQuantity={selectedDatasets?.length} onRemoveAll={onDatasetRemoveAll} />
    </Wrapper>
  );
};

export default Selection;

const Wrapper = styled.div`
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: space-between;
`;

const InnerWrapper = styled.div`
  padding: 16px;
  flex: 1;
  overflow: auto;
`;

const StyledButton = styled.button`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  border: none;
  border-radius: 4px;
  background: #00bebe;
  color: #fff;
  padding: 10px;
  cursor: pointer;
`;

const ButtonText = styled.p`
  margin: 0;
  user-select: none;
`;

const StyledIcon = styled(Icon)`
  margin-right: 8px;
`;
