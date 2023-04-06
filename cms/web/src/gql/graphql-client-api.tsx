import { gql } from '@apollo/client';
import * as Apollo from '@apollo/client';
export type Maybe<T> = T | null;
export type InputMaybe<T> = Maybe<T>;
export type Exact<T extends { [key: string]: unknown }> = { [K in keyof T]: T[K] };
export type MakeOptional<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]?: Maybe<T[SubKey]> };
export type MakeMaybe<T, K extends keyof T> = Omit<T, K> & { [SubKey in K]: Maybe<T[SubKey]> };
const defaultOptions = {} as const;
/** All built-in and custom scalars, mapped to their actual values */
export type Scalars = {
  ID: string;
  String: string;
  Boolean: boolean;
  Int: number;
  Float: number;
  Any: any;
  Cursor: string;
  DateTime: Date;
  FileSize: number;
  Lang: string;
  URL: string;
  Upload: any;
};

export type AddCommentInput = {
  content: Scalars['String'];
  threadId: Scalars['ID'];
};

export type AddIntegrationToWorkspaceInput = {
  integrationId: Scalars['ID'];
  role: Role;
  workspaceId: Scalars['ID'];
};

export type AddUsersToWorkspaceInput = {
  users: Array<MemberInput>;
  workspaceId: Scalars['ID'];
};

export type AddUsersToWorkspacePayload = {
  __typename?: 'AddUsersToWorkspacePayload';
  workspace: Workspace;
};

export type ApproveRequestInput = {
  requestId: Scalars['ID'];
};

export enum ArchiveExtractionStatus {
  Done = 'DONE',
  Failed = 'FAILED',
  InProgress = 'IN_PROGRESS',
  Pending = 'PENDING',
  Skipped = 'SKIPPED'
}

export type Asset = Node & {
  __typename?: 'Asset';
  archiveExtractionStatus?: Maybe<ArchiveExtractionStatus>;
  createdAt: Scalars['DateTime'];
  createdBy: Operator;
  createdById: Scalars['ID'];
  createdByType: OperatorType;
  id: Scalars['ID'];
  items?: Maybe<Array<AssetItem>>;
  previewType?: Maybe<PreviewType>;
  project: Project;
  projectId: Scalars['ID'];
  size: Scalars['FileSize'];
  thread?: Maybe<Thread>;
  threadId: Scalars['ID'];
  url: Scalars['String'];
  uuid: Scalars['String'];
};

export type AssetConnection = {
  __typename?: 'AssetConnection';
  edges: Array<AssetEdge>;
  nodes: Array<Maybe<Asset>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type AssetEdge = {
  __typename?: 'AssetEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<Asset>;
};

export type AssetFile = {
  __typename?: 'AssetFile';
  children?: Maybe<Array<AssetFile>>;
  contentType?: Maybe<Scalars['String']>;
  name: Scalars['String'];
  path: Scalars['String'];
  size: Scalars['FileSize'];
};

export type AssetItem = {
  __typename?: 'AssetItem';
  itemId: Scalars['ID'];
  modelId: Scalars['ID'];
};

export type AssetSort = {
  direction?: InputMaybe<SortDirection>;
  sortBy: AssetSortType;
};

export enum AssetSortType {
  Date = 'DATE',
  Name = 'NAME',
  Size = 'SIZE'
}

export type Comment = {
  __typename?: 'Comment';
  author?: Maybe<Operator>;
  authorId: Scalars['ID'];
  authorType: OperatorType;
  content: Scalars['String'];
  createdAt: Scalars['DateTime'];
  id: Scalars['ID'];
  threadId: Scalars['ID'];
  workspaceId: Scalars['ID'];
};

export type CommentPayload = {
  __typename?: 'CommentPayload';
  comment: Comment;
  thread: Thread;
};

export type CreateAssetInput = {
  file?: InputMaybe<Scalars['Upload']>;
  projectId: Scalars['ID'];
  skipDecompression?: InputMaybe<Scalars['Boolean']>;
  url?: InputMaybe<Scalars['String']>;
};

export type CreateAssetPayload = {
  __typename?: 'CreateAssetPayload';
  asset: Asset;
};

export type CreateFieldInput = {
  description?: InputMaybe<Scalars['String']>;
  key: Scalars['String'];
  modelId: Scalars['ID'];
  multiple: Scalars['Boolean'];
  required: Scalars['Boolean'];
  title: Scalars['String'];
  type: SchemaFieldType;
  typeProperty: SchemaFieldTypePropertyInput;
  unique: Scalars['Boolean'];
};

export type CreateIntegrationInput = {
  description?: InputMaybe<Scalars['String']>;
  logoUrl: Scalars['URL'];
  name: Scalars['String'];
  type: IntegrationType;
};

export type CreateItemInput = {
  fields: Array<ItemFieldInput>;
  modelId: Scalars['ID'];
  schemaId: Scalars['ID'];
};

export type CreateModelInput = {
  description?: InputMaybe<Scalars['String']>;
  key?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  projectId: Scalars['ID'];
};

export type CreateProjectInput = {
  alias?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  workspaceId: Scalars['ID'];
};

export type CreateRequestInput = {
  description?: InputMaybe<Scalars['String']>;
  items: Array<RequestItemInput>;
  projectId: Scalars['ID'];
  reviewersId?: InputMaybe<Array<Scalars['ID']>>;
  state?: InputMaybe<RequestState>;
  title: Scalars['String'];
};

export type CreateThreadInput = {
  workspaceId: Scalars['ID'];
};

export type CreateWebhookInput = {
  active: Scalars['Boolean'];
  integrationId: Scalars['ID'];
  name: Scalars['String'];
  secret: Scalars['String'];
  trigger: WebhookTriggerInput;
  url: Scalars['URL'];
};

export type CreateWorkspaceInput = {
  name: Scalars['String'];
};

export type CreateWorkspacePayload = {
  __typename?: 'CreateWorkspacePayload';
  workspace: Workspace;
};

export type DecompressAssetInput = {
  assetId: Scalars['ID'];
};

export type DecompressAssetPayload = {
  __typename?: 'DecompressAssetPayload';
  asset: Asset;
};

export type DeleteAssetInput = {
  assetId: Scalars['ID'];
};

export type DeleteAssetPayload = {
  __typename?: 'DeleteAssetPayload';
  assetId: Scalars['ID'];
};

export type DeleteCommentInput = {
  commentId: Scalars['ID'];
  threadId: Scalars['ID'];
};

export type DeleteCommentPayload = {
  __typename?: 'DeleteCommentPayload';
  commentId: Scalars['ID'];
  thread: Thread;
};

export type DeleteFieldInput = {
  fieldId: Scalars['ID'];
  modelId: Scalars['ID'];
};

export type DeleteFieldPayload = {
  __typename?: 'DeleteFieldPayload';
  fieldId: Scalars['ID'];
};

export type DeleteIntegrationInput = {
  integrationId: Scalars['ID'];
};

export type DeleteIntegrationPayload = {
  __typename?: 'DeleteIntegrationPayload';
  integrationId: Scalars['ID'];
};

export type DeleteItemInput = {
  itemId: Scalars['ID'];
};

export type DeleteItemPayload = {
  __typename?: 'DeleteItemPayload';
  itemId: Scalars['ID'];
};

export type DeleteMeInput = {
  userId: Scalars['ID'];
};

export type DeleteMePayload = {
  __typename?: 'DeleteMePayload';
  userId: Scalars['ID'];
};

export type DeleteModelInput = {
  modelId: Scalars['ID'];
};

export type DeleteModelPayload = {
  __typename?: 'DeleteModelPayload';
  modelId: Scalars['ID'];
};

export type DeleteProjectInput = {
  projectId: Scalars['ID'];
};

export type DeleteProjectPayload = {
  __typename?: 'DeleteProjectPayload';
  projectId: Scalars['ID'];
};

export type DeleteRequestInput = {
  projectId: Scalars['ID'];
  requestsId: Array<Scalars['ID']>;
};

export type DeleteRequestPayload = {
  __typename?: 'DeleteRequestPayload';
  requests: Array<Scalars['ID']>;
};

export type DeleteWebhookInput = {
  integrationId: Scalars['ID'];
  webhookId: Scalars['ID'];
};

export type DeleteWebhookPayload = {
  __typename?: 'DeleteWebhookPayload';
  webhookId: Scalars['ID'];
};

export type DeleteWorkspaceInput = {
  workspaceId: Scalars['ID'];
};

export type DeleteWorkspacePayload = {
  __typename?: 'DeleteWorkspacePayload';
  workspaceId: Scalars['ID'];
};

export type FieldPayload = {
  __typename?: 'FieldPayload';
  field: SchemaField;
};

export type FieldsPayload = {
  __typename?: 'FieldsPayload';
  fields: Array<SchemaField>;
};

export type Integration = Node & {
  __typename?: 'Integration';
  config?: Maybe<IntegrationConfig>;
  createdAt: Scalars['DateTime'];
  description?: Maybe<Scalars['String']>;
  developer: User;
  developerId: Scalars['ID'];
  iType: IntegrationType;
  id: Scalars['ID'];
  logoUrl: Scalars['URL'];
  name: Scalars['String'];
  updatedAt: Scalars['DateTime'];
};

export type IntegrationConfig = {
  __typename?: 'IntegrationConfig';
  token: Scalars['String'];
  webhooks: Array<Webhook>;
};

export type IntegrationPayload = {
  __typename?: 'IntegrationPayload';
  integration: Integration;
};

export enum IntegrationType {
  Private = 'Private',
  Public = 'Public'
}

export type Item = Node & {
  __typename?: 'Item';
  assets: Array<Maybe<Asset>>;
  createdAt: Scalars['DateTime'];
  fields: Array<ItemField>;
  id: Scalars['ID'];
  integration?: Maybe<Integration>;
  integrationId?: Maybe<Scalars['ID']>;
  model: Model;
  modelId: Scalars['ID'];
  project: Project;
  projectId: Scalars['ID'];
  schema: Schema;
  schemaId: Scalars['ID'];
  status: ItemStatus;
  thread: Thread;
  threadId: Scalars['ID'];
  updatedAt: Scalars['DateTime'];
  user?: Maybe<User>;
  userId?: Maybe<Scalars['ID']>;
};

export type ItemConnection = {
  __typename?: 'ItemConnection';
  edges: Array<ItemEdge>;
  nodes: Array<Maybe<Item>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type ItemEdge = {
  __typename?: 'ItemEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<Item>;
};

export type ItemField = {
  __typename?: 'ItemField';
  schemaFieldId: Scalars['ID'];
  type: SchemaFieldType;
  value?: Maybe<Scalars['Any']>;
};

export type ItemFieldInput = {
  schemaFieldId: Scalars['ID'];
  type: SchemaFieldType;
  value: Scalars['Any'];
};

export type ItemPayload = {
  __typename?: 'ItemPayload';
  item: Item;
};

export type ItemQuery = {
  project: Scalars['ID'];
  q?: InputMaybe<Scalars['String']>;
  schema?: InputMaybe<Scalars['ID']>;
};

export type ItemSort = {
  direction?: InputMaybe<SortDirection>;
  sortBy: ItemSortType;
};

export enum ItemSortType {
  CreationDate = 'CREATION_DATE',
  ModificationDate = 'MODIFICATION_DATE'
}

export enum ItemStatus {
  Draft = 'DRAFT',
  Public = 'PUBLIC',
  PublicDraft = 'PUBLIC_DRAFT',
  PublicReview = 'PUBLIC_REVIEW',
  Review = 'REVIEW'
}

export type KeyAvailability = {
  __typename?: 'KeyAvailability';
  available: Scalars['Boolean'];
  key: Scalars['String'];
};

export type Me = {
  __typename?: 'Me';
  auths: Array<Scalars['String']>;
  email: Scalars['String'];
  id: Scalars['ID'];
  integrations: Array<Integration>;
  lang: Scalars['Lang'];
  myWorkspace: Workspace;
  myWorkspaceId: Scalars['ID'];
  name: Scalars['String'];
  theme: Theme;
  workspaces: Array<Workspace>;
};

export type MemberInput = {
  role: Role;
  userId: Scalars['ID'];
};

export type Model = Node & {
  __typename?: 'Model';
  createdAt: Scalars['DateTime'];
  description: Scalars['String'];
  id: Scalars['ID'];
  key: Scalars['String'];
  name: Scalars['String'];
  project: Project;
  projectId: Scalars['ID'];
  public: Scalars['Boolean'];
  schema: Schema;
  schemaId: Scalars['ID'];
  updatedAt: Scalars['DateTime'];
};

export type ModelConnection = {
  __typename?: 'ModelConnection';
  edges: Array<ModelEdge>;
  nodes: Array<Maybe<Model>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type ModelEdge = {
  __typename?: 'ModelEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<Model>;
};

export type ModelPayload = {
  __typename?: 'ModelPayload';
  model: Model;
};

export type Mutation = {
  __typename?: 'Mutation';
  addComment?: Maybe<CommentPayload>;
  addIntegrationToWorkspace?: Maybe<AddUsersToWorkspacePayload>;
  addUsersToWorkspace?: Maybe<AddUsersToWorkspacePayload>;
  approveRequest?: Maybe<RequestPayload>;
  createAsset?: Maybe<CreateAssetPayload>;
  createField?: Maybe<FieldPayload>;
  createIntegration?: Maybe<IntegrationPayload>;
  createItem?: Maybe<ItemPayload>;
  createModel?: Maybe<ModelPayload>;
  createProject?: Maybe<ProjectPayload>;
  createRequest?: Maybe<RequestPayload>;
  createThread?: Maybe<ThreadPayload>;
  createWebhook?: Maybe<WebhookPayload>;
  createWorkspace?: Maybe<CreateWorkspacePayload>;
  decompressAsset?: Maybe<DecompressAssetPayload>;
  deleteAsset?: Maybe<DeleteAssetPayload>;
  deleteComment?: Maybe<DeleteCommentPayload>;
  deleteField?: Maybe<DeleteFieldPayload>;
  deleteIntegration?: Maybe<DeleteIntegrationPayload>;
  deleteItem?: Maybe<DeleteItemPayload>;
  deleteMe?: Maybe<DeleteMePayload>;
  deleteModel?: Maybe<DeleteModelPayload>;
  deleteProject?: Maybe<DeleteProjectPayload>;
  deleteRequest?: Maybe<DeleteRequestPayload>;
  deleteWebhook?: Maybe<DeleteWebhookPayload>;
  deleteWorkspace?: Maybe<DeleteWorkspacePayload>;
  publishModel?: Maybe<PublishModelPayload>;
  removeIntegrationFromWorkspace?: Maybe<RemoveMemberFromWorkspacePayload>;
  removeMyAuth?: Maybe<UpdateMePayload>;
  removeUserFromWorkspace?: Maybe<RemoveMemberFromWorkspacePayload>;
  unpublishItem?: Maybe<UnpublishItemPayload>;
  updateAsset?: Maybe<UpdateAssetPayload>;
  updateComment?: Maybe<CommentPayload>;
  updateField?: Maybe<FieldPayload>;
  updateFields?: Maybe<FieldsPayload>;
  updateIntegration?: Maybe<IntegrationPayload>;
  updateIntegrationOfWorkspace?: Maybe<UpdateMemberOfWorkspacePayload>;
  updateItem?: Maybe<ItemPayload>;
  updateMe?: Maybe<UpdateMePayload>;
  updateModel?: Maybe<ModelPayload>;
  updateProject?: Maybe<ProjectPayload>;
  updateRequest?: Maybe<RequestPayload>;
  updateUserOfWorkspace?: Maybe<UpdateMemberOfWorkspacePayload>;
  updateWebhook?: Maybe<WebhookPayload>;
  updateWorkspace?: Maybe<UpdateWorkspacePayload>;
};


export type MutationAddCommentArgs = {
  input: AddCommentInput;
};


export type MutationAddIntegrationToWorkspaceArgs = {
  input: AddIntegrationToWorkspaceInput;
};


export type MutationAddUsersToWorkspaceArgs = {
  input: AddUsersToWorkspaceInput;
};


export type MutationApproveRequestArgs = {
  input: ApproveRequestInput;
};


export type MutationCreateAssetArgs = {
  input: CreateAssetInput;
};


export type MutationCreateFieldArgs = {
  input: CreateFieldInput;
};


export type MutationCreateIntegrationArgs = {
  input: CreateIntegrationInput;
};


export type MutationCreateItemArgs = {
  input: CreateItemInput;
};


export type MutationCreateModelArgs = {
  input: CreateModelInput;
};


export type MutationCreateProjectArgs = {
  input: CreateProjectInput;
};


export type MutationCreateRequestArgs = {
  input: CreateRequestInput;
};


export type MutationCreateThreadArgs = {
  input: CreateThreadInput;
};


export type MutationCreateWebhookArgs = {
  input: CreateWebhookInput;
};


export type MutationCreateWorkspaceArgs = {
  input: CreateWorkspaceInput;
};


export type MutationDecompressAssetArgs = {
  input: DecompressAssetInput;
};


export type MutationDeleteAssetArgs = {
  input: DeleteAssetInput;
};


export type MutationDeleteCommentArgs = {
  input: DeleteCommentInput;
};


export type MutationDeleteFieldArgs = {
  input: DeleteFieldInput;
};


export type MutationDeleteIntegrationArgs = {
  input: DeleteIntegrationInput;
};


export type MutationDeleteItemArgs = {
  input: DeleteItemInput;
};


export type MutationDeleteMeArgs = {
  input: DeleteMeInput;
};


export type MutationDeleteModelArgs = {
  input: DeleteModelInput;
};


export type MutationDeleteProjectArgs = {
  input: DeleteProjectInput;
};


export type MutationDeleteRequestArgs = {
  input: DeleteRequestInput;
};


export type MutationDeleteWebhookArgs = {
  input: DeleteWebhookInput;
};


export type MutationDeleteWorkspaceArgs = {
  input: DeleteWorkspaceInput;
};


export type MutationPublishModelArgs = {
  input: PublishModelInput;
};


export type MutationRemoveIntegrationFromWorkspaceArgs = {
  input: RemoveIntegrationFromWorkspaceInput;
};


export type MutationRemoveMyAuthArgs = {
  input: RemoveMyAuthInput;
};


export type MutationRemoveUserFromWorkspaceArgs = {
  input: RemoveUserFromWorkspaceInput;
};


export type MutationUnpublishItemArgs = {
  input: UnpublishItemInput;
};


export type MutationUpdateAssetArgs = {
  input: UpdateAssetInput;
};


export type MutationUpdateCommentArgs = {
  input: UpdateCommentInput;
};


export type MutationUpdateFieldArgs = {
  input: UpdateFieldInput;
};


export type MutationUpdateFieldsArgs = {
  input: Array<UpdateFieldInput>;
};


export type MutationUpdateIntegrationArgs = {
  input: UpdateIntegrationInput;
};


export type MutationUpdateIntegrationOfWorkspaceArgs = {
  input: UpdateIntegrationOfWorkspaceInput;
};


export type MutationUpdateItemArgs = {
  input: UpdateItemInput;
};


export type MutationUpdateMeArgs = {
  input: UpdateMeInput;
};


export type MutationUpdateModelArgs = {
  input: UpdateModelInput;
};


export type MutationUpdateProjectArgs = {
  input: UpdateProjectInput;
};


export type MutationUpdateRequestArgs = {
  input: UpdateRequestInput;
};


export type MutationUpdateUserOfWorkspaceArgs = {
  input: UpdateUserOfWorkspaceInput;
};


export type MutationUpdateWebhookArgs = {
  input: UpdateWebhookInput;
};


export type MutationUpdateWorkspaceArgs = {
  input: UpdateWorkspaceInput;
};

export type Node = {
  id: Scalars['ID'];
};

export enum NodeType {
  Asset = 'ASSET',
  Integration = 'Integration',
  Item = 'Item',
  Model = 'Model',
  Project = 'PROJECT',
  Request = 'REQUEST',
  Schema = 'Schema',
  User = 'USER',
  Workspace = 'WORKSPACE'
}

export type Operator = Integration | User;

export enum OperatorType {
  Integration = 'Integration',
  User = 'User'
}

export type PageInfo = {
  __typename?: 'PageInfo';
  endCursor?: Maybe<Scalars['Cursor']>;
  hasNextPage: Scalars['Boolean'];
  hasPreviousPage: Scalars['Boolean'];
  startCursor?: Maybe<Scalars['Cursor']>;
};

export type Pagination = {
  after?: InputMaybe<Scalars['Cursor']>;
  before?: InputMaybe<Scalars['Cursor']>;
  first?: InputMaybe<Scalars['Int']>;
  last?: InputMaybe<Scalars['Int']>;
  offset?: InputMaybe<Scalars['Int']>;
};

export enum PreviewType {
  Geo = 'GEO',
  Geo_3DTiles = 'GEO_3D_TILES',
  GeoMvt = 'GEO_MVT',
  Image = 'IMAGE',
  ImageSvg = 'IMAGE_SVG',
  Model_3D = 'MODEL_3D',
  Unknown = 'UNKNOWN'
}

export type Project = Node & {
  __typename?: 'Project';
  alias: Scalars['String'];
  createdAt: Scalars['DateTime'];
  description: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
  publication?: Maybe<ProjectPublication>;
  updatedAt: Scalars['DateTime'];
  workspace?: Maybe<Workspace>;
  workspaceId: Scalars['ID'];
};

export type ProjectAliasAvailability = {
  __typename?: 'ProjectAliasAvailability';
  alias: Scalars['String'];
  available: Scalars['Boolean'];
};

export type ProjectConnection = {
  __typename?: 'ProjectConnection';
  edges: Array<ProjectEdge>;
  nodes: Array<Maybe<Project>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type ProjectEdge = {
  __typename?: 'ProjectEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<Project>;
};

export type ProjectPayload = {
  __typename?: 'ProjectPayload';
  project: Project;
};

export type ProjectPublication = {
  __typename?: 'ProjectPublication';
  assetPublic: Scalars['Boolean'];
  scope: ProjectPublicationScope;
};

export enum ProjectPublicationScope {
  Limited = 'LIMITED',
  Private = 'PRIVATE',
  Public = 'PUBLIC'
}

export type PublishModelInput = {
  modelId: Scalars['ID'];
  status: Scalars['Boolean'];
};

export type PublishModelPayload = {
  __typename?: 'PublishModelPayload';
  modelId: Scalars['ID'];
  status: Scalars['Boolean'];
};

export type Query = {
  __typename?: 'Query';
  assetFile: AssetFile;
  assets: AssetConnection;
  checkModelKeyAvailability: KeyAvailability;
  checkProjectAlias: ProjectAliasAvailability;
  items: ItemConnection;
  me?: Maybe<Me>;
  models: ModelConnection;
  node?: Maybe<Node>;
  nodes: Array<Maybe<Node>>;
  projects: ProjectConnection;
  requests: RequestConnection;
  searchItem: ItemConnection;
  searchUser?: Maybe<User>;
  versionsByItem: Array<VersionedItem>;
};


export type QueryAssetFileArgs = {
  assetId: Scalars['ID'];
};


export type QueryAssetsArgs = {
  keyword?: InputMaybe<Scalars['String']>;
  pagination?: InputMaybe<Pagination>;
  projectId: Scalars['ID'];
  sort?: InputMaybe<AssetSort>;
};


export type QueryCheckModelKeyAvailabilityArgs = {
  key: Scalars['String'];
  projectId: Scalars['ID'];
};


export type QueryCheckProjectAliasArgs = {
  alias: Scalars['String'];
};


export type QueryItemsArgs = {
  pagination?: InputMaybe<Pagination>;
  schemaId: Scalars['ID'];
  sort?: InputMaybe<ItemSort>;
};


export type QueryModelsArgs = {
  pagination?: InputMaybe<Pagination>;
  projectId: Scalars['ID'];
};


export type QueryNodeArgs = {
  id: Scalars['ID'];
  type: NodeType;
};


export type QueryNodesArgs = {
  id: Array<Scalars['ID']>;
  type: NodeType;
};


export type QueryProjectsArgs = {
  pagination?: InputMaybe<Pagination>;
  workspaceId: Scalars['ID'];
};


export type QueryRequestsArgs = {
  createdBy?: InputMaybe<Scalars['ID']>;
  key?: InputMaybe<Scalars['String']>;
  pagination?: InputMaybe<Pagination>;
  projectId: Scalars['ID'];
  reviewer?: InputMaybe<Scalars['ID']>;
  sort?: InputMaybe<Sort>;
  state?: InputMaybe<Array<RequestState>>;
};


export type QuerySearchItemArgs = {
  pagination?: InputMaybe<Pagination>;
  query: ItemQuery;
  sort?: InputMaybe<ItemSort>;
};


export type QuerySearchUserArgs = {
  nameOrEmail: Scalars['String'];
};


export type QueryVersionsByItemArgs = {
  itemId: Scalars['ID'];
};

export type RemoveIntegrationFromWorkspaceInput = {
  integrationId: Scalars['ID'];
  workspaceId: Scalars['ID'];
};

export type RemoveMemberFromWorkspacePayload = {
  __typename?: 'RemoveMemberFromWorkspacePayload';
  workspace: Workspace;
};

export type RemoveMyAuthInput = {
  auth: Scalars['String'];
};

export type RemoveUserFromWorkspaceInput = {
  userId: Scalars['ID'];
  workspaceId: Scalars['ID'];
};

export type Request = Node & {
  __typename?: 'Request';
  approvedAt?: Maybe<Scalars['DateTime']>;
  closedAt?: Maybe<Scalars['DateTime']>;
  createdAt: Scalars['DateTime'];
  createdBy?: Maybe<User>;
  createdById: Scalars['ID'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  items: Array<RequestItem>;
  project?: Maybe<Project>;
  projectId: Scalars['ID'];
  reviewers: Array<User>;
  reviewersId: Array<Scalars['ID']>;
  state: RequestState;
  thread?: Maybe<Thread>;
  threadId: Scalars['ID'];
  title: Scalars['String'];
  updatedAt: Scalars['DateTime'];
  workspace?: Maybe<Workspace>;
  workspaceId: Scalars['ID'];
};

export type RequestConnection = {
  __typename?: 'RequestConnection';
  edges: Array<RequestEdge>;
  nodes: Array<Maybe<Request>>;
  pageInfo: PageInfo;
  totalCount: Scalars['Int'];
};

export type RequestEdge = {
  __typename?: 'RequestEdge';
  cursor: Scalars['Cursor'];
  node?: Maybe<Request>;
};

export type RequestItem = {
  __typename?: 'RequestItem';
  item?: Maybe<VersionedItem>;
  itemId: Scalars['ID'];
  ref?: Maybe<Scalars['String']>;
  version?: Maybe<Scalars['String']>;
};

export type RequestItemInput = {
  itemId: Scalars['ID'];
};

export type RequestPayload = {
  __typename?: 'RequestPayload';
  request: Request;
};

export enum RequestState {
  Approved = 'APPROVED',
  Closed = 'CLOSED',
  Draft = 'DRAFT',
  Waiting = 'WAITING'
}

export enum Role {
  Maintainer = 'MAINTAINER',
  Owner = 'OWNER',
  Reader = 'READER',
  Writer = 'WRITER'
}

export type Schema = Node & {
  __typename?: 'Schema';
  fields: Array<SchemaField>;
  id: Scalars['ID'];
  project: Project;
  projectId: Scalars['ID'];
};

export type SchemaField = {
  __typename?: 'SchemaField';
  createdAt: Scalars['DateTime'];
  description?: Maybe<Scalars['String']>;
  id: Scalars['ID'];
  key: Scalars['String'];
  model: Model;
  modelId: Scalars['ID'];
  multiple: Scalars['Boolean'];
  order?: Maybe<Scalars['Int']>;
  required: Scalars['Boolean'];
  title: Scalars['String'];
  type: SchemaFieldType;
  typeProperty?: Maybe<SchemaFieldTypeProperty>;
  unique: Scalars['Boolean'];
  updatedAt: Scalars['DateTime'];
};

export type SchemaFieldAsset = {
  __typename?: 'SchemaFieldAsset';
  defaultValue?: Maybe<Scalars['Any']>;
};

export type SchemaFieldAssetInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
};

export type SchemaFieldBool = {
  __typename?: 'SchemaFieldBool';
  defaultValue?: Maybe<Scalars['Any']>;
};

export type SchemaFieldBoolInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
};

export type SchemaFieldDate = {
  __typename?: 'SchemaFieldDate';
  defaultValue?: Maybe<Scalars['Any']>;
};

export type SchemaFieldDateInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
};

export type SchemaFieldInteger = {
  __typename?: 'SchemaFieldInteger';
  defaultValue?: Maybe<Scalars['Any']>;
  max?: Maybe<Scalars['Int']>;
  min?: Maybe<Scalars['Int']>;
};

export type SchemaFieldIntegerInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  max?: InputMaybe<Scalars['Int']>;
  min?: InputMaybe<Scalars['Int']>;
};

export type SchemaFieldMarkdown = {
  __typename?: 'SchemaFieldMarkdown';
  defaultValue?: Maybe<Scalars['Any']>;
  maxLength?: Maybe<Scalars['Int']>;
};

export type SchemaFieldReference = {
  __typename?: 'SchemaFieldReference';
  modelId: Scalars['ID'];
};

export type SchemaFieldReferenceInput = {
  modelId: Scalars['ID'];
};

export type SchemaFieldRichText = {
  __typename?: 'SchemaFieldRichText';
  defaultValue?: Maybe<Scalars['Any']>;
  maxLength?: Maybe<Scalars['Int']>;
};

export type SchemaFieldRichTextInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  maxLength?: InputMaybe<Scalars['Int']>;
};

