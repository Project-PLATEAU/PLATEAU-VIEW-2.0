import { Maybe, Model } from "@reearth-cms/gql/graphql-client-api";

export const fromGraphQLModel = (model: Maybe<Model>) => {
  if (!model) return;
  return {
    id: model.id,
    description: model.description,
    name: model.name,
    key: model.key,
    public: model.public,
    schema: {
      id: model.schema?.id,
      fields: model.schema?.fields.map(field => ({
        id: field.id,
        description: field.description,
        title: field.title,
        type: field.type,
        key: field.key,
        unique: field.unique,
        multiple: field.multiple,
        required: field.required,
        typeProperty: field.typeProperty,
      })),
    },
  };
};
