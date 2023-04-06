import omit from "lodash/omit";
import pick from "lodash/pick";

const APPEARANCES = {
  marker: 1,
  polyline: 1,
  polygon: 1,
  model: 1,
  "3dtiles": 1,
  ellipsoid: 1,
  box: 1,
  photooverlay: 1,
  resource: 1,
  raster: 1,
};

// GTFS needs to hide entity(like point) when camera distance is near to model.
// And entity is hidden, then bus model is appeared.
export const proxyGTFS = (layer: any) => {
  const appearancesNeedNearFar = omit(APPEARANCES, "3dtiles", "resource", "raster", "model");
  const layerByAppearance = pick(layer, Object.keys(appearancesNeedNearFar));

  const result: Record<string, any> = {};
  Object.keys(layerByAppearance).forEach(k => {
    result[k] = {
      ...layerByAppearance[k],
      near: 1000,
      clampToGround: true,
    };
  });

  return { ...layer, ...result };
};
