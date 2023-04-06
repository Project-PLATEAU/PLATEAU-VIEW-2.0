import shadow from "@web/extensions/sidebar/core/assets/shadowImg.png";
import { styled } from "@web/theme";

const ShadowFunction: React.FC = () => {
  return (
    <Wrapper>
      <Title>日影機能について</Title>

      <ContentWrapper>
        <Paragraph>
          PLATEAU
          VIEWでは建物モデルなどが見やすいように、デフォルトでは日影の効果を「なし」に設定しています。
          日影の効果をONにするための手順は以下の通りです。
        </Paragraph>

        <Paragraph>
          建物モデルをVIEWに追加した時に、左側の一覧に表示される「影」というメニューのプルダウンから「投影のみ」、「受光のみ」、「投影と受光」のいずれかを選択します。
        </Paragraph>

        <ImgWrapper>
          <img src={shadow} />
        </ImgWrapper>

        <Paragraph>
          PLATEAU
          VIEWではSimonら（1994）の手法で推定された太陽の位置を用いて日影を計算し、表現しています。タイムバーが出ていないときは、地図が見やすいように仮想の光源から光を当てて表示しています。
          Simon, J., Bretagnon, P., Chapront, J., & Chapront-Touze, M. (1994). Numerical expressions
          for precession formulae and mean elements for the Moon and the planets. Astronomy and
          Astrophysics (Berlin. Print), 282(2), 663–683.
        </Paragraph>
      </ContentWrapper>
    </Wrapper>
  );
};

export default ShadowFunction;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 16px;
  gap: 24px;
  width: 333px;
  height: 781px;
`;

const Title = styled.p`
  margin: 0;
  font-size: 16px;
  line-height: 24px;
  color: rgba(0, 0, 0, 0.85);
`;

const ContentWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  gap: 20px;
  width: 301px;
  height: 906px;
`;

const Paragraph = styled.p`
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.45);
`;

const ImgWrapper = styled.div`
  width: 301px;
  height: 188px;
`;
