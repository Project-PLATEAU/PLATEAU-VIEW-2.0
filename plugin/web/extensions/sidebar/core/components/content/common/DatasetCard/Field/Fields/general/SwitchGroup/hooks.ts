import { array_move, generateID } from "@web/extensions/sidebar/utils";
import { fieldGroups } from "@web/extensions/sidebar/utils/fieldGroups";
import { useCallback, useMemo, useState } from "react";

import { GroupItem, SwitchGroup } from "../../types";

export default ({
  value,
  selectedGroup,
  onUpdate,
  onCurrentGroupUpdate,
}: {
  value: SwitchGroup;
  selectedGroup?: string;
  onUpdate: (property: SwitchGroup) => void;
  onCurrentGroupUpdate?: (fieldGroupID: string) => void;
}) => {
  const [groupItems, updateGroupItems] = useState<GroupItem[]>(value.groups);
  const [title, setTitle] = useState(value.title);

  const currentGroup = useMemo(
    () => groupItems.find(fg => fg.fieldGroupID === selectedGroup),
    [groupItems, selectedGroup],
  );

  const handleTitleChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setTitle(e.target.value);
      onUpdate({ ...value, title: e.target.value });
    },
    [value, onUpdate],
  );

  const handleGroupChoose = useCallback(
    (id: string) => {
      const selected = groupItems?.find(gi => gi.id === id);
      if (!selected) return;
      onCurrentGroupUpdate?.(selected.fieldGroupID);
    },
    [groupItems, onCurrentGroupUpdate],
  );

  const handleItemAdd = useCallback(() => {
    const newItem: GroupItem = {
      id: generateID(),
      title: `新グループ${value.groups.length ? value.groups.length + 1 : 1}`,
      fieldGroupID: fieldGroups[0].id,
    };
    updateGroupItems(gi => (gi ? [...gi, newItem] : [newItem]));
    onUpdate({ ...value, groups: value.groups ? [...value.groups, newItem] : [newItem] });
  }, [value, onUpdate]);

  const handleItemRemove = useCallback(
    (id: string) => {
      if (value.groups.length < 2) return;
      const newGroups = value.groups?.filter(g => g.id !== id);
      if (!newGroups) return;
      onUpdate({ ...value, groups: newGroups });
    },
    [value, onUpdate],
  );

  const handleItemGroupChange = useCallback(
    (idx: number, fieldGroupID?: string) => {
      if (!fieldGroupID || !value.groups) return;
      const updatedGroups = value.groups;
      updatedGroups[idx].fieldGroupID = fieldGroupID;
      updateGroupItems(updatedGroups);
      onUpdate({ ...value, groups: updatedGroups });
    },
    [value, onUpdate],
  );

  const handleItemTitleChange = useCallback(
    (title: string, idx: number) => {
      if (!value.groups) return;
      const updatedGroups = value.groups;
      updatedGroups[idx].title = title;
      onUpdate({ ...value, groups: updatedGroups });
    },
    [value, onUpdate],
  );

  const handleItemMoveUp = useCallback(
    (idx: number) => {
      if (idx === 0 || !value.groups) return;
      const newGroups = value.groups;
      array_move(newGroups, idx, idx - 1);
      onUpdate({ ...value, groups: newGroups });
    },
    [value, onUpdate],
  );

  const handleItemMoveDown = useCallback(
    (idx: number) => {
      if (!value.groups || idx >= value.groups.length - 1) return;
      const newGroups = value.groups;
      array_move(newGroups, idx, idx + 1);
      onUpdate({ ...value, groups: newGroups });
    },
    [value, onUpdate],
  );

  return {
    title,
    groupItems,
    currentGroup,
    fieldGroups,
    handleTitleChange,
    handleGroupChoose,
    handleItemGroupChange,
    handleItemTitleChange,
    handleItemAdd,
    handleItemRemove,
    handleItemMoveUp,
    handleItemMoveDown,
  };
};
