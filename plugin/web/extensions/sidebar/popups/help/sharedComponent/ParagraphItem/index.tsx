import { styled } from "@web/theme";
import { ReactNode } from "react";

const ParagraphItem: React.FC<{ children?: ReactNode }> = ({ children }) => {
  return <Wrapper>{children}</Wrapper>;
};

export default ParagraphItem;

const Wrapper = styled.div`
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  padding: 0px;
  gap: 8px;
  width: 301px;
`;
