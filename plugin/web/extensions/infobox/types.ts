import type { Template } from "@web/extensions/sidebar/core/types";

export type { InfoboxField as Field } from "@web/extensions/sidebar/core/types";
export type InfoboxTemplate = Omit<Template, "components">;

export type ActionType = "getInEditor" | "saveTemplate" | "init";
export type PostMessageProps = { action: ActionType; payload?: any };

// Reearth types
export type PluginExtensionInstance = {
  id: string;
  pluginId: string;
  name: string;
  extensionId: string;
  extensionType: "widget" | "block";
};

export type PluginMessage = {
  data: {
    action: "infoboxFieldsFetch";
    payload: any;
  };
  sender: string;
};

export type FixedProperties = {
  gml_id?: string;
  名称?: string;
  用途?: string;
  住所?: string;
  地上階数?: number;
  地下階数?: number;
  建築年?: number;
  計測高さ?: number;
  "建物利用現況（大分類）"?: string;
  "建物利用現況（中分類）"?: string;
  "建物利用現況（小分類）"?: string;
  "建物利用現況（詳細分類）"?: string;
  構造種別?: string;
  "構造種別（自治体独自）"?: string;
  耐火構造種別?: string;
};

export type Properties = FixedProperties & {
  attributes?: any;
  [key: string]: any;
};

export type FldInfo = {
  name?: string;
  datasetName?: string;
};
