import AccessibilityMolecule from "@reearth-cms/components/molecules/Accessibility";

import useHooks from "./hooks";

const Accessibility: React.FC = () => {
  const { projectScope, models, handlePublicUpdate, alias, assetPublic } = useHooks();
  return (
    <AccessibilityMolecule
      assetPublic={assetPublic}
      projectScope={projectScope}
      models={models}
      alias={alias}
      onPublicUpdate={handlePublicUpdate}
    />
  );
};

export default Accessibility;
