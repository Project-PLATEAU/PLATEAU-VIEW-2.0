import { BuildingSearch, DataCatalogItem, Template } from "@web/extensions/sidebar/core/types";
import { ReearthApi } from "@web/extensions/sidebar/types";
import { postMsg } from "@web/extensions/sidebar/utils";
import {
  getNameFromPath,
  createFileName,
  normalizeExtension,
} from "@web/extensions/sidebar/utils/file";
import { Dropdown, Icon, Menu, Spin } from "@web/sharedComponents";
import { DownloadOutlined } from "@web/sharedComponents/Icon/icons";
import { styled } from "@web/theme";
import type { Identifier, XYCoord } from "dnd-core";
import fileDownload from "js-file-download";
import { useCallback, useEffect, useMemo, useRef, useState } from "react";
import {
  Accordion,
  AccordionItem,
  AccordionItemHeading,
  AccordionItemButton,
  AccordionItemPanel,
  AccordionItemState,
} from "react-accessible-accordion";
import { useDrag, useDrop } from "react-dnd";

import AddButton from "./AddButton";
import Field from "./Field";
import { IdealZoom } from "./Field/Fields/types";
import useHooks from "./hooks";

type Tabs = "default" | "edit";

type DragItem = {
  index: number;
  id: string;
  type: string;
};

type BaseFieldType = Partial<DataCatalogItem> & {
  title?: string;
  icon?: string;
  value?: string | number;
  onClick?: () => Promise<void> | void;
};

