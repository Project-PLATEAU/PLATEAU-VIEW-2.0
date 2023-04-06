import ReactDOM from "react-dom/client";

import MobileDropdown from ".";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<MobileDropdown />);
  }
})();

export {};
