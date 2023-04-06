import styled from "@emotion/styled";

import Icon from "@reearth-cms/components/atoms/Icon";
import { Integration } from "@reearth-cms/components/molecules/Integration/types";

export type Props = {
  integration: Integration;
  integrationSelected: boolean;
  onClick: () => void;
};

const IntegrationCard: React.FC<Props> = ({ integration, integrationSelected, onClick }) => (
  <CardWrapper onClick={onClick} isSelected={integrationSelected}>
    <Icon icon="api" size={64} color="#00000040" />
    <CardTitle>{integration.name}</CardTitle>
  </CardWrapper>
);

const CardWrapper = styled.div<{ isSelected?: boolean }>`
  cursor: pointer;
  box-shadow: 0px 2px 8px #00000026;
  border: 1px solid #f0f0f0;
  padding: 12px;
  min-height: 88px;
  display: flex;
  align-items: center;
  background-color: ${({ isSelected }) => (isSelected ? "#E6F7FF" : "#FFF")};
  margin-bottom: 10px;
`;

const CardTitle = styled.h5`
  font-weight: 500;
  font-size: 16px;
  line-height: 24px;
  margin: 0;
  padding-left: 12px;
`;

export default IntegrationCard;
