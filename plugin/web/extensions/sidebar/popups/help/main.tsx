import ReactDOM from "react-dom/client";

import Help from ".";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<Help />);
  }
})();

export {};
