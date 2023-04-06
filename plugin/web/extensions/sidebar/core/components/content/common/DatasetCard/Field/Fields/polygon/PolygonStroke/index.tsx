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

import { BaseFieldProps, Cond, PolygonStroke as PolygonStrokeType } from "../../types";

import PolygonStrokeItem from "./PolygonStrokeItem";

const PolygonStroke: React.FC<BaseFieldProps<"polygonStroke">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  const [items, updateItems] = useState(value.items);

  const handleMoveUp = useCallback((idx: number) => {
    updateItems(c => {
      const newItems = moveItemUp(idx, c) ?? c;
      return newItems;
    });
  }, []);

  const handleMoveDown = useCallback((idx: number) => {
    updateItems(c => {
      const newItems = moveItemDown(idx, c) ?? c;
      return newItems;
    });
  }, []);

  const handleAdd = useCallback(() => {
    updateItems(c => {
      const newItem: {
        strokeColor: string;
        strokeWidth: number;
        condition: Cond<any>;
      } = {
        strokeColor: "",
        strokeWidth: 5,
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
        <PolygonStrokeItem
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

export default PolygonStroke;

const generateOverride = (items: PolygonStrokeType["items"]) => {
  const strokeConditions: [string, string][] = [["true", "true"]];
  const strokeColorConditions: [string, string][] = [["true", 'color("white")']];
  const strokeWidthConditions: [string, string][] = [["true", "1"]];
  items?.forEach(item => {
    const resStrokeColor = "color" + `("${item.strokeColor}")`;
    const resStrokeWidth = String(item.strokeWidth);
    const cond = stringifyCondition(item.condition);
    strokeColorConditions.unshift([cond, resStrokeColor]);
    strokeWidthConditions.unshift([cond, resStrokeWidth]);
    strokeConditions.unshift([cond, cond]);
  });
  return {
    polygon: {
      stroke: {
        expression: {
          conditions: strokeConditions,
        },
      },
      strokeColor: {
        expression: {
          conditions: strokeColorConditions,
        },
      },
      strokeWidth: {
        expression: {
          conditions: strokeWidthConditions,
        },
      },
    },
  };
};
