import {
  TextField,
  ConditionField,
  ItemControls,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Item } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { useCallback, ChangeEvent } from "react";

import { Cond } from "../../types";

import { ConditionItemType } from ".";

const ConditionItem: React.FC<{
  index: number;
  item: ConditionItemType;
  handleMoveDown: (index: number) => void;
  handleMoveUp: (index: number) => void;
  handleRemove: (index: number) => void;
  onItemUpdate: (item: ConditionItemType, index: number) => void;
}> = ({ index, item, handleMoveDown, handleMoveUp, handleRemove, onItemUpdate }) => {
  const handleTitleChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      onItemUpdate({ ...item, title: e.target.value }, index);
    },
    [index, item, onItemUpdate],
  );

  const handleConditionUpdate = useCallback(
    (condition: Cond<any>) => {
      if (condition) {
        onItemUpdate({ ...item, condition }, index);
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
      <TextField title="Title" titleWidth={82} value={item.title} onChange={handleTitleChange} />
    </Item>
  );
};

export default ConditionItem;
