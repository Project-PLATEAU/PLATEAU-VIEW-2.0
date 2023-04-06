import styled from "@emotion/styled";
import { useCallback } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import ContentSection from "@reearth-cms/components/atoms/InnerContents/ContentSection";
import Modal from "@reearth-cms/components/atoms/Modal";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  onIntegrationDelete: () => Promise<void>;
};

const DangerZone: React.FC<Props> = ({ onIntegrationDelete }) => {
  const t = useT();
  const { confirm } = Modal;

  const handleWorkspaceDeleteConfirmation = useCallback(() => {
    confirm({
      title: t("Are you sure to remove this integration?"),
      icon: <Icon icon="exclamationCircle" />,
      content: (
        <>
          {t("Permanently remove your Integration and all of its contents from the Re:Earth CMS.")}
          <br />
          {t("Once the integration is removed, it will disappear from all workspaces.")}
        </>
      ),
      onOk() {
        onIntegrationDelete();
      },
    });
  }, [confirm, onIntegrationDelete, t]);

  return (
    <ContentSection title={t("Danger Zone")} danger>
      <Title>{t("Remove Integration")}</Title>
      <Text>
        {t(
          "Permanently remove your Integration and all of its contents from the Re:Earth CMS. This action is not reversible – please continue with caution.",
        )}
      </Text>

      <Button onClick={handleWorkspaceDeleteConfirmation} type="primary" danger>
        {t("Remove Integration")}
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
