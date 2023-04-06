export type Model = {
  id: string;
  name: string;
  description?: string;
  key: string;
  schema: Schema;
  public: boolean;
};

export type Schema = {
  id: string;
  fields: Field[];
};

export type Field = {
  id: string;
  type: FieldType;
  title: string;
  key: string;
  description: string | null | undefined;
  required: boolean;
  unique: boolean;
  multiple: boolean;
  typeProperty?: TypeProperty;
};

export type FieldType =
  | "Text"
  | "TextArea"
  | "RichText"
  | "MarkdownText"
  | "Asset"
  | "Date"
  | "Bool"
  | "Select"
  | "Tag"
  | "Integer"
  | "Reference"
  | "URL";

export type TypeProperty =
  | {
      defaultValue?: string | number;
      maxLength?: number;
      assetDefaultValue?: string;
      selectDefaultValue?: string;
      integerDefaultValue?: number;
      min?: number;
      max?: number;
    }
  | any;

export type CreationFieldTypePropertyInput = {
  asset?: { defaultValue: string };
  integer?: { defaultValue: number; min: number; max: number };
  markdownText?: { defaultValue: string; maxLength: number };
  select?: { defaultValue: string; values: string[] };
  text?: { defaultValue: string; maxLength: number };
  textArea?: { defaultValue: string; maxLength: number };
  url?: { defaultValue: string };
};

export type FieldModalTabs = "settings" | "validation" | "defaultValue";
