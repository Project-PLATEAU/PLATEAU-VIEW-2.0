import styled from "@emotion/styled";
import { CSSProperties } from "react";

type CustomTagProps = {
  value?: number | string;
  color?: CSSProperties["color"];
};

const CustomTag: React.FC<CustomTagProps> = ({ value, color }) => {
  return (
    <CustomTagWrapper color={color ?? "#bfbfbf"}>
      <span>{value ?? ""}</span>
    </CustomTagWrapper>
  );
};

const CustomTagWrapper = styled.div`
  padding: 0px 6px;
  width: 20px;
  height: 16px;
  background-color: ${props => props.color};
  color: #ffffff;
  border-radius: 100px;
  font-family: Roboto Mono;
  font-style: normal;
  font-weight: 400;
  font-size: 12px;
  line-height: 20px;
  display: flex;
  justify-content: center;
  align-items: center;
`;

export default CustomTag;
