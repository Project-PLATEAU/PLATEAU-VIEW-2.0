import {
  TextField,
  ColorField,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import { Wrapper } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { generateColorGradient } from "@web/extensions/sidebar/utils/color";
import { isEqual, omit } from "lodash";
import { useCallback, useEffect, useMemo, useState } from "react";

import { BaseFieldProps } from "../types";

const generateConditions = (
  field?: string,
  min?: number,
  max?: number,
  step?: number,
  startColor?: string,
  endColor?: string,
) => {
  if (!field || !min || !max || !step || !startColor || !endColor || min >= max || step >= max) {
    return [];
  }

  const values = [];
  for (let i = min; i <= max; i += step) values.push(i);
  const colors = generateColorGradient(startColor, endColor, values.length);
  const fieldName = "${" + String(field) + "}";
  return values
    .map((value, index) => [`${fieldName} >= ${String(value)}`, `'${String(colors[index])}'`])
    .reverse();
};

const PointColorGradient: React.FC<BaseFieldProps<"pointColorGradient">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  const [colorGradient, setColorGradient] = useState(value);

  const handleFieldChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      setColorGradient({ ...colorGradient, field: e.target.value });
    },
    [colorGradient],
  );

  const handleMinValueUpdate = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const min = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 1;
      setColorGradient({ ...colorGradient, min });
    },
    [colorGradient],
  );

  const handleMaxValueUpdate = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const max = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 1;
      setColorGradient({ ...colorGradient, max });
    },
    [colorGradient],
  );

  const handleStartColorUpdate = useCallback(
    (color: string) => {
      if (color) {
        setColorGradient({ ...colorGradient, startColor: color });
      }
    },
    [colorGradient],
  );

  const handleEndColorUpdate = useCallback(
    (color: string) => {
      if (color) {
        setColorGradient({ ...colorGradient, endColor: color });
      }
    },
    [colorGradient],
  );

  const handleStepUpdate = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const step = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 1;
      setColorGradient({ ...colorGradient, step });
    },
    [colorGradient],
  );

  const conditions = useMemo(() => {
    const { field, min, max, step, startColor, endColor } = colorGradient;
    return generateConditions(field, min, max, step, startColor, endColor);
  }, [colorGradient]);

  useEffect(() => {
    if (isEqual(omit(colorGradient, "override"), omit(value, "override"))) return;
    const timer = setTimeout(() => {
      onUpdate({
        ...value,
        field: colorGradient.field,
        min: colorGradient.min,
        max: colorGradient.max,
        step: colorGradient.step,
        startColor: colorGradient.startColor,
        endColor: colorGradient.endColor,
        override: {
          marker: {
            style: "point",
            pointColor: {
              expression: {
                conditions,
              },
            },
          },
        },
      });
    }, 500);
    return () => {
      clearTimeout(timer);
    };
  }, [value, onUpdate, colorGradient, conditions]);

  return editMode ? (
    <Wrapper>
      <TextField
        title="Field"
        titleWidth={82}
        defaultValue={colorGradient.field}
        onChange={handleFieldChange}
      />
      <TextField
        title="Min Value"
        titleWidth={82}
        defaultValue={colorGradient.min}
        onChange={handleMinValueUpdate}
      />
      <TextField
        title="Max Value"
        titleWidth={82}
        defaultValue={colorGradient.max}
        onChange={handleMaxValueUpdate}
      />
      <ColorField
        title="Start Color"
        titleWidth={82}
        color={colorGradient.startColor}
        onChange={handleStartColorUpdate}
      />
      <ColorField
        title="End Color"
        titleWidth={82}
        color={colorGradient.endColor}
        onChange={handleEndColorUpdate}
      />
      <TextField
        title="Step"
        titleWidth={82}
        defaultValue={colorGradient.step}
        onChange={handleStepUpdate}
      />
    </Wrapper>
  ) : null;
};

export default PointColorGradient;
