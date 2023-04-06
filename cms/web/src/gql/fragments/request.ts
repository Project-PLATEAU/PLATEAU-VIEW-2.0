import { gql } from "@apollo/client";

import { threadFragment } from "@reearth-cms/gql/fragments";

export const requestFragment = gql`
  fragment requestFragment on Request {
    id
    items {
      itemId
      version
      ref
      item {
        version
        parents
        refs
        value {
          id
          schemaId
          modelId
          model {
            name
          }
          fields {
            schemaFieldId
            type
            value
          }
          schema {
            id
            fields {
              id
              type
              title
              key
              description
              required
              unique
              multiple
              typeProperty {
                ... on SchemaFieldText {
                  defaultValue
                  maxLength
                }
                ... on SchemaFieldTextArea {
                  defaultValue
                  maxLength
                }
                ... on SchemaFieldMarkdown {
                  defaultValue
                  maxLength
                }
                ... on SchemaFieldAsset {
                  assetDefaultValue: defaultValue
                }
                ... on SchemaFieldSelect {
                  selectDefaultValue: defaultValue
                  values
                }
                ... on SchemaFieldInteger {
                  integerDefaultValue: defaultValue
                  min
                  max
                }
                ... on SchemaFieldBool {
                  defaultValue
                }
                ... on SchemaFieldURL {
                  defaultValue
                }
              }
            }
          }
        }
      }
    }
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
    state
    createdAt
    updatedAt
    approvedAt
    closedAt
    thread {
      ...threadFragment
    }
    project {
      id
      name
      createdAt
      updatedAt
    }
    reviewers {
      id
      name
      email
    }
  }
  ${threadFragment}
`;

export default requestFragment;
