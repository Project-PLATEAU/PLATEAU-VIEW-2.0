import { styled } from "@web/theme";

import useHooks from "./hooks";

type Props = {};

const SceneEditor: React.FC<Props> = () => {
  const { titleRef, descriptionRef, canSave, onTitleChange, onSave, onCancel } = useHooks();

  return (
    <Wrapper>
      <TitleInput placeholder="タイトル" ref={titleRef} onChange={onTitleChange} />
      <ContentInput placeholder="内容" ref={descriptionRef} />
      <Actions>
        <Button primary onClick={onSave} disabled={!canSave}>
          保存
        </Button>
        <Button onClick={onCancel}>キャンセル</Button>
      </Actions>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  position: relative;
  width: 100%;
  height: 100%;
  padding: 12px;
  display: flex;
  flex-direction: column;
  gap: 12px;

  .editor-toolbar {
    padding: 0 !important;
    border-color: #d9d9d9 !important;
    border-radius: 2px 2px 0 0 !important;
  }
  .editor-toolbar button {
    border: none !important;
    border-radius: 0 !important;
  }
  .editor-toolbar button.active,
  .editor-toolbar button:hover {
    background: #a7e3e3 !important;
  }
  .CodeMirror {
    border-color: #d9d9d9 !important;
    border-radius: 0 0 2px 2px !important;
  }
  .CodeMirror .cm-spell-error:not(.cm-url):not(.cm-comment):not(.cm-tag):not(.cm-word) {
    background: none !important;
  }
`;

const TitleInput = styled.input`
  display: block;
  width: 100%;
  height: 32px;
  padding: 4px 12px;
  background-color: #fff;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  font-size: 14px;
  line-height: 24px;
  outline: none;
  flex-shrink: 0;
`;

const ContentInput = styled.textarea`
  display: block;
  width: 100%;
  height: 100%;
  padding: 4px 12px;
  background-color: #fff;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  font-size: 14px;
  line-height: 22px;
  outline: none;
  resize: none;
`;

const Actions = styled.div`
  display: flex;
  flex-direction: row-reverse;
  gap: 12px;
`;

const Button = styled.div<{ primary?: boolean; disabled?: boolean }>`
  display: flex;
  align-items: center;
  justify-content: center;
  height: 29px;
  padding: 4px 12px;
  border-radius: 4px;
  background-color: ${({ primary, disabled }) =>
    disabled ? "#d1d1d1" : primary ? "#00bebe" : "#d1d1d1"};
  color: #fff;
  font-size: 14px;
  line-height: 21px;
  cursor: pointer;
  pointer-events: ${({ disabled }) => (disabled ? "none" : "all")};
`;

export default SceneEditor;
