import { Slider } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { CSSProperties } from "react";

import { BaseFieldProps } from "../../types";

import useHooks from "./hooks";

const styleProps = {
  trackStyle: {
    backgroundColor: "#00BEBE",
  } as CSSProperties,
  handleStyle: {
    border: "2px solid #00BEBE",
  } as CSSProperties,
};

const BuildingTransparency: React.FC<BaseFieldProps<"buildingTransparency">> = ({
  value,
  dataID,
  editMode,
  onUpdate,
}) => {
  const { options, handleUpdateNumber } = useHooks({
    value,
    dataID,
    onUpdate,
  });

  return editMode ? null : (
    <FieldWrapper>
      <Slider
        value={options.transparency}
        defaultValue={100}
        max={100}
        onChange={handleUpdateNumber("transparency")}
        {...styleProps}
      />
    </FieldWrapper>
  );
};

export default BuildingTransparency;

const FieldWrapper = styled.div`
  width: 100%;
`;
