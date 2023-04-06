import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

import useHooks from "./hooks";

const Pedestrian: React.FC = () => {
  const { onPedestrian } = useHooks();

  return (
    <MainButton onClick={onPedestrian}>
      <Icon icon="personSimpleWalk" size={20} />
    </MainButton>
  );
};

const MainButton = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  width: 32px;
  height: 32px;
  background-color: #fff;
  border-radius: 16px;
  cursor: pointer;
`;

export default Pedestrian;
