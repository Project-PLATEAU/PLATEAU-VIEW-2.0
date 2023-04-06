import { useCallback, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { Project } from "@reearth-cms/components/molecules/Workspace/types";
import {
  useGetProjectsQuery,
  useCreateProjectMutation,
  useCreateWorkspaceMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useWorkspace } from "@reearth-cms/state";

export default () => {
  const t = useT();
  const navigate = useNavigate();
  const coverImageUrl = window.REEARTH_CONFIG?.coverImageUrl;

  const [currentWorkspace, setCurrentWorkspace] = useWorkspace();

  const [searchedProjectName, setSearchedProjectName] = useState<string>("");

  const [workspaceModalShown, setWorkspaceModalShown] = useState(false);
  const [projectModalShown, setProjectModalShown] = useState(false);

  const workspaceId = currentWorkspace?.id;

  const {
    data,
    loading: loadingProjects,
    refetch,
  } = useGetProjectsQuery({
    variables: { workspaceId: workspaceId ?? "", pagination: { first: 100 } },
    skip: !workspaceId,
  });

  const projects = useMemo(() => {
    return data?.projects.nodes
      .map<Project | undefined>(project =>
        project
          ? {
              id: project.id,
              description: project.description,
              name: project.name,
            }
          : undefined,
      )
      .filter(
        (project): project is Project =>
          !!project &&
          (!searchedProjectName ||
            (!!searchedProjectName &&
              project.name.toLocaleLowerCase().includes(searchedProjectName.toLocaleLowerCase()))),
      );
  }, [data?.projects.nodes, searchedProjectName]);

  const [createNewProject] = useCreateProjectMutation({
    refetchQueries: ["GetProjects"],
  });

  const handleProjectSearch = useCallback(
    (value: string) => {
      setSearchedProjectName?.(value);
    },
    [setSearchedProjectName],
  );

  const handleProjectCreate = useCallback(
    async (data: { name: string; alias: string; description: string }) => {
      if (!workspaceId) return;
      const project = await createNewProject({
        variables: {
          workspaceId,
          name: data.name,
          alias: data.alias,
          description: data.description,
        },
      });
      if (project.errors || !project.data?.createProject) {
        Notification.error({ message: t("Failed to create project.") });
        return;
      }
      Notification.success({ message: t("Successfully created project!") });
      setProjectModalShown(false);
      refetch();
    },
    [createNewProject, workspaceId, refetch, t],
  );

  const handleProjectModalClose = useCallback(() => {
    setProjectModalShown(false);
  }, []);

  const handleProjectModalOpen = useCallback(() => setProjectModalShown(true), []);

  const handleProjectNavigation = useCallback(
    (project?: Project) => {
      navigate(`/workspace/${currentWorkspace?.id}/project/${project?.id}`);
    },
    [currentWorkspace, navigate],
  );

  const [createWorkspaceMutation] = useCreateWorkspaceMutation();
  const handleWorkspaceCreate = useCallback(
    async (data: { name: string }) => {
      const results = await createWorkspaceMutation({
        variables: { name: data.name },
        refetchQueries: ["GetWorkspaces"],
      });
      if (results.data?.createWorkspace) {
        Notification.success({ message: t("Successfully created workspace!") });
        setCurrentWorkspace(results.data.createWorkspace.workspace);
        navigate(`/workspace/${results.data.createWorkspace.workspace.id}`);
      }
      refetch();
    },
    [createWorkspaceMutation, setCurrentWorkspace, refetch, navigate, t],
  );

  const handleWorkspaceModalClose = useCallback(() => {
    setWorkspaceModalShown(false);
  }, []);

  const handleWorkspaceModalOpen = useCallback(() => setWorkspaceModalShown(true), []);

  return {
    projects,
    projectModalShown,
    loadingProjects,
    workspaceModalShown,
    handleProjectSearch,
    handleProjectCreate,
    handleProjectModalOpen,
    handleProjectModalClose,
    handleProjectNavigation,
    handleWorkspaceModalClose,
    handleWorkspaceModalOpen,
    handleWorkspaceCreate,
    coverImageUrl,
  };
};
