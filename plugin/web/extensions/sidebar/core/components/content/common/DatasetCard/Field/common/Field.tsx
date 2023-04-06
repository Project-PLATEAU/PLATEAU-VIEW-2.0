import { FieldTitle, FieldValue, FieldWrapper } from "../commonComponents";

type Props = {
  title: string;
  titleWidth?: number;
  noBorder?: boolean;
  value: JSX.Element;
};

const Field: React.FC<Props> = ({ title, titleWidth, noBorder, value }) => {
  return (
    <FieldWrapper>
      <FieldTitle width={titleWidth}>{title}</FieldTitle>
      <FieldValue noBorder={noBorder}>{value}</FieldValue>
    </FieldWrapper>
  );
};

export default Field;
