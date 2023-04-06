import type {
  DataCatalogGroup,
  DataCatalogItem,
  DataCatalogTreeItem,
} from "@web/extensions/sidebar/core/types";
import { omit } from "lodash";

import { makeTree, mapTree } from "./utils";

// TODO: REFACTOR: CONFUSING REEXPORT
export type { DataCatalogItem, DataCatalogGroup, DataCatalogTreeItem };

export type RawDataCatalogTreeItem = RawDataCatalogGroup | RawDataCatalogItem;

export type RawDataCatalogGroup = {
  id: string;
  name: string;
  desc?: string;
  children: RawDataCatalogTreeItem[];
};

type RawRawDataCatalogItem = {
  id: string;
  itemId?: string;
  name?: string;
  pref: string;
  pref_code?: string;
  pref_code_i: number;
  city?: string;
  city_en?: string;
  city_code?: string;
  city_code_i: number;
  ward?: string;
  ward_en?: string;
  ward_code?: string;
  ward_code_i: number;
  type: string;
  type_en: string;
  type2?: string;
  type2_en?: string;
  format: string;
  layers?: string[] | string;
  layer?: string[] | string;
  url: string;
  desc: string;
  year: number;
  tags?: { type: "type" | "location"; value: string }[];
  openDataUrl?: string;
  config?: {
    data?: {
      name: string;
      type: string;
      url: string;
      layers?: string[] | string;
      layer?: string[] | string;
    }[];
  };
  order?: number;
  // bldg only fields
  bldg_low_texture_url?: string;
  bldg_no_texture_url?: string;
  search_index?: string;
  // internal
  path?: string[];
  code: number;
};

export type RawDataCatalogItem = Omit<RawRawDataCatalogItem, "layers" | "layer" | "config"> & {
  layers?: string[];
  config?: {
    data?: {
      name: string;
      type: string;
      url: string;
      layer?: string[];
    }[];
  };
};

export type GroupBy = "city" | "type" | "tag"; // Tag not implemented yet

export async function getDataCatalog(
  base: string,
  project?: string,
): Promise<RawDataCatalogItem[]> {
  const res = await fetch(`${base}/datacatalog${project ? `/${project}` : ""}`);
  if (!res.ok) {
    throw new Error("failed to fetch data catalog");
  }

  const data: RawRawDataCatalogItem[] = await res.json();
  return data.map(modifyDataCatalog);
}

export function modifyDataCatalog(
  d: Omit<RawRawDataCatalogItem, "pref_code_i" | "city_code_i" | "ward_code_i" | "tags" | "code">,
): RawDataCatalogItem {
  const pref = d.pref === "全国" || d.pref === "全球" ? zenkyu : d.pref;
  const pref_code = d.pref === "全国" || d.pref === "全球" || d.pref === zenkyu ? "0" : d.pref_code;
  const pref_code_i = parseInt(pref_code ?? "");
  const city_code_i = parseInt(d.city_code ?? "");
  const ward_code_i = parseInt(d.ward_code ?? "");
  return {
    ...omit(d, ["layers", "layer", "config"]),
    pref,
    pref_code,
    pref_code_i,
    city_code_i,
    ward_code_i,
    code: !isNaN(ward_code_i)
      ? ward_code_i
      : !isNaN(city_code_i)
      ? city_code_i
      : !isNaN(pref_code_i)
      ? pref_code_i * 1000
      : pref === zenkyu
      ? 0
      : 99999,
    tags: [
      { type: "type", value: d.type },
      ...(d.type2 ? [{ type: "type", value: d.type2 } as const] : []),
      ...(d.city ? [{ type: "location", value: d.city } as const] : []),
      ...(d.ward ? [{ type: "location", value: d.ward } as const] : []),
    ],
    ...(d.layers || d.layer ? { layers: [...getLayers(d.layers), ...getLayers(d.layer)] } : {}),
    ...(d.config
      ? {
          config: {
            ...(d.config.data
              ? {
                  data: d.config.data.map(dd => ({
                    ...omit(dd, ["layers", "layer"]),
                    layer: [...getLayers(dd.layers), ...getLayers(dd.layer)],
                  })),
                }
              : {}),
          },
        }
      : {}),
  };
}

