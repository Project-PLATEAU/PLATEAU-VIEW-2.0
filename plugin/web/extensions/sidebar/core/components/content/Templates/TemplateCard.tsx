import { DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import { Dropdown, Icon, Menu, Spin } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { useCallback, useMemo, useState } from "react";
import {
  Accordion,
  AccordionItem,
  AccordionItemHeading,
  AccordionItemButton,
  AccordionItemPanel,
  AccordionItemState,
} from "react-accessible-accordion";

import AddButton from "../common/DatasetCard/AddButton";
import Field from "../common/DatasetCard/Field";

import useHooks from "./hooks";

type BaseFieldType = Partial<DataCatalogItem> & {
  title?: string;
  icon?: string;
  value?: string | number;
  onClick?: () => void;
};

type Tabs = "default" | "edit";

export type Props = {
  template: Template;
  templates: Template[];
  savingTemplate: boolean;
  onTemplateSave: (template: Template) => Promise<void>;
  onTemplateUpdate?: (template: Template) => void;
};
const TemplateCard: React.FC<Props> = ({
  template,
  templates,
  savingTemplate,
  onTemplateSave,
  onTemplateUpdate,
}) => {
  const [currentTab, changeTab] = useState<Tabs>("edit");
  const [hidden, setHidden] = useState(false);

  const [editTitle, setEditTitle] = useState(false);

  const {
    fieldComponentsList,
    handleFieldUpdate,
    handleFieldRemove,
    handleMoveFieldUp,
    handleMoveFieldDown,
    handleGroupsUpdate,
  } = useHooks({
    template,
    onTemplateUpdate,
  });

  const baseFields: BaseFieldType[] = useMemo(
    () => [
      {
        id: "zoom",
        title: "カメラ",
        icon: "mapPin",
        value: 1,
      },
      { id: "about", title: "About Data", icon: "about", value: "www.plateau.org/data-url" },
      {
        id: "remove",
        icon: "trash",
      },
    ],
    [],
  );

  const handleTabChange: React.MouseEventHandler<HTMLParagraphElement> = useCallback(e => {
    e.stopPropagation();
    changeTab(e.currentTarget.id as Tabs);
  }, []);

  const handleTemplateSave = useCallback(() => {
    onTemplateSave(template);
  }, [template, onTemplateSave]);

  const handleTemplateNameUpdate = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      e.stopPropagation();
      onTemplateUpdate?.({ ...template, name: e.currentTarget.value });
    },
    [template, onTemplateUpdate],
  );

  const handleToggleTitleEdit = useCallback(
    (e: React.MouseEvent<HTMLDivElement, MouseEvent> | undefined) => {
      e?.stopPropagation();
      setEditTitle(!editTitle);
    },
    [editTitle],
  );

  const menuGenerator = (menuItems: { [key: string]: any }) => (
    <Menu>
      {Object.keys(menuItems).map(i => {
        if (menuItems[i].fields) {
          return (
            <Menu.Item key={menuItems[i].key}>
              <Dropdown
                overlay={menuGenerator(menuItems[i].fields)}
                placement="bottom"
                trigger={["click"]}
                getPopupContainer={trigger => trigger.parentElement ?? document.body}>
                <div onClick={e => e.stopPropagation()}>
                  <p style={{ margin: 0 }}>{menuItems[i].name}</p>
                </div>
              </Dropdown>
            </Menu.Item>
          );
        } else {
          return (
            <Menu.Item key={i} onClick={menuItems[i]?.onClick}>
              <p style={{ margin: 0 }}>{menuItems[i].name}</p>
            </Menu.Item>
          );
        }
      })}
    </Menu>
  );

  return (
    <StyledAccordionComponent allowZeroExpanded preExpanded={["templatecard"]}>
      <AccordionItem uuid="templatecard">
        <AccordionItemState>
          {({ expanded }) => (
            <Header expanded={expanded}>
              <StyledAccordionItemButton>
                <HeaderContents>
                  <LeftMain>
                    <Icon
                      icon={hidden ? "hidden" : "visible"}
                      size={20}
                      onClick={e => {
                        e?.stopPropagation();
                        setHidden(!hidden);
                      }}
                    />
                    <NameWrapper>
                      {editTitle ? (
                        <input
                          defaultValue={template.name}
                          onChange={handleTemplateNameUpdate}
                          onClick={e => e.stopPropagation()}
                        />
                      ) : (
                        <Title>{template.name}</Title>
                      )}
                      {currentTab === "edit" && (
                        <EditIcon icon="edit" size={16} onClick={handleToggleTitleEdit} />
                      )}
                    </NameWrapper>
                  </LeftMain>
                  <ArrowIcon icon="arrowDown" size={16} expanded={expanded} />
                </HeaderContents>
                {expanded && (
                  <TabWrapper>
                    <Tab id="default" selected={currentTab === "default"} onClick={handleTabChange}>
                      公開
                    </Tab>
                    <Tab id="edit" selected={currentTab === "edit"} onClick={handleTabChange}>
                      設定
                    </Tab>
                  </TabWrapper>
                )}
              </StyledAccordionItemButton>
            </Header>
          )}
        </AccordionItemState>
        <BodyWrapper>
          <Content>
            {baseFields.map((field, idx) => (
              <BaseField key={idx} onClick={field.onClick}>
                {field.icon && <Icon icon={field.icon} size={20} color="#00BEBE" />}
                {field.title && <FieldName>{field.title}</FieldName>}
              </BaseField>
            ))}
            {template.components?.map((c, idx) => (
              <Field
                key={idx}
                index={idx}
                field={c}
                isActive={true}
                editMode={currentTab === "edit"}
                templates={templates}
                onUpdate={handleFieldUpdate}
                onRemove={handleFieldRemove}
                onMoveUp={handleMoveFieldUp}
                onMoveDown={handleMoveFieldDown}
                onGroupsUpdate={handleGroupsUpdate}
              />
            ))}
          </Content>
          {currentTab === "edit" && (
            <>
              <StyledAddButton text="フィルドを追加" items={menuGenerator(fieldComponentsList)} />
              <SaveButton onClick={handleTemplateSave} disabled={savingTemplate}>
                <Icon icon="save" size={14} />
                <Text>保存</Text>
              </SaveButton>
              {savingTemplate && (
                <Loading>
                  <Spin />
                </Loading>
              )}
            </>
          )}
        </BodyWrapper>
      </AccordionItem>
    </StyledAccordionComponent>
  );
};

