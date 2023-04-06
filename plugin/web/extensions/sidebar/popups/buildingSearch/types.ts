export type InitData = {
  viewport: Viewport;
  data: RawDatasetData;
};

export type RawDatasetData = {
  title: string;
  dataID: string;
  searchIndex: {
    url: string;
  }[];
};

type EnumIndex = {
  kind: string;
  values: {
    [key: string]: EnumValue;
  };
};

type EnumValue = {
  count: number;
  url: string;
};

export type SearchIndex = {
  baseURL: string;
  indexRoot: {
    indexes: {
      [key: string]: EnumIndex;
    };
  };
  resultsData?: any[];
};

export type SearchResults = {
  tilesetId: string;
  results: Result[];
};

//
export type Dataset = {
  title: string;
  dataID: string;
  indexes: IndexData[];
};

export type IndexData = {
  field: string;
  values: string[];
};

export type Condition = {
  field: string;
  values: string[];
};

// from index
export type Result = {
  gml_id: string;
  Longitude: string;
  Latitude: string;
  Height: string;
};

// reearth types
export type Viewport = {
  width: number;
  height: number;
  isMobile: boolean;
};
