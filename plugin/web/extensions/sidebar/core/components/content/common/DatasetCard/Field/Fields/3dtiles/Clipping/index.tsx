import { Checkbox } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { CheckboxChangeEvent } from "antd/lib/checkbox";

import { BaseFieldProps } from "../../types";

import { Text } from "./commonStyles";
import useHooks from "./hooks";
import SelectField from "./SelectField";
import SwitchField from "./SwitchField";

const Clipping: React.FC<BaseFieldProps<"clipping">> = ({ value, dataID, editMode, onUpdate }) => {
  const { options, handleUpdateBool, handleUpdateSelect } = useHooks({
    value,
    dataID,
    onUpdate,
  });

  return editMode ? null : (
    <div>
      <FieldWrapper>
        <SwitchField
          style={{ margin: 0 }}
          title="有効にする"
          titleWidth={87}
          checked={options.enabled}
          onChange={handleUpdateBool("enabled")}
        />
      </FieldWrapper>
      {options.enabled && (
        <>
          <FieldWrapper>
            <Checkbox
              style={{ margin: 0 }}
              checked={options.show}
              onChange={(e: CheckboxChangeEvent) => handleUpdateBool("show")(e.target.checked)}>
              <Text>クリップボックスを表示する</Text>
            </Checkbox>
          </FieldWrapper>
          <FieldWrapper>
            <Checkbox
              style={{ margin: 0 }}
              checked={options.aboveGroundOnly}
              onChange={(e: CheckboxChangeEvent) =>
                handleUpdateBool("aboveGroundOnly")(e.target.checked)
              }>
              <Text>クリップボックスを地面にスナップする</Text>
            </Checkbox>
          </FieldWrapper>
          <SelectField
            title="クリップ方法の選択"
            defaultValue="inside"
            style={{ width: "100%" }}
            value={options.direction}
            onChange={handleUpdateSelect("direction")}
            getPopupContainer={trigger => trigger.parentElement ?? document.body}
            options={[
              {
                value: "inside",
                label: "ボックス内をクリップ",
              },
              {
                value: "outside",
                label: "ボックス外をクリップ",
              },
            ]}
          />
        </>
      )}
    </div>
  );
};

export default Clipping;

const FieldWrapper = styled.div`
  margin-bottom: 8px;
`;
