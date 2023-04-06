import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback } from "react";

import type { Result } from "../../types";

import ResultItem from "./ResultItem";

type Props = {
  active: boolean;
  results: Result[];
  highlightAll: boolean;
  showMatchingOnly: boolean;
  selected: Result[];
  isSearching: boolean;
  setHighlightAll: React.Dispatch<React.SetStateAction<boolean>>;
  setShowMatchingOnly: React.Dispatch<React.SetStateAction<boolean>>;
  setSelected: React.Dispatch<React.SetStateAction<Result[]>>;
};

const ResultPanel: React.FC<Props> = ({
  active,
  results,
  highlightAll,
  showMatchingOnly,
  selected,
  isSearching,
  setHighlightAll,
  setShowMatchingOnly,
  setSelected,
}) => {
  const onHighlightAll = useCallback(() => {
    setHighlightAll(highlightAll => {
      if (!highlightAll && selected.length > 0) {
        setSelected([]);
      }
      return !highlightAll;
    });
  }, [setHighlightAll, selected, setSelected]);

  const onShowMatchingOnly = useCallback(() => {
    setShowMatchingOnly(showMatchingOnly => !showMatchingOnly);
  }, [setShowMatchingOnly]);

  const onSelect = useCallback(
    (selected: Result[]) => {
      setSelected(selected);
      if (highlightAll) {
        setHighlightAll(false);
      }
    },
    [setSelected, highlightAll, setHighlightAll],
  );

  return (
    <Wrapper active={active}>
      <ResultInfo>{isSearching ? "検索中..." : `${results.length} 件が見つかりました`}</ResultInfo>
      <ResultWrapper>
        {!isSearching &&
          results?.map((item, index) => (
            <ResultItem
              key={index}
              item={item}
              onSelect={onSelect}
              selected={selected}
              hasBorderBottom={results.length < 10}
            />
          ))}
        {!isSearching && results.length === 0 && (
          <EmptyWrapper>
            <Empty>
              <Icon icon="fileDotted" size={24} />
              <EmptyInfo>検索結果がありません</EmptyInfo>
            </Empty>
          </EmptyWrapper>
        )}
      </ResultWrapper>
      <ButtonWrapper>
        <Button
          active={highlightAll}
          onClick={onHighlightAll}
          disabled={isSearching || results.length === 0}>
          結果をハイライト表示
        </Button>
        <Button
          active={showMatchingOnly}
          onClick={onShowMatchingOnly}
          disabled={isSearching || results.length === 0}>
          結果のみ表示
        </Button>
      </ButtonWrapper>
    </Wrapper>
  );
};

const Wrapper = styled.div<{ active: boolean }>`
  display: ${({ active }) => (active ? "flex" : "none")};
  padding: 8px 0;
  flex-direction: column;
`;

const ResultInfo = styled.div`
  height: 40px;
  display: flex;
  align-items: flex-start;
  justify-content: flex-start;
  flex-shrink: 0;
  width: 100%;
  padding: 9px 20px;
  font-size: 16px;
`;

const ResultWrapper = styled.div`
  position: relative;
  height: 350px;
  overflow: auto;
  border: 1px solid #d9d9d9;
  margin: 0 12px;
`;

const EmptyWrapper = styled.div`
  position: absolute;
  width: 100%;
  height: 100%;
  display: flex;
  align-items: center;
  justify-content: center;
`;

const Empty = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
  color: #bfbfbf;
`;

const EmptyInfo = styled.span``;

const ButtonWrapper = styled.div`
  padding: 6px 12px;
  display: flex;
  gap: 12px;
`;

const Button = styled.div<{ active: boolean; disabled?: boolean }>`
  width: 50%;
  height: 32px;
  display: flex;
  align-items: center;
  justify-content: center;
  color: ${({ active, disabled }) => (disabled ? "rgba(0, 0, 0, 0.25)" : active ? "#fff" : "#000")};
  background: ${({ active, disabled }) =>
    disabled ? "none" : active ? "var(--theme-color)" : "#fff"};
  border: ${({ active, disabled }) =>
    disabled ? "1px solid #D9D9D9" : active ? "1px solid var(--theme-color)" : "1px solid #e6e6e6"};
  border-radius: 4px;
  cursor: ${({ disabled }) => (disabled ? "default" : "pointer")};
  pointer-events: ${({ disabled }) => (disabled ? "none" : "all")};
`;

export default ResultPanel;
