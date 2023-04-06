import { DataCatalogGroup, DataCatalogItem } from "@web/extensions/sidebar/core/types";
import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import { checkKeyPress } from "@web/extensions/sidebar/utils";
import { getNameFromPath } from "@web/extensions/sidebar/utils/file";
import { Icon, Input } from "@web/sharedComponents";
import Popconfirm, { PopconfirmProps } from "@web/sharedComponents/Popconfirm";
import { styled } from "@web/theme";
import { ComponentType, useCallback, useMemo, ChangeEvent, useState } from "react";

export type Props = {
  dataset: DataCatalogItem | DataCatalogGroup | UserDataItem;
  isShareable?: boolean;
  isPublishable?: boolean;
  addDisabled: boolean;
  inEditor?: boolean;
  requireLayerName?: boolean;
  contentSection?: ComponentType;
  onDatasetAdd: (dataset: DataCatalogItem | UserDataItem, keepModalOpen?: boolean) => void;
  onDatasetPublish?: (dataID: string, publish: boolean) => void;
};

const showShareButton = false; // This code can be removed when decision about share button is made

const DatasetDetails: React.FC<Props> = ({
  dataset,
  isShareable,
  isPublishable,
  addDisabled,
  inEditor,
  requireLayerName,
  contentSection: ContentSection,
  onDatasetAdd,
  onDatasetPublish,
}) => {
  const published = useMemo(() => (dataset as DataCatalogItem).public, [dataset]);
  const [layers, setLayers] = useState<string[]>([]);

  const handleDatasetPublish = useCallback(() => {
    if (!("dataID" in dataset)) return;
    const datasetToUpdate = dataset as DataCatalogItem;
    onDatasetPublish?.(datasetToUpdate.dataID, !datasetToUpdate.public);
  }, [dataset, onDatasetPublish]);

  const handleDatasetAdd = useCallback(
    (e: React.MouseEvent<HTMLButtonElement>) => {
      if (!dataset || !("dataID" in dataset) || addDisabled) return;
      const terminalDataset = dataset;
      if (layers.length) terminalDataset.layers = layers;

      const keyPressed = checkKeyPress(e, ["shift", "meta", "ctrl"]);
      onDatasetAdd(terminalDataset, keyPressed);
    },
    [dataset, addDisabled, layers, onDatasetAdd],
  );

  const handleLayersAddOnDataset = useCallback(
    (event: ChangeEvent<HTMLInputElement>) => {
      if (!dataset || addDisabled) return;
      const newValue = event.target.value;
      const newLayersArray = newValue.split(",");
      setLayers(newLayersArray);
    },
    [addDisabled, dataset],
  );

  const popConfirmProps: PopconfirmProps = {
    title: (
      <>
        <PopConfirmText>本当に公開してもよろしいですか？</PopConfirmText>
        {isPublishable && (
          <PopConfirmText>
            この都市のCityGMLなどのPLATEAUデータがG空間情報センターにも公開されます。
          </PopConfirmText>
        )}
      </>
    ),
    placement: "topRight",
    onConfirm: handleDatasetPublish,
    okText: "OK",
    cancelText: "キャンセル",
    okButtonProps: {
      style: {
        backgroundColor: "#00BEBE",
        width: "48%",
        border: "none",
      },
    },
    cancelButtonProps: {
      type: "primary",
      style: {
        color: "#000000d9",
        backgroundColor: "#FFFFFF",
        width: "48%",
        border: "1px solid #D9D9D9",
      },
    },
  };

  const title = useMemo(() => getNameFromPath(dataset.name), [dataset.name]);

  return (
    <>
      <TopWrapper>
        <HeaderWrapper>
          <Title>{title}</Title>
          {"dataID" in dataset &&
            inEditor &&
            (!published ? (
              <Popconfirm {...popConfirmProps}>
                <PublishButton published={published}>
                  <HoverText published={published}>公開</HoverText>
                  <Text published={published}>未公開</Text>
                </PublishButton>
              </Popconfirm>
            ) : (
              <PublishButton published={published} onClick={handleDatasetPublish}>
                <HoverText published={published}>未公開</HoverText>
                <Text published={published}>公開済み</Text>
              </PublishButton>
            ))}
        </HeaderWrapper>
        {"dataID" in dataset && (
          <>
            <ButtonWrapper>
              <AddButton disabled={addDisabled} onClick={handleDatasetAdd}>
                {!addDisabled && <Icon icon="plusCircle" />}
                {addDisabled ? "シーンに追加済み" : "シーンに追加"}
              </AddButton>
              {showShareButton && (
                <ShareButton isShareable={isShareable}>
                  <Icon icon="share" />
                  シェア
                </ShareButton>
              )}
            </ButtonWrapper>
            {requireLayerName && (
              <LayerNamesWrapper>
                <Text>表示したいレイヤー名を入力してください。</Text>
                <Input placeholder="レイヤー名" onChange={handleLayersAddOnDataset} />
              </LayerNamesWrapper>
            )}
          </>
        )}
      </TopWrapper>
      {ContentSection && (
        <Wrapper>
          <ContentSection />
        </Wrapper>
      )}
    </>
  );
};

export default DatasetDetails;

const Wrapper = styled.div`
  padding: 16px 24px;
`;

const TopWrapper = styled(Wrapper)`
  border-bottom: 1px solid #c7c5c5;
`;

const HeaderWrapper = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 12px;
  margin-bottom: 16px;
`;

const Title = styled.p`
  font-size: 16px;
  font-weight: 700;
  line-height: 22px;
  margin: 0;
`;

const ButtonWrapper = styled.div`
  display: flex;
  justify-content: space-between;
  gap: 12px;
`;

const LayerNamesWrapper = styled.div`
  margin-bottom: 16px;
`;

const BaseButton = styled.button<{ disabled?: boolean }>`
  display: flex;
  justify-content: center;
  align-items: center;
  gap: 8px;
  height: 40px;
  font-weight: 500;
  border-radius: 4px;
  border: 1px solid #e6e6e6;
`;

const AddButton = styled(BaseButton)<{ disabled?: boolean }>`
  color: ${({ disabled }) => (disabled ? "grey" : "#ffffff")};
  background-color: ${({ disabled }) => (disabled ? "#dcdcdc" : "#00bebe")};
  flex: 1;
  ${({ disabled }) => !disabled && "cursor: pointer;"}
`;

const ShareButton = styled(BaseButton)<{ isShareable?: boolean }>`
  display: ${({ isShareable }) => (isShareable !== false ? "flex" : "none")};
  color: #00bebe;
  background-color: #ffffff;
  flex: 1;
  cursor: pointer;
`;

const PublishButton = styled(BaseButton)<{ published?: boolean }>`
  display: flex;
  min-width: 120px;
  color: #ffffff;
  background-color: ${({ published }) => (published ? "#00bebe" : "#bfbfbf")};
  cursor: pointer;

  &:hover {
    background-color: ${({ published }) => (published ? "#bfbfbf" : "#00bebe")};
  }
`;

const HoverText = styled.p<{ published?: boolean }>`
  display: none;
  margin-bottom: 0;
  ${PublishButton}:hover & {
    display: initial;
  }
`;

const Text = styled.p<{ published?: boolean }>`
  display: initial;
  margin-bottom: 0;
  ${PublishButton}:hover & {
    display: none;
  }
`;

const PopConfirmText = styled.p`
  margin-bottom: 0;
`;