export default TemplateCard;

const StyledAccordionComponent = styled(Accordion)`
  width: 100%;
  border-radius: 4px;
  box-shadow: 1px 2px 4px rgba(0, 0, 0, 0.25);
  margin: 8px 0;
  background: #ffffff;
`;

const Header = styled(AccordionItemHeading)<{ expanded?: boolean }>`
  border-bottom-width: 1px;
  border-bottom-style: solid;
  border-bottom-color: transparent;
  ${({ expanded }) => expanded && "border-bottom-color: #e0e0e0;"}
`;

const StyledAccordionItemButton = styled(AccordionItemButton)`
  display: flex;
  flex-direction: column;
`;

const HeaderContents = styled.div`
  display: flex;
  align-items: center;
  height: 46px;
  padding: 0 12px;
  outline: none;
  cursor: pointer;
`;

const BodyWrapper = styled(AccordionItemPanel)<{ noTransition?: boolean }>`
  position: relative;
  width: 100%;
  border-radius: 0px 0px 4px 4px;
  background: #fafafa;
  padding: 12px;
`;

const LeftMain = styled.div`
  flex: 1;
  display: flex;
  align-items: center;
  gap: 8px;
`;

const Title = styled.p`
  margin: 0;
  font-size: 16px;
  text-overflow: ellipsis;
  overflow: hidden;
  max-width: 230px;
  white-space: nowrap;
`;

const Content = styled.div`
  display: flex;
  align-content: center;
  align-items: center;
  flex-wrap: wrap;
  gap: 6px;
`;

const BaseField = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  flex: 1 0 auto;
  padding: 8px;
  background: #ffffff;
  border: 1px solid #e6e6e6;
  border-radius: 4px;
  cursor: pointer;

  :hover {
    background: #f4f4f4;
  }
`;

const ArrowIcon = styled(Icon)<{ expanded?: boolean }>`
  transition: transform 0.15s ease;
  transform: ${({ expanded }) => !expanded && "rotate(90deg)"};
`;

const FieldName = styled.p`
  margin: 0;
`;

const TabWrapper = styled.div`
  display: flex;
  gap: 12px;
  padding: 0 12px;
`;

const Tab = styled.p<{ selected?: boolean }>`
  margin: 0;
  padding: 0 0 10px 0;
  border-bottom-width: 2px;
  border-bottom-style: solid;
  border-bottom-color: ${({ selected }) => (selected ? "#1890FF" : "transparent")};
  color: ${({ selected }) => (selected ? "#1890FF" : "inherit")};
  cursor: pointer;
`;

const StyledAddButton = styled(AddButton)`
  margin-top: 12px;
`;

const SaveButton = styled.div<{ disabled?: boolean }>`
  margin-top: 12px;
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  background: #ffffff;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 5px;
  height: 32px;
  cursor: pointer;

  :hover {
    background: #f4f4f4;
  }
  ${({ disabled }) =>
    disabled &&
    `
      color: rgb(209, 209, 209);
      pointer-events: none;
    `}
`;

const Text = styled.p`
  margin: 0;
  line-height: 15px;
`;

const NameWrapper = styled.div`
  display: flex;
  align-items: center;
  gap: 5px;
`;

const EditIcon = styled(Icon)`
  padding: 3px;

  :hover {
    cursor: pointer;
  }
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
