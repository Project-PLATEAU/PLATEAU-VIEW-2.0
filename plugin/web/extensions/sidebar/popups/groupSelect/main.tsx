import ReactDOM from "react-dom/client";

import GroupSelect from ".";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<GroupSelect />);
  }
})();

export {};
