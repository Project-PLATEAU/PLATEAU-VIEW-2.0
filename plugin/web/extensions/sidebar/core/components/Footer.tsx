import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { memo } from "react";

export type Props = {
  datasetQuantity?: number;
  onRemoveAll?: () => void;
};

const Footer: React.FC<Props> = ({ datasetQuantity, onRemoveAll }) => {
  return (
    <FooterBan>
      <RemoveBtn onClick={onRemoveAll}>
        <Icon icon="trash" />
        全てを削除
      </RemoveBtn>
      <Text>データセット x {datasetQuantity ?? 0}</Text>
    </FooterBan>
  );
};

export default memo(Footer);

const FooterBan = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 8px 16px;
  height: 48px;
  border-top: 1px solid #c7c5c5;
  background: #f4f4f4;
  color: #4a4a4a;
`;

const Text = styled.p`
  margin: 0;
  user-select: none;
`;

const RemoveBtn = styled.button`
  display: flex;
  flex-direction: row;
  align-items: flex-start;
  gap: 10px;
  width: 131px;
  height: 32px;
  border: 1px solid #c7c5c5;
  border-radius: 4px;
  background: inherit;
  padding: 4px 10px;
  cursor: pointer;
`;
