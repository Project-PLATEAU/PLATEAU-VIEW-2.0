import styled from "@emotion/styled";
import { useCallback } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Content from "@reearth-cms/components/atoms/Content";
import Icon from "@reearth-cms/components/atoms/Icon";
import Modal from "@reearth-cms/components/atoms/Modal";
import PageHeader from "@reearth-cms/components/atoms/PageHeader";
import Search from "@reearth-cms/components/atoms/Search";
import Table from "@reearth-cms/components/atoms/Table";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";
import MemberAddModal from "@reearth-cms/components/molecules/Member/MemberAddModal";
import MemberRoleModal from "@reearth-cms/components/molecules/Member/MemberRoleModal";
import { Member } from "@reearth-cms/components/molecules/Workspace/types";
import { useT } from "@reearth-cms/i18n";

import useHooks from "./hooks";

const columns = [
  {
    title: "Name",
    dataIndex: "name",
    key: "name",
  },
  {
    title: "Thumbnail",
    dataIndex: "thumbnail",
    key: "thumbnail",
  },
  {
    title: "Email",
    dataIndex: "email",
    key: "email",
  },
  {
    title: "Role",
    dataIndex: "role",
    key: "role",
  },
  {
    title: "Action",
    dataIndex: "action",
    key: "action",
  },
];

const Members: React.FC = () => {
  const t = useT();

  const { confirm } = Modal;

  const {
    me,
    owner,
    searchedUser,
    handleSearchTerm,
    changeSearchedUser,
    searchedUserList,
    changeSearchedUserList,
    handleUserSearch,
    handleUserAdd,
    handleUsersAddToWorkspace,
    handleMemberOfWorkspaceUpdate,
    selectedMember,
    roleModalShown,
    handleMemberRemoveFromWorkspace,
    handleRoleModalClose,
    handleRoleModalOpen,
    handleMemberAddModalClose,
    handleMemberAddModalOpen,
    MemberAddModalShown,
    workspaceUserMembers,
  } = useHooks();

  const handleMemberDelete = useCallback(
    (member: Member) => {
      confirm({
        title: t("Are you sure to remove this member?"),
        icon: <Icon icon="exclamationCircle" />,
        content: t(
          "Remove this member from workspace means this member will not view any content of this workspace.",
        ),
        onOk() {
          handleMemberRemoveFromWorkspace(member?.userId);
        },
      });
    },
    [confirm, handleMemberRemoveFromWorkspace, t],
  );

  const dataSource = workspaceUserMembers?.map(member => ({
    key: member.userId,
    name: member.user.name,
    thumbnail: <UserAvatar username={member.user.name} />,
    email: member.user.email,
    role: member.role,
    action: (
      <>
        {member.userId !== me?.id && (
          <Button type="link" onClick={() => handleRoleModalOpen(member)} disabled={!owner}>
            {t("Change Role?")}
          </Button>
        )}
        {member.role !== "OWNER" && (
          <Button
            type="link"
            style={{ marginLeft: "8px" }}
            onClick={() => handleMemberDelete(member)}
            disabled={!owner}>
            {t("Remove")}
          </Button>
        )}
      </>
    ),
  }));

  return (
    <>
      <PaddedContent>
        <PageHeader
          title={t("Members")}
          extra={
            <Button
              type="primary"
              onClick={handleMemberAddModalOpen}
              icon={<Icon icon="userGroupAdd" />}>
              {t("New Member")}
            </Button>
          }
        />
        <ActionHeader>
          <Search
            onSearch={handleSearchTerm}
            placeholder={t("search for a member")}
            allowClear
            style={{ width: 264 }}
          />
        </ActionHeader>
        <Table
          dataSource={dataSource}
          columns={columns}
          style={{ padding: "24px", overflowX: "auto" }}
        />
      </PaddedContent>
      <MemberRoleModal
        member={selectedMember}
        open={roleModalShown}
        onClose={handleRoleModalClose}
        onSubmit={handleMemberOfWorkspaceUpdate}
      />
      <MemberAddModal
        open={MemberAddModalShown}
        searchedUser={searchedUser}
        searchedUserList={searchedUserList}
        changeSearchedUserList={changeSearchedUserList}
        onClose={handleMemberAddModalClose}
        onUserSearch={handleUserSearch}
        onUserAdd={handleUserAdd}
        changeSearchedUser={changeSearchedUser}
        onSubmit={handleUsersAddToWorkspace}
      />
    </>
  );
};

const PaddedContent = styled(Content)`
  margin: 16px;
  background-color: #fff;
  min-height: 100%;
`;

const ActionHeader = styled(Content)`
  border-top: 1px solid #f0f0f0;
  padding: 16px;
  display: flex;
  justify-content: space-between;
`;

export default Members;
