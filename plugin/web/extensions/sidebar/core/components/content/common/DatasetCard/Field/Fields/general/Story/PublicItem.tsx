import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback } from "react";

import { StoryItem } from "../../types";

type Props = {
  story: StoryItem;
  handleStoryPlay: (story: StoryItem) => void;
};

const PublicStoryItem: React.FC<Props> = ({ story, handleStoryPlay }) => {
  const onPlay = useCallback(() => {
    handleStoryPlay(story);
  }, [handleStoryPlay, story]);

  return (
    <StoryButton key={story.id} onClick={onPlay}>
      <StyledIcon icon="circledPlay" size={24} />
      <Text>{story.title}</Text>
    </StoryButton>
  );
};

const Text = styled.p`
  margin: 0;
  line-height: 24px;
  font-weight: 400;
  font-size: 14px;
`;

const StyledIcon = styled(Icon)`
  color: #00bebe;
`;

const StoryButton = styled.div`
  display: flex;
  align-items: center;
  background: #ffffff;
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.15);
  border-radius: 4px;
  padding: 12px;
  gap: 12px;
  height: 48px;
  cursor: pointer;
  :hover {
    background: #f4f4f4;
  }
`;

export default PublicStoryItem;
