import welcomeScreenVideo from "@web/extensions/sidebar/core/assets/welcomeScreenVideo.png";
import useHooks from "@web/extensions/sidebar/modals/welcomescreen/hooks";
import { Checkbox, Icon } from "@web/sharedComponents";
import Video from "@web/sharedComponents/Video";
import { styled } from "@web/theme";

const WelcomeScreen: React.FC = () => {
  const {
    isMobile,
    showVideo,
    dontShowAgain,
    handleDontShowAgain,
    handleCloseVideo,
    handleClose,
    handleOpenHelp,
    handleOpenCatalog,
  } = useHooks();

  return (
    <Wrapper>
      {!showVideo ? (
        <InnerWrapper isMobile={isMobile}>
          <WelcomeCloseButton size={40} icon="close" onClick={handleClose} isMobile={isMobile} />
          <TextWrapper isMobile={isMobile}>
            <Title weight={700} size={isMobile ? 24 : 48}>
              ようこそ
            </Title>
            <Text weight={500} size={isMobile ? 16 : 20}>
              {isMobile ? "データがお好きですか？" : "マップを使ってみる"}
            </Text>
          </TextWrapper>
          <ContentWrapper isMobile={isMobile}>
            {!isMobile && (
              <ImgWrapper
                type="text"
                href="https://www.mlit.go.jp/plateau/learning/?topic=plateau-view"
                target="_blank"
                imgUrl={welcomeScreenVideo}>
                <Icon icon="playCircle" size={48} color="#fff" />
              </ImgWrapper>
            )}
            <BtnsWrapper isMobile={isMobile}>
              {!isMobile && (
                <ButtonWrapper onClick={handleOpenHelp}>
                  <Text weight={500} size={14}>
                    ヘルプをみる
                  </Text>
                </ButtonWrapper>
              )}
              <ButtonWrapper onClick={handleOpenCatalog}>
                <Icon size={20} icon="plusCircle" color="#fafafa" />
                <Text weight={500} size={14}>
                  カタログから検索する
                </Text>
              </ButtonWrapper>
            </BtnsWrapper>
          </ContentWrapper>
          <CheckWrapper>
            <Checkbox checked={dontShowAgain} onClick={handleDontShowAgain} />
            <Text weight={700} size={14}>
              閉じて今後は表示しない
            </Text>
          </CheckWrapper>
        </InnerWrapper>
      ) : (
        <CloseBtnWrapper isMobile={isMobile}>
          <VideoCloseButton size={40} icon="close" onClick={handleCloseVideo} isMobile={isMobile} />
          <VideoWrapper>
            <Video width=" 1142" height="543" src="https://www.youtube.com/embed/pY2dM-eG5mA" />
          </VideoWrapper>
        </CloseBtnWrapper>
      )}
    </Wrapper>
  );
};

export default WelcomeScreen;

const Wrapper = styled.div`
  width: 100%;
  height: 100%;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  background: rgba(0, 0, 0, 0.7);
`;

const InnerWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  flex-direction: column;
  position: relative;
  width: 100%;
  max-width: ${({ isMobile }) => (isMobile ? "70vw" : "742px")};
`;

const Text = styled.p<{ weight: number; size: number }>`
  font-weight: ${({ weight }) => weight}px;
  font-size: ${({ size }) => size}px;
  margin: 0px;
  color: #fafafa;
`;

const Title = styled(Text)<{ isMobile?: boolean }>`
  ${({ isMobile }) => !isMobile && `margin-bottom: 24px;`}
`;

const TextWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  flex-direction: column;
  align-items: ${({ isMobile }) => (isMobile ? "center" : "flex-start")};
  justify-content: flex-end;
  margin-bottom: ${({ isMobile }) => (isMobile ? "60px" : "24px")};
`;

const ContentWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  ${({ isMobile }) =>
    isMobile &&
    `
  flex-direction: column;
  align-items: center;
  `};
  justify-content: space-between;
`;

const BtnsWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  gap: 24px;
  width: ${({ isMobile }) => (isMobile ? "100%" : "318px")};
`;

const ImgWrapper = styled.a<{ imgUrl: string }>`
  display: flex;
  justify-content: center;
  align-items: center;
  width: 305px;
  height: 159px;
  cursor: pointer;
  background-color: rgba(0, 0, 0, 0.5);
  background-image: ${props =>
    `linear-gradient(rgba(0, 0, 0, 0.5), rgba(0, 0, 0, 0.5)), url(${props.imgUrl})`};
  background-size: cover;
  background-repeat: no-repeat;
  background-position: center center;
`;

const CloseButton = styled(Icon)<{ isMobile?: boolean }>`
  width: ${({ isMobile }) => (isMobile ? "48px" : "40px")};
  height: ${({ isMobile }) => (isMobile ? "48px" : "40px")};
  border: none;
  color: white;
  cursor: pointer;
`;

const WelcomeCloseButton = styled(CloseButton)`
  position: absolute;
  right: ${({ isMobile }) => (isMobile ? "-48px" : "-40px")};
  top: ${({ isMobile }) => (isMobile ? "-48px" : "-40px")};
  flex-grow: 1;
`;

const VideoCloseButton = styled(CloseButton)`
  align-self: end;
`;

const ButtonWrapper = styled.div<{ selected?: boolean }>`
  display: flex;
  flex-direction: row;
  justify-content: center;
  align-items: center;
  padding: 10px;
  width: 100%;
  height: 41px;
  background: ${({ selected }) => (selected ? "#d1d1d1" : "#00bebe")};
  border-radius: 4px;
  border: none;
  gap: 8px;
  cursor: pointer;
  transition: background 0.3s;
  :hover {
    background: #d1d1d1;
  }
`;

const CheckWrapper = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  align-self: center;
  gap: 8px;
  margin-top: 50px;
`;

const VideoWrapper = styled.div`
  width: 1142px;
  height: 543px;
`;
const CloseBtnWrapper = styled.div<{ isMobile?: boolean }>`
  display: flex;
  flex-direction: column;
  justify-content: center;
  position: relative;
  width: ${({ isMobile }) => (isMobile ? "348px" : "1182px")};
  height: ${({ isMobile }) => (isMobile ? "390px" : "635px")};
`;
