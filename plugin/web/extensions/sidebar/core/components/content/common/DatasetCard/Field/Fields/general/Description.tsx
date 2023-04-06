import { Switch } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useMemo, useState } from "react";
import { Remarkable } from "remarkable";

import { BaseFieldProps } from "../types";

const Description: React.FC<BaseFieldProps<"description">> = ({ value, editMode, onUpdate }) => {
  const [isMarkdown, setIsMarkdown] = useState(!!value.isMarkdown);
  const [content, setContent] = useState(value.content ?? "");

  const description = useMemo(() => {
    const md = new Remarkable({
      html: false,
      breaks: true,
      typographer: true,
      linkTarget: "__blank",
    });
    return content ? (isMarkdown ? md.render(content) : content) : undefined;
  }, [content, isMarkdown]);

  const handleIsMarkdownChange = useCallback(() => {
    setIsMarkdown(!isMarkdown);
    onUpdate({
      ...value,
      isMarkdown: !isMarkdown,
    });
  }, [value, isMarkdown, onUpdate]);

  const handleContentChange = useCallback(
    (e: React.ChangeEvent<HTMLTextAreaElement>) => {
      setContent(e.currentTarget.value);
      onUpdate({
        ...value,
        content: e.currentTarget.value,
      });
    },
    [value, onUpdate],
  );

  return editMode ? (
    <div>
      <Text>内容</Text>
      <TextBox rows={4} defaultValue={content} onChange={handleContentChange} />
      <SwitchWrapper>
        <Switch checked={isMarkdown} size="small" onChange={handleIsMarkdownChange} />
        <Text>マークダウン</Text>
      </SwitchWrapper>
    </div>
  ) : isMarkdown && description ? (
    <Markdown dangerouslySetInnerHTML={{ __html: description }} />
  ) : (
    <div>{description}</div>
  );
};

export default Description;

const Text = styled.p`
  margin: 0;
`;

const TextBox = styled.textarea`
  width: 100%;
`;

const SwitchWrapper = styled.div`
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
`;

const Markdown = styled.div`
  img {
    max-width: 100%;
  }
`;
