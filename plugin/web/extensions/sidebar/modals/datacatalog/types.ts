import { DataCatalogItem } from "./api/api";

export type UserDataItem = Partial<DataCatalogItem> & {
  description?: string;
};
