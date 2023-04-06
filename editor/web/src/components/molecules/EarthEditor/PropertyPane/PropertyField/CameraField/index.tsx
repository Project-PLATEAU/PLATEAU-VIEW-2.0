import React from "react";

import Button from "@reearth/components/atoms/Button";
import Icon from "@reearth/components/atoms/Icon";
import Text from "@reearth/components/atoms/Text";
import { useT } from "@reearth/i18n";
import { styled, css, useTheme } from "@reearth/theme";
import { Camera } from "@reearth/util/value";

import { FieldProps } from "../types";

import useHooks from "./hooks";

export type Props = FieldProps<Camera> & {
  onDelete?: () => void;
  isCapturing?: boolean;
  camera?: Camera;
  onlyPose?: boolean;
  onlyPosition?: boolean;
  onIsCapturingChange?: (isCapturing: boolean) => void;
  onCameraChange?: (camera: Partial<Camera>) => void;
};

const CameraField: React.FC<Props> = ({
  value,
  onChange,
  onDelete,
  disabled,
  isCapturing,
  onIsCapturingChange,
  camera: cameraState,
  onCameraChange,
  onlyPose,
  onlyPosition,
}) => {
  const t = useT();

  const {
    wrapperRef,
    cameraWrapperRef,
    popper,
    camera,
    open,
    openPopup,
    startCapture,
    handleLatChange,
    handleLngChange,
    handleAltitudeChange,
    handleHeadingChange,
    handlePitchChange,
    handleRollChange,
    handleClickCancelButton,
    handleClickSubmitButton,
    jump,
  } = useHooks({
    cameraValue: value,
    onSubmit: onChange,
    isCapturing,
    onIsCapturingChange,
    cameraState,
    onCameraChange,
    disabled,
    onlyPose,
    onlyPosition,
  });
  const theme = useTheme();

  const heading =
    typeof camera?.heading === "number"
      ? Math.round(((camera?.heading * 180) / Math.PI) * 1000) / 1000
      : "";
  const pitch =
    typeof camera?.pitch === "number"
      ? Math.round(((camera?.pitch * 180) / Math.PI) * 1000) / 1000
      : "";
  const roll =
    typeof camera?.roll === "number"
      ? Math.round(((camera?.roll * 180) / Math.PI) * 1000) / 1000
      : "";

  return (
    <Wrapper ref={wrapperRef} onClick={value ? undefined : startCapture} data-camera-popup>
      <CameraWrapper ref={cameraWrapperRef}>
        <StyledText
          size="xs"
          color={value ? theme.properties.contentsFloatText : theme.properties.contentsText}
          onClick={value ? openPopup : undefined}>
          {value ? (onlyPose ? t("Pose Set") : t("Position Set")) : t("Not Set")}
        </StyledText>
        {value ? (
          <StyledIcon icon="bin" size={16} onClick={onDelete} />
        ) : (
          <StyledIcon icon="capture" size={16} onClick={value ? openPopup : undefined} />
        )}
      </CameraWrapper>
      <Popup ref={popper.ref} open={open} style={popper.styles} {...popper.attributes}>
        {!onlyPose && (
          <FormGroup>
            <FormIcon>
              <Icon icon="marker" size={16} color={theme.properties.contentsText} />
            </FormIcon>
            <FormFieldGroup>
              <FormFieldRow>
                <FormWrapper>
                  <Input
                    type="number"
                    value={camera?.lat}
                    step={0.01}
                    readOnly={!isCapturing}
                    onChange={handleLatChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Latitude")}
                  </FloatText>
                </FormWrapper>
                <FormWrapper>
                  <Input
                    type="number"
                    value={camera?.lng}
                    step={0.01}
                    readOnly={!isCapturing}
                    onChange={handleLngChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Longtitude")}
                  </FloatText>
                </FormWrapper>
                <FormWrapper>
                  <Input
                    type="number"
                    value={camera?.height}
                    step={1000}
                    readOnly={!isCapturing}
                    onChange={handleAltitudeChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Altitude")}
                  </FloatText>
                </FormWrapper>
              </FormFieldRow>
            </FormFieldGroup>
          </FormGroup>
        )}
        {!onlyPosition && (
          <FormGroup>
            <FormIcon>
              <Icon icon="camera" size={16} color={theme.properties.contentsText} />
            </FormIcon>
            <FormFieldGroup>
              <FormFieldRow>
                <FormWrapper>
                  <Input
                    type="number"
                    value={heading}
                    step={1}
                    readOnly={!isCapturing}
                    onChange={handleHeadingChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Heading")}
                  </FloatText>
                </FormWrapper>
                <FormWrapper>
                  <Input
                    type="number"
                    value={pitch}
                    step={1}
                    readOnly={!isCapturing}
                    onChange={handlePitchChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Pitch")}
                  </FloatText>
                </FormWrapper>
                <FormWrapper>
                  <Input
                    type="number"
                    value={roll}
                    step={1}
                    readOnly={!isCapturing}
                    onChange={handleRollChange}
                  />
                  <FloatText size="2xs" color={theme.properties.contentsFloatText}>
                    {t("Roll")}
                  </FloatText>
                </FormWrapper>
              </FormFieldRow>
            </FormFieldGroup>
          </FormGroup>
        )}
        <FormGroup>
          {value && !isCapturing && (
            <FormButtonGroup>
              <Button
                buttonType="secondary"
                text={onlyPose ? t("Check Pose") : t("Jump")}
                onClick={jump}
              />
            </FormButtonGroup>
          )}
        </FormGroup>
        <FormGroup>
          <FormButtonGroup>
            <Button buttonType="secondary" text={t("Cancel")} onClick={handleClickCancelButton} />
            {!isCapturing && (
              <Button
                buttonType="primary"
                text={value && onlyPose ? t("Edit Pose") : t("Edit Position")}
                onClick={startCapture}
              />
            )}
            {isCapturing && (
              <Button buttonType="primary" text={t("Capture")} onClick={handleClickSubmitButton} />
            )}
          </FormButtonGroup>
        </FormGroup>
      </Popup>
    </Wrapper>
  );
};

const Wrapper = styled.div`
  width: 100%;
  border: solid 1px ${props => props.theme.properties.border};
`;

const CameraWrapper = styled.div`
  border-radius: 3px;
  display: flex;
  align-items: center;
  width: 100%;
  box-sizing: border-box;
  cursor: pointer;
  user-select: none;
`;

const StyledIcon = styled(Icon)`
  color: ${props => props.theme.properties.contentsText};
  margin-right: 5px;
`;

const StyledText = styled(Text)`
  flex: 1;
  padding: 8px;
  margin: auto;
`;

const Popup = styled.ul<{ open: boolean }>`
  ${({ open }) =>
    !open &&
    css`
      visibility: hidden;
      pointer-events: none;
    `}
  display: flex;
  flex-direction: column;
  width: 286px;
  margin: 0;
  border: solid 1px ${props => props.theme.properties.border};
  border-radius: 5px;
  background: ${props => props.theme.properties.bg};
  box-sizing: border-box;
  padding: 10px 16px;
  z-index: ${props => props.theme.zIndexes.propertyFieldPopup};
`;

const FormGroup = styled.div`
  display: flex;
  align-items: center;
`;

const FormIcon = styled.div`
  margin-top: auto;
  margin-bottom: 20px;
  width: 35.78px;
`;

const FormButtonGroup = styled.div`
  display: flex;
  justify-content: flex-end;
  flex: 1;
`;

const FormFieldGroup = styled.div`
  display: flex;
  flex-direction: column;
`;

const FormFieldRow = styled.div`
  display: flex;
  margin: 5px;
`;

const FormWrapper = styled.div`
  display: flex;
  align-items: center;
  flex: 1;
  flex-direction: column;
  box-sizing: border-box;

  &:not(:last-child) {
    margin-right: 5px;
  }
`;

const Input = styled.input`
  font-size: 11px;
  display: block;
  border: solid 1px ${props => props.theme.properties.border};
  border-radius: 3px;
  background: ${({ theme }) => theme.main.deepBg};
  outline: none;
  color: ${({ theme }) => theme.properties.contentsText};
  width: 100%;
  padding: 5px;
  box-sizing: border-box;
  -moz-appearance: textfield;

  ::-webkit-outer-spin-button,
  ::-webkit-inner-spin-button {
    -webkit-appearance: none;
    margin: 0;
  }
`;

const FloatText = styled(Text)`
  user-select: none;
`;

export default CameraField;
