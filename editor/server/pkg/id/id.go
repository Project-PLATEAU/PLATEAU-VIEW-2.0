package id

import "github.com/reearth/reearthx/idx"

type Asset struct{}
type AuthRequest struct{}
type Dataset struct{}
type DatasetField struct{}
type DatasetSchema struct{}
type Cluster struct{}
type InfoboxField struct{}
type Layer struct{}
type PluginExtension struct{}
type Project struct{}
type Property struct{}
type PropertyItem struct{}
type PropertyField struct{}
type PropertySchemaGroup struct{}
type Scene struct{}
type Tag struct{}
type Workspace struct{}
type User struct{}
type Widget struct{}

func (Asset) Type() string               { return "asset" }
func (AuthRequest) Type() string         { return "authRequest" }
func (Dataset) Type() string             { return "dataset" }
func (DatasetField) Type() string        { return "datasetField" }
func (DatasetSchema) Type() string       { return "datasetSchema" }
func (Cluster) Type() string             { return "cluster" }
func (InfoboxField) Type() string        { return "infoboxField" }
func (Layer) Type() string               { return "layer" }
func (PluginExtension) Type() string     { return "pluginExtension" }
func (Project) Type() string             { return "project" }
func (Property) Type() string            { return "property" }
func (PropertyItem) Type() string        { return "propertyItem" }
func (PropertyField) Type() string       { return "propertyField" }
func (PropertySchemaGroup) Type() string { return "propertySchemaGroup" }
func (Scene) Type() string               { return "scene" }
func (Tag) Type() string                 { return "tag" }
func (Workspace) Type() string           { return "workspace" }
func (User) Type() string                { return "user" }
func (Widget) Type() string              { return "widget" }

type AssetID = idx.ID[Asset]
type AuthRequestID = idx.ID[AuthRequest]
type DatasetID = idx.ID[Dataset]
type DatasetFieldID = idx.ID[DatasetField]
type DatasetSchemaID = idx.ID[DatasetSchema]
type ClusterID = idx.ID[Cluster]
type InfoboxFieldID = idx.ID[InfoboxField]
type LayerID = idx.ID[Layer]
type ProjectID = idx.ID[Project]
type PropertyID = idx.ID[Property]
type PropertyItemID = idx.ID[PropertyItem]
type SceneID = idx.ID[Scene]
type TagID = idx.ID[Tag]
type WorkspaceID = idx.ID[Workspace]
type UserID = idx.ID[User]
type WidgetID = idx.ID[Widget]

type PluginExtensionID = idx.StringID[PluginExtension]
type PropertySchemaGroupID = idx.StringID[PropertySchemaGroup]
type PropertyFieldID = idx.StringID[PropertyField]

var NewAssetID = idx.New[Asset]
var NewAuthRequestID = idx.New[AuthRequest]
var NewDatasetID = idx.New[Dataset]
var NewDatasetFieldID = idx.New[DatasetField]
var NewDatasetSchemaID = idx.New[DatasetSchema]
var NewClusterID = idx.New[Cluster]
var NewInfoboxFieldID = idx.New[InfoboxField]
var NewLayerID = idx.New[Layer]
var NewProjectID = idx.New[Project]
var NewPropertyID = idx.New[Property]
var NewPropertyItemID = idx.New[PropertyItem]
var NewSceneID = idx.New[Scene]
var NewTagID = idx.New[Tag]
var NewWorkspaceID = idx.New[Workspace]
var NewUserID = idx.New[User]
var NewWidgetID = idx.New[Widget]

var MustAssetID = idx.Must[Asset]
var MustAuthRequestID = idx.Must[AuthRequest]
var MustDatasetID = idx.Must[Dataset]
var MustDatasetFieldID = idx.Must[DatasetField]
var MustDatasetSchemaID = idx.Must[DatasetSchema]
var MustClusterID = idx.Must[Cluster]
var MustInfoboxFieldID = idx.Must[InfoboxField]
var MustLayerID = idx.Must[Layer]
var MustProjectID = idx.Must[Project]
var MustPropertyID = idx.Must[Property]
var MustPropertyItemID = idx.Must[PropertyItem]
var MustSceneID = idx.Must[Scene]
var MustTagID = idx.Must[Tag]
var MustWorkspaceID = idx.Must[Workspace]
var MustUserID = idx.Must[User]
var MustWidgetID = idx.Must[Widget]

var AssetIDFrom = idx.From[Asset]
var AuthRequestIDFrom = idx.From[AuthRequest]
var DatasetIDFrom = idx.From[Dataset]
var DatasetFieldIDFrom = idx.From[DatasetField]
var DatasetSchemaIDFrom = idx.From[DatasetSchema]
var ClusterIDFrom = idx.From[Cluster]
var InfoboxFieldIDFrom = idx.From[InfoboxField]
var LayerIDFrom = idx.From[Layer]
var ProjectIDFrom = idx.From[Project]
var PropertyIDFrom = idx.From[Property]
var PropertyItemIDFrom = idx.From[PropertyItem]
var SceneIDFrom = idx.From[Scene]
var TagIDFrom = idx.From[Tag]
var WorkspaceIDFrom = idx.From[Workspace]
var UserIDFrom = idx.From[User]
var WidgetIDFrom = idx.From[Widget]

