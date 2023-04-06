import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

type Props = {
  icon: string;
  text: string;
  disabled?: boolean;
  active?: boolean;
  onClick?: () => void;
};

const ControlButton: React.FC<Props> = ({ icon, text, disabled, active, onClick }) => {
  return (
    <StyledButton onClick={onClick} disabled={disabled} active={active}>
      <Icon icon={icon} size={16} />
      <Text>{text}</Text>
    </StyledButton>
  );
};

const StyledButton = styled.div<{ disabled?: boolean; active?: boolean }>`
  width: 100%;
  height: 30px;
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  border: ${({ disabled }) => (disabled ? "1px solid #595959" : "1px solid var(--theme-color)")};
  background-color: ${({ active }) => (active ? "var(--theme-color)" : "#ffffff")};
  color: ${({ disabled, active }) =>
    disabled ? "#595959" : active ? "#ffffff" : "var(--theme-color)"};
  cursor: ${({ disabled }) => (disabled ? "default" : "pointer")};
  pointer-events: ${({ disabled }) => (disabled ? "none" : "all")};
  border-radius: 4px;
  user-select: none;
  transition: all 0.25s ease;

  &:hover {
    background-color: var(--theme-color);
    color: #fff;
  }
`;

const Text = styled.div``;

export default ControlButton;
