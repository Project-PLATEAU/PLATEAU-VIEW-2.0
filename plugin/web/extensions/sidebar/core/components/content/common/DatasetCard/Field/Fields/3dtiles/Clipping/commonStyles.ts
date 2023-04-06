import { styled } from "@web/theme";

export const FieldWrapper = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

export const ColumnFieldWrapper = styled.div<{ gap?: number }>`
  display: flex;
  flex-direction: column;
  ${({ gap }) => gap && `gap: ${gap}px;`}
`;

export const FieldValue = styled.div<{ noBorder?: boolean }>`
  display: flex;
  justify-content: start;
  align-items: center;
  ${({ noBorder }) => !noBorder && "border: 1px solid #d9d9d9;"}
  ${({ noBorder }) => !noBorder && "border-radius: 2px;"}
  flex: 1;
  height: 100%;
  width: 100%;
`;

export const Text = styled.p`
  margin: 0;
`;

export const FieldTitle = styled(Text)<{ width?: number }>`
  ${({ width }) => width && `width: ${width}px;`}
`;
