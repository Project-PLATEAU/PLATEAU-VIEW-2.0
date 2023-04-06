import omit from "lodash/omit";
import { useCallback, useEffect, useRef, useState } from "react";

import { BaseFieldProps } from "../../types";
import { useObservingDataURL } from "../hooks";

import { FILTERING_FIELD_DEFINITION, OptionsState, USE_MIN_FIELD_PROPERTIES } from "./constants";
import { useBuildingFilter } from "./useBuildingFilter";

const useHooks = ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"buildingFilter">, "value" | "dataID" | "onUpdate">) => {
  const [options, setOptions] = useState<OptionsState>(() =>
    value.userSettings
      ? Object.fromEntries(
          Object.entries(FILTERING_FIELD_DEFINITION)
            .map(([k_, v]) => {
              const k = k_ as keyof typeof FILTERING_FIELD_DEFINITION;
              return value.userSettings[k]
                ? [k, { ...v, ...omit(value.userSettings[k], "override") }]
                : undefined;
            })
            .filter(
              (
                f,
              ): f is [
                keyof typeof FILTERING_FIELD_DEFINITION,
                (typeof FILTERING_FIELD_DEFINITION)[keyof typeof FILTERING_FIELD_DEFINITION],
              ] => !!f,
            ),
        )
      : {},
  );
  const url = useObservingDataURL(dataID);

  const handleUpdate = useCallback(
    (property: any) => {
      onUpdate({
        ...value,
        userSettings: {
          height: options.height,
          abovegroundFloor: options.abovegroundFloor,
          basementFloor: options.basementFloor,
          buildingAge: options.buildingAge,
          override: { ["3dtiles"]: property },
        },
      });
    },
    [onUpdate, value, options],
  );

  const handleUpdateOptions = useCallback(
    <P extends keyof OptionsState>(prop: P, v?: Exclude<OptionsState[P], undefined>["value"]) => {
      setOptions(o => {
        return {
          ...o,
          [prop]: {
            ...o[prop],
            value: v,
          },
        };
      });
    },
    [],
  );

  const handleUpdateRange = useCallback(
    (prop: keyof OptionsState) => (value: number | number[]) => {
      if (value && Array.isArray(value)) {
        handleUpdateOptions(prop, value as [from: number, to: number]);
      }
    },
    [handleUpdateOptions],
  );

  const fetchedUrlRef = useRef<string>();
  useEffect(() => {
    const handleFilteringFields = (data: any) => {
      const tempOptions: typeof options = {};
      Object.entries(data?.properties || {}).forEach(([propertyKey, propertyValue]) => {
        Object.entries(FILTERING_FIELD_DEFINITION).forEach(([k_, type]) => {
          const k = k_ as keyof OptionsState;
          if (
            propertyKey === type.featurePropertyName &&
            propertyValue &&
            typeof propertyValue === "object" &&
            Object.keys(propertyValue).length
          ) {
            const customType = (() => {
              const min =
                USE_MIN_FIELD_PROPERTIES.includes(k) &&
                "minimum" in propertyValue &&
                type.min &&
                Number(propertyValue.minimum) >= type.min
                  ? Number(propertyValue.minimum)
                  : type.min;
              const max = type.max;
              const shouldChangeMin =
                options[k]?.min !== min && options[k]?.value[0] === options[k]?.min;
              const shouldChangeMax =
                options[k]?.max !== max && options[k]?.value[1] === options[k]?.max;
              return {
                ...type,
                value: [
                  (shouldChangeMin ? min : options[k]?.value[0]) ?? type.value[0],
                  (shouldChangeMax ? max : options[k]?.value[1]) ?? type.value[1],
                ].filter(v => v !== undefined) as typeof type.value,
                min,
              };
            })();
            tempOptions[k] = customType;
          }
        });
      });
      setOptions(tempOptions);
    };
    const fetchTileset = async () => {
      if (!url || fetchedUrlRef.current === url) {
        return;
      }
      fetchedUrlRef.current = url;
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

  useBuildingFilter({ options, dataID, onUpdate: handleUpdate });

  return {
    options,
    handleUpdateRange,
  };
};

export default useHooks;
