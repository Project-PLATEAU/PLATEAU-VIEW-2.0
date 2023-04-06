import { generateID, moveItemDown, moveItemUp, postMsg } from "@web/extensions/sidebar/utils";
import { useCallback, useState } from "react";

import { BaseFieldProps, StoryItem } from "../../types";

export default ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"story">, "value" | "dataID" | "onUpdate">) => {
  const [stories, updateStories] = useState(value.stories);

  const handleStoryAdd = useCallback(() => {
    updateStories(s => {
      const newItem = {
        id: generateID(),
        dataID,
        title: "タイトル",
      };
      const newStories = s ? [...s, newItem] : [newItem];
      onUpdate({ ...value, stories: newStories });
      return newStories;
    });
  }, [onUpdate, value, dataID]);

  const handleItemMoveUp = useCallback(
    (idx: number) => {
      updateStories(s => {
        if (!s) return s;
        const newStories = moveItemUp(idx, s) ?? s;
        onUpdate({ ...value, stories: newStories });
        return newStories;
      });
    },
    [onUpdate, value],
  );

  const handleItemMoveDown = useCallback(
    (idx: number) => {
      updateStories(s => {
        if (!s) return s;
        const newStories = moveItemDown(idx, s) ?? s;
        onUpdate({ ...value, stories: newStories });
        return newStories;
      });
    },
    [onUpdate, value],
  );

  const handleItemRemove = useCallback(
    (id: string) => {
      updateStories(s => {
        const newStories = s?.filter(st => st.id !== id);
        onUpdate({ ...value, stories: newStories });
        return newStories;
      });
      postMsg({
        action: "storyDelete",
        payload: {
          id,
        },
      });
    },
    [onUpdate, value],
  );

  const handleStoryTitleChange = useCallback(
    (id: string, title: string) => {
      updateStories(s => {
        if (!s) return s;
        const story = s.find(story => story.id === id);
        if (story) story.title = title;
        const newStories = [...s];
        onUpdate({ ...value, stories: newStories });
        return newStories;
      });
    },
    [onUpdate, value],
  );

  const handleStoryEdit = useCallback(
    (story: StoryItem) => {
      postMsg({
        action: "storyEdit",
        payload: { ...story, dataID },
      });
    },
    [dataID],
  );

  const handleStoryEditFinish = useCallback((id: string) => {
    postMsg({
      action: "storyEditFinish",
      payload: { id },
    });
  }, []);

  const handleStoryPlay = useCallback((story: StoryItem) => {
    postMsg({
      action: "storyPlay",
      payload: story,
    });
  }, []);

  return {
    stories,
    handleStoryAdd,
    handleItemMoveUp,
    handleItemMoveDown,
    handleItemRemove,
    handleStoryTitleChange,
    handleStoryEdit,
    handleStoryEditFinish,
    handleStoryPlay,
  };
};
