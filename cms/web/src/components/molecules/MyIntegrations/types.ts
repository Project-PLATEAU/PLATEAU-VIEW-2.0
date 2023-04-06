export type Integration = {
  id: string;
  name: string;
  description?: string | null;
  logoUrl: string;
  developerId: string;
  iType: IntegrationType;
  config: {
    token?: string;
    webhooks?: Webhook[];
  };
};

export enum IntegrationType {
  Private = "Private",
  Public = "Public",
}

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
  onItemPublish?: boolean | null;
  onItemUnPublish?: boolean | null;
  onAssetUpload?: boolean | null;
  onAssetDecompress?: boolean | null;
  onAssetDelete?: boolean | null;
};
