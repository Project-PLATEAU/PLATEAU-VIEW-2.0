import { CSSProperties } from "react";

import Select, { DefaultOptionType } from "@reearth-cms/components/atoms/Select";
import { useT } from "@reearth-cms/i18n";

export type PreviewType =
  | "GEO"
  | "GEO_3D_TILES"
  | "GEO_MVT"
  | "IMAGE"
  | "IMAGE_SVG"
  | "MODEL_3D"
  | "UNKNOWN";

type Props = {
  onTypeChange: (
    value: PreviewType,
    option: DefaultOptionType | DefaultOptionType[],
  ) => void | undefined;
  style?: CSSProperties;
  value?: PreviewType;
};

type PreviewTypeListItem = {
  id: number;
  name: string;
  value: PreviewType;
};

export const PreviewTypeSelect: React.FC<Props> = ({ onTypeChange, style, value }) => {
  const t = useT();
  const previewTypeList: PreviewTypeListItem[] = [
    { id: 1, name: t("PNG/JPEG/TIFF/GIF"), value: "IMAGE" },
    { id: 2, name: t("SVG"), value: "IMAGE_SVG" },
    {
      id: 3,
      name: t("GEOJSON/KML/CZML"),
      value: "GEO",
    },
    { id: 4, name: t("3D Tiles"), value: "GEO_3D_TILES" },
    { id: 5, name: t("MVT"), value: "GEO_MVT" },
    { id: 6, name: t("GLTF/GLB"), value: "MODEL_3D" },
    { id: 7, name: t("Unknown Type"), value: "UNKNOWN" },
  ];
  return (
    <Select style={style} value={value} onChange={onTypeChange}>
      {previewTypeList.map((type: PreviewTypeListItem) => (
        <Select.Option key={type.id} value={type.value}>
          {type.name}
        </Select.Option>
      ))}
    </Select>
  );
};
