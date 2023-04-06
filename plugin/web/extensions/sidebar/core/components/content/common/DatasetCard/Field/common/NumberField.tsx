import { ComponentProps, useEffect, useState } from "react";

import { FieldTitle, FieldValue, FieldWrapper, NumberInput, TextInput } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
  value?: number;
  onChange?: (value: number) => void;
} & ComponentProps<typeof TextInput>;

const NumberField: React.FC<Props> = ({ title, titleWidth, value, onChange, ...props }) => {
  const [number, setNumber] = useState(value);

  useEffect(() => {
    setNumber(value);
  }, [value]);

  const handleChange = (n: number) => {
    setNumber(n);
    onChange?.(n);
  };

  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue>
        <NumberInput value={number} onChange={handleChange} {...props} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default NumberField;
