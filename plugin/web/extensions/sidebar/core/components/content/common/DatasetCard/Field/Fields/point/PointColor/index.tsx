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

import { BaseFieldProps, Cond, PointColor as PointColorType } from "../../types";

import PointColorItem from "./PointColorItem";

const PointColor: React.FC<BaseFieldProps<"pointColor">> = ({ value, editMode, onUpdate }) => {
  const [pointColors, updatePointColors] = useState(value.pointColors);

  const handleMoveUp = useCallback((idx: number) => {
    updatePointColors(c => {
      const newPointColors = moveItemUp(idx, c) ?? c;
      return newPointColors;
    });
  }, []);

  const handleMoveDown = useCallback((idx: number) => {
    updatePointColors(c => {
      const newPointColors = moveItemDown(idx, c) ?? c;
      return newPointColors;
    });
  }, []);

  const handleAdd = useCallback(() => {
    updatePointColors(c => {
      const newPointColor: { condition: Cond<any>; color: string } = {
        condition: {
          key: generateID(),
          operator: "===",
          operand: true,
          value: true,
        },
        color: "",
      };
      return c ? [...c, newPointColor] : [newPointColor];
    });
  }, []);

  const handleRemove = useCallback((idx: number) => {
    updatePointColors(c => {
      const newPointColors = removeItem(idx, c) ?? c;
      return newPointColors;
    });
  }, []);

  const handleItemUpdate = (item: { condition: Cond<number>; color: string }, index: number) => {
    updatePointColors(c => {
      const newPointColors = [...(c ?? [])];
      newPointColors.splice(index, 1, item);
      return newPointColors;
    });
  };

  const handleApply = useCallback(() => {
    onUpdate({
      ...value,
      pointColors,
      override: generateOverride(pointColors),
    });
  }, [value, pointColors, onUpdate]);

  return editMode ? (
    <Wrapper>
      {pointColors?.map((c, idx) => (
        <PointColorItem
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

export default PointColor;

const generateOverride = (pointColors: PointColorType["pointColors"]) => {
  const conditions: [string, string][] = [["true", 'color("white")']];
  pointColors?.forEach(item => {
    const res = "color" + `("${item.color}")`;
    const cond = stringifyCondition(item.condition);
    conditions.unshift([cond, res]);
  });
  return {
    marker: {
      style: "point",
      pointColor: {
        expression: {
          conditions,
        },
      },
    },
  };
};
