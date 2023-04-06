export type IntegrationMember = {
  id: string;
  integration: Integration;
  integrationRole: Role;
  invitedById: string;
  active: boolean;
};

export type Role = "WRITER" | "READER" | "OWNER" | "MAINTAINER";

export type Integration = {
  id: string;
  name: string;
  description?: string | null;
  logoUrl: string;
  developerId: string;
  developer: Developer;
  iType: IntegrationType;
  config: {
    token?: string;
    webhooks?: Webhook[];
  };
};

export type Developer = {
  id: string;
  name: string;
  email: string;
};

export type IntegrationType = "Private" | "Public";

export type Webhook = {
  id: string;
  name: string;
  url: string;
  active: boolean;
  trigger: WebhookTrigger;
};

export type WebhookTrigger = {
  onItemCreate?: boolean | null;
  onItemUpdate?: boolean | null;
  onItemDelete?: boolean | null;
  onAssetUpload?: boolean | null;
  onAssetDelete?: boolean | null;
  onItemPublish?: boolean | null;
  onItemUnPublish?: boolean | null;
};
