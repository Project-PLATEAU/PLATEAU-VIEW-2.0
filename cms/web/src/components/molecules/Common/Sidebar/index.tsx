import styled from "@emotion/styled";

import Sider, { SiderProps } from "@reearth-cms/components/atoms/Sider";

const Sidebar: React.FC<SiderProps> = ({ collapsible = true, ...siderProps }) => {
  return <StyledSidebar collapsible={collapsible} {...siderProps} />;
};

export default Sidebar;

const StyledSidebar = styled(Sider)`
  background-color: #fff;

  .ant-layout-sider-trigger {
    background-color: #fff;
    border-top: 1px solid #f0f0f0;
    border-right: 1px solid #f0f0f0;
    color: #002140;
    text-align: left;
    padding: 0 20px;
    margin: 0;
    height: 38px;
    line-height: 38px;
    cursor: pointer;
  }
  .ant-layout-sider-children {
    height: calc(100% + 12px);
  }
  .ant-menu-inline {
    border-right: 1px solid white;

    & > li {
      padding: 0 20px;
    }
  }
  .ant-menu-vertical {
    border-right: none;
  }
`;
