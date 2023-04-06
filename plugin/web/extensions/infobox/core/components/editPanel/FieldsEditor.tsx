import type { Field } from "@web/extensions/infobox/types";
import { Button } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { DndProvider } from "react-dnd";
import { HTML5Backend } from "react-dnd-html5-backend";

import FieldItem from "./FieldItem";

type Props = {
  fields: Field[];
  isSaving: boolean;
  onFieldCheckChange: (e: any) => void;
  onFieldTitleChange: (e: any) => void;
  onFieldMove: (dragIndex: number, hoverIndex: number) => void;
  saveTemplate: () => void;
};

const FieldsEditor: React.FC<Props> = ({
  fields,
  isSaving,
  onFieldCheckChange,
  onFieldTitleChange,
  onFieldMove,
  saveTemplate,
}) => {
  return (
    <Wrapper>
      <FieldsHeader>
        <IconsWrapper />
        <ContentWrapper>
          <JsonPath>Key</JsonPath>
          <Title>
            Display Title
            <StyledButton size="small" onClick={saveTemplate} loading={isSaving}>
              保存
            </StyledButton>
          </Title>
        </ContentWrapper>
      </FieldsHeader>
      <FieldsWrapper>
        <DndProvider backend={HTML5Backend}>
          {fields.map((field, index) => (
            <FieldItem
              id={field.path}
              index={index}
              key={field.path}
              field={field}
              onCheckChange={onFieldCheckChange}
              onTitleChange={onFieldTitleChange}
              moveProperty={onFieldMove}
            />
          ))}
        </DndProvider>
      </FieldsWrapper>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  background-color: #fff;
`;

const FieldsWrapper = styled.div`
  padding: 0 12px 4px;
`;

const FieldsHeader = styled.div`
  position: relative;
  display: flex;
  align-items: center;
  height: 40px;
  padding: 4px 12px;
  gap: 12px;
  border-bottom: 1px solid #d9d9d9;
  font-size: 14px;
`;

const IconsWrapper = styled.div`
  width: 56px;
  flex-shrink: 0;
`;

const ContentWrapper = styled.div`
  width: 100%;
  display: flex;
  gap: 12px;
`;

const JsonPath = styled.div`
  width: 50%;
`;

const Title = styled.div`
  display: flex;
  justify-content: space-between;
  width: 50%;
`;

const StyledButton = styled(Button)`
  border-radius: 4px;
`;

export default FieldsEditor;
