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

import { BaseFieldProps, Cond, PolygonColor as PolygonColorType } from "../../types";

import PolygonColorItem from "./PolygonColorItem";

const PolygonColor: React.FC<BaseFieldProps<"polygonColor">> = ({ value, editMode, onUpdate }) => {
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
      const newItem: { condition: Cond<any>; color: string } = {
        condition: {
          key: generateID(),
          operator: "===",
          operand: true,
          value: true,
        },
        color: "",
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

  const handleItemUpdate = (item: { condition: Cond<number>; color: string }, index: number) => {
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
        <PolygonColorItem
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

export default PolygonColor;

const generateOverride = (items: PolygonColorType["items"]) => {
  const fillColorConditions: [string, string][] = [["true", 'color("white")']];
  const fillConditions: [string, string][] = [["true", "true"]];
  items?.forEach(item => {
    const resFillColor = "color" + `("${item.color}")`;
    const cond = stringifyCondition(item.condition);
    fillColorConditions.unshift([cond, resFillColor]);
    fillConditions.unshift([cond, cond]);
  });
  return {
    polygon: {
      fill: {
        expression: { conditions: fillConditions },
      },
      fillColor: { expression: { conditions: fillColorConditions } },
    },
  };
};
