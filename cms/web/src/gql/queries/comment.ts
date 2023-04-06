import { gql } from "@apollo/client";

export const ADD_COMMENT = gql`
  mutation AddComment($threadId: ID!, $content: String!) {
    addComment(input: { threadId: $threadId, content: $content }) {
      comment {
        id
        author {
          ... on User {
            id
            name
            email
          }
          ... on Integration {
            id
            name
          }
        }
        authorType
        authorId
        content
        createdAt
      }
    }
  }
`;

export const UPDATE_COMMENT = gql`
  mutation UpdateComment($commentId: ID!, $threadId: ID!, $content: String!) {
    updateComment(input: { commentId: $commentId, threadId: $threadId, content: $content }) {
      comment {
        id
        author {
          ... on User {
            id
            name
            email
          }
          ... on Integration {
            id
            name
          }
        }
        authorType
        authorId
        content
        createdAt
      }
    }
  }
`;

export const DELETE_COMMENT = gql`
  mutation DeleteComment($commentId: ID!, $threadId: ID!) {
    deleteComment(input: { commentId: $commentId, threadId: $threadId }) {
      commentId
    }
  }
`;
