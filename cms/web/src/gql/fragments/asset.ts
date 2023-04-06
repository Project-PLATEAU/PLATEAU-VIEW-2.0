import { gql } from "@apollo/client";

import { assetFileFragment, threadFragment, integrationFragment } from "@reearth-cms/gql/fragments";

export const assetFragment = gql`
  fragment assetFragment on Asset {
    id
    projectId
    createdAt
    createdBy {
      ... on User {
        id
        name
        email
      }
      ... on Integration {
        ...integrationFragment
      }
    }
    size
    previewType
    uuid
    url
    thread {
      ...threadFragment
    }
    archiveExtractionStatus
  }
  ${assetFileFragment}
  ${integrationFragment}
  ${threadFragment}
`;

export default assetFragment;
