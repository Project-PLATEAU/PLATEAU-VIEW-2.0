import {
  Cartesian3,
  Model,
  Viewer,
  HeadingPitchRoll,
  Transforms,
  Ellipsoid,
  BoundingSphere,
} from "cesium";
import { useCallback, useEffect, useRef } from "react";
import { useCesium } from "resium";

type Props = {
  url: string;
};

export const Imagery: React.FC<Props> = ({ url }) => {
  const { viewer } = useCesium() as { viewer: Viewer | undefined };
  const modelRef = useRef<Model | undefined>();

  const loadModel = useCallback(async () => {
    if (!viewer) return;
    const position = Cartesian3.ZERO;
    const hpr = new HeadingPitchRoll(0, 0, 0);

    const model = viewer.scene.primitives.add(
      Model.fromGltf({
        url,
        modelMatrix: Transforms.headingPitchRollToFixedFrame(position, hpr, Ellipsoid.WGS84),
        show: true,
      }),
    );

    modelRef.current = model;
    viewer.scene.primitives.add(model);
    model.readyPromise.then(function (model: { boundingSphere: BoundingSphere }) {
      viewer.camera.flyToBoundingSphere(model.boundingSphere, {
        duration: 0.5,
      });
    });
  }, [url, viewer]);

  useEffect(() => {
    loadModel();
  }, [loadModel]);

  return null;
};
