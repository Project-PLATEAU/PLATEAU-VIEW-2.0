import Storytelling from "@web/extensions/storytelling/core";
import ReactDOM from "react-dom/client";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<Storytelling />);
  }
})();

export {};
