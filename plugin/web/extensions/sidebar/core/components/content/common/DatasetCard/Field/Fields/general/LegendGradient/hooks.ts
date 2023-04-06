import { generateColorGradient } from "@web/extensions/sidebar/utils/color";
import { useCallback, useEffect, useState } from "react";

import { BaseFieldProps, LegendStyleType } from "../../types";

const legendStyles: { [key: string]: string } = {
  square: "四角",
  circle: "丸",
  line: "線",
};

export default ({
  value,
  onUpdate,
}: Pick<BaseFieldProps<"legendGradient">, "value" | "onUpdate">) => {
  const [legendGradient, updateLegendGradient] = useState(value);
  const [displayValues, setDisplayValues] = useState<{ [step: number]: string }>();

  const handleStyleChange = useCallback(
    (style: Omit<LegendStyleType, "icon">) => {
      updateLegendGradient(lg => {
        const newLegendGradient = {
          ...lg,
          style,
        };
        onUpdate(newLegendGradient);
        return newLegendGradient;
      });
    },
    [onUpdate],
  );

  const handleStartColorChange = useCallback(
    (color: string) => {
      if (color) {
        updateLegendGradient(lg => {
          const newLegendGradient = {
            ...lg,
            startColor: color,
          };
          onUpdate(newLegendGradient);
          return newLegendGradient;
        });
      }
    },
    [onUpdate],
  );

  const handleEndColorChange = useCallback(
    (color: string) => {
      if (color) {
        updateLegendGradient(lg => {
          const newLegendGradient = {
            ...lg,
            endColor: color,
          };
          onUpdate(newLegendGradient);
          return newLegendGradient;
        });
      }
    },
    [onUpdate],
  );

  const handleStepChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const step = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 1;
      updateLegendGradient(lg => {
        const newLegendGradient = {
          ...lg,
          step: step,
        };
        onUpdate(newLegendGradient);
        return newLegendGradient;
      });
    },
    [onUpdate],
  );

  const handleMinValueChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const min = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 1;
      updateLegendGradient(lg => {
        const newLegendGradient = {
          ...lg,
          min,
        };
        onUpdate(newLegendGradient);
        return newLegendGradient;
      });
    },
    [onUpdate],
  );

  const handleMaxValueChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const max = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 1;
      updateLegendGradient(lg => {
        const newLegendGradient = {
          ...lg,
          max,
        };
        onUpdate(newLegendGradient);
        return newLegendGradient;
      });
    },
    [onUpdate],
  );
  const generateDisplayValue = useCallback(
    (min?: number, max?: number, step?: number, startColor?: string, endColor?: string) => {
      if (!min || !max || !step || !startColor || !endColor || min >= max || step >= max) {
        return {};
      }
      const values = [];
      for (let i = min; i <= max; i += step) values.push(i);
      const colors = generateColorGradient(startColor, endColor, values.length);
      const newDisplayValue: { [key: number]: string } = values.reduce(
        (acc: { [key: number]: string }, value, index) => {
          acc[value] = colors[index];
          return acc;
        },
        {},
      );
      return newDisplayValue;
    },
    [],
  );

  useEffect(() => {
    setDisplayValues(
      generateDisplayValue(
        legendGradient.min,
        legendGradient.max,
        legendGradient.step,
        legendGradient.startColor,
        legendGradient.endColor,
      ),
    );
  }, [
    generateDisplayValue,
    legendGradient.endColor,
    legendGradient.max,
    legendGradient.min,
    legendGradient.startColor,
    legendGradient.step,
  ]);

  return {
    legendStyles,
    legendGradient,
    displayValues,
    handleStyleChange,
    handleStepChange,
    handleStartColorChange,
    handleEndColorChange,
    handleMinValueChange,
    handleMaxValueChange,
  };
};