export type SchemaFieldSelect = {
  __typename?: 'SchemaFieldSelect';
  defaultValue?: Maybe<Scalars['Any']>;
  values: Array<Scalars['String']>;
};

export type SchemaFieldSelectInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  values: Array<Scalars['String']>;
};

export type SchemaFieldTag = {
  __typename?: 'SchemaFieldTag';
  defaultValue?: Maybe<Scalars['Any']>;
  values: Array<Scalars['String']>;
};

export type SchemaFieldTagInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  values: Array<Scalars['String']>;
};

export type SchemaFieldText = {
  __typename?: 'SchemaFieldText';
  defaultValue?: Maybe<Scalars['Any']>;
  maxLength?: Maybe<Scalars['Int']>;
};

export type SchemaFieldTextArea = {
  __typename?: 'SchemaFieldTextArea';
  defaultValue?: Maybe<Scalars['Any']>;
  maxLength?: Maybe<Scalars['Int']>;
};

export type SchemaFieldTextAreaInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  maxLength?: InputMaybe<Scalars['Int']>;
};

export type SchemaFieldTextInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  maxLength?: InputMaybe<Scalars['Int']>;
};

export enum SchemaFieldType {
  Asset = 'Asset',
  Bool = 'Bool',
  Date = 'Date',
  Integer = 'Integer',
  MarkdownText = 'MarkdownText',
  Reference = 'Reference',
  RichText = 'RichText',
  Select = 'Select',
  Tag = 'Tag',
  Text = 'Text',
  TextArea = 'TextArea',
  Url = 'URL'
}

export type SchemaFieldTypeProperty = SchemaFieldAsset | SchemaFieldBool | SchemaFieldDate | SchemaFieldInteger | SchemaFieldMarkdown | SchemaFieldReference | SchemaFieldRichText | SchemaFieldSelect | SchemaFieldTag | SchemaFieldText | SchemaFieldTextArea | SchemaFieldUrl;

export type SchemaFieldTypePropertyInput = {
  asset?: InputMaybe<SchemaFieldAssetInput>;
  bool?: InputMaybe<SchemaFieldBoolInput>;
  date?: InputMaybe<SchemaFieldDateInput>;
  integer?: InputMaybe<SchemaFieldIntegerInput>;
  markdownText?: InputMaybe<SchemaMarkdownTextInput>;
  reference?: InputMaybe<SchemaFieldReferenceInput>;
  richText?: InputMaybe<SchemaFieldRichTextInput>;
  select?: InputMaybe<SchemaFieldSelectInput>;
  tag?: InputMaybe<SchemaFieldTagInput>;
  text?: InputMaybe<SchemaFieldTextInput>;
  textArea?: InputMaybe<SchemaFieldTextAreaInput>;
  url?: InputMaybe<SchemaFieldUrlInput>;
};

export type SchemaFieldUrl = {
  __typename?: 'SchemaFieldURL';
  defaultValue?: Maybe<Scalars['Any']>;
};

export type SchemaFieldUrlInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
};

export type SchemaMarkdownTextInput = {
  defaultValue?: InputMaybe<Scalars['Any']>;
  maxLength?: InputMaybe<Scalars['Int']>;
};

export type Sort = {
  key: Scalars['String'];
  reverted?: InputMaybe<Scalars['Boolean']>;
};

export enum SortDirection {
  Asc = 'ASC',
  Desc = 'DESC'
}

export enum Theme {
  Dark = 'DARK',
  Default = 'DEFAULT',
  Light = 'LIGHT'
}

export type Thread = {
  __typename?: 'Thread';
  comments: Array<Comment>;
  id: Scalars['ID'];
  workspace?: Maybe<Workspace>;
  workspaceId: Scalars['ID'];
};

export type ThreadPayload = {
  __typename?: 'ThreadPayload';
  thread: Thread;
};

export type UnpublishItemInput = {
  itemId: Array<Scalars['ID']>;
};

export type UnpublishItemPayload = {
  __typename?: 'UnpublishItemPayload';
  items: Array<Item>;
};

export type UpdateAssetInput = {
  id: Scalars['ID'];
  previewType?: InputMaybe<PreviewType>;
};

export type UpdateAssetPayload = {
  __typename?: 'UpdateAssetPayload';
  asset: Asset;
};

export type UpdateCommentInput = {
  commentId: Scalars['ID'];
  content: Scalars['String'];
  threadId: Scalars['ID'];
};

export type UpdateFieldInput = {
  description?: InputMaybe<Scalars['String']>;
  fieldId: Scalars['ID'];
  key?: InputMaybe<Scalars['String']>;
  modelId: Scalars['ID'];
  multiple?: InputMaybe<Scalars['Boolean']>;
  order?: InputMaybe<Scalars['Int']>;
  required?: InputMaybe<Scalars['Boolean']>;
  title?: InputMaybe<Scalars['String']>;
  typeProperty?: InputMaybe<SchemaFieldTypePropertyInput>;
  unique?: InputMaybe<Scalars['Boolean']>;
};

export type UpdateIntegrationInput = {
  description?: InputMaybe<Scalars['String']>;
  integrationId: Scalars['ID'];
  logoUrl?: InputMaybe<Scalars['URL']>;
  name?: InputMaybe<Scalars['String']>;
};

export type UpdateIntegrationOfWorkspaceInput = {
  integrationId: Scalars['ID'];
  role: Role;
  workspaceId: Scalars['ID'];
};

export type UpdateItemInput = {
  fields: Array<ItemFieldInput>;
  itemId: Scalars['ID'];
};

export type UpdateMeInput = {
  email?: InputMaybe<Scalars['String']>;
  lang?: InputMaybe<Scalars['Lang']>;
  name?: InputMaybe<Scalars['String']>;
  password?: InputMaybe<Scalars['String']>;
  passwordConfirmation?: InputMaybe<Scalars['String']>;
  theme?: InputMaybe<Theme>;
};

export type UpdateMePayload = {
  __typename?: 'UpdateMePayload';
  me: Me;
};

export type UpdateMemberOfWorkspacePayload = {
  __typename?: 'UpdateMemberOfWorkspacePayload';
  workspace: Workspace;
};

export type UpdateModelInput = {
  description?: InputMaybe<Scalars['String']>;
  key?: InputMaybe<Scalars['String']>;
  modelId: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  public: Scalars['Boolean'];
};

export type UpdateProjectInput = {
  alias?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  name?: InputMaybe<Scalars['String']>;
  projectId: Scalars['ID'];
  publication?: InputMaybe<UpdateProjectPublicationInput>;
};

export type UpdateProjectPublicationInput = {
  assetPublic?: InputMaybe<Scalars['Boolean']>;
  scope?: InputMaybe<ProjectPublicationScope>;
};

export type UpdateRequestInput = {
  description?: InputMaybe<Scalars['String']>;
  items?: InputMaybe<Array<RequestItemInput>>;
  requestId: Scalars['ID'];
  reviewersId?: InputMaybe<Array<Scalars['ID']>>;
  state?: InputMaybe<RequestState>;
  title?: InputMaybe<Scalars['String']>;
};

export type UpdateUserOfWorkspaceInput = {
  role: Role;
  userId: Scalars['ID'];
  workspaceId: Scalars['ID'];
};

export type UpdateWebhookInput = {
  active?: InputMaybe<Scalars['Boolean']>;
  integrationId: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  secret?: InputMaybe<Scalars['String']>;
  trigger?: InputMaybe<WebhookTriggerInput>;
  url?: InputMaybe<Scalars['URL']>;
  webhookId: Scalars['ID'];
};

export type UpdateWorkspaceInput = {
  name: Scalars['String'];
  workspaceId: Scalars['ID'];
};

export type UpdateWorkspacePayload = {
  __typename?: 'UpdateWorkspacePayload';
  workspace: Workspace;
};

export type User = Node & {
  __typename?: 'User';
  email: Scalars['String'];
  id: Scalars['ID'];
  name: Scalars['String'];
};

export type VersionedItem = {
  __typename?: 'VersionedItem';
  parents?: Maybe<Array<Scalars['String']>>;
  refs: Array<Scalars['String']>;
  value: Item;
  version: Scalars['String'];
};

export type Webhook = {
  __typename?: 'Webhook';
  active: Scalars['Boolean'];
  createdAt: Scalars['DateTime'];
  id: Scalars['ID'];
  name: Scalars['String'];
  secret: Scalars['String'];
  trigger: WebhookTrigger;
  updatedAt: Scalars['DateTime'];
  url: Scalars['URL'];
};

export type WebhookPayload = {
  __typename?: 'WebhookPayload';
  webhook: Webhook;
};

export type WebhookTrigger = {
  __typename?: 'WebhookTrigger';
  onAssetDecompress?: Maybe<Scalars['Boolean']>;
  onAssetDelete?: Maybe<Scalars['Boolean']>;
  onAssetUpload?: Maybe<Scalars['Boolean']>;
  onItemCreate?: Maybe<Scalars['Boolean']>;
  onItemDelete?: Maybe<Scalars['Boolean']>;
  onItemPublish?: Maybe<Scalars['Boolean']>;
  onItemUnPublish?: Maybe<Scalars['Boolean']>;
  onItemUpdate?: Maybe<Scalars['Boolean']>;
};

export type WebhookTriggerInput = {
  onAssetDecompress?: InputMaybe<Scalars['Boolean']>;
  onAssetDelete?: InputMaybe<Scalars['Boolean']>;
  onAssetUpload?: InputMaybe<Scalars['Boolean']>;
  onItemCreate?: InputMaybe<Scalars['Boolean']>;
  onItemDelete?: InputMaybe<Scalars['Boolean']>;
  onItemPublish?: InputMaybe<Scalars['Boolean']>;
  onItemUnPublish?: InputMaybe<Scalars['Boolean']>;
  onItemUpdate?: InputMaybe<Scalars['Boolean']>;
};

export type Workspace = Node & {
  __typename?: 'Workspace';
  id: Scalars['ID'];
  members: Array<WorkspaceMember>;
  name: Scalars['String'];
  personal: Scalars['Boolean'];
};

export type WorkspaceIntegrationMember = {
  __typename?: 'WorkspaceIntegrationMember';
  active: Scalars['Boolean'];
  integration?: Maybe<Integration>;
  integrationId: Scalars['ID'];
  invitedBy?: Maybe<User>;
  invitedById: Scalars['ID'];
  role: Role;
};

export type WorkspaceMember = WorkspaceIntegrationMember | WorkspaceUserMember;

export type WorkspaceUserMember = {
  __typename?: 'WorkspaceUserMember';
  role: Role;
  user?: Maybe<User>;
  userId: Scalars['ID'];
};

export type AssetFragmentFragment = { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null };

export type AssetFileFragmentFragment = { __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string };

export type AssetFile5FragmentFragment = { __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string }> | null }> | null }> | null }> | null }> | null };

export type IntegrationFragmentFragment = { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null };

export type RequestFragmentFragment = { __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, items: Array<{ __typename?: 'RequestItem', itemId: string, version?: string | null, ref?: string | null, item?: { __typename?: 'VersionedItem', version: string, parents?: Array<string> | null, refs: Array<string>, value: { __typename?: 'Item', id: string, schemaId: string, modelId: string, model: { __typename?: 'Model', name: string }, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } } | null }>, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null, project?: { __typename?: 'Project', id: string, name: string, createdAt: Date, updatedAt: Date } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }> };

