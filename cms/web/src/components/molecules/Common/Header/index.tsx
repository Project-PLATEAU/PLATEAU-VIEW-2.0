import styled from "@emotion/styled";
import { useCallback, useMemo } from "react";
import { useNavigate } from "react-router-dom";

import { useAuth } from "@reearth-cms/auth";
import Header from "@reearth-cms/components/atoms/Header";
import Icon from "@reearth-cms/components/atoms/Icon";
import Menu from "@reearth-cms/components/atoms/Menu";
import Tooltip from "@reearth-cms/components/atoms/Tooltip";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";
import { useT } from "@reearth-cms/i18n";
import { Project, Workspace } from "@reearth-cms/state";

import HeaderDropdown from "./Dropdown";

export type { User } from "./types";

export interface Props {
  username?: string;
  personalWorkspace?: Workspace;
  currentWorkspace?: Workspace;
  workspaces?: any[];
  currentProject?: Project;
  onWorkspaceModalOpen: () => void;
  onNavigateToSettings: () => void;
  logoUrl?: string;
}

const HeaderMolecule: React.FC<Props> = ({
  username,
  personalWorkspace,
  currentWorkspace,
  workspaces,
  currentProject,
  onWorkspaceModalOpen,
  onNavigateToSettings,
  logoUrl,
}) => {
  const t = useT();
  const { logout } = useAuth();
  const navigate = useNavigate();

  const currentIsPersonal = useMemo(
    () => currentWorkspace?.id === personalWorkspace?.id,
    [currentWorkspace?.id, personalWorkspace?.id],
  );

  const handleWorkspaceNavigation = useCallback(
    (id: number) => {
      navigate(`/workspace/${id}`);
    },
    [navigate],
  );

  const handleHomeNavigation = useCallback(() => {
    navigate(`/workspace/${currentWorkspace?.id}`);
  }, [currentWorkspace?.id, navigate]);

  const WorkspacesMenu = (
    <HeaderMenu
      items={[
        {
          label: t("Personal Account"),
          key: "personal-account",
          type: "group",
          children: workspaces
            ?.filter(workspace => workspace.id === personalWorkspace?.id)
            ?.map(workspace => ({
              label: (
                <Tooltip title={workspace.name} placement="right">
                  <MenuText>{workspace.name}</MenuText>
                </Tooltip>
              ),
              key: workspace.id,
              icon: <UserAvatar username={workspace.name} size="small" />,
              style: { paddingLeft: 0, paddingRight: 0 },
              onClick: () => handleWorkspaceNavigation(workspace.id),
            })),
        },
        {
          type: "divider",
        },
        {
          label: t("Workspaces"),
          key: "workspaces",
          type: "group",
          children: workspaces
            ?.filter(workspace => workspace.id !== personalWorkspace?.id)
            ?.map(workspace => ({
              label: (
                <Tooltip title={workspace.name} placement="right">
                  <MenuText>{workspace.name}</MenuText>
                </Tooltip>
              ),
              key: workspace.id,
              icon: <UserAvatar username={workspace.name} size="small" shape="square" />,
              style: { paddingLeft: 0, paddingRight: 0 },
              onClick: () => handleWorkspaceNavigation(workspace.id),
            })),
        },
        {
          label: t("Create Workspace"),
          key: "new-workspace",
          icon: <Icon icon="userGroupAdd" />,
          onClick: onWorkspaceModalOpen,
        },
      ]}
    />
  );

  const AccountMenu = (
    <HeaderMenu
      items={[
        {
          label: t("Account Settings"),
          key: "account-settings",
          icon: <Icon icon="user" />,
          onClick: onNavigateToSettings,
        },
        {
          label: t("Logout"),
          key: "logout",
          icon: <Icon icon="logout" />,
          onClick: logout,
        },
      ]}
    />
  );

  return (
    <MainHeader>
      {logoUrl ? (
        <LogoIcon src={logoUrl} onClick={handleHomeNavigation} />
      ) : (
        <Logo onClick={handleHomeNavigation}>{t("Re:Earth CMS")}</Logo>
      )}
      <VerticalDivider />
      <WorkspaceDropdown
        name={currentWorkspace?.name}
        menu={WorkspacesMenu}
        personal={currentIsPersonal}
      />
      {currentProject?.name && (
        <CurrentProject>
          <Break>/</Break>
          <ProjectText>{currentProject.name}</ProjectText>
        </CurrentProject>
      )}
      <Spacer />
      <AccountDropdown name={username} menu={AccountMenu} personal={true} />
    </MainHeader>
  );
};

const MainHeader = styled(Header)`
  display: flex;
  align-items: center;
  height: 48px;
  line-height: 41px;
  padding: 0;
  background-color: #1d1d1d;

  .ant-space-item {
    color: #dbdbdb;
  }
`;

const Logo = styled.div`
  display: inline-block;
  color: #df3013;
  font-weight: 500;
  font-size: 14px;
  line-height: 48px;
  cursor: pointer;
  padding: 0 40px 0 20px;
`;

const LogoIcon = styled.img`
  width: 100px;
  margin: 0 0 0 10px;
  cursor: pointer;
`;

const Spacer = styled.div`
  flex: 1;
`;

const HeaderMenu = styled(Menu)`
  background-color: #141414;
  width: 190px;

  .ant-dropdown-menu-item-divider {
    background-color: #303030;
  }
  .ant-dropdown-menu-item-group-title,
  .ant-dropdown-menu-item,
  .ant-dropdown-menu-submenu-title {
    color: #fff;
  }
  .ant-dropdown-menu-item-group-title {
    font-weight: 400;
    font-size: 12px;
    line-height: 22px;
    user-select: none;
    color: #dbdbdb;
  }
  .ant-dropdown-menu-item-active {
    background-color: #1d1d1d;
  }

  svg {
    width: 16px;
    height: 16px;
  }
`;

const VerticalDivider = styled.div`
  display: inline-block;
  height: 32px;
  color: #fff;
  margin: 0;
  vertical-align: middle;
  border-top: 0;
  border-left: 1px solid #303030;
`;

const WorkspaceDropdown = styled(HeaderDropdown)`
  margin-left: 20px;
  padding-left: 20px;
`;

const AccountDropdown = styled(HeaderDropdown)`
  padding-right: 20px;
`;

const ProjectText = styled.p`
  margin: 0;
`;

const Break = styled.p`
  margin: 0 10px 0 10px;
`;

const CurrentProject = styled.div`
  height: 100%;
  margin: 0;
  display: flex;
  align-items: center;
  color: #dbdbdb;
`;

const MenuText = styled.p`
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
  width: 140px;
`;

export default HeaderMolecule;
