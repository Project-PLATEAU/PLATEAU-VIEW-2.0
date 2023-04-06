import AddButton from "@web/extensions/sidebar/core/components/content/common/DatasetCard/AddButton";
import { styled } from "@web/theme";

import { BaseFieldProps } from "../../types";

import EditStoryItem from "./EditStoryItem";
import useHooks from "./hooks";
import PublicStoryItem from "./PublicItem";

const Story: React.FC<BaseFieldProps<"story">> = ({ value, editMode, dataID, onUpdate }) => {
  const {
    stories,
    handleStoryAdd,
    handleItemMoveUp,
    handleItemMoveDown,
    handleItemRemove,
    handleStoryTitleChange,
    handleStoryEdit,
    handleStoryEditFinish,
    handleStoryPlay,
  } = useHooks({ value, dataID, onUpdate });

  return editMode ? (
    <Wrapper>
      <AddButton text="新しいストーリーを追加" onClick={handleStoryAdd} />
      {stories?.map((story, idx) => (
        <EditStoryItem
          story={story}
          key={story.id}
          idx={idx}
          handleItemMoveUp={handleItemMoveUp}
          handleItemMoveDown={handleItemMoveDown}
          handleItemRemove={handleItemRemove}
          handleStoryTitleChange={handleStoryTitleChange}
          handleStoryEdit={handleStoryEdit}
          handleStoryEditFinish={handleStoryEditFinish}
        />
      ))}
    </Wrapper>
  ) : (
    <Wrapper>
      {stories?.map(story => (
        <PublicStoryItem story={story} key={story.id} handleStoryPlay={handleStoryPlay} />
      ))}
    </Wrapper>
  );
};

export default Story;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;
