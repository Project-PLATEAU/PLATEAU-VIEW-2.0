import { Comment } from "@reearth-cms/components/molecules/Asset/asset.type";
import CommentsPanelMolecule from "@reearth-cms/components/molecules/Common/CommentsPanel";

import useHooks from "./hooks";

export type Props = {
  emptyText?: string;
  threadId?: string;
  comments?: Comment[];
  collapsed: boolean;
  onCollapse: (value: boolean) => void;
};

const CommentsPanel: React.FC<Props> = ({
  emptyText,
  threadId,
  comments,
  collapsed,
  onCollapse,
}) => {
  const { me, handleCommentCreate, handleCommentUpdate, handleCommentDelete } = useHooks({
    threadId,
  });

  return (
    <CommentsPanelMolecule
      me={me}
      emptyText={emptyText}
      comments={comments}
      collapsed={collapsed}
      onCollapse={onCollapse}
      onCommentCreate={handleCommentCreate}
      onCommentUpdate={handleCommentUpdate}
      onCommentDelete={handleCommentDelete}
    />
  );
};

export default CommentsPanel;
