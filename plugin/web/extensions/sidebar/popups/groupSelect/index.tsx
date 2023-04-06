import { postMsg } from "@web/extensions/sidebar/utils";
import { Button, Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useEffect, useState } from "react";

import { fieldGroups } from "../../utils/fieldGroups";

const GroupSelect: React.FC = () => {
  const [selectedGroup, selectGroup] = useState<string>();

  const handleSave = useCallback(() => {
    postMsg({ action: "saveGroups", payload: { selected: selectedGroup } });
  }, [selectedGroup]);

  const handleCancel = useCallback(() => {
    postMsg({ action: "popupClose" });
  }, []);

  const handleGroupSelect = useCallback(
    (id: string) => {
      if (selectedGroup === id) {
        selectGroup(undefined);
      } else {
        selectGroup(id);
      }
    },
    [selectedGroup],
  );

  useEffect(() => {
    if ((window as any).groupSelectInit) {
      const init = (window as any).groupSelectInit;
      if (init.selected) {
        selectGroup(init.selected);
      }
    }
  }, []);

  return (
    <Wrapper>
      <Header>
        <Text>グルーブ管理</Text>
        <StyledIcon icon="close" size={16} onClick={handleCancel} />
      </Header>
      <Content>
        <Text>グループリスト</Text>
        <List>
          {fieldGroups.map(g => (
            <ListItem
              key={g.id}
              selected={selectedGroup === g.id}
              onClick={() => handleGroupSelect(g.id)}>
              {g.name}
            </ListItem>
          ))}
        </List>
      </Content>
      <Footer>
        <Button type="default" onClick={handleCancel}>
          キャンセル
        </Button>
        <Button type="primary" onClick={handleSave}>
          確認
        </Button>
      </Footer>
    </Wrapper>
  );
};

export default GroupSelect;

const Wrapper = styled.div`
  height: 100%;
  width: 100%;
  background: #ffffff;
  border-radius: 2px;
`;

const Header = styled.div`
  display: flex;
  justify-content: space-between;
  padding: 16px 24px;
  border-bottom: 1px solid lightgrey;
`;

const Text = styled.p`
  margin: 0;
  user-select: none;
`;

const StyledIcon = styled(Icon)`
  align-self: center;
  cursor: pointer;
`;

const Content = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
  padding: 24px;
`;

const List = styled.div`
  height: 90px;
  background: #f5f5f5;
  overflow: scroll;
`;

const ListItem = styled.div<{ selected?: boolean }>`
  display: flex;
  align-items: center;
  padding: 0 12px;
  height: 26px;
  background: ${({ selected }) => (selected ? "#1890FF" : "#ffffff")};
  color: ${({ selected }) => (selected ? "#ffffff" : "#00000")};
  box-shadow: 0px 2px 8px rgba(0, 0, 0, 0.15);
  cursor: pointer;
  transition: 0.3s background;
  user-select: none;

  :hover {
    ${({ selected }) => !selected && "background: #BAE7FF"};
  }
`;

const Footer = styled.div`
  display: flex;
  justify-content: flex-end;
  gap: 8px;
  height: 52px;
  padding: 10px;
  border-top: 1px solid lightgrey;
`;
