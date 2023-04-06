import { Template as TemplateType } from "@web/extensions/sidebar/core/types";
import { ReearthApi } from "@web/extensions/sidebar/types";

export const generalFieldName = {
  idealZoom: "カメラ",
  legend: "凡例",
  legendGradient: "凡例（グラデーション)",
  realtime: "リアルタイム",
  story: "ストーリー",
  timeline: "タイムライン",
  currentTime: "現在時刻",
  switchGroup: "グループの切り替え",
  buttonLink: "リンク",
  styleCode: "スタイルコード",
  switchDataset: "データセットの切り替え",
  switchVisibility: "表示の切り替え",
  point: "ポイント",
  description: "説明",
  template: "テンプレート",
  eventField: "イベント",
  infoboxStyle: "インフォボックス スタイル",
  heightReference: "高さ基準",
};

export const pointFieldName = {
  pointColor: "色",
  pointColorGradient: "色（グラデーション）",
  pointSize: "サイズ",
  pointIcon: "アイコン",
  pointLabel: "ラベル",
  pointModel: "モデル",
  pointStroke: "ストローク",
  pointCSV: "ポイントに変換（CSV）",
};

export const polygonFieldName = {
  polygonColor: "ポリゴン色",
  polygonColorGradient: "ポリゴン色（グラデーション）",
  polygonStroke: "ポリゴンストローク",
  polygonClassificationType: "被せて表示",
};

export const threeDFieldName = {
  clipping: "クリッピング",
  buildingFilter: "フィルター（建築物）",
  buildingTransparency: "透明度",
  buildingColor: "色分け（建築物）",
  buildingShadow: "影",
  floodColor: "色分け（浸水想定区域）",
  floodFilter: "フィルター（浸水想定区域）",
};

export const polylineFieldName = {
  polylineColor: "ポリライン色",
  polylineColorGradient: "ポリライン色（グラデーション）",
  polylineStrokeWeight: "ポリラインストローク",
  polylineClassificationType: "被せて表示",
};

export const fieldName = {
  ...generalFieldName,
  ...pointFieldName,
  ...polygonFieldName,
  ...threeDFieldName,
  ...polylineFieldName,
};

export type FieldComponent =
  | IdealZoom
  | Legend
  | LegendGradient
  | StyleCode
  | ButtonLink
  | Description
  | SwitchGroup
  | ButtonLink
  | Story
  | Realtime
  | Timeline
  | CurrentTime
  | SwitchDataset
  | SwitchVisibility
  | EventField
  | InfoboxStyle
  | HeightReference
  | Template
  | PointColor
  | PointColorGradient
  | PointSize
  | PointIcon
  | PointLabel
  | PointModel
  | PointStroke
  | PointCSV
  | PolylineColor
  | PolylineColorGradient
  | PolylineStrokeWeight
  | PolylineClassificationType
  | PolygonColor
  | PolygonColorGradient
  | PolygonStroke
  | PolygonClassificationType
  | Clipping
  | BuildingFilter
  | BuildingTransparency
  | BuildingColor
  | BuildingShadow
  | FloodColor
  | FloodFilter;

type FieldBase<T extends keyof typeof fieldName> = {
  id: string;
  type: T;
  group?: string;
  override?: any;
  cleanseOverride?: any;
};

type CameraPosition = {
  lng: number;
  lat: number;
  height: number;
  pitch: number;
  heading: number;
  roll: number;
};

export type IdealZoom = FieldBase<"idealZoom"> & {
  position: CameraPosition;
};

export type LegendStyleType = "square" | "circle" | "line" | "icon";

export type LegendItem = {
  title?: string;
  color?: string;
  url?: string;
};

export type Legend = FieldBase<"legend"> & {
  style: LegendStyleType;
  items?: LegendItem[];
};

type LegendGradient = FieldBase<"legendGradient"> & {
  style: Omit<LegendStyleType, "icon">;
  min?: number;
  max?: number;
  startColor?: string;
  endColor?: string;
  step?: number;
};

type CurrentTime = FieldBase<"currentTime"> & {
  currentDate: string;
  currentTime: string;
  startDate: string;
  startTime: string;
  stopDate: string;
  stopTime: string;
};

type Realtime = FieldBase<"realtime"> & {
  updateInterval: number; // 1000 * 60 -> 1m
  userSettings: {
    enabled?: boolean;
  };
};

