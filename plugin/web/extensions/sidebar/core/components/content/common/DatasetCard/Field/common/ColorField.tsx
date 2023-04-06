import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { ChangeEvent, useEffect, useState } from "react";

import { FieldTitle, FieldValue, FieldWrapper, TextInput } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
  color?: string;
  onChange?: (color: string) => void;
};

function isValidColor(color: string) {
  return CSS.supports("color", color);
}

const ColorField: React.FC<Props> = ({ title, titleWidth, color, onChange }) => {
  const [text, setText] = useState(color);
  const [selectedColor, setSelectedColor] = useState(color);

  useEffect(() => {
    setText(color);
    setSelectedColor(color);
  }, [color]);

  const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
    setText(e.target.value);
    setSelectedColor(() => {
      const result = isValidColor(e.target.value) ? e.target.value : "";
      onChange?.(result);
      return result;
    });
  };

  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue>
        {selectedColor ? (
          <ColorBlock color={selectedColor} />
        ) : (
          <Icon icon="transparent" size={30} />
        )}
        <TextInput value={text} placeholder="#FFFFFF" onChange={handleChange} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default ColorField;

const ColorBlock = styled.div<{ color: string; legendStyle?: "circle" | "square" | "line" }>`
  width: 30px;
  height: ${({ legendStyle }) => (legendStyle === "line" ? "3px" : "30px")};
  background: ${({ color }) => color ?? "#d9d9d9"};
  border-radius: ${({ legendStyle }) =>
    legendStyle
      ? legendStyle === "circle"
        ? "50%"
        : legendStyle === "line"
        ? "5px"
        : "2px"
      : "1px 0 0 1px"};
`;
