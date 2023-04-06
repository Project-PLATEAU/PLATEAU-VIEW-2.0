import { styled } from "@web/theme";

export type Props = {
  number?: number;
};
const NumberingWrapper: React.FC<Props> = ({ number }) => {
  return (
    <Wrapper>
      <Text>{number}</Text>
    </Wrapper>
  );
};

export default NumberingWrapper;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 4px;
  gap: 10px;
  width: 16px;
  height: 22px;
  background: #595959;
  border-radius: 8px;
`;

const Text = styled.p`
  width: 8px;
  height: 22px;
  color: #ffffff;
`;
