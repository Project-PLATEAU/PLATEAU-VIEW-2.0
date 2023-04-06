import { ComponentProps } from "react";

import { FieldTitle, FieldValue, FieldWrapper, TextInput } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
} & ComponentProps<typeof TextInput>;

const TextField: React.FC<Props> = ({ title, titleWidth, ...props }) => {
  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue>
        <TextInput {...props} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default TextField;
