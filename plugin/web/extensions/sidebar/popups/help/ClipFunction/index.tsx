import clip from "@web/extensions/sidebar/core/assets/clip.png";
import { postMsg } from "@web/extensions/sidebar/utils";
import { styled } from "@web/theme";
import { useCallback } from "react";

import { NumberingWrapper, ParagraphItem } from "../sharedComponent";

const ClipFunction: React.FC = () => {
  const handleShowClipModal = useCallback(() => {
    postMsg({ action: "clipModalOpen" });
  }, []);

  return (
    <Wrapper>
      <Title>クリップ機能について</Title>
      <ContentWrapper>
        <ImgWrapper>
          <img src={clip} onClick={handleShowClipModal} />
        </ImgWrapper>
        <Paragraph>
          この動画では、3Dモデルの断面を表示するための、クリップボックスの使い方を紹介しています。
          （クリップ機能の使い方）
        </Paragraph>
        <ParagraphItem>
          <NumberingWrapper number={1} />
          <Paragraph>
            クリップ機能が使えるデータ（建物モデルやBIMデータ）を表示すると、左側の一覧の中に「クリップ機能」というチェックボックスが表示されます。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={2} />
          <Paragraph>チェックボックスにチェックを入れると、機能を有効にできます。</Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={3} />
          <Paragraph>機能を有効にすると、画面中央に立方体が表示されます。</Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={4} />
          <Paragraph>
            この立方体を移動して、3Dモデルに重ねると、立方体の内側、あるいは外側のみ3Dモデルを表示することができます。内側、外側の切り替えにはプルダウンの「ボックス内をクリップ」、「ボックス外をクリップ」を選択します。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={5} />
          <Paragraph>
            立方体のサイズや形は頂点の赤いポイント、あるいは各面の中央に配置された青いポイントを使って変更することができます。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={6} />
          <Paragraph>
            「クリップボックスを表示」のチェックボックスはデフォルトではチェックが入っていますが、このチェックを外すと、最初に表示された立方体を非表示にすることができます。立方体は表示されませんが、クリップ機能は有効です。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={7} />
          <Paragraph>
            「クリップボックスを地面にスナップ」のチェックボックスはデフォルトではチェックが入っていて、地下に潜らないようになっています。地下のオブジェクトに対してもクリップ機能を使いたい場合は、このチェックを外してください。
          </Paragraph>
        </ParagraphItem>
      </ContentWrapper>
    </Wrapper>
  );
};

export default ClipFunction;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 16px;
  gap: 24px;
  width: 333px;
  height: 1100px;
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
  height: 1052px;
`;

const Paragraph = styled.p`
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.45);
`;

const ImgWrapper = styled.div`
  width: 301px;
  height: 144px;
  cursor: pointer;
`;
