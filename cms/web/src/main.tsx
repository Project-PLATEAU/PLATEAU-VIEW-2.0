import { StrictMode } from "react";
import * as ReactDOM from "react-dom/client";

import App from "./App";
import loadConfig from "./config";

import "antd/dist/antd.css";
import "./index.css";

(async function () {
  try {
    await loadConfig();
  } finally {
    const element = document.getElementById("root");
    if (element) {
      const root = ReactDOM.createRoot(element);
      root.render(
        <StrictMode>
          <App />
        </StrictMode>,
      );
    }
  }
})();
