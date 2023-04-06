import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

import ControlButton from "./ControlButton";
import useHooks from "./hooks";

const PedestrianController: React.FC = () => {
  const {
    mode,
    mainButtonText,
    moveForwardOn,
    moveBackwardOn,
    moveLeftOn,
    moveRightOn,
    moveUpOn,
    moveDownOn,
    miniMapViewRotate,
    visible,
    handleMoveForwardClick,
    handleMoveBackwardClick,
    handleMoveLeftClick,
    handleMoveRightClick,
    handleMoveUpClick,
    handleMoveDownClick,
    onClose,
    onMainButtonClick,
  } = useHooks();

  return (
    <Wrapper visible={visible}>
      <Header>
        <TitleWrapper>
          <Icon icon="personSimpleWalk" size={20} />
          <Title>歩行者モード</Title>
        </TitleWrapper>
        <CloseButton onClick={onClose}>
          <Icon icon="cross" />
        </CloseButton>
      </Header>
      <Content>
        <ControlButton onClick={onMainButtonClick} icon="crosshair" text={mainButtonText} />
        <MouseTip>
          <Icon icon="mousetip" size={54} />
        </MouseTip>
        <Discription>
          地図上のスタートしたい地点をクリックします。マウス操作で右や左に曲がります。
        </Discription>
        <Directions>
          <Line>
            <EmptySpace />
            <ControlButton
              icon="arrowUpRegular"
              text="W"
              disabled={mode !== "pedestrian"}
              active={moveForwardOn}
              onClick={handleMoveForwardClick}
            />
            <EmptySpace />
          </Line>
          <Line>
            <ControlButton
              icon="arrowLeftRegular"
              text="A"
              disabled={mode !== "pedestrian"}
              active={moveLeftOn}
              onClick={handleMoveLeftClick}
            />
            <ControlButton
              icon="arrowDownRegular"
              text="S"
              disabled={mode !== "pedestrian"}
              active={moveBackwardOn}
              onClick={handleMoveBackwardClick}
            />
            <ControlButton
              icon="arrowRightRegular"
              text="D"
              disabled={mode !== "pedestrian"}
              active={moveRightOn}
              onClick={handleMoveRightClick}
            />
          </Line>
        </Directions>
        <UpAndDown>
          <Line>
            <ControlButton
              icon="arrowLineUpRegular"
              text="Space"
              disabled={mode !== "pedestrian"}
              active={moveUpOn}
              onClick={handleMoveUpClick}
            />
            <ControlButton
              icon="arrowLineDownRegular"
              text="Shift"
              disabled={mode !== "pedestrian"}
              active={moveDownOn}
              onClick={handleMoveDownClick}
            />
          </Line>
        </UpAndDown>
        <MiniMapWrapper hidden={mode !== "pedestrian"}>
          <ViewIndicator>
            <StyledIcon icon={"viewIndicator"} size={20} rotate={miniMapViewRotate} />
          </ViewIndicator>
          <MiniMap id="minimap" />
        </MiniMapWrapper>
      </Content>
    </Wrapper>
  );
};

const Wrapper = styled.div<{ visible: boolean }>`
  width: 100%;
  height: 100%;
  background-color: #fff;
  border-radius: 4px;
  opacity: ${({ visible }) => (visible ? 1 : 0)};
  transition: all 0.25s ease;
`;

const Header = styled.div`
  display: flex;
  align-items: center;
  justify-content: space-between;
  padding: 12px;
`;

const TitleWrapper = styled.div`
  display: flex;
  gap: 8px;
  align-items: center;
`;

const Title = styled.div`
  font-size: 14px;
  font-weight: 700;
  color: #262626;
`;

const CloseButton = styled.div`
  cursor: pointer;
`;

const Content = styled.div`
  display: flex;
  flex-direction: column;
  padding: 0 12px 12px 12px;
  gap: 12px;
`;

const MouseTip = styled.div`
  width: 100%;
  display: flex;
  justify-content: center;
  color: #595959;
`;

const Discription = styled.div`
  color: #595959;
  font-size: 14px;
  line-height: 22px;
`;

const Directions = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Line = styled.div`
  display: flex;
  justify-content: space-between;
  gap: 8px;
`;

const EmptySpace = styled.div`
  width: 100%;
`;

const UpAndDown = styled.div``;

const MiniMapWrapper = styled.div<{ hidden: boolean }>`
  position: relative;
  height: 200px;
  width: 100%;
  opacity: ${({ hidden }) => (hidden ? 0 : 1)};
  border-radius: 4px;
  overflow: hidden;
`;

const ViewIndicator = styled.div`
  position: absolute;
  display: flex;
  align-items: center;
  justify-content: center;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  z-index: 401;
`;

const StyledIcon = styled(Icon)<{ rotate?: number }>`
  transform: ${({ rotate }) => `rotate(${rotate ?? 0}deg)`};
`;

const MiniMap = styled.div`
  position: relative;
  width: 184px;
  height: 200px;
`;

export default PedestrianController;
