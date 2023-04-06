import { Viewer as CesiumViewer } from "cesium";
import { ComponentProps, useMemo, useState } from "react";

import ResiumViewer from "@reearth-cms/components/atoms/ResiumViewer";

import { Imagery, Property } from "./Imagery";

type Props = {
  viewerProps?: ComponentProps<typeof ResiumViewer>;
  url: string;
  onGetViewer: (viewer: CesiumViewer | undefined) => void;
};

const MvtViewer: React.FC<Props> = ({ viewerProps, url, onGetViewer }) => {
  const [properties, setProperties] = useState<Property>();
  const properties2 = useMemo(() => {
    if (typeof properties !== "object" || !properties) return properties;
    const attributes = properties.attributes;
    if (typeof attributes !== "string") {
      return properties;
    }
    try {
      return { ...properties, attributes: JSON.parse(attributes) };
    } catch {
      return properties;
    }
  }, [properties]);

  return (
    <ResiumViewer {...viewerProps} onGetViewer={onGetViewer} properties={properties2}>
      <Imagery url={url} handleProperties={setProperties} />
    </ResiumViewer>
  );
};

export default MvtViewer;