export type ThreadFragmentFragment = { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> };

export type WorkspaceFragmentFragment = { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> };

export type GetAssetsQueryVariables = Exact<{
  projectId: Scalars['ID'];
  keyword?: InputMaybe<Scalars['String']>;
  sort?: InputMaybe<AssetSort>;
  pagination?: InputMaybe<Pagination>;
}>;


export type GetAssetsQuery = { __typename?: 'Query', assets: { __typename?: 'AssetConnection', totalCount: number, edges: Array<{ __typename?: 'AssetEdge', cursor: string, node?: { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } | null }>, nodes: Array<{ __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } | null>, pageInfo: { __typename?: 'PageInfo', startCursor?: string | null, endCursor?: string | null, hasNextPage: boolean, hasPreviousPage: boolean } } };

export type GetAssetsItemsQueryVariables = Exact<{
  projectId: Scalars['ID'];
  keyword?: InputMaybe<Scalars['String']>;
  sort?: InputMaybe<AssetSort>;
  pagination?: InputMaybe<Pagination>;
}>;


export type GetAssetsItemsQuery = { __typename?: 'Query', assets: { __typename?: 'AssetConnection', edges: Array<{ __typename?: 'AssetEdge', cursor: string, node?: { __typename?: 'Asset', id: string, items?: Array<{ __typename?: 'AssetItem', itemId: string, modelId: string }> | null } | null }>, nodes: Array<{ __typename?: 'Asset', id: string, items?: Array<{ __typename?: 'AssetItem', itemId: string, modelId: string }> | null } | null> } };

export type GetAssetQueryVariables = Exact<{
  assetId: Scalars['ID'];
}>;


export type GetAssetQuery = { __typename?: 'Query', node?: { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } | { __typename?: 'Integration' } | { __typename?: 'Item' } | { __typename?: 'Model' } | { __typename?: 'Project' } | { __typename?: 'Request' } | { __typename?: 'Schema' } | { __typename?: 'User' } | { __typename?: 'Workspace' } | null };

export type GetAssetFileQueryVariables = Exact<{
  assetId: Scalars['ID'];
}>;


export type GetAssetFileQuery = { __typename?: 'Query', assetFile: { __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string, children?: Array<{ __typename?: 'AssetFile', name: string, size: number, contentType?: string | null, path: string }> | null }> | null }> | null }> | null }> | null } };

export type GetAssetItemQueryVariables = Exact<{
  assetId: Scalars['ID'];
}>;


export type GetAssetItemQuery = { __typename?: 'Query', node?: { __typename?: 'Asset', id: string, items?: Array<{ __typename?: 'AssetItem', itemId: string, modelId: string }> | null } | { __typename?: 'Integration' } | { __typename?: 'Item' } | { __typename?: 'Model' } | { __typename?: 'Project' } | { __typename?: 'Request' } | { __typename?: 'Schema' } | { __typename?: 'User' } | { __typename?: 'Workspace' } | null };

export type CreateAssetMutationVariables = Exact<{
  projectId: Scalars['ID'];
  file?: InputMaybe<Scalars['Upload']>;
  url?: InputMaybe<Scalars['String']>;
  skipDecompression?: InputMaybe<Scalars['Boolean']>;
}>;


export type CreateAssetMutation = { __typename?: 'Mutation', createAsset?: { __typename?: 'CreateAssetPayload', asset: { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } } | null };

export type UpdateAssetMutationVariables = Exact<{
  id: Scalars['ID'];
  previewType?: InputMaybe<PreviewType>;
}>;


export type UpdateAssetMutation = { __typename?: 'Mutation', updateAsset?: { __typename?: 'UpdateAssetPayload', asset: { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } } | null };

export type DeleteAssetMutationVariables = Exact<{
  assetId: Scalars['ID'];
}>;


export type DeleteAssetMutation = { __typename?: 'Mutation', deleteAsset?: { __typename?: 'DeleteAssetPayload', assetId: string } | null };

export type DecompressAssetMutationVariables = Exact<{
  assetId: Scalars['ID'];
}>;


export type DecompressAssetMutation = { __typename?: 'Mutation', decompressAsset?: { __typename?: 'DecompressAssetPayload', asset: { __typename?: 'Asset', id: string, projectId: string, createdAt: Date, size: number, previewType?: PreviewType | null, uuid: string, url: string, archiveExtractionStatus?: ArchiveExtractionStatus | null, createdBy: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | { __typename?: 'User', id: string, name: string, email: string }, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } } | null };

export type AddCommentMutationVariables = Exact<{
  threadId: Scalars['ID'];
  content: Scalars['String'];
}>;


export type AddCommentMutation = { __typename?: 'Mutation', addComment?: { __typename?: 'CommentPayload', comment: { __typename?: 'Comment', id: string, authorType: OperatorType, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null } } | null };

export type UpdateCommentMutationVariables = Exact<{
  commentId: Scalars['ID'];
  threadId: Scalars['ID'];
  content: Scalars['String'];
}>;


export type UpdateCommentMutation = { __typename?: 'Mutation', updateComment?: { __typename?: 'CommentPayload', comment: { __typename?: 'Comment', id: string, authorType: OperatorType, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null } } | null };

export type DeleteCommentMutationVariables = Exact<{
  commentId: Scalars['ID'];
  threadId: Scalars['ID'];
}>;


export type DeleteCommentMutation = { __typename?: 'Mutation', deleteComment?: { __typename?: 'DeleteCommentPayload', commentId: string } | null };

export type CreateFieldMutationVariables = Exact<{
  modelId: Scalars['ID'];
  type: SchemaFieldType;
  title: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  key: Scalars['String'];
  multiple: Scalars['Boolean'];
  unique: Scalars['Boolean'];
  required: Scalars['Boolean'];
  typeProperty: SchemaFieldTypePropertyInput;
}>;


export type CreateFieldMutation = { __typename?: 'Mutation', createField?: { __typename?: 'FieldPayload', field: { __typename?: 'SchemaField', id: string } } | null };

export type UpdateFieldMutationVariables = Exact<{
  modelId: Scalars['ID'];
  fieldId: Scalars['ID'];
  title: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  order?: InputMaybe<Scalars['Int']>;
  key: Scalars['String'];
  multiple: Scalars['Boolean'];
  unique: Scalars['Boolean'];
  required: Scalars['Boolean'];
  typeProperty: SchemaFieldTypePropertyInput;
}>;


export type UpdateFieldMutation = { __typename?: 'Mutation', updateField?: { __typename?: 'FieldPayload', field: { __typename?: 'SchemaField', id: string } } | null };

export type UpdateFieldsMutationVariables = Exact<{
  updateFieldInput: Array<UpdateFieldInput> | UpdateFieldInput;
}>;


export type UpdateFieldsMutation = { __typename?: 'Mutation', updateFields?: { __typename?: 'FieldsPayload', fields: Array<{ __typename?: 'SchemaField', id: string }> } | null };

export type DeleteFieldMutationVariables = Exact<{
  modelId: Scalars['ID'];
  fieldId: Scalars['ID'];
}>;


export type DeleteFieldMutation = { __typename?: 'Mutation', deleteField?: { __typename?: 'DeleteFieldPayload', fieldId: string } | null };

export type CreateIntegrationMutationVariables = Exact<{
  name: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  logoUrl: Scalars['URL'];
  type: IntegrationType;
}>;


export type CreateIntegrationMutation = { __typename?: 'Mutation', createIntegration?: { __typename?: 'IntegrationPayload', integration: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType } } | null };

export type UpdateIntegrationMutationVariables = Exact<{
  integrationId: Scalars['ID'];
  name: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  logoUrl: Scalars['URL'];
}>;


export type UpdateIntegrationMutation = { __typename?: 'Mutation', updateIntegration?: { __typename?: 'IntegrationPayload', integration: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType } } | null };

export type DeleteIntegrationMutationVariables = Exact<{
  integrationId: Scalars['ID'];
}>;


export type DeleteIntegrationMutation = { __typename?: 'Mutation', deleteIntegration?: { __typename?: 'DeleteIntegrationPayload', integrationId: string } | null };

export type GetItemsQueryVariables = Exact<{
  schemaId: Scalars['ID'];
  pagination?: InputMaybe<Pagination>;
}>;


export type GetItemsQuery = { __typename?: 'Query', items: { __typename?: 'ItemConnection', nodes: Array<{ __typename?: 'Item', id: string, schemaId: string, createdAt: Date, updatedAt: Date, status: ItemStatus, user?: { __typename?: 'User', name: string } | null, integration?: { __typename?: 'Integration', name: string } | null, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, thread: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null> } };

export type GetItemQueryVariables = Exact<{
  id: Scalars['ID'];
}>;


export type GetItemQuery = { __typename?: 'Query', node?: { __typename?: 'Asset' } | { __typename?: 'Integration' } | { __typename?: 'Item', id: string, schemaId: string, createdAt: Date, updatedAt: Date, status: ItemStatus, assets: Array<{ __typename?: 'Asset', id: string, url: string } | null>, user?: { __typename?: 'User', name: string } | null, integration?: { __typename?: 'Integration', name: string } | null, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, thread: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } } | { __typename?: 'Model' } | { __typename?: 'Project' } | { __typename?: 'Request' } | { __typename?: 'Schema' } | { __typename?: 'User' } | { __typename?: 'Workspace' } | null };

export type SearchItemQueryVariables = Exact<{
  query: ItemQuery;
  sort?: InputMaybe<ItemSort>;
  pagination?: InputMaybe<Pagination>;
}>;


export type SearchItemQuery = { __typename?: 'Query', searchItem: { __typename?: 'ItemConnection', totalCount: number, nodes: Array<{ __typename?: 'Item', id: string, schemaId: string, createdAt: Date, updatedAt: Date, status: ItemStatus, assets: Array<{ __typename?: 'Asset', id: string, url: string } | null>, user?: { __typename?: 'User', name: string } | null, integration?: { __typename?: 'Integration', name: string } | null, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, thread: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null> } };

export type CreateItemMutationVariables = Exact<{
  modelId: Scalars['ID'];
  schemaId: Scalars['ID'];
  fields: Array<ItemFieldInput> | ItemFieldInput;
}>;


export type CreateItemMutation = { __typename?: 'Mutation', createItem?: { __typename?: 'ItemPayload', item: { __typename?: 'Item', id: string, schemaId: string, fields: Array<{ __typename?: 'ItemField', value?: any | null, type: SchemaFieldType, schemaFieldId: string }> } } | null };

export type DeleteItemMutationVariables = Exact<{
  itemId: Scalars['ID'];
}>;


export type DeleteItemMutation = { __typename?: 'Mutation', deleteItem?: { __typename?: 'DeleteItemPayload', itemId: string } | null };

export type UpdateItemMutationVariables = Exact<{
  itemId: Scalars['ID'];
  fields: Array<ItemFieldInput> | ItemFieldInput;
}>;


export type UpdateItemMutation = { __typename?: 'Mutation', updateItem?: { __typename?: 'ItemPayload', item: { __typename?: 'Item', id: string, schemaId: string, fields: Array<{ __typename?: 'ItemField', value?: any | null, type: SchemaFieldType, schemaFieldId: string }> } } | null };

export type UnpublishItemMutationVariables = Exact<{
  itemId: Array<Scalars['ID']> | Scalars['ID'];
}>;


export type UnpublishItemMutation = { __typename?: 'Mutation', unpublishItem?: { __typename?: 'UnpublishItemPayload', items: Array<{ __typename?: 'Item', id: string }> } | null };

export type GetModelsQueryVariables = Exact<{
  projectId: Scalars['ID'];
  pagination?: InputMaybe<Pagination>;
}>;


export type GetModelsQuery = { __typename?: 'Query', models: { __typename?: 'ModelConnection', nodes: Array<{ __typename?: 'Model', id: string, name: string, description: string, key: string, public: boolean, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, order?: number | null, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } | null> } };

export type CreateModelMutationVariables = Exact<{
  projectId: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  key?: InputMaybe<Scalars['String']>;
}>;


export type CreateModelMutation = { __typename?: 'Mutation', createModel?: { __typename?: 'ModelPayload', model: { __typename?: 'Model', id: string, name: string } } | null };

export type DeleteModelMutationVariables = Exact<{
  modelId: Scalars['ID'];
}>;


export type DeleteModelMutation = { __typename?: 'Mutation', deleteModel?: { __typename?: 'DeleteModelPayload', modelId: string } | null };

export type UpdateModelMutationVariables = Exact<{
  modelId: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  key?: InputMaybe<Scalars['String']>;
  public: Scalars['Boolean'];
}>;


export type UpdateModelMutation = { __typename?: 'Mutation', updateModel?: { __typename?: 'ModelPayload', model: { __typename?: 'Model', id: string, name: string } } | null };

export type CheckModelKeyAvailabilityQueryVariables = Exact<{
  projectId: Scalars['ID'];
  key: Scalars['String'];
}>;


export type CheckModelKeyAvailabilityQuery = { __typename?: 'Query', checkModelKeyAvailability: { __typename?: 'KeyAvailability', key: string, available: boolean } };

export type GetProjectQueryVariables = Exact<{
  projectId: Scalars['ID'];
}>;


export type GetProjectQuery = { __typename?: 'Query', node?: { __typename?: 'Asset', id: string } | { __typename?: 'Integration', id: string } | { __typename?: 'Item', id: string } | { __typename?: 'Model', id: string } | { __typename?: 'Project', name: string, description: string, alias: string, id: string, publication?: { __typename?: 'ProjectPublication', scope: ProjectPublicationScope, assetPublic: boolean } | null } | { __typename?: 'Request', id: string } | { __typename?: 'Schema', id: string } | { __typename?: 'User', id: string } | { __typename?: 'Workspace', id: string } | null };

export type GetProjectsQueryVariables = Exact<{
  workspaceId: Scalars['ID'];
  pagination?: InputMaybe<Pagination>;
}>;


export type GetProjectsQuery = { __typename?: 'Query', projects: { __typename?: 'ProjectConnection', nodes: Array<{ __typename?: 'Project', id: string, name: string, description: string, alias: string, publication?: { __typename?: 'ProjectPublication', scope: ProjectPublicationScope, assetPublic: boolean } | null } | null> } };

export type CheckProjectAliasQueryVariables = Exact<{
  alias: Scalars['String'];
}>;


export type CheckProjectAliasQuery = { __typename?: 'Query', checkProjectAlias: { __typename?: 'ProjectAliasAvailability', alias: string, available: boolean } };

export type CreateProjectMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  name: Scalars['String'];
  description: Scalars['String'];
  alias: Scalars['String'];
}>;


export type CreateProjectMutation = { __typename?: 'Mutation', createProject?: { __typename?: 'ProjectPayload', project: { __typename?: 'Project', id: string, name: string, description: string, alias: string, publication?: { __typename?: 'ProjectPublication', scope: ProjectPublicationScope, assetPublic: boolean } | null } } | null };

export type DeleteProjectMutationVariables = Exact<{
  projectId: Scalars['ID'];
}>;


export type DeleteProjectMutation = { __typename?: 'Mutation', deleteProject?: { __typename?: 'DeleteProjectPayload', projectId: string } | null };

export type UpdateProjectMutationVariables = Exact<{
  projectId: Scalars['ID'];
  name?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  alias?: InputMaybe<Scalars['String']>;
  publication?: InputMaybe<UpdateProjectPublicationInput>;
}>;


export type UpdateProjectMutation = { __typename?: 'Mutation', updateProject?: { __typename?: 'ProjectPayload', project: { __typename?: 'Project', id: string, name: string, description: string, alias: string, publication?: { __typename?: 'ProjectPublication', scope: ProjectPublicationScope, assetPublic: boolean } | null } } | null };

export type GetRequestsQueryVariables = Exact<{
  projectId: Scalars['ID'];
  key?: InputMaybe<Scalars['String']>;
  state?: InputMaybe<Array<RequestState> | RequestState>;
  pagination?: InputMaybe<Pagination>;
  createdBy?: InputMaybe<Scalars['ID']>;
  reviewer?: InputMaybe<Scalars['ID']>;
  sort?: InputMaybe<Sort>;
}>;


export type GetRequestsQuery = { __typename?: 'Query', requests: { __typename?: 'RequestConnection', totalCount: number, nodes: Array<{ __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }>, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null } | null> } };

export type GetModalRequestsQueryVariables = Exact<{
  projectId: Scalars['ID'];
  key?: InputMaybe<Scalars['String']>;
  state?: InputMaybe<Array<RequestState> | RequestState>;
  pagination?: InputMaybe<Pagination>;
  createdBy?: InputMaybe<Scalars['ID']>;
  reviewer?: InputMaybe<Scalars['ID']>;
  sort?: InputMaybe<Sort>;
}>;


export type GetModalRequestsQuery = { __typename?: 'Query', requests: { __typename?: 'RequestConnection', totalCount: number, nodes: Array<{ __typename?: 'Request', id: string, title: string, description?: string | null, state: RequestState, createdAt: Date, createdBy?: { __typename?: 'User', name: string } | null, items: Array<{ __typename?: 'RequestItem', itemId: string }>, reviewers: Array<{ __typename?: 'User', id: string, name: string }> } | null> } };

export type GetRequestQueryVariables = Exact<{
  requestId: Scalars['ID'];
}>;


export type GetRequestQuery = { __typename?: 'Query', node?: { __typename?: 'Asset', id: string } | { __typename?: 'Integration', id: string } | { __typename?: 'Item', id: string } | { __typename?: 'Model', id: string } | { __typename?: 'Project', id: string } | { __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, items: Array<{ __typename?: 'RequestItem', itemId: string, version?: string | null, ref?: string | null, item?: { __typename?: 'VersionedItem', version: string, parents?: Array<string> | null, refs: Array<string>, value: { __typename?: 'Item', id: string, schemaId: string, modelId: string, model: { __typename?: 'Model', name: string }, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } } | null }>, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null, project?: { __typename?: 'Project', id: string, name: string, createdAt: Date, updatedAt: Date } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }> } | { __typename?: 'Schema', id: string } | { __typename?: 'User', id: string } | { __typename?: 'Workspace', id: string } | null };

export type CreateRequestMutationVariables = Exact<{
  projectId: Scalars['ID'];
  title: Scalars['String'];
  description?: InputMaybe<Scalars['String']>;
  state?: InputMaybe<RequestState>;
  reviewersId?: InputMaybe<Array<Scalars['ID']> | Scalars['ID']>;
  items: Array<RequestItemInput> | RequestItemInput;
}>;


export type CreateRequestMutation = { __typename?: 'Mutation', createRequest?: { __typename?: 'RequestPayload', request: { __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, items: Array<{ __typename?: 'RequestItem', itemId: string, version?: string | null, ref?: string | null, item?: { __typename?: 'VersionedItem', version: string, parents?: Array<string> | null, refs: Array<string>, value: { __typename?: 'Item', id: string, schemaId: string, modelId: string, model: { __typename?: 'Model', name: string }, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } } | null }>, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null, project?: { __typename?: 'Project', id: string, name: string, createdAt: Date, updatedAt: Date } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }> } } | null };

export type UpdateRequestMutationVariables = Exact<{
  requestId: Scalars['ID'];
  title?: InputMaybe<Scalars['String']>;
  description?: InputMaybe<Scalars['String']>;
  state?: InputMaybe<RequestState>;
  reviewersId?: InputMaybe<Array<Scalars['ID']> | Scalars['ID']>;
  items?: InputMaybe<Array<RequestItemInput> | RequestItemInput>;
}>;


export type UpdateRequestMutation = { __typename?: 'Mutation', updateRequest?: { __typename?: 'RequestPayload', request: { __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, items: Array<{ __typename?: 'RequestItem', itemId: string, version?: string | null, ref?: string | null, item?: { __typename?: 'VersionedItem', version: string, parents?: Array<string> | null, refs: Array<string>, value: { __typename?: 'Item', id: string, schemaId: string, modelId: string, model: { __typename?: 'Model', name: string }, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } } | null }>, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null, project?: { __typename?: 'Project', id: string, name: string, createdAt: Date, updatedAt: Date } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }> } } | null };

export type ApproveRequestMutationVariables = Exact<{
  requestId: Scalars['ID'];
}>;


