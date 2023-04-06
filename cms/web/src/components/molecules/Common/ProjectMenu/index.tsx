import { ItemType } from "antd/lib/menu/hooks/useItems";
import { useCallback, useEffect, useState } from "react";
import { useNavigate } from "react-router-dom";

import Icon from "@reearth-cms/components/atoms/Icon";
import Menu from "@reearth-cms/components/atoms/Menu";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  inlineCollapsed: boolean;
  workspaceId?: string;
  projectId?: string;
  defaultSelectedKey?: string;
};

const ProjectMenu: React.FC<Props> = ({
  inlineCollapsed,
  workspaceId,
  projectId,
  defaultSelectedKey,
}) => {
  const t = useT();
  const navigate = useNavigate();

  const topItems: ItemType[] = [
    { label: t("Overview"), key: "home", icon: <Icon icon="dashboard" /> },
    { label: t("Schema"), key: "schema", icon: <Icon icon="unorderedList" /> },
    { label: t("Content"), key: "content", icon: <Icon icon="table" /> },
    { label: t("Asset"), key: "asset", icon: <Icon icon="file" /> },
    { label: t("Request"), key: "request", icon: <Icon icon="pullRequest" /> },
  ];
  const [selected, changeSelected] = useState([defaultSelectedKey ?? "home"]);

  useEffect(() => {
    if (defaultSelectedKey && defaultSelectedKey !== selected[0]) {
      changeSelected([defaultSelectedKey]);
    }
  }, [selected, defaultSelectedKey]);

  const items: ItemType[] = [
    {
      label: t("Accessibility"),
      key: "accessibility",
      icon: <Icon icon="send" />,
    },
    {
      label: t("Settings"),
      key: "settings",
      icon: <Icon icon="settings" />,
    },
  ];

  const onClick = useCallback(
    (e: any) => {
      changeSelected([e.key]);
      if (e.key === "schema") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/schema`);
      } else if (e.key === "content") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/content`);
      } else if (e.key === "asset") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/asset`);
      } else if (e.key === "request") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/request`);
      } else if (e.key === "accessibility") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/accessibility`);
      } else if (e.key === "settings") {
        navigate(`/workspace/${workspaceId}/project/${projectId}/settings`);
      } else {
        navigate(`/workspace/${workspaceId}/project/${projectId}`);
      }
    },
    [navigate, workspaceId, projectId],
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

export default ProjectMenu;
