import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { ReactNode } from "react";

type Props = {
  children?: ReactNode;
  handleClose?: () => void;
};

const PopupWrapper: React.FC<Props> = ({ children, handleClose }) => {
  return (
    <Wrapper>
      <Header>
        <CloseButton>
          <Icon size={32} icon="close" onClick={handleClose} />
        </CloseButton>
      </Header>
      {children}
    </Wrapper>
  );
};

export default PopupWrapper;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  background: #e7e7e7;
`;

const Header = styled.div`
  display: flex;
  background: #e7e7e7;
  position: relative;
  width: 100%;
  height: 32px;
`;

const CloseButton = styled.button`
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  right: 0;
  height: 32px;
  width: 32px;
  border: none;
  background: #00bebe;
  color: white;
  cursor: pointer;
`;
