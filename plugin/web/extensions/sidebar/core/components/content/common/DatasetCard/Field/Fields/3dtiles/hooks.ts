import { postMsg } from "@web/extensions/sidebar/utils";
import { useEffect, useState } from "react";

export const useObservingDataURL = (dataID: string | undefined) => {
  const [url, setURL] = useState<string>();
  useEffect(() => {
    const waitReturnedPostMsg = async (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "findTileset" && e.data.payload.dataID === dataID) {
        const layer = e.data.payload.layer;
        if (layer.data.url) {
          setURL(layer.data.url);
        }
      }
    };
    addEventListener("message", waitReturnedPostMsg);
    // Wait until layer is overridden
    const timeoutId = window.setTimeout(() => {
      postMsg({
        action: "findTileset",
        payload: {
          dataID,
        },
      });
    }, 300);

    return () => {
      if (timeoutId) window.clearTimeout(timeoutId);
    };
  });

  return url;
};
