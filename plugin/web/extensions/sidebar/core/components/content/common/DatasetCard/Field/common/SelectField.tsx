import { Select } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { ComponentProps } from "react";

import { FieldTitle, FieldValue, FieldWrapper } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
} & ComponentProps<typeof Select>;

const SelectField: React.FC<Props> = ({ title, titleWidth, ...props }) => {
  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue noBorder>
        <StyledSelect {...props} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default SelectField;

const StyledSelect = styled(Select)`
  width: 100%;
`;
