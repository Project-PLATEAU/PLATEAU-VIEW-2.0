import { gql } from "@apollo/client";

import { threadFragment } from "@reearth-cms/gql/fragments";

export const GET_ITEMS = gql`
  query GetItems($schemaId: ID!, $pagination: Pagination) {
    items(schemaId: $schemaId, pagination: $pagination) {
      nodes {
        id
        schemaId
        createdAt
        updatedAt
        status
        user {
          name
        }
        integration {
          name
        }
        fields {
          schemaFieldId
          type
          value
        }
        thread {
          ...threadFragment
        }
      }
    }
  }

  ${threadFragment}
`;

export const GET_ITEM_NODE = gql`
  query GetItem($id: ID!) {
    node(id: $id, type: Item) {
      ... on Item {
        id
        schemaId
        createdAt
        updatedAt
        status
        assets {
          id
          url
        }
        user {
          name
        }
        integration {
          name
        }
        fields {
          schemaFieldId
          type
          value
        }
        thread {
          ...threadFragment
        }
      }
    }
  }
`;

export const SEARCH_ITEM = gql`
  query SearchItem($query: ItemQuery!, $sort: ItemSort, $pagination: Pagination) {
    searchItem(query: $query, sort: $sort, pagination: $pagination) {
      nodes {
        id
        schemaId
        createdAt
        updatedAt
        status
        assets {
          id
          url
        }
        user {
          name
        }
        integration {
          name
        }
        fields {
          schemaFieldId
          type
          value
        }
        thread {
          ...threadFragment
        }
      }
      totalCount
    }
  }

  ${threadFragment}
`;

export const CREATE_ITEM = gql`
  mutation CreateItem($modelId: ID!, $schemaId: ID!, $fields: [ItemFieldInput!]!) {
    createItem(input: { modelId: $modelId, schemaId: $schemaId, fields: $fields }) {
      item {
        id
        schemaId
        fields {
          value
          type
          schemaFieldId
        }
      }
    }
  }
`;

export const DELETE_ITEM = gql`
  mutation DeleteItem($itemId: ID!) {
    deleteItem(input: { itemId: $itemId }) {
      itemId
    }
  }
`;

export const UPDATE_ITEM = gql`
  mutation UpdateItem($itemId: ID!, $fields: [ItemFieldInput!]!) {
    updateItem(input: { itemId: $itemId, fields: $fields }) {
      item {
        id
        schemaId
        fields {
          value
          type
          schemaFieldId
        }
      }
    }
  }
`;

export const UNPUBLISH_ITEM = gql`
  mutation UnpublishItem($itemId: [ID!]!) {
    unpublishItem(input: { itemId: $itemId }) {
      items {
        id
      }
    }
  }
`;
