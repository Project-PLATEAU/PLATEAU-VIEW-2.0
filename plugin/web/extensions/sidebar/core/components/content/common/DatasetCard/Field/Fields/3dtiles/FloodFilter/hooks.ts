import omit from "lodash/omit";
import { useCallback, useEffect, useRef, useState } from "react";

import { BaseFieldProps } from "../../types";
import { useObservingDataURL } from "../hooks";

import {
  FEATURE_PROPERTY_NAME_RANK_CODE,
  FEATURE_PROPERTY_NAME_RANK_ORG_CODE,
  FilteringField,
} from "./constants";
import { useFloodFilter } from "./useFloodFilter";

const useHooks = ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"floodFilter">, "value" | "dataID" | "onUpdate" | "configData">) => {
  const [options, setOptions] = useState<FilteringField>({
    ...omit(value.userSettings, "override"),
  });
  const url = useObservingDataURL(dataID);

  const handleUpdate = useCallback(
    (property: any) => {
      onUpdate({
        ...value,
        userSettings: {
          ...value.userSettings,
          ...options,
          override: { ["3dtiles"]: property },
        },
      });
    },
    [onUpdate, value, options],
  );

  const handleUpdateRange = useCallback((v: number | number[]) => {
    if (v && Array.isArray(v)) {
      const range = v as [from: number, to: number];
      setOptions(o => {
        return {
          ...o,
          value: range,
        };
      });
    }
  }, []);

  const fetchedURLRef = useRef<string>();
  useEffect(() => {
    const handleFilteringFields = (data: any) => {
      let tempOptions: typeof options = {};
      Object.entries(data?.properties || {}).forEach(([propertyKey, propertyValue]) => {
        if (
          [FEATURE_PROPERTY_NAME_RANK_CODE, FEATURE_PROPERTY_NAME_RANK_ORG_CODE].includes(
            propertyKey,
          ) &&
          propertyValue &&
          typeof propertyValue === "object" &&
          Object.keys(propertyValue).length
        ) {
          const obj = propertyValue as any;
          const min = obj.minimum;
          const max = obj.maximum;
          const shouldChangeMin = options.min !== min && options.value?.[0] === options.min;
          const shouldChangeMax = options.max !== max && options.value?.[1] === options.max;
          tempOptions = {
            min,
            max,
            value: [
              (shouldChangeMin ? min : options.value?.[0]) ?? min,
              (shouldChangeMax ? max : options.value?.[1]) ?? max,
            ].filter(v => v !== undefined) as typeof options.value,
            isOrg: propertyKey.includes(FEATURE_PROPERTY_NAME_RANK_ORG_CODE),
          };
        }
      });
      setOptions(tempOptions);
    };
    const fetchTileset = async () => {
      if (!url || fetchedURLRef.current === url) {
        return;
      }
      fetchedURLRef.current = url;
      const data = await (async () => {
        try {
          return await fetch(url).then(r => r.json());
        } catch (e) {
          console.error(e);
        }
      })();
      handleFilteringFields(data);
    };
    fetchTileset();
  }, [dataID, url, options]);

  useFloodFilter({ options, dataID, onUpdate: handleUpdate });

  return {
    options,
    handleUpdateRange,
  };
};

export default useHooks;