export type ApproveRequestMutation = { __typename?: 'Mutation', approveRequest?: { __typename?: 'RequestPayload', request: { __typename?: 'Request', id: string, title: string, description?: string | null, workspaceId: string, projectId: string, threadId: string, reviewersId: Array<string>, state: RequestState, createdAt: Date, updatedAt: Date, approvedAt?: Date | null, closedAt?: Date | null, items: Array<{ __typename?: 'RequestItem', itemId: string, version?: string | null, ref?: string | null, item?: { __typename?: 'VersionedItem', version: string, parents?: Array<string> | null, refs: Array<string>, value: { __typename?: 'Item', id: string, schemaId: string, modelId: string, model: { __typename?: 'Model', name: string }, fields: Array<{ __typename?: 'ItemField', schemaFieldId: string, type: SchemaFieldType, value?: any | null }>, schema: { __typename?: 'Schema', id: string, fields: Array<{ __typename?: 'SchemaField', id: string, type: SchemaFieldType, title: string, key: string, description?: string | null, required: boolean, unique: boolean, multiple: boolean, typeProperty?: { __typename?: 'SchemaFieldAsset', assetDefaultValue?: any | null } | { __typename?: 'SchemaFieldBool', defaultValue?: any | null } | { __typename?: 'SchemaFieldDate' } | { __typename?: 'SchemaFieldInteger', min?: number | null, max?: number | null, integerDefaultValue?: any | null } | { __typename?: 'SchemaFieldMarkdown', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldReference' } | { __typename?: 'SchemaFieldRichText' } | { __typename?: 'SchemaFieldSelect', values: Array<string>, selectDefaultValue?: any | null } | { __typename?: 'SchemaFieldTag' } | { __typename?: 'SchemaFieldText', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldTextArea', defaultValue?: any | null, maxLength?: number | null } | { __typename?: 'SchemaFieldURL', defaultValue?: any | null } | null }> } } } | null }>, createdBy?: { __typename?: 'User', id: string, name: string, email: string } | null, thread?: { __typename?: 'Thread', id: string, workspaceId: string, comments: Array<{ __typename?: 'Comment', id: string, authorId: string, content: string, createdAt: Date, author?: { __typename?: 'Integration', id: string, name: string } | { __typename?: 'User', id: string, name: string, email: string } | null }> } | null, project?: { __typename?: 'Project', id: string, name: string, createdAt: Date, updatedAt: Date } | null, reviewers: Array<{ __typename?: 'User', id: string, name: string, email: string }> } } | null };

export type DeleteRequestMutationVariables = Exact<{
  projectId: Scalars['ID'];
  requestsId: Array<Scalars['ID']> | Scalars['ID'];
}>;


export type DeleteRequestMutation = { __typename?: 'Mutation', deleteRequest?: { __typename?: 'DeleteRequestPayload', requests: Array<string> } | null };

export type GetUserBySearchQueryVariables = Exact<{
  nameOrEmail: Scalars['String'];
}>;


export type GetUserBySearchQuery = { __typename?: 'Query', searchUser?: { __typename?: 'User', id: string, name: string, email: string } | null };

export type GetMeQueryVariables = Exact<{ [key: string]: never; }>;


export type GetMeQuery = { __typename?: 'Query', me?: { __typename?: 'Me', id: string, name: string, email: string, lang: string, auths: Array<string>, myWorkspace: { __typename?: 'Workspace', id: string, name: string }, workspaces: Array<{ __typename?: 'Workspace', id: string, name: string, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> }>, integrations: Array<{ __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null }> } | null };

export type GetProfileQueryVariables = Exact<{ [key: string]: never; }>;


export type GetProfileQuery = { __typename?: 'Query', me?: { __typename?: 'Me', id: string, name: string, email: string, lang: string, theme: Theme, auths: Array<string>, myWorkspace: { __typename?: 'Workspace', id: string, name: string } } | null };

export type GetLanguageQueryVariables = Exact<{ [key: string]: never; }>;


export type GetLanguageQuery = { __typename?: 'Query', me?: { __typename?: 'Me', id: string, lang: string } | null };

export type GetThemeQueryVariables = Exact<{ [key: string]: never; }>;


export type GetThemeQuery = { __typename?: 'Query', me?: { __typename?: 'Me', id: string, theme: Theme } | null };

export type UpdateMeMutationVariables = Exact<{
  name?: InputMaybe<Scalars['String']>;
  email?: InputMaybe<Scalars['String']>;
  lang?: InputMaybe<Scalars['Lang']>;
  theme?: InputMaybe<Theme>;
  password?: InputMaybe<Scalars['String']>;
  passwordConfirmation?: InputMaybe<Scalars['String']>;
}>;


export type UpdateMeMutation = { __typename?: 'Mutation', updateMe?: { __typename?: 'UpdateMePayload', me: { __typename?: 'Me', id: string, name: string, email: string, lang: string, theme: Theme, myWorkspace: { __typename?: 'Workspace', id: string, name: string } } } | null };

export type DeleteMeMutationVariables = Exact<{
  userId: Scalars['ID'];
}>;


export type DeleteMeMutation = { __typename?: 'Mutation', deleteMe?: { __typename?: 'DeleteMePayload', userId: string } | null };

export type CreateWebhookMutationVariables = Exact<{
  integrationId: Scalars['ID'];
  name: Scalars['String'];
  url: Scalars['URL'];
  active: Scalars['Boolean'];
  trigger: WebhookTriggerInput;
  secret: Scalars['String'];
}>;


export type CreateWebhookMutation = { __typename?: 'Mutation', createWebhook?: { __typename?: 'WebhookPayload', webhook: { __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } } } | null };

export type UpdateWebhookMutationVariables = Exact<{
  integrationId: Scalars['ID'];
  webhookId: Scalars['ID'];
  name: Scalars['String'];
  url: Scalars['URL'];
  active: Scalars['Boolean'];
  trigger: WebhookTriggerInput;
  secret?: InputMaybe<Scalars['String']>;
}>;


export type UpdateWebhookMutation = { __typename?: 'Mutation', updateWebhook?: { __typename?: 'WebhookPayload', webhook: { __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } } } | null };

export type DeleteWebhookMutationVariables = Exact<{
  integrationId: Scalars['ID'];
  webhookId: Scalars['ID'];
}>;


export type DeleteWebhookMutation = { __typename?: 'Mutation', deleteWebhook?: { __typename?: 'DeleteWebhookPayload', webhookId: string } | null };

export type GetWorkspacesQueryVariables = Exact<{ [key: string]: never; }>;


export type GetWorkspacesQuery = { __typename?: 'Query', me?: { __typename?: 'Me', id: string, name: string, myWorkspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> }, workspaces: Array<{ __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> }> } | null };

export type UpdateWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  name: Scalars['String'];
}>;


export type UpdateWorkspaceMutation = { __typename?: 'Mutation', updateWorkspace?: { __typename?: 'UpdateWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type DeleteWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
}>;


export type DeleteWorkspaceMutation = { __typename?: 'Mutation', deleteWorkspace?: { __typename?: 'DeleteWorkspacePayload', workspaceId: string } | null };

export type AddUsersToWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  users: Array<MemberInput> | MemberInput;
}>;


