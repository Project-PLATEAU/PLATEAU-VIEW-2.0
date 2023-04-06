import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useMemo } from "react";

import type { Mode } from "../../hooks";

export type TabProps = {
  mode: Mode;
  text: string;
  icon: string;
  currentMode: Mode;
  theme?: string;
  onClick: (m: Mode) => void;
};

const Tab: React.FC<TabProps> = ({ mode, text, icon, currentMode, theme, onClick }) => {
  const handleClick = useCallback(() => {
    onClick(mode);
  }, [mode, onClick]);

  const active = useMemo(() => currentMode === mode, [currentMode, mode]);

  return (
    <StyledTab>
      <TabContent onClick={handleClick} active={active} theme={theme}>
        <Icon icon={icon} size={16} />
        <Text>{text}</Text>
      </TabContent>
    </StyledTab>
  );
};

const StyledTab = styled.div`
  display: flex;
  height: 100%;
  align-items: flex-end;
`;

const TabContent = styled.div<{ active: boolean; theme?: string }>`
  display: flex;
  align-items: center;
  gap: 4px;
  height: 37px;
  padding: 8px 12px;
  cursor: pointer;
  background: ${({ active, theme }) =>
    active ? (theme === "grey" ? "#F4F4F4" : "#FFFFFF") : "#dfdfdf"};
  color: ${({ active }) => (active ? "var(--theme-color)" : "#898989")};
  border-width: 1px 1px 0px 1px;
  border-style: solid;
  border-color: ${({ active, theme }) =>
    active ? (theme === "grey" ? "#F4F4F4" : "#FFFFFF") : "#c8c8c8"};
  border-radius: 3px 3px 0px 0px;

  :hover {
    color: ${({ active }) => (active ? "var(--theme-color)" : "#666")};
  }
`;

const Text = styled.div`
  font-weight: 500;
  font-size: 14px;
  line-height: 21px;
`;

export default Tab;
