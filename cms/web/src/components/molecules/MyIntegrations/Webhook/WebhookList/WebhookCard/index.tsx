import styled from "@emotion/styled";
import { Switch } from "antd";
import { useCallback } from "react";

import Card from "@reearth-cms/components/atoms/Card";
import Icon from "@reearth-cms/components/atoms/Icon";
import { Webhook, WebhookTrigger } from "@reearth-cms/components/molecules/MyIntegrations/types";

export type Props = {
  webhook: Webhook;
  onWebhookDelete: (webhookId: string) => Promise<void>;
  onWebhookUpdate: (data: {
    webhookId: string;
    name: string;
    url: string;
    active: boolean;
    trigger: WebhookTrigger;
  }) => Promise<void>;
  onWebhookSettings: (webhookId: string) => void;
};

const WebhookCard: React.FC<Props> = ({
  webhook,
  onWebhookDelete,
  onWebhookUpdate,
  onWebhookSettings,
}) => {
  const handleWebhookDelete = useCallback(() => {
    onWebhookDelete(webhook.id);
  }, [onWebhookDelete, webhook.id]);

  const handleWebhookUpdate = useCallback(
    (active: boolean) => {
      onWebhookUpdate({ ...webhook, active, webhookId: webhook.id });
    },
    [onWebhookUpdate, webhook],
  );

  return (
    <Card
      style={{ marginTop: 16 }}
      title={
        <>
          <WebhookTitle>{webhook.name}</WebhookTitle>
          <Switch
            checkedChildren="ON"
            unCheckedChildren="OFF"
            checked={webhook.active}
            onChange={handleWebhookUpdate}
          />
        </>
      }
      extra={
        <>
          <Icon icon="settings" size={16} onClick={() => onWebhookSettings(webhook.id)} />
          <Icon icon="delete" size={16} onClick={handleWebhookDelete} style={{ marginLeft: 12 }} />
        </>
      }>
      {webhook.url}
    </Card>
  );
};

const WebhookTitle = styled.span`
  display: inline-block;
  margin-right: 8px;
`;

export default WebhookCard;
