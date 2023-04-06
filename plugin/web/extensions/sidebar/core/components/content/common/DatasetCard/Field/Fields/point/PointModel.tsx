import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import {
  TextInput,
  Wrapper,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import { useCallback, useEffect, useState } from "react";

import { BaseFieldProps } from "../types";

const PointModel: React.FC<BaseFieldProps<"pointModel">> = ({ value, editMode, onUpdate }) => {
  const [modelURL, setModelURL] = useState(value.modelURL ?? "");
  const [scale, setImageSize] = useState(value.scale);

  const handleURLUpdate = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setModelURL(e.currentTarget.value);
  }, []);

  const handleScaleUpdate = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    const size = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 1;
    setImageSize(size);
  }, []);

  useEffect(() => {
    if (value.scale === scale && value.modelURL === modelURL) return;
    const timer = setTimeout(() => {
      onUpdate({
        ...value,
        modelURL,
        scale,
        override: { model: { url: modelURL, scale } },
      });
    }, 500);
    return () => {
      clearTimeout(timer);
    };
  }, [modelURL, scale, value, onUpdate]);

  return editMode ? (
    <Wrapper>
      <Field
        title="モデルURL"
        titleWidth={82}
        value={<TextInput defaultValue={modelURL} onChange={handleURLUpdate} />}
      />
      <Field
        title="目盛"
        titleWidth={82}
        value={<TextInput defaultValue={scale} onChange={handleScaleUpdate} />}
      />
    </Wrapper>
  ) : null;
};

export default PointModel;
