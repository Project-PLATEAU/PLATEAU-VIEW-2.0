import IntegrationCreationModal from "@reearth-cms/components/molecules/MyIntegrations/CreationModal";
import MyIntegrationList from "@reearth-cms/components/molecules/MyIntegrations/List";

import useHooks from "./hooks";

const MyIntegrations: React.FC = () => {
  const {
    integrations,
    integrationModalShown,
    handleIntegrationCreate,
    handleIntegrationModalOpen,
    handleIntegrationModalClose,
  } = useHooks();

  return (
    <>
      <MyIntegrationList
        integrations={integrations}
        onIntegrationModalOpen={handleIntegrationModalOpen}
      />
      <IntegrationCreationModal
        open={integrationModalShown}
        onClose={handleIntegrationModalClose}
        onSubmit={handleIntegrationCreate}
      />
    </>
  );
};

export default MyIntegrations;
