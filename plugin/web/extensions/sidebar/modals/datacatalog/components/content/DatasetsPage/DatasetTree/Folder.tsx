import { postMsg } from "@web/extensions/sidebar/utils";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useState, useEffect, useCallback, useMemo } from "react";

import { DataCatalogGroup, DataCatalogItem } from "../../../../api/api";

export type Props = {
  item: DataCatalogGroup;
  id: string;
  name: string;
  isMobile?: boolean;
  nestLevel: number;
  expandedFolders?: { id?: string; name?: string }[];
  selectedID?: string;
  onSelect?: (item: DataCatalogItem | DataCatalogGroup) => void;
  setExpandedFolders?: React.Dispatch<React.SetStateAction<{ id?: string; name?: string }[]>>;
  children?: React.ReactNode;
};

const Folder: React.FC<Props> = ({
  item,
  id,
  name,
  isMobile,
  nestLevel,
  expandedFolders,
  selectedID,
  onSelect,
  setExpandedFolders,
  children,
}) => {
  const [isOpen, open] = useState(false);

  const findCb = useCallback(
    (item: { id?: string; name?: string }) => (item.id ? item.id === id : item.name === name),
    [id, name],
  );

  useEffect(() => {
    if (!!expandedFolders?.find(findCb) && !isOpen) {
      open(true);
    } else if (!expandedFolders?.find(findCb) && isOpen) {
      open(false);
    }
  }, [expandedFolders, findCb, id, isOpen, name]);

  const selected = useMemo(() => selectedID === item.id, [selectedID, item]);

  const handleExpand = useCallback(
    (folder: { id?: string; name?: string }) => {
      onSelect?.(item);
      setExpandedFolders?.((prevState: { id?: string; name?: string }[]) => {
        const newExpandedFolders = [...prevState];
        const index = prevState.findIndex(findCb);
        index >= 0 ? newExpandedFolders.splice(index, 1) : newExpandedFolders.push(folder);
        postMsg({
          action: "saveExpandedFolders",
          payload: { expandedFolders: newExpandedFolders },
        });
        return newExpandedFolders;
      });
    },
    [item, findCb, onSelect, setExpandedFolders],
  );

  return (
    <Wrapper key={id} isOpen={isOpen}>
      <FolderItem
        selected={selected}
        nestLevel={nestLevel}
        onClick={() => handleExpand({ id, name })}>
        <NameWrapper isMobile={isMobile}>
          <Icon icon={isOpen ? "folderOpen" : "folder"} size={20} />
          <Name>{name}</Name>
        </NameWrapper>
      </FolderItem>
      {isOpen ? children : null}
    </Wrapper>
  );
};

export default Folder;

const Wrapper = styled.div<{ isOpen?: boolean }>`
  width: 100%;
  ${({ isOpen }) =>
    isOpen
      ? "height: 100%;"
      : `
  height: 29px; 
  overflow: hidden;
  `}
`;

const FolderItem = styled.div<{ nestLevel: number; selected?: boolean }>`
  display: flex;
  align-items: center;
  justify-content: space-between;
  box-sizing: border-box;
  gap: 8px;
  height: 29px;

  ${({ selected }) =>
    selected &&
    `
  background: #00BEBE;
  color: white;
  `}

  padding-left: ${({ nestLevel }) => (nestLevel ? `${nestLevel * 8}px` : "8px")};
  padding-right: 8px;
  cursor: pointer;

  :hover {
    background: #00bebe;
    color: white;
  }
`;

const NameWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  align-items: center;
  gap: 8px;
  flex: 1;
`;

const Name = styled.p`
  margin: 0;
  user-select: none;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
`;
