import ReactDOM from "react-dom/client";

import WelcomeScreen from "./index";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<WelcomeScreen />);
  }
})();

export {};
