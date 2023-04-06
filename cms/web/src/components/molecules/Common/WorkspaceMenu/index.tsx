import { ItemType } from "antd/lib/menu/hooks/useItems";
import React, { useCallback, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Icon from "@reearth-cms/components/atoms/Icon";
import Menu from "@reearth-cms/components/atoms/Menu";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  inlineCollapsed: boolean;
  isPersonalWorkspace?: boolean;
  workspaceId?: string;
  defaultSelectedKey?: string;
};

export type MenuShowType = "personal" | "notPersonal" | "both";

export type WorkspaceItemType = ItemType & { show: MenuShowType };

const WorkspaceMenu: React.FC<Props> = ({
  inlineCollapsed,
  isPersonalWorkspace,
  workspaceId,
  defaultSelectedKey,
}) => {
  const t = useT();
  const navigate = useNavigate();
  const [selected, changeSelected] = useState([defaultSelectedKey ?? "home"]);

  useEffect(() => {
    if (defaultSelectedKey) {
      changeSelected([defaultSelectedKey]);
    }
  }, [defaultSelectedKey]);

  const topItems: WorkspaceItemType[] = [
    { label: t("Home"), key: "home", icon: <Icon icon="home" />, show: "both" },
  ];

  const items: WorkspaceItemType[] = [
    {
      label: t("Member"),
      key: "members",
      icon: <Icon icon="userGroupAdd" />,
      show: "notPersonal" as MenuShowType,
    },
    {
      label: t("Integrations"),
      key: "integrations",
      icon: <Icon icon="api" />,
      show: "both" as MenuShowType,
    },
    {
      label: t("My Integrations"),
      key: "myIntegrations",
      icon: <Icon icon="myIntegrations" />,
      show: "personal" as MenuShowType,
    },
    // {
    //   label: t("Role"),
    //   key: "role",
    //   icon: <Icon icon="userSwitch" />,
    //   show: "notPersonal" as MenuShowType,
    // },
    // {
    //   label: t("API key"),
    //   key: "apiKey",
    //   icon: <Icon icon="key" />,
    //   show: "both" as MenuShowType,
    // },
    {
      label: t("Account"),
      key: "account",
      icon: <Icon icon="user" />,
      show: "personal" as MenuShowType,
    },
    {
      label: t("Settings"),
      key: "settings",
      icon: <Icon icon="settings" />,
      show: "notPersonal" as MenuShowType,
    },
  ].filter(
    item =>
      (isPersonalWorkspace && item.show === "personal") ||
      (!isPersonalWorkspace && item.show === "notPersonal") ||
      item.show === "both",
  );

  const onClick = useCallback(
    (e: any) => {
      changeSelected([e.key]);
      if (e.key === "members") {
        navigate(`/workspace/${workspaceId}/members`);
      } else if (e.key === "myIntegrations") {
        navigate(`/workspace/${workspaceId}/myIntegrations`);
      } else if (e.key === "integrations") {
        navigate(`/workspace/${workspaceId}/integrations`);
      } else if (e.key === "role") {
        navigate(`/workspace/${workspaceId}/role`);
      } else if (e.key === "apiKey") {
        navigate(`/workspace/${workspaceId}/apiKey`);
      } else if (e.key === "settings") {
        navigate(`/workspace/${workspaceId}/settings`);
      } else if (e.key === "account") {
        navigate(`/workspace/${workspaceId}/account`);
      } else {
        navigate(`/workspace/${workspaceId}`);
      }
    },
    [navigate, workspaceId],
  );

  return (
    <>
      <Menu
        onClick={onClick}
        selectedKeys={selected}
        inlineCollapsed={inlineCollapsed}
        mode="inline"
        items={topItems}
      />
      <Menu
        onClick={onClick}
        selectedKeys={selected}
        inlineCollapsed={inlineCollapsed}
        mode="inline"
        items={items}
      />
    </>
  );
};

export default WorkspaceMenu;
