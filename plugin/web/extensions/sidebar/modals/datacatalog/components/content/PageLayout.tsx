import { styled } from "@web/theme";
import { ReactNode } from "react";

export type Props = {
  left: ReactNode;
  right: ReactNode;
};

const PageLayout: React.FC<Props> = ({ left, right }) => {
  return (
    <Body>
      <LeftPane>{left}</LeftPane>
      <Divider />
      <RightPane>{right}</RightPane>
    </Body>
  );
};

export default PageLayout;

const Body = styled.div`
  display: flex;
  height: 541px;
`;

const LeftPane = styled.div`
  width: 322px;
`;

const RightPane = styled.div`
  flex: 1;
  height: 100%;
  overflow: auto;
`;

const Divider = styled.div`
  border-right: 1px solid #c7c5c5;
  margin: 24px 10px;
`;
