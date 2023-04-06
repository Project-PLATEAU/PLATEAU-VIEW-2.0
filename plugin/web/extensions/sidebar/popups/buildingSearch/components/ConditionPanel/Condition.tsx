import { Select } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useMemo } from "react";

import { IndexData, Condition as ConditionType } from "../../types";
import "./index.css";

type Props = {
  indexItem: IndexData;
  setConditions: React.Dispatch<React.SetStateAction<ConditionType[]>>;
};

const Condition: React.FC<Props> = ({ indexItem, setConditions }) => {
  const options = useMemo(() => indexItem.values.map(v => ({ key: v, value: v })), [indexItem]);

  const handleChange = useCallback(
    (value: string[] | unknown) => {
      setConditions((conditions: ConditionType[]) => {
        const field = conditions.find(c => c.field === indexItem.field);
        if (field) field.values = value as string[];
        return [...conditions];
      });
    },
    [indexItem, setConditions],
  );

  return (
    <Wrapper key={indexItem.field}>
      <Title>{indexItem.field}：</Title>
      <StyledSelect
        mode="multiple"
        showArrow
        placeholder="キーワードを入力"
        listHeight={200}
        onChange={handleChange}
        options={options}
        getPopupContainer={trigger => trigger.parentElement ?? document.body}
      />
    </Wrapper>
  );
};

const Wrapper = styled.div`
  position: relative;
  display: flex;
  flex-direction: column;
  gap: 4px;
  padding: 6px 20px;
`;

const Title = styled.div`
  font-size: 14px;
`;

const StyledSelect = styled(Select)`
  &.ant-select:not(.ant-select-disabled):hover .ant-select-selector {
    border-color: var(--theme-color);
  }
  &.ant-select-focused:not(.ant-select-disabled).ant-select:not(.ant-select-customize-input)
    .ant-select-selector {
    border-color: var(--theme-color);
    box-shadow: 0 0 0 2px #bee5e5;
  }
`;

export default Condition;
