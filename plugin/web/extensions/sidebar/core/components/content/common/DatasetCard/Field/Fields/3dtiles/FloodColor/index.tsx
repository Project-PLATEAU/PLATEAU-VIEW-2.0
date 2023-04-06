import {
  getRGBAFromString,
  RGBA,
  rgbaToString,
  getOverriddenLayerByDataID,
} from "@web/extensions/sidebar/utils";
import { styled } from "@web/theme";
import { Radio } from "antd";
import { ComponentProps, useCallback, useEffect, useRef, useState } from "react";

import { BaseFieldProps } from "../../types";

import { CONDITIONS, DEFAULT_TRANSPARENCY } from "./conditions";

const FloodColor: React.FC<BaseFieldProps<"floodColor">> = ({
  dataID,
  onUpdate,
  value,
  editMode,
}) => {
  const [colorType, setColorType] = useState<
    BaseFieldProps<"floodColor">["value"]["userSettings"]["colorType"]
  >(value.userSettings?.colorType ?? "water");

  const handleUpdateColorType: Exclude<ComponentProps<typeof Radio>["onChange"], undefined> =
    useCallback(e => {
      setColorType(e.target.value);
    }, []);

  const handleUpdate = useCallback(
    (property: any) => {
      onUpdate({
        ...value,
        userSettings: {
          colorType,
          updatedAt: new Date(),
          override: { "3dtiles": property },
        },
      });
    },
    [onUpdate, value, colorType],
  );

  const onUpdateRef = useRef(handleUpdate);
  useEffect(() => {
    onUpdateRef.current = handleUpdate;
  }, [handleUpdate]);

  useEffect(() => {
    const updateTileset = async () => {
      const overriddenLayer = await getOverriddenLayerByDataID(dataID);

      // We can get transparency from RGBA. Because the color is defined as RGBA.
      const overriddenColor = overriddenLayer?.["3dtiles"]?.color;
      const transparency = getRGBAFromString(
        typeof overriddenColor === "string"
          ? overriddenColor
          : overriddenColor?.expression?.conditions?.[0]?.[1],
      )?.[3];

      const expression = {
        expression: {
          conditions: CONDITIONS[colorType].map(([k, v]: [string, string]) => {
            const rgba = getRGBAFromString(v);
            if (!rgba) {
              return [k, v];
            }
            const composedRGBA = [
              ...rgba.slice(0, -1),
              !transparency || transparency === 1 ? DEFAULT_TRANSPARENCY : transparency,
            ] as RGBA;
            return [k, rgbaToString(composedRGBA)];
          }),
        },
      };

      onUpdateRef.current({
        color: expression,
        colorBlendMode: colorType === "water" ? "highlight" : "replace",
      });
    };
    updateTileset();
  }, [dataID, colorType]);

  return editMode ? null : (
    <Radio.Group onChange={handleUpdateColorType} value={colorType} defaultValue="water">
      <StyledRadio value="water">
        <Label>水面表現</Label>
      </StyledRadio>
      <StyledRadio value="rank">
        <Label>浸水ランク</Label>
      </StyledRadio>
    </Radio.Group>
  );
};

export default FloodColor;

const StyledRadio = styled(Radio)`
  width: 100%;
  margin-top: 8px;
`;

const Label = styled.span`
  font-size: 14px;
`;
