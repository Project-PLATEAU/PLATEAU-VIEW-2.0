import { useEffect, useState } from "react";

import { postMsg } from "../utils";

import DesktopSidebar from "./components/desktop";
import MobileSidebar from "./components/mobile";

export type Props = {
  className?: string;
};

const Sidebar: React.FC<Props> = ({ className }) => {
  const [isMobile, setIsMobile] = useState<boolean>();

  useEffect(() => {
    postMsg({ action: "checkIfMobile" });
  }, []);

  useEffect(() => {
    const eventListenerCallback = (e: MessageEvent<any>) => {
      if (e.source !== parent) return;
      if (e.data.action === "checkIfMobile") {
        setIsMobile(e.data.payload);
      }
    };
    addEventListener("message", e => eventListenerCallback(e));
    return () => {
      removeEventListener("message", eventListenerCallback);
    };
  }, []); // eslint-disable-line react-hooks/exhaustive-deps

  return isMobile !== undefined ? (
    isMobile ? (
      <MobileSidebar className={className} />
    ) : (
      <DesktopSidebar className={className} />
    )
  ) : null;
};

export default Sidebar;
