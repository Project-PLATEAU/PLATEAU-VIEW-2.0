import { KmlDataSource } from "cesium";
import { ComponentProps, useCallback } from "react";
import { KmlDataSource as ResiumKmlDataSource, useCesium } from "resium";

type Props = ComponentProps<typeof ResiumKmlDataSource>;

const KmlComponent: React.FC<Props> = ({ data }) => {
  const { viewer } = useCesium();

  const handleLoad = useCallback(
    async (ds: KmlDataSource) => {
      try {
        await viewer?.zoomTo(ds);
        ds.show = true;
      } catch (error) {
        console.error(error);
      }
    },
    [viewer],
  );

  return <ResiumKmlDataSource data={data} onLoad={handleLoad} />;
};

export default KmlComponent;
