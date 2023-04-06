import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

import ConditionPanel from "./ConditionPanel";
import useHooks, { type Size } from "./hooks";
import ResultPanel from "./ResultPanel";

const BuildingSearch: React.FC = () => {
  const {
    minimized,
    size,
    activeTab,
    dataset,
    results,
    highlightAll,
    showMatchingOnly,
    selected,
    isSearching,
    conditionsState,
    onClickCondition,
    onClickResult,
    toggleMinimize,
    popupClose,
    setConditions,
    conditionApply,
    setHighlightAll,
    setShowMatchingOnly,
    setSelected,
  } = useHooks();

  return (
    <Wrapper size={size}>
      <Header>
        <TitleWrapper>
          <Icon icon="magnifyingGlass" size={20} />
          <Title>データを検索</Title>
        </TitleWrapper>
        <ButtonWrapper>
          <Button onClick={toggleMinimize}>
            <Icon icon={minimized ? "rectMaximize" : "rectMinimize"} />
          </Button>
          <Button onClick={popupClose}>
            <Icon icon="cross" />
          </Button>
        </ButtonWrapper>
      </Header>
      <MiniContent active={minimized} disabled={results.length === 0}>
        {results.length === 0 ? `検索結果がありません` : `${results.length} 件が見つかりました`}
      </MiniContent>
      <Content active={!minimized}>
        <Tabs>
          <Tab active={activeTab === "condition"} onClick={onClickCondition}>
            <TabIcon>
              <Icon icon="funnel" size={24} />
            </TabIcon>
            <TabTitle>条件</TabTitle>
          </Tab>
          <Tab active={activeTab === "result"} onClick={onClickResult}>
            <TabIcon>
              <Icon icon="listNumbers" size={24} />
            </TabIcon>
            <TabTitle>結果</TabTitle>
          </Tab>
        </Tabs>
        <TabContent>
          <ConditionPanel
            active={activeTab === "condition"}
            dataset={dataset}
            conditionsState={conditionsState}
            setConditions={setConditions}
            conditionApply={conditionApply}
          />
          <ResultPanel
            active={activeTab === "result"}
            results={results}
            highlightAll={highlightAll}
            showMatchingOnly={showMatchingOnly}
            selected={selected}
            isSearching={isSearching}
            setHighlightAll={setHighlightAll}
            setShowMatchingOnly={setShowMatchingOnly}
            setSelected={setSelected}
          />
        </TabContent>
      </Content>
    </Wrapper>
  );
};

const Wrapper = styled.div<{ size: Size }>`
  display: flex;
  flex-direction: column;
  width: ${({ size }) => `${size.width}px`};
  height: ${({ size }) => `${size.height}px`};
  background-color: #dfdfdf;
  border-radius: 4px;
  overflow: hidden;
  transition: all 0.5s ease;
`;

const Header = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  height: 34px;
  background: #fff;
`;

const TitleWrapper = styled.div`
  display: flex;
  align-items: center;
  gap: 8px;
  color: #000;
  padding: 0 12px;
`;

const Title = styled.div`
  font-size: 16px;
`;

const ButtonWrapper = styled.div`
  display: flex;
  gap: 4px;
`;

const Button = styled.div`
  width: 34px;
  height: 34px;
  display: flex;
  align-items: center;
  justify-content: center;
  background-color: var(--theme-color);
  color: #fff;
  cursor: pointer;
`;

const MiniContent = styled.div<{ active: boolean; disabled: boolean }>`
  display: ${({ active }) => (active ? "flex" : "none")};
  height: 100%;
  align-items: center;
  justify-content: center;
  font-size: 14px;
  color: ${({ disabled }) => (disabled ? "#bfbfbf" : "#000")};
`;

const Content = styled.div<{ active: boolean }>`
  display: ${({ active }) => (active ? "flex" : "none")};
  flex-direction: column;
  height: 100%;
`;

const Tabs = styled.div`
  display: flex;
  align-items: flex-end;
  height: 44px;
  padding: 0 12px;
  gap: 12px;
  flex-shrink: 0;
`;

const Tab = styled.div<{ active: boolean }>`
  display: flex;
  align-items: center;
  gap: 10px;
  height: 40px;
  padding: 0 12px;
  background-color: #fff;
  cursor: pointer;
  background: ${({ active }) => (active ? "#F4F4F4" : "#DFDFDF")};
  color: ${({ active }) => (active ? "var(--theme-color)" : "#8C8C8C")};
  border-width: 1px 1px 0px 1px;
  border-style: solid;
  border-color: ${({ active }) => (active ? "#F4F4F4" : "#BFBFBF")};
  border-radius: 4px 4px 0px 0px;
`;

const TabIcon = styled.div``;

const TabTitle = styled.div``;

const TabContent = styled.div`
  height: 100%;
  background: #f4f4f4;
`;

export default BuildingSearch;
