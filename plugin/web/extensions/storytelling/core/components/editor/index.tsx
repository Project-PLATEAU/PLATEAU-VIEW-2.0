import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useRef, useCallback, type WheelEvent } from "react";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import type { Camera, Scene as SceneType } from "../../../types";

import Scene from "./scene";

type Props = {
  scenes: SceneType[];
  sceneCapture: () => void;
  sceneView: (camera: Camera) => void;
  sceneRecapture: (id: string) => void;
  sceneDelete: (id: string) => void;
  sceneEdit: (id: string) => void;
  sceneMove: (dragIndex: number, hoverIndex: number) => void;
};

const Editor: React.FC<Props> = ({
  scenes,
  sceneCapture,
  sceneView,
  sceneRecapture,
  sceneDelete,
  sceneEdit,
  sceneMove,
}) => {
  const wrapperRef = useRef<HTMLDivElement>(null);

  const handleWheel = useCallback((e: WheelEvent) => {
    if (wrapperRef.current) {
      wrapperRef.current.scrollLeft += e.deltaY < 0 ? -30 : 30;
    }
  }, []);

  return (
    <Wrapper onWheel={handleWheel} ref={wrapperRef}>
      <DndProvider backend={HTML5Backend}>
        <Content>
          {scenes?.map((scene, index) => (
            <ItemWrapper key={scene.id}>
              <Scene
                index={index}
                sceneView={sceneView}
                sceneRecapture={sceneRecapture}
                sceneDelete={sceneDelete}
                sceneEdit={sceneEdit}
                sceneMove={sceneMove}
                {...scene}
              />
            </ItemWrapper>
          ))}
          <CreateStoryWrapper>
            <CreateStory onClick={sceneCapture}>
              <Icon icon="cornersOut" size={24} />
              <CreateText>キャプチャ</CreateText>
            </CreateStory>
          </CreateStoryWrapper>
        </Content>
      </DndProvider>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  height: 100%;
  flex: 1;
  overflow-x: auto;
  overflow-y: hidden;
  padding: 12px 6px;
  scrollbar-width: thin;
`;

const Content = styled.div`
  display: flex;
  height: 100%;
`;

const ItemWrapper = styled.div`
  padding: 0 6px;
`;
const CreateStoryWrapper = styled.div`
  padding: 0 12px 0 6px;
`;

const CreateStory = styled.div`
  width: 170px;
  height: 114px;
  background: #fff;
  border-radius: 8px;
  border: 1px solid var(--theme-color);
  color: var(--theme-color);
  cursor: pointer;
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  gap: 4px;
`;

const CreateText = styled.div`
  font-weight: 500;
  font-size: 14px;
  line-height: 21px;
`;

export default Editor;
