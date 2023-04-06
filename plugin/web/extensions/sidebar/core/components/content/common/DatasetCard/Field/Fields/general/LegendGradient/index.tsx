import { Icon, Dropdown, Menu } from "@web/sharedComponents";
import { styled } from "@web/theme";

import { ColorField, TextField } from "../../../common";
import { FieldTitle, FieldValue, FieldWrapper } from "../../../commonComponents";
import { BaseFieldProps, LegendStyleType } from "../../types";

import useHooks from "./hooks";

const LegendGradient: React.FC<BaseFieldProps<"legendGradient">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  const {
    legendGradient,
    legendStyles,
    displayValues,
    handleStyleChange,
    handleStepChange,
    handleStartColorChange,
    handleEndColorChange,
    handleMinValueChange,
    handleMaxValueChange,
  } = useHooks({ value, onUpdate });

  const stylesMenu = (
    <Menu
      items={Object.keys(legendStyles).map(ls => {
        return {
          key: ls,
          label: (
            <p
              style={{ margin: 0 }}
              onClick={() => handleStyleChange(ls as Omit<LegendStyleType, "icon">)}>
              {legendStyles[ls]}
            </p>
          ),
        };
      })}
    />
  );

  return editMode ? (
    <Wrapper>
      <FieldWrapper>
        <FieldTitle width={82}>スタイル</FieldTitle>
        <FieldValue>
          <Dropdown
            overlay={stylesMenu}
            placement="bottom"
            trigger={["click"]}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}>
            <StyledDropdownButton>
              <p style={{ margin: 0 }}>{legendStyles[`${legendGradient.style}`]}</p>
              <StyledIcon icon="arrowDownSimple" size={12} />
            </StyledDropdownButton>
          </Dropdown>
        </FieldValue>
      </FieldWrapper>
      <TextField
        title="Minimum value"
        titleWidth={82}
        defaultValue={legendGradient.min}
        onChange={handleMinValueChange}
      />
      <TextField
        title="Maximum value"
        titleWidth={82}
        defaultValue={legendGradient.max}
        onChange={handleMaxValueChange}
      />
      <ColorField
        title="Start Color"
        titleWidth={82}
        color={legendGradient.startColor}
        onChange={handleStartColorChange}
      />
      <ColorField
        title="End Color"
        titleWidth={82}
        color={legendGradient.endColor}
        onChange={handleEndColorChange}
      />
      <TextField
        title="Step width"
        titleWidth={82}
        defaultValue={legendGradient.step}
        onChange={handleStepChange}
      />
    </Wrapper>
  ) : (
    <Wrapper>
      {displayValues &&
        Object.entries(displayValues).map(([step, color], index) => (
          <FieldWrapper key={index} gap={12}>
            <ColorBlock color={color} legendStyle={legendGradient.style} />
            <Text>{step}</Text>
          </FieldWrapper>
        ))}
    </Wrapper>
  );
};
export default LegendGradient;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const StyledDropdownButton = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  align-content: center;
  padding: 0 16px;
  cursor: pointer;
`;

const StyledIcon = styled(Icon)`
  font-size: 0;
`;

const Text = styled.p`
  margin: 0;
`;

const ColorBlock = styled.div<{
  color?: string;
  legendStyle: Omit<LegendStyleType, "icon">;
}>`
  width: 30px;
  height: ${({ legendStyle }) => (legendStyle == "line" ? "3px" : "30px")};
  background: ${({ color }) => color ?? "#d9d9d9"};
  border-radius: ${({ legendStyle }) =>
    legendStyle
      ? legendStyle == "circle"
        ? "50%"
        : legendStyle == "line"
        ? "5px"
        : "2px"
      : "1px 0 0 1px"};
`;
