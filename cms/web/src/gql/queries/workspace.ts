import { gql } from "@apollo/client";

import { workspaceFragment } from "@reearth-cms/gql/fragments";

export const GET_WORKSPACES = gql`
  query GetWorkspaces {
    me {
      id
      name
      myWorkspace {
        id
        ...WorkspaceFragment
      }
      workspaces {
        id
        ...WorkspaceFragment
      }
    }
  }

  ${workspaceFragment}
`;

export const UPDATE_WORKSPACE = gql`
  mutation UpdateWorkspace($workspaceId: ID!, $name: String!) {
    updateWorkspace(input: { workspaceId: $workspaceId, name: $name }) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const DELETE_WORKSPACE = gql`
  mutation DeleteWorkspace($workspaceId: ID!) {
    deleteWorkspace(input: { workspaceId: $workspaceId }) {
      workspaceId
    }
  }
`;

export const ADD_USERS_TO_WORKSPACE = gql`
  mutation AddUsersToWorkspace($workspaceId: ID!, $users: [MemberInput!]!) {
    addUsersToWorkspace(input: { workspaceId: $workspaceId, users: $users }) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const UPDATE_MEMBER_OF_WORKSPACE = gql`
  mutation UpdateMemberOfWorkspace($workspaceId: ID!, $userId: ID!, $role: Role!) {
    updateUserOfWorkspace(input: { workspaceId: $workspaceId, userId: $userId, role: $role }) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const REMOVE_MEMBER_FROM_WORKSPACE = gql`
  mutation RemoveMemberFromWorkspace($workspaceId: ID!, $userId: ID!) {
    removeUserFromWorkspace(input: { workspaceId: $workspaceId, userId: $userId }) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const ADD_INTEGRATION_TO_WORKSPACE = gql`
  mutation AddIntegrationToWorkspace($workspaceId: ID!, $integrationId: ID!, $role: Role!) {
    addIntegrationToWorkspace(
      input: { workspaceId: $workspaceId, integrationId: $integrationId, role: $role }
    ) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const UPDATE_INTEGRATION_OF_WORKSPACE = gql`
  mutation UpdateIntegrationOfWorkspace($workspaceId: ID!, $integrationId: ID!, $role: Role!) {
    updateIntegrationOfWorkspace(
      input: { workspaceId: $workspaceId, integrationId: $integrationId, role: $role }
    ) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const REMOVE_INTEGRATION_FROM_WORKSPACE = gql`
  mutation RemoveIntegrationFromWorkspace($workspaceId: ID!, $integrationId: ID!) {
    removeIntegrationFromWorkspace(
      input: { workspaceId: $workspaceId, integrationId: $integrationId }
    ) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;

export const CREATE_WORKSPACE = gql`
  mutation CreateWorkspace($name: String!) {
    createWorkspace(input: { name: $name }) {
      workspace {
        id
        ...WorkspaceFragment
      }
    }
  }
`;
