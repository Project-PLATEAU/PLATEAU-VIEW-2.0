import { postMsg } from "@web/extensions/sidebar/utils";
import { useCallback, useEffect, useState } from "react";

export type Tab = "basic" | "map" | "shadow" | "clip";

type Items = {
  label: string;
  key: Tab;
};

const items: Items[] = [
  { label: "基本操作", key: "basic" },
  { label: "マップを使ってみる", key: "map" },
  { label: "日影機能について", key: "shadow" },
  { label: "クリップ機能", key: "clip" },
];

export default () => {
  const [selectedTab, changeTab] = useState<Tab>("basic");

  useEffect(() => {
    postMsg({ action: "msgToPopup", payload: selectedTab });
  }, [selectedTab]);

  useEffect(() => {
    const eventListenerCallback = (e: any) => {
      if (e.source !== parent) return null;
      if (e.data.action) {
        if (e.data.action === "initPopup") {
          postMsg({ action: "msgToPopup", payload: selectedTab });
        }
      }
    };
    (globalThis as any).addEventListener("message", eventListenerCallback);
    return () => {
      (globalThis as any).removeEventListener("message", eventListenerCallback);
    };
  });

  const handleItemClicked = useCallback((key: Tab) => {
    changeTab(key);
    postMsg({ action: "helpPopupOpen" });
  }, []);

  return {
    items,
    selectedTab,
    handleItemClicked,
  };
};
