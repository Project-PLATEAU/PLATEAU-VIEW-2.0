import styled from "@emotion/styled";
import { CSSProperties, ReactNode } from "react";

type Props = {
  title: string | JSX.Element;
  toolbar?: ReactNode;
  children?: ReactNode;
  style?: CSSProperties;
};

const Card: React.FC<Props> = ({ title, toolbar, children, style }) => {
  return (
    <CardWrapper style={style}>
      <CardHeader>
        <Title>{title}</Title>
        <Toolbar>{toolbar}</Toolbar>
      </CardHeader>
      <CardBody>{children}</CardBody>
    </CardWrapper>
  );
};

const CardWrapper = styled.div`
  padding: 0;
  border: 1px solid #f5f5f5;
  margin-bottom: 24px;
`;

const CardHeader = styled.div`
  height: 64px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 12px;
  font-size: 24px;
`;

const Title = styled.p`
  margin: 0;
  font-size: 14px;
  font-weight: 500;
`;

const Toolbar = styled.div`
  display: flex;
  align-items: center;
`;

const CardBody = styled.div`
  padding: 10px;
  background-color: #f5f5f5;
  text-align: center;
`;

export default Card;
