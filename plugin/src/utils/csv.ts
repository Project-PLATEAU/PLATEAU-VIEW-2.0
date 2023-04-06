const LNG_PROPERTY_NAMES = ["lng", "longitude"];
const LAT_PROPERTY_NAMES = ["lat", "latitude"];
const HEIGHT_PROPERTY_NAMES = ["height"];

export const getLocationNamesFromFeatureProperties = (
  feature: { properties?: any } | undefined,
) => {
  return Object.keys(feature?.properties || {}).reduce((res, v) => {
    if (!res.lng && LNG_PROPERTY_NAMES.includes(v)) {
      res.lng = v;
    }
    if (!res.lat && LAT_PROPERTY_NAMES.includes(v)) {
      res.lat = v;
    }
    if (!res.height && HEIGHT_PROPERTY_NAMES.includes(v)) {
      res.height = v;
    }
    return res;
  }, {} as { lng?: string; lat?: string; height?: string });
};
