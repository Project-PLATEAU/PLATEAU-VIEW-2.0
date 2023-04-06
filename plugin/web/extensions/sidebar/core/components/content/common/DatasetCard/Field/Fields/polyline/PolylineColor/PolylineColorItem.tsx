import {
  ColorField,
  ConditionField,
  ItemControls,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Item } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { useCallback } from "react";

import { Cond } from "../../types";

const PolylineColorItem: React.FC<{
  index: number;
  item: { condition: Cond<number>; color: string };
  handleMoveDown: (index: number) => void;
  handleMoveUp: (index: number) => void;
  handleRemove: (index: number) => void;
  onItemUpdate: (item: { condition: Cond<number>; color: string }, index: number) => void;
}> = ({ index, item, handleMoveDown, handleMoveUp, handleRemove, onItemUpdate }) => {
  const handleBackgroundColorUpdate = useCallback(
    (color: string) => {
      if (color) {
        const copy = { ...item, color };
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
        title="色"
        titleWidth={82}
        color={item.color}
        onChange={handleBackgroundColorUpdate}
      />
    </Item>
  );
};

export default PolylineColorItem;
