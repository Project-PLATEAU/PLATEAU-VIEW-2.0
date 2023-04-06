import { useCallback, useMemo } from "react";

import Notification from "@reearth-cms/components/atoms/Notification";
import { PublicScope, Model } from "@reearth-cms/components/molecules/Accessibility";
import {
  useUpdateModelMutation,
  useGetModelsQuery,
  Model as GQLModel,
  ProjectPublicationScope,
  useUpdateProjectMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useProject } from "@reearth-cms/state";
import { fromGraphQLModel } from "@reearth-cms/utils/values";

export default () => {
  const t = useT();
  const [currentProject] = useProject();

  const { data: modelsData } = useGetModelsQuery({
    variables: {
      projectId: currentProject?.id ?? "",
      pagination: { first: 100 },
    },
    skip: !currentProject?.id,
  });

  const models = useMemo(() => {
    return modelsData?.models.nodes
      ?.map<Model | undefined>(model => fromGraphQLModel(model as GQLModel))
      .filter((model): model is Model => !!model);
  }, [modelsData?.models.nodes]);

  const [updateProjectMutation] = useUpdateProjectMutation();
  const [updateModelMutation] = useUpdateModelMutation();

  const handlePublicUpdate = useCallback(
    async (
      alias?: string,
      scope?: PublicScope,
      modelsToUpdate?: Model[],
      assetPublic?: boolean,
    ) => {
      if (!currentProject?.id) return;
      let errors = false;

      if ((scope && scope !== currentProject.scope) || alias) {
        const gqlScope =
          scope === "public" ? ProjectPublicationScope.Public : ProjectPublicationScope.Private;
        const projRes = await updateProjectMutation({
          variables: {
            alias: alias,
            projectId: currentProject.id,
            publication: { scope: gqlScope, assetPublic },
          },
        });
        if (projRes.errors) {
          errors = true;
        }
      }

      if (modelsToUpdate) {
        modelsToUpdate.forEach(async model => {
          const modelRes = await updateModelMutation({
            variables: { modelId: model.id, public: model.public },
          });
          if (modelRes.errors) {
            errors = true;
          }
        });
      }
      if (errors) {
        Notification.error({ message: t("Failed to update publication settings.") });
      } else {
        Notification.success({
          message: t("Successfully updated publication settings!"),
        });
      }
    },
    [currentProject, t, updateProjectMutation, updateModelMutation],
  );

  return {
    projectScope: currentProject?.scope,
    assetPublic: currentProject?.assetPublic,
    models,
    alias: currentProject?.alias,
    handlePublicUpdate,
  };
};
