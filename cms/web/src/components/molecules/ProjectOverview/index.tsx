import styled from "@emotion/styled";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import InnerContent from "@reearth-cms/components/atoms/InnerContents/basic";
import ContentSection from "@reearth-cms/components/atoms/InnerContents/ContentSection";
import { useT } from "@reearth-cms/i18n";

import ModelCard, { Model as ModelType } from "./ModelCard";

export type Model = ModelType;

export type Props = {
  projectName?: string;
  projectDescription?: string;
  models?: Model[];
  onModelModalOpen?: () => void;
  onSchemaNavigate?: (modelId: string) => void;
  onContentNavigate?: (modelId: string) => void;
  onModelDeletionModalOpen: (model: Model) => Promise<void>;
  onModelUpdateModalOpen: (model: Model) => Promise<void>;
};

const ProjectOverview: React.FC<Props> = ({
  projectName,
  projectDescription,
  models,
  onModelModalOpen,
  onSchemaNavigate,
  onContentNavigate,
  onModelDeletionModalOpen,
  onModelUpdateModalOpen,
}) => {
  const t = useT();

  return (
    <InnerContent title={projectName} subtitle={projectDescription} flexChildren>
      <ContentSection
        title={t("Models")}
        headerActions={
          <Button type="primary" icon={<Icon icon="plus" />} onClick={onModelModalOpen}>
            {t("New Model")}
          </Button>
        }>
        <GridArea>
          {models?.map(m => (
            <ModelCard
              key={m.id}
              model={m}
              onSchemaNavigate={onSchemaNavigate}
              onContentNavigate={onContentNavigate}
              onModelDeletionModalOpen={onModelDeletionModalOpen}
              onModelUpdateModalOpen={onModelUpdateModalOpen}
            />
          ))}
        </GridArea>
      </ContentSection>
    </InnerContent>
  );
};

export default ProjectOverview;

const GridArea = styled.div`
  display: grid;
  gap: 24px;
  grid-template-columns: repeat(auto-fill, minmax(240px, 1fr));
  justify-content: space-between;
`;
