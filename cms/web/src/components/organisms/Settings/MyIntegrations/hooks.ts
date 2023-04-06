import { useCallback, useMemo, useState } from "react";

import Notification from "@reearth-cms/components/atoms/Notification";
import {
  Integration,
  IntegrationType,
} from "@reearth-cms/components/molecules/MyIntegrations/types";
import { useCreateIntegrationMutation, useGetMeQuery } from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";

export default () => {
  const [integrationModalShown, setIntegrationModalShown] = useState(false);
  const t = useT();

  const { data } = useGetMeQuery();

  const [createNewIntegration] = useCreateIntegrationMutation({
    refetchQueries: ["GetMe"],
  });

  const integrations = useMemo(() => {
    return data?.me?.integrations
      ?.map<Integration | undefined>(integration =>
        integration
          ? {
              id: integration.id,
              name: integration.name,
              description: integration.description,
              logoUrl: integration.logoUrl,
              developerId: integration.developerId,
              iType: integration.iType,
              config: {
                token: integration.config?.token,
                webhooks: integration.config?.webhooks,
              },
            }
          : undefined,
      )
      .filter((integration): integration is Integration => !!integration);
  }, [data?.me?.integrations]);

  const handleIntegrationCreate = useCallback(
    async (data: { name: string; description: string; logoUrl: string; type: IntegrationType }) => {
      const integration = await createNewIntegration({
        variables: {
          name: data.name,
          description: data.description,
          logoUrl: data.logoUrl,
          type: data.type,
        },
      });
      if (integration.errors || !integration.data?.createIntegration) {
        Notification.error({ message: t("Failed to create integration.") });
        return;
      }
      Notification.success({ message: t("Successfully created integration!") });
      setIntegrationModalShown(false);
    },
    [createNewIntegration, t],
  );

  const handleIntegrationModalClose = useCallback(() => {
    setIntegrationModalShown(false);
  }, []);

  const handleIntegrationModalOpen = useCallback(() => setIntegrationModalShown(true), []);

  return {
    integrations,
    integrationModalShown,
    handleIntegrationCreate,
    handleIntegrationModalOpen,
    handleIntegrationModalClose,
  };
};
