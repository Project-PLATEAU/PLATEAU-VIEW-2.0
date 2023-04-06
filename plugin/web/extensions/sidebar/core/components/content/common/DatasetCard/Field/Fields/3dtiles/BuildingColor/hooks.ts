import { Radio } from "@web/sharedComponents";
import { ComponentProps, useCallback, useEffect, useState } from "react";

import { BaseFieldProps } from "../../types";
import { useObservingDataURL } from "../hooks";

import { INDEPENDENT_COLOR_TYPE } from "./constants";
import { useBuildingColor } from "./useBuildingColor";

type OptionsState = BaseFieldProps<"buildingColor">["value"]["userSettings"];

type RadioItem = {
  id: string;
  label: string;
  featurePropertyName: string;
  order?: number;
  useOwnData?: boolean;
  floodScale?: number;
};

const useHooks = ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"buildingColor">, "value" | "dataID" | "onUpdate">) => {
  const [options, setOptions] = useState<OptionsState>({
    colorType: value.userSettings?.colorType ?? "none",
  });
  const [independentColorTypes, setIndependentColorTypes] = useState<RadioItem[]>([]);
  const [floods, setFloods] = useState<RadioItem[]>([]);
  const [initialized, setInitialized] = useState(false);
  const url = useObservingDataURL(dataID);

  const handleUpdate = useCallback(
    (property: any) => {
      if (!initialized) return;
      onUpdate({
        ...value,
        userSettings: { ...options, updatedAt: new Date(), override: { "3dtiles": property } },
      });
    },
    [onUpdate, options, value, initialized],
  );

  const handleUpdateOptions = useCallback(
    <P extends keyof OptionsState>(prop: P, v?: OptionsState[P]) => {
      setOptions(o => {
        const next = { ...o, [prop]: v };
        return next;
      });
    },
    [],
  );

  const handleUpdateColorType: Exclude<ComponentProps<typeof Radio>["onChange"], undefined> =
    useCallback(
      e => {
        handleUpdateOptions("colorType", e.target.value);
      },
      [handleUpdateOptions],
    );

  useEffect(() => {
    const handleIndependentColorTypes = (data: any) => {
      const tempTypes: typeof independentColorTypes = [];
      Object.entries(data?.properties || {}).forEach(([k]) => {
        Object.entries(INDEPENDENT_COLOR_TYPE).forEach(([, type]) => {
          if (!type.always && k === type.featurePropertyName) {
            tempTypes.push(type);
          }
        });
      });
      Object.entries(INDEPENDENT_COLOR_TYPE).forEach(([, type]) => {
        if (type.always) {
          tempTypes.push(type);
        }
      });
      tempTypes.sort((a, b) => (a.order && b.order ? a.order - b.order : 0));
      setIndependentColorTypes(tempTypes);
    };
    const handleFloods = (data: any) => {
      const tempFloods: typeof floods = [];
      Object.entries(data?.properties || {}).forEach(([k, v]) => {
        if (k.endsWith("_浸水ランク") && v && typeof v === "object") {
          const useOwnData = !Object.keys(v).length;
          const featurePropertyName = (() => {
            if (!useOwnData) return k;
            return k.split(/[(_（＿]/)[0];
          })();
          const scale = k.match(/L([1,2])/)?.[1];
          tempFloods.push({
            id: `floods-${tempFloods.length}`,
            label: k.replaceAll("_", " "),
            featurePropertyName,
            useOwnData,
            floodScale: scale ? Number(scale) : undefined,
          });
        }
      });
      setFloods(tempFloods);
    };
    const fetchTileset = async () => {
      setInitialized(true);

      if (!url) {
        return;
      }
      const data = await (async () => {
        try {
          return await fetch(url).then(r => r.json());
        } catch (e) {
          console.error(e);
        }
      })();
      handleIndependentColorTypes(data);
      handleFloods(data);

      setInitialized(true);
    };
    fetchTileset();
  }, [dataID, url]);

  useBuildingColor({ options, dataID, floods, initialized, onUpdate: handleUpdate });

  return {
    initialized,
    options,
    independentColorTypes,
    floods,
    handleUpdateColorType,
  };
};

export default useHooks;
