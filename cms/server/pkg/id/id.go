package id

import "github.com/reearth/reearthx/idx"

type Workspace struct{}
type User struct{}
type Asset struct{}
type Event struct{}

func (Workspace) Type() string { return "workspace" }
func (User) Type() string      { return "user" }
func (Asset) Type() string     { return "asset" }
func (Event) Type() string     { return "event" }

type WorkspaceID = idx.ID[Workspace]
type UserID = idx.ID[User]
type AssetID = idx.ID[Asset]
type EventID = idx.ID[Event]

var NewWorkspaceID = idx.New[Workspace]
var NewUserID = idx.New[User]
var NewAssetID = idx.New[Asset]
var NewEventID = idx.New[Event]

var MustWorkspaceID = idx.Must[Workspace]
var MustUserID = idx.Must[User]
var MustAssetID = idx.Must[Asset]
var MustEventID = idx.Must[Event]

var WorkspaceIDFrom = idx.From[Workspace]
var UserIDFrom = idx.From[User]
var AssetIDFrom = idx.From[Asset]
var EventIDFrom = idx.From[Event]

var WorkspaceIDFromRef = idx.FromRef[Workspace]
var UserIDFromRef = idx.FromRef[User]
var AssetIDFromRef = idx.FromRef[Asset]
var EventIDFromRef = idx.FromRef[Event]

type WorkspaceIDList = idx.List[Workspace]
type UserIDList = idx.List[User]
type AssetIDList = idx.List[Asset]

var WorkspaceIDListFrom = idx.ListFrom[Workspace]
var UserIDListFrom = idx.ListFrom[User]
var AssetIDListFrom = idx.ListFrom[Asset]

type WorkspaceIDSet = idx.Set[Workspace]
type UserIDSet = idx.Set[User]
type AssetIDSet = idx.Set[Asset]

var NewWorkspaceIDSet = idx.NewSet[Workspace]
var NewUserIDSet = idx.NewSet[User]
var NewAssetIDSet = idx.NewSet[Asset]

type Project struct{}

func (Project) Type() string { return "project" }

type ProjectID = idx.ID[Project]
type ProjectIDList = idx.List[Project]

var MustProjectID = idx.Must[Project]
var NewProjectID = idx.New[Project]
var ProjectIDFrom = idx.From[Project]
var ProjectIDFromRef = idx.FromRef[Project]
var ProjectIDListFrom = idx.ListFrom[Project]

type Model struct{}

func (Model) Type() string { return "model" }

type ModelID = idx.ID[Model]
type ModelIDList = idx.List[Model]

var MustModelID = idx.Must[Model]
var NewModelID = idx.New[Model]
var ModelIDFrom = idx.From[Model]
var ModelIDFromRef = idx.FromRef[Model]
var ModelIDListFrom = idx.ListFrom[Model]

type Field struct{}

func (Field) Type() string { return "field" }

type FieldID = idx.ID[Field]
type FieldIDList = idx.List[Field]

var MustFieldID = idx.Must[Field]
var NewFieldID = idx.New[Field]
var FieldIDFrom = idx.From[Field]
var FieldIDFromRef = idx.FromRef[Field]
var FieldIDListFrom = idx.ListFrom[Field]

type Schema struct{}

func (Schema) Type() string { return "schema" }

type SchemaID = idx.ID[Schema]
type SchemaIDList = idx.List[Schema]

var MustSchemaID = idx.Must[Schema]
var NewSchemaID = idx.New[Schema]
var SchemaIDFrom = idx.From[Schema]
var SchemaIDFromRef = idx.FromRef[Schema]
var SchemaIDListFrom = idx.ListFrom[Schema]

type Thread struct{}

func (Thread) Type() string { return "thread" }

type ThreadID = idx.ID[Thread]
type ThreadIDList = idx.List[Thread]

var NewThreadID = idx.New[Thread]
var MustThreadID = idx.Must[Thread]
var ThreadIDFrom = idx.From[Thread]
var ThreadIDFromRef = idx.FromRef[Thread]

type Comment struct{}

func (Comment) Type() string { return "comment" }

type CommentID = idx.ID[Comment]
type CommentIDList = idx.List[Comment]

var NewCommentID = idx.New[Comment]
var MustCommentID = idx.Must[Comment]
var CommentIDFrom = idx.From[Comment]
var CommentIDFromRef = idx.FromRef[Comment]

type Item struct{}

func (Item) Type() string { return "item" }

type ItemID = idx.ID[Item]
type ItemIDList = idx.List[Item]

var MustItemID = idx.Must[Item]
var NewItemID = idx.New[Item]
var ItemIDFrom = idx.From[Item]
var ItemIDFromRef = idx.FromRef[Item]
var ItemIDListFrom = idx.ListFrom[Item]

type Integration struct{}

func (Integration) Type() string { return "integration" }

type IntegrationID = idx.ID[Integration]
type IntegrationIDList = idx.List[Integration]

var MustIntegrationID = idx.Must[Integration]
var NewIntegrationID = idx.New[Integration]
var IntegrationIDFrom = idx.From[Integration]
var IntegrationIDFromRef = idx.FromRef[Integration]
var IntegrationIDListFrom = idx.ListFrom[Integration]

type Webhook struct{}

func (Webhook) Type() string { return "webhook" }

type WebhookID = idx.ID[Webhook]
type WebhookIDList = idx.List[Webhook]

var MustWebhookID = idx.Must[Webhook]
var NewWebhookID = idx.New[Webhook]
var WebhookIDFrom = idx.From[Webhook]
var WebhookIDFromRef = idx.FromRef[Webhook]
var WebhookIDListFrom = idx.ListFrom[Webhook]

type Task struct{}

func (Task) Type() string { return "task" }

type TaskID = idx.ID[Task]

var NewTaskID = idx.New[Task]
var MustTaskID = idx.Must[Task]
var TaskIDFrom = idx.From[Task]
var TaskIDFromRef = idx.FromRef[Task]

type TaskIDList = idx.List[Task]

var TaskIDListFrom = idx.ListFrom[Task]

type TaskIDSet = idx.Set[Task]

var NewTaskIDSet = idx.NewSet[Task]

type Request struct{}

func (Request) Type() string { return "request" }

type RequestID = idx.ID[Request]
type RequestIDList = idx.List[Request]

var NewRequestID = idx.New[Request]
var MustRequestID = idx.Must[Request]
var RequestIDFrom = idx.From[Request]
var RequestIDFromRef = idx.FromRef[Request]
