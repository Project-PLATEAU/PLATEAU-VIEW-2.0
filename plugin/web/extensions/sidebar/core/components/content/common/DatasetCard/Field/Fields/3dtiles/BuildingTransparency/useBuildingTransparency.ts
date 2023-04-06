import {
  getTransparencyExpression,
  getOverriddenLayerByDataID,
} from "@web/extensions/sidebar/utils";
import debounce from "lodash/debounce";
import { MutableRefObject, RefObject, useCallback, useEffect, useMemo, useRef } from "react";

import { BaseFieldProps } from "../../types";

export const useBuildingTransparency = ({
  options,
  dataID,
  onUpdate,
  onChangeTransparency,
}: Pick<BaseFieldProps<"buildingTransparency">, "dataID"> & {
  options: BaseFieldProps<"buildingTransparency">["value"]["userSettings"];
  onUpdate: (property: any) => void;
  onChangeTransparency: (transparency: number) => void;
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

  const render = useCallback(() => {
    renderTileset(
      {
        dataID,
        transparency: options.transparency,
        isInitializedRef,
      },
      onUpdateRef,
      onChangeTransparency,
    );
  }, [options.transparency, dataID, onChangeTransparency]);

  useEffect(() => {
    renderRef.current = render;
    debouncedRender();
  }, [render, debouncedRender]);
};

export type State = {
  dataID: string | undefined;
  transparency: number;
  isInitializedRef: MutableRefObject<boolean>;
};

const renderTileset = (
  state: State,
  onUpdateRef: RefObject<(property: any) => void>,
  onChangeTransparency: (transparency: number) => void,
) => {
  const updateTileset = async () => {
    const overriddenLayer = await getOverriddenLayerByDataID(state.dataID);
    const transparency = (state.transparency ?? 100) / 100;
    const { expression, updatedTransparency } = getTransparencyExpression(
      overriddenLayer,
      transparency,
      !state.isInitializedRef.current,
    );

    if (!state.isInitializedRef.current) {
      onChangeTransparency(updatedTransparency * 100);
      state.isInitializedRef.current = true;
    } else {
      onUpdateRef.current?.({
        color: expression,
      });
    }
  };

  updateTileset();
};
