import { useCallback } from "react";
import { useNavigate } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import {
  useUpdateWorkspaceMutation,
  useDeleteWorkspaceMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useWorkspace } from "@reearth-cms/state";

export default () => {
  const navigate = useNavigate();
  const [currentWorkspace, setCurrentWorkspace] = useWorkspace();
  const t = useT();

  const workspaceId = currentWorkspace?.id;
  const workspaceName = currentWorkspace?.name;

  const [updateWorkspaceMutation] = useUpdateWorkspaceMutation();
  const [deleteWorkspaceMutation] = useDeleteWorkspaceMutation({
    refetchQueries: ["GetMe"],
  });

  const handleWorkspaceUpdate = useCallback(
    async (name?: string) => {
      if (!workspaceId || !name) return;
      const res = await updateWorkspaceMutation({
        variables: {
          workspaceId,
          name,
        },
      });
      if (res.errors || !res.data?.updateWorkspace) {
        Notification.error({ message: t("Failed to update workspace.") });
      } else {
        setCurrentWorkspace(res.data.updateWorkspace.workspace);
        Notification.success({ message: t("Successfully updated workspace!") });
      }
    },
    [workspaceId, updateWorkspaceMutation, setCurrentWorkspace, t],
  );

  const handleWorkspaceDelete = useCallback(async () => {
    if (!workspaceId) return;
    const results = await deleteWorkspaceMutation({ variables: { workspaceId } });
    if (results.errors) {
      Notification.error({ message: t("Failed to delete workspace.") });
    } else {
      setCurrentWorkspace(undefined);
      Notification.success({ message: t("Successfully deleted workspace!") });
      navigate(`/workspace`);
    }
  }, [workspaceId, deleteWorkspaceMutation, setCurrentWorkspace, navigate, t]);

  return {
    workspaceName,
    handleWorkspaceUpdate,
    handleWorkspaceDelete,
  };
};
