import { gql } from "@apollo/client";

import { requestFragment } from "@reearth-cms/gql/fragments";

export const GET_REQUESTS = gql`
  query GetRequests(
    $projectId: ID!
    $key: String
    $state: [RequestState!]
    $pagination: Pagination
    $createdBy: ID
    $reviewer: ID
    $sort: Sort
  ) {
    requests(
      projectId: $projectId
      key: $key
      state: $state
      pagination: $pagination
      createdBy: $createdBy
      reviewer: $reviewer
      sort: $sort
    ) {
      nodes {
        id
        title
        description
        createdBy {
          id
          name
          email
        }
        workspaceId
        projectId
        threadId
        reviewersId
        reviewers {
          id
          name
          email
        }
        state
        createdAt
        updatedAt
        approvedAt
        closedAt
        thread {
          ...threadFragment
        }
      }
      totalCount
    }
  }

  ${requestFragment}
`;

export const GET_MODAL_REQUESTS = gql`
  query GetModalRequests(
    $projectId: ID!
    $key: String
    $state: [RequestState!]
    $pagination: Pagination
    $createdBy: ID
    $reviewer: ID
    $sort: Sort
  ) {
    requests(
      projectId: $projectId
      key: $key
      state: $state
      pagination: $pagination
      createdBy: $createdBy
      reviewer: $reviewer
      sort: $sort
    ) {
      nodes {
        id
        title
        description
        createdBy {
          name
        }
        items {
          itemId
        }
        reviewers {
          id
          name
        }
        state
        createdAt
      }
      totalCount
    }
  }

  ${requestFragment}
`;

export const GET_REQUEST = gql`
  query GetRequest($requestId: ID!) {
    node(id: $requestId, type: REQUEST) {
      id
      ... on Request {
        ...requestFragment
      }
    }
  }
`;

export const CREATE_REQUEST = gql`
  mutation CreateRequest(
    $projectId: ID!
    $title: String!
    $description: String
    $state: RequestState
    $reviewersId: [ID!]
    $items: [RequestItemInput!]!
  ) {
    createRequest(
      input: {
        projectId: $projectId
        title: $title
        description: $description
        state: $state
        reviewersId: $reviewersId
        items: $items
      }
    ) {
      request {
        ...requestFragment
      }
    }
  }
`;

export const UPDATE_REQUEST = gql`
  mutation UpdateRequest(
    $requestId: ID!
    $title: String
    $description: String
    $state: RequestState
    $reviewersId: [ID!]
    $items: [RequestItemInput!]
  ) {
    updateRequest(
      input: {
        requestId: $requestId
        title: $title
        description: $description
        state: $state
        reviewersId: $reviewersId
        items: $items
      }
    ) {
      request {
        ...requestFragment
      }
    }
  }
`;

export const APPROVE_REQUEST = gql`
  mutation ApproveRequest($requestId: ID!) {
    approveRequest(input: { requestId: $requestId }) {
      request {
        ...requestFragment
      }
    }
  }
`;

export const DELETE_REQUEST = gql`
  mutation DeleteRequest($projectId: ID!, $requestsId: [ID!]!) {
    deleteRequest(input: { projectId: $projectId, requestsId: $requestsId }) {
      requests
    }
  }
`;
