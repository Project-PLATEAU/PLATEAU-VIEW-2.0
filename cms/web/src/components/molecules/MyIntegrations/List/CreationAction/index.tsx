import styled from "@emotion/styled";

import Icon from "@reearth-cms/components/atoms/Icon";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  onIntegrationModalOpen: () => void;
};

const IntegrationCreationAction: React.FC<Props> = ({ onIntegrationModalOpen }) => {
  const t = useT();

  return (
    <CardWrapper>
      <Card onClick={onIntegrationModalOpen}>
        <Icon icon="plus" style={{ fontSize: 36 }} />
        <CardTitle>{t("Create new integration")}</CardTitle>
      </Card>
    </CardWrapper>
  );
};

const CardWrapper = styled.div`
  padding: 12px;
`;

const Card = styled.div`
  justify-content: center;
  height: 180px;
  width: 240px;
  display: flex;
  flex-direction: column;
  align-items: center;
  padding: 24px;
  border: 1px solid #f0f0f0;
  box-shadow: 0px 2px 8px #00000026;
  border-radius: 4px;
  color: #00000073;
  cursor: pointer;
  &:hover {
    color: #1890ff;
    background-color: #e6f7ff;
  }
`;

const CardTitle = styled.p`
  margin-top: 8px;
  font-weight: 500;
  font-size: 14px;
  line-height: 22px;
`;

export default IntegrationCreationAction;
