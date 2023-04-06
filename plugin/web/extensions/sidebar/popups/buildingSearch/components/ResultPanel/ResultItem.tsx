import { styled } from "@web/theme";
import { useCallback, useMemo } from "react";

import { Result } from "../../types";

type Props = {
  item: Result;
  selected: Result[];
  hasBorderBottom: boolean;
  onSelect: (selected: Result[]) => void;
};

const ResultItem: React.FC<Props> = ({ item, selected, hasBorderBottom, onSelect }) => {
  const onClick = useCallback(() => {
    if (selected.find(s => s.gml_id === item.gml_id)) {
      onSelect(selected.filter(s => s.gml_id !== item.gml_id));
    } else {
      onSelect([item]);
    }
  }, [onSelect, item, selected]);

  const isActive = useMemo(() => {
    return !!selected.find(s => s.gml_id === item.gml_id);
  }, [selected, item]);

  return (
    <StyledResultItem onClick={onClick} active={isActive} hasBorderBottom={hasBorderBottom}>
      {item.gml_id}
    </StyledResultItem>
  );
};

const StyledResultItem = styled.div<{ active: boolean; hasBorderBottom: boolean }>`
  display: flex;
  align-items: center;
  width: 100%;
  height: 38px;
  padding: 0 12px;
  border-bottom: 1px solid #d9d9d9;
  font-size: 12px;
  background: ${({ active }) => (active ? "var(--theme-color)" : "#fff")};
  color: ${({ active }) => (active ? "#fff" : "#000")};
  cursor: pointer;

  &:last-child {
    border-bottom: ${({ hasBorderBottom }) => (hasBorderBottom ? "1px solid #d9d9d9" : "none")};
  }
`;

export default ResultItem;
