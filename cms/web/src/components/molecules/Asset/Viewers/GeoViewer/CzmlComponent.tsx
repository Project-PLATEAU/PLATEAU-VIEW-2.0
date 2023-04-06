import { CzmlDataSource } from "cesium";
import { ComponentProps, useCallback } from "react";
import { CzmlDataSource as ResiumCzmlDataSource, useCesium } from "resium";

type Props = ComponentProps<typeof ResiumCzmlDataSource>;

const CzmlComponent: React.FC<Props> = ({ data }) => {
  const { viewer } = useCesium();

  const handleLoad = useCallback(
    async (ds: CzmlDataSource) => {
      try {
        await viewer?.zoomTo(ds);
        ds.show = true;
      } catch (error) {
        console.error(error);
      }
    },
    [viewer],
  );

  return <ResiumCzmlDataSource data={data} onLoad={handleLoad} />;
};

export default CzmlComponent;
