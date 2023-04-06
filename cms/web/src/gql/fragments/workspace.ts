import { gql } from "@apollo/client";

import { integrationFragment } from "@reearth-cms/gql/fragments";

export const workspaceFragment = gql`
  fragment WorkspaceFragment on Workspace {
    id
    name
    members {
      ... on WorkspaceUserMember {
        user {
          id
          name
          email
        }
        userId
        role
      }
      ... on WorkspaceIntegrationMember {
        integration {
          ...integrationFragment
        }
        integrationRole: role
        active
        invitedBy {
          id
          name
          email
        }
        invitedById
      }
    }
    personal
  }
  ${integrationFragment}
`;

export default workspaceFragment;
