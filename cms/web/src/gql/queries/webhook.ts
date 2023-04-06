import { gql } from "@apollo/client";

export const CREATE_WEBHOOK = gql`
  mutation CreateWebhook(
    $integrationId: ID!
    $name: String!
    $url: URL!
    $active: Boolean!
    $trigger: WebhookTriggerInput!
    $secret: String!
  ) {
    createWebhook(
      input: {
        integrationId: $integrationId
        name: $name
        url: $url
        active: $active
        trigger: $trigger
        secret: $secret
      }
    ) {
      webhook {
        id
        name
        url
        active
        trigger {
          onItemCreate
          onItemUpdate
          onItemDelete
          onItemPublish
          onItemUnPublish
          onAssetUpload
          onAssetDecompress
          onAssetDelete
        }
        secret
        createdAt
        updatedAt
      }
    }
  }
`;

export const UPDATE_WEBHOOK = gql`
  mutation UpdateWebhook(
    $integrationId: ID!
    $webhookId: ID!
    $name: String!
    $url: URL!
    $active: Boolean!
    $trigger: WebhookTriggerInput!
    $secret: String
  ) {
    updateWebhook(
      input: {
        integrationId: $integrationId
        webhookId: $webhookId
        name: $name
        url: $url
        active: $active
        trigger: $trigger
        secret: $secret
      }
    ) {
      webhook {
        id
        name
        url
        active
        trigger {
          onItemCreate
          onItemUpdate
          onItemDelete
          onItemPublish
          onItemUnPublish
          onAssetUpload
          onAssetDecompress
          onAssetDelete
        }
        secret
        createdAt
        updatedAt
      }
    }
  }
`;

export const DELETE_WEBHOOK = gql`
  mutation DeleteWebhook($integrationId: ID!, $webhookId: ID!) {
    deleteWebhook(input: { integrationId: $integrationId, webhookId: $webhookId }) {
      webhookId
    }
  }
`;
