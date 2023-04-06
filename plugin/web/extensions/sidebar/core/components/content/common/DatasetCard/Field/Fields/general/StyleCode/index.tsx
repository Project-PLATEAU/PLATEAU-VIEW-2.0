import { styled } from "@web/theme";

import { BaseFieldProps } from "../../types";

import useHooks from "./hooks";

const StyleCode: React.FC<BaseFieldProps<"styleCode">> = ({ value, editMode, onUpdate }) => {
  const { code, onEdit, onApply } = useHooks({ value, onUpdate });

  return editMode ? (
    <Wrapper>
      <CodeEditor value={code} onChange={onEdit} />
      <Button onClick={onApply}>Apply</Button>
    </Wrapper>
  ) : null;
};

export default StyleCode;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-end;
  gap: 8px;
`;

const CodeEditor = styled.textarea`
  height: 144px;
  width: 100%;
  padding: 12px;
  border: none;
  overflow: auto;
  background: #f3f3f3;
  outline: none;
  resize: none;
  :focus {
    border: none;
  }
`;

const Button = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  padding: 0 16px;
  height: 32px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  box-shadow: 0px 2px 0px rgba(0, 0, 0, 0.016);
  cursor: pointer;
`;
