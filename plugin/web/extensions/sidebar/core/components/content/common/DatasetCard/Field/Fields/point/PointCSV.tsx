import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import {
  TextInput,
  Wrapper,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import debounce from "lodash/debounce";
import isEqual from "lodash/isEqual";
import pick from "lodash/pick";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";

import { BaseFieldProps } from "../types";

type LocationNameField = {
  lng?: string;
  lat?: string;
  height?: string;
};

const PointCSV: React.FC<BaseFieldProps<"pointCSV">> = ({ value, editMode, onUpdate }) => {
  const [locationNames, setLocationNames] = useState<LocationNameField>({
    lat: value?.lat,
    lng: value?.lng,
    height: value?.height,
  });
  const updaterRef = useRef<() => void>();
  const debouncedUpdater = useMemo(
    () => debounce(() => updaterRef.current?.(), 500, { maxWait: 1000 }),
    [],
  );

  const isInitializedRef = useRef(false);
  const handleUpdate = useCallback(() => {
    if (isEqual(pick(value, "lng", "lat", "height"), locationNames) && isInitializedRef.current) {
      return;
    }
    isInitializedRef.current = true;
    onUpdate({
      ...value,
      ...locationNames,
      override: {
        data: {
          csv: {
            lngColumn: locationNames.lng,
            latColumn: locationNames.lat,
            heightColumn: locationNames.height,
          },
        },
      },
    });
  }, [value, locationNames, onUpdate]);

  const handleChange = useCallback(
    (prop: keyof LocationNameField) => (e: React.ChangeEvent<HTMLInputElement>) => {
      const text = e.currentTarget.value;
      setLocationNames(v => {
        const next = { ...v, [prop]: text };
        return next;
      });
    },
    [],
  );

  useEffect(() => {
    updaterRef.current = handleUpdate;
    debouncedUpdater();
  }, [handleUpdate, debouncedUpdater]);

  // FIXME(@keiya01): support auto csv field complement
  // useEffect(() => {
  //   const waitReturnedPostMsg = async (e: MessageEvent<any>) => {
  //     if (e.source !== parent) return;
  //     if (e.data.action === "getLocationNamesFromCSVFeatureProperty") {
  //       const locationNames = e.data.locationNames;
  //       setLocationNames(locationNames);
  //       removeEventListener("message", waitReturnedPostMsg);
  //     }
  //   };
  //   addEventListener("message", waitReturnedPostMsg);
  //   postMsg({
  //     action: "getLocationNamesFromCSVFeatureProperty",
  //     payload: {
  //       dataID,
  //     },
  //   });
  // }, [dataID]);

  return editMode ? (
    <Wrapper>
      <Field
        title="経度フィールド名"
        titleWidth={87}
        value={<TextInput value={locationNames.lng} onChange={handleChange("lng")} />}
      />
      <Field
        title="緯度フィールド名"
        titleWidth={87}
        value={<TextInput value={locationNames.lat} onChange={handleChange("lat")} />}
      />
      <Field
        title="高さフィールド名"
        titleWidth={87}
        value={<TextInput value={locationNames.height} onChange={handleChange("height")} />}
      />
    </Wrapper>
  ) : null;
};

export default PointCSV;