export type AddUsersToWorkspaceMutation = { __typename?: 'Mutation', addUsersToWorkspace?: { __typename?: 'AddUsersToWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type UpdateMemberOfWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  userId: Scalars['ID'];
  role: Role;
}>;


export type UpdateMemberOfWorkspaceMutation = { __typename?: 'Mutation', updateUserOfWorkspace?: { __typename?: 'UpdateMemberOfWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type RemoveMemberFromWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  userId: Scalars['ID'];
}>;


export type RemoveMemberFromWorkspaceMutation = { __typename?: 'Mutation', removeUserFromWorkspace?: { __typename?: 'RemoveMemberFromWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type AddIntegrationToWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  integrationId: Scalars['ID'];
  role: Role;
}>;


export type AddIntegrationToWorkspaceMutation = { __typename?: 'Mutation', addIntegrationToWorkspace?: { __typename?: 'AddUsersToWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type UpdateIntegrationOfWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  integrationId: Scalars['ID'];
  role: Role;
}>;


export type UpdateIntegrationOfWorkspaceMutation = { __typename?: 'Mutation', updateIntegrationOfWorkspace?: { __typename?: 'UpdateMemberOfWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type RemoveIntegrationFromWorkspaceMutationVariables = Exact<{
  workspaceId: Scalars['ID'];
  integrationId: Scalars['ID'];
}>;


export type RemoveIntegrationFromWorkspaceMutation = { __typename?: 'Mutation', removeIntegrationFromWorkspace?: { __typename?: 'RemoveMemberFromWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export type CreateWorkspaceMutationVariables = Exact<{
  name: Scalars['String'];
}>;


export type CreateWorkspaceMutation = { __typename?: 'Mutation', createWorkspace?: { __typename?: 'CreateWorkspacePayload', workspace: { __typename?: 'Workspace', id: string, name: string, personal: boolean, members: Array<{ __typename?: 'WorkspaceIntegrationMember', active: boolean, invitedById: string, integrationRole: Role, integration?: { __typename?: 'Integration', id: string, name: string, description?: string | null, logoUrl: string, iType: IntegrationType, developerId: string, createdAt: Date, updatedAt: Date, developer: { __typename?: 'User', id: string, name: string, email: string }, config?: { __typename?: 'IntegrationConfig', token: string, webhooks: Array<{ __typename?: 'Webhook', id: string, name: string, url: string, active: boolean, secret: string, createdAt: Date, updatedAt: Date, trigger: { __typename?: 'WebhookTrigger', onItemCreate?: boolean | null, onItemUpdate?: boolean | null, onItemDelete?: boolean | null, onItemPublish?: boolean | null, onItemUnPublish?: boolean | null, onAssetUpload?: boolean | null, onAssetDecompress?: boolean | null, onAssetDelete?: boolean | null } }> } | null } | null, invitedBy?: { __typename?: 'User', id: string, name: string, email: string } | null } | { __typename?: 'WorkspaceUserMember', userId: string, role: Role, user?: { __typename?: 'User', id: string, name: string, email: string } | null }> } } | null };

export const IntegrationFragmentFragmentDoc = gql`
    fragment integrationFragment on Integration {
  id
  name
  description
  logoUrl
  iType
  developerId
  developer {
    id
    name
    email
  }
  config {
    token
    webhooks {
      id
      name
      url
      active
      secret
      trigger {
        onItemCreate
        onItemUpdate
        onItemDelete
        onItemPublish
        onItemUnPublish
        onAssetUpload
        onAssetDecompress
        onAssetDelete
      }
      createdAt
      updatedAt
    }
  }
  createdAt
  updatedAt
}
    `;
export const ThreadFragmentFragmentDoc = gql`
    fragment threadFragment on Thread {
  id
  workspaceId
  comments {
    id
    author {
      ... on User {
        id
        name
        email
      }
      ... on Integration {
        id
        name
      }
    }
    authorId
    content
    createdAt
  }
}
    `;
export const AssetFragmentFragmentDoc = gql`
    fragment assetFragment on Asset {
  id
  projectId
  createdAt
  createdBy {
    ... on User {
      id
      name
      email
    }
    ... on Integration {
      ...integrationFragment
    }
  }
  size
  previewType
  uuid
  url
  thread {
    ...threadFragment
  }
  archiveExtractionStatus
}
    ${IntegrationFragmentFragmentDoc}
${ThreadFragmentFragmentDoc}`;
export const AssetFileFragmentFragmentDoc = gql`
    fragment assetFileFragment on AssetFile {
  name
  size
  contentType
  path
}
    `;
export const AssetFile5FragmentFragmentDoc = gql`
    fragment assetFile5Fragment on AssetFile {
  ...assetFileFragment
  children {
    ...assetFileFragment
    children {
      ...assetFileFragment
      children {
        ...assetFileFragment
        children {
          ...assetFileFragment
          children {
            ...assetFileFragment
          }
        }
      }
    }
  }
}
    ${AssetFileFragmentFragmentDoc}`;
export const RequestFragmentFragmentDoc = gql`
    fragment requestFragment on Request {
  id
  items {
    itemId
    version
    ref
    item {
      version
      parents
      refs
      value {
        id
        schemaId
        modelId
        model {
          name
        }
        fields {
          schemaFieldId
          type
          value
        }
        schema {
          id
          fields {
            id
            type
            title
            key
            description
            required
            unique
            multiple
            typeProperty {
              ... on SchemaFieldText {
                defaultValue
                maxLength
              }
              ... on SchemaFieldTextArea {
                defaultValue
                maxLength
              }
              ... on SchemaFieldMarkdown {
                defaultValue
                maxLength
              }
              ... on SchemaFieldAsset {
                assetDefaultValue: defaultValue
              }
              ... on SchemaFieldSelect {
                selectDefaultValue: defaultValue
                values
              }
              ... on SchemaFieldInteger {
                integerDefaultValue: defaultValue
                min
                max
              }
              ... on SchemaFieldBool {
                defaultValue
              }
              ... on SchemaFieldURL {
                defaultValue
              }
            }
          }
        }
      }
    }
  }
  title
  description
  createdBy {
    id
    name
    email
  }
  workspaceId
  projectId
  threadId
  reviewersId
  state
  createdAt
  updatedAt
  approvedAt
  closedAt
  thread {
    ...threadFragment
  }
  project {
    id
    name
    createdAt
    updatedAt
  }
  reviewers {
    id
    name
    email
  }
}
    ${ThreadFragmentFragmentDoc}`;
export const WorkspaceFragmentFragmentDoc = gql`
    fragment WorkspaceFragment on Workspace {
  id
  name
  members {
    ... on WorkspaceUserMember {
      user {
        id
        name
        email
      }
      userId
      role
    }
    ... on WorkspaceIntegrationMember {
      integration {
        ...integrationFragment
      }
      integrationRole: role
      active
      invitedBy {
        id
        name
        email
      }
      invitedById
    }
  }
  personal
}
    ${IntegrationFragmentFragmentDoc}`;
export const GetAssetsDocument = gql`
    query GetAssets($projectId: ID!, $keyword: String, $sort: AssetSort, $pagination: Pagination) {
  assets(
    projectId: $projectId
    keyword: $keyword
    sort: $sort
    pagination: $pagination
  ) {
    edges {
      cursor
      node {
        ...assetFragment
      }
    }
    nodes {
      ...assetFragment
    }
    pageInfo {
      startCursor
      endCursor
      hasNextPage
      hasPreviousPage
    }
    totalCount
  }
}
    ${AssetFragmentFragmentDoc}`;

/**
 * __useGetAssetsQuery__
 *
 * To run a query within a React component, call `useGetAssetsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      keyword: // value for 'keyword'
 *      sort: // value for 'sort'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useGetAssetsQuery(baseOptions: Apollo.QueryHookOptions<GetAssetsQuery, GetAssetsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetsQuery, GetAssetsQueryVariables>(GetAssetsDocument, options);
      }
export function useGetAssetsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetsQuery, GetAssetsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetsQuery, GetAssetsQueryVariables>(GetAssetsDocument, options);
        }
export type GetAssetsQueryHookResult = ReturnType<typeof useGetAssetsQuery>;
export type GetAssetsLazyQueryHookResult = ReturnType<typeof useGetAssetsLazyQuery>;
export type GetAssetsQueryResult = Apollo.QueryResult<GetAssetsQuery, GetAssetsQueryVariables>;
export const GetAssetsItemsDocument = gql`
    query GetAssetsItems($projectId: ID!, $keyword: String, $sort: AssetSort, $pagination: Pagination) {
  assets(
    projectId: $projectId
    keyword: $keyword
    sort: $sort
    pagination: $pagination
  ) {
    edges {
      cursor
      node {
        id
        items {
          itemId
          modelId
        }
      }
    }
    nodes {
      id
      items {
        itemId
        modelId
      }
    }
  }
}
    `;

/**
 * __useGetAssetsItemsQuery__
 *
 * To run a query within a React component, call `useGetAssetsItemsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetsItemsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetsItemsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      keyword: // value for 'keyword'
 *      sort: // value for 'sort'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useGetAssetsItemsQuery(baseOptions: Apollo.QueryHookOptions<GetAssetsItemsQuery, GetAssetsItemsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetsItemsQuery, GetAssetsItemsQueryVariables>(GetAssetsItemsDocument, options);
      }
export function useGetAssetsItemsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetsItemsQuery, GetAssetsItemsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetsItemsQuery, GetAssetsItemsQueryVariables>(GetAssetsItemsDocument, options);
        }
export type GetAssetsItemsQueryHookResult = ReturnType<typeof useGetAssetsItemsQuery>;
export type GetAssetsItemsLazyQueryHookResult = ReturnType<typeof useGetAssetsItemsLazyQuery>;
export type GetAssetsItemsQueryResult = Apollo.QueryResult<GetAssetsItemsQuery, GetAssetsItemsQueryVariables>;
export const GetAssetDocument = gql`
    query GetAsset($assetId: ID!) {
  node(id: $assetId, type: ASSET) {
    ...assetFragment
  }
}
    ${AssetFragmentFragmentDoc}`;

/**
 * __useGetAssetQuery__
 *
 * To run a query within a React component, call `useGetAssetQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetQuery({
 *   variables: {
 *      assetId: // value for 'assetId'
 *   },
 * });
 */
export function useGetAssetQuery(baseOptions: Apollo.QueryHookOptions<GetAssetQuery, GetAssetQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetQuery, GetAssetQueryVariables>(GetAssetDocument, options);
      }
export function useGetAssetLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetQuery, GetAssetQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetQuery, GetAssetQueryVariables>(GetAssetDocument, options);
        }
export type GetAssetQueryHookResult = ReturnType<typeof useGetAssetQuery>;
export type GetAssetLazyQueryHookResult = ReturnType<typeof useGetAssetLazyQuery>;
export type GetAssetQueryResult = Apollo.QueryResult<GetAssetQuery, GetAssetQueryVariables>;
export const GetAssetFileDocument = gql`
    query GetAssetFile($assetId: ID!) {
  assetFile(assetId: $assetId) {
    ...assetFile5Fragment
  }
}
    ${AssetFile5FragmentFragmentDoc}`;

/**
 * __useGetAssetFileQuery__
 *
 * To run a query within a React component, call `useGetAssetFileQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetFileQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetFileQuery({
 *   variables: {
 *      assetId: // value for 'assetId'
 *   },
 * });
 */
export function useGetAssetFileQuery(baseOptions: Apollo.QueryHookOptions<GetAssetFileQuery, GetAssetFileQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetFileQuery, GetAssetFileQueryVariables>(GetAssetFileDocument, options);
      }
export function useGetAssetFileLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetFileQuery, GetAssetFileQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetFileQuery, GetAssetFileQueryVariables>(GetAssetFileDocument, options);
        }
export type GetAssetFileQueryHookResult = ReturnType<typeof useGetAssetFileQuery>;
export type GetAssetFileLazyQueryHookResult = ReturnType<typeof useGetAssetFileLazyQuery>;
export type GetAssetFileQueryResult = Apollo.QueryResult<GetAssetFileQuery, GetAssetFileQueryVariables>;
export const GetAssetItemDocument = gql`
    query GetAssetItem($assetId: ID!) {
  node(id: $assetId, type: ASSET) {
    ... on Asset {
      id
      items {
        itemId
        modelId
      }
    }
  }
}
    `;

/**
 * __useGetAssetItemQuery__
 *
 * To run a query within a React component, call `useGetAssetItemQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetAssetItemQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetAssetItemQuery({
 *   variables: {
 *      assetId: // value for 'assetId'
 *   },
 * });
 */
export function useGetAssetItemQuery(baseOptions: Apollo.QueryHookOptions<GetAssetItemQuery, GetAssetItemQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetAssetItemQuery, GetAssetItemQueryVariables>(GetAssetItemDocument, options);
      }
export function useGetAssetItemLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetAssetItemQuery, GetAssetItemQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetAssetItemQuery, GetAssetItemQueryVariables>(GetAssetItemDocument, options);
        }
export type GetAssetItemQueryHookResult = ReturnType<typeof useGetAssetItemQuery>;
export type GetAssetItemLazyQueryHookResult = ReturnType<typeof useGetAssetItemLazyQuery>;
export type GetAssetItemQueryResult = Apollo.QueryResult<GetAssetItemQuery, GetAssetItemQueryVariables>;
export const CreateAssetDocument = gql`
    mutation CreateAsset($projectId: ID!, $file: Upload, $url: String, $skipDecompression: Boolean) {
  createAsset(
    input: {projectId: $projectId, file: $file, url: $url, skipDecompression: $skipDecompression}
  ) {
    asset {
      ...assetFragment
    }
  }
}
    ${AssetFragmentFragmentDoc}`;
export type CreateAssetMutationFn = Apollo.MutationFunction<CreateAssetMutation, CreateAssetMutationVariables>;

/**
 * __useCreateAssetMutation__
 *
 * To run a mutation, you first call `useCreateAssetMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateAssetMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createAssetMutation, { data, loading, error }] = useCreateAssetMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      file: // value for 'file'
 *      url: // value for 'url'
 *      skipDecompression: // value for 'skipDecompression'
 *   },
 * });
 */
export function useCreateAssetMutation(baseOptions?: Apollo.MutationHookOptions<CreateAssetMutation, CreateAssetMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateAssetMutation, CreateAssetMutationVariables>(CreateAssetDocument, options);
      }
export type CreateAssetMutationHookResult = ReturnType<typeof useCreateAssetMutation>;
export type CreateAssetMutationResult = Apollo.MutationResult<CreateAssetMutation>;
export type CreateAssetMutationOptions = Apollo.BaseMutationOptions<CreateAssetMutation, CreateAssetMutationVariables>;
export const UpdateAssetDocument = gql`
    mutation UpdateAsset($id: ID!, $previewType: PreviewType) {
  updateAsset(input: {id: $id, previewType: $previewType}) {
    asset {
      ...assetFragment
    }
  }
}
    ${AssetFragmentFragmentDoc}`;
export type UpdateAssetMutationFn = Apollo.MutationFunction<UpdateAssetMutation, UpdateAssetMutationVariables>;

/**
 * __useUpdateAssetMutation__
 *
 * To run a mutation, you first call `useUpdateAssetMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateAssetMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateAssetMutation, { data, loading, error }] = useUpdateAssetMutation({
 *   variables: {
 *      id: // value for 'id'
 *      previewType: // value for 'previewType'
 *   },
 * });
 */
export function useUpdateAssetMutation(baseOptions?: Apollo.MutationHookOptions<UpdateAssetMutation, UpdateAssetMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateAssetMutation, UpdateAssetMutationVariables>(UpdateAssetDocument, options);
      }
export type UpdateAssetMutationHookResult = ReturnType<typeof useUpdateAssetMutation>;
export type UpdateAssetMutationResult = Apollo.MutationResult<UpdateAssetMutation>;
export type UpdateAssetMutationOptions = Apollo.BaseMutationOptions<UpdateAssetMutation, UpdateAssetMutationVariables>;
export const DeleteAssetDocument = gql`
    mutation DeleteAsset($assetId: ID!) {
  deleteAsset(input: {assetId: $assetId}) {
    assetId
  }
}
    `;
export type DeleteAssetMutationFn = Apollo.MutationFunction<DeleteAssetMutation, DeleteAssetMutationVariables>;

/**
 * __useDeleteAssetMutation__
 *
 * To run a mutation, you first call `useDeleteAssetMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteAssetMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteAssetMutation, { data, loading, error }] = useDeleteAssetMutation({
 *   variables: {
 *      assetId: // value for 'assetId'
 *   },
 * });
 */
export function useDeleteAssetMutation(baseOptions?: Apollo.MutationHookOptions<DeleteAssetMutation, DeleteAssetMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteAssetMutation, DeleteAssetMutationVariables>(DeleteAssetDocument, options);
      }
export type DeleteAssetMutationHookResult = ReturnType<typeof useDeleteAssetMutation>;
export type DeleteAssetMutationResult = Apollo.MutationResult<DeleteAssetMutation>;
export type DeleteAssetMutationOptions = Apollo.BaseMutationOptions<DeleteAssetMutation, DeleteAssetMutationVariables>;
export const DecompressAssetDocument = gql`
    mutation DecompressAsset($assetId: ID!) {
  decompressAsset(input: {assetId: $assetId}) {
    asset {
      ...assetFragment
    }
  }
}
    ${AssetFragmentFragmentDoc}`;
export type DecompressAssetMutationFn = Apollo.MutationFunction<DecompressAssetMutation, DecompressAssetMutationVariables>;

/**
 * __useDecompressAssetMutation__
 *
 * To run a mutation, you first call `useDecompressAssetMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDecompressAssetMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [decompressAssetMutation, { data, loading, error }] = useDecompressAssetMutation({
 *   variables: {
 *      assetId: // value for 'assetId'
 *   },
 * });
 */
export function useDecompressAssetMutation(baseOptions?: Apollo.MutationHookOptions<DecompressAssetMutation, DecompressAssetMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DecompressAssetMutation, DecompressAssetMutationVariables>(DecompressAssetDocument, options);
      }
export type DecompressAssetMutationHookResult = ReturnType<typeof useDecompressAssetMutation>;
export type DecompressAssetMutationResult = Apollo.MutationResult<DecompressAssetMutation>;
export type DecompressAssetMutationOptions = Apollo.BaseMutationOptions<DecompressAssetMutation, DecompressAssetMutationVariables>;
export const AddCommentDocument = gql`
    mutation AddComment($threadId: ID!, $content: String!) {
  addComment(input: {threadId: $threadId, content: $content}) {
    comment {
      id
      author {
        ... on User {
          id
          name
          email
        }
        ... on Integration {
          id
          name
        }
      }
      authorType
      authorId
      content
      createdAt
    }
  }
}
    `;
export type AddCommentMutationFn = Apollo.MutationFunction<AddCommentMutation, AddCommentMutationVariables>;

/**
 * __useAddCommentMutation__
 *
 * To run a mutation, you first call `useAddCommentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddCommentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addCommentMutation, { data, loading, error }] = useAddCommentMutation({
 *   variables: {
 *      threadId: // value for 'threadId'
 *      content: // value for 'content'
 *   },
 * });
 */
export function useAddCommentMutation(baseOptions?: Apollo.MutationHookOptions<AddCommentMutation, AddCommentMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddCommentMutation, AddCommentMutationVariables>(AddCommentDocument, options);
      }
export type AddCommentMutationHookResult = ReturnType<typeof useAddCommentMutation>;
export type AddCommentMutationResult = Apollo.MutationResult<AddCommentMutation>;
export type AddCommentMutationOptions = Apollo.BaseMutationOptions<AddCommentMutation, AddCommentMutationVariables>;
export const UpdateCommentDocument = gql`
    mutation UpdateComment($commentId: ID!, $threadId: ID!, $content: String!) {
  updateComment(
    input: {commentId: $commentId, threadId: $threadId, content: $content}
  ) {
    comment {
      id
      author {
        ... on User {
          id
          name
          email
        }
        ... on Integration {
          id
          name
        }
      }
      authorType
      authorId
      content
      createdAt
    }
  }
}
    `;
export type UpdateCommentMutationFn = Apollo.MutationFunction<UpdateCommentMutation, UpdateCommentMutationVariables>;

/**
 * __useUpdateCommentMutation__
 *
 * To run a mutation, you first call `useUpdateCommentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateCommentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateCommentMutation, { data, loading, error }] = useUpdateCommentMutation({
 *   variables: {
 *      commentId: // value for 'commentId'
 *      threadId: // value for 'threadId'
 *      content: // value for 'content'
 *   },
 * });
 */
export function useUpdateCommentMutation(baseOptions?: Apollo.MutationHookOptions<UpdateCommentMutation, UpdateCommentMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateCommentMutation, UpdateCommentMutationVariables>(UpdateCommentDocument, options);
      }
export type UpdateCommentMutationHookResult = ReturnType<typeof useUpdateCommentMutation>;
export type UpdateCommentMutationResult = Apollo.MutationResult<UpdateCommentMutation>;
export type UpdateCommentMutationOptions = Apollo.BaseMutationOptions<UpdateCommentMutation, UpdateCommentMutationVariables>;
export const DeleteCommentDocument = gql`
    mutation DeleteComment($commentId: ID!, $threadId: ID!) {
  deleteComment(input: {commentId: $commentId, threadId: $threadId}) {
    commentId
  }
}
    `;
export type DeleteCommentMutationFn = Apollo.MutationFunction<DeleteCommentMutation, DeleteCommentMutationVariables>;

/**
 * __useDeleteCommentMutation__
 *
 * To run a mutation, you first call `useDeleteCommentMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteCommentMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteCommentMutation, { data, loading, error }] = useDeleteCommentMutation({
 *   variables: {
 *      commentId: // value for 'commentId'
 *      threadId: // value for 'threadId'
 *   },
 * });
 */
export function useDeleteCommentMutation(baseOptions?: Apollo.MutationHookOptions<DeleteCommentMutation, DeleteCommentMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteCommentMutation, DeleteCommentMutationVariables>(DeleteCommentDocument, options);
      }
export type DeleteCommentMutationHookResult = ReturnType<typeof useDeleteCommentMutation>;
export type DeleteCommentMutationResult = Apollo.MutationResult<DeleteCommentMutation>;
export type DeleteCommentMutationOptions = Apollo.BaseMutationOptions<DeleteCommentMutation, DeleteCommentMutationVariables>;
export const CreateFieldDocument = gql`
    mutation CreateField($modelId: ID!, $type: SchemaFieldType!, $title: String!, $description: String, $key: String!, $multiple: Boolean!, $unique: Boolean!, $required: Boolean!, $typeProperty: SchemaFieldTypePropertyInput!) {
  createField(
    input: {modelId: $modelId, type: $type, title: $title, description: $description, key: $key, multiple: $multiple, unique: $unique, required: $required, typeProperty: $typeProperty}
  ) {
    field {
      id
    }
  }
}
    `;
export type CreateFieldMutationFn = Apollo.MutationFunction<CreateFieldMutation, CreateFieldMutationVariables>;

/**
 * __useCreateFieldMutation__
 *
 * To run a mutation, you first call `useCreateFieldMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateFieldMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createFieldMutation, { data, loading, error }] = useCreateFieldMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *      type: // value for 'type'
 *      title: // value for 'title'
 *      description: // value for 'description'
 *      key: // value for 'key'
 *      multiple: // value for 'multiple'
 *      unique: // value for 'unique'
 *      required: // value for 'required'
 *      typeProperty: // value for 'typeProperty'
 *   },
 * });
 */
export function useCreateFieldMutation(baseOptions?: Apollo.MutationHookOptions<CreateFieldMutation, CreateFieldMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateFieldMutation, CreateFieldMutationVariables>(CreateFieldDocument, options);
      }
export type CreateFieldMutationHookResult = ReturnType<typeof useCreateFieldMutation>;
export type CreateFieldMutationResult = Apollo.MutationResult<CreateFieldMutation>;
export type CreateFieldMutationOptions = Apollo.BaseMutationOptions<CreateFieldMutation, CreateFieldMutationVariables>;
export const UpdateFieldDocument = gql`
    mutation UpdateField($modelId: ID!, $fieldId: ID!, $title: String!, $description: String, $order: Int, $key: String!, $multiple: Boolean!, $unique: Boolean!, $required: Boolean!, $typeProperty: SchemaFieldTypePropertyInput!) {
  updateField(
    input: {modelId: $modelId, fieldId: $fieldId, title: $title, description: $description, order: $order, key: $key, multiple: $multiple, unique: $unique, required: $required, typeProperty: $typeProperty}
  ) {
    field {
      id
    }
  }
}
    `;
export type UpdateFieldMutationFn = Apollo.MutationFunction<UpdateFieldMutation, UpdateFieldMutationVariables>;

/**
 * __useUpdateFieldMutation__
 *
 * To run a mutation, you first call `useUpdateFieldMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateFieldMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateFieldMutation, { data, loading, error }] = useUpdateFieldMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *      fieldId: // value for 'fieldId'
 *      title: // value for 'title'
 *      description: // value for 'description'
 *      order: // value for 'order'
 *      key: // value for 'key'
 *      multiple: // value for 'multiple'
 *      unique: // value for 'unique'
 *      required: // value for 'required'
 *      typeProperty: // value for 'typeProperty'
 *   },
 * });
 */
export function useUpdateFieldMutation(baseOptions?: Apollo.MutationHookOptions<UpdateFieldMutation, UpdateFieldMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateFieldMutation, UpdateFieldMutationVariables>(UpdateFieldDocument, options);
      }
export type UpdateFieldMutationHookResult = ReturnType<typeof useUpdateFieldMutation>;
export type UpdateFieldMutationResult = Apollo.MutationResult<UpdateFieldMutation>;
export type UpdateFieldMutationOptions = Apollo.BaseMutationOptions<UpdateFieldMutation, UpdateFieldMutationVariables>;
export const UpdateFieldsDocument = gql`
    mutation UpdateFields($updateFieldInput: [UpdateFieldInput!]!) {
  updateFields(input: $updateFieldInput) {
    fields {
      id
    }
  }
}
    `;
export type UpdateFieldsMutationFn = Apollo.MutationFunction<UpdateFieldsMutation, UpdateFieldsMutationVariables>;

/**
 * __useUpdateFieldsMutation__
 *
 * To run a mutation, you first call `useUpdateFieldsMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateFieldsMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateFieldsMutation, { data, loading, error }] = useUpdateFieldsMutation({
 *   variables: {
 *      updateFieldInput: // value for 'updateFieldInput'
 *   },
 * });
 */
export function useUpdateFieldsMutation(baseOptions?: Apollo.MutationHookOptions<UpdateFieldsMutation, UpdateFieldsMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateFieldsMutation, UpdateFieldsMutationVariables>(UpdateFieldsDocument, options);
      }
export type UpdateFieldsMutationHookResult = ReturnType<typeof useUpdateFieldsMutation>;
export type UpdateFieldsMutationResult = Apollo.MutationResult<UpdateFieldsMutation>;
export type UpdateFieldsMutationOptions = Apollo.BaseMutationOptions<UpdateFieldsMutation, UpdateFieldsMutationVariables>;
export const DeleteFieldDocument = gql`
    mutation DeleteField($modelId: ID!, $fieldId: ID!) {
  deleteField(input: {modelId: $modelId, fieldId: $fieldId}) {
    fieldId
  }
}
    `;
export type DeleteFieldMutationFn = Apollo.MutationFunction<DeleteFieldMutation, DeleteFieldMutationVariables>;

/**
 * __useDeleteFieldMutation__
 *
 * To run a mutation, you first call `useDeleteFieldMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteFieldMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteFieldMutation, { data, loading, error }] = useDeleteFieldMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *      fieldId: // value for 'fieldId'
 *   },
 * });
 */
export function useDeleteFieldMutation(baseOptions?: Apollo.MutationHookOptions<DeleteFieldMutation, DeleteFieldMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteFieldMutation, DeleteFieldMutationVariables>(DeleteFieldDocument, options);
      }
export type DeleteFieldMutationHookResult = ReturnType<typeof useDeleteFieldMutation>;
export type DeleteFieldMutationResult = Apollo.MutationResult<DeleteFieldMutation>;
export type DeleteFieldMutationOptions = Apollo.BaseMutationOptions<DeleteFieldMutation, DeleteFieldMutationVariables>;
export const CreateIntegrationDocument = gql`
    mutation CreateIntegration($name: String!, $description: String, $logoUrl: URL!, $type: IntegrationType!) {
  createIntegration(
    input: {name: $name, description: $description, logoUrl: $logoUrl, type: $type}
  ) {
    integration {
      id
      name
      description
      logoUrl
      iType
    }
  }
}
    `;
export type CreateIntegrationMutationFn = Apollo.MutationFunction<CreateIntegrationMutation, CreateIntegrationMutationVariables>;

/**
 * __useCreateIntegrationMutation__
 *
 * To run a mutation, you first call `useCreateIntegrationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateIntegrationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createIntegrationMutation, { data, loading, error }] = useCreateIntegrationMutation({
 *   variables: {
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      logoUrl: // value for 'logoUrl'
 *      type: // value for 'type'
 *   },
 * });
 */
export function useCreateIntegrationMutation(baseOptions?: Apollo.MutationHookOptions<CreateIntegrationMutation, CreateIntegrationMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateIntegrationMutation, CreateIntegrationMutationVariables>(CreateIntegrationDocument, options);
      }
export type CreateIntegrationMutationHookResult = ReturnType<typeof useCreateIntegrationMutation>;
export type CreateIntegrationMutationResult = Apollo.MutationResult<CreateIntegrationMutation>;
export type CreateIntegrationMutationOptions = Apollo.BaseMutationOptions<CreateIntegrationMutation, CreateIntegrationMutationVariables>;
export const UpdateIntegrationDocument = gql`
    mutation UpdateIntegration($integrationId: ID!, $name: String!, $description: String, $logoUrl: URL!) {
  updateIntegration(
    input: {integrationId: $integrationId, name: $name, description: $description, logoUrl: $logoUrl}
  ) {
    integration {
      id
      name
      description
      logoUrl
      iType
    }
  }
}
    `;
export type UpdateIntegrationMutationFn = Apollo.MutationFunction<UpdateIntegrationMutation, UpdateIntegrationMutationVariables>;

/**
 * __useUpdateIntegrationMutation__
 *
 * To run a mutation, you first call `useUpdateIntegrationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateIntegrationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateIntegrationMutation, { data, loading, error }] = useUpdateIntegrationMutation({
 *   variables: {
 *      integrationId: // value for 'integrationId'
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      logoUrl: // value for 'logoUrl'
 *   },
 * });
 */
export function useUpdateIntegrationMutation(baseOptions?: Apollo.MutationHookOptions<UpdateIntegrationMutation, UpdateIntegrationMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateIntegrationMutation, UpdateIntegrationMutationVariables>(UpdateIntegrationDocument, options);
      }
export type UpdateIntegrationMutationHookResult = ReturnType<typeof useUpdateIntegrationMutation>;
export type UpdateIntegrationMutationResult = Apollo.MutationResult<UpdateIntegrationMutation>;
export type UpdateIntegrationMutationOptions = Apollo.BaseMutationOptions<UpdateIntegrationMutation, UpdateIntegrationMutationVariables>;
export const DeleteIntegrationDocument = gql`
    mutation DeleteIntegration($integrationId: ID!) {
  deleteIntegration(input: {integrationId: $integrationId}) {
    integrationId
  }
}
    `;
export type DeleteIntegrationMutationFn = Apollo.MutationFunction<DeleteIntegrationMutation, DeleteIntegrationMutationVariables>;

/**
 * __useDeleteIntegrationMutation__
 *
 * To run a mutation, you first call `useDeleteIntegrationMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteIntegrationMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteIntegrationMutation, { data, loading, error }] = useDeleteIntegrationMutation({
 *   variables: {
 *      integrationId: // value for 'integrationId'
 *   },
 * });
 */
export function useDeleteIntegrationMutation(baseOptions?: Apollo.MutationHookOptions<DeleteIntegrationMutation, DeleteIntegrationMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteIntegrationMutation, DeleteIntegrationMutationVariables>(DeleteIntegrationDocument, options);
      }
export type DeleteIntegrationMutationHookResult = ReturnType<typeof useDeleteIntegrationMutation>;
export type DeleteIntegrationMutationResult = Apollo.MutationResult<DeleteIntegrationMutation>;
export type DeleteIntegrationMutationOptions = Apollo.BaseMutationOptions<DeleteIntegrationMutation, DeleteIntegrationMutationVariables>;
export const GetItemsDocument = gql`
    query GetItems($schemaId: ID!, $pagination: Pagination) {
  items(schemaId: $schemaId, pagination: $pagination) {
    nodes {
      id
      schemaId
      createdAt
      updatedAt
      status
      user {
        name
      }
      integration {
        name
      }
      fields {
        schemaFieldId
        type
        value
      }
      thread {
        ...threadFragment
      }
    }
  }
}
    ${ThreadFragmentFragmentDoc}`;

/**
 * __useGetItemsQuery__
 *
 * To run a query within a React component, call `useGetItemsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetItemsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetItemsQuery({
 *   variables: {
 *      schemaId: // value for 'schemaId'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useGetItemsQuery(baseOptions: Apollo.QueryHookOptions<GetItemsQuery, GetItemsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetItemsQuery, GetItemsQueryVariables>(GetItemsDocument, options);
      }
export function useGetItemsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetItemsQuery, GetItemsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetItemsQuery, GetItemsQueryVariables>(GetItemsDocument, options);
        }
export type GetItemsQueryHookResult = ReturnType<typeof useGetItemsQuery>;
export type GetItemsLazyQueryHookResult = ReturnType<typeof useGetItemsLazyQuery>;
export type GetItemsQueryResult = Apollo.QueryResult<GetItemsQuery, GetItemsQueryVariables>;
export const GetItemDocument = gql`
    query GetItem($id: ID!) {
  node(id: $id, type: Item) {
    ... on Item {
      id
      schemaId
      createdAt
      updatedAt
      status
      assets {
        id
        url
      }
      user {
        name
      }
      integration {
        name
      }
      fields {
        schemaFieldId
        type
        value
      }
      thread {
        ...threadFragment
      }
    }
  }
}
    ${ThreadFragmentFragmentDoc}`;

/**
 * __useGetItemQuery__
 *
 * To run a query within a React component, call `useGetItemQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetItemQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetItemQuery({
 *   variables: {
 *      id: // value for 'id'
 *   },
 * });
 */
export function useGetItemQuery(baseOptions: Apollo.QueryHookOptions<GetItemQuery, GetItemQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetItemQuery, GetItemQueryVariables>(GetItemDocument, options);
      }
export function useGetItemLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetItemQuery, GetItemQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetItemQuery, GetItemQueryVariables>(GetItemDocument, options);
        }
export type GetItemQueryHookResult = ReturnType<typeof useGetItemQuery>;
export type GetItemLazyQueryHookResult = ReturnType<typeof useGetItemLazyQuery>;
export type GetItemQueryResult = Apollo.QueryResult<GetItemQuery, GetItemQueryVariables>;
export const SearchItemDocument = gql`
    query SearchItem($query: ItemQuery!, $sort: ItemSort, $pagination: Pagination) {
  searchItem(query: $query, sort: $sort, pagination: $pagination) {
    nodes {
      id
      schemaId
      createdAt
      updatedAt
      status
      assets {
        id
        url
      }
      user {
        name
      }
      integration {
        name
      }
      fields {
        schemaFieldId
        type
        value
      }
      thread {
        ...threadFragment
      }
    }
    totalCount
  }
}
    ${ThreadFragmentFragmentDoc}`;

/**
 * __useSearchItemQuery__
 *
 * To run a query within a React component, call `useSearchItemQuery` and pass it any options that fit your needs.
 * When your component renders, `useSearchItemQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useSearchItemQuery({
 *   variables: {
 *      query: // value for 'query'
 *      sort: // value for 'sort'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useSearchItemQuery(baseOptions: Apollo.QueryHookOptions<SearchItemQuery, SearchItemQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<SearchItemQuery, SearchItemQueryVariables>(SearchItemDocument, options);
      }
export function useSearchItemLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<SearchItemQuery, SearchItemQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<SearchItemQuery, SearchItemQueryVariables>(SearchItemDocument, options);
        }
export type SearchItemQueryHookResult = ReturnType<typeof useSearchItemQuery>;
export type SearchItemLazyQueryHookResult = ReturnType<typeof useSearchItemLazyQuery>;
export type SearchItemQueryResult = Apollo.QueryResult<SearchItemQuery, SearchItemQueryVariables>;
export const CreateItemDocument = gql`
    mutation CreateItem($modelId: ID!, $schemaId: ID!, $fields: [ItemFieldInput!]!) {
  createItem(input: {modelId: $modelId, schemaId: $schemaId, fields: $fields}) {
    item {
      id
      schemaId
      fields {
        value
        type
        schemaFieldId
      }
    }
  }
}
    `;
export type CreateItemMutationFn = Apollo.MutationFunction<CreateItemMutation, CreateItemMutationVariables>;

/**
 * __useCreateItemMutation__
 *
 * To run a mutation, you first call `useCreateItemMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateItemMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createItemMutation, { data, loading, error }] = useCreateItemMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *      schemaId: // value for 'schemaId'
 *      fields: // value for 'fields'
 *   },
 * });
 */
export function useCreateItemMutation(baseOptions?: Apollo.MutationHookOptions<CreateItemMutation, CreateItemMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateItemMutation, CreateItemMutationVariables>(CreateItemDocument, options);
      }
export type CreateItemMutationHookResult = ReturnType<typeof useCreateItemMutation>;
export type CreateItemMutationResult = Apollo.MutationResult<CreateItemMutation>;
export type CreateItemMutationOptions = Apollo.BaseMutationOptions<CreateItemMutation, CreateItemMutationVariables>;
export const DeleteItemDocument = gql`
    mutation DeleteItem($itemId: ID!) {
  deleteItem(input: {itemId: $itemId}) {
    itemId
  }
}
    `;
export type DeleteItemMutationFn = Apollo.MutationFunction<DeleteItemMutation, DeleteItemMutationVariables>;

/**
 * __useDeleteItemMutation__
 *
 * To run a mutation, you first call `useDeleteItemMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteItemMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteItemMutation, { data, loading, error }] = useDeleteItemMutation({
 *   variables: {
 *      itemId: // value for 'itemId'
 *   },
 * });
 */
export function useDeleteItemMutation(baseOptions?: Apollo.MutationHookOptions<DeleteItemMutation, DeleteItemMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteItemMutation, DeleteItemMutationVariables>(DeleteItemDocument, options);
      }
export type DeleteItemMutationHookResult = ReturnType<typeof useDeleteItemMutation>;
export type DeleteItemMutationResult = Apollo.MutationResult<DeleteItemMutation>;
export type DeleteItemMutationOptions = Apollo.BaseMutationOptions<DeleteItemMutation, DeleteItemMutationVariables>;
export const UpdateItemDocument = gql`
    mutation UpdateItem($itemId: ID!, $fields: [ItemFieldInput!]!) {
  updateItem(input: {itemId: $itemId, fields: $fields}) {
    item {
      id
      schemaId
      fields {
        value
        type
        schemaFieldId
      }
    }
  }
}
    `;
export type UpdateItemMutationFn = Apollo.MutationFunction<UpdateItemMutation, UpdateItemMutationVariables>;

/**
 * __useUpdateItemMutation__
 *
 * To run a mutation, you first call `useUpdateItemMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateItemMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateItemMutation, { data, loading, error }] = useUpdateItemMutation({
 *   variables: {
 *      itemId: // value for 'itemId'
 *      fields: // value for 'fields'
 *   },
 * });
 */
export function useUpdateItemMutation(baseOptions?: Apollo.MutationHookOptions<UpdateItemMutation, UpdateItemMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateItemMutation, UpdateItemMutationVariables>(UpdateItemDocument, options);
      }
export type UpdateItemMutationHookResult = ReturnType<typeof useUpdateItemMutation>;
export type UpdateItemMutationResult = Apollo.MutationResult<UpdateItemMutation>;
export type UpdateItemMutationOptions = Apollo.BaseMutationOptions<UpdateItemMutation, UpdateItemMutationVariables>;
export const UnpublishItemDocument = gql`
    mutation UnpublishItem($itemId: [ID!]!) {
  unpublishItem(input: {itemId: $itemId}) {
    items {
      id
    }
  }
}
    `;
export type UnpublishItemMutationFn = Apollo.MutationFunction<UnpublishItemMutation, UnpublishItemMutationVariables>;

/**
 * __useUnpublishItemMutation__
 *
 * To run a mutation, you first call `useUnpublishItemMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUnpublishItemMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [unpublishItemMutation, { data, loading, error }] = useUnpublishItemMutation({
 *   variables: {
 *      itemId: // value for 'itemId'
 *   },
 * });
 */
export function useUnpublishItemMutation(baseOptions?: Apollo.MutationHookOptions<UnpublishItemMutation, UnpublishItemMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UnpublishItemMutation, UnpublishItemMutationVariables>(UnpublishItemDocument, options);
      }
export type UnpublishItemMutationHookResult = ReturnType<typeof useUnpublishItemMutation>;
export type UnpublishItemMutationResult = Apollo.MutationResult<UnpublishItemMutation>;
export type UnpublishItemMutationOptions = Apollo.BaseMutationOptions<UnpublishItemMutation, UnpublishItemMutationVariables>;
export const GetModelsDocument = gql`
    query GetModels($projectId: ID!, $pagination: Pagination) {
  models(projectId: $projectId, pagination: $pagination) {
    nodes {
      id
      name
      description
      key
      public
      schema {
        id
        fields {
          id
          type
          title
          key
          description
          required
          unique
          multiple
          order
          typeProperty {
            ... on SchemaFieldText {
              defaultValue
              maxLength
            }
            ... on SchemaFieldTextArea {
              defaultValue
              maxLength
            }
            ... on SchemaFieldMarkdown {
              defaultValue
              maxLength
            }
            ... on SchemaFieldAsset {
              assetDefaultValue: defaultValue
            }
            ... on SchemaFieldSelect {
              selectDefaultValue: defaultValue
              values
            }
            ... on SchemaFieldInteger {
              integerDefaultValue: defaultValue
              min
              max
            }
            ... on SchemaFieldBool {
              defaultValue
            }
            ... on SchemaFieldURL {
              defaultValue
            }
          }
        }
      }
    }
  }
}
    `;

/**
 * __useGetModelsQuery__
 *
 * To run a query within a React component, call `useGetModelsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetModelsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetModelsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useGetModelsQuery(baseOptions: Apollo.QueryHookOptions<GetModelsQuery, GetModelsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetModelsQuery, GetModelsQueryVariables>(GetModelsDocument, options);
      }
export function useGetModelsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetModelsQuery, GetModelsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetModelsQuery, GetModelsQueryVariables>(GetModelsDocument, options);
        }
export type GetModelsQueryHookResult = ReturnType<typeof useGetModelsQuery>;
export type GetModelsLazyQueryHookResult = ReturnType<typeof useGetModelsLazyQuery>;
export type GetModelsQueryResult = Apollo.QueryResult<GetModelsQuery, GetModelsQueryVariables>;
export const CreateModelDocument = gql`
    mutation CreateModel($projectId: ID!, $name: String, $description: String, $key: String) {
  createModel(
    input: {projectId: $projectId, name: $name, description: $description, key: $key}
  ) {
    model {
      id
      name
    }
  }
}
    `;
export type CreateModelMutationFn = Apollo.MutationFunction<CreateModelMutation, CreateModelMutationVariables>;

/**
 * __useCreateModelMutation__
 *
 * To run a mutation, you first call `useCreateModelMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateModelMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createModelMutation, { data, loading, error }] = useCreateModelMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      key: // value for 'key'
 *   },
 * });
 */
export function useCreateModelMutation(baseOptions?: Apollo.MutationHookOptions<CreateModelMutation, CreateModelMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateModelMutation, CreateModelMutationVariables>(CreateModelDocument, options);
      }
export type CreateModelMutationHookResult = ReturnType<typeof useCreateModelMutation>;
export type CreateModelMutationResult = Apollo.MutationResult<CreateModelMutation>;
export type CreateModelMutationOptions = Apollo.BaseMutationOptions<CreateModelMutation, CreateModelMutationVariables>;
export const DeleteModelDocument = gql`
    mutation DeleteModel($modelId: ID!) {
  deleteModel(input: {modelId: $modelId}) {
    modelId
  }
}
    `;
export type DeleteModelMutationFn = Apollo.MutationFunction<DeleteModelMutation, DeleteModelMutationVariables>;

/**
 * __useDeleteModelMutation__
 *
 * To run a mutation, you first call `useDeleteModelMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteModelMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteModelMutation, { data, loading, error }] = useDeleteModelMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *   },
 * });
 */
export function useDeleteModelMutation(baseOptions?: Apollo.MutationHookOptions<DeleteModelMutation, DeleteModelMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteModelMutation, DeleteModelMutationVariables>(DeleteModelDocument, options);
      }
export type DeleteModelMutationHookResult = ReturnType<typeof useDeleteModelMutation>;
export type DeleteModelMutationResult = Apollo.MutationResult<DeleteModelMutation>;
export type DeleteModelMutationOptions = Apollo.BaseMutationOptions<DeleteModelMutation, DeleteModelMutationVariables>;
export const UpdateModelDocument = gql`
    mutation UpdateModel($modelId: ID!, $name: String, $description: String, $key: String, $public: Boolean!) {
  updateModel(
    input: {modelId: $modelId, name: $name, description: $description, key: $key, public: $public}
  ) {
    model {
      id
      name
    }
  }
}
    `;
export type UpdateModelMutationFn = Apollo.MutationFunction<UpdateModelMutation, UpdateModelMutationVariables>;

/**
 * __useUpdateModelMutation__
 *
 * To run a mutation, you first call `useUpdateModelMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateModelMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateModelMutation, { data, loading, error }] = useUpdateModelMutation({
 *   variables: {
 *      modelId: // value for 'modelId'
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      key: // value for 'key'
 *      public: // value for 'public'
 *   },
 * });
 */
export function useUpdateModelMutation(baseOptions?: Apollo.MutationHookOptions<UpdateModelMutation, UpdateModelMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateModelMutation, UpdateModelMutationVariables>(UpdateModelDocument, options);
      }
export type UpdateModelMutationHookResult = ReturnType<typeof useUpdateModelMutation>;
export type UpdateModelMutationResult = Apollo.MutationResult<UpdateModelMutation>;
export type UpdateModelMutationOptions = Apollo.BaseMutationOptions<UpdateModelMutation, UpdateModelMutationVariables>;
export const CheckModelKeyAvailabilityDocument = gql`
    query CheckModelKeyAvailability($projectId: ID!, $key: String!) {
  checkModelKeyAvailability(projectId: $projectId, key: $key) {
    key
    available
  }
}
    `;

/**
 * __useCheckModelKeyAvailabilityQuery__
 *
 * To run a query within a React component, call `useCheckModelKeyAvailabilityQuery` and pass it any options that fit your needs.
 * When your component renders, `useCheckModelKeyAvailabilityQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCheckModelKeyAvailabilityQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      key: // value for 'key'
 *   },
 * });
 */
export function useCheckModelKeyAvailabilityQuery(baseOptions: Apollo.QueryHookOptions<CheckModelKeyAvailabilityQuery, CheckModelKeyAvailabilityQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CheckModelKeyAvailabilityQuery, CheckModelKeyAvailabilityQueryVariables>(CheckModelKeyAvailabilityDocument, options);
      }
export function useCheckModelKeyAvailabilityLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CheckModelKeyAvailabilityQuery, CheckModelKeyAvailabilityQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CheckModelKeyAvailabilityQuery, CheckModelKeyAvailabilityQueryVariables>(CheckModelKeyAvailabilityDocument, options);
        }
export type CheckModelKeyAvailabilityQueryHookResult = ReturnType<typeof useCheckModelKeyAvailabilityQuery>;
export type CheckModelKeyAvailabilityLazyQueryHookResult = ReturnType<typeof useCheckModelKeyAvailabilityLazyQuery>;
export type CheckModelKeyAvailabilityQueryResult = Apollo.QueryResult<CheckModelKeyAvailabilityQuery, CheckModelKeyAvailabilityQueryVariables>;
export const GetProjectDocument = gql`
    query GetProject($projectId: ID!) {
  node(id: $projectId, type: PROJECT) {
    id
    ... on Project {
      name
      description
      alias
      publication {
        scope
        assetPublic
      }
    }
  }
}
    `;

/**
 * __useGetProjectQuery__
 *
 * To run a query within a React component, call `useGetProjectQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetProjectQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetProjectQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useGetProjectQuery(baseOptions: Apollo.QueryHookOptions<GetProjectQuery, GetProjectQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetProjectQuery, GetProjectQueryVariables>(GetProjectDocument, options);
      }
export function useGetProjectLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetProjectQuery, GetProjectQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetProjectQuery, GetProjectQueryVariables>(GetProjectDocument, options);
        }
export type GetProjectQueryHookResult = ReturnType<typeof useGetProjectQuery>;
export type GetProjectLazyQueryHookResult = ReturnType<typeof useGetProjectLazyQuery>;
export type GetProjectQueryResult = Apollo.QueryResult<GetProjectQuery, GetProjectQueryVariables>;
export const GetProjectsDocument = gql`
    query GetProjects($workspaceId: ID!, $pagination: Pagination) {
  projects(workspaceId: $workspaceId, pagination: $pagination) {
    nodes {
      id
      name
      description
      alias
      publication {
        scope
        assetPublic
      }
    }
  }
}
    `;

/**
 * __useGetProjectsQuery__
 *
 * To run a query within a React component, call `useGetProjectsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetProjectsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetProjectsQuery({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      pagination: // value for 'pagination'
 *   },
 * });
 */
export function useGetProjectsQuery(baseOptions: Apollo.QueryHookOptions<GetProjectsQuery, GetProjectsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetProjectsQuery, GetProjectsQueryVariables>(GetProjectsDocument, options);
      }
export function useGetProjectsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetProjectsQuery, GetProjectsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetProjectsQuery, GetProjectsQueryVariables>(GetProjectsDocument, options);
        }
