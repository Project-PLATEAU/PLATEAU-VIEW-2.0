import styled from "@emotion/styled";
import { FocusEventHandler, useCallback, useState } from "react";

import Badge from "@reearth-cms/components/atoms/Badge";
import Button from "@reearth-cms/components/atoms/Button";
import Select, { SelectProps } from "@reearth-cms/components/atoms/Select";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";
import SidebarCard from "@reearth-cms/components/molecules/Request/Details/SidebarCard";
import { Request, RequestUpdatePayload } from "@reearth-cms/components/molecules/Request/types";
import { Member } from "@reearth-cms/components/molecules/Workspace/types";
import { useT } from "@reearth-cms/i18n";
import { dateTimeFormat } from "@reearth-cms/utils/format";

export type Props = {
  currentRequest?: Request;
  workspaceUserMembers: Member[];
  onRequestUpdate: (data: RequestUpdatePayload) => Promise<void>;
};

const RequestSidebarWrapper: React.FC<Props> = ({
  currentRequest,
  workspaceUserMembers,
  onRequestUpdate,
}) => {
  const t = useT();
  const formattedCreatedAt = dateTimeFormat(currentRequest?.createdAt);
  const [selectedReviewers, setSelectedReviewers] = useState<string[]>([]);
  const [viewReviewers, toggleViewReviewers] = useState<boolean>(false);
  const currentReviewers = currentRequest?.reviewers;
  const reviewers: SelectProps["options"] = [];
  // TODO: this needs performance improvement
  workspaceUserMembers
    ?.filter(
      member =>
        currentReviewers?.findIndex(currentReviewer => currentReviewer.id === member.userId) === -1,
    )
    .forEach(member => {
      reviewers.push({
        label: member.user.name,
        value: member.userId,
        name: member.user.name,
      });
    });

  const displayViewReviewers = () => {
    toggleViewReviewers(true);
  };

  const hideViewReviewers = () => {
    toggleViewReviewers(false);
  };

  const handleSubmit: FocusEventHandler<HTMLElement> | undefined = useCallback(async () => {
    const requestId = currentRequest?.id;
    if (!requestId || selectedReviewers.length === 0) {
      hideViewReviewers();
      return;
    }

    const currentReviewersId = currentReviewers?.map(reviewer => reviewer.id) ?? [];
    const reviewersId: string[] | undefined = [...currentReviewersId, ...selectedReviewers];

    try {
      await onRequestUpdate?.({
        requestId: requestId,
        title: currentRequest?.title,
        description: currentRequest?.description,
        state: currentRequest?.state,
        reviewersId: reviewersId,
      });
    } catch (error) {
      console.error("Validate Failed:", error);
    } finally {
      hideViewReviewers();
    }
  }, [
    currentRequest?.description,
    currentRequest?.id,
    currentRequest?.state,
    currentRequest?.title,
    currentReviewers,
    onRequestUpdate,
    selectedReviewers,
  ]);

  return (
    <SideBarWrapper>
      <SidebarCard title={t("State")}>
        <Badge
          color={
            currentRequest?.state === "APPROVED"
              ? "#52C41A"
              : currentRequest?.state === "CLOSED"
              ? "#F5222D"
              : currentRequest?.state === "WAITING"
              ? "#FA8C16"
              : ""
          }
          text={currentRequest?.state}
        />
      </SidebarCard>
      <SidebarCard title={t("Created By")}>
        <UserAvatar username={currentRequest?.createdBy?.name} />
      </SidebarCard>
      <SidebarCard title={t("Reviewer")}>
        <div style={{ display: "flex", margin: "4px 0" }}>
          {currentRequest?.reviewers.map((reviewer, index) => (
            <UserAvatar username={reviewer.name} key={index} style={{ marginRight: "8px" }} />
          ))}
        </div>
        <Select
          placeholder={t("Reviewer")}
          mode="multiple"
          options={reviewers}
          filterOption={(input, option) =>
            option?.name.toLowerCase().indexOf(input.toLowerCase()) >= 0
          }
          style={{ width: "100%", display: viewReviewers ? "initial" : "none" }}
          onChange={setSelectedReviewers}
          onBlur={handleSubmit}
          allowClear
        />
        <div style={{ display: viewReviewers ? "none" : "flex", justifyContent: "end" }}>
          <Button type="link" onClick={displayViewReviewers} style={{ paddingRight: "0" }}>
            Assign to
          </Button>
        </div>
      </SidebarCard>
      <SidebarCard title={t("Created Time")}>{formattedCreatedAt}</SidebarCard>
    </SideBarWrapper>
  );
};

const SideBarWrapper = styled.div`
  background-color: #fafafa;
  padding: 8px;
  width: 272px;
`;

export default RequestSidebarWrapper;
