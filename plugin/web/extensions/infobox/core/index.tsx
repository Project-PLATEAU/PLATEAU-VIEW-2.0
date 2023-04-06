import { Spin } from "@web/sharedComponents";
import { styled } from "@web/theme";

import EditPanel from "./components/editPanel";
import ViewPanel from "./components/viewPanel";
import useHooks from "./hooks";

const Infobox: React.FC = () => {
  const {
    inEditor,
    dataState,
    properties,
    fields,
    template,
    wrapperRef,
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
  } = useHooks();

  return (
    <Wrapper ref={wrapperRef}>
      {dataState !== "empty" && (
        <ContentPanel ready={dataState === "ready"}>
          {template &&
            (inEditor ? (
              <EditPanel
                template={template}
                fields={fields}
                properties={properties}
                isSaving={isSaving}
                editorTab={editorTab}
                commonProperties={commonProperties}
                attributesKey={attributesKey}
                attributesName={attributesName}
                handleEditorTab={handleEditorTab}
                saveTemplate={saveTemplate}
                onFieldCheckChange={onFieldCheckChange}
                onFieldTitleChange={onFieldTitleChange}
                onFieldMove={onFieldMove}
              />
            ) : (
              <ViewPanel
                name={template.name}
                fields={fields}
                properties={properties}
                commonProperties={commonProperties}
                attributesKey={attributesKey}
                attributesName={attributesName}
              />
            ))}
        </ContentPanel>
      )}
      {dataState === "empty" && <SimplePane>NO DATA</SimplePane>}
      <Loading active={dataState === "loading"}>
        <Spin />
      </Loading>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  position: relative;
  padding: 0 12px;
`;

const ContentPanel = styled.div<{ ready: boolean }>`
  background: #f4f4f4;
  margin-bottom: 6px;
  box-shadow: 1px 2px 4px rgba(0, 0, 0, 0.25);
  border-radius: 4px !important;
  overflow: hidden;
  opacity: ${({ ready }) => (ready ? 1 : 0.2)};
  transition: all 0.25s ease;
`;

const Loading = styled.div<{ active: boolean }>`
  position: absolute;
  width: 100%;
  height: 200px;
  left: 0;
  top: 0;
  display: flex;
  align-items: center;
  justify-content: center;
  opacity: ${({ active }) => (active ? 1 : 0)};
  pointer-events: none;
  z-index: 10;
`;

const SimplePane = styled.div`
  width: 100%;
  padding: 12px 0;
  display: flex;
  align-items: center;
  justify-content: center;
  color: #666;
`;

export default Infobox;
