import styled from "@emotion/styled";

import Icon from "@reearth-cms/components/atoms/Icon";
import { User } from "@reearth-cms/components/molecules/AccountSettings/types";
import { Comment } from "@reearth-cms/components/molecules/Asset/asset.type";
import Sidebar from "@reearth-cms/components/molecules/Common/Sidebar";
import { useT } from "@reearth-cms/i18n";

import Editor from "./Editor";
import Thread from "./Thread";

export type Props = {
  me?: User;
  comments?: Comment[];
  emptyText?: string;
  collapsed: boolean;
  onCollapse: (value: boolean) => void;
  onCommentCreate: (content: string) => Promise<void>;
  onCommentUpdate: (commentId: string, content: string) => Promise<void>;
  onCommentDelete: (commentId: string) => Promise<void>;
};

const CommentsPanel: React.FC<Props> = ({
  me,
  comments,
  emptyText,
  collapsed,
  onCollapse,
  onCommentCreate,
  onCommentUpdate,
  onCommentDelete,
}) => {
  const t = useT();

  return (
    <Sidebar
      collapsible
      width={274}
      collapsedWidth={54}
      collapsed={collapsed}
      onCollapse={value => onCollapse(value)}
      trigger={<Icon icon={collapsed ? "panelToggleLeft" : "panelToggleRight"} />}>
      <ContentWrapper>
        {collapsed ? (
          <StyledIcon icon="message" />
        ) : (
          <>
            <ThreadWrapper>
              <Title>{t("Comments")}</Title>
              <CommentsContainer>
                <Thread
                  me={me}
                  comments={comments}
                  onCommentUpdate={onCommentUpdate}
                  onCommentDelete={onCommentDelete}
                />
              </CommentsContainer>
            </ThreadWrapper>

            {!comments || comments.length === 0 ? (
              <EmptyTextWrapper>{emptyText}</EmptyTextWrapper>
            ) : null}

            <Editor disabled={!comments} onCommentCreate={onCommentCreate} />
          </>
        )}
      </ContentWrapper>
    </Sidebar>
  );
};

export default CommentsPanel;

const StyledIcon = styled(Icon)`
  padding: 12px 20px;
`;

const ThreadWrapper = styled.div`
  padding: 12px;
  overflow: auto;
`;

const Title = styled.h3`
  font-size: 18px;
`;

const CommentsContainer = styled.div`
  overflow: auto;
`;

const EmptyTextWrapper = styled.div`
  padding: 12px;
  color: #00000073;
  text-align: center;
`;

const ContentWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  height: 100%;
`;
