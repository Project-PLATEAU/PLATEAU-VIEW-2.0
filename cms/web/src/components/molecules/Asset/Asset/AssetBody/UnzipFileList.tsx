import styled from "@emotion/styled";
import { Key } from "rc-table/lib/interface";
import { useCallback, useEffect, useState } from "react";

import Icon from "@reearth-cms/components/atoms/Icon";
import Spin from "@reearth-cms/components/atoms/Spin";
import Tree, { DataNode, TreeProps } from "@reearth-cms/components/atoms/Tree";
import {
  ArchiveExtractionStatus,
  AssetFile,
} from "@reearth-cms/components/molecules/Asset/asset.type";
import { useT } from "@reearth-cms/i18n";

type Props = {
  file: AssetFile;
  assetBaseUrl: string;
  archiveExtractionStatus: ArchiveExtractionStatus;
  setAssetUrl: (url: string) => void;
};

type FileNode = DataNode & {
  file: AssetFile;
};

const UnzipFileList: React.FC<Props> = ({
  file,
  assetBaseUrl,
  archiveExtractionStatus,
  setAssetUrl,
}) => {
  const t = useT();

  const [expandedKeys, setExpandedKeys] = useState<FileNode["key"][]>(["0"]);
  const [selectedKeys, setSelectedKeys] = useState<FileNode["key"][]>([]);
  const [treeData, setTreeData] = useState<FileNode[]>([]);

  const getTreeData = useCallback(
    (assetFile: AssetFile, parentKey?: string): FileNode[] =>
      assetFile?.children?.map((file: AssetFile, index: number) => {
        let children: FileNode[] = [];
        const key = parentKey ? parentKey + "-" + index.toString() : index.toString();

        if (file.children && file.children.length > 0) {
          children = getTreeData(file, key);
        }

        return {
          title: (
            <>
              {file.name}
              <CopyIcon
                selected={selectedKeys[0] === key}
                icon="copy"
                onClick={() => {
                  navigator.clipboard.writeText(assetBaseUrl + file.path);
                }}
              />
            </>
          ),
          key: key,
          children: children,
          file: file,
        };
      }) || [],
    [selectedKeys, assetBaseUrl],
  );

  useEffect(() => {
    setTreeData(getTreeData(file));
  }, [file, getTreeData]);

  const previewFile = useCallback(
    (file: AssetFile) => {
      setAssetUrl(assetBaseUrl + file.path);
    },
    [assetBaseUrl, setAssetUrl],
  );

  const onSelect: TreeProps<FileNode>["onSelect"] = useCallback(
    (keys: Key[], { node: { file } }: { node: FileNode }) => {
      if (!keys[0] || keys[0] === selectedKeys[0]) return;
      previewFile(file);
      setSelectedKeys(keys);
    },
    [previewFile, selectedKeys],
  );

  const onExpand: TreeProps["onExpand"] = (keys: Key[]) => {
    setExpandedKeys([...keys]);
  };

  return (
    <UnzipFileListWrapper>
      {archiveExtractionStatus === "IN_PROGRESS" || archiveExtractionStatus === "PENDING" ? (
        <ExtractionInProgressWrapper>
          <Spin tip={t("Decompressing...")} size="large" />
        </ExtractionInProgressWrapper>
      ) : archiveExtractionStatus === "FAILED" ? (
        <ExtractionFailedWrapper>
          <ExtractionFailedIcon icon="closeCircle" color="#FF4D4F" size="56px" />
          <ExtractionFailedText>
            {t("Failed to decompress. Please check the file and try again.")}
          </ExtractionFailedText>
        </ExtractionFailedWrapper>
      ) : (
        <Tree
          switcherIcon={<Icon icon="caretDown" />}
          expandedKeys={[...expandedKeys]}
          selectedKeys={[...selectedKeys]}
          onSelect={onSelect}
          onExpand={onExpand}
          treeData={treeData}
          multiple={false}
          showLine
        />
      )}
    </UnzipFileListWrapper>
  );
};

const UnzipFileListWrapper = styled.div`
  height: 250px;
  overflow-y: scroll;
  background-color: #f5f5f5;
`;

const ExtractionInProgressWrapper = styled.div`
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
`;

const ExtractionFailedWrapper = styled.div`
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100%;
`;

const ExtractionFailedIcon = styled(Icon)`
  margin-bottom: 28px;
`;

const CopyIcon = styled(Icon)<{ selected?: boolean }>`
  margin-left: 16px;
  visibility: ${({ selected }) => (selected ? "visible" : "hidden")};
  &:active {
    color: #096dd9;
  }
`;

const ExtractionFailedText = styled.p`
  margin-bottom: 0;
  font-family: Roboto;
  font-style: normal;
  font-weight: 400;
  font-size: 14px;
  line-height: 22px;
  color: rgba(0, 0, 0, 0.85);
`;

export default UnzipFileList;
