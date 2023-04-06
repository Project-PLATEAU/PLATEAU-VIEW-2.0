import AddButton from "@web/extensions/sidebar/core/components/content/common/DatasetCard/AddButton";
import { array_move } from "@web/extensions/sidebar/utils";
import { Icon, Dropdown, Menu } from "@web/sharedComponents";
import ModifiedImage from "@web/sharedComponents/ModifiedImage";
import { styled } from "@web/theme";
import { useCallback, useState } from "react";

import { BaseFieldProps, LegendItem, LegendStyleType } from "../../types";

import LegendItemComponent from "./LegendItem";

const legendStyles: { [key: string]: string } = {
  square: "四角",
  circle: "丸",
  line: "線",
  icon: "アイコン",
};

const Legend: React.FC<BaseFieldProps<"legend">> = ({ value, editMode, onUpdate }) => {
  const [legend, updateLegend] = useState(value);

  const handleStyleChange = useCallback(
    (style: LegendStyleType) => {
      updateLegend(l => {
        const newLegend = {
          ...l,
          style,
        };
        onUpdate(newLegend);
        return newLegend;
      });
    },
    [onUpdate],
  );

  const handleMoveUp = useCallback(
    (idx: number) => {
      if (idx === 0) return;
      updateLegend(l => {
        let newItems: LegendItem[] | undefined = undefined;
        if (l.items) {
          newItems = l.items;
          array_move(newItems, idx, idx - 1);
        }
        const newLegend = { ...l, items: newItems };
        onUpdate(newLegend);
        return newLegend;
      });
    },
    [onUpdate],
  );

  const handleMoveDown = useCallback(
    (idx: number) => {
      if (legend.items && idx >= legend.items.length - 1) return;
      updateLegend(l => {
        let newItems: LegendItem[] | undefined = undefined;
        if (l.items) {
          newItems = l.items;
          array_move(newItems, idx, idx + 1);
        }
        const newLegend = { ...l, items: newItems };
        onUpdate(newLegend);
        return newLegend;
      });
    },
    [legend, onUpdate],
  );

  const handleAdd = useCallback(() => {
    updateLegend(l => {
      const newItem = { title: "新しいアイテム", color: "#00bebe" };
      const newLegend = {
        ...l,
        items: l.items ? [...l.items, newItem] : [newItem],
      };
      onUpdate(newLegend);
      return newLegend;
    });
  }, [onUpdate]);

  const handleRemove = useCallback(
    (idx: number) => {
      updateLegend(l => {
        let newItems: LegendItem[] | undefined = undefined;
        if (l.items) {
          newItems = l.items.filter((_, idx2) => idx2 != idx);
        }
        const newLegend = { ...l, items: newItems };
        onUpdate(newLegend);
        return newLegend;
      });
    },
    [onUpdate],
  );

  const handleItemUpdate = (item: LegendItem, index: number) => {
    updateLegend(l => {
      const newLegendsItems = [...(l.items ?? [])];
      newLegendsItems.splice(index, 1, item);
      const newLegend = {
        ...value,
        items: newLegendsItems,
      };
      onUpdate(newLegend);
      return newLegend;
    });
  };

  const menu = (
    <Menu
      items={Object.keys(legendStyles).map(ls => {
        return {
          key: ls,
          label: (
            <p style={{ margin: 0 }} onClick={() => handleStyleChange(ls as LegendStyleType)}>
              {legendStyles[ls]}
            </p>
          ),
        };
      })}
    />
  );

  return editMode ? (
    <Wrapper>
      <Field>
        <FieldTitle>スタイル</FieldTitle>
        <FieldValue>
          <Dropdown
            overlay={menu}
            placement="bottom"
            trigger={["click"]}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}>
            <StyledDropdownButton>
              <p style={{ margin: 0 }}>{legendStyles[legend.style]}</p>
              <StyledIcon icon="arrowDownSimple" size={12} />
            </StyledDropdownButton>
          </Dropdown>
        </FieldValue>
      </Field>
      <AddButton text="項目" onClick={handleAdd} />
      {legend.items?.map((item, idx) => (
        <LegendItemComponent
          key={idx}
          index={idx}
          item={item}
          legendStyle={legend.style}
          handleMoveDown={handleMoveDown}
          handleMoveUp={handleMoveUp}
          handleRemove={handleRemove}
          onItemUpdate={handleItemUpdate}
        />
      ))}
    </Wrapper>
  ) : (
    <Wrapper>
      {legend.items?.map((item, idx) => (
        <Field key={idx} gap={12}>
          {legend.style === "icon" ? (
            <StyledImgWrapper>
              {item.url && (
                <ModifiedImage
                  imageUrl={item.url}
                  blendColor={item.color ?? " #d9d9d9"}
                  width={30}
                  height={30}
                />
              )}
            </StyledImgWrapper>
          ) : (
            <ColorBlock color={item.color} legendStyle={legend.style} />
          )}
          <Text>{item.title}</Text>
        </Field>
      ))}
    </Wrapper>
  );
};

export default Legend;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const StyledDropdownButton = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  align-content: center;
  padding: 0 16px;
  cursor: pointer;
`;

const StyledIcon = styled(Icon)`
  font-size: 0;
`;

const Text = styled.p`
  margin: 0;
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  min-height: 32px;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  position: relative;
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const ColorBlock = styled.div<{ color?: string; legendStyle?: "circle" | "square" | "line" }>`
  flex: none;
  width: 30px;
  height: ${({ legendStyle }) => (legendStyle === "line" ? "3px" : "30px")};
  background: ${({ color }) => color ?? "#d9d9d9"};
  border-radius: ${({ legendStyle }) =>
    legendStyle
      ? legendStyle === "circle"
        ? "50%"
        : legendStyle === "line"
        ? "5px"
        : "2px"
      : "1px 0 0 1px"};
`;

const StyledImgWrapper = styled.div`
  flex: none;
  width: 30px;
  position: relative;
`;
