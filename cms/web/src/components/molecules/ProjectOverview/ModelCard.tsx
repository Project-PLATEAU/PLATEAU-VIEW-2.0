import styled from "@emotion/styled";

import Card from "@reearth-cms/components/atoms/Card";
import Dropdown from "@reearth-cms/components/atoms/Dropdown";
import Icon from "@reearth-cms/components/atoms/Icon";
import Menu from "@reearth-cms/components/atoms/Menu";
import { useT } from "@reearth-cms/i18n";

export type Model = {
  id: string;
  name: string;
  description?: string;
  key: string;
};

export type Props = {
  model: Model;
  onSchemaNavigate?: (modelId: string) => void;
  onContentNavigate?: (modelId: string) => void;
  onModelDeletionModalOpen: (model: Model) => Promise<void>;
  onModelUpdateModalOpen: (model: Model) => Promise<void>;
};

const ModelCard: React.FC<Props> = ({
  model,
  onSchemaNavigate,
  onContentNavigate,
  onModelDeletionModalOpen,
  onModelUpdateModalOpen,
}) => {
  const { Meta } = Card;
  const t = useT();

  const ModelMenu = (
    <Menu
      items={[
        {
          key: "edit",
          label: t("Edit"),
          onClick: () => onModelUpdateModalOpen(model),
        },
        {
          key: "delete",
          label: t("Delete"),
          onClick: () => onModelDeletionModalOpen(model),
        },
      ]}
    />
  );

  return (
    <StyledCard
      actions={[
        <Icon icon="unorderedList" key="schema" onClick={() => onSchemaNavigate?.(model.id)} />,
        <Icon icon="table" key="content" onClick={() => onContentNavigate?.(model.id)} />,
        <Dropdown key="options" overlay={ModelMenu} trigger={["click"]}>
          <Icon icon="ellipsis" />
        </Dropdown>,
      ]}>
      <Meta title={model.name} description={model.description} />
    </StyledCard>
  );
};

export default ModelCard;

const StyledCard = styled(Card)`
  .ant-card-body {
    height: 102px;
  }
  .ant-card-meta-description {
    height: 40px;
    overflow: hidden;
    text-overflow: ellipsis;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
  }
`;
