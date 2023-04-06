import styled from "@emotion/styled";
import { useCallback, useState } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Modal from "@reearth-cms/components/atoms/Modal";
import IntegrationCard from "@reearth-cms/components/molecules/Integration/IntegrationConnectModal/IntegrationCard";
import { Integration } from "@reearth-cms/components/molecules/Integration/types";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  integrations?: Integration[];
  open?: boolean;
  onClose: () => void;
  onSubmit: (integration?: Integration) => Promise<void> | void;
};

const IntegrationConnectModal: React.FC<Props> = ({ integrations, open, onClose, onSubmit }) => {
  const t = useT();
  const [selectedIntegration, SetSelectedIntegration] = useState<Integration | undefined>();

  const handleIntegrationSelect = useCallback(
    (integration: Integration) => {
      SetSelectedIntegration(integration);
    },
    [SetSelectedIntegration],
  );

  return (
    <Modal
      afterClose={() => SetSelectedIntegration(undefined)}
      title={t("Connect Integration")}
      visible={open}
      onCancel={onClose}
      footer={[
        <Button key="back" onClick={onClose}>
          {t("Cancel")}
        </Button>,
        <Button key="submit" type="primary" onClick={() => onSubmit(selectedIntegration)}>
          {t("Connect")}
        </Button>,
      ]}>
      <ModalContent>
        {integrations?.map(integration => (
          <IntegrationCard
            onClick={() => handleIntegrationSelect(integration)}
            key={integration.id}
            integration={integration}
            integrationSelected={integration.id === selectedIntegration?.id}
          />
        ))}
      </ModalContent>
    </Modal>
  );
};

const ModalContent = styled.div`
  max-height: 274px;
  overflow-y: scroll;
  overflow-x: hidden;
`;

export default IntegrationConnectModal;
