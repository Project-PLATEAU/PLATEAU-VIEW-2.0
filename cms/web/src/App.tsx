import { BrowserRouter as Router, Route, Routes } from "react-router-dom";

import { Provider as Auth0Provider } from "@reearth-cms/auth";
import NotFound from "@reearth-cms/components/atoms/NotFound";
import AccountSettings from "@reearth-cms/components/organisms/Account";
import Asset from "@reearth-cms/components/organisms/Asset/Asset";
import AssetList from "@reearth-cms/components/organisms/Asset/AssetList";
import Accessibility from "@reearth-cms/components/organisms/Project/Accessibility";
import ContentDetails from "@reearth-cms/components/organisms/Project/Content/ContentDetails";
import Content from "@reearth-cms/components/organisms/Project/Content/ContentList";
import ProjectOverview from "@reearth-cms/components/organisms/Project/Overview";
import RequestDetails from "@reearth-cms/components/organisms/Project/Request/RequestDetails";
import RequestList from "@reearth-cms/components/organisms/Project/Request/RequestList";
import Schema from "@reearth-cms/components/organisms/Project/Schema";
import ProjectSettings from "@reearth-cms/components/organisms/Project/settings";
import Integration from "@reearth-cms/components/organisms/Settings/Integration";
import Members from "@reearth-cms/components/organisms/Settings/Members";
import MyIntegrationDetails from "@reearth-cms/components/organisms/Settings/MyIntegrationDetails";
import MyIntegrations from "@reearth-cms/components/organisms/Settings/MyIntegrations";
import Workspace from "@reearth-cms/components/organisms/Workspace";
import WorkspaceSettings from "@reearth-cms/components/organisms/Workspace/Settings";
import CMSPageWrapper from "@reearth-cms/components/pages/CMSPage";
import RootPage from "@reearth-cms/components/pages/RootPage";
import { Provider as GqlProvider } from "@reearth-cms/gql";
import { Provider as I18nProvider } from "@reearth-cms/i18n";

function App() {
  return (
    <Auth0Provider>
      <GqlProvider>
        <I18nProvider>
          <Router>
            <Routes>
              <Route path="/" element={<RootPage />} />
              <Route path="workspace" element={<CMSPageWrapper />}>
                <Route path=":workspaceId" element={<Workspace />} />
                <Route path=":workspaceId/account" element={<AccountSettings />} />
                <Route path=":workspaceId/members" element={<Members />} />
                <Route path=":workspaceId/myIntegrations" element={<MyIntegrations />} />
                <Route
                  path=":workspaceId/myIntegrations/:integrationId"
                  element={<MyIntegrationDetails />}
                />
                <Route path=":workspaceId/integrations" element={<Integration />} />
                {/* <Route
                    path=":workspaceId/role"
                    element={<div>Role page - GOTTA DO THIS PAGE!!!</div>}
                  />
                  <Route
                    path=":workspaceId/apiKey"
                    element={<div>API Key page - GOTTA DO THIS PAGE!!!</div>}
                  /> */}
                <Route path=":workspaceId/settings" element={<WorkspaceSettings />} />
                <Route path=":workspaceId/project/:projectId" element={<ProjectOverview />} />
                <Route path=":workspaceId/project/:projectId/schema" element={<Schema />} />
                <Route
                  path=":workspaceId/project/:projectId/schema/:modelId"
                  element={<Schema />}
                />
                <Route
                  path=":workspaceId/project/:projectId/accessibility"
                  element={<Accessibility />}
                />
                <Route
                  path=":workspaceId/project/:projectId/settings"
                  element={<ProjectSettings />}
                />
                <Route path=":workspaceId/project/:projectId/content" element={<Content />} />
                <Route
                  path=":workspaceId/project/:projectId/content/:modelId"
                  element={<Content />}
                />
                <Route
                  path=":workspaceId/project/:projectId/content/:modelId/details"
                  element={<ContentDetails />}
                />
                <Route
                  path=":workspaceId/project/:projectId/content/:modelId/details/:itemId"
                  element={<ContentDetails />}
                />
                <Route path=":workspaceId/project/:projectId/asset" element={<AssetList />} />
                <Route path=":workspaceId/project/:projectId/asset/:assetId" element={<Asset />} />
                <Route path=":workspaceId/project/:projectId/request" element={<RequestList />} />
                <Route
                  path=":workspaceId/project/:projectId/request/:requestId"
                  element={<RequestDetails />}
                />
              </Route>
              <Route path="*" element={<NotFound />} />
            </Routes>
          </Router>
        </I18nProvider>
      </GqlProvider>
    </Auth0Provider>
  );
}

export default App;
