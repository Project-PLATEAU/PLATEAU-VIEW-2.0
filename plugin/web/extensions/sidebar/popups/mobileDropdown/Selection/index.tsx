import SelectionComponent from "@web/extensions/sidebar/core/components/content/Selection";
import { DataCatalogItem, BuildingSearch, Template } from "@web/extensions/sidebar/core/types";
import { ReearthApi } from "@web/extensions/sidebar/types";
import { postMsg } from "@web/extensions/sidebar/utils";
import { styled } from "@web/theme";
import { useEffect } from "react";

import PopupItem from "../sharedComponents/PopupItem";

type Props = {
  selectedDatasets: DataCatalogItem[];
  savingDataset?: boolean;
  buildingSearch?: BuildingSearch;
  templates?: Template[];
  onDatasetUpdate: (updatedDataset: DataCatalogItem) => void;
  onDatasetRemove: (id: string) => void;
  onDatasetRemoveAll: () => void;
  onProjectDatasetsUpdate: (datasets: DataCatalogItem[]) => void;
  onBuildingSearch: (id: string) => void;
  onSceneUpdate: (updatedProperties: Partial<ReearthApi>) => void;
};

const Selection: React.FC<Props> = ({
  selectedDatasets,
  savingDataset,
  buildingSearch,
  templates,
  onDatasetUpdate,
  onDatasetRemove,
  onDatasetRemoveAll,
  onProjectDatasetsUpdate,
  onBuildingSearch,
  onSceneUpdate,
}) => {
  useEffect(() => {
    postMsg({ action: "extendPopup" });
  }, []);

  return (
    <Wrapper>
      <PopupItem>
        <Title>データスタイル設定</Title>
      </PopupItem>
      <SelectionComponent
        selectedDatasets={selectedDatasets}
        savingDataset={savingDataset}
        buildingSearch={buildingSearch}
        templates={templates}
        isMobile
        onDatasetUpdate={onDatasetUpdate}
        onDatasetRemove={onDatasetRemove}
        onDatasetRemoveAll={onDatasetRemoveAll}
        onProjectDatasetsUpdate={onProjectDatasetsUpdate}
        onBuildingSearch={onBuildingSearch}
        onSceneUpdate={onSceneUpdate}
      />
    </Wrapper>
  );
};

export default Selection;

const Wrapper = styled.div`
  border-top: 1px solid #d9d9d9;
  height: calc(100% - 47px);
`;

const Title = styled.p`
  margin: 0;
`;
