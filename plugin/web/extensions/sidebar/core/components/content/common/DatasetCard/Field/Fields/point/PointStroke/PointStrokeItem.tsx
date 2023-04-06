import {
  ColorField,
  ConditionField,
  ItemControls,
  NumberField,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Item } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { useCallback } from "react";

import { Cond } from "../../types";

const PointStrokeItem: React.FC<{
  index: number;
  item: { condition: Cond<string | number>; strokeColor: string; strokeWidth: number };
  handleMoveDown: (index: number) => void;
  handleMoveUp: (index: number) => void;
  handleRemove: (index: number) => void;
  onItemUpdate: (
    item: { condition: Cond<string | number>; strokeColor: string; strokeWidth: number },
    index: number,
  ) => void;
}> = ({ index, item, handleMoveDown, handleMoveUp, handleRemove, onItemUpdate }) => {
  const handleStrokeColorUpdate = useCallback(
    (color: string) => {
      if (color) {
        const copy = { ...item, strokeColor: color };
        onItemUpdate(copy, index);
      }
    },
    [index, item, onItemUpdate],
  );

  const handleConditionUpdate = useCallback(
    (condition: Cond<number>) => {
      if (condition) {
        const copy = { ...item, condition };
        onItemUpdate(copy, index);
      }
    },
    [index, item, onItemUpdate],
  );

  const handleStrokeWidthUpdate = useCallback(
    (strokeWidth: number) => {
      const copy = { ...item, strokeWidth };
      onItemUpdate(copy, index);
    },
    [index, item, onItemUpdate],
  );

  return (
    <Item>
      <ItemControls
        index={index}
        handleMoveDown={handleMoveDown}
        handleMoveUp={handleMoveUp}
        handleRemove={handleRemove}
      />
      <ConditionField
        title="if"
        fieldGap={8}
        condition={item.condition}
        onChange={handleConditionUpdate}
      />
      <ColorField
        title="strokeColor"
        titleWidth={82}
        color={item.strokeColor}
        onChange={handleStrokeColorUpdate}
      />
      <NumberField
        title="StrokeWidth"
        titleWidth={82}
        value={item.strokeWidth}
        onChange={handleStrokeWidthUpdate}
      />
    </Item>
  );
};

export default PointStrokeItem;
