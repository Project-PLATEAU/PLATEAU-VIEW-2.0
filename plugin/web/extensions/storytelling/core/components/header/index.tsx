import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

import type { Mode } from "../../hooks";

import Tab from "./Tab";

type Props = {
  mode: Mode;
  isMobile: boolean;
  setMode: (m: Mode) => void;
  storyShare: () => void;
  storyClear: () => void;
  handleMinimize: () => void;
};

const Header: React.FC<Props> = ({
  mode,
  isMobile,
  setMode,
  storyShare,
  storyClear,
  handleMinimize,
}) => {
  return (
    <StyledHeader isMobile={isMobile}>
      <HeaderMain isMobile={isMobile}>
        <WidgetTitle isMobile={isMobile}>ストーリー</WidgetTitle>
        {!isMobile && (
          <>
            <Tab
              mode="editor"
              icon="pencil"
              text="編集モード"
              currentMode={mode}
              onClick={setMode}
            />
            <Tab
              mode="player"
              icon="play"
              text="再生モード"
              theme="grey"
              currentMode={mode}
              onClick={setMode}
            />
          </>
        )}
      </HeaderMain>
      <HeaderBtns isMobile={isMobile}>
        {!isMobile && mode === "editor" && (
          <IconBtn onClick={storyClear} isMobile={isMobile}>
            <Icon icon="eraser" size={24} />
          </IconBtn>
        )}
        {!isMobile && (
          <IconBtn onClick={storyShare} isMobile={isMobile}>
            <Icon icon="paperPlane" size={isMobile ? 16 : 24} />
          </IconBtn>
        )}
        <IconBtn onClick={handleMinimize} isMobile={isMobile}>
          <Icon icon="cross" size={isMobile ? 16 : 24} />
        </IconBtn>
      </HeaderBtns>
    </StyledHeader>
  );
};

const StyledHeader = styled.div<{ isMobile: boolean }>`
  height: ${({ isMobile }) => (isMobile ? "26px" : "40px")};
  background: #dfdfdf;
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-shrink: 0;
`;

const HeaderMain = styled.div<{ isMobile: boolean }>`
  display: flex;
  gap: 10px;
  height: 100%;
  min-width: ${({ isMobile }) => (isMobile ? "110px" : "362px")};
`;
const HeaderBtns = styled.div<{ isMobile: boolean }>`
  display: flex;
  gap: ${({ isMobile }) => (isMobile ? "1px" : "2px")};
  height: 100%;
`;

const WidgetTitle = styled.div<{ isMobile: boolean }>`
  display: flex;
  align-items: center;
  color: #4a4a4a;
  font-weight: 700;
  font-size: 14px;
  padding: ${({ isMobile }) => (isMobile ? "0 10px" : "0 20px")};
`;

const IconBtn = styled.div<{ isMobile: boolean }>`
  display: flex;
  width: ${({ isMobile }) => (isMobile ? "26px" : "40px")};
  height: 100%;
  align-items: center;
  justify-content: center;
  background: var(--theme-color);
  color: #fff;
  cursor: pointer;
  line-height: 1;
`;

export default Header;
