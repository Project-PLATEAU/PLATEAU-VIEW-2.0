export type ActionType =
  | "resize"
  | "getViewport"
  | "sceneCapture"
  | "sceneView"
  | "sceneRecapture"
  | "sceneEdit"
  | "sceneSave"
  | "sceneEditorClose"
  | "storyShare"
  | "storySaveData"
  | "storyCancelPlay";

export type CommunicationActionType = "storyEdit" | "storySave" | "storyDelete" | "storyPlay";

export type PostMessageProps = { action: ActionType; payload?: any };

export type Story = {
  id?: string;
  title?: string;
  scenes: Scene[];
};

export type Scene = {
  id: string;
  title: string;
  description: string;
  camera: Camera | undefined;
};

// Reearth types
export type Camera = {
  lat: number;
  lng: number;
  height: number;
  heading: number;
  pitch: number;
  roll: number;
  fov: number;
};

export type Viewport = {
  width: number;
  height: number;
  isMobile: boolean;
};

export type PluginExtensionInstance = {
  id: string;
  pluginId: string;
  name: string;
  extensionId: string;
  extensionType: "widget" | "block";
};

// Communications
export type PluginMessage = {
  data: StoryEdit | StoryDelete | StoryPlay | StoryCancelPlay;
  sender: string;
};

// sidebar -> storytelling
export type StoryEdit = {
  action: "storyEdit";
  payload: {
    id: string;
    dataID?: string;
    scenes: string;
    title?: string;
  };
};

export type StoryEditFinish = {
  action: "storyEditFinish";
  payload: {
    id: string;
  };
};

export type StoryDelete = {
  action: "storyDelete";
  payload: {
    id: string;
  };
};

export type StoryPlay = {
  action: "storyPlay";
  payload: {
    id: string;
    scenes: string;
    title?: string;
  };
};

// sidebar -> storytelling
// storytelling -> sidebar
export type StoryCancelPlay = {
  action: "storyCancelPlay";
  payload: {
    id: string;
  };
};

// storytelling -> sidebar
export type StoryShare = {
  action: "storyShare";
};

export type StorySaveData = {
  action: "storySaveData";
  payload: {
    id: string;
    scenes: string;
  };
};
