import { Select } from "@web/sharedComponents";
import { ChangeEvent, useEffect, useState } from "react";

import { FieldTitle, FieldValue, FieldWrapper, TextInput } from "../commonComponents";
import { Cond } from "../Fields/types";

const operatorOptions = [
  { value: ">", label: ">" },
  { value: "<", label: "<" },
  { value: ">=", label: ">=" },
  { value: "<=", label: "<=" },
  { value: "===", label: "=" },
  { value: "!==", label: "!=" },
];

type Props = {
  title: string;
  fieldGap?: number;
  condition: Cond<any>;
  onChange?: (condition: Cond<any>) => void;
};

const ConditionField: React.FC<Props> = ({ title, fieldGap, condition, onChange }) => {
  const [cond, setCond] = useState<Cond<any>>(condition);

  useEffect(() => {
    setCond(condition);
  }, [condition]);

  const handleOperatorChange = (operator: any) => {
    setCond(prevCond => {
      const copy = { ...prevCond, operator };
      onChange?.(copy);
      return copy;
    });
  };

  const handleOperandChange = (e: ChangeEvent<HTMLInputElement>) => {
    const operand = e.target.value;
    setCond(prevCond => {
      const copy = { ...prevCond, operand };
      onChange?.(copy);
      return copy;
    });
  };

  const handleValueChange = (e: ChangeEvent<HTMLInputElement>) => {
    const value = e.target.value;
    setCond(prevCond => {
      const copy = { ...prevCond, value };
      onChange?.(copy);
      return copy;
    });
  };

  return (
    <FieldWrapper gap={fieldGap}>
      <FieldTitle>{title}</FieldTitle>
      <FieldValue>
        <TextInput value={cond.operand} onChange={handleOperandChange} />
      </FieldValue>
      <FieldValue noBorder>
        <Select
          defaultValue={"="}
          options={operatorOptions}
          style={{ width: "100%" }}
          value={cond.operator}
          onChange={handleOperatorChange}
          getPopupContainer={trigger => trigger.parentElement ?? document.body}
        />
      </FieldValue>
      <FieldValue>
        <TextInput value={cond.value} onChange={handleValueChange} />
      </FieldValue>
    </FieldWrapper>
  );
};

export default ConditionField;
