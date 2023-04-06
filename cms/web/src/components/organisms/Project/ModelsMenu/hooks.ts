import { useCallback, useEffect, useMemo, useState } from "react";
import { useNavigate } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { Model } from "@reearth-cms/components/molecules/Schema/types";
import {
  useGetModelsQuery,
  useCreateModelMutation,
  useCheckModelKeyAvailabilityLazyQuery,
  Model as GQLModel,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useModel, useWorkspace, useProject } from "@reearth-cms/state";
import { fromGraphQLModel } from "@reearth-cms/utils/values";

type Params = {
  modelId?: string;
};

export default ({ modelId }: Params) => {
  const t = useT();
  const [currentModel, setCurrentModel] = useModel();
  const [currentWorkspace] = useWorkspace();
  const [currentProject] = useProject();
  const navigate = useNavigate();

  const projectId = useMemo(() => currentProject?.id, [currentProject]);
  const [modelModalShown, setModelModalShown] = useState(false);
  const [isKeyAvailable, setIsKeyAvailable] = useState(false);

  const [CheckModelKeyAvailability, { data: keyData }] = useCheckModelKeyAvailabilityLazyQuery({
    fetchPolicy: "no-cache",
  });

  const handleModelKeyCheck = useCallback(
    async (key: string, ignoredKey?: string) => {
      if (!projectId || !key) return false;
      if (ignoredKey && key === ignoredKey) return true;
      const response = await CheckModelKeyAvailability({ variables: { projectId, key } });
      return response.data ? response.data.checkModelKeyAvailability.available : false;
    },
    [projectId, CheckModelKeyAvailability],
  );

  useEffect(() => {
    setIsKeyAvailable(!!keyData?.checkModelKeyAvailability.available);
  }, [keyData?.checkModelKeyAvailability]);

  const { data } = useGetModelsQuery({
    variables: { projectId: projectId ?? "", pagination: { first: 100 } },
    skip: !projectId,
  });

  const models = useMemo(() => {
    return data?.models.nodes
      ?.map<Model | undefined>(model => (model ? fromGraphQLModel(model as GQLModel) : undefined))
      .filter((model): model is Model => !!model);
  }, [data?.models.nodes]);

  const rawModel = useMemo(
    () => data?.models.nodes?.find(node => node?.id === modelId),
    [data, modelId],
  );

  const model = useMemo<Model | undefined>(
    () => (rawModel?.id ? fromGraphQLModel(rawModel as GQLModel) : undefined),
    [rawModel],
  );

  useEffect(() => {
    setCurrentModel(model ?? undefined);
  }, [model, currentModel, modelId, setCurrentModel]);

  const [createNewModel] = useCreateModelMutation({
    refetchQueries: ["GetModels"],
  });

  const handleModelCreate = useCallback(
    async (data: { name: string; description: string; key: string }) => {
      if (!projectId) return;
      const model = await createNewModel({
        variables: {
          projectId,
          name: data.name,
          description: data.description,
          key: data.key,
        },
      });
      if (model.errors || !model.data?.createModel) {
        Notification.error({ message: t("Failed to create model.") });
        return;
      }
      Notification.success({ message: t("Successfully created model!") });
      setModelModalShown(false);
      navigate(
        `/workspace/${currentWorkspace?.id}/project/${projectId}/schema/${model.data?.createModel.model.id}`,
      );
    },
    [currentWorkspace?.id, projectId, createNewModel, navigate, t],
  );

  const handleModalClose = useCallback(() => setModelModalShown(false), []);

  const handleModalOpen = useCallback(() => setModelModalShown(true), []);

  return {
    model,
    models,
    modelModalShown,
    isKeyAvailable,
    handleModalOpen,
    handleModalClose,
    handleModelCreate,
    handleModelKeyCheck,
  };
};