var AssetIDFromRef = idx.FromRef[Asset]
var AuthRequestIDFromRef = idx.FromRef[AuthRequest]
var DatasetIDFromRef = idx.FromRef[Dataset]
var DatasetFieldIDFromRef = idx.FromRef[DatasetField]
var DatasetSchemaIDFromRef = idx.FromRef[DatasetSchema]
var ClusterIDFromRef = idx.FromRef[Cluster]
var InfoboxFieldIDFromRef = idx.FromRef[InfoboxField]
var LayerIDFromRef = idx.FromRef[Layer]
var ProjectIDFromRef = idx.FromRef[Project]
var PropertyIDFromRef = idx.FromRef[Property]
var PropertyItemIDFromRef = idx.FromRef[PropertyItem]
var SceneIDFromRef = idx.FromRef[Scene]
var TagIDFromRef = idx.FromRef[Tag]
var WorkspaceIDFromRef = idx.FromRef[Workspace]
var UserIDFromRef = idx.FromRef[User]
var WidgetIDFromRef = idx.FromRef[Widget]

var PluginExtensionIDFromRef = idx.StringIDFromRef[PluginExtension]
var PropertyFieldIDFromRef = idx.StringIDFromRef[PropertyField]
var PropertySchemaGroupIDFromRef = idx.StringIDFromRef[PropertySchemaGroup]

type AssetIDList = idx.List[Asset]
type AuthRequestIDList = idx.List[AuthRequest]
type DatasetIDList = idx.List[Dataset]
type DatasetFieldIDList = idx.List[DatasetField]
type DatasetSchemaIDList = idx.List[DatasetSchema]
type ClusterIDList = idx.List[Cluster]
type InfoboxFieldIDList = idx.List[InfoboxField]
type LayerIDList = idx.List[Layer]
type ProjectIDList = idx.List[Project]
type PropertyIDList = idx.List[Property]
type PropertyItemIDList = idx.List[PropertyItem]
type SceneIDList = idx.List[Scene]
type TagIDList = idx.List[Tag]
type WorkspaceIDList = idx.List[Workspace]
type UserIDList = idx.List[User]
type WidgetIDList = idx.List[Widget]

var AssetIDListFrom = idx.ListFrom[Asset]
var AuthRequestIDListFrom = idx.ListFrom[AuthRequest]
var DatasetIDListFrom = idx.ListFrom[Dataset]
var DatasetFieldIDListFrom = idx.ListFrom[DatasetField]
var DatasetSchemaIDListFrom = idx.ListFrom[DatasetSchema]
var ClusterIDListFrom = idx.ListFrom[Cluster]
var InfoboxFieldIDListFrom = idx.ListFrom[InfoboxField]
var LayerIDListFrom = idx.ListFrom[Layer]
var ProjectIDListFrom = idx.ListFrom[Project]
var PropertyIDListFrom = idx.ListFrom[Property]
var PropertyItemIDListFrom = idx.ListFrom[PropertyItem]
var SceneIDListFrom = idx.ListFrom[Scene]
var TagIDListFrom = idx.ListFrom[Tag]
var WorkspaceIDListFrom = idx.ListFrom[Workspace]
var UserIDListFrom = idx.ListFrom[User]
var WidgetIDListFrom = idx.ListFrom[Widget]

type AssetIDSet = idx.Set[Asset]
type AuthRequestIDSet = idx.Set[AuthRequest]
type DatasetIDSet = idx.Set[Dataset]
type DatasetFieldIDSet = idx.Set[DatasetField]
type DatasetSchemaIDSet = idx.Set[DatasetSchema]
type ClusterIDSet = idx.Set[Cluster]
type InfoboxFieldIDSet = idx.Set[InfoboxField]
type LayerIDSet = idx.Set[Layer]
type ProjectIDSet = idx.Set[Project]
type PropertyIDSet = idx.Set[Property]
type PropertyItemIDSet = idx.Set[PropertyItem]
type SceneIDSet = idx.Set[Scene]
type TagIDSet = idx.Set[Tag]
type WorkspaceIDSet = idx.Set[Workspace]
type UserIDSet = idx.Set[User]
type WidgetIDSet = idx.Set[Widget]

var NewAssetIDSet = idx.NewSet[Asset]
var NewAuthRequestIDSet = idx.NewSet[AuthRequest]
var NewDatasetIDSet = idx.NewSet[Dataset]
var NewDatasetFieldIDSet = idx.NewSet[DatasetField]
var NewDatasetSchemaIDSet = idx.NewSet[DatasetSchema]
var NewClusterIDSet = idx.NewSet[Cluster]
var NewInfoboxFieldIDSet = idx.NewSet[InfoboxField]
var NewLayerIDSet = idx.NewSet[Layer]
var NewProjectIDSet = idx.NewSet[Project]
var NewPropertyIDSet = idx.NewSet[Property]
var NewPropertyItemIDSet = idx.NewSet[PropertyItem]
var NewSceneIDSet = idx.NewSet[Scene]
var NewTagIDSet = idx.NewSet[Tag]
var NewWorkspaceIDSet = idx.NewSet[Workspace]
var NewUserIDSet = idx.NewSet[User]
var NewWidgetIDSet = idx.NewSet[Widget]
