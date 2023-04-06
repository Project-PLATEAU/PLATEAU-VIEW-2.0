import { gql } from "@apollo/client";

export const CREATE_INTEGRATION = gql`
  mutation CreateIntegration(
    $name: String!
    $description: String
    $logoUrl: URL!
    $type: IntegrationType!
  ) {
    createIntegration(
      input: { name: $name, description: $description, logoUrl: $logoUrl, type: $type }
    ) {
      integration {
        id
        name
        description
        logoUrl
        iType
      }
    }
  }
`;

export const UPDATE_INTEGRATION = gql`
  mutation UpdateIntegration(
    $integrationId: ID!
    $name: String!
    $description: String
    $logoUrl: URL!
  ) {
    updateIntegration(
      input: {
        integrationId: $integrationId
        name: $name
        description: $description
        logoUrl: $logoUrl
      }
    ) {
      integration {
        id
        name
        description
        logoUrl
        iType
      }
    }
  }
`;

export const DELETE_INTEGRATION = gql`
  mutation DeleteIntegration($integrationId: ID!) {
    deleteIntegration(input: { integrationId: $integrationId }) {
      integrationId
    }
  }
`;
