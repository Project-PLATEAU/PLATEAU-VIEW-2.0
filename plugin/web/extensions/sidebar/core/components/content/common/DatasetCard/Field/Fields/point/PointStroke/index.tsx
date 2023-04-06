import AddButton from "@web/extensions/sidebar/core/components/content/common/DatasetCard/AddButton";
import {
  ButtonWrapper,
  Wrapper,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import {
  generateID,
  moveItemDown,
  moveItemUp,
  removeItem,
  stringifyCondition,
} from "@web/extensions/sidebar/utils";
import { styled, commonStyles } from "@web/theme";
import { useCallback, useState } from "react";

import { BaseFieldProps, Cond, PointStroke as PointStrokeType } from "../../types";

import PointStrokeItem from "./PointStrokeItem";

const PointStroke: React.FC<BaseFieldProps<"pointStroke">> = ({ value, editMode, onUpdate }) => {
  const [items, updateItems] = useState(value.items);

  const handleMoveUp = useCallback((idx: number) => {
    if (idx === 0) return;
    updateItems(c => {
      const newItems = moveItemUp(idx, c) ?? c;
      return newItems;
    });
  }, []);

  const handleMoveDown = useCallback(
    (idx: number) => {
      if (items && idx >= items.length - 1) return;
      updateItems(c => {
        const newItems = moveItemDown(idx, c) ?? c;
        return newItems;
      });
    },
    [items],
  );

  const handleAdd = useCallback(() => {
    updateItems(c => {
      const newItem: {
        strokeColor: string;
        strokeWidth: number;
        condition: Cond<any>;
      } = {
        strokeColor: "",
        strokeWidth: 0,
        condition: {
          key: generateID(),
          operator: "===",
          operand: true,
          value: true,
        },
      };
      return c ? [...c, newItem] : [newItem];
    });
  }, []);

  const handleRemove = useCallback((idx: number) => {
    updateItems(c => {
      const newItems = removeItem(idx, c) ?? c;
      return newItems;
    });
  }, []);

  const handleItemUpdate = (
    item: { condition: Cond<string | number>; strokeColor: string; strokeWidth: number },
    index: number,
  ) => {
    updateItems(c => {
      const newItems = [...(c ?? [])];
      newItems.splice(index, 1, item);
      return newItems;
    });
  };

  const handleApply = useCallback(() => {
    onUpdate({
      ...value,
      items,
      override: generateOverride(items),
    });
  }, [value, items, onUpdate]);

  return editMode ? (
    <Wrapper>
      {items?.map((c, idx) => (
        <PointStrokeItem
          key={idx}
          index={idx}
          item={c}
          handleMoveDown={handleMoveDown}
          handleMoveUp={handleMoveUp}
          handleRemove={handleRemove}
          onItemUpdate={handleItemUpdate}
        />
      ))}
      <ButtonWrapper>
        <AddButton text="Add Condition" height={24} onClick={handleAdd} />
      </ButtonWrapper>
      <Button onClick={handleApply}>Apply</Button>
    </Wrapper>
  ) : null;
};

const Button = styled.div`
  ${commonStyles.simpleButton}
`;

export default PointStroke;

const generateOverride = (items: PointStrokeType["items"]) => {
  const pointOutlineColorConditions: [string, string][] = [["true", 'color("white")']];
  const pointOutlineWidthConditions: [string, string][] = [["true", "1"]];
  items?.forEach(item => {
    const resStrokeColor = "color" + `("${item.strokeColor}")`;
    const resStrokeWidth = String(item.strokeWidth);
    const cond = stringifyCondition(item.condition);
    pointOutlineColorConditions.unshift([cond, resStrokeColor]);
    pointOutlineWidthConditions.unshift([cond, resStrokeWidth]);
  });
  return {
    marker: {
      style: "point",
      pointOutlineColor: {
        expression: {
          conditions: pointOutlineColorConditions,
        },
      },
      pointOutlineWidth: {
        expression: {
          conditions: pointOutlineWidthConditions,
        },
      },
    },
  };
};