export type Props = {
  index: number;
  id: string;
  dataset: DataCatalogItem;
  templates?: Template[];
  buildingSearch?: BuildingSearch;
  inEditor?: boolean;
  savingDataset?: boolean;
  isMobile?: boolean;
  moveCard: (dragIndex: number, hoverIndex: number) => void;
  onDatasetSave?: (dataID: string) => void;
  onDatasetRemove?: (dataID: string) => void;
  onDatasetUpdate: (dataset: DataCatalogItem, cleanseOverride?: any) => void;
  onBuildingSearch: (id: string) => void;
  onOverride?: (dataID: string, activeIDs?: string[]) => void;
  onSceneUpdate: (updatedProperties: Partial<ReearthApi>) => void;
};
const DatasetCard: React.FC<Props> = ({
  index,
  id,
  dataset,
  templates,
  buildingSearch,
  inEditor,
  savingDataset,
  isMobile,
  moveCard,
  onDatasetSave,
  onDatasetRemove,
  onDatasetUpdate,
  onBuildingSearch,
  onOverride,
  onSceneUpdate,
}) => {
  const [currentTab, changeTab] = useState<Tabs>("default");

  const dragRef = useRef<HTMLDivElement>(null);
  const previewRef = useRef<HTMLDivElement>(null);

  const {
    activeComponentIDs,
    fieldComponentsList,
    handleFieldUpdate,
    handleFieldRemove,
    handleMoveUp,
    handleMoveDown,
    handleCurrentGroupUpdate,
    handleCurrentDatasetUpdate,
    handleGroupsUpdate,
  } = useHooks({
    dataset,
    inEditor,
    templates,
    buildingSearch,
    onDatasetUpdate,
    onOverride,
  });
  const readyMVTPosition = useRef<
    Promise<
      | {
          lng?: number;
          lat?: number;
          height?: number;
          heading?: number;
          pitch?: number;
          roll?: number;
        }
      | undefined
    >
  >();

  // Fetch mvt position
  useEffect(() => {
    const fetchMetadataJSONForMVT = async () => {
      const layer = await (() =>
        new Promise<any>(resolve => {
          let success = false;
          const handleMessage = (e: any) => {
            if (e.source !== parent) return;
            if (e.data.action === "findLayerByDataID" && e.data.payload.dataID === dataset.dataID) {
              success = true;
              removeEventListener("message", handleMessage);
              resolve(e.data.payload.layer);
              return;
            }
          };
          addEventListener("message", handleMessage);
          postMsg({
            action: "findLayerByDataID",
            payload: {
              dataID: dataset.dataID,
            },
          });
          setTimeout(() => {
            if (!success) {
              removeEventListener("message", handleMessage);
              resolve(undefined);
            }
          }, 3000);
        }))();

      if (layer?.data?.type !== "mvt") return;

      const mvtBaseURL = layer?.data?.url?.match(/(.+)(\/{z}\/{x}\/{y}.mvt)/)?.[1];
      if (!mvtBaseURL) return;

      const json = await fetch(`${mvtBaseURL}/metadata.json`).then(d => d.json());
      const center = json?.center?.split(",")?.map((s: string) => Number(s));
      if (center?.length < 2) return;

      // TODO: Add minzoom later once it is improved
      const maxzoom = json?.maxzoom;
      if (layer?.id && maxzoom) {
        postMsg({
          action: "updateMVTRaster",
          payload: {
            layerId: layer?.id,
            maxzoom,
          },
        });
      }

      return {
        lng: center[0],
        lat: center[1],
        height: 30000,
        pitch: -(Math.PI / 2),
        heading: 0,
        roll: 0,
      };
    };

    // Wait until reearth layer is overridden with updated dataset
    readyMVTPosition.current = new Promise(resolve =>
      setTimeout(
        () =>
          fetchMetadataJSONForMVT()
            .then(resolve)
            .catch(() => resolve(undefined)),
        100,
      ),
    );
  }, [dataset]);

  const baseFields: BaseFieldType[] = useMemo(
    () => [
      {
        id: "zoom",
        title: "カメラ",
        icon: "mapPin",
        value: 1,
        onClick: async () => {
          const idealZoomField = dataset.components?.find(c => c.type === "idealZoom");
          const mvtPosition = await readyMVTPosition.current;
          postMsg({
            action: "cameraFlyTo",
            payload: idealZoomField
              ? [(idealZoomField as IdealZoom).position, { duration: 2 }]
              : mvtPosition
              ? [mvtPosition, { duration: 2 }]
              : dataset.dataID,
          });
          if (isMobile) postMsg({ action: "popupClose" });
        },
      },
      {
        id: "about",
        title: "About Data",
        icon: "about",
        onClick: () => {
          if (isMobile) {
            postMsg({ action: "mobileCatalogOpen", payload: dataset });
          } else {
            postMsg({ action: "catalogModalOpen" });
          }
          postMsg({ action: "saveDataset", payload: { dataset } });
        },
      },
      {
        id: "remove",
        icon: "trash",
        onClick: () => onDatasetRemove?.(dataset.dataID),
      },
      ...(currentTab === "default" && dataset.type === "建築物モデル"
        ? [
            {
              id: "search",
              title: "データを検索",
              icon: "search",
              value: 1,
              onClick: () => {
                onBuildingSearch(dataset.dataID);
              },
            },
          ]
        : []),
    ],
    [currentTab, dataset, isMobile, onDatasetRemove, onBuildingSearch],
  );

  const handleTabChange: React.MouseEventHandler<HTMLParagraphElement> = useCallback(e => {
    e.stopPropagation();
    changeTab(e.currentTarget.id as Tabs);
  }, []);

  const handleFieldSave = useCallback(() => {
    if (!inEditor) return;
    onDatasetSave?.(dataset.dataID);
  }, [dataset.dataID, inEditor, onDatasetSave]);

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

  const title = useMemo(() => getNameFromPath(dataset.name), [dataset.name]);

  const [{ handlerId }, drop] = useDrop<DragItem, void, { handlerId: Identifier | null }>({
    accept: "card",
    collect(monitor) {
      return {
        handlerId: monitor.getHandlerId(),
      };
    },
    hover(item: DragItem, monitor) {
      if (!previewRef.current) {
        return;
      }
      const dragIndex = item.index;
      const hoverIndex = index;
      if (dragIndex === hoverIndex) {
        return;
      }
      const hoverBoundingRect = previewRef.current?.getBoundingClientRect();
      const hoverMiddleY = (hoverBoundingRect.bottom - hoverBoundingRect.top) / 2;
      const clientOffset = monitor.getClientOffset();
      const hoverClientY = (clientOffset as XYCoord).y - hoverBoundingRect.top;
      if (dragIndex < hoverIndex && hoverClientY < hoverMiddleY) {
        return;
      }
      if (dragIndex > hoverIndex && hoverClientY > hoverMiddleY) {
        return;
      }
      moveCard(dragIndex, hoverIndex);
      item.index = hoverIndex;
    },
  });

  const [{ isDragging }, drag, preview] = useDrag({
    type: "card",
    item: () => {
      return { id, index };
    },
    collect: (monitor: any) => ({
      isDragging: monitor.isDragging(),
    }),
  });

  const opacity = useMemo(() => (isDragging ? 0 : 1), [isDragging]);

  drag(dragRef);
  drop(preview(previewRef));

  const handleDataDownload = useCallback(async (url: string, name?: string) => {
    const res = await fetch(url, {
      method: "GET",
    });
    const blob = await res.blob();
    fileDownload(blob, name ?? "export-data");
  }, []);

  const handleDataExport = useCallback(async () => {
    if (dataset.selectedDataset) {
      const name = getNameFromPath(dataset.selectedDataset.name);
      const format = normalizeExtension(dataset.format);
      const filename = createFileName(name, format);
      handleDataDownload(dataset.selectedDataset.url, filename);
    } else if (dataset.url) {
      const name = getNameFromPath(dataset.name);
      const format = normalizeExtension(dataset.format);
      const filename = createFileName(name, format);
      handleDataDownload(dataset.url, filename);
    }
  }, [dataset.format, dataset.name, dataset.selectedDataset, dataset.url, handleDataDownload]);

  return (
    <div ref={previewRef} style={{ opacity }} data-handler-id={handlerId}>
      <StyledAccordionComponent allowZeroExpanded preExpanded={["datasetcard"]}>
        <AccordionItem uuid="datasetcard">
          <AccordionItemState>
            {({ expanded }) => (
              <Header expanded={expanded}>
                <StyledAccordionItemButton>
                  <HeaderContents>
                    <LeftMain>
                      <div ref={dragRef}>
                        <Icon icon="holder" cursor="move" size={20} />
                      </div>
                      <Icon
                        icon={!dataset.visible ? "hidden" : "visible"}
                        size={20}
                        onClick={e => {
                          e?.stopPropagation();
                          onDatasetUpdate({ ...dataset, visible: !dataset.visible });
                        }}
                      />
                      <Title>{title}</Title>
                    </LeftMain>
                    <ArrowIcon icon="arrowDown" size={16} expanded={expanded} />
                  </HeaderContents>
                  {inEditor && expanded && (
                    <TabWrapper>
                      <Tab
                        id="default"
                        selected={currentTab === "default"}
                        onClick={handleTabChange}>
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
              {(dataset.format === "czml" || dataset.format === "geojson") && (
                <BaseButton onClick={handleDataExport}>
                  <DownloadOutlined style={{ fontSize: 20, color: "#00BEBE", marginRight: 8 }} />
                  <Text>データをエクスポート</Text>
                </BaseButton>
              )}
              {dataset.openDataUrl && (
                <BaseButton onClick={() => window.open(dataset.openDataUrl, "_blank", "noopener")}>
                  <Text>オープンデータを入手</Text>
                </BaseButton>
              )}
              {dataset.components?.map((c, idx) => (
                <Field
                  key={c.id}
                  index={idx}
                  field={c}
                  activeIDs={activeComponentIDs}
                  isActive={!!activeComponentIDs?.find(id => id === c.id)}
                  isEditing={currentTab === "edit"}
                  dataID={dataset.dataID}
                  editMode={inEditor && currentTab === "edit"}
                  templates={templates}
                  configData={dataset.config?.data}
                  selectedGroup={dataset.selectedGroup}
                  onUpdate={handleFieldUpdate}
                  onRemove={handleFieldRemove}
                  onMoveUp={handleMoveUp}
                  onMoveDown={handleMoveDown}
                  onGroupsUpdate={handleGroupsUpdate}
                  onCurrentGroupUpdate={handleCurrentGroupUpdate}
                  onCurrentDatasetUpdate={handleCurrentDatasetUpdate}
                  onSceneUpdate={onSceneUpdate}
                />
              ))}
            </Content>
            {inEditor && currentTab === "edit" && (
              <>
                <StyledAddButton text="フィルドを追加" items={menuGenerator(fieldComponentsList)} />
                <SaveButton onClick={handleFieldSave} disabled={!!savingDataset}>
                  <Icon icon="save" size={14} />
                  <Text>保存</Text>
                </SaveButton>
                {savingDataset && (
                  <Loading>
                    <Spin />
                  </Loading>
                )}
              </>
            )}
          </BodyWrapper>
        </AccordionItem>
      </StyledAccordionComponent>
    </div>
  );
};

export default DatasetCard;

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
  height: auto;
  padding: 12px;
  gap: 12px;
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
  min-width: 0;
`;

const Title = styled.p`
  margin: 0;
  font-size: 16px;
  width: 230px;
  user-select: none;
  overflow-wrap: break-word;
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
  user-select: none;
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
  user-select: none;
`;

const StyledAddButton = styled(AddButton)`
  margin-top: 12px;
`;

const SaveButton = styled.div<{ disabled: boolean }>`
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
  user-select: none;
`;

const BaseButton = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 38px;
  width: 100%;
  background: #ffffff;
  border: 1px solid #e6e6e6;
  border-radius: 4px;
  cursor: pointer;
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
