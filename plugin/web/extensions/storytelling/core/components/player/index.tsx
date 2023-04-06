import { Carousel, Icon, Pagination } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useEffect, useRef, useState, useMemo } from "react";
import { Remarkable } from "remarkable";

import type { Camera, Scene as SceneType } from "../../../types";

type Props = {
  scenes: SceneType[];
  isMobile: boolean;
  contentWidth: number;
  sceneView: (camera: Camera) => void;
  setPlayerHeight: (height: number) => void;
};

const Player: React.FC<Props> = ({
  scenes,
  isMobile,
  contentWidth,
  sceneView,
  setPlayerHeight,
}) => {
  const minCarouselHeight = 131;
  const maxCarouselHeight = useMemo(() => {
    return isMobile ? 162 : 331;
  }, [isMobile]);

  const sceneRefs = useRef<HTMLDivElement[]>([]);
  const setSceneRef = useCallback((dom: HTMLDivElement) => {
    sceneRefs.current.push(dom);
  }, []);

  const sceneTitleRefs = useRef<HTMLDivElement[]>([]);
  const setSceneTitleRef = useCallback((dom: HTMLDivElement) => {
    sceneTitleRefs.current.push(dom);
  }, []);

  const sceneContentRefs = useRef<HTMLDivElement[]>([]);
  const setSceneContentRef = useCallback((dom: HTMLDivElement) => {
    sceneContentRefs.current.push(dom);
  }, []);

  const updateHeight = useCallback(
    (index: number) => {
      const outerHeight = isMobile ? 12 + 26 : 24 + 40;
      if (scenes.length > 0 && sceneRefs.current[index]) {
        let carouselHeight =
          sceneTitleRefs.current[index].clientHeight +
          sceneContentRefs.current[index].clientHeight +
          8 + // gap
          24 + // description padding bottom
          12; // space

        carouselHeight =
          carouselHeight > maxCarouselHeight
            ? maxCarouselHeight
            : carouselHeight < minCarouselHeight
            ? minCarouselHeight
            : carouselHeight;
        sceneRefs.current[index].style.height = `${carouselHeight}px`;
        setPlayerHeight(carouselHeight + outerHeight);
      } else {
        setPlayerHeight(minCarouselHeight + outerHeight);
      }
    },
    [scenes, setPlayerHeight, isMobile, maxCarouselHeight],
  );

  const carouselRef = useRef<any>(null);
  const [current, setCurrent] = useState<number>(0);
  const currentRef = useRef<number>(current);
  currentRef.current = current;

  const onSlideChange = useCallback(
    (oldSlide: number, currentSlide: number) => {
      if (currentSlide !== current) {
        const camera = scenes[currentSlide]?.camera;
        if (camera) {
          sceneView(camera);
        }
        setCurrent(currentSlide);
        updateHeight(currentSlide);
      }
    },
    [scenes, sceneView, current, setCurrent, updateHeight],
  );

  const prev = useCallback(() => {
    if (carouselRef.current) {
      carouselRef.current.prev();
    }
  }, []);

  const next = useCallback(() => {
    if (carouselRef.current) {
      carouselRef.current.next();
    }
  }, []);

  const onPaginationChange = useCallback((current: number) => {
    if (carouselRef.current) {
      carouselRef.current.goTo(current - 1);
    }
  }, []);

  const md = useRef(
    new Remarkable({
      html: false,
      breaks: true,
      typographer: true,
      linkTarget: "_blank",
    }),
  );

  const [carouselWidth, setCarouselWidth] = useState<number>(document.body.clientWidth);
  useEffect(() => {
    setCarouselWidth(isMobile ? contentWidth - 12 - 12 - 48 : contentWidth - 24 - 24 - 80);
  }, [contentWidth, isMobile]);

  useEffect(() => {
    if (scenes.length === 0) {
      sceneRefs.current = [];
      sceneTitleRefs.current = [];
      sceneContentRefs.current = [];
      carouselRef.current = undefined;
      updateHeight(0);
    } else {
      if (carouselRef.current && currentRef.current !== 0) {
        carouselRef.current.goTo(0);
      } else {
        if (scenes[0]?.camera) {
          sceneView(scenes[0].camera);
        }
        updateHeight(0);
      }
    }
  }, [scenes, sceneView, updateHeight]);

  return (
    <Wrapper isMobile={isMobile}>
      <NavButton onClick={prev} disabled={current === 0} isMobile={isMobile}>
        <Icon icon="caretLeft" size={isMobile ? 24 : 32} />
      </NavButton>
      <MainContent isMobile={isMobile}>
        <CarouselContainer>
          <CarouselArea minWidth={carouselWidth}>
            {scenes.length > 0 && (
              <Carousel
                beforeChange={onSlideChange}
                dots={false}
                ref={carouselRef}
                infinite={false}
                draggable={true}
                speed={250}>
                {scenes.map((scene, index) => (
                  <div key={index}>
                    <StoryItem ref={setSceneRef}>
                      <Title ref={setSceneTitleRef}>{scene.title}</Title>
                      <Description>
                        <div
                          ref={setSceneContentRef}
                          dangerouslySetInnerHTML={{
                            __html: md.current.render(scene.description),
                          }}
                        />
                      </Description>
                    </StoryItem>
                  </div>
                ))}
              </Carousel>
            )}
          </CarouselArea>
        </CarouselContainer>
        <PaginationContainer>
          {scenes.length > 0 && (
            <StyledPagination
              current={current + 1}
              size="small"
              total={scenes.length}
              pageSize={1}
              showSizeChanger={false}
              onChange={onPaginationChange}
            />
          )}
        </PaginationContainer>
      </MainContent>
      <NavButton
        onClick={next}
        disabled={current >= scenes.length - 1}
        className="next"
        isMobile={isMobile}>
        <Icon icon="caretLeft" size={isMobile ? 24 : 32} />
      </NavButton>
    </Wrapper>
  );
};

