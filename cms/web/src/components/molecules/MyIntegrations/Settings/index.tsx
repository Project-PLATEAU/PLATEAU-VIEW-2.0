import styled from "@emotion/styled";

import Content from "@reearth-cms/components/atoms/Content";
import DangerZone from "@reearth-cms/components/molecules/MyIntegrations/Settings/DangerZone";
import MyIntegrationForm from "@reearth-cms/components/molecules/MyIntegrations/Settings/Form";
import { Integration } from "@reearth-cms/components/molecules/MyIntegrations/types";

export type Props = {
  integration: Integration;
  onIntegrationUpdate: (data: { name: string; description: string; logoUrl: string }) => void;
  onIntegrationDelete: () => Promise<void>;
};

const MyIntegrationSettings: React.FC<Props> = ({
  integration,
  onIntegrationUpdate,
  onIntegrationDelete,
}) => {
  return (
    <Wrapper>
      <MyIntegrationForm integration={integration} onIntegrationUpdate={onIntegrationUpdate} />
      <DangerZone onIntegrationDelete={onIntegrationDelete} />
    </Wrapper>
  );
};

export default MyIntegrationSettings;

const Wrapper = styled(Content)`
  display: flex;
  flex-direction: column;
  height: calc(100% - 32px);
`;
