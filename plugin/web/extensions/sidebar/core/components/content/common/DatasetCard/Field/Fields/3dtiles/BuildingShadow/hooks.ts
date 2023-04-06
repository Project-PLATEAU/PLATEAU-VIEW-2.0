import { useCallback, useState } from "react";

import { BaseFieldProps } from "../../types";

import { useBuildingShadow } from "./useBuildingShadow";

type OptionsState = BaseFieldProps<"buildingShadow">["value"]["userSettings"];

const useHooks = ({
  value,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"buildingShadow">, "value" | "dataID" | "onUpdate">) => {
  const [options, setOptions] = useState<OptionsState>({
    shadow: value.userSettings?.shadow ?? "disabled",
  });

  const handleUpdate = useCallback(
    (property: any) => {
      onUpdate({
        ...value,
        userSettings: {
          ...options,
          override: { ["3dtiles"]: property },
        },
      });
    },
    [onUpdate, value, options],
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

  const handleUpdateSelect = useCallback(
    (prop: keyof OptionsState) => (value: any) => {
      handleUpdateOptions(prop, value as OptionsState["shadow"]);
    },
    [handleUpdateOptions],
  );

  useBuildingShadow({ options, dataID, onUpdate: handleUpdate });

  return {
    options,
    handleUpdateSelect,
  };
};

export default useHooks;
