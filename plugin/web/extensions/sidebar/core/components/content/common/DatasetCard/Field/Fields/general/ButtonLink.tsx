import { styled } from "@web/theme";
import { useCallback, useState } from "react";

import { BaseFieldProps } from "../types";

const ButtonLink: React.FC<BaseFieldProps<"buttonLink">> = ({ value, editMode, onUpdate }) => {
  const [title, setTitle] = useState(value.title);
  const [link, setLink] = useState(value.link);

  const handleChangeButtonTitle = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setTitle(e.currentTarget.value);
      onUpdate({
        ...value,
        title: e.currentTarget.value,
      });
    },
    [value, onUpdate],
  );

  const handleChangeButtonLink = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setLink(e.currentTarget.value);
      onUpdate({
        ...value,
        link: link,
      });
    },
    [onUpdate, value, link],
  );

  const handleButtonClick = useCallback(() => {
    if (!link) return;
    const prefix = "https://";
    let url = link;
    if (!url.match(/^[a-zA-Z]+:\/\//)) {
      url = prefix + url;
    }
    window.open(url, "_blank", "noopener");
  }, [link]);

  return editMode ? (
    <Wrapper>
      <Field>
        <FieldTitle>タイトル</FieldTitle>
        <FieldValue>
          <TextInput defaultValue={title} onChange={handleChangeButtonTitle} />
        </FieldValue>
      </Field>

      <Field>
        <FieldTitle>リンク</FieldTitle>
        <FieldValue>
          <TextInput defaultValue={link} onChange={handleChangeButtonLink} />
        </FieldValue>
      </Field>
    </Wrapper>
  ) : (
    <StyledButton onClick={handleButtonClick}>{title && <Text>{title}</Text>}</StyledButton>
  );
};

export default ButtonLink;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Text = styled.p`
  margin: 0;
  font-weight: 400;
  font-size: 14px;
  line-height: 22px;
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const TextInput = styled.input.attrs({ type: "text" })`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;

  :focus {
    border: none;
  }
`;

const StyledButton = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  background: #00bebe;
  color: #fff;
  border-radius: 2px;
  height: 24px;
  cursor: pointer;
`;
