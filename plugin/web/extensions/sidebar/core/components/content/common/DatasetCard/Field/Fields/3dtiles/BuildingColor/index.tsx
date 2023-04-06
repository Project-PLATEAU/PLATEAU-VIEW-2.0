import { Checkbox, Radio } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useState } from "react";

import { BaseFieldProps } from "../../types";

import { COLOR_TYPE_CONDITIONS } from "./conditions";
import { LAND_SLIDE_RISK_FIELD, LEGEND_IMAGES } from "./constants";
import useHooks from "./hooks";

const BuildingColor: React.FC<BaseFieldProps<"buildingColor">> = ({
  value,
  dataID,
  editMode,
  onUpdate,
}) => {
  const { options, floods, initialized, independentColorTypes, handleUpdateColorType } = useHooks({
    value,
    dataID,
    onUpdate,
  });

  const [disableFloodRankLegend, setDisableFloodRankLegend] = useState(
    value.disableFloodRankLegend,
  );
  const handleDisableFloodRankLegend = useCallback(() => {
    setDisableFloodRankLegend(e => {
      onUpdate({ ...value, disableFloodRankLegend: !e });
      return !e;
    });
  }, [value, onUpdate]);

  return editMode ? (
    <Checkbox
      style={{ margin: 0 }}
      checked={disableFloodRankLegend}
      onChange={handleDisableFloodRankLegend}>
      <Text>浸水ランクの凡例を非表示にする</Text>
    </Checkbox>
  ) : (
    <>
      <Radio.Group onChange={handleUpdateColorType} value={options.colorType} defaultValue="none">
        <StyledRadio value="none">
          <Label>色分けなし</Label>
        </StyledRadio>
        {initialized && (
          <>
            {independentColorTypes.map(type => (
              <StyledRadio key={type.id} value={type.id}>
                <Label>{type.label}</Label>
              </StyledRadio>
            ))}
            {Object.entries(LAND_SLIDE_RISK_FIELD).map(([, val]) => (
              <StyledRadio key={value.id} value={val.id}>
                <Label>{val.label}</Label>
              </StyledRadio>
            ))}
            {floods.map(flood => (
              <StyledRadio key={flood.id} value={flood.id}>
                <Label>{flood.label}</Label>
              </StyledRadio>
            ))}
          </>
        )}
      </Radio.Group>
      {initialized && (
        <LegendContainer>
          {options.colorType.startsWith("floods") && !disableFloodRankLegend ? (
            <LegendImage src={LEGEND_IMAGES.floods} />
          ) : options.colorType &&
            options.colorType !== "none" &&
            !options.colorType.startsWith("floods") ? (
            <>
              <LegendLabel>凡例</LegendLabel>
              <LegendList>
                {COLOR_TYPE_CONDITIONS[options.colorType as keyof typeof COLOR_TYPE_CONDITIONS].map(
                  cond =>
                    !cond.default && (
                      <LegendItem key={cond.condition}>
                        <ColorBlock color={cond.color} />
                        <LegendText>{cond.label}</LegendText>
                      </LegendItem>
                    ),
                )}
              </LegendList>
            </>
          ) : undefined}
        </LegendContainer>
      )}
    </>
  );
};

export default BuildingColor;

const Text = styled.p`
  margin: 0;
  font-size: 14px;
`;

const StyledRadio = styled(Radio)`
  width: 100%;
  margin-top: 8px;
`;

const Label = styled.span`
  font-size: 14px;
`;

const LegendContainer = styled.div`
  margin-top: 12px;
`;

const LegendImage = styled.img`
  height: auto;
  width: 100%;
`;

const LegendLabel = styled.span`
  color: #000000;
  font-weight: bold;
  font-size: 14px;
  line-height: 1.6;
`;

const LegendList = styled.ul`
  display: flex;
  flex-direction: column;
  gap: 11px 0;
  margin: 12px 0 0 0;
  padding: 0;
`;

const LegendItem = styled.li`
  display: flex;
  align-items: center;
  list-item: none;
`;

const ColorBlock = styled.div<{ color?: string }>`
  width: 30px;
  height: 30px;
  background: ${({ color }) => color};
`;

const LegendText = styled.span`
  color: #000000;
  font-weight: 400;
  font-size: 14px;
  line-height: 1.3;
  margin-left: 12px;
`;
