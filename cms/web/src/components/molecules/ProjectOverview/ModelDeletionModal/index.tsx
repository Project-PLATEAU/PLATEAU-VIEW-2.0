import styled from "@emotion/styled";

import Alert from "@reearth-cms/components/atoms/Alert";
import Button from "@reearth-cms/components/atoms/Button";
import Modal from "@reearth-cms/components/atoms/Modal";
import { Model } from "@reearth-cms/components/molecules/ProjectOverview";
import { useT } from "@reearth-cms/i18n";

export interface Props {
  open: boolean;
  model?: Model;
  onClose: (refetch?: boolean) => void;
  onDelete: (modelId?: string) => Promise<void> | void;
}

const ModelDeletionModal: React.FC<Props> = ({ open, model, onClose, onDelete }) => {
  const t = useT();

  return (
    <Modal
      title={t("Delete Model")}
      visible={open}
      onCancel={() => onClose()}
      footer={[
        <Button key="back" onClick={() => onClose()}>
          {t("Cancel")}
        </Button>,
        <Button key="submit" type="primary" onClick={() => onDelete(model?.id)} danger>
          {t("Delete Model")}
        </Button>,
      ]}>
      <p>
        {t("Are you sure you want to delete the model")} <ModelName> {model?.name} </ModelName>?
      </p>
      <Alert
        message={t("Warning")}
        description={t(
          "This action will permanently delete the selected model and cannot be reversed.",
        )}
        type="warning"
        showIcon
      />
    </Modal>
  );
};

export default ModelDeletionModal;

const ModelName = styled.span`
  font-weight: 600;
`;
