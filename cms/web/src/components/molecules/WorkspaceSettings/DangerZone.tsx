import styled from "@emotion/styled";
import { useCallback } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import ContentSection from "@reearth-cms/components/atoms/InnerContents/ContentSection";
import Modal from "@reearth-cms/components/atoms/Modal";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  onWorkspaceDelete: () => Promise<void>;
};

const DangerZone: React.FC<Props> = ({ onWorkspaceDelete }) => {
  const t = useT();
  const { confirm } = Modal;

  const handleWorkspaceDeleteConfirmation = useCallback(() => {
    confirm({
      title: t("Are you sure you want to delete this workspace?"),
      icon: <Icon icon="exclamationCircle" />,
      onOk() {
        onWorkspaceDelete();
      },
    });
  }, [confirm, onWorkspaceDelete, t]);

  return (
    <ContentSection title={t("Danger Zone")} danger>
      <Title>{t("Remove Workspace")}</Title>
      <Text>
        {t(
          "Permanently removes the current workspace and all of its contents from Re:Earth CMS. This action is not reversible, so please continue with caution.",
        )}
      </Text>

      <Button onClick={handleWorkspaceDeleteConfirmation} type="primary" danger>
        {t("Remove Workspace")}
      </Button>
    </ContentSection>
  );
};

export default DangerZone;

const Title = styled.h1`
  font-weight: 500;
  font-size: 16px;
  line-height: 24px;
  color: #000000d9;
`;

const Text = styled.p`
  font-weight: 400;
  font-size: 14px;
  line-height: 22px;
  color: #000000d9;
  margin: 24px 0;
`;
