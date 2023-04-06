import { User } from "../Member/types";

export type RequestState = "APPROVED" | "CLOSED" | "DRAFT" | "WAITING";

export type Request = {
  id: string;
  threadId: string;
  title: string;
  description: string;
  comments: Comment[];
  createdAt: Date;
  reviewers: User[];
  state: RequestState;
  createdBy?: User;
  updatedAt: Date;
  approvedAt?: Date;
  closedAt?: Date;
  items: {
    id: string;
    modelName?: string;
    schema?: any;
    initialValues: any;
  }[];
};

export type RequestUpdatePayload = {
  requestId: string;
  title?: string;
  description?: string;
  state?: RequestState;
  reviewersId?: string[];
  items?: {
    itemId: string;
  }[];
};

export type Comment = {
  id: string;
  author: { id?: string; name: string; type: "User" | "Integration" | null };
  content: string;
  createdAt: string;
};