export type GetProjectsQueryHookResult = ReturnType<typeof useGetProjectsQuery>;
export type GetProjectsLazyQueryHookResult = ReturnType<typeof useGetProjectsLazyQuery>;
export type GetProjectsQueryResult = Apollo.QueryResult<GetProjectsQuery, GetProjectsQueryVariables>;
export const CheckProjectAliasDocument = gql`
    query CheckProjectAlias($alias: String!) {
  checkProjectAlias(alias: $alias) {
    alias
    available
  }
}
    `;

/**
 * __useCheckProjectAliasQuery__
 *
 * To run a query within a React component, call `useCheckProjectAliasQuery` and pass it any options that fit your needs.
 * When your component renders, `useCheckProjectAliasQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useCheckProjectAliasQuery({
 *   variables: {
 *      alias: // value for 'alias'
 *   },
 * });
 */
export function useCheckProjectAliasQuery(baseOptions: Apollo.QueryHookOptions<CheckProjectAliasQuery, CheckProjectAliasQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<CheckProjectAliasQuery, CheckProjectAliasQueryVariables>(CheckProjectAliasDocument, options);
      }
export function useCheckProjectAliasLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<CheckProjectAliasQuery, CheckProjectAliasQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<CheckProjectAliasQuery, CheckProjectAliasQueryVariables>(CheckProjectAliasDocument, options);
        }
export type CheckProjectAliasQueryHookResult = ReturnType<typeof useCheckProjectAliasQuery>;
export type CheckProjectAliasLazyQueryHookResult = ReturnType<typeof useCheckProjectAliasLazyQuery>;
export type CheckProjectAliasQueryResult = Apollo.QueryResult<CheckProjectAliasQuery, CheckProjectAliasQueryVariables>;
export const CreateProjectDocument = gql`
    mutation CreateProject($workspaceId: ID!, $name: String!, $description: String!, $alias: String!) {
  createProject(
    input: {workspaceId: $workspaceId, name: $name, description: $description, alias: $alias}
  ) {
    project {
      id
      name
      description
      alias
      publication {
        scope
        assetPublic
      }
    }
  }
}
    `;
export type CreateProjectMutationFn = Apollo.MutationFunction<CreateProjectMutation, CreateProjectMutationVariables>;

/**
 * __useCreateProjectMutation__
 *
 * To run a mutation, you first call `useCreateProjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateProjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createProjectMutation, { data, loading, error }] = useCreateProjectMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      alias: // value for 'alias'
 *   },
 * });
 */
export function useCreateProjectMutation(baseOptions?: Apollo.MutationHookOptions<CreateProjectMutation, CreateProjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateProjectMutation, CreateProjectMutationVariables>(CreateProjectDocument, options);
      }
export type CreateProjectMutationHookResult = ReturnType<typeof useCreateProjectMutation>;
export type CreateProjectMutationResult = Apollo.MutationResult<CreateProjectMutation>;
export type CreateProjectMutationOptions = Apollo.BaseMutationOptions<CreateProjectMutation, CreateProjectMutationVariables>;
export const DeleteProjectDocument = gql`
    mutation DeleteProject($projectId: ID!) {
  deleteProject(input: {projectId: $projectId}) {
    projectId
  }
}
    `;
export type DeleteProjectMutationFn = Apollo.MutationFunction<DeleteProjectMutation, DeleteProjectMutationVariables>;

/**
 * __useDeleteProjectMutation__
 *
 * To run a mutation, you first call `useDeleteProjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteProjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteProjectMutation, { data, loading, error }] = useDeleteProjectMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *   },
 * });
 */
export function useDeleteProjectMutation(baseOptions?: Apollo.MutationHookOptions<DeleteProjectMutation, DeleteProjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteProjectMutation, DeleteProjectMutationVariables>(DeleteProjectDocument, options);
      }
