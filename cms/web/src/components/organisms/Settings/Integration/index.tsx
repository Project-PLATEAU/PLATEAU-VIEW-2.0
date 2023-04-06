import { useParams } from "react-router-dom";

import IntegrationConnectModal from "@reearth-cms/components/molecules/Integration/IntegrationConnectModal";
import IntegrationSettingsModal from "@reearth-cms/components/molecules/Integration/IntegrationSettingsModal";
import IntegrationTable from "@reearth-cms/components/molecules/Integration/IntegrationTable";

import useHooks from "./hooks";

const Integration: React.FC = () => {
  const { workspaceId } = useParams();

  const {
    integrations,
    workspaceIntegrationMembers,
    handleIntegrationConnect,
    handleIntegrationConnectModalClose,
    handleIntegrationConnectModalOpen,
    integrationConnectModalShown,
    handleIntegrationSettingsModalClose,
    handleIntegrationSettingsModalOpen,
    integrationSettingsModalShown,
    handleUpdateIntegration,
    selectedIntegrationMember,
    selection,
    handleSearchTerm,
    setSelection,
    handleIntegrationRemove,
  } = useHooks(workspaceId);

  return (
    <>
      <IntegrationTable
        integrationMembers={workspaceIntegrationMembers}
        selection={selection}
        onSearchTerm={handleSearchTerm}
        onIntegrationSettingsModalOpen={handleIntegrationSettingsModalOpen}
        onIntegrationConnectModalOpen={handleIntegrationConnectModalOpen}
        setSelection={setSelection}
        onIntegrationRemove={handleIntegrationRemove}
      />
      <IntegrationConnectModal
        integrations={integrations}
        onSubmit={handleIntegrationConnect}
        open={integrationConnectModalShown}
        onClose={handleIntegrationConnectModalClose}
      />
      <IntegrationSettingsModal
        selectedIntegrationMember={selectedIntegrationMember}
        onSubmit={handleUpdateIntegration}
        open={integrationSettingsModalShown}
        onClose={handleIntegrationSettingsModalClose}
      />
    </>
  );
};

export default Integration;
