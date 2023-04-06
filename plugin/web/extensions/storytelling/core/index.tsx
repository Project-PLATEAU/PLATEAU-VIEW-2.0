import Editor from "@web/extensions/storytelling/core/components/editor";
import Header from "@web/extensions/storytelling/core/components/header";
import Player from "@web/extensions/storytelling/core/components/player";
import useHooks, { Size, sizes } from "@web/extensions/storytelling/core/hooks";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

const Storytelling: React.FC = () => {
  const {
    size,
    mode,
    minimized,
    scenes,
    isMobile,
    contentWidth,
    setPlayerHeight,
    handleMinimize,
    handleSetMode,
    sceneCapture,
    sceneView,
    sceneRecapture,
    sceneDelete,
    sceneEdit,
    sceneMove,
    storyClear,
    storyShare,
  } = useHooks();

  return (
    <Wrapper size={size} mode={mode} minimized={minimized}>
      <MiniPane onClick={handleMinimize} minimized={minimized}>
        <Icon icon="cornersOut" color="#4A4A4A" size={24} />
        <MiniTitle>ストーリー</MiniTitle>
      </MiniPane>
      <ContentPane minimized={minimized}>
        <Header
          mode={mode}
          setMode={handleSetMode}
          storyShare={storyShare}
          storyClear={storyClear}
          handleMinimize={handleMinimize}
          isMobile={isMobile}
        />
        {!isMobile && mode === "editor" && (
          <Editor
            scenes={scenes}
            sceneCapture={sceneCapture}
            sceneView={sceneView}
            sceneRecapture={sceneRecapture}
            sceneDelete={sceneDelete}
            sceneEdit={sceneEdit}
            sceneMove={sceneMove}
          />
        )}
        {mode === "player" && (
          <Player
            scenes={scenes}
            isMobile={isMobile}
            contentWidth={contentWidth}
            sceneView={sceneView}
            setPlayerHeight={setPlayerHeight}
          />
        )}
      </ContentPane>
    </Wrapper>
  );
};

const Wrapper = styled.div<{
  size: Size;
  mode?: string;
  minimized: boolean;
}>`
  position: relative;
  display: inline-block;
  border-radius: 8px;
  background: ${({ mode, minimized }) => (minimized || mode === "editor" ? "#fff" : "#F4F4F4")};
  transition: min-width 0.5s, min-height 0.5s;
  min-width: ${({ minimized }) => (minimized ? `${sizes.mini.width}px` : "100%")};
  min-height: ${({ size }) => `${size.height}px`};
  overflow: hidden;
`;

const MiniPane = styled.div<{ minimized: boolean }>`
  position: absolute;
  left: 0;
  top: 0;
  display: flex;
  align-items: center;
  gap: 4px;
  padding: 8px 12px;
  width: ${sizes.mini.width};
  cursor: pointer;
  pointer-events: ${({ minimized }) => (minimized ? "all" : "none")};
  opacity: ${({ minimized }) => (minimized ? 1 : 0)};
  transition: opacity 0.25s;
`;

const MiniTitle = styled.div`
  font-weight: 700;
  font-size: 14px;
  width: auto;
  white-space: nowrap;
`;

const ContentPane = styled.div<{ minimized: boolean }>`
  position: absolute;
  width: 100%;
  height: 100%;
  left: 0;
  top: 0;
  display: flex;
  flex-direction: column;
  pointer-events: ${({ minimized }) => (minimized ? "none" : "all")};
  opacity: ${({ minimized }) => (minimized ? 0 : 1)};
  transition: opacity 0.25s;
`;

export default Storytelling;
