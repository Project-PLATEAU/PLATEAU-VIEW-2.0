import { Select } from "@web/sharedComponents";

export const fileFormats = ".kml,.csv,.czml,.gpx,.geojson,.georss,.shapefile,.zip,.glb,.gltf";

export type FileType =
  | "auto"
  | "geojson"
  | "kml"
  | "csv"
  | "czml"
  | "gpx"
  | "georss"
  | "shapefile"
  | "gltf";

type Props = {
  onFileTypeSelect: (value: string) => void;
};

const FileTypeSelect: React.FC<Props> = ({ onFileTypeSelect }) => {
  const options = [
    {
      value: "auto",
      label: "自動検出",
    },
    {
      value: "geojson",
      label: "GeoJSON",
    },
    {
      value: "kml",
      label: "KML・KMZ",
    },
    {
      value: "csv",
      label: "CSV",
    },
    {
      value: "czml",
      label: "CZML",
    },
    {
      value: "gpx",
      label: "GPX",
    },
    {
      value: "georss",
      label: "GeoRSS",
    },
    {
      value: "shapefile",
      label: "ShapeFile (zip)",
    },
    {
      value: "gltf",
      label: "GLTF/GLB",
    },
  ];

  return (
    <Select
      defaultValue="auto"
      style={{ width: "100%" }}
      onChange={onFileTypeSelect}
      options={options}
    />
  );
};

export default FileTypeSelect;
