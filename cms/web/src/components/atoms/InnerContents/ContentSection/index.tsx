import styled from "@emotion/styled";
import { ReactNode } from "react";

export type Props = {
  className?: string;
  title?: string;
  danger?: boolean;
  headerActions?: ReactNode;
  children?: ReactNode;
};

const ContentSection: React.FC<Props> = ({ title, headerActions, children, danger }) => {
  return (
    <Wrapper danger={danger}>
      {title && (
        <Header>
          <Title>{title}</Title>
          {headerActions}
        </Header>
      )}
      <GridArea>{children}</GridArea>
    </Wrapper>
  );
};

export default ContentSection;

const Wrapper = styled.div<{ danger?: boolean }>`
  display: flex;
  flex-direction: column;
  height: 100%;
  background-color: #fff;
  color: rgba(0, 0, 0, 0.85);
  ${({ danger }) => danger && "border: 1px solid #FF4D4F;"}
`;

const Header = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: space-between;
  border-bottom: 1px solid rgba(0, 0, 0, 0.03);
  align-items: center;
  padding: 10px 24px;
`;

const Title = styled.p`
  font-size: 16px;
  font-weight: 500;
  margin: 0 8px 0 0;
`;

const GridArea = styled.div`
  padding: 24px;
`;
