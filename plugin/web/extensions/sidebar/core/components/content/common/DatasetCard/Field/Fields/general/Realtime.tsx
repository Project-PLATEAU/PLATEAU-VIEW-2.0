import { Icon, Switch } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useEffect, useRef, useState } from "react";

import { BaseFieldProps } from "../types";

const Realtime: React.FC<BaseFieldProps<"realtime">> = ({ value, editMode, onUpdate }) => {
  const [intervalInSecond, setIntervalValue] = useState<number>(value.updateInterval ?? 30);
  const [timer, setTimer] = useState<string>("00:00:00");
  const [enableUpdate, changeUpdateState] = useState<boolean>(value.userSettings?.enabled ?? true);
  const Ref = useRef<any>(null);

  const getTimeRemaining = useCallback((interval: number) => {
    const seconds = Math.floor(interval % 60);
    const minutes = Math.floor((interval / 60) % 60);
    const hours = Math.floor((interval / 60 / 60) % 24);
    if (interval >= 0) {
      setTimer(
        (hours > 9 ? hours : "0" + hours) +
          ":" +
          (minutes > 9 ? minutes : "0" + minutes) +
          ":" +
          (seconds > 9 ? seconds : "0" + seconds),
      );
    }
  }, []);

  const startTimer = useCallback(
    (interval: number) => {
      if (Ref.current) clearInterval(Ref.current);
      const remainingTime = setInterval(() => {
        if (enableUpdate) {
          getTimeRemaining(interval);
          interval--;
          if (interval === 0) interval = intervalInSecond;
        }
      }, 1000);

      Ref.current = remainingTime;
    },
    [enableUpdate, getTimeRemaining, intervalInSecond],
  );
  const clearTimer = useCallback(() => {
    if (Ref.current) clearInterval(Ref.current);
  }, []);

  const handleChangeupdateState = useCallback(() => {
    changeUpdateState(!enableUpdate);
  }, [enableUpdate]);

  const handleChangeUpdateTime = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const inputValue = parseFloat(e.currentTarget.value);
      if (isNaN(inputValue)) setTimer("00:00:00");
      setIntervalValue(inputValue);
      startTimer(inputValue);
    },
    [startTimer],
  );

  const isInitializedRef = useRef(false);
  const propagateRealTimeToLayer = useCallback(() => {
    if (
      value.updateInterval === intervalInSecond &&
      value.userSettings?.enabled === enableUpdate &&
      isInitializedRef.current
    ) {
      return;
    }

    isInitializedRef.current = true;

    if (enableUpdate) {
      onUpdate({
        ...value,
        updateInterval: intervalInSecond,
        userSettings: {
          enabled: true,
        },
        override: {
          data: {
            updateInterval: intervalInSecond * 1000, // to ms
          },
        },
      });
    } else {
      onUpdate({
        ...value,
        updateInterval: intervalInSecond,
        userSettings: {
          enabled: false,
        },
        override: {
          data: {
            updateInterval: undefined, // to ms
          },
        },
      });
    }
  }, [intervalInSecond, enableUpdate, onUpdate, value]);

  useEffect(() => {
    if (enableUpdate) {
      startTimer(intervalInSecond);
    } else clearTimer();
  }, [clearTimer, enableUpdate, intervalInSecond, startTimer]);

  useEffect(() => {
    propagateRealTimeToLayer();
  }, [propagateRealTimeToLayer, enableUpdate]);

  return editMode ? (
    <Wrapper>
      <Field>
        <FieldTitle>Update time</FieldTitle>
        <FieldValue>
          <TextInput
            type={"number"}
            readOnly={!enableUpdate}
            value={intervalInSecond}
            onChange={handleChangeUpdateTime}
          />
        </FieldValue>
      </Field>
    </Wrapper>
  ) : (
    <Wrapper>
      <Field gap={10}>
        <Icon icon="clock" size={24} />
        <Text>次のデータ更新まであと {timer}</Text>
      </Field>
      <Field gap={10}>
        <SwitchWrapper>
          <Switch checked={enableUpdate} size="small" onChange={handleChangeupdateState} />
          <Text>リアルタイム更新を有効にする</Text>
        </SwitchWrapper>
      </Field>
    </Wrapper>
  );
};

export default Realtime;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Text = styled.p`
  margin: 0;
  size: 14px;
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const TextInput = styled.input`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;
  :focus {
    border: none;
  }
`;
const SwitchWrapper = styled.div`
  display: flex;
  align-items: center;
  gap: 10px;
  margin-top: 12px;
`;
