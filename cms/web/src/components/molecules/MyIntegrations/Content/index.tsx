import styled from "@emotion/styled";

import PageHeader from "@reearth-cms/components/atoms/PageHeader";
import Tabs from "@reearth-cms/components/atoms/Tabs";
import MyIntegrationSettings from "@reearth-cms/components/molecules/MyIntegrations/Settings";
import {
  Integration,
  WebhookTrigger,
} from "@reearth-cms/components/molecules/MyIntegrations/types";
import Webhook from "@reearth-cms/components/molecules/MyIntegrations/Webhook";

export type Props = {
  integration: Integration;
  webhookInitialValues?: any;
  // activeTab?: string;
  // showWebhookForm: boolean;
  onIntegrationUpdate: (data: { name: string; description: string; logoUrl: string }) => void;
  onWebhookCreate: (data: {
    name: string;
    url: string;
    active: boolean;
    trigger: WebhookTrigger;
    secret: string;
  }) => Promise<void>;
  onWebhookDelete: (webhookId: string) => Promise<void>;
  onWebhookUpdate: (data: {
    webhookId: string;
    name: string;
    url: string;
    active: boolean;
    trigger: WebhookTrigger;
    secret?: string;
  }) => Promise<void>;
  onIntegrationHeaderBack: () => void;
  onWebhookSelect: (id: string) => void;
  onIntegrationDelete: () => Promise<void>;
  // onWebhookFormHeaderBack: () => void;
  // onWebhookFormNavigation: () => void;
  // onWebhookEditNavigation: (webhookId: string) => void;
};

const MyIntegrationContent: React.FC<Props> = ({
  integration,
  webhookInitialValues,
  // activeTab,
  // showWebhookForm,
  onIntegrationUpdate,
  onWebhookCreate,
  onWebhookDelete,
  onWebhookUpdate,
  onIntegrationHeaderBack,
  onIntegrationDelete,
  onWebhookSelect,
  // onWebhookFormHeaderBack,
  // onWebhookFormNavigation,
  // onWebhookEditNavigation,
}) => {
  const { TabPane } = Tabs;

  return (
    <MyIntegrationWrapper>
      <PageHeader title={integration.name} onBack={onIntegrationHeaderBack} />
      <MyIntegrationTabs
        defaultActiveKey="integration"
        // activeKey={activeTab}
        // onChange={onTabChange}
      >
        <TabPane tab="General" key="integration">
          <MyIntegrationSettings
            integration={integration}
            onIntegrationUpdate={onIntegrationUpdate}
            onIntegrationDelete={onIntegrationDelete}
          />
        </TabPane>
        <TabPane tab="Webhook" key="webhooks">
          <Webhook
            integration={integration}
            webhookInitialValues={webhookInitialValues}
            onWebhookCreate={onWebhookCreate}
            onWebhookDelete={onWebhookDelete}
            onWebhookUpdate={onWebhookUpdate}
            onWebhookSelect={onWebhookSelect}
            // onWebhookFormHeaderBack={onWebhookFormHeaderBack}
            // onWebhookFormNavigation={onWebhookFormNavigation}
            // onWebhookEditNavigation={onWebhookEditNavigation}
          />
        </TabPane>
      </MyIntegrationTabs>
    </MyIntegrationWrapper>
  );
};

const MyIntegrationWrapper = styled.div`
  min-height: 100%;
  background-color: #fff;
`;

const MyIntegrationTabs = styled(Tabs)`
  padding: 0 24px;
`;

export default MyIntegrationContent;
