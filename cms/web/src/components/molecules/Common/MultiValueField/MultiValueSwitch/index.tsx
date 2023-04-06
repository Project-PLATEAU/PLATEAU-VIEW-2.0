import styled from "@emotion/styled";
import { useCallback, useEffect } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Icon from "@reearth-cms/components/atoms/Icon";
import Switch from "@reearth-cms/components/atoms/Switch";
import { useT } from "@reearth-cms/i18n";

import { moveItemInArray } from "../moveItemArray";

type Props = {
  className?: string;
  checked?: boolean[];
  onChange?: (value: (string | number | boolean)[]) => void;
  disabled?: boolean;
};

const MultiValueSwitch: React.FC<Props> = ({ className, checked = [], onChange, disabled }) => {
  const t = useT();
  const handleInput = useCallback(
    (e: boolean, id: number) => {
      onChange?.(checked?.map((valueItem, index) => (index === id ? e : valueItem)));
    },
    [onChange, checked],
  );

  useEffect(() => {
    if (typeof checked === "string" || typeof checked === "boolean") onChange?.([checked]);
    else if (!checked) onChange?.([]);
  }, [onChange, checked]);

  const handleInputDelete = useCallback(
    (key: number) => {
      onChange?.(
        checked.filter((_, index) => {
          return index !== key;
        }),
      );
    },
    [onChange, checked],
  );

  return (
    <div className={className}>
      {Array.isArray(checked) &&
        checked?.map((valueItem, key) => (
          <FieldWrapper key={key}>
            {!disabled && (
              <>
                <FieldButton
                  type="link"
                  icon={<Icon icon="arrowUp" />}
                  onClick={() => onChange?.(moveItemInArray(checked, key, key - 1))}
                  disabled={key === 0}
                />
                <FieldButton
                  type="link"
                  icon={<Icon icon="arrowDown" />}
                  onClick={() => onChange?.(moveItemInArray(checked, key, key + 1))}
                  disabled={key === checked.length - 1}
                />
              </>
            )}
            <Switch onChange={(e: boolean) => handleInput(e, key)} checked={valueItem} />
            <FlexSpace />
            {!disabled && (
              <FieldButton
                type="link"
                icon={<Icon icon="delete" />}
                onClick={() => handleInputDelete(key)}
              />
            )}
          </FieldWrapper>
        ))}
      {!disabled && (
        <Button
          icon={<Icon icon="plus" />}
          type="primary"
          onClick={() => {
            if (!checked) checked = [];
            onChange?.([...checked, false]);
          }}>
          {t("New")}
        </Button>
      )}
    </div>
  );
};

export default MultiValueSwitch;

const FieldWrapper = styled.div`
  display: flex;
  align-items: center;
  margin: 8px 0;
`;

const FieldButton = styled(Button)`
  color: #000000d9;
  margin-top: 4px;
`;

const FlexSpace = styled.div`
  flex: 1;
`;
