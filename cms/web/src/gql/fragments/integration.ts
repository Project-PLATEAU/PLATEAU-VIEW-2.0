import { gql } from "@apollo/client";

export const integrationFragment = gql`
  fragment integrationFragment on Integration {
    id
    name
    description
    logoUrl
    iType
    developerId
    developer {
      id
      name
      email
    }
    config {
      token
      webhooks {
        id
        name
        url
        active
        secret
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
        createdAt
        updatedAt
      }
    }
    createdAt
    updatedAt
  }
`;

export default integrationFragment;
