import debounce from "lodash/debounce";
import { RefObject, useCallback, useEffect, useMemo, useRef } from "react";

import { BaseFieldProps } from "../../types";

import { FILTERING_FIELD_DEFINITION, OptionsState } from "./constants";

export const useBuildingFilter = ({
  options,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"buildingFilter">, "dataID"> & {
  options: OptionsState;
  onUpdate: (property: any) => void;
}) => {
  const renderRef = useRef<() => void>();
  const debouncedRender = useMemo(
    () => debounce(() => renderRef.current?.(), 100, { maxWait: 300 }),
    [],
  );
  const initializedRef = useRef(false);

  const onUpdateRef = useRef(onUpdate);
  useEffect(() => {
    onUpdateRef.current = onUpdate;
  }, [onUpdate]);

  const render = useCallback(async () => {
    renderTileset(
      {
        dataID,
        options,
      },
      onUpdateRef,
    );
  }, [options, dataID]);

  useEffect(() => {
    if (!initializedRef.current) return;
    renderRef.current = render;
    debouncedRender();
  }, [render, debouncedRender]);

  useEffect(() => {
    initializedRef.current = true;
  }, []);
};

export type State = {
  dataID: string | undefined;
  options: OptionsState;
};

const renderTileset = (state: State, onUpdateRef: RefObject<(property: any) => void>) => {
  const updateTileset = () => {
    if (!Object.keys(state.options || {}).length) {
      return;
    }

    const defaultConditionalValue = (prop: string, startValue: number) =>
      `((\${${prop}} === "" || \${${prop}} === null || isNaN(Number(\${${prop}}))) ? ${startValue} : Number(\${${prop}}))`;
    const condition = (
      featurePropertyName: string,
      max: number,
      min: number | undefined,
      range: [from: number, to: number] | undefined,
      conditionalValue: string,
    ) =>
      featurePropertyName === FILTERING_FIELD_DEFINITION.buildingAge.featurePropertyName &&
      min &&
      min === range?.[0] &&
      max === range?.[1]
        ? `true`
        : min === range?.[0]
        ? `${conditionalValue} <= ${range?.[1]}`
        : max === range?.[1]
        ? `${conditionalValue} >= ${range?.[0]}`
        : `${conditionalValue} >= ${range?.[0]} && ${conditionalValue} <= ${range?.[1]}`;
    const conditions = Object.entries(state.options || {}).reduce((res, [, v]) => {
      const conditionalValue = defaultConditionalValue(v.featurePropertyName, v.min ?? 0);
      const conditionDef = condition(
        v.featurePropertyName,
        v.max,
        v.min,
        v.value,
        conditionalValue,
      );
      return `${res ? `${res} && ` : ""}${conditionDef}`;
    }, "");
    onUpdateRef.current?.({
      show: {
        expression: {
          conditions: [...(conditions ? [[conditions, "true"]] : []), ["true", "false"]],
        },
      },
    });
  };

  updateTileset();
};
