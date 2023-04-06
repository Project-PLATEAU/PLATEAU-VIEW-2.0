import styled from "@emotion/styled";

import PageHeader from "@reearth-cms/components/atoms/PageHeader";
import MyIntegrationCard from "@reearth-cms/components/molecules/MyIntegrations/List/Card";
import IntegrationCreationAction from "@reearth-cms/components/molecules/MyIntegrations/List/CreationAction";
import { Integration } from "@reearth-cms/components/molecules/MyIntegrations/types";
import { useT } from "@reearth-cms/i18n";

export type Props = {
  integrations?: Integration[];
  onIntegrationModalOpen: () => void;
};

const MyIntegrationList: React.FC<Props> = ({ integrations, onIntegrationModalOpen }) => {
  const t = useT();

  return (
    <Wrapper>
      <PageHeader title={t("My Integrations")} />
      <ListWrapper>
        {integrations?.map((integration: Integration) => (
          <MyIntegrationCard key={integration.id} integration={integration} />
        ))}
        <IntegrationCreationAction onIntegrationModalOpen={onIntegrationModalOpen} />
      </ListWrapper>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background: #fff;
  min-height: 100%;
`;

const ListWrapper = styled.div`
  border-top: 1px solid #f0f0f0;
  padding: 12px;
  display: flex;
  flex-wrap: wrap;
  align-items: stretch;
`;

export default MyIntegrationList;
