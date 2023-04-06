import styled from "@emotion/styled";

import Button from "@reearth-cms/components/atoms/Button";
import { useT } from "@reearth-cms/i18n";

const NotFound: React.FC = () => {
  const t = useT();

  return (
    <Wrapper>
      <CircleWrapper>
        <Circle>404</Circle>
      </CircleWrapper>
      <Content>
        <StyledTitle>{t("Oops!")}</StyledTitle>
        <StyledText>{t("PAGE NOT FOUND ON SERVER")}</StyledText>
        <Button href="/" type="primary">
          {t("Go back Home")}
        </Button>
      </Content>
    </Wrapper>
  );
};

export default NotFound;

const Wrapper = styled.div`
  display: flex;
  flex-wrap: wrap;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #f0f2f5;
`;

const CircleWrapper = styled.div`
  padding: 32px;
`;

const Circle = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  font-size: 96px;
  color: #bfbfbf;
  font-weight: 700;
  background-color: #d9d9d9;
  width: 240px;
  height: 240px;
  padding: 0;
  border-radius: 50%;
`;

const Content = styled.div`
  padding: 32px;
  text-align: center;
`;

const StyledTitle = styled.h1`
  text-align: center;
  color: #1890ff;
  font-weight: 500;
  font-size: 38px;
  line-height: 46px;
  margin-bottom: 24px;
`;

const StyledText = styled.p`
  font-weight: 500;
  font-size: 16px;
  line-height: 24px;
  color: #00000073;
  margin-bottom: 24px;
`;
