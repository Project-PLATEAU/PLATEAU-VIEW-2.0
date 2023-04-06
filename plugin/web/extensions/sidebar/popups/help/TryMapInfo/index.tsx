import mapVideo from "@web/extensions/sidebar/core/assets/mapVideo.png";
import { postMsg } from "@web/extensions/sidebar/utils";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback } from "react";

import { NumberingWrapper, ParagraphItem } from "../sharedComponent";

const TryMapInfo: React.FC = () => {
  const handleShowMapModal = useCallback(() => {
    postMsg({ action: "mapModalOpen" });
  }, []);

  return (
    <Wrapper>
      <Title>マップを使ってみる</Title>

      <ContentWrapper>
        <ImgWrapper>
          <img src={mapVideo} onClick={handleShowMapModal} />
        </ImgWrapper>
        <Paragraph>
          上の動画では、空間データや地図を扱うのが初めての方が、データを追加したり地図上に表現するために必要な、基本的な機能について紹介しています。
          動画を見る時間がない方は、次の手順をお試しください。
        </Paragraph>
        <ParagraphItem>
          <NumberingWrapper number={1} />
          <Paragraph>
            「
            <InlineIcon icon="plusCircle" size={16} color={" #00bebe"} />
            <BlueText>カタログから検索する</BlueText>
            」ボタンで使用可能なデータを表示し、マップに追加してみましょう。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={2} />
          <Paragraph>
            画面の左側に現れる
            <InlineIcon icon={"visible"} size={16} />
            ボタンで表示/非表示ボタンを切り替えてみましょう。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={3} />
          <Paragraph>
            「<InlineIcon icon="sliders" size={16} />
            <BlueText>マップ設定</BlueText>」ボタンをクリックして、背景図を変更してみましょう。
          </Paragraph>
        </ParagraphItem>
        <ParagraphItem>
          <NumberingWrapper number={4} />
          <Paragraph>画面右上のズームや回転ボタンを使って、視点を変化させてみましょう。</Paragraph>
        </ParagraphItem>
      </ContentWrapper>
    </Wrapper>
  );
};

export default TryMapInfo;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px 16px;
  gap: 24px;
  width: 333px;
  height: 670px;
`;

const ImgWrapper = styled.div`
  width: 301px;
  height: 170px;
  cursor: pointer;
`;

const Title = styled.p`
  margin: 0;
  font-size: 16px;
  line-height: 24px;
  color: rgba(0, 0, 0, 0.85);
`;

const Paragraph = styled.p`
  color: rgba(0, 0, 0, 0.45);
`;

const ContentWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
  padding: 0px;
  gap: 20px;
  width: 301px;
  height: 664px;
`;

const BlueText = styled.span`
  color: #00bebe;
`;
const InlineIcon = styled(Icon)`
  display: inline-block;
  color: #00bebe;
`;
