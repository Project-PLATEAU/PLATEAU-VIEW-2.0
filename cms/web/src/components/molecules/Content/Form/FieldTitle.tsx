import styled from "@emotion/styled";

import { useT } from "@reearth-cms/i18n";

export type Props = {
  title: string;
  isUnique: boolean;
};

const FieldTitle: React.FC<Props> = ({ title, isUnique }) => {
  const t = useT();

  return (
    <Title>
      {title}
      {isUnique ? <FieldUnique>({t("unique")})</FieldUnique> : ""}
    </Title>
  );
};

export default FieldTitle;

const Title = styled.p`
  color: #000000d9;
  font-weight: 400;
  margin: 0;
`;

const FieldUnique = styled.span`
  margin-left: 4px;
  color: rgba(0, 0, 0, 0.45);
  font-weight: 400;
`;
