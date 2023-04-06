import { Field } from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/common";
import {
  TextInput,
  Wrapper,
} from "@web/extensions/sidebar/core/components/content/common/DatasetCard/Field/commonComponents";
import Divider from "@web/sharedComponents/Divider";
import { isEqual, pick } from "lodash";
import debounce from "lodash/debounce";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";

import { BaseFieldProps } from "../types";

const formatDateTime = (d: string, t: string) => {
  const date = d
    ?.split(/-|\//)
    ?.map(s => s.padStart(2, "0"))
    ?.join("-");
  const Time = t
    ?.split(/:/)
    ?.map(s => s.padStart(2, "0"))
    ?.join(":");
  const dateStr = [date, Time].filter(s => !!s).join("T");

  try {
    return new Date(dateStr).toISOString();
  } catch {
    return new Date().toISOString();
  }
};

const CurrentTime: React.FC<BaseFieldProps<"currentTime">> = ({
  value,
  editMode,
  onUpdate,
  onSceneUpdate,
}) => {
  const [options, setOptions] = useState({
    currentDate: value.currentDate,
    currentTime: value.currentTime,
    startDate: value.startDate,
    startTime: value.startTime,
    stopDate: value.stopDate,
    stopTime: value.stopTime,
  });
  const updaterRef = useRef<() => void>();
  const debouncedUpdater = useMemo(
    () => debounce(() => updaterRef.current?.(), 500, { maxWait: 1000 }),
    [],
  );

  const currentTimeStr = useMemo(
    () => formatDateTime(options.currentDate, options.currentTime),
    [options.currentDate, options.currentTime],
  );

  const startTimeStr = useMemo(
    () => formatDateTime(options.startDate, options.startTime),
    [options.startDate, options.startTime],
  );

  const stopTimeStr = useMemo(
    () => formatDateTime(options.stopDate, options.stopTime),
    [options.stopDate, options.stopTime],
  );

  const isInitialized = useRef(false);
  const handleUpdate = useCallback(() => {
    if (
      isInitialized.current &&
      isEqual(
        options,
        pick(value, "currentDate", "currentTime", "startDate", "startTime", "stopDate", "stopTime"),
      )
    ) {
      return;
    }
    isInitialized.current = true;

    onUpdate({
      ...value,
      currentDate: options.currentDate,
      currentTime: options.currentTime,
      startDate: options.startDate,
      startTime: options.startTime,
      stopDate: options.stopDate,
      stopTime: options.stopTime,
    });

    onSceneUpdate?.({
      timeline: {
        current: currentTimeStr,
        start: startTimeStr,
        stop: stopTimeStr,
      },
    });
  }, [currentTimeStr, onSceneUpdate, onUpdate, options, startTimeStr, stopTimeStr, value]);

  const handleChange = useCallback(
    (prop: keyof typeof options) => (e: React.ChangeEvent<HTMLInputElement>) => {
      const text = e.currentTarget.value;
      setOptions(v => {
        const next = { ...v, [prop]: text };
        return next;
      });
    },
    [],
  );

  useEffect(() => {
    updaterRef.current = handleUpdate;
    debouncedUpdater();
  }, [handleUpdate, debouncedUpdater]);

  const isStopDisabled = useMemo(
    () => !options.startDate || !options.startTime,
    [options.startDate, options.startTime],
  );

  useEffect(() => {
    if (isStopDisabled && (options.stopDate || options.stopTime)) {
      setOptions(v => {
        return {
          ...v,
          stopDate: "",
          stopTime: "",
        };
      });
    }
  }, [isStopDisabled, options.stopDate, options.stopTime]);

  useEffect(
    () => () => {
      // TODO: Use undefined to reset time.
      // But currently we can not override scene property with undefined.
      const now = Date.now();
      const start = new Date(now - 86400000).toISOString();
      const stop = new Date(now).toISOString();
      onSceneUpdate?.({
        timeline: {
          current: start,
          start,
          stop,
        },
      });
    },
    [onSceneUpdate],
  );

  return editMode ? (
    <Wrapper>
      <Divider plain>Current Time</Divider>
      <Field
        title="日付"
        titleWidth={87}
        value={
          <TextInput
            value={options.currentDate}
            placeholder="YYYY-MM-DD"
            onChange={handleChange("currentDate")}
          />
        }
      />
      <Field
        title="時間"
        titleWidth={87}
        value={
          <TextInput
            value={options.currentTime}
            placeholder="HH:mm:ss.sss"
            onChange={handleChange("currentTime")}
          />
        }
      />
      <Divider plain>Start Time</Divider>
      <Field
        title="日付"
        titleWidth={87}
        value={
          <TextInput
            value={options.startDate}
            placeholder="YYYY-MM-DD"
            onChange={handleChange("startDate")}
          />
        }
      />
      <Field
        title="時間"
        titleWidth={87}
        value={
          <TextInput
            value={options.startTime}
            placeholder="HH:mm:ss.sss"
            onChange={handleChange("startTime")}
          />
        }
      />
      <Divider plain>Stop Time</Divider>
      <Field
        title="日付"
        titleWidth={87}
        value={
          <TextInput
            value={options.stopDate}
            placeholder="YYYY-MM-DD"
            disabled={isStopDisabled}
            onChange={handleChange("stopDate")}
          />
        }
      />
      <Field
        title="時間"
        titleWidth={87}
        value={
          <TextInput
            value={options.stopTime}
            placeholder="HH:mm:ss.sss"
            disabled={isStopDisabled}
            onChange={handleChange("stopTime")}
          />
        }
      />
    </Wrapper>
  ) : null;
};

export default CurrentTime;
