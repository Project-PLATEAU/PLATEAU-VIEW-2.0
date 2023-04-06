import styled from "@emotion/styled";

import Content from "@reearth-cms/components/atoms/Content";
import Layout from "@reearth-cms/components/atoms/Layout";
import Sider from "@reearth-cms/components/atoms/Sider";

export type InnerProps = {
  onWorkspaceModalOpen?: () => void;
};

export type Props = {
  headerComponent: React.ReactNode;
  contentComponent: React.ReactNode;
  sidebarComponent: React.ReactNode;
  collapsed: boolean;
  onCollapse: (collapse: boolean) => void;
};

const CMSWrapper: React.FC<Props> = ({
  contentComponent,
  sidebarComponent,
  headerComponent,
  collapsed,
  onCollapse,
}) => {
  return (
    <Wrapper>
      <HeaderWrapper>{headerComponent}</HeaderWrapper>
      <BodyWrapper>
        <CMSSidebar collapsible collapsed={collapsed} onCollapse={onCollapse} collapsedWidth={54}>
          {sidebarComponent}
        </CMSSidebar>
        <ContentWrapper>{contentComponent}</ContentWrapper>
      </BodyWrapper>
    </Wrapper>
  );
};

const Wrapper = styled(Layout)`
  height: 100vh;
`;

const BodyWrapper = styled(Layout)`
  margin-top: 48px;
`;

const ContentWrapper = styled(Content)`
  overflow: auto;
  height: 100%;
`;

const CMSSidebar = styled(Sider)`
  background-color: #fff;

  .ant-layout-sider-trigger {
    background-color: #fff;
    border-top: 1px solid #f0f0f0;
    color: #002140;
    text-align: left;
    padding: 0 20px;
    margin: 0;
    height: 38px;
    line-height: 38px;
  }
  .ant-layout-sider-children {
    display: flex;
    flex-direction: column;
    justify-content: space-between;
  }
  .ant-menu-inline {
    border-right: 1px solid white;
  }
  .ant-menu-vertical {
    border-right: none;
  }
`;

const HeaderWrapper = styled.div`
  position: fixed;
  z-index: 1;
  width: 100%;
`;

export default CMSWrapper;
