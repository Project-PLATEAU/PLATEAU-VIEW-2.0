import { useCallback, useMemo } from "react";
import { useNavigate, useParams } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { User } from "@reearth-cms/components/molecules/AccountSettings/types";
import { Request } from "@reearth-cms/components/molecules/Request/types";
import { getInitialFormValues } from "@reearth-cms/components/organisms/Project/Request/convertRequest";
import {
  useDeleteRequestMutation,
  useApproveRequestMutation,
  useAddCommentMutation,
  useGetMeQuery,
  useUpdateCommentMutation,
  useDeleteCommentMutation,
  useGetRequestQuery,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useProject, useWorkspace } from "@reearth-cms/state";

export default () => {
  const t = useT();
  const navigate = useNavigate();
  const [currentProject] = useProject();
  const [currentWorkspace] = useWorkspace();
  const { requestId } = useParams();

  const { data: userData } = useGetMeQuery();
  const { data: rawRequest, loading: requestLoading } = useGetRequestQuery({
    variables: { requestId: requestId ?? "" },
    skip: !requestId,
  });

  const me: User | undefined = useMemo(() => {
    return userData?.me
      ? {
          id: userData.me.id,
          name: userData.me.name,
          lang: userData.me.lang,
          email: userData.me.email,
        }
      : undefined;
  }, [userData]);

  const myRole = useMemo(
    () => currentWorkspace?.members?.find(m => m.userId === me?.id).role,
    [currentWorkspace?.members, me?.id],
  );

  const projectId = useMemo(() => currentProject?.id, [currentProject]);

  const currentRequest: Request | undefined = useMemo(() => {
    if (!rawRequest) return;
    if (rawRequest.node?.__typename !== "Request") return;
    const r = rawRequest.node;
    return {
      id: r.id,
      threadId: r.thread?.id ?? "",
      title: r.title,
      description: r.description ?? "",
      comments:
        r.thread?.comments?.map(c => ({
          id: c.id,
          author: {
            id: c.author?.id,
            name: c.author?.name ?? "Anonymous",
            type: c.author ? (c.author.__typename === "User" ? "User" : "Integration") : null,
          },
          content: c.content,
          createdAt: c.createdAt.toString(),
        })) ?? [],
      createdAt: r.createdAt,
      reviewers: r.reviewers,
      state: r.state,
      createdBy: r.createdBy ?? undefined,
      updatedAt: r.updatedAt,
      approvedAt: r.approvedAt ?? undefined,
      closedAt: r.closedAt ?? undefined,
      items: r.items.map(item => ({
        id: item.itemId,
        modelName: item?.item?.value.model.name,
        initialValues: getInitialFormValues(item.item?.value.fields),
        schema: item.item?.value.schema ? item.item?.value.schema : undefined,
      })),
    };
  }, [rawRequest]);

  const isCloseActionEnabled: boolean = useMemo(
    () =>
      currentRequest?.state !== "CLOSED" &&
      currentRequest?.state !== "APPROVED" &&
      !!currentRequest?.reviewers.find(reviewer => reviewer.id === me?.id) &&
      myRole !== "READER" &&
      myRole !== "WRITER",
    [currentRequest?.reviewers, currentRequest?.state, me?.id, myRole],
  );

  const isApproveActionEnabled: boolean = useMemo(
    () =>
      currentRequest?.state === "WAITING" &&
      !!currentRequest?.reviewers.find(reviewer => reviewer.id === me?.id) &&
      myRole !== "READER" &&
      myRole !== "WRITER",
    [currentRequest?.reviewers, currentRequest?.state, me?.id, myRole],
  );

  const [deleteRequestMutation] = useDeleteRequestMutation();
  const handleRequestDelete = useCallback(
    (requestsId: string[]) =>
      (async () => {
        if (!projectId) return;
        const result = await deleteRequestMutation({
          variables: { projectId, requestsId },
          refetchQueries: ["GetRequests"],
        });
        if (result.errors) {
          Notification.error({ message: t("Failed to delete one or more requests.") });
        }
        if (result) {
          Notification.success({ message: t("One or more requests were successfully closed!") });
          navigate(`/workspace/${currentWorkspace?.id}/project/${projectId}/request`);
        }
      })(),
    [t, projectId, currentWorkspace?.id, navigate, deleteRequestMutation],
  );

  const [approveRequestMutation] = useApproveRequestMutation();
  const handleRequestApprove = useCallback(
    (requestId: string) =>
      (async () => {
        const result = await approveRequestMutation({
          variables: { requestId },
          refetchQueries: ["GetRequests"],
        });
        if (result.errors) {
          Notification.error({ message: t("Failed to approve request.") });
        }
        if (result) {
          Notification.success({ message: t("Successfully approved request!") });
          navigate(`/workspace/${currentWorkspace?.id}/project/${projectId}/request`);
        }
      })(),
    [currentWorkspace?.id, projectId, navigate, t, approveRequestMutation],
  );

  const [createComment] = useAddCommentMutation({
    refetchQueries: ["GetRequests"],
  });

  const handleCommentCreate = useCallback(
    async (content: string) => {
      if (!currentRequest?.threadId) return;
      const comment = await createComment({
        variables: {
          threadId: currentRequest.threadId,
          content,
        },
      });
      if (comment.errors || !comment.data?.addComment) {
        Notification.error({ message: t("Failed to create comment.") });
        return;
      }
      Notification.success({ message: t("Successfully created comment!") });
    },
    [createComment, currentRequest?.threadId, t],
  );

  const handleNavigateToRequestsList = () => {
    navigate(`/workspace/${currentWorkspace?.id}/project/${projectId}/request`);
  };

  const handleNavigateToItemEditForm = useCallback(
    (itemId: string, modelId?: string) => {
      if (!modelId) return;
      window.open(
        `/workspace/${currentWorkspace?.id}/project/${currentProject?.id}/content/${modelId}/details/${itemId}`,
      );
    },
    [currentWorkspace?.id, currentProject?.id],
  );

  const [updateComment] = useUpdateCommentMutation({
    refetchQueries: ["GetRequests"],
  });

  const handleCommentUpdate = useCallback(
    async (commentId: string, content: string) => {
      if (!currentRequest?.threadId) return;
      const comment = await updateComment({
        variables: {
          threadId: currentRequest.threadId,
          commentId,
          content,
        },
      });
      if (comment.errors || !comment.data?.updateComment) {
        Notification.error({ message: t("Failed to update comment.") });
        return;
      }
      Notification.success({ message: t("Successfully updated comment!") });
    },
    [updateComment, currentRequest?.threadId, t],
  );

  const [deleteComment] = useDeleteCommentMutation({
    refetchQueries: ["GetRequests"],
  });

  const handleCommentDelete = useCallback(
    async (commentId: string) => {
      if (!currentRequest?.threadId) return;
      const comment = await deleteComment({
        variables: {
          threadId: currentRequest.threadId,
          commentId,
        },
      });
      if (comment.errors || !comment.data?.deleteComment) {
        Notification.error({ message: t("Failed to delete comment.") });
        return;
      }
      Notification.success({ message: t("Successfully deleted comment!") });
    },
    [deleteComment, currentRequest?.threadId, t],
  );

  return {
    me,
    isCloseActionEnabled,
    loading: requestLoading,
    isApproveActionEnabled,
    currentRequest,
    handleRequestDelete,
    handleRequestApprove,
    handleCommentCreate,
    handleCommentUpdate,
    handleCommentDelete,
    handleNavigateToRequestsList,
    handleNavigateToItemEditForm,
  };
};
