import { Button, Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { ReactNode } from "react";

export type Props = {
  title?: string;
  children?: ReactNode;
  onModalClose: () => void;
};

const CommonModalWrapper: React.FC<Props> = ({ title, children, onModalClose }) => {
  return (
    <Wrapper>
      {title && (
        <HeaderWrapper>
          <Title>{title}</Title>
          <CloseButton size={16} icon="close" color="#00000073" onClick={onModalClose} />
        </HeaderWrapper>
      )}

      <ContentWrapper>{children}</ContentWrapper>

      <FooterWrapper>
        <OkButton type="primary" onClick={onModalClose}>
          <Text>OK</Text>
        </OkButton>
      </FooterWrapper>
    </Wrapper>
  );
};

export default CommonModalWrapper;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  background: #ffffff;
  box-shadow: 0px 3px 6px -4px rgba(0, 0, 0, 0.12), 0px 6px 16px rgba(0, 0, 0, 0.08),
    0px 9px 28px 8px rgba(0, 0, 0, 0.05);
  border-radius: 2px;
  height: 100%;
`;

const Title = styled.p`
  font-size: 16px;
  margin-bottom: 0;
`;

const CloseButton = styled(Icon)`
  cursor: pointer;
`;

const HeaderWrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: space-between;
  align-items: center;
  padding: 16px 24px;
  gap: 36px;
  height: 56px;
  background: #ffffff;
  box-shadow: inset 0px -1px 0px #f0f0f0;
  flex: none;
  order: 0;
  align-self: stretch;
  flex-grow: 0;
`;

const ContentWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 24px 24px;
  height: 100%;
`;

const FooterWrapper = styled.div`
  display: flex;
  flex-direction: row;
  justify-content: flex-end;
  align-items: center;
  padding: 10px 16px;
  height: 52px;
  background: #ffffff;
  box-shadow: inset 0px 1px 0px #f0f0f0;
  flex: none;
  order: 2;
  align-self: stretch;
  flex-grow: 0;
`;

const OkButton = styled(Button)`
  padding: 5px 16px;
  width: 51px;
  height: 32px;
  background: #1890ff;
  box-shadow: 0px 2px 0px rgba(0, 0, 0, 0.043);
  border-radius: 2px;
  text-align: center;
`;

const Text = styled.p`
  font-size: 14px;
`;