const Wrapper = styled.div<{ isMobile: boolean }>`
  position: relative;
  display: flex;
  justify-content: space-between;
  height: 100%;
  padding: ${({ isMobile }) => (isMobile ? "6px" : "12px")};
  gap: ${({ isMobile }) => (isMobile ? "6px" : "12px")};
`;

const NavButton = styled.div<{ disabled: boolean; isMobile: boolean }>`
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  width: ${({ isMobile }) => (isMobile ? "24px" : "40px")};
  color: ${({ disabled }) => (disabled ? "#ccc" : "var(--theme-color)")};
  pointer-events: ${({ disabled }) => (disabled ? "none" : "all")};
  cursor: pointer;
  z-index: 2;

  &.next {
    transform: rotate(180deg);
  }
`;

const MainContent = styled.div<{ isMobile: boolean }>`
  position: relative;
  height: 100%;
  width: 100%;
  overflow: hidden;
  border-radius: 6px;
  border: ${({ isMobile }) =>
    isMobile ? "1px solid rgba(0, 0, 0, 0.1)" : "1px solid rgba(0, 0, 0, 0.45)"};
`;

const CarouselContainer = styled.div`
  height: 100%;
  width: 100%;
  overflow: hidden;
`;

const CarouselArea = styled.div<{ minWidth: number }>`
  position: absolute;
  width: 100%;
  height: 100%;
  min-width: ${({ minWidth }) => `${minWidth}px`};
`;

const PaginationContainer = styled.div`
  position: absolute;
  display: flex;
  flex-direction: row-reverse;
  right: 6px;
  bottom: 1px;
  background-color: #fff;
`;

const StyledPagination = styled(Pagination)`
  .ant-pagination-prev,
  .ant-pagination-next {
    display: none;
  }

  .ant-pagination-item-active {
    border-color: rgba(0, 0, 0, 0);
    background: none;
  }

  .ant-pagination-item:hover a,
  .ant-pagination-item-active a {
    color: var(--theme-color);
  }
`;

const StoryItem = styled.div`
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Title = styled.div`
  font-size: 14px;
  font-weight: 700;
  line-height: 19px;
  flex-shrink: 0;
  padding: 12px 12px 0;
`;

const Description = styled.div`
  height: 100%;
  overflow: auto;
  font-size: 12px;
  line-height: 1.5;
  padding: 0 12px 24px;

  img {
    max-width: 100%;
  }
`;

export default Player;