export type DeleteProjectMutationHookResult = ReturnType<typeof useDeleteProjectMutation>;
export type DeleteProjectMutationResult = Apollo.MutationResult<DeleteProjectMutation>;
export type DeleteProjectMutationOptions = Apollo.BaseMutationOptions<DeleteProjectMutation, DeleteProjectMutationVariables>;
export const UpdateProjectDocument = gql`
    mutation UpdateProject($projectId: ID!, $name: String, $description: String, $alias: String, $publication: UpdateProjectPublicationInput) {
  updateProject(
    input: {projectId: $projectId, name: $name, description: $description, alias: $alias, publication: $publication}
  ) {
    project {
      id
      name
      description
      alias
      publication {
        scope
        assetPublic
      }
    }
  }
}
    `;
export type UpdateProjectMutationFn = Apollo.MutationFunction<UpdateProjectMutation, UpdateProjectMutationVariables>;

/**
 * __useUpdateProjectMutation__
 *
 * To run a mutation, you first call `useUpdateProjectMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateProjectMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateProjectMutation, { data, loading, error }] = useUpdateProjectMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      name: // value for 'name'
 *      description: // value for 'description'
 *      alias: // value for 'alias'
 *      publication: // value for 'publication'
 *   },
 * });
 */
export function useUpdateProjectMutation(baseOptions?: Apollo.MutationHookOptions<UpdateProjectMutation, UpdateProjectMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateProjectMutation, UpdateProjectMutationVariables>(UpdateProjectDocument, options);
      }
export type UpdateProjectMutationHookResult = ReturnType<typeof useUpdateProjectMutation>;
export type UpdateProjectMutationResult = Apollo.MutationResult<UpdateProjectMutation>;
export type UpdateProjectMutationOptions = Apollo.BaseMutationOptions<UpdateProjectMutation, UpdateProjectMutationVariables>;
export const GetRequestsDocument = gql`
    query GetRequests($projectId: ID!, $key: String, $state: [RequestState!], $pagination: Pagination, $createdBy: ID, $reviewer: ID, $sort: Sort) {
  requests(
    projectId: $projectId
    key: $key
    state: $state
    pagination: $pagination
    createdBy: $createdBy
    reviewer: $reviewer
    sort: $sort
  ) {
    nodes {
      id
      title
      description
      createdBy {
        id
        name
        email
      }
      workspaceId
      projectId
      threadId
      reviewersId
      reviewers {
        id
        name
        email
      }
      state
      createdAt
      updatedAt
      approvedAt
      closedAt
      thread {
        ...threadFragment
      }
    }
    totalCount
  }
}
    ${ThreadFragmentFragmentDoc}`;

/**
 * __useGetRequestsQuery__
 *
 * To run a query within a React component, call `useGetRequestsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetRequestsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetRequestsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      key: // value for 'key'
 *      state: // value for 'state'
 *      pagination: // value for 'pagination'
 *      createdBy: // value for 'createdBy'
 *      reviewer: // value for 'reviewer'
 *      sort: // value for 'sort'
 *   },
 * });
 */
export function useGetRequestsQuery(baseOptions: Apollo.QueryHookOptions<GetRequestsQuery, GetRequestsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetRequestsQuery, GetRequestsQueryVariables>(GetRequestsDocument, options);
      }
export function useGetRequestsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetRequestsQuery, GetRequestsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetRequestsQuery, GetRequestsQueryVariables>(GetRequestsDocument, options);
        }
export type GetRequestsQueryHookResult = ReturnType<typeof useGetRequestsQuery>;
export type GetRequestsLazyQueryHookResult = ReturnType<typeof useGetRequestsLazyQuery>;
export type GetRequestsQueryResult = Apollo.QueryResult<GetRequestsQuery, GetRequestsQueryVariables>;
export const GetModalRequestsDocument = gql`
    query GetModalRequests($projectId: ID!, $key: String, $state: [RequestState!], $pagination: Pagination, $createdBy: ID, $reviewer: ID, $sort: Sort) {
  requests(
    projectId: $projectId
    key: $key
    state: $state
    pagination: $pagination
    createdBy: $createdBy
    reviewer: $reviewer
    sort: $sort
  ) {
    nodes {
      id
      title
      description
      createdBy {
        name
      }
      items {
        itemId
      }
      reviewers {
        id
        name
      }
      state
      createdAt
    }
    totalCount
  }
}
    `;

/**
 * __useGetModalRequestsQuery__
 *
 * To run a query within a React component, call `useGetModalRequestsQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetModalRequestsQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetModalRequestsQuery({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      key: // value for 'key'
 *      state: // value for 'state'
 *      pagination: // value for 'pagination'
 *      createdBy: // value for 'createdBy'
 *      reviewer: // value for 'reviewer'
 *      sort: // value for 'sort'
 *   },
 * });
 */
export function useGetModalRequestsQuery(baseOptions: Apollo.QueryHookOptions<GetModalRequestsQuery, GetModalRequestsQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetModalRequestsQuery, GetModalRequestsQueryVariables>(GetModalRequestsDocument, options);
      }
export function useGetModalRequestsLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetModalRequestsQuery, GetModalRequestsQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetModalRequestsQuery, GetModalRequestsQueryVariables>(GetModalRequestsDocument, options);
        }
export type GetModalRequestsQueryHookResult = ReturnType<typeof useGetModalRequestsQuery>;
export type GetModalRequestsLazyQueryHookResult = ReturnType<typeof useGetModalRequestsLazyQuery>;
export type GetModalRequestsQueryResult = Apollo.QueryResult<GetModalRequestsQuery, GetModalRequestsQueryVariables>;
export const GetRequestDocument = gql`
    query GetRequest($requestId: ID!) {
  node(id: $requestId, type: REQUEST) {
    id
    ... on Request {
      ...requestFragment
    }
  }
}
    ${RequestFragmentFragmentDoc}`;

/**
 * __useGetRequestQuery__
 *
 * To run a query within a React component, call `useGetRequestQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetRequestQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetRequestQuery({
 *   variables: {
 *      requestId: // value for 'requestId'
 *   },
 * });
 */
export function useGetRequestQuery(baseOptions: Apollo.QueryHookOptions<GetRequestQuery, GetRequestQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetRequestQuery, GetRequestQueryVariables>(GetRequestDocument, options);
      }
export function useGetRequestLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetRequestQuery, GetRequestQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetRequestQuery, GetRequestQueryVariables>(GetRequestDocument, options);
        }
export type GetRequestQueryHookResult = ReturnType<typeof useGetRequestQuery>;
export type GetRequestLazyQueryHookResult = ReturnType<typeof useGetRequestLazyQuery>;
export type GetRequestQueryResult = Apollo.QueryResult<GetRequestQuery, GetRequestQueryVariables>;
export const CreateRequestDocument = gql`
    mutation CreateRequest($projectId: ID!, $title: String!, $description: String, $state: RequestState, $reviewersId: [ID!], $items: [RequestItemInput!]!) {
  createRequest(
    input: {projectId: $projectId, title: $title, description: $description, state: $state, reviewersId: $reviewersId, items: $items}
  ) {
    request {
      ...requestFragment
    }
  }
}
    ${RequestFragmentFragmentDoc}`;
export type CreateRequestMutationFn = Apollo.MutationFunction<CreateRequestMutation, CreateRequestMutationVariables>;

/**
 * __useCreateRequestMutation__
 *
 * To run a mutation, you first call `useCreateRequestMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateRequestMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createRequestMutation, { data, loading, error }] = useCreateRequestMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      title: // value for 'title'
 *      description: // value for 'description'
 *      state: // value for 'state'
 *      reviewersId: // value for 'reviewersId'
 *      items: // value for 'items'
 *   },
 * });
 */
export function useCreateRequestMutation(baseOptions?: Apollo.MutationHookOptions<CreateRequestMutation, CreateRequestMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateRequestMutation, CreateRequestMutationVariables>(CreateRequestDocument, options);
      }
export type CreateRequestMutationHookResult = ReturnType<typeof useCreateRequestMutation>;
export type CreateRequestMutationResult = Apollo.MutationResult<CreateRequestMutation>;
export type CreateRequestMutationOptions = Apollo.BaseMutationOptions<CreateRequestMutation, CreateRequestMutationVariables>;
export const UpdateRequestDocument = gql`
    mutation UpdateRequest($requestId: ID!, $title: String, $description: String, $state: RequestState, $reviewersId: [ID!], $items: [RequestItemInput!]) {
  updateRequest(
    input: {requestId: $requestId, title: $title, description: $description, state: $state, reviewersId: $reviewersId, items: $items}
  ) {
    request {
      ...requestFragment
    }
  }
}
    ${RequestFragmentFragmentDoc}`;
export type UpdateRequestMutationFn = Apollo.MutationFunction<UpdateRequestMutation, UpdateRequestMutationVariables>;

/**
 * __useUpdateRequestMutation__
 *
 * To run a mutation, you first call `useUpdateRequestMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateRequestMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateRequestMutation, { data, loading, error }] = useUpdateRequestMutation({
 *   variables: {
 *      requestId: // value for 'requestId'
 *      title: // value for 'title'
 *      description: // value for 'description'
 *      state: // value for 'state'
 *      reviewersId: // value for 'reviewersId'
 *      items: // value for 'items'
 *   },
 * });
 */
export function useUpdateRequestMutation(baseOptions?: Apollo.MutationHookOptions<UpdateRequestMutation, UpdateRequestMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateRequestMutation, UpdateRequestMutationVariables>(UpdateRequestDocument, options);
      }
export type UpdateRequestMutationHookResult = ReturnType<typeof useUpdateRequestMutation>;
export type UpdateRequestMutationResult = Apollo.MutationResult<UpdateRequestMutation>;
export type UpdateRequestMutationOptions = Apollo.BaseMutationOptions<UpdateRequestMutation, UpdateRequestMutationVariables>;
export const ApproveRequestDocument = gql`
    mutation ApproveRequest($requestId: ID!) {
  approveRequest(input: {requestId: $requestId}) {
    request {
      ...requestFragment
    }
  }
}
    ${RequestFragmentFragmentDoc}`;
export type ApproveRequestMutationFn = Apollo.MutationFunction<ApproveRequestMutation, ApproveRequestMutationVariables>;

/**
 * __useApproveRequestMutation__
 *
 * To run a mutation, you first call `useApproveRequestMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useApproveRequestMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [approveRequestMutation, { data, loading, error }] = useApproveRequestMutation({
 *   variables: {
 *      requestId: // value for 'requestId'
 *   },
 * });
 */
export function useApproveRequestMutation(baseOptions?: Apollo.MutationHookOptions<ApproveRequestMutation, ApproveRequestMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<ApproveRequestMutation, ApproveRequestMutationVariables>(ApproveRequestDocument, options);
      }
export type ApproveRequestMutationHookResult = ReturnType<typeof useApproveRequestMutation>;
export type ApproveRequestMutationResult = Apollo.MutationResult<ApproveRequestMutation>;
export type ApproveRequestMutationOptions = Apollo.BaseMutationOptions<ApproveRequestMutation, ApproveRequestMutationVariables>;
export const DeleteRequestDocument = gql`
    mutation DeleteRequest($projectId: ID!, $requestsId: [ID!]!) {
  deleteRequest(input: {projectId: $projectId, requestsId: $requestsId}) {
    requests
  }
}
    `;
export type DeleteRequestMutationFn = Apollo.MutationFunction<DeleteRequestMutation, DeleteRequestMutationVariables>;

/**
 * __useDeleteRequestMutation__
 *
 * To run a mutation, you first call `useDeleteRequestMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteRequestMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteRequestMutation, { data, loading, error }] = useDeleteRequestMutation({
 *   variables: {
 *      projectId: // value for 'projectId'
 *      requestsId: // value for 'requestsId'
 *   },
 * });
 */
export function useDeleteRequestMutation(baseOptions?: Apollo.MutationHookOptions<DeleteRequestMutation, DeleteRequestMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteRequestMutation, DeleteRequestMutationVariables>(DeleteRequestDocument, options);
      }
export type DeleteRequestMutationHookResult = ReturnType<typeof useDeleteRequestMutation>;
export type DeleteRequestMutationResult = Apollo.MutationResult<DeleteRequestMutation>;
export type DeleteRequestMutationOptions = Apollo.BaseMutationOptions<DeleteRequestMutation, DeleteRequestMutationVariables>;
export const GetUserBySearchDocument = gql`
    query GetUserBySearch($nameOrEmail: String!) {
  searchUser(nameOrEmail: $nameOrEmail) {
    id
    name
    email
  }
}
    `;

/**
 * __useGetUserBySearchQuery__
 *
 * To run a query within a React component, call `useGetUserBySearchQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetUserBySearchQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetUserBySearchQuery({
 *   variables: {
 *      nameOrEmail: // value for 'nameOrEmail'
 *   },
 * });
 */
export function useGetUserBySearchQuery(baseOptions: Apollo.QueryHookOptions<GetUserBySearchQuery, GetUserBySearchQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetUserBySearchQuery, GetUserBySearchQueryVariables>(GetUserBySearchDocument, options);
      }
export function useGetUserBySearchLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetUserBySearchQuery, GetUserBySearchQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetUserBySearchQuery, GetUserBySearchQueryVariables>(GetUserBySearchDocument, options);
        }
export type GetUserBySearchQueryHookResult = ReturnType<typeof useGetUserBySearchQuery>;
export type GetUserBySearchLazyQueryHookResult = ReturnType<typeof useGetUserBySearchLazyQuery>;
export type GetUserBySearchQueryResult = Apollo.QueryResult<GetUserBySearchQuery, GetUserBySearchQueryVariables>;
export const GetMeDocument = gql`
    query GetMe {
  me {
    id
    name
    email
    lang
    myWorkspace {
      id
      name
    }
    workspaces {
      id
      name
      members {
        ... on WorkspaceUserMember {
          user {
            id
            name
            email
          }
          userId
          role
        }
        ... on WorkspaceIntegrationMember {
          integration {
            ...integrationFragment
          }
          integrationRole: role
          active
          invitedById
        }
      }
    }
    auths
    integrations {
      ...integrationFragment
    }
  }
}
    ${IntegrationFragmentFragmentDoc}`;

/**
 * __useGetMeQuery__
 *
 * To run a query within a React component, call `useGetMeQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetMeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetMeQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetMeQuery(baseOptions?: Apollo.QueryHookOptions<GetMeQuery, GetMeQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetMeQuery, GetMeQueryVariables>(GetMeDocument, options);
      }
export function useGetMeLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetMeQuery, GetMeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetMeQuery, GetMeQueryVariables>(GetMeDocument, options);
        }
export type GetMeQueryHookResult = ReturnType<typeof useGetMeQuery>;
export type GetMeLazyQueryHookResult = ReturnType<typeof useGetMeLazyQuery>;
export type GetMeQueryResult = Apollo.QueryResult<GetMeQuery, GetMeQueryVariables>;
export const GetProfileDocument = gql`
    query GetProfile {
  me {
    id
    name
    email
    lang
    theme
    myWorkspace {
      id
      name
    }
    auths
  }
}
    `;

/**
 * __useGetProfileQuery__
 *
 * To run a query within a React component, call `useGetProfileQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetProfileQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetProfileQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetProfileQuery(baseOptions?: Apollo.QueryHookOptions<GetProfileQuery, GetProfileQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetProfileQuery, GetProfileQueryVariables>(GetProfileDocument, options);
      }
export function useGetProfileLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetProfileQuery, GetProfileQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetProfileQuery, GetProfileQueryVariables>(GetProfileDocument, options);
        }
export type GetProfileQueryHookResult = ReturnType<typeof useGetProfileQuery>;
export type GetProfileLazyQueryHookResult = ReturnType<typeof useGetProfileLazyQuery>;
export type GetProfileQueryResult = Apollo.QueryResult<GetProfileQuery, GetProfileQueryVariables>;
export const GetLanguageDocument = gql`
    query GetLanguage {
  me {
    id
    lang
  }
}
    `;

/**
 * __useGetLanguageQuery__
 *
 * To run a query within a React component, call `useGetLanguageQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetLanguageQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetLanguageQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetLanguageQuery(baseOptions?: Apollo.QueryHookOptions<GetLanguageQuery, GetLanguageQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetLanguageQuery, GetLanguageQueryVariables>(GetLanguageDocument, options);
      }
export function useGetLanguageLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetLanguageQuery, GetLanguageQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetLanguageQuery, GetLanguageQueryVariables>(GetLanguageDocument, options);
        }
export type GetLanguageQueryHookResult = ReturnType<typeof useGetLanguageQuery>;
export type GetLanguageLazyQueryHookResult = ReturnType<typeof useGetLanguageLazyQuery>;
export type GetLanguageQueryResult = Apollo.QueryResult<GetLanguageQuery, GetLanguageQueryVariables>;
export const GetThemeDocument = gql`
    query GetTheme {
  me {
    id
    theme
  }
}
    `;

/**
 * __useGetThemeQuery__
 *
 * To run a query within a React component, call `useGetThemeQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetThemeQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetThemeQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetThemeQuery(baseOptions?: Apollo.QueryHookOptions<GetThemeQuery, GetThemeQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetThemeQuery, GetThemeQueryVariables>(GetThemeDocument, options);
      }
export function useGetThemeLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetThemeQuery, GetThemeQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetThemeQuery, GetThemeQueryVariables>(GetThemeDocument, options);
        }
export type GetThemeQueryHookResult = ReturnType<typeof useGetThemeQuery>;
export type GetThemeLazyQueryHookResult = ReturnType<typeof useGetThemeLazyQuery>;
export type GetThemeQueryResult = Apollo.QueryResult<GetThemeQuery, GetThemeQueryVariables>;
export const UpdateMeDocument = gql`
    mutation UpdateMe($name: String, $email: String, $lang: Lang, $theme: Theme, $password: String, $passwordConfirmation: String) {
  updateMe(
    input: {name: $name, email: $email, lang: $lang, theme: $theme, password: $password, passwordConfirmation: $passwordConfirmation}
  ) {
    me {
      id
      name
      email
      lang
      theme
      myWorkspace {
        id
        name
      }
    }
  }
}
    `;
export type UpdateMeMutationFn = Apollo.MutationFunction<UpdateMeMutation, UpdateMeMutationVariables>;

/**
 * __useUpdateMeMutation__
 *
 * To run a mutation, you first call `useUpdateMeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateMeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateMeMutation, { data, loading, error }] = useUpdateMeMutation({
 *   variables: {
 *      name: // value for 'name'
 *      email: // value for 'email'
 *      lang: // value for 'lang'
 *      theme: // value for 'theme'
 *      password: // value for 'password'
 *      passwordConfirmation: // value for 'passwordConfirmation'
 *   },
 * });
 */
export function useUpdateMeMutation(baseOptions?: Apollo.MutationHookOptions<UpdateMeMutation, UpdateMeMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateMeMutation, UpdateMeMutationVariables>(UpdateMeDocument, options);
      }
export type UpdateMeMutationHookResult = ReturnType<typeof useUpdateMeMutation>;
export type UpdateMeMutationResult = Apollo.MutationResult<UpdateMeMutation>;
export type UpdateMeMutationOptions = Apollo.BaseMutationOptions<UpdateMeMutation, UpdateMeMutationVariables>;
export const DeleteMeDocument = gql`
    mutation DeleteMe($userId: ID!) {
  deleteMe(input: {userId: $userId}) {
    userId
  }
}
    `;
export type DeleteMeMutationFn = Apollo.MutationFunction<DeleteMeMutation, DeleteMeMutationVariables>;

/**
 * __useDeleteMeMutation__
 *
 * To run a mutation, you first call `useDeleteMeMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteMeMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteMeMutation, { data, loading, error }] = useDeleteMeMutation({
 *   variables: {
 *      userId: // value for 'userId'
 *   },
 * });
 */
export function useDeleteMeMutation(baseOptions?: Apollo.MutationHookOptions<DeleteMeMutation, DeleteMeMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteMeMutation, DeleteMeMutationVariables>(DeleteMeDocument, options);
      }
