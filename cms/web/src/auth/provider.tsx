import { Auth0Provider } from "@auth0/auth0-react";

type Props = {
  children?: React.ReactNode;
};

const Provider: React.FC<Props> = ({ children }) => {
  const domain = window.REEARTH_CONFIG?.auth0Domain;
  const clientId = window.REEARTH_CONFIG?.auth0ClientId;
  const audience = window.REEARTH_CONFIG?.auth0Audience;

  return domain && clientId ? (
    <Auth0Provider
      domain={domain}
      clientId={clientId}
      audience={audience}
      useRefreshTokens
      scope="openid profile email"
      cacheLocation="localstorage"
      redirectUri={window.location.origin}>
      {children}
    </Auth0Provider>
  ) : (
    // TODO
    <>{children}</>
  );
};

export default Provider;
