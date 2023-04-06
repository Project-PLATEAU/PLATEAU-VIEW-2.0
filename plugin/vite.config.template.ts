/// <reference types="vitest" />
/// <reference types="vite/client" />

import react from "@vitejs/plugin-react";
import type { UserConfigExport, Plugin } from "vite";
import importToCDN, { autoComplete } from "vite-plugin-cdn-import";
import { viteSingleFile } from "vite-plugin-singlefile";
import svgr from "vite-plugin-svgr";
import tsconfigPaths from "vite-tsconfig-paths";

export const plugin = (name: string): UserConfigExport => ({
  build: {
    outDir: "dist/plugin",
    emptyOutDir: false,
    lib: {
      formats: ["iife"],
      // https://github.com/vitejs/vite/pull/7047
      entry: `src/${name}.ts`,
      name: `ReearthPluginPV_${name}`,
      fileName: () => `${name}.js`,
    },
  },
});

export const web =
  ({
    name,
    parent,
    type = "core",
  }: {
    name: string;
    parent?: string;
    type?: "modal" | "core" | "popup";
  }): UserConfigExport =>
  () => {
    const root =
      parent && type !== "core"
        ? `./web/extensions/${parent}/${type}s/${name}`
        : `./web/extensions/${name}/${type}`;
    const outDir =
      parent && type !== "core"
        ? `../../../../../dist/web/${parent}/${type}s/${name}`
        : `../../../../dist/web/${name}/${type}`;

    return {
      plugins: [
        tsconfigPaths(),
        react(),
        serverHeaders(),
        viteSingleFile(),
        svgr(),
        (importToCDN /* workaround */ as any as { default: typeof importToCDN }).default({
          modules: [
            autoComplete("react"),
            autoComplete("react-dom"),
            {
              name: "react-is",
              var: "react-is",
              path: "https://unpkg.com/react-is@18.2.0/umd/react-is.production.min.js",
            },
            {
              name: "antd",
              var: "antd",
              path: "https://cdnjs.cloudflare.com/ajax/libs/antd/4.22.8/antd.min.js",
              css: "https://cdnjs.cloudflare.com/ajax/libs/antd/4.22.8/antd.min.css",
            },
            {
              name: "styled-components",
              var: "styled-components",
              path: "https://unpkg.com/styled-components@5.3.6/dist/styled-components.min.js",
            },
          ],
        }),
      ],
      publicDir: false,
      emptyOutDir: false,
      root,
      build: {
        outDir,
      },
      css: {
        preprocessorOptions: {
          less: {
            javascriptEnabled: true,
            modifyVars: {
              "primary-color": "#00BEBE",
              "font-family": "Noto Sans",
              "typography-title-font-weight": "500",
              "typography-title-font-height": "21.79px",
            },
          },
        },
      },
      test: {
        globals: true,
        environment: "jsdom",
        setupFiles: "./web/test/setup.ts",
      },
    };
  };

function serverHeaders(): Plugin {
  return {
    name: "server-headers",
    configureServer(server) {
      server.middlewares.use((_req, res, next) => {
        res.setHeader("Service-Worker-Allowed", "/");
        next();
      });
    },
  };
}
