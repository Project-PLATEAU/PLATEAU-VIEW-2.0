export type JsonPrimitive = boolean | number | string | null;

export type JsonArray = (JsonPrimitive | JsonObject | JsonArray)[];

export type JsonObject = {
  [key: string]: JsonPrimitive | JsonObject | JsonArray;
};

export type Json = JsonPrimitive | JsonArray | JsonObject;
