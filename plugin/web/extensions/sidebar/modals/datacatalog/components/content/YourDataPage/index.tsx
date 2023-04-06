import { UserDataItem } from "@web/extensions/sidebar/modals/datacatalog/types";
import { useCallback, useState } from "react";

import PageLayout from "../PageLayout";

import Details from "./Details";
import FileSelectPane from "./FileSelect";

export type Props = {
  onDatasetAdd: (dataset: UserDataItem, keepModalOpen?: boolean) => void;
};

const YourDataPage: React.FC<Props> = ({ onDatasetAdd }) => {
  const [selectedDataset, setDataset] = useState<UserDataItem>();
  const [requireLayerName, setRequireLayerName] = useState<boolean>(false);

  const handleOpenDetails = useCallback((data?: UserDataItem, needLayerName?: boolean) => {
    setDataset(data);
    setRequireLayerName(!!needLayerName);
  }, []);

  return (
    <PageLayout
      left={<FileSelectPane onOpenDetails={handleOpenDetails} />}
      right={
        <Details
          isShareable={false}
          requireLayerName={requireLayerName}
          dataset={selectedDataset}
          onDatasetAdd={onDatasetAdd}
        />
      }
    />
  );
};

export default YourDataPage;
