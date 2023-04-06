/// <reference types="vite/client" />

declare module "*.yml" {
  const yml: any;
  export default yml;
}

declare module "*.yaml" {
  const yml: any;
  export default yml;
}

interface ImportMetaEnv {
  readonly REEARTH_CMS_API: string;
  readonly REEARTH_CMS_AUTH0_DOMAIN: string;
  readonly REEARTH_CMS_AUTH0_AUDIENCE: string;
  readonly REEARTH_CMS_AUTH0_CLIENT_ID: string;
}

interface ImportMeta {
  readonly env: ImportMetaEnv;
}