// TODO: REFACTOR: confusing typing
export function getDataCatalogTree(
  items: DataCatalogItem[],
  groupBy: GroupBy,
  q?: string | undefined,
): DataCatalogTreeItem[] {
  return getRawDataCatalogTree(items, groupBy, q) as DataCatalogTreeItem[];
}

export function getRawDataCatalogTree(
  items: RawDataCatalogItem[],
  groupBy: GroupBy,
  q?: string | undefined,
): (RawDataCatalogGroup | RawDataCatalogItem)[] {
  return mapTree(
    makeTree(sortInternal(items, groupBy, q)),
    (item): RawDataCatalogGroup | RawDataCatalogItem =>
      item.item ?? {
        id: item.id,
        name: item.name,
        desc: item.desc,
        children: [],
      },
  );
}

type InternalDataCatalogItem = RawDataCatalogItem & {
  path: string[];
};

function sortInternal(
  items: RawDataCatalogItem[],
  groupBy: GroupBy,
  q?: string | undefined,
): InternalDataCatalogItem[] {
  return filter(q, items)
    .map(
      (i): InternalDataCatalogItem => ({
        ...i,
        path: path(i, groupBy),
      }),
    )
    .sort((a, b) => sortBy(a, b, groupBy));
}

function path(i: RawDataCatalogItem, groupBy: GroupBy): string[] {
  return groupBy === "type"
    ? [
        i.type,
        i.pref,
        ...((i.ward || i.type2) && i.city ? [i.city] : []),
        ...(i.name || "（名称未決定）").split("/"),
      ]
    : [
        i.pref,
        ...(i.city ? [i.city] : []),
        ...(i.ward ? [i.ward] : []),
        ...(i.type2 ||
        ((i.type_en === "usecase" ||
          i.type_en === "fld" ||
          i.type_en === "htd" ||
          i.type_en === "tnm" ||
          i.type_en === "ifld") &&
          i.pref !== zenkyu)
          ? [i.type]
          : []),
        ...(i.name || "（名称未決定）").split("/"),
      ];
}

function sortBy(a: InternalDataCatalogItem, b: InternalDataCatalogItem, sort: GroupBy): number {
  return sort === "type"
    ? sortByType(a, b) || sortByCity(a, b) || sortByOrder(a.order, b.order)
    : sortByCity(a, b) || sortByType(a, b) || sortByOrder(a.order, b.order);
}

function sortByCity(a: InternalDataCatalogItem, b: InternalDataCatalogItem): number {
  return clamp(
    (a.pref === zenkyu ? 0 : 1) - (b.pref === zenkyu ? 0 : 1) || // items whose prefecture is zenkyu is upper
      (a.pref === tokyo ? 0 : 1) - (b.pref === tokyo ? 0 : 1) || // items whose prefecture is tokyo is upper
      a.pref_code_i - b.pref_code_i ||
      (a.ward ? 0 : 1) - (b.ward ? 0 : 1) || // items that have a ward is upper
      (a.city ? 0 : 1) - (b.city ? 0 : 1) || // items that have a city is upper
      a.code - b.code ||
      types.indexOf(a.type_en) - types.indexOf(b.type_en),
  );
}

function sortByType(a: RawDataCatalogItem, b: RawDataCatalogItem): number {
  return clamp(types.indexOf(a.type_en) - types.indexOf(b.type_en));
}

function sortByOrder(a: number | undefined, b: number | undefined): number {
  return clamp(Math.min(0, a ?? 0) - Math.min(0, b ?? 0));
}

function filter(q: string | undefined, items: RawDataCatalogItem[]): RawDataCatalogItem[] {
  if (!q) return items;
  return items.filter(
    i => i.name?.includes(q) || i.pref.includes(q) || i.city?.includes(q) || i.ward?.includes(q),
  );
}

function clamp(n: number): number {
  return Math.max(-1, Math.min(1, n));
}

function getLayers(layers?: string[] | string): string[] {
  return layers ? (typeof layers === "string" ? layers.split(/, */).filter(Boolean) : layers) : [];
}

const zenkyu = "全球データ";
const tokyo = "東京都";
const types = [
  "bldg",
  "tran",
  "brid",
  "rail",
  "veg",
  "frn",
  "luse",
  "lsld",
  "urf",
  "fld",
  "tnm",
  "htd",
  "ifld",
  "gen",
  "shelter",
  "landmark",
  "station",
  "emergency_route",
  "railway",
  "park",
  "border",
  "usecase",
];
