import { useCallback, useState } from "react";

import { BaseFieldProps } from "../../types";

import { useClippingBox } from "./useClippingBox";

type OptionsState = Omit<BaseFieldProps<"clipping">["value"], "id" | "group" | "type">;

const useHooks = ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"clipping">, "value" | "dataID" | "onUpdate">) => {
  const [options, setOptions] = useState<OptionsState["userSettings"]>({
    enabled: !!value.userSettings?.enabled,
    show: !!value.userSettings?.show || !!value.userSettings?.enabled,
    aboveGroundOnly: !!value.userSettings?.aboveGroundOnly || !!value.userSettings?.enabled,
    direction: value.userSettings?.direction ?? "inside",
  });

  const handleUpdate = useCallback(
    (tilesetProperty: any, boxProperty: any) => {
      onUpdate({
        ...value,
        userSettings: {
          ...options,
          override: { ["3dtiles"]: tilesetProperty, box: boxProperty },
        },
      });
    },
    [onUpdate, value, options],
  );

  const handleUpdateOptions = useCallback(
    <P extends keyof OptionsState["userSettings"]>(
      prop: P,
      v?: OptionsState["userSettings"][P],
    ) => {
      setOptions(o => {
        const next = { ...o, [prop]: v ?? !o[prop] };
        return next;
      });
    },
    [],
  );

  const handleUpdateBool = useCallback(
    (prop: keyof OptionsState["userSettings"]) => (value: boolean) => {
      if (prop === "enabled") {
        setOptions({
          ...options,
          [prop]: value,
          show: value,
          aboveGroundOnly: value,
          direction: "inside",
        });
      } else {
        setOptions({
          ...options,
          [prop]: value,
        });
      }
    },
    [options],
  );

  const handleUpdateSelect = useCallback(
    (prop: keyof OptionsState["userSettings"]) => (value: unknown) => {
      handleUpdateOptions(prop, value as OptionsState["userSettings"]["direction"]);
    },
    [handleUpdateOptions],
  );

  useClippingBox({ options, dataID, onUpdate: handleUpdate });

  return {
    options,
    handleUpdateBool,
    handleUpdateSelect,
  };
};

export default useHooks;
