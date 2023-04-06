import { Icon, Dropdown, Menu } from "@web/sharedComponents";
import { styled } from "@web/theme";
import type { Identifier, XYCoord } from "dnd-core";
import { useCallback, useRef } from "react";
import { useDrag, useDrop } from "react-dnd";

import type { Camera, Scene as SceneType } from "../../../types";

type Props = SceneType & {
  index: number;
  sceneView: (camera: Camera) => void;
  sceneRecapture: (id: string) => void;
  sceneDelete: (id: string) => void;
  sceneEdit: (id: string) => void;
  sceneMove: (dragIndex: number, hoverIndex: number) => void;
};

const Scene: React.FC<Props> = ({
  id,
  title,
  description,
  camera,
  index,
  sceneView,
  sceneRecapture,
  sceneDelete,
  sceneEdit,
  sceneMove,
}) => {
  const hendleView = useCallback(() => {
    if (camera) {
      sceneView(camera);
    }
  }, [sceneView, camera]);

  const handleEdit = useCallback(() => {
    sceneEdit(id);
  }, [sceneEdit, id]);

  const handleRecapture = useCallback(() => {
    sceneRecapture(id);
  }, [sceneRecapture, id]);

  const handleDelete = useCallback(() => {
    sceneDelete(id);
  }, [sceneDelete, id]);

  const items = [
    { label: "ビュー", key: "view", onClick: hendleView },
    { label: "編集", key: "edit", onClick: handleEdit },
    { label: "再キャプチャ", key: "recapture", onClick: handleRecapture },
    { label: "削除", key: "delete", onClick: handleDelete },
  ];
  const menu = <Menu items={items} />;

  interface DragItem {
    index: number;
    id: string;
    type: string;
  }

  const ref = useRef<HTMLDivElement>(null);

  const [{ handlerId }, drop] = useDrop<DragItem, void, { handlerId: Identifier | null }>({
    accept: "scene",
    collect(monitor) {
      return {
        handlerId: monitor.getHandlerId(),
      };
    },
    hover(item: DragItem, monitor) {
      if (!ref.current) {
        return;
      }
      const dragIndex = item.index;
      const hoverIndex = index;

      if (dragIndex === hoverIndex) {
        return;
      }

      const hoverBoundingRect = ref.current?.getBoundingClientRect();
      const hoverMiddleX = (hoverBoundingRect.right - hoverBoundingRect.left) / 2;
      const clientOffset = monitor.getClientOffset();
      const hoverClientX = (clientOffset as XYCoord).x - hoverBoundingRect.left;

      if (dragIndex < hoverIndex && hoverClientX < hoverMiddleX) {
        return;
      }

      if (dragIndex > hoverIndex && hoverClientX > hoverMiddleX) {
        return;
      }

      sceneMove(dragIndex, hoverIndex);

      item.index = hoverIndex;
    },
  });

  const [{ isDragging }, drag] = useDrag({
    type: "scene",
    item: () => {
      return { id, index };
    },
    collect: (monitor: any) => ({
      isDragging: monitor.isDragging(),
    }),
  });

  const opacity = isDragging ? 0 : 1;
  drag(drop(ref));

  return (
    <StyledScene ref={ref} style={{ opacity }} data-handler-id={handlerId}>
      <Header>
        <Title>{title}</Title>
        <ActionsBtn>
          <Dropdown
            trigger={["click"]}
            overlay={menu}
            placement={"bottomRight"}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}>
            <Icon icon="dotsThreeVertical" size={24} color={"var(--theme-color)"} />
          </Dropdown>
        </ActionsBtn>
      </Header>
      <Description>{description}</Description>
    </StyledScene>
  );
};

const StyledScene = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  width: 170px;
  height: 114px;
  background: #f8f8f8;
  border-radius: 8px;
  border: 1px solid #c7c5c5;
  padding: 6px 0;
`;

const Header = styled.div`
  display: flex;
  width: 100%;
  height: 36px;
  padding: 6px 12px;
  align-items: center;
  justify-content: space-between;
`;

const Title = styled.div`
  width: 100%;
  font-weight: 700;
  font-size: 14px;
  line-height: 19px;
  color: #000;
  text-overflow: ellipsis;
  white-space: nowrap;
  overflow: hidden;
`;

const ActionsBtn = styled.a`
  position: relative;
  display: flex;
  flex-shrink: 0;

  .ant-dropdown-menu-item,
  .ant-dropdown-menu-submenu-title {
    line-height: 18px;
  }
`;

const Description = styled.div`
  height: 100%;
  padding: 6px 12px;
  color: #3a3a3a;
  font-size: 12px;
  line-height: 18px;
  font-weight: 500;
  overflow: hidden;
  display: -webkit-box !important;
  -webkit-line-clamp: 3;
  -webkit-box-orient: vertical;
  white-space: normal;
`;

export default Scene;
