import { gql } from "@apollo/client";

export const threadFragment = gql`
  fragment threadFragment on Thread {
    id
    workspaceId
    comments {
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
      authorId
      content
      createdAt
    }
  }
`;

export default threadFragment;
