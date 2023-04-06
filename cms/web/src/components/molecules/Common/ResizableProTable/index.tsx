import styled from "@emotion/styled";
import { useEffect, useState } from "react";

import ProTable, {
  ProColumns,
  ProTableProps,
  ParamsType,
} from "@reearth-cms/components/atoms/ProTable";
import { ResizableTitle } from "@reearth-cms/components/molecules/Common/ResizableProTable/resizable";
import type { ResizeCallbackData } from "@reearth-cms/components/molecules/Common/ResizableProTable/resizable";

export type Props = ProTableProps<Record<string, any> | any, ParamsType, "text">;

const ResizableProTable: React.FC<Props> = ({
  dataSource,
  columns,
  loading,
  options,
  toolbar,
  rowSelection,
  tableAlertOptionRender,
  pagination,
  onChange,
}) => {
  const [resizableColumns, setResizableColumns] = useState<ProColumns<any, "text">[]>([]);

  useEffect(() => {
    if (columns) {
      setResizableColumns(columns);
    }
  }, [columns, setResizableColumns]);

  const handleResize =
    (index: number) =>
    (_: React.SyntheticEvent<Element>, { size }: ResizeCallbackData) => {
      const newColumns = [...resizableColumns];
      newColumns[index] = {
        ...newColumns[index],
        width: size.width,
      };
      setResizableColumns(newColumns);
    };

  const mergeColumns: ProColumns<any, "text">[] = resizableColumns?.map((col, index) => ({
    ...col,
    onHeaderCell: column => ({
      minWidth: (column as ProColumns<any, "text"> & { minWidth: number }).minWidth,
      width: (column as ProColumns<any, "text">).width,
      onResize: handleResize(index),
    }),
  }));

  return (
    <Wrapper>
      <ProTable
        dataSource={dataSource}
        columns={mergeColumns}
        components={{
          header: {
            cell: ResizableTitle,
          },
        }}
        rowKey="id"
        search={false}
        loading={loading}
        toolbar={toolbar}
        options={options}
        tableAlertOptionRender={tableAlertOptionRender}
        rowSelection={rowSelection}
        pagination={pagination}
        onChange={onChange}
      />
    </Wrapper>
  );
};

export default ResizableProTable;

const Wrapper = styled.div`
  .ant-table.ant-table-middle {
    overflow-x: scroll;
  }
  .ant-table-content {
    width: 1px;
    max-width: 100%;
    table {
      width: 100%;
    }
  }
`;
