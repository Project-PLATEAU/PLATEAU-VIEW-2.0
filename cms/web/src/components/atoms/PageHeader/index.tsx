import styled from "@emotion/styled";
import { PageHeader, PageHeaderProps } from "antd";

export type Props = PageHeaderProps;

const Header: React.FC<Props> = props => {
  return <StyledPageHeader {...props} />;
};

const StyledPageHeader = styled(PageHeader)`
  background-color: #fff;
`;

export default Header;
