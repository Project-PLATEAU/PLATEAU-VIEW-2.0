import { gql } from "@apollo/client";

export const GET_PROJECT = gql`
  query GetProject($projectId: ID!) {
    node(id: $projectId, type: PROJECT) {
      id
      ... on Project {
        name
        description
        alias
        publication {
          scope
          assetPublic
        }
      }
    }
  }
`;

export const GET_PROJECTS = gql`
  query GetProjects($workspaceId: ID!, $pagination: Pagination) {
    projects(workspaceId: $workspaceId, pagination: $pagination) {
      nodes {
        id
        name
        description
        alias
        publication {
          scope
          assetPublic
        }
      }
    }
  }
`;

export const CHECK_PROJECT_ALIAS = gql`
  query CheckProjectAlias($alias: String!) {
    checkProjectAlias(alias: $alias) {
      alias
      available
    }
  }
`;

export const CREATE_PROJECT = gql`
  mutation CreateProject(
    $workspaceId: ID!
    $name: String!
    $description: String!
    $alias: String!
  ) {
    createProject(
      input: { workspaceId: $workspaceId, name: $name, description: $description, alias: $alias }
    ) {
      project {
        id
        name
        description
        alias
        publication {
          scope
          assetPublic
        }
      }
    }
  }
`;

export const DELETE_PROJECT = gql`
  mutation DeleteProject($projectId: ID!) {
    deleteProject(input: { projectId: $projectId }) {
      projectId
    }
  }
`;

export const UPDATE_PROJECT = gql`
  mutation UpdateProject(
    $projectId: ID!
    $name: String
    $description: String
    $alias: String
    $publication: UpdateProjectPublicationInput
  ) {
    updateProject(
      input: {
        projectId: $projectId
        name: $name
        description: $description
        alias: $alias
        publication: $publication
      }
    ) {
      project {
        id
        name
        description
        alias
        publication {
          scope
          assetPublic
        }
      }
    }
  }
`;
