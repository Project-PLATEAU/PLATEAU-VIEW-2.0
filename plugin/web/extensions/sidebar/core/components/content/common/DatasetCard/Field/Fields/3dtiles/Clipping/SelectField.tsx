import { Select } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { ComponentProps } from "react";

import { ColumnFieldWrapper, Text, FieldValue } from "./commonStyles";

type Props = {
  title: string;
} & ComponentProps<typeof Select>;

const SelectField: React.FC<Props> = ({ title, ...props }) => {
  return (
    <ColumnFieldWrapper>
      <Text>{title}</Text>
      <FieldValue noBorder>
        <StyledSelect {...props} />
      </FieldValue>
    </ColumnFieldWrapper>
  );
};

export default SelectField;

const StyledSelect = styled(Select)`
  width: 100%;
`;