export type DeleteMeMutationHookResult = ReturnType<typeof useDeleteMeMutation>;
export type DeleteMeMutationResult = Apollo.MutationResult<DeleteMeMutation>;
export type DeleteMeMutationOptions = Apollo.BaseMutationOptions<DeleteMeMutation, DeleteMeMutationVariables>;
export const CreateWebhookDocument = gql`
    mutation CreateWebhook($integrationId: ID!, $name: String!, $url: URL!, $active: Boolean!, $trigger: WebhookTriggerInput!, $secret: String!) {
  createWebhook(
    input: {integrationId: $integrationId, name: $name, url: $url, active: $active, trigger: $trigger, secret: $secret}
  ) {
    webhook {
      id
      name
      url
      active
      trigger {
        onItemCreate
        onItemUpdate
        onItemDelete
        onItemPublish
        onItemUnPublish
        onAssetUpload
        onAssetDecompress
        onAssetDelete
      }
      secret
      createdAt
      updatedAt
    }
  }
}
    `;
export type CreateWebhookMutationFn = Apollo.MutationFunction<CreateWebhookMutation, CreateWebhookMutationVariables>;

/**
 * __useCreateWebhookMutation__
 *
 * To run a mutation, you first call `useCreateWebhookMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateWebhookMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createWebhookMutation, { data, loading, error }] = useCreateWebhookMutation({
 *   variables: {
 *      integrationId: // value for 'integrationId'
 *      name: // value for 'name'
 *      url: // value for 'url'
 *      active: // value for 'active'
 *      trigger: // value for 'trigger'
 *      secret: // value for 'secret'
 *   },
 * });
 */
export function useCreateWebhookMutation(baseOptions?: Apollo.MutationHookOptions<CreateWebhookMutation, CreateWebhookMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateWebhookMutation, CreateWebhookMutationVariables>(CreateWebhookDocument, options);
      }
export type CreateWebhookMutationHookResult = ReturnType<typeof useCreateWebhookMutation>;
export type CreateWebhookMutationResult = Apollo.MutationResult<CreateWebhookMutation>;
export type CreateWebhookMutationOptions = Apollo.BaseMutationOptions<CreateWebhookMutation, CreateWebhookMutationVariables>;
export const UpdateWebhookDocument = gql`
    mutation UpdateWebhook($integrationId: ID!, $webhookId: ID!, $name: String!, $url: URL!, $active: Boolean!, $trigger: WebhookTriggerInput!, $secret: String) {
  updateWebhook(
    input: {integrationId: $integrationId, webhookId: $webhookId, name: $name, url: $url, active: $active, trigger: $trigger, secret: $secret}
  ) {
    webhook {
      id
      name
      url
      active
      trigger {
        onItemCreate
        onItemUpdate
        onItemDelete
        onItemPublish
        onItemUnPublish
        onAssetUpload
        onAssetDecompress
        onAssetDelete
      }
      secret
      createdAt
      updatedAt
    }
  }
}
    `;
export type UpdateWebhookMutationFn = Apollo.MutationFunction<UpdateWebhookMutation, UpdateWebhookMutationVariables>;

/**
 * __useUpdateWebhookMutation__
 *
 * To run a mutation, you first call `useUpdateWebhookMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateWebhookMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateWebhookMutation, { data, loading, error }] = useUpdateWebhookMutation({
 *   variables: {
 *      integrationId: // value for 'integrationId'
 *      webhookId: // value for 'webhookId'
 *      name: // value for 'name'
 *      url: // value for 'url'
 *      active: // value for 'active'
 *      trigger: // value for 'trigger'
 *      secret: // value for 'secret'
 *   },
 * });
 */
export function useUpdateWebhookMutation(baseOptions?: Apollo.MutationHookOptions<UpdateWebhookMutation, UpdateWebhookMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateWebhookMutation, UpdateWebhookMutationVariables>(UpdateWebhookDocument, options);
      }
export type UpdateWebhookMutationHookResult = ReturnType<typeof useUpdateWebhookMutation>;
export type UpdateWebhookMutationResult = Apollo.MutationResult<UpdateWebhookMutation>;
export type UpdateWebhookMutationOptions = Apollo.BaseMutationOptions<UpdateWebhookMutation, UpdateWebhookMutationVariables>;
export const DeleteWebhookDocument = gql`
    mutation DeleteWebhook($integrationId: ID!, $webhookId: ID!) {
  deleteWebhook(input: {integrationId: $integrationId, webhookId: $webhookId}) {
    webhookId
  }
}
    `;
export type DeleteWebhookMutationFn = Apollo.MutationFunction<DeleteWebhookMutation, DeleteWebhookMutationVariables>;

/**
 * __useDeleteWebhookMutation__
 *
 * To run a mutation, you first call `useDeleteWebhookMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteWebhookMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteWebhookMutation, { data, loading, error }] = useDeleteWebhookMutation({
 *   variables: {
 *      integrationId: // value for 'integrationId'
 *      webhookId: // value for 'webhookId'
 *   },
 * });
 */
export function useDeleteWebhookMutation(baseOptions?: Apollo.MutationHookOptions<DeleteWebhookMutation, DeleteWebhookMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteWebhookMutation, DeleteWebhookMutationVariables>(DeleteWebhookDocument, options);
      }
export type DeleteWebhookMutationHookResult = ReturnType<typeof useDeleteWebhookMutation>;
export type DeleteWebhookMutationResult = Apollo.MutationResult<DeleteWebhookMutation>;
export type DeleteWebhookMutationOptions = Apollo.BaseMutationOptions<DeleteWebhookMutation, DeleteWebhookMutationVariables>;
export const GetWorkspacesDocument = gql`
    query GetWorkspaces {
  me {
    id
    name
    myWorkspace {
      id
      ...WorkspaceFragment
    }
    workspaces {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;

/**
 * __useGetWorkspacesQuery__
 *
 * To run a query within a React component, call `useGetWorkspacesQuery` and pass it any options that fit your needs.
 * When your component renders, `useGetWorkspacesQuery` returns an object from Apollo Client that contains loading, error, and data properties
 * you can use to render your UI.
 *
 * @param baseOptions options that will be passed into the query, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options;
 *
 * @example
 * const { data, loading, error } = useGetWorkspacesQuery({
 *   variables: {
 *   },
 * });
 */
export function useGetWorkspacesQuery(baseOptions?: Apollo.QueryHookOptions<GetWorkspacesQuery, GetWorkspacesQueryVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useQuery<GetWorkspacesQuery, GetWorkspacesQueryVariables>(GetWorkspacesDocument, options);
      }
export function useGetWorkspacesLazyQuery(baseOptions?: Apollo.LazyQueryHookOptions<GetWorkspacesQuery, GetWorkspacesQueryVariables>) {
          const options = {...defaultOptions, ...baseOptions}
          return Apollo.useLazyQuery<GetWorkspacesQuery, GetWorkspacesQueryVariables>(GetWorkspacesDocument, options);
        }
export type GetWorkspacesQueryHookResult = ReturnType<typeof useGetWorkspacesQuery>;
export type GetWorkspacesLazyQueryHookResult = ReturnType<typeof useGetWorkspacesLazyQuery>;
export type GetWorkspacesQueryResult = Apollo.QueryResult<GetWorkspacesQuery, GetWorkspacesQueryVariables>;
export const UpdateWorkspaceDocument = gql`
    mutation UpdateWorkspace($workspaceId: ID!, $name: String!) {
  updateWorkspace(input: {workspaceId: $workspaceId, name: $name}) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type UpdateWorkspaceMutationFn = Apollo.MutationFunction<UpdateWorkspaceMutation, UpdateWorkspaceMutationVariables>;

/**
 * __useUpdateWorkspaceMutation__
 *
 * To run a mutation, you first call `useUpdateWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateWorkspaceMutation, { data, loading, error }] = useUpdateWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      name: // value for 'name'
 *   },
 * });
 */
export function useUpdateWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<UpdateWorkspaceMutation, UpdateWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateWorkspaceMutation, UpdateWorkspaceMutationVariables>(UpdateWorkspaceDocument, options);
      }
export type UpdateWorkspaceMutationHookResult = ReturnType<typeof useUpdateWorkspaceMutation>;
export type UpdateWorkspaceMutationResult = Apollo.MutationResult<UpdateWorkspaceMutation>;
export type UpdateWorkspaceMutationOptions = Apollo.BaseMutationOptions<UpdateWorkspaceMutation, UpdateWorkspaceMutationVariables>;
export const DeleteWorkspaceDocument = gql`
    mutation DeleteWorkspace($workspaceId: ID!) {
  deleteWorkspace(input: {workspaceId: $workspaceId}) {
    workspaceId
  }
}
    `;
export type DeleteWorkspaceMutationFn = Apollo.MutationFunction<DeleteWorkspaceMutation, DeleteWorkspaceMutationVariables>;

/**
 * __useDeleteWorkspaceMutation__
 *
 * To run a mutation, you first call `useDeleteWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useDeleteWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [deleteWorkspaceMutation, { data, loading, error }] = useDeleteWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *   },
 * });
 */
export function useDeleteWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<DeleteWorkspaceMutation, DeleteWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<DeleteWorkspaceMutation, DeleteWorkspaceMutationVariables>(DeleteWorkspaceDocument, options);
      }
export type DeleteWorkspaceMutationHookResult = ReturnType<typeof useDeleteWorkspaceMutation>;
export type DeleteWorkspaceMutationResult = Apollo.MutationResult<DeleteWorkspaceMutation>;
export type DeleteWorkspaceMutationOptions = Apollo.BaseMutationOptions<DeleteWorkspaceMutation, DeleteWorkspaceMutationVariables>;
export const AddUsersToWorkspaceDocument = gql`
    mutation AddUsersToWorkspace($workspaceId: ID!, $users: [MemberInput!]!) {
  addUsersToWorkspace(input: {workspaceId: $workspaceId, users: $users}) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type AddUsersToWorkspaceMutationFn = Apollo.MutationFunction<AddUsersToWorkspaceMutation, AddUsersToWorkspaceMutationVariables>;

/**
 * __useAddUsersToWorkspaceMutation__
 *
 * To run a mutation, you first call `useAddUsersToWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddUsersToWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addUsersToWorkspaceMutation, { data, loading, error }] = useAddUsersToWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      users: // value for 'users'
 *   },
 * });
 */
export function useAddUsersToWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<AddUsersToWorkspaceMutation, AddUsersToWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddUsersToWorkspaceMutation, AddUsersToWorkspaceMutationVariables>(AddUsersToWorkspaceDocument, options);
      }
export type AddUsersToWorkspaceMutationHookResult = ReturnType<typeof useAddUsersToWorkspaceMutation>;
export type AddUsersToWorkspaceMutationResult = Apollo.MutationResult<AddUsersToWorkspaceMutation>;
export type AddUsersToWorkspaceMutationOptions = Apollo.BaseMutationOptions<AddUsersToWorkspaceMutation, AddUsersToWorkspaceMutationVariables>;
export const UpdateMemberOfWorkspaceDocument = gql`
    mutation UpdateMemberOfWorkspace($workspaceId: ID!, $userId: ID!, $role: Role!) {
  updateUserOfWorkspace(
    input: {workspaceId: $workspaceId, userId: $userId, role: $role}
  ) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type UpdateMemberOfWorkspaceMutationFn = Apollo.MutationFunction<UpdateMemberOfWorkspaceMutation, UpdateMemberOfWorkspaceMutationVariables>;

/**
 * __useUpdateMemberOfWorkspaceMutation__
 *
 * To run a mutation, you first call `useUpdateMemberOfWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateMemberOfWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateMemberOfWorkspaceMutation, { data, loading, error }] = useUpdateMemberOfWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      userId: // value for 'userId'
 *      role: // value for 'role'
 *   },
 * });
 */
export function useUpdateMemberOfWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<UpdateMemberOfWorkspaceMutation, UpdateMemberOfWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateMemberOfWorkspaceMutation, UpdateMemberOfWorkspaceMutationVariables>(UpdateMemberOfWorkspaceDocument, options);
      }
export type UpdateMemberOfWorkspaceMutationHookResult = ReturnType<typeof useUpdateMemberOfWorkspaceMutation>;
export type UpdateMemberOfWorkspaceMutationResult = Apollo.MutationResult<UpdateMemberOfWorkspaceMutation>;
export type UpdateMemberOfWorkspaceMutationOptions = Apollo.BaseMutationOptions<UpdateMemberOfWorkspaceMutation, UpdateMemberOfWorkspaceMutationVariables>;
export const RemoveMemberFromWorkspaceDocument = gql`
    mutation RemoveMemberFromWorkspace($workspaceId: ID!, $userId: ID!) {
  removeUserFromWorkspace(input: {workspaceId: $workspaceId, userId: $userId}) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type RemoveMemberFromWorkspaceMutationFn = Apollo.MutationFunction<RemoveMemberFromWorkspaceMutation, RemoveMemberFromWorkspaceMutationVariables>;

/**
 * __useRemoveMemberFromWorkspaceMutation__
 *
 * To run a mutation, you first call `useRemoveMemberFromWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveMemberFromWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeMemberFromWorkspaceMutation, { data, loading, error }] = useRemoveMemberFromWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      userId: // value for 'userId'
 *   },
 * });
 */
export function useRemoveMemberFromWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<RemoveMemberFromWorkspaceMutation, RemoveMemberFromWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<RemoveMemberFromWorkspaceMutation, RemoveMemberFromWorkspaceMutationVariables>(RemoveMemberFromWorkspaceDocument, options);
      }
export type RemoveMemberFromWorkspaceMutationHookResult = ReturnType<typeof useRemoveMemberFromWorkspaceMutation>;
export type RemoveMemberFromWorkspaceMutationResult = Apollo.MutationResult<RemoveMemberFromWorkspaceMutation>;
export type RemoveMemberFromWorkspaceMutationOptions = Apollo.BaseMutationOptions<RemoveMemberFromWorkspaceMutation, RemoveMemberFromWorkspaceMutationVariables>;
export const AddIntegrationToWorkspaceDocument = gql`
    mutation AddIntegrationToWorkspace($workspaceId: ID!, $integrationId: ID!, $role: Role!) {
  addIntegrationToWorkspace(
    input: {workspaceId: $workspaceId, integrationId: $integrationId, role: $role}
  ) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type AddIntegrationToWorkspaceMutationFn = Apollo.MutationFunction<AddIntegrationToWorkspaceMutation, AddIntegrationToWorkspaceMutationVariables>;

/**
 * __useAddIntegrationToWorkspaceMutation__
 *
 * To run a mutation, you first call `useAddIntegrationToWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useAddIntegrationToWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [addIntegrationToWorkspaceMutation, { data, loading, error }] = useAddIntegrationToWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      integrationId: // value for 'integrationId'
 *      role: // value for 'role'
 *   },
 * });
 */
export function useAddIntegrationToWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<AddIntegrationToWorkspaceMutation, AddIntegrationToWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<AddIntegrationToWorkspaceMutation, AddIntegrationToWorkspaceMutationVariables>(AddIntegrationToWorkspaceDocument, options);
      }
export type AddIntegrationToWorkspaceMutationHookResult = ReturnType<typeof useAddIntegrationToWorkspaceMutation>;
export type AddIntegrationToWorkspaceMutationResult = Apollo.MutationResult<AddIntegrationToWorkspaceMutation>;
export type AddIntegrationToWorkspaceMutationOptions = Apollo.BaseMutationOptions<AddIntegrationToWorkspaceMutation, AddIntegrationToWorkspaceMutationVariables>;
export const UpdateIntegrationOfWorkspaceDocument = gql`
    mutation UpdateIntegrationOfWorkspace($workspaceId: ID!, $integrationId: ID!, $role: Role!) {
  updateIntegrationOfWorkspace(
    input: {workspaceId: $workspaceId, integrationId: $integrationId, role: $role}
  ) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type UpdateIntegrationOfWorkspaceMutationFn = Apollo.MutationFunction<UpdateIntegrationOfWorkspaceMutation, UpdateIntegrationOfWorkspaceMutationVariables>;

/**
 * __useUpdateIntegrationOfWorkspaceMutation__
 *
 * To run a mutation, you first call `useUpdateIntegrationOfWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useUpdateIntegrationOfWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [updateIntegrationOfWorkspaceMutation, { data, loading, error }] = useUpdateIntegrationOfWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      integrationId: // value for 'integrationId'
 *      role: // value for 'role'
 *   },
 * });
 */
export function useUpdateIntegrationOfWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<UpdateIntegrationOfWorkspaceMutation, UpdateIntegrationOfWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<UpdateIntegrationOfWorkspaceMutation, UpdateIntegrationOfWorkspaceMutationVariables>(UpdateIntegrationOfWorkspaceDocument, options);
      }
export type UpdateIntegrationOfWorkspaceMutationHookResult = ReturnType<typeof useUpdateIntegrationOfWorkspaceMutation>;
export type UpdateIntegrationOfWorkspaceMutationResult = Apollo.MutationResult<UpdateIntegrationOfWorkspaceMutation>;
export type UpdateIntegrationOfWorkspaceMutationOptions = Apollo.BaseMutationOptions<UpdateIntegrationOfWorkspaceMutation, UpdateIntegrationOfWorkspaceMutationVariables>;
export const RemoveIntegrationFromWorkspaceDocument = gql`
    mutation RemoveIntegrationFromWorkspace($workspaceId: ID!, $integrationId: ID!) {
  removeIntegrationFromWorkspace(
    input: {workspaceId: $workspaceId, integrationId: $integrationId}
  ) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type RemoveIntegrationFromWorkspaceMutationFn = Apollo.MutationFunction<RemoveIntegrationFromWorkspaceMutation, RemoveIntegrationFromWorkspaceMutationVariables>;

/**
 * __useRemoveIntegrationFromWorkspaceMutation__
 *
 * To run a mutation, you first call `useRemoveIntegrationFromWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useRemoveIntegrationFromWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [removeIntegrationFromWorkspaceMutation, { data, loading, error }] = useRemoveIntegrationFromWorkspaceMutation({
 *   variables: {
 *      workspaceId: // value for 'workspaceId'
 *      integrationId: // value for 'integrationId'
 *   },
 * });
 */
export function useRemoveIntegrationFromWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<RemoveIntegrationFromWorkspaceMutation, RemoveIntegrationFromWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<RemoveIntegrationFromWorkspaceMutation, RemoveIntegrationFromWorkspaceMutationVariables>(RemoveIntegrationFromWorkspaceDocument, options);
      }
export type RemoveIntegrationFromWorkspaceMutationHookResult = ReturnType<typeof useRemoveIntegrationFromWorkspaceMutation>;
export type RemoveIntegrationFromWorkspaceMutationResult = Apollo.MutationResult<RemoveIntegrationFromWorkspaceMutation>;
export type RemoveIntegrationFromWorkspaceMutationOptions = Apollo.BaseMutationOptions<RemoveIntegrationFromWorkspaceMutation, RemoveIntegrationFromWorkspaceMutationVariables>;
export const CreateWorkspaceDocument = gql`
    mutation CreateWorkspace($name: String!) {
  createWorkspace(input: {name: $name}) {
    workspace {
      id
      ...WorkspaceFragment
    }
  }
}
    ${WorkspaceFragmentFragmentDoc}`;
export type CreateWorkspaceMutationFn = Apollo.MutationFunction<CreateWorkspaceMutation, CreateWorkspaceMutationVariables>;

/**
 * __useCreateWorkspaceMutation__
 *
 * To run a mutation, you first call `useCreateWorkspaceMutation` within a React component and pass it any options that fit your needs.
 * When your component renders, `useCreateWorkspaceMutation` returns a tuple that includes:
 * - A mutate function that you can call at any time to execute the mutation
 * - An object with fields that represent the current status of the mutation's execution
 *
 * @param baseOptions options that will be passed into the mutation, supported options are listed on: https://www.apollographql.com/docs/react/api/react-hooks/#options-2;
 *
 * @example
 * const [createWorkspaceMutation, { data, loading, error }] = useCreateWorkspaceMutation({
 *   variables: {
 *      name: // value for 'name'
 *   },
 * });
 */
export function useCreateWorkspaceMutation(baseOptions?: Apollo.MutationHookOptions<CreateWorkspaceMutation, CreateWorkspaceMutationVariables>) {
        const options = {...defaultOptions, ...baseOptions}
        return Apollo.useMutation<CreateWorkspaceMutation, CreateWorkspaceMutationVariables>(CreateWorkspaceDocument, options);
      }
export type CreateWorkspaceMutationHookResult = ReturnType<typeof useCreateWorkspaceMutation>;
export type CreateWorkspaceMutationResult = Apollo.MutationResult<CreateWorkspaceMutation>;
export type CreateWorkspaceMutationOptions = Apollo.BaseMutationOptions<CreateWorkspaceMutation, CreateWorkspaceMutationVariables>;