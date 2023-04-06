import {
  ColorField,
  TextField,
  ItemControls,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Item } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { useCallback, ChangeEvent } from "react";

import { LegendItem, LegendStyleType } from "../../types";

const LegendItemComponent: React.FC<{
  index: number;
  item: LegendItem;
  legendStyle: LegendStyleType;
  handleMoveDown: (index: number) => void;
  handleMoveUp: (index: number) => void;
  handleRemove: (index: number) => void;
  onItemUpdate: (property: LegendItem, index: number) => void;
}> = ({ index, item, legendStyle, handleMoveDown, handleMoveUp, handleRemove, onItemUpdate }) => {
  const handleColorChange = useCallback(
    (color?: string) => {
      const copy = { ...item, color: color };
      onItemUpdate(copy, index);
    },
    [index, item, onItemUpdate],
  );

  const handleTitleChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      const copy = { ...item, title: e.currentTarget.value };
      onItemUpdate(copy, index);
    },
    [index, item, onItemUpdate],
  );

  const handleURLChange = useCallback(
    (e: ChangeEvent<HTMLInputElement>) => {
      const copy = { ...item, url: e.currentTarget.value };
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
      {legendStyle === "icon" && (
        <TextField title="URL" titleWidth={82} value={item.url} onChange={handleURLChange} />
      )}

      <ColorField title="色" titleWidth={82} color={item.color} onChange={handleColorChange} />
      <TextField title="タイトル" titleWidth={82} value={item.title} onChange={handleTitleChange} />
    </Item>
  );
};

export default LegendItemComponent;