export type Timeline = FieldBase<"timeline"> & {
  timeFieldName: string;
  userSettings: {
    timeBasedDisplay: boolean;
  };
};

export type Description = FieldBase<"description"> & {
  content?: string;
  isMarkdown?: boolean;
};

export type StyleCode = FieldBase<"styleCode"> & {
  src: string;
};

export type GroupItem = {
  id: string;
  title: string;
  fieldGroupID: string;
};

export type SwitchGroup = FieldBase<"switchGroup"> & {
  title: string;
  groups: GroupItem[];
};

export type SwitchDataset = FieldBase<"switchDataset"> & {
  uiStyle?: "dropdown" | "radio";
  userSettings: {
    selected?: ConfigData;
    override?: any;
  };
};

export type SwitchVisibility = FieldBase<"switchVisibility"> & {
  uiStyle: "dropdown" | "radio";
  conditions: {
    id: string;
    condition: Cond<any>;
    title: string;
  }[];
  userSettings: {
    selected?: string;
  };
};

export type ButtonLink = FieldBase<"buttonLink"> & {
  title?: string;
  link?: string;
};

export type StoryItem = {
  id: string;
  title?: string;
  scenes?: string;
};

export type Story = FieldBase<"story"> & {
  stories?: StoryItem[];
};

type Template = FieldBase<"template"> & {
  templateID?: string;
  userSettings: {
    components?: FieldComponent[];
    override?: any;
  };
};

type EventField = FieldBase<"eventField"> & {
  eventType: string;
  triggerEvent: string;
  urlType: "manual" | "fromData";
  url?: string;
  field?: string;
};

type InfoboxStyle = FieldBase<"infoboxStyle"> & {
  displayStyle: "attributes" | "description" | null;
};

type HeightReference = FieldBase<"heightReference"> & {
  heightReferenceType: "clamp" | "relative" | "none";
};

export type PointColor = FieldBase<"pointColor"> & {
  pointColors?: {
    condition: Cond<number>;
    color: string;
  }[];
};

type PointColorGradient = FieldBase<"pointColorGradient"> & {
  field?: string;
  min?: number;
  max?: number;
  startColor?: string;
  endColor?: string;
  step?: number;
};

type PointSize = FieldBase<"pointSize"> & {
  pointSize?: number;
};

type PointIcon = FieldBase<"pointIcon"> & {
  url?: string;
  size: number;
  sizeInMeters: boolean;
};

type PointLabel = FieldBase<"pointLabel"> & {
  field?: string;
  fontSize?: number;
  fontColor?: string;
  height?: number;
  extruded?: boolean;
  useBackground?: boolean;
  backgroundColor?: string;
};

type PointModel = FieldBase<"pointModel"> & {
  modelURL?: string;
  scale?: number;
};

export type PointStroke = FieldBase<"pointStroke"> & {
  items?: {
    strokeColor: string;
    strokeWidth: number;
    condition: Cond<string | number>;
  }[];
};

type PointCSV = FieldBase<"pointCSV"> & {
  lng?: string;
  lat?: string;
  height?: string;
};

export type PolylineColor = FieldBase<"polylineColor"> & {
  items?: {
    condition: Cond<number>;
    color: string;
  }[];
};

type PolylineColorGradient = FieldBase<"polylineColorGradient"> & {
  field?: string;
  startColor?: string;
  endColor?: string;
  step?: number;
};

type PolylineStrokeWeight = FieldBase<"polylineStrokeWeight"> & {
  strokeWidth: number;
};

export type ClassificationType = "both" | "terrain" | "3dtiles";
export type PolylineClassificationType = FieldBase<"polylineClassificationType"> & {
  classificationType: ClassificationType;
};

export type PolygonColor = FieldBase<"polygonColor"> & {
  items?: {
    condition: Cond<number>;
    color: string;
  }[];
};

type PolygonColorGradient = FieldBase<"polygonColorGradient"> & {
  field?: string;
  startColor?: string;
  endColor?: string;
  step?: number;
};

export type PolygonStroke = FieldBase<"polygonStroke"> & {
  items?: {
    strokeColor: string;
    strokeWidth: number;
    condition: Cond<string | number>;
  }[];
};

