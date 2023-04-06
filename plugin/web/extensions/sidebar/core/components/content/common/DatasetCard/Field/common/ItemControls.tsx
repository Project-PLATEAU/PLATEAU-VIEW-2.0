import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

type Props = {
  handleMoveUp: (index: number) => void;
  handleMoveDown: (index: number) => void;
  handleRemove: (index: number) => void;
  index: number;
};

const ItemControls: React.FC<Props> = ({ handleMoveUp, handleMoveDown, handleRemove, index }) => {
  return (
    <Wrapper>
      <Icon icon="arrowUpThin" size={16} onClick={() => handleMoveUp(index)} />
      <Icon icon="arrowDownThin" size={16} onClick={() => handleMoveDown(index)} />
      <Icon icon="trash" size={16} onClick={() => handleRemove(index)} />
    </Wrapper>
  );
};

export default ItemControls;

const Wrapper = styled.div`
  display: flex;
  justify-content: right;
  gap: 4px;
  cursor: pointer;
`;
