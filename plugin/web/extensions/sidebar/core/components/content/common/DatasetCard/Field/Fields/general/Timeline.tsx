import { Checkbox } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { debounce } from "lodash";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";

import { Field } from "../../common";
import { TextInput } from "../../commonComponents";
import { BaseFieldProps } from "../types";

const Timeline: React.FC<BaseFieldProps<"timeline">> = ({ value, editMode, onUpdate }) => {
  const [timeFieldName, setTimeFieldName] = useState(value.timeFieldName);
  const [timeBasedDisplay, setTimeBasedDisplay] = useState(
    value.userSettings?.timeBasedDisplay ?? true,
  );
  const updaterRef = useRef<() => void>();
  const debouncedUpdater = useMemo(
    () => debounce(() => updaterRef.current?.(), 500, { maxWait: 1000 }),
    [],
  );
  const shouldUpdate = useRef(true);

  const handleUpdate = useCallback(() => {
    if (timeBasedDisplay) {
      onUpdate({
        ...value,
        timeFieldName,
        override: {
          data: {
            time: {
              property: timeFieldName,
              interval: 86400000,
            },
          },
        },
      });
    } else {
      onUpdate({
        ...value,
        timeFieldName,
        override: {
          data: {
            time: undefined,
          },
        },
      });
    }
  }, [timeFieldName, timeBasedDisplay, onUpdate, value]);

  const handleTimeFieldName = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    const text = e.currentTarget.value;
    setTimeFieldName(text);
  }, []);

  useEffect(() => {
    updaterRef.current = handleUpdate;
    if (timeFieldName !== value.timeFieldName || shouldUpdate.current) {
      debouncedUpdater();
      shouldUpdate.current = false;
    }
  }, [handleUpdate, debouncedUpdater, timeFieldName, value.timeFieldName]);

  const handleTimeBasedDisplay = useCallback(() => {
    setTimeBasedDisplay(!timeBasedDisplay);
    onUpdate({
      ...value,
      userSettings: {
        timeBasedDisplay: !timeBasedDisplay,
      },
    });
    shouldUpdate.current = true;
  }, [value, timeBasedDisplay, onUpdate]);

  return editMode ? (
    <Field
      title="時間フィールド名"
      titleWidth={87}
      value={<TextInput value={timeFieldName} onChange={handleTimeFieldName} />}
    />
  ) : (
    <div>
      <Checkbox style={{ margin: 0 }} checked={timeBasedDisplay} onChange={handleTimeBasedDisplay}>
        <Text>時刻ベースのデータを表示</Text>
      </Checkbox>
    </div>
  );
};

export default Timeline;

const Text = styled.p`
  margin: 0;
  size: 14px;
`;
