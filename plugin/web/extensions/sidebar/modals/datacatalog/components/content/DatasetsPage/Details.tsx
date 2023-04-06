import { DataCatalogGroup, DataCatalogItem } from "@web/extensions/sidebar/core/types";
import DetailsComponent from "@web/extensions/sidebar/modals/datacatalog/components/content/DatasetDetails";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

import { UserDataItem } from "../../../types";

import { Tag as TagType } from "./Tags";
// import Tags, {Tag as TagType} from "./Tags";

export type Tag = TagType;

export type Props = {
  dataset?: DataCatalogItem | DataCatalogGroup;
  isMobile?: boolean;
  inEditor?: boolean;
  addDisabled: (dataID: string) => boolean;
  onTagSelect?: (tag: TagType) => void;
  onDatasetAdd: (dataset: DataCatalogItem | UserDataItem, keepModalOpen?: boolean) => void;
  onDatasetPublish?: (dataID: string, publish: boolean) => void;
};

const DatasetDetails: React.FC<Props> = ({
  dataset,
  // isMobile,
  inEditor,
  addDisabled,
  // onTagSelect,
  onDatasetAdd,
  onDatasetPublish,
}) => {
  // const datasetTags = useMemo(
  //   () => (dataset?.type !== "group" ? dataset?.tags?.map(tag => tag) : undefined),
  //   [dataset],
  // );

  const ContentComponent: React.FC = () => (
    <>
      {/* {!isMobile && <Tags tags={datasetTags} onTagSelect={onTagSelect} />} */}
      {dataset && (
        <Content>
          {dataset?.desc
            ?.split(
              /(https?:\/\/[-a-zA-Z0-9@:%._+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b[-a-zA-Z0-9(@:%_+.~#?&//=]*)/,
            )
            .map((e, i) =>
              (i + 1) % 2 === 0 ? (
                <a key={i} onClick={() => window.open(e, "_blank")}>
                  {e}
                  <LinkIcon icon="externalLink" size={16} />
                </a>
              ) : (
                <span key={i}>{e}</span>
              ),
            )}
        </Content>
      )}
    </>
  );
  return dataset && ("dataID" in dataset || dataset.desc) ? (
    <DetailsComponent
      dataset={dataset}
      addDisabled={"dataID" in dataset && addDisabled(dataset.dataID)}
      inEditor={inEditor}
      isPublishable={"dataID" in dataset && !!dataset.itemId}
      contentSection={ContentComponent}
      onDatasetAdd={onDatasetAdd}
      onDatasetPublish={onDatasetPublish}
    />
  ) : (
    <NoData>
      <NoDataMain>
        <Icon icon="empty" size={64} />
        <br />
        <StyledP>データセットを選択してください</StyledP>
      </NoDataMain>
      <NoDataFooter
        onClick={() =>
          window.open("https://www.geospatial.jp/ckan/dataset/plateau-tokyo23ku", "_blank")
        }>
        <Icon icon="newPage" size={16} />
        <StyledP> オープンデータ・ダウンロード(G空間情報センターへのリンク)</StyledP>
      </NoDataFooter>
    </NoData>
  );
};

export default DatasetDetails;

const NoData = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  color: rgba(0, 0, 0, 0.25);
  height: calc(100% - 24px);
  margin-bottom: 24px;
`;

const NoDataMain = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  flex: 1;
  flex-direction: column;
`;

const NoDataFooter = styled.div`
  display: flex;
  align-items: center;
  gap: 4px;
  margin: 0;
  color: #00bebe;
  cursor: pointer;
`;

const StyledP = styled.p`
  margin: 0;
  text-align: center;
`;

const Content = styled.div`
  margin-top: 16px;
  white-space: pre-wrap;
  a {
    color: #00bebe;
    svg {
      transform: translateY(2px);
    }
  }
`;

const LinkIcon = styled(Icon)`
  display: inline;
  padding-left: 4px;
`;
