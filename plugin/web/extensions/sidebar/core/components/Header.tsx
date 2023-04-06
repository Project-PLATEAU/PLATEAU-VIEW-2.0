import { Icon, Col } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { memo, ReactNode, useMemo } from "react";

export type Pages = "data" | "map" | "share" | "help" | "feedback" | "template";

export type TabProps = {
  key: Pages;
  icon: ReactNode;
};

type Props = {
  className?: string;
  current: string;
  inEditor?: boolean;
  minimized: boolean;
  onMinimize: () => void;
  onClick: (p: Pages) => void;
};

const Header: React.FC<Props> = ({
  className,
  current,
  inEditor,
  minimized,
  onMinimize,
  onClick,
}) => {
  const headerItems = useMemo(() => {
    const items: TabProps[] = [
      {
        key: "data",
        icon: <StyledIcon icon="database" />,
      },
      {
        key: "map",
        icon: <StyledIcon icon="sliders" />,
      },
      {
        key: "share",
        icon: <StyledIcon icon="share" />,
      },
      {
        key: "help",
        icon: <StyledIcon icon="help" />,
      },
      {
        key: "feedback",
        icon: <StyledIcon icon="feedback" />,
      },
      {
        key: "template",
        icon: <StyledIcon icon="templates" />,
      },
    ];

    return !inEditor ? [...items.slice(0, -1)] : [...items];
  }, [inEditor]);

  return (
    <HeaderWrapper className={className} minimized={minimized}>
      <TopSection>
        <PlateauIcon icon="plateauLogo" size={114} wide />
        <MinimizeButton minimized={minimized}>
          <Icon icon={minimized ? "menu" : "close"} onClick={onMinimize} />
        </MinimizeButton>
      </TopSection>
      {!minimized && (
        <Nav>
          {headerItems?.map((i, idx) => (
            <IconWrapper key={idx} current={current === i.key} onClick={() => onClick(i.key)}>
              {i.icon}
            </IconWrapper>
          ))}
        </Nav>
      )}
    </HeaderWrapper>
  );
};

export default memo(Header);

const HeaderWrapper = styled(Col)<{ minimized?: boolean }>`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  align-items: center;
  height: 164px;
  width: 100%;
  background-color: #ffff;
  border-radius: ${({ minimized }) => (minimized ? "0 0 8px 0" : 0)};
  transition: height 0.5s, width 0.5s, border-radius 0.5s;
`;

const Nav = styled.div`
  display: flex;
  justify-content: space-around;
  width: 100%;
  padding: 0 53px;
`;

const TopSection = styled.div`
  display: flex;
  position: relative;
  width: 100%;
  height: 100%;
  padding: 24px;
  border-radius: 0 0 8px 0;
`;

const MinimizeButton = styled.div<{ minimized?: boolean }>`
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  right: 0;
  top: 0;
  border: none;
  height: 32px;
  width: 32px;
  background: #00bebe;
  cursor: pointer;
  transition: background 0.3s;
  color: white;

  ${({ minimized }) => minimized && "position: static;"}
`;

const PlateauIcon = styled(Icon)<{ minimized?: boolean }>`
  margin: auto;

  ${({ minimized }) => minimized && "text-align: left;"}
`;

const IconWrapper = styled.div<{ current?: boolean }>`
  cursor: pointer;
  padding: 6px;
  border-bottom-style: solid;
  border-bottom-width: 1px;
  border-bottom-color: ${({ current }) => (current ? "#00bebe" : "transparent")};
  color: ${({ current }) => (current ? "#00bebe" : "#C7C5C5")};
  transition: border-bottom-color 0.5s, color 0.5s;

  :hover {
    border-bottom-color: #00bebe;
    color: #00bebe;
  }
`;

const StyledIcon = styled(Icon)`
  width: 100%;
`;