export type PolygonClassificationType = FieldBase<"polygonClassificationType"> & {
  classificationType: ClassificationType;
};

type Clipping = FieldBase<"clipping"> & {
  userSettings: {
    enabled: boolean;
    show: boolean;
    aboveGroundOnly: boolean;
    direction: "inside" | "outside";
    override?: any;
  };
};

type BuildingFilter = FieldBase<"buildingFilter"> & {
  userSettings: Record<
    "height" | "abovegroundFloor" | "basementFloor" | "buildingAge",
    | {
        value?: [from: number, to: number];
        min?: number;
        max?: number;
      }
    | undefined
  > & {
    override?: any;
  };
};

type BuildingShadow = FieldBase<"buildingShadow"> & {
  userSettings: {
    shadow: "disabled" | "enabled" | "cast_only" | "receive_only";
    override?: any;
  };
};

type BuildingTransparency = FieldBase<"buildingTransparency"> & {
  userSettings: {
    transparency: number;
    updatedAt?: Date;
    override?: any;
  };
};

type BuildingColor = FieldBase<"buildingColor"> & {
  disableFloodRankLegend?: boolean;
  userSettings: {
    colorType: string;
    updatedAt?: Date;
    override?: any;
  };
};

type FloodColor = FieldBase<"floodColor"> & {
  userSettings: {
    colorType: "water" | "rank";
    updatedAt?: Date;
    override?: any;
  };
};

type FloodFilter = FieldBase<"floodFilter"> & {
  userSettings: {
    value?: [from: number, to: number];
    min?: number;
    max?: number;
    isOrg?: boolean;
    override?: any;
  };
};

export type Fields = {
  // general
  idealZoom: IdealZoom;
  legend: Legend;
  legendGradient: LegendGradient;
  description: Description;
  styleCode: StyleCode;
  switchGroup: SwitchGroup;
  buttonLink: ButtonLink;
  story: Story;
  currentTime: CurrentTime;
  realtime: Realtime;
  timeline: Timeline;
  switchDataset: SwitchDataset;
  switchVisibility: SwitchVisibility;
  eventField: EventField;
  infoboxStyle: InfoboxStyle;
  heightReference: HeightReference;
  // point
  pointColor: PointColor;
  pointColorGradient: PointColorGradient;
  pointSize: PointSize;
  pointIcon: PointIcon;
  pointLabel: PointLabel;
  pointModel: PointModel;
  pointStroke: PointStroke;
  pointCSV: PointCSV;
  // polyline
  polylineColor: PolylineColor;
  polylineColorGradient: PolylineColorGradient;
  polylineStrokeWeight: PolylineStrokeWeight;
  polylineClassificationType: PolylineClassificationType;
  // polygon
  polygonColor: PolygonColor;
  polygonColorGradient: PolygonColorGradient;
  polygonStroke: PolygonStroke;
  polygonClassificationType: PolygonClassificationType;
  // 3d-model
  // 3d-tile
  clipping: Clipping;
  buildingFilter: BuildingFilter;
  buildingTransparency: BuildingTransparency;
  buildingColor: BuildingColor;
  buildingShadow: BuildingShadow;
  floodColor: FloodColor;
  floodFilter: FloodFilter;
  // template
  template: Template;
};

export type BaseFieldProps<T extends keyof Fields> = {
  value: Fields[T];
  dataID?: string;
  editMode?: boolean;
  isActive?: boolean;
  activeIDs?: string[];
  templates?: TemplateType[];
  selectedGroup?: string;
  configData?: ConfigData[];
  onUpdate: (property: Fields[T]) => void;
  onSceneUpdate: (updatedProperties: Partial<ReearthApi>) => void;
  onCurrentGroupUpdate?: (fieldGroupID: string) => void;
  onCurrentDatasetUpdate?: (data?: ConfigData) => void;
};

export type ConfigData = { name: string; type: string; url: string; layer?: string[] };

export type Expression<T extends string | number | boolean = string | number | boolean> =
  | T
  | {
      conditions: Cond<T>[];
    }
  | {
      gradient: {
        key: string;
        defaultValue?: T;
        steps: { min?: number; max: number; value: T }[];
      };
    };

export type Cond<T> = {
  key: string;
  operator: "===" | ">=" | "<=" | ">" | "<" | "!==";
  operand: T;
  value: T;
};
