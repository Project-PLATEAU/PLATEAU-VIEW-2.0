import { useTranslation } from "react-i18next";

export { default as Provider } from "./provider";
export { localesWithLabel } from "./locale";

export const useT = () => useTranslation().t;
export const useLang = () => useTranslation().i18n.language;
