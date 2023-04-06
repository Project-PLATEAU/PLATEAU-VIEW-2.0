import { useParams } from "react-router-dom";

import ModelsList from "@reearth-cms/components/molecules/Model/ModelsList/ModelsList";
import ModelFormModal from "@reearth-cms/components/molecules/Schema/ModelFormModal";

import useHooks from "./hooks";

export interface Props {
  className?: string;
  title: string;
  collapsed?: boolean;
  onModelSelect: (modelId: string) => void;
}

const ModelsMenu: React.FC<Props> = ({ className, title, collapsed, onModelSelect }) => {
  const { modelId } = useParams();

  const {
    model,
    models,
    modelModalShown,
    isKeyAvailable,
    handleModalOpen,
    handleModalClose,
    handleModelCreate,
    handleModelKeyCheck,
  } = useHooks({
    modelId,
  });

  return (
    <>
      <ModelsList
        className={className}
        title={title}
        selectedKey={model?.id}
        models={models}
        collapsed={collapsed}
        onModelSelect={onModelSelect}
        onModalOpen={handleModalOpen}
      />
      <ModelFormModal
        isKeyAvailable={isKeyAvailable}
        open={modelModalShown}
        onModelKeyCheck={handleModelKeyCheck}
        onClose={handleModalClose}
        onCreate={handleModelCreate}
      />
    </>
  );
};

export default ModelsMenu;
