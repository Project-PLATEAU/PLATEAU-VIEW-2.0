import styled from "@emotion/styled";
import { ReactNode } from "react";

export type Props = {
  left?: ReactNode;
  center?: ReactNode;
  right?: ReactNode;
};

const ComplexInnerContents: React.FC<Props> = ({ left, center, right }) => {
  return (
    <PaddedContent>
      <Main>
        {left}
        <Center>{center}</Center>
      </Main>
      <Right>{right}</Right>
    </PaddedContent>
  );
};

export default ComplexInnerContents;

const PaddedContent = styled.div`
  display: flex;
  margin: 16px 0 0 16px;
  height: calc(100vh - 66px);
`;

const Main = styled.div`
  display: flex;
  margin-right: 10px;
  flex: 1;
`;

const Center = styled.div`
  max-height: 100%;
  display: flex;
  flex: 1;
`;

const Right = styled.div`
  // overflow-y: auto;
  display: flex;
`;
