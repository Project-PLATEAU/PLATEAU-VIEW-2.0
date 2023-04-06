import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useEffect } from "react";

import type { StoryItem } from "../../types";

type Props = {
  story: StoryItem;
  idx: number;
  handleItemMoveUp: (idx: number) => void;
  handleItemMoveDown: (idx: number) => void;
  handleItemRemove: (id: string) => void;
  handleStoryTitleChange: (id: string, title: string) => void;
  handleStoryEdit: (story: StoryItem) => void;
  handleStoryEditFinish: (id: string) => void;
};

const EditStoryItem: React.FC<Props> = ({
  story,
  idx,
  handleItemMoveUp,
  handleItemMoveDown,
  handleItemRemove,
  handleStoryTitleChange,
  handleStoryEdit,
  handleStoryEditFinish,
}: Props) => {
  const onMoveUp = useCallback(() => {
    handleItemMoveUp(idx);
  }, [handleItemMoveUp, idx]);

  const onMoveDown = useCallback(() => {
    handleItemMoveDown(idx);
  }, [handleItemMoveDown, idx]);

  const onRemove = useCallback(() => {
    handleItemRemove(story.id);
  }, [handleItemRemove, story]);

  const onTitleChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      handleStoryTitleChange(story.id, e.currentTarget.value);
    },
    [handleStoryTitleChange, story],
  );

  const onEdit = useCallback(() => {
    handleStoryEdit(story);
  }, [handleStoryEdit, story]);

  const onFinish = useCallback(() => {
    handleStoryEditFinish(story.id);
  }, [handleStoryEditFinish, story]);

  useEffect(() => {
    return () => {
      onFinish();
    };
  }, [onFinish]);

  return (
    <Item>
      <ItemControls>
        <Icon icon="arrowUpThin" size={16} onClick={onMoveUp} />
        <Icon icon="arrowDownThin" size={16} onClick={onMoveDown} />
        <TrashIcon icon="trash" size={16} onClick={onRemove} />
      </ItemControls>
      <Field>
        <FieldTitle>タイトル</FieldTitle>
        <FieldValue>
          <TextInput value={story.title} onChange={onTitleChange} />
        </FieldValue>
      </Field>
      <EditButton onClick={onEdit}>
        <StyledIcon icon="editUnderline" size={14} />
        <Text>編集</Text>
      </EditButton>
      <EditButton onClick={onFinish}>
        <StyledIcon icon="editStop" size={14} />
        <Text>編集を終了する</Text>
      </EditButton>
    </Item>
  );
};

const Item = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 8px;
`;

const ItemControls = styled.div`
  display: flex;
  justify-content: right;
  gap: 4px;
  cursor: pointer;
`;

const TrashIcon = styled(Icon)<{ disabled?: boolean }>`
  ${({ disabled }) =>
    disabled &&
    `
      color: rgb(209, 209, 209);
      pointer-events: none;
    `}
`;

const EditButton = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  background: #ffffff;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 5px;
  height: 32px;
  cursor: pointer;

  :hover {
    background: #f4f4f4;
  }
`;

const StyledIcon = styled(Icon)`
  font-size: 0;
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

const Text = styled.p`
  margin: 0;
  line-height: 24px;
  font-weight: 400;
  font-size: 14px;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const TextInput = styled.input.attrs({ type: "text" })`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;

  :focus {
    border: none;
  }
`;

export default EditStoryItem;
