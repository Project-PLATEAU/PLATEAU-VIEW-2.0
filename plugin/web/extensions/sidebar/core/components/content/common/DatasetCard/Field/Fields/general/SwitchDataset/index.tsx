import { Icon, Dropdown, Menu, Radio } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useEffect, useRef, useState } from "react";

import { BaseFieldProps, ConfigData } from "../../types";

type UIStyles = "dropdown" | "radio";

const uiStyles: { [key: string]: string } = {
  dropdown: "ドロップダウン",
  radio: "ラジオ",
};

const SwitchDataset: React.FC<BaseFieldProps<"switchDataset">> = ({
  value,
  editMode,
  configData,
  onUpdate,
  onCurrentDatasetUpdate,
}) => {
  const initialized = useRef(false);
  const [selectedStyle, selectStyle] = useState(value.uiStyle ?? "dropdown");
  const [selectedDataset, selectDataset] = useState(
    value.userSettings?.selected ?? configData?.[0],
  );

  const styleOptions = (
    <Menu
      selectable
      selectedKeys={[selectedStyle]}
      items={Object.keys(uiStyles).map(key => {
        return {
          key: key,
          label: (
            <p style={{ margin: 0 }} onClick={() => handleStyleChange(key as UIStyles)}>
              {uiStyles[key]}
            </p>
          ),
        };
      })}
    />
  );

  const datasetOptions = (
    <Menu
      selectable
      selectedKeys={selectedDataset ? [selectedDataset.name] : undefined}
      items={configData?.map(d => {
        return {
          key: d.name,
          label: (
            <p style={{ margin: 0 }} onClick={() => handleDatasetSelect(d)}>
              {d.name}
            </p>
          ),
        };
      })}
    />
  );

  const handleStyleChange = useCallback((style: UIStyles) => {
    selectStyle(style);
  }, []);

  const handleDatasetSelect = useCallback((dataset: ConfigData) => {
    selectDataset(dataset);
  }, []);

  useEffect(() => {
    if (!initialized.current) {
      initialized.current = true;
      return;
    }

    if (selectedDataset === value.userSettings?.selected && selectedStyle === value.uiStyle) return;
    onUpdate({
      ...value,
      uiStyle: selectedStyle,
      userSettings: {
        selected: selectedDataset,
        override: {
          data: {
            url: selectedDataset?.url,
            type: selectedDataset?.type.toLowerCase().replace(/\s/g, ""),
            layers: selectedDataset?.layer,
            time: {
              updateClockOnLoad: true,
            },
          },
        },
      },
      cleanseOverride: {
        data: {
          url: configData?.[0].url,
          time: {
            updateClockOnLoad: false,
          },
        },
      },
    });
    onCurrentDatasetUpdate?.(selectedDataset);
  }, [selectedDataset, selectedStyle, configData, value, onUpdate, onCurrentDatasetUpdate]);

  return editMode ? (
    <Wrapper>
      <Field>
        <FieldTitle>UIスタイル</FieldTitle>
        <FieldValue>
          <Dropdown
            overlay={styleOptions}
            placement="bottom"
            trigger={["click"]}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}>
            <StyledDropdownButton>
              <p style={{ margin: 0 }}>{uiStyles[selectedStyle]}</p>
              <StyledIcon icon="arrowDownSimple" size={12} />
            </StyledDropdownButton>
          </Dropdown>
        </FieldValue>
      </Field>
    </Wrapper>
  ) : (
    <Wrapper>
      <Field>
        {selectedDataset ? (
          value.uiStyle === "radio" && configData ? (
            <Radio.Group
              onChange={e =>
                handleDatasetSelect(
                  configData.find(cd => cd.name === e.target.value) ?? selectedDataset,
                )
              }
              value={selectedDataset.name}>
              {configData?.map(cd => (
                <StyledRadio key={cd.name} value={cd.name}>
                  <Label>{cd.name}</Label>
                </StyledRadio>
              ))}
            </Radio.Group>
          ) : (
            <FieldValue>
              <Dropdown
                overlay={datasetOptions}
                placement="bottom"
                trigger={["click"]}
                getPopupContainer={trigger => trigger.parentElement ?? document.body}>
                <StyledDropdownButton>
                  <p style={{ margin: 0 }}>{selectedDataset.name}</p>
                  <StyledIcon icon="arrowDownSimple" size={12} />
                </StyledDropdownButton>
              </Dropdown>
            </FieldValue>
          )
        ) : (
          <Text>対応されているデータがない</Text>
        )}
      </Field>
    </Wrapper>
  );
};

export default SwitchDataset;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const StyledDropdownButton = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  align-content: center;
  width: 100%;
  height: 32px;
  padding: 0 16px;
  cursor: pointer;
`;

const StyledIcon = styled(Icon)`
  font-size: 0;
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
`;

const Text = styled.p`
  margin: 0;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  position: relative;
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const StyledRadio = styled(Radio)`
  width: 100%;
  margin-top: 8px;
`;

const Label = styled.span`
  font-size: 14px;
`;
