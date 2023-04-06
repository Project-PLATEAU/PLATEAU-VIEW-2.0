import { Switch } from "@web/sharedComponents";
import { ComponentProps } from "react";

import { FieldTitle, FieldValue, FieldWrapper } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
} & ComponentProps<typeof Switch>;

const SwitchField: React.FC<Props> = ({ title, titleWidth, ...props }) => {
  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue noBorder>
        <Switch {...props} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default SwitchField;
