import { gql } from "@apollo/client";

export const assetFileFragment = gql`
  fragment assetFileFragment on AssetFile {
    name
    size
    contentType
    path
  }
`;

export const assetFile5Fragment = gql`
  fragment assetFile5Fragment on AssetFile {
    ...assetFileFragment
    children {
      ...assetFileFragment
      children {
        ...assetFileFragment
        children {
          ...assetFileFragment
          children {
            ...assetFileFragment
            children {
              ...assetFileFragment
            }
          }
        }
      }
    }
  }

  ${assetFileFragment}
`;

export default assetFileFragment;
