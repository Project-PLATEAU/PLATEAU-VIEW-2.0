/// <reference types="vitest" />
/// <reference types="vite/client" />

import { defineConfig } from "vite";

import { web } from "../../../../vite.config.template";

// https://vitejs.dev/config/
export default defineConfig(web({ name: "groupSelect", parent: "sidebar", type: "popup" }));
