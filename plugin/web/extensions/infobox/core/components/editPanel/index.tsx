import type { Properties, InfoboxTemplate, Field } from "@web/extensions/infobox/types";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback } from "react";

import { EditorTab } from "../../hooks";
import PropertyBrowser from "../common/PropertyBrowser";

import FieldsEditor from "./FieldsEditor";

type Props = {
  properties?: Properties;
  fields: Field[];
  template: InfoboxTemplate;
  isSaving: boolean;
  editorTab: string;
  commonProperties: string[];
  attributesKey?: string;
  attributesName?: string;
  handleEditorTab: (tab: EditorTab) => void;
  onFieldCheckChange: (e: any) => void;
  onFieldTitleChange: (e: any) => void;
  onFieldMove: (dragIndex: number, hoverIndex: number) => void;
  saveTemplate: () => void;
};

const EditPanel: React.FC<Props> = ({
  template,
  fields,
  properties,
  isSaving,
  editorTab,
  commonProperties,
  attributesKey,
  attributesName,
  handleEditorTab,
  onFieldCheckChange,
  onFieldTitleChange,
  onFieldMove,
  saveTemplate,
}) => {
  const onTabView = useCallback(() => {
    handleEditorTab("view");
  }, [handleEditorTab]);
  const onTabEdit = useCallback(() => {
    handleEditorTab("edit");
  }, [handleEditorTab]);

  return (
    <>
      <Header>
        <Tab active={editorTab === "view"} onClick={onTabView}>
          {template.name}
        </Tab>
        <Tab active={editorTab === "edit"} onClick={onTabEdit}>
          <Icon icon="gearWheel" size={16} />
        </Tab>
      </Header>
      {editorTab === "view" && (
        <PropertyBrowser
          properties={properties}
          fields={fields}
          commonProperties={commonProperties}
          attributesKey={attributesKey}
          attributesName={attributesName}
        />
      )}
      {editorTab === "edit" && (
        <FieldsEditor
          fields={fields}
          isSaving={isSaving}
          saveTemplate={saveTemplate}
          onFieldCheckChange={onFieldCheckChange}
          onFieldTitleChange={onFieldTitleChange}
          onFieldMove={onFieldMove}
        />
      )}
    </>
  );
};

const Header = styled.div`
  display: flex;
  align-items: flex-end;
  justify-content: space-between;
  height: 48px;
  padding: 8px 8px 0 8px;
  background-color: #f5f5f5;
  border-bottom: 1px solid #f0f0f0;
`;

const Tab = styled.div<{ active: boolean }>`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 40px;
  padding: 0 16px;
  box-shadow: inset -1px 0px 0px #f0f0f0, inset 0px 1px 0px #f0f0f0, inset 1px 0px 0px #f0f0f0;
  border-radius: 2px 2px 0px 0px;
  background-color: ${({ active }) => (active ? "#fff" : "#fafafa")};
  color: ${({ active }) => (active ? "#1890FF" : "rgba(0, 0, 0, 0.85)")};
  border-bottom: 1px solid #f0f0f0;
  border-color: ${({ active }) => (active ? "#fff" : "#f0f0f0")};
  transform: translateY(1px);
  cursor: pointer;
`;

export default EditPanel;
