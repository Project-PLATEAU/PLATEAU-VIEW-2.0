import AddButton from "@web/extensions/sidebar/core/components/content/common/DatasetCard/AddButton";
import { Icon, Dropdown, Menu } from "@web/sharedComponents";
import { styled } from "@web/theme";

import { BaseFieldProps } from "../../types";

import useHooks from "./hooks";

const SwitchGroup: React.FC<BaseFieldProps<"switchGroup">> = ({
  value,
  editMode,
  selectedGroup,
  onUpdate,
  onCurrentGroupUpdate,
}) => {
  const {
    title,
    groupItems,
    fieldGroups,
    currentGroup,
    handleTitleChange,
    handleGroupChoose,
    handleItemGroupChange,
    handleItemTitleChange,
    handleItemAdd,
    handleItemRemove,
    handleItemMoveUp,
    handleItemMoveDown,
  } = useHooks({
    value,
    selectedGroup,
    onUpdate,
    onCurrentGroupUpdate,
  });

  const uiMenu = (
    <Menu
      selectable
      items={groupItems?.map(gi => {
        return {
          key: gi.id,
          label: (
            <p style={{ margin: 0 }} onClick={() => handleGroupChoose(gi.id)}>
              {gi.title}
            </p>
          ),
        };
      })}
    />
  );

  const editGroupMenu = (groupItemIndex: number) => (
    <Menu
      selectable
      selectedKeys={[groupItems[groupItemIndex].fieldGroupID]}
      items={fieldGroups.map(fg => {
        return {
          key: fg.id,
          label: (
            <p style={{ margin: 0 }} onClick={() => handleItemGroupChange(groupItemIndex, fg.id)}>
              {fg.name}
            </p>
          ),
        };
      })}
    />
  );

  return editMode ? (
    <Wrapper>
      <Field>
        <FieldTitle>タイトル</FieldTitle>
        <FieldValue>
          <TextInput defaultValue={title} onChange={handleTitleChange} />
        </FieldValue>
      </Field>
      <AddButton text="Add Item" onClick={handleItemAdd} />
      {value.groups.map((g, idx) => (
        <Item key={idx}>
          <ItemControls>
            <Icon icon="arrowUpThin" size={16} onClick={() => handleItemMoveUp(idx)} />
            <Icon icon="arrowDownThin" size={16} onClick={() => handleItemMoveDown(idx)} />
            <TrashIcon icon="trash" size={16} onClick={() => handleItemRemove(g.id)} />
          </ItemControls>
          <Field>
            <FieldTitle>グループ</FieldTitle>
            <FieldValue>
              <Dropdown
                overlay={editGroupMenu(idx)}
                placement="bottom"
                trigger={["click"]}
                getPopupContainer={trigger => trigger.parentElement ?? document.body}>
                <StyledDropdownButton>
                  <p style={{ margin: 0 }}>
                    {fieldGroups?.find(fg => fg.id === g.fieldGroupID)?.name ?? "-"}
                  </p>
                  <StyledIcon icon="arrowDownSimple" size={12} />
                </StyledDropdownButton>
              </Dropdown>
            </FieldValue>
          </Field>
          <Field>
            <FieldTitle>名前</FieldTitle>
            <FieldValue>
              <TextInput
                defaultValue={g.title}
                onChange={(e: React.ChangeEvent<HTMLInputElement>) => {
                  handleItemTitleChange(e.target.value, idx);
                }}
              />
            </FieldValue>
          </Field>
        </Item>
      ))}
    </Wrapper>
  ) : (
    <Wrapper>
      <Field>
        <FieldTitle>{title}</FieldTitle>
        <FieldValue>
          <Dropdown overlay={uiMenu} placement="bottom" trigger={["click"]}>
            <StyledDropdownButton>
              <p style={{ margin: 0 }}>{currentGroup ? currentGroup.title : "-"}</p>
              <StyledIcon icon="arrowDownSimple" size={12} />
            </StyledDropdownButton>
          </Dropdown>
        </FieldValue>
      </Field>
    </Wrapper>
  );
};

export default SwitchGroup;

const Wrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const StyledDropdownButton = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  width: 100%;
  align-content: center;
  padding: 0 16px;
  cursor: pointer;
`;

const StyledIcon = styled(Icon)`
  font-size: 0;
`;

const Text = styled.p`
  margin: 0;
`;

const TrashIcon = styled(Icon)<{ disabled?: boolean }>`
  ${({ disabled }) =>
    disabled &&
    `
      color: rgb(209, 209, 209);
      pointer-events: none;
    `}
`;

const Field = styled.div<{ gap?: number }>`
  display: flex;
  align-items: center;
  ${({ gap }) => gap && `gap: ${gap}px;`}
  height: 32px;
`;

const FieldTitle = styled(Text)`
  width: 82px;
`;

const FieldValue = styled.div`
  position: relative;
  display: flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  flex: 1;
  height: 100%;
  width: 100%;
`;

const TextInput = styled.input.attrs({ type: "text" })`
  height: 100%;
  width: 100%;
  flex: 1;
  padding: 0 12px;
  border: none;
  outline: none;

  :focus {
    border: none;
  }
`;

const Item = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 8px;
`;

const ItemControls = styled.div`
  display: flex;
  justify-content: right;
  gap: 4px;
  cursor: pointer;
`;
