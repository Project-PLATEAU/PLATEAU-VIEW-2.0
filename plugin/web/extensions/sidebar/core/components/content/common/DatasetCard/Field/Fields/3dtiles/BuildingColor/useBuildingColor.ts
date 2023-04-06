import {
  getRGBAFromString,
  RGBA,
  rgbaToString,
  getOverriddenLayerByDataID,
} from "@web/extensions/sidebar/utils";
import debounce from "lodash/debounce";
import pick from "lodash/pick";
import { MutableRefObject, RefObject, useCallback, useEffect, useMemo, useRef } from "react";

import { BaseFieldProps } from "../../types";

import { COLOR_TYPE_CONDITIONS, makeSelectedFloodCondition } from "./conditions";
import { INDEPENDENT_COLOR_TYPE } from "./constants";

export const useBuildingColor = ({
  options,
  initialized,
  floods,
  dataID,
  onUpdate,
}: Pick<BaseFieldProps<"buildingColor">, "dataID"> & {
  initialized: boolean;
  options: BaseFieldProps<"buildingColor">["value"]["userSettings"];
  floods: {
    id: string;
    label: string;
    featurePropertyName: string;
    useOwnData?: boolean;
    floodScale?: number;
  }[];
  onUpdate: (property: any) => void;
}) => {
  const renderRef = useRef<() => void>();
  const debouncedRender = useMemo(
    () => debounce(() => renderRef.current?.(), 100, { maxWait: 300 }),
    [],
  );

  const isInitializedRef = useRef(false);

  const onUpdateRef = useRef(onUpdate);
  useEffect(() => {
    onUpdateRef.current = onUpdate;
  }, [onUpdate]);

  const render = useCallback(async () => {
    renderTileset(
      {
        dataID,
        floods,
        colorType: options.colorType,
        isInitializedRef,
      },
      onUpdateRef,
    );
  }, [options.colorType, dataID, floods]);

  useEffect(() => {
    if (!initialized) {
      return;
    }
    renderRef.current = render;
    debouncedRender();
  }, [render, debouncedRender, initialized]);
};

export type State = {
  dataID: string | undefined;
  floods: {
    id: string;
    label: string;
    featurePropertyName: string;
    useOwnData?: boolean;
    floodScale?: number;
  }[];
  colorType: string;
  isInitializedRef: MutableRefObject<boolean>;
};

const renderTileset = (state: State, onUpdateRef: RefObject<(property: any) => void>) => {
  const updateTileset = async () => {
    const overriddenLayer = await getOverriddenLayerByDataID(state.dataID);

    // We can get transparency from RGBA. Because the color is defined as RGBA.
    const overriddenColor = overriddenLayer?.["3dtiles"]?.color;
    const transparency =
      getRGBAFromString(
        typeof overriddenColor === "string"
          ? overriddenColor
          : overriddenColor?.expression?.conditions?.[0]?.[1],
      )?.[3] || 1;

    const expression = {
      expression: {
        conditions: (
          COLOR_TYPE_CONDITIONS[
            (state.colorType as keyof typeof INDEPENDENT_COLOR_TYPE) || "none"
          ]?.map((cond): [string, string] => [cond.condition, cond.color]) ??
          makeSelectedFloodCondition(
            pick(
              state.floods?.find(f => f.id === state.colorType),
              "featurePropertyName",
              "useOwnData",
              "floodScale",
            ) as Pick<
              (typeof state.floods)[number],
              "featurePropertyName" | "useOwnData" | "floodScale"
            >,
          )
        ).map(([k, v]: [string, string]) => {
          const rgba = getRGBAFromString(v);
          if (!rgba) {
            return [k, v];
          }
          const composedRGBA = [...rgba.slice(0, -1), transparency || rgba[3]] as RGBA;
          return [k, rgbaToString(composedRGBA)];
        }),
      },
    };

    if (!state.isInitializedRef.current) {
      state.isInitializedRef.current = true;
    } else {
      onUpdateRef.current?.({
        colorBlendMode: "replace",
        color: expression,
      });
    }
  };

  updateTileset();
};
