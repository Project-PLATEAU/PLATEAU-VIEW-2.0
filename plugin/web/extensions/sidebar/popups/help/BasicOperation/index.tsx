import BasicOperButton from "@web/extensions/sidebar/core/assets/BasicOperButton.png";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

const BasicOperation: React.FC = () => {
  return (
    <Wrapper>
      <TopWrapper>
        <Title>視点や画面移動</Title>
        <InstructionWrapper>
          <ImgWrapper>
            <Icon icon="mouseMiddleButton" width={65} height={95} />
            <CaptionText>
              ドラッグ：視点移動
              <br />
              <br />
              回転：拡大縮小
            </CaptionText>
          </ImgWrapper>
          <ImgWrapper>
            <Icon icon="mouseLeftButton" width={65} height={95} />
            <CaptionText>ドラッグ：画面移動</CaptionText>
          </ImgWrapper>
          <ImgWrapper>
            <Icon icon="mouseRightButton" width={65} height={95} />
            <CaptionText>ドラッグ：拡大縮小</CaptionText>
          </ImgWrapper>
        </InstructionWrapper>
      </TopWrapper>
      <BottomWrapper>
        <Title>データの追加</Title>
        <Paragraph>以下のボタンで建物やデータを地図に追加してください</Paragraph>
        <img src={BasicOperButton} />
        <Paragraph>
          カタログから検索するウインドウが表示されたら、
          <br /> ① 表示したいエリアに対応するフォルダをクリックして開く
          <br />② 建物モデルや重ね合わせたいデータを
          <InlineIcon icon="plusCircle" size={16} />
          で選択する
        </Paragraph>
      </BottomWrapper>
    </Wrapper>
  );
};

export default BasicOperation;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 16px;
  gap: 24px;
  width: 333px;
  height: 800px;
`;
const Title = styled.p`
  margin: 0;
  font-size: 16px;
  line-height: 24px;
  color: rgba(0, 0, 0, 0.85);
`;

const BottomWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  gap: 24px;
  width: 301px;
  height: 354px;
`;

const Paragraph = styled.p`
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.45);
`;

const InlineIcon = styled(Icon)`
  display: inline-block;
  color: #00bebe;
`;
const TopWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  gap: 24px;
  width: 301px;
  height: 405px;
`;

const InstructionWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  padding: 0px;
  gap: 36px;
  width: 301px;
  height: 357px;
`;

const ImgWrapper = styled.div`
  display: flex;
  flex-direction: row;
  align-items: center;
  padding: 0px;
  gap: 36px;
  width: 230px;
`;

const CaptionText = styled.p`
  font-weight: 700;
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.45);
`;
