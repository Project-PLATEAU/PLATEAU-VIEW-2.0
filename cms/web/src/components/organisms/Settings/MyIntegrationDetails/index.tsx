import { useCallback } from "react";
import { useNavigate, useParams } from "react-router-dom";

import MyIntegrationContent from "@reearth-cms/components/molecules/MyIntegrations/Content";

import useHooks from "./hooks";

const MyIntegrationDetails: React.FC = () => {
  const { workspaceId, integrationId } = useParams();
  const navigate = useNavigate();

  const handleIntegrationHeaderBack = useCallback(() => {
    navigate(`/workspace/${workspaceId}/myIntegrations`);
  }, [navigate, workspaceId]);

  const {
    selectedIntegration,
    webhookInitialValues,
    handleIntegrationUpdate,
    handleIntegrationDelete,
    handleWebhookCreate,
    handleWebhookDelete,
    handleWebhookUpdate,
    handleWebhookSelect,
  } = useHooks({
    integrationId,
  });

  return selectedIntegration ? (
    <MyIntegrationContent
      integration={selectedIntegration}
      webhookInitialValues={webhookInitialValues}
      onIntegrationUpdate={handleIntegrationUpdate}
      onIntegrationDelete={handleIntegrationDelete}
      onWebhookCreate={handleWebhookCreate}
      onWebhookDelete={handleWebhookDelete}
      onWebhookUpdate={handleWebhookUpdate}
      onIntegrationHeaderBack={handleIntegrationHeaderBack}
      onWebhookSelect={handleWebhookSelect}
      // onWebhookFormHeaderBack={handleWebhookFormHeaderBack}
      // onWebhookFormNavigation={handleWebhookFormNavigation}
      // onWebhookEditNavigation={handleWebhookEditNavigation}
      // onTabChange={handleTabChange}
      // activeTab={tab}
      // showWebhookForm={showWebhookForm}
    />
  ) : null;
};

export default MyIntegrationDetails;
