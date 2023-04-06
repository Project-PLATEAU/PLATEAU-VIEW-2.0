import CommonPage from "@web/extensions/sidebar/core/components/content/CommonPage";
import { postMsg } from "@web/extensions/sidebar/utils";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useEffect } from "react";

import useHooks from "./hooks";

const Help: React.FC = () => {
  const { items, selectedTab, handleItemClicked } = useHooks();

  useEffect(() => {
    postMsg({ action: "helpPopupOpen" });
    return () => {
      postMsg({ action: "popupClose" });
    };
  }, []);

  return (
    <CommonPage title="使い方">
      <MenuWrapper>
        {items.map(item => (
          <MenuItem
            key={item.key}
            selected={item.key === selectedTab}
            onClick={() => handleItemClicked(item.key)}>
            <Text>{item?.label}</Text>
            <Icon icon="arrowRight" size={16} />
          </MenuItem>
        ))}
      </MenuWrapper>
    </CommonPage>
  );
};

export default Help;

const MenuWrapper = styled.div`
  display: flex;
  flex-direction: column;
  align-items: space-between;
  padding: 0px;
  gap: 12px;
  background: #e7e7e7;
  width: 326px;
`;

const MenuItem = styled.div<{ selected?: boolean }>`
  display: flex;
  justify-content: space-between;
  align-items: center;
  height: 42px;
  padding: 10px 12px;
  cursor: pointer;

  ${({ selected }) =>
    selected &&
    `
  background: #00BEBE;
  color: #fff;
  `}
`;

const Text = styled.p`
  margin: 0;
`;
