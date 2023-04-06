import { useCallback, useEffect, useMemo, useState } from "react";

import Notification from "@reearth-cms/components/atoms/Notification";
import { User } from "@reearth-cms/components/molecules/Member/types";
import { Member, MemberInput } from "@reearth-cms/components/molecules/Workspace/types";
import {
  useGetWorkspacesQuery,
  useAddUsersToWorkspaceMutation,
  useUpdateMemberOfWorkspaceMutation,
  Role,
  useRemoveMemberFromWorkspaceMutation,
  Workspace,
  useGetUserBySearchLazyQuery,
  MemberInput as GQLMemberInput,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useWorkspace } from "@reearth-cms/state";
import { stringSortCallback } from "@reearth-cms/utils/sort";

export type RoleUnion = "READER" | "WRITER" | "MAINTAINER" | "OWNER";

export default () => {
  const [currentWorkspace, setWorkspace] = useWorkspace();
  const [roleModalShown, setRoleModalShown] = useState(false);
  const [MemberAddModalShown, setMemberAddModalShown] = useState(false);
  const [selectedMember, setSelectedMember] = useState<Member | undefined>(undefined);
  const [searchTerm, setSearchTerm] = useState<string>();
  const [owner, setOwner] = useState(false);
  const t = useT();

  const handleSearchTerm = useCallback((term?: string) => {
    setSearchTerm(term);
  }, []);

  const [searchedUser, changeSearchedUser] = useState<User>();
  const [searchedUserList, changeSearchedUserList] = useState<User[]>([]);

  const { data, loading } = useGetWorkspacesQuery();
  const me = { id: data?.me?.id, myWorkspace: data?.me?.myWorkspace.id };
  const workspaces = data?.me?.workspaces as Workspace[];
  const workspaceId = currentWorkspace?.id;

  const isOwner = useMemo(
    () => currentWorkspace?.members?.find(m => m.userId === me?.id && m.role === "OWNER"),
    [currentWorkspace?.members, me?.id],
  );

  useEffect(() => {
    setOwner(isOwner);
  }, [isOwner]);

  useEffect(() => {
    if (!workspaceId) return;
    if (!currentWorkspace) {
      setWorkspace(
        workspaceId
          ? workspaces?.find(t => t.id === workspaceId)
          : data?.me?.myWorkspace ?? undefined,
      );
    }
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [currentWorkspace, setWorkspace, workspaces, data?.me]);

  const [searchUserQuery, { data: searchUserData }] = useGetUserBySearchLazyQuery({
    fetchPolicy: "no-cache",
  });

  useEffect(() => {
    changeSearchedUser(
      searchUserData?.searchUser && searchUserData?.searchUser?.id !== data?.me?.id
        ? searchUserData.searchUser
        : undefined,
    );
  }, [searchUserData?.searchUser, data?.me?.id]);

  const handleUserSearch = useCallback(
    (nameOrEmail: string) => nameOrEmail && searchUserQuery({ variables: { nameOrEmail } }),
    [searchUserQuery],
  );

  const handleUserAdd = useCallback(() => {
    if (
      searchedUser &&
      searchedUser.id !== data?.me?.id &&
      !searchedUserList.find(user => user.id === searchedUser.id)
    ) {
      changeSearchedUserList([...searchedUserList, searchedUser]);
    }
  }, [data?.me?.id, searchedUser, searchedUserList]);

  const workspaceUserMembers = useMemo((): Member[] | undefined => {
    return currentWorkspace?.members
      ?.map<Member | undefined>(member =>
        member && member.__typename === "WorkspaceUserMember" && member.user
          ? {
              userId: member.userId,
              user: member.user,
              role: member.role,
            }
          : undefined,
      )
      .filter(
        (user): user is Member =>
          !!user && user.user.name.toLowerCase().includes(searchTerm?.toLowerCase() ?? ""),
      )
      .sort((user1, user2) => stringSortCallback(user1.userId, user2.userId));
  }, [currentWorkspace, searchTerm]);

  const [addUsersToWorkspaceMutation] = useAddUsersToWorkspaceMutation();

  const handleUsersAddToWorkspace = useCallback(
    (users: MemberInput[]) =>
      (async () => {
        if (!workspaceId) return;
        const result = await addUsersToWorkspaceMutation({
          variables: { workspaceId, users: users as GQLMemberInput[] },
          refetchQueries: ["GetWorkspaces"],
        });
        const workspace = result.data?.addUsersToWorkspace?.workspace;
        if (result.errors || !workspace) {
          Notification.error({ message: t("Failed to add one or more members.") });
          return;
        }
        setWorkspace(workspace);

        if (result.data) {
          Notification.success({ message: t("Successfully added member(s) to the workspace!") });
        }
      })(),
    [workspaceId, addUsersToWorkspaceMutation, setWorkspace, t],
  );

  const [updateMemberOfWorkspaceMutation] = useUpdateMemberOfWorkspaceMutation();

  const handleMemberOfWorkspaceUpdate = useCallback(
    async (userId: string, role: RoleUnion) => {
      if (workspaceId) {
        const results = await updateMemberOfWorkspaceMutation({
          variables: {
            workspaceId,
            userId,
            role: {
              READER: Role.Reader,
              WRITER: Role.Writer,
              MAINTAINER: Role.Maintainer,
              OWNER: Role.Owner,
            }[role],
          },
        });
        const workspace = results.data?.updateUserOfWorkspace?.workspace;
        if (workspace) {
          setWorkspace(workspace);
        }
      }
    },
    [workspaceId, setWorkspace, updateMemberOfWorkspaceMutation],
  );

  const [removeMemberFromWorkspaceMutation] = useRemoveMemberFromWorkspaceMutation();

  const handleMemberRemoveFromWorkspace = useCallback(
    async (userId: string) => {
      if (!workspaceId) return;
      const result = await removeMemberFromWorkspaceMutation({
        variables: { workspaceId, userId },
        refetchQueries: ["GetWorkspaces"],
      });
      const workspace = result.data?.removeUserFromWorkspace?.workspace;
      if (result.errors || !workspace) {
        Notification.error({ message: t("Failed to delete member from the workspace.") });
        return;
      }
      setWorkspace(workspace);
      Notification.success({ message: t("Successfully removed member from the workspace!") });
    },
    [workspaceId, removeMemberFromWorkspaceMutation, setWorkspace, t],
  );

  const handleRoleModalClose = useCallback(() => {
    setRoleModalShown(false);
    setSelectedMember(undefined);
  }, []);

  const handleRoleModalOpen = useCallback((member: Member) => {
    setRoleModalShown(true);
    setSelectedMember(member);
  }, []);

  const handleMemberAddModalClose = useCallback(() => {
    setMemberAddModalShown(false);
    setSelectedMember(undefined);
  }, []);

  const handleMemberAddModalOpen = useCallback(() => {
    setMemberAddModalShown(true);
    setSelectedMember(undefined);
  }, []);

  return {
    me,
    owner,
    workspaces,
    currentWorkspace,
    searchedUser,
    handleSearchTerm,
    changeSearchedUser,
    searchedUserList,
    changeSearchedUserList,
    handleUserSearch,
    handleUserAdd,
    handleUsersAddToWorkspace,
    handleMemberOfWorkspaceUpdate,
    handleMemberRemoveFromWorkspace,
    handleRoleModalClose,
    handleRoleModalOpen,
    handleMemberAddModalClose,
    handleMemberAddModalOpen,
    MemberAddModalShown,
    setSelectedMember,
    selectedMember,
    roleModalShown,
    loading,
    workspaceUserMembers,
  };
};
