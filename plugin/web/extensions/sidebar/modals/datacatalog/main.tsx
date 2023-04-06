import ReactDOM from "react-dom/client";

import DataCatalog from "./components";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<DataCatalog />);
  }
})();

export {};
