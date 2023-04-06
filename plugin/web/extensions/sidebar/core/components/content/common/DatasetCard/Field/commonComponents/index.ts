import { Input, InputNumber } from "@web/sharedComponents";
import { styled } from "@web/theme";

export const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

export const Item = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 8px;
`;

export const TextInput = styled(Input)`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;

  :focus {
    border: none;
  }
`;

export const NumberInput = styled(InputNumber)`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;

  :focus {
    border: none;
  }
`;

export const ButtonWrapper = styled.div`
  width: 125px;
  align-self: flex-end;
`;

export const FieldWrapper = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

export const FieldValue = styled.div<{ noBorder?: boolean }>`
  position: relative;
  display: flex;
  justify-content: start;
  align-items: center;
  ${({ noBorder }) => !noBorder && "border: 1px solid #d9d9d9;"}
  ${({ noBorder }) => !noBorder && "border-radius: 2px;"}
  flex: 1;
  height: 100%;
  width: 100%;
`;

const Text = styled.p`
  margin: 0;
`;

export const FieldTitle = styled(Text)<{ width?: number }>`
  ${({ width }) => width && `width: ${width}px;`}
`;
