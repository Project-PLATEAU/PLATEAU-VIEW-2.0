import styled from "@emotion/styled";

import Icon from "@reearth-cms/components/atoms/Icon";
import ComplexInnerContents from "@reearth-cms/components/atoms/InnerContents/complex";
import PageHeader from "@reearth-cms/components/atoms/PageHeader";
import Sidebar from "@reearth-cms/components/molecules/Common/Sidebar";
import FieldList from "@reearth-cms/components/molecules/Schema/FieldList";
import ModelFieldList from "@reearth-cms/components/molecules/Schema/ModelFieldList";
import { Field, FieldType, Model } from "@reearth-cms/components/molecules/Schema/types";

export type Props = {
  collapsed?: boolean;
  model?: Model;
  modelsMenu?: JSX.Element;
  onCollapse?: (collapse: boolean) => void;
  onFieldReorder: (data: Field[]) => Promise<void> | void;
  onFieldUpdateModalOpen: (field: Field) => void;
  onFieldCreationModalOpen: (fieldType: FieldType) => void;
  onFieldDelete: (fieldId: string) => Promise<void>;
};

const Schema: React.FC<Props> = ({
  collapsed,
  model,
  modelsMenu,
  onCollapse,
  onFieldReorder,
  onFieldUpdateModalOpen,
  onFieldCreationModalOpen,
  onFieldDelete,
}) => {
  return (
    <ComplexInnerContents
      left={
        <Sidebar
          collapsed={collapsed}
          onCollapse={onCollapse}
          collapsedWidth={54}
          width={208}
          trigger={<Icon icon={collapsed ? "panelToggleRight" : "panelToggleLeft"} />}>
          {modelsMenu}
        </Sidebar>
      }
      center={
        <Content>
          <PageHeader title={model?.name} subTitle={model?.key ? `#${model.key}` : null} />
          <ModelFieldListWrapper>
            <ModelFieldList
              fields={model?.schema.fields}
              handleFieldUpdateModalOpen={onFieldUpdateModalOpen}
              onFieldReorder={onFieldReorder}
              onFieldDelete={onFieldDelete}
            />
          </ModelFieldListWrapper>
        </Content>
      }
      right={
        <FieldListWrapper>
          <FieldList addField={onFieldCreationModalOpen} />
        </FieldListWrapper>
      }
    />
  );
};

export default Schema;

const Content = styled.div`
  width: 100%;
  height: 100%;
  overflow-y: auto;
  background: #fafafa;
`;

const ModelFieldListWrapper = styled.div`
  padding: 24px;
`;

const FieldListWrapper = styled.div`
  height: 100%;
  width: 272px;
  padding: 12px;
  overflow-y: auto;
`;
