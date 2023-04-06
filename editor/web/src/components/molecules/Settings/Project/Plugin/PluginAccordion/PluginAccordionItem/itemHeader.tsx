import React, { useState } from "react";

import Box from "@reearth/components/atoms/Box";
import Button from "@reearth/components/atoms/Button";
import Flex from "@reearth/components/atoms/Flex";
import Text from "@reearth/components/atoms/Text";
import { useT } from "@reearth/i18n";
import { fonts, styled } from "@reearth/theme";

import DeleteModal from "./deleteModal";

export type PluginItemProps = {
  className?: string;
  thumbnail?: string;
  title?: string;
  version?: string;
  isInstalled?: boolean;
  onUninstall: () => void;
};

const PluginAccordionItemHeader: React.FC<PluginItemProps> = ({
  className,
  thumbnail,
  title,
  version,
  isInstalled,
  onUninstall,
}) => {
  const t = useT();
  const [isModalOpen, setIsModalOpen] = useState(false);
  const [hovered, setHovered] = useState(false);
  const handleMouseEnter = () => {
    setHovered(true);
  };
  const handleMouseLeave = () => {
    setHovered(false);
  };
  return (
    <Wrapper align="center" justify="space-between" className={className}>
      <Flex align="center">
        <TitleWrapper>
          {thumbnail && (
            <Box borderRadius={8} mh="m">
              <Thumbnail src={thumbnail} alt="plugin thumbnail" />
            </Box>
          )}
          <Text size="m" weight="bold" otherProperties={{ marginRight: "20px", maxWidth: "200px" }}>
            {title}
          </Text>
        </TitleWrapper>
        <Text size="m">v{version}</Text>
      </Flex>
      <StyledButton
        buttonType={isInstalled && hovered ? "danger" : "secondary"}
        type="button"
        large
        icon={isInstalled ? (hovered ? "bin" : "check") : "install"}
        text={isInstalled ? (hovered ? t("Uninstall") : t("Installed")) : t("Install")}
        onMouseEnter={handleMouseEnter}
        onMouseLeave={handleMouseLeave}
        onClick={isInstalled ? () => setIsModalOpen(true) : undefined}
      />
      <DeleteModal
        onCancel={() => setIsModalOpen(false)}
        isOpen={isModalOpen}
        onClose={() => setIsModalOpen(false)}
        onProceed={onUninstall}
      />
    </Wrapper>
  );
};

const Wrapper = styled(Flex)`
  width: 100%;
  padding: ${props => `${props.theme.metrics.xl}px 0`};
`;

const TitleWrapper = styled(Flex)`
  width: 250px;
  margin-right: 32px;
`;

const Thumbnail = styled.img`
  border-radius: 8px;
  width: 64px;
  height: 64px;
`;

const StyledButton = styled(Button)`
  font-weight: ${fonts.weight.bold};
  width: 153px;
  border-radius: ${props => props.theme.metrics.s}px;
  padding: ${({ theme }) => `${theme.metrics.s}px ${theme.metrics["2xl"]}`};
  transition: all 0.3s;
`;

export default PluginAccordionItemHeader;
