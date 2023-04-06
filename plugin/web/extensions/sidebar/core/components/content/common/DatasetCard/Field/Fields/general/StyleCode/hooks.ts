import JSON5 from "json5";
import { useCallback, useState } from "react";

import { BaseFieldProps } from "../../types";

export default ({ value, onUpdate }: Pick<BaseFieldProps<"styleCode">, "value" | "onUpdate">) => {
  const [code, editCode] = useState(value.src);

  const onApply = useCallback(() => {
    const styleObject = JSON5.parse(code);
    onUpdate({
      ...value,
      src: code,
      override: styleObject,
    });
  }, [onUpdate, code, value]);

  const onEdit = useCallback((e: React.ChangeEvent<HTMLTextAreaElement>) => {
    const newValue = e.target.value;
    editCode(newValue);
  }, []);

  return {
    code,
    editCode,
    onApply,
    onEdit,
  };
};
