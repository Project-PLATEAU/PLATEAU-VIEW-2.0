import { DataCatalogGroup, DataCatalogItem } from "../../../../api/api";

import File from "./File";
import Folder from "./Folder";

type Props = {
  catalogItem: DataCatalogGroup | DataCatalogItem | (DataCatalogItem | DataCatalogGroup)[];
  isMobile?: boolean;
  addedDatasetDataIDs?: string[];
  selectedID?: string;
  nestLevel: number;
  expandedFolders?: { id?: string; name?: string }[];
  addDisabled: (dataID: string) => boolean;
  onDatasetAdd: (dataset: DataCatalogItem, keepModalOpen?: boolean) => void;
  onSelect?: (item: DataCatalogItem | DataCatalogGroup) => void;
  setExpandedFolders?: React.Dispatch<React.SetStateAction<{ id?: string; name?: string }[]>>;
};

const TreeBuilder: React.FC<Props> = ({
  catalogItem,
  isMobile,
  addedDatasetDataIDs,
  selectedID,
  nestLevel,
  expandedFolders,
  addDisabled,
  onDatasetAdd,
  onSelect,
  setExpandedFolders,
}) => {
  return (
    <>
      {Array.isArray(catalogItem) ? (
        catalogItem.map(item =>
          "children" in item ? (
            <Folder
              item={item}
              key={item.id}
              id={item.id}
              name={item.name}
              nestLevel={nestLevel + 1}
              expandedFolders={expandedFolders}
              isMobile={isMobile}
              selectedID={selectedID}
              onSelect={onSelect}
              setExpandedFolders={setExpandedFolders}>
              <TreeBuilder
                catalogItem={item.children}
                addedDatasetDataIDs={addedDatasetDataIDs}
                selectedID={selectedID}
                nestLevel={nestLevel + 1}
                expandedFolders={expandedFolders}
                addDisabled={addDisabled}
                onDatasetAdd={onDatasetAdd}
                onSelect={onSelect}
                setExpandedFolders={setExpandedFolders}
              />
            </Folder>
          ) : (
            <TreeBuilder
              catalogItem={item}
              addedDatasetDataIDs={addedDatasetDataIDs}
              selectedID={selectedID}
              nestLevel={nestLevel + 1}
              expandedFolders={expandedFolders}
              addDisabled={addDisabled}
              onDatasetAdd={onDatasetAdd}
              onSelect={onSelect}
              setExpandedFolders={setExpandedFolders}
            />
          ),
        )
      ) : "children" in catalogItem ? (
        <Folder
          item={catalogItem}
          key={catalogItem.id}
          id={catalogItem.id}
          name={catalogItem.name}
          nestLevel={nestLevel + 1}
          expandedFolders={expandedFolders}
          isMobile={isMobile}
          selectedID={selectedID}
          onSelect={onSelect}
          setExpandedFolders={setExpandedFolders}>
          <TreeBuilder
            catalogItem={catalogItem.children}
            addedDatasetDataIDs={addedDatasetDataIDs}
            selectedID={selectedID}
            nestLevel={nestLevel + 1}
            expandedFolders={expandedFolders}
            addDisabled={addDisabled}
            onDatasetAdd={onDatasetAdd}
            onSelect={onSelect}
            setExpandedFolders={setExpandedFolders}
          />
        </Folder>
      ) : (
        <File
          item={catalogItem}
          isMobile={isMobile}
          nestLevel={nestLevel}
          selectedID={selectedID}
          addDisabled={addDisabled}
          onDatasetAdd={onDatasetAdd}
          onSelect={onSelect}
        />
      )}
    </>
  );
};

export default TreeBuilder;
