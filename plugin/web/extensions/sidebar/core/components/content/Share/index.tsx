import CommonPage from "@web/extensions/sidebar/core/components/content/CommonPage";
import { Row, Icon, message, Spin } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { memo, useState } from "react";

import useHooks, { Project } from "./hooks";

export type Props = {
  project: Project;
  reearthURL?: string;
  backendURL?: string;
  backendProjectName?: string;
  isMobile?: boolean;
};

const Share: React.FC<Props> = ({
  project,
  reearthURL,
  backendURL,
  backendProjectName,
  isMobile,
}) => {
  const [copiedUrl, setCopiedUrl] = useState(false);
  const [copiedIframe, setCopiedIframe] = useState(false);

  const [messageApi, contextHolder] = message.useMessage();

  const {
    shareDisabled,
    publishedUrl,
    handleProjectShare,
    handleScreenshotShow,
    handleScreenshotSave,
  } = useHooks({
    project,
    reearthURL,
    backendURL,
    backendProjectName,
    messageApi,
  });

  const handleCopyToClipboard = (type: "url" | "iframe", value?: string) => {
    if (!value) return;
    navigator.clipboard.writeText(value);
    if (type === "url") {
      setCopiedUrl(true);
    } else if (type === "iframe") {
      setCopiedIframe(true);
    }
    setTimeout(() => {
      setCopiedUrl(false);
      setCopiedIframe(false);
    }, 2000);
  };

  const iframeCode = `<iframe src="${publishedUrl}" />`;

  return (
    <CommonPage title="共有・印刷" isMobile={isMobile}>
      <>
        {contextHolder}
        <ShareButton onClick={handleProjectShare} disabled={shareDisabled}>
          共有
        </ShareButton>
        {shareDisabled && (
          <Loading>
            <Spin />
          </Loading>
        )}
        {publishedUrl && (
          <>
            <Subtitle>URLで共有</Subtitle>
            <FlexWrapper>
              <ShareTextWrapper>
                <ShareText>{publishedUrl}</ShareText>
              </ShareTextWrapper>
              <StyledButton onClick={() => handleCopyToClipboard("url", publishedUrl)}>
                <Icon icon={copiedUrl ? "check" : "copy"} />
              </StyledButton>
            </FlexWrapper>
            <SubText>このURLを使えば誰でもこのマップにアクセスできます。</SubText>
            <Subtitle>HTMLページへの埋め込みは下記のコードをお使いください：</Subtitle>
            <FlexWrapper>
              <ShareTextWrapper>
                <ShareText>{iframeCode}</ShareText>
              </ShareTextWrapper>
              <StyledButton onClick={() => handleCopyToClipboard("iframe", iframeCode)}>
                <Icon icon={copiedIframe ? "check" : "copy"} />
              </StyledButton>
            </FlexWrapper>
            <SubText>このURLを使えば誰でもこのマップにアクセスできます。</SubText>
          </>
        )}
      </>
      <>
        <Subtitle>印刷</Subtitle>
        <SectionWrapper>
          <ButtonWrapper>
            <Button onClick={handleScreenshotSave}>ダウンロード</Button>
            <Button onClick={handleScreenshotShow}>プリントビュー</Button>
          </ButtonWrapper>
          <SubText>このマップを印刷できる状態で表示</SubText>
        </SectionWrapper>
      </>
    </CommonPage>
  );
};

export default memo(Share);

const Text = styled.p`
  font-size: 14px;
  margin: 0;
`;

const Subtitle = styled(Text)`
  margin-bottom: 15px;
`;

const SubText = styled.p`
  font-size: 12px;
  color: #b1b1b1;
  margin: 8px 0 16px;
`;

const SectionWrapper = styled(Row)`
  display: flex;
  flex-direction: column;
  align-items: flex-start;
`;

const Button = styled.button`
  height: 37px;
  width: 160px;
  border: none;
  border-radius: 3px;
  background: #ffffff;
  font-size: 14px;
  line-height: 21px;
  cursor: pointer;
`;

const ShareButton = styled.button<{ disabled?: boolean }>`
  height: 32px;
  width: 100%;
  background: ${({ disabled }) => (disabled ? "#D1D1D1" : "#00bebe")};
  color: white;
  border: none;
  border-radius: 2px;
  margin-bottom: 15px;
  cursor: ${({ disabled }) => (disabled ? "not-allowed" : "pointer")};
`;

const FlexWrapper = styled.div`
  display: flex;
  align-items: center;
  width: 100%;
  height: 32px;
`;

const ButtonWrapper = styled(FlexWrapper)`
  gap: 8px;
`;

const StyledButton = styled.button`
  background: #00bebe;
  border: none;
  border-radius: 2px;
  width: 40px;
  height: 100%;
  cursor: pointer;
  color: white;

  :hover {
    background: #00bebe;
    border-color: #00bebe;
  }
`;

const ShareTextWrapper = styled.div`
  display: flex;
  align-items: center;
  height: 100%;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 0 12px;
  white-space: nowrap;
  overflow: auto;
  -ms-overflow-style: none;
  ::-webkit-scrollbar {
    display: none;
  }
`;

const ShareText = styled.p`
  margin: 0;
  color: rgba(0, 0, 0, 0.45);
`;
const Loading = styled.div`
  position: absolute;
  width: 100%;
  height: 100%;
  min-height: 200px;
  left: 0;
  top: 0;
  display: flex;
  align-items: center;
  justify-content: center;
`;
