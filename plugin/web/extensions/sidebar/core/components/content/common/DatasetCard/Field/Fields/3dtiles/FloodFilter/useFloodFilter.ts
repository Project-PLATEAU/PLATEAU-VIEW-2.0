import debounce from "lodash/debounce";
import { RefObject, useCallback, useEffect, useMemo, useRef } from "react";

import { BaseFieldProps } from "../../types";

import {
  FEATURE_PROPERTY_NAME_RANK_CODE,
  FEATURE_PROPERTY_NAME_RANK_ORG_CODE,
  FilteringField,
} from "./constants";

export const useFloodFilter = ({
  options,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"floodFilter">, "dataID"> & {
  options: FilteringField;
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
  options: FilteringField;
};

const renderTileset = (state: State, onUpdateRef: RefObject<(property: any) => void>) => {
  const updateTileset = () => {
    if (!Object.keys(state.options || {}).length) {
      return;
    }

    const defaultConditionalValue = (prop: string) =>
      `((\${${prop}} === "" || \${${prop}} === null || isNaN(Number(\${${prop}}))) ? 1 : Number(\${${prop}}))`;
    const condition = (
      max: number | undefined,
      range: [from: number, to: number] | undefined,
      conditionalValue: string,
    ) =>
      max === range?.[1]
        ? `${conditionalValue} >= ${range?.[0]}`
        : `${conditionalValue} >= ${range?.[0]} && ${conditionalValue} <= ${range?.[1]}`;
    const conditions = (() => {
      const conditionalRankCode = defaultConditionalValue(FEATURE_PROPERTY_NAME_RANK_CODE);
      const conditionalRankOrgCode = defaultConditionalValue(FEATURE_PROPERTY_NAME_RANK_ORG_CODE);
      return state.options.isOrg
        ? [condition(state.options.max, state.options.value, conditionalRankOrgCode), "true"]
        : [condition(state.options.max, state.options.value, conditionalRankCode), "true"];
    })();
    onUpdateRef.current?.({
      show: {
        expression: {
          conditions: [[...(conditions ? conditions : [])], ["true", "false"]],
        },
      },
    });
  };

  updateTileset();
};
