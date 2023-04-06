export type Thread = {
  id: string;
  comments: CommentItem[];
};

export type CommentItem = {
  id: string;
  author: string;
  authorId: string;
  content: string;
  createdAt: string;
};
