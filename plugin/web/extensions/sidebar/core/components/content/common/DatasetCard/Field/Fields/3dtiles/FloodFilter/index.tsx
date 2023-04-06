import { Slider } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { CSSProperties } from "react";

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

const FloodFilter: React.FC<BaseFieldProps<"floodFilter">> = ({
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

  return editMode || (options.value?.length ?? 0) !== 2 ? null : (
    <div>
      <FieldWrapper>
        <LabelWrapper>
          <Label>浸水ランク</Label>
          <Range>{rangeToText(options.value || [1, 1])}</Range>
        </LabelWrapper>
        <Slider
          range={true}
          value={options.value}
          max={options.max || 1}
          min={options.min || 1}
          onChange={handleUpdateRange}
          {...styleProps}
        />
      </FieldWrapper>
    </div>
  );
};

export default FloodFilter;

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
