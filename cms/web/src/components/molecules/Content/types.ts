import { FieldType } from "@reearth-cms/components/molecules/Schema/types";

export type ItemStatus = "DRAFT" | "PUBLIC" | "REVIEW" | "PUBLIC_REVIEW" | "PUBLIC_DRAFT";

export type ItemField = {
  schemaFieldId: string;
  type: FieldType;
  value: any;
};

export type Item = {
  id: string;
  schemaId: string;
  fields: ItemField[] | undefined | null;
  threadId: string;
  comments: Comment[];
};

export type ContentTableField = {
  id: string;
  author: string;
  schemaId: string;
  status: ItemStatus;
  modelId?: string;
  fields: { [key: string]: any };
  comments: Comment[];
  createdAt: Date;
  updatedAt: Date;
};

export type Comment = {
  id: string;
  author: { id?: string; name: string; type: "User" | "Integration" | null };
  content: string;
  createdAt: string;
};
