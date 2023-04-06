import i18next from "i18next";

export const fieldTypes: {
  [P: string]: { icon: string; title: string; description: string; color: string };
} = {
  Text: {
    icon: "textT",
    title: i18next.t("Text"),
    description: i18next.t("Heading and titles, one-line field"),
    color: "#FF7875",
  },
  TextArea: {
    icon: "textAlignLeft",
    title: i18next.t("TextArea"),
    description: i18next.t("Multi line text"),
    color: "#FF7875",
  },
  MarkdownText: {
    icon: "markdown",
    title: i18next.t("Markdown text"),
    description: i18next.t("Rich text which supports md style"),
    color: "#FF7875",
  },
  Asset: {
    icon: "asset",
    title: i18next.t("Asset"),
    description: i18next.t("Asset file"),
    color: "#FF9C6E",
  },
  Bool: {
    icon: "boolean",
    title: i18next.t("Boolean"),
    description: i18next.t("true/false field"),
    color: "#FFD666",
  },
  Select: {
    icon: "listBullets",
    title: i18next.t("Option"),
    description: i18next.t("Multiple select"),
    color: "#7CB305",
  },
  Integer: {
    icon: "numberNine",
    title: i18next.t("Int"),
    description: i18next.t("Integer"),
    color: "#36CFC9",
  },
  URL: {
    icon: "link",
    title: i18next.t("URL"),
    description: i18next.t("http/https URL"),
    color: "#9254DE",
  },
};
