import RequestListMolecule from "@reearth-cms/components/molecules/Request/List";
import CommentsPanel from "@reearth-cms/components/organisms/Common/CommentsPanel";
import { useT } from "@reearth-cms/i18n";

import useHooks from "./hooks";

const RequestList: React.FC = () => {
  const t = useT();

  const {
    requests,
    loading,
    collapsedCommentsPanel,
    selectedRequest,
    // selectRequest,
    // selectedRequests,
    selection,
    setSelection,
    collapseCommentsPanel,
    handleRequestSelect,
    handleRequestsReload,
    handleRequestDelete,
    handleSearchTerm,
    handleNavigateToRequest,
    totalCount,
    searchTerm,
    reviewedByMe,
    createdByMe,
    requestState,
    page,
    pageSize,
    handleRequestTableChange,
  } = useHooks();

  return (
    <RequestListMolecule
      commentsPanel={
        <CommentsPanel
          collapsed={collapsedCommentsPanel}
          onCollapse={collapseCommentsPanel}
          emptyText={
            selectedRequest
              ? t("No comments.")
              : t("Please click the comment bubble in the table to check comments.")
          }
          comments={selectedRequest?.comments}
          threadId={selectedRequest?.threadId}
        />
      }
      requests={requests}
      onRequestSelect={handleRequestSelect}
      loading={loading}
      onRequestsReload={handleRequestsReload}
      onRequestDelete={handleRequestDelete}
      selectedRequest={selectedRequest}
      onSearchTerm={handleSearchTerm}
      selection={selection}
      setSelection={setSelection}
      onEdit={handleNavigateToRequest}
      totalCount={totalCount}
      searchTerm={searchTerm}
      reviewedByMe={reviewedByMe}
      createdByMe={createdByMe}
      requestState={requestState}
      page={page}
      onRequestTableChange={handleRequestTableChange}
      pageSize={pageSize}
    />
  );
};

export default RequestList;
