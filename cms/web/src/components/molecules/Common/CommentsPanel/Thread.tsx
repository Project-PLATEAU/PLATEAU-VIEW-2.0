import { User } from "@reearth-cms/components/molecules/AccountSettings/types";
import { Comment } from "@reearth-cms/components/molecules/Asset/asset.type";

import { CommentList } from "./CommentList";

type Props = {
  me?: User;
  comments?: Comment[];
  onCommentUpdate: (commentId: string, content: string) => Promise<void>;
  onCommentDelete: (commentId: string) => Promise<void>;
};

export const Thread: React.FC<Props> = ({ me, comments, onCommentUpdate, onCommentDelete }) => {
  return (
    <>
      {comments && comments?.length > 0 && (
        <CommentList
          me={me}
          comments={comments}
          onCommentUpdate={onCommentUpdate}
          onCommentDelete={onCommentDelete}
        />
      )}
    </>
  );
};

export default Thread;
