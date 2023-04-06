import { gql } from "@apollo/client";

import { integrationFragment } from "@reearth-cms/gql/fragments";

export const GET_USER_BY_SEARCH = gql`
  query GetUserBySearch($nameOrEmail: String!) {
    searchUser(nameOrEmail: $nameOrEmail) {
      id
      name
      email
    }
  }
`;

export const GET_ME = gql`
  query GetMe {
    me {
      id
      name
      email
      lang
      myWorkspace {
        id
        name
      }
      workspaces {
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
            invitedById
          }
        }
      }
      auths
      integrations {
        ...integrationFragment
      }
    }
  }
  ${integrationFragment}
`;

export const GET_PROFILE = gql`
  query GetProfile {
    me {
      id
      name
      email
      lang
      theme
      myWorkspace {
        id
        name
      }
      auths
    }
  }
`;

export const GET_LANGUAGE = gql`
  query GetLanguage {
    me {
      id
      lang
    }
  }
`;

export const GET_THEME = gql`
  query GetTheme {
    me {
      id
      theme
    }
  }
`;

export const UPDATE_ME = gql`
  mutation UpdateMe(
    $name: String
    $email: String
    $lang: Lang
    $theme: Theme
    $password: String
    $passwordConfirmation: String
  ) {
    updateMe(
      input: {
        name: $name
        email: $email
        lang: $lang
        theme: $theme
        password: $password
        passwordConfirmation: $passwordConfirmation
      }
    ) {
      me {
        id
        name
        email
        lang
        theme
        myWorkspace {
          id
          name
        }
      }
    }
  }
`;

export const DELETE_ME = gql`
  mutation DeleteMe($userId: ID!) {
    deleteMe(input: { userId: $userId }) {
      userId
    }
  }
`;
