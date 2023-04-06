import { Slider } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { CSSProperties, useMemo } from "react";

import { BaseFieldProps } from "../../types";

import useHooks from "./hooks";

const rangeToText = (range: [from: number, to: number]) => range.join(" ~ ");

const styleProps = {
  trackStyle: [
    {
      backgroundColor: "#00BEBE",
    },
  ] as CSSProperties[],
  handleStyle: [
    {
      border: "2px solid #00BEBE",
    },
    {
      border: "2px solid #00BEBE",
    },
  ] as CSSProperties[],
};

const BuildingFilter: React.FC<BaseFieldProps<"buildingFilter">> = ({
  value,
  dataID,
  editMode,
  onUpdate,
}) => {
  const { options, handleUpdateRange } = useHooks({
    value,
    dataID,
    onUpdate,
  });
  const fields = useMemo(
    () =>
      Object.entries(options)
        .map(([, v]) => v)
        .sort((a, b) => a.order - b.order),
    [options],
  );

  return editMode ? null : (
    <div>
      {fields.map(f =>
        f.value.length === 2 ? (
          <FieldWrapper key={f.id}>
            <LabelWrapper>
              <Label>{f.label}</Label>
              <Range>{rangeToText(f.value)}</Range>
            </LabelWrapper>
            <Slider
              range={true}
              value={f.value}
              max={f.max}
              min={f.min}
              onChange={handleUpdateRange(f.id)}
              {...styleProps}
            />
          </FieldWrapper>
        ) : null,
      )}
    </div>
  );
};

export default BuildingFilter;

const FieldWrapper = styled.div`
  width: 100%;
`;

const LabelWrapper = styled.div`
  width: 100%;
  display: flex;
  align-items: center;
`;

const Label = styled.label`
  display: flex;
  flex: 1;
  font-size: 12px;
`;

const Range = styled.span`
  font-size: 12px;
`;
