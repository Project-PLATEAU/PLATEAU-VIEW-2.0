import GeolocationWrapper from "@web/extensions/geolocation/core";
import ReactDOM from "react-dom/client";

(async () => {
  const element = document.getElementById("root");
  if (element) {
    const root = ReactDOM.createRoot(element);
    root.render(<GeolocationWrapper />);
  }
})();

export {};
