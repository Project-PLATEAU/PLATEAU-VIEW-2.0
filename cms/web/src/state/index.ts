import { atom, useAtom } from "jotai";
import { atomWithStorage } from "jotai/utils";

import { PublicScope } from "@reearth-cms/components/molecules/Accessibility";
import { Model } from "@reearth-cms/components/molecules/Schema/types";

export type Workspace = {
  id: string;
  name: string;
  members?: Array<any>;
  assets?: any;
  projects?: any;
  personal?: boolean;
};

const workspace = atom<Workspace | undefined>(undefined);
export const useWorkspace = () => useAtom(workspace);

const workspaceId = atomWithStorage<string | undefined>("workspaceId", undefined);
export const useWorkspaceId = () => useAtom(workspaceId);

const userId = atomWithStorage<string | undefined>("userId", undefined);
export const useUserId = () => useAtom(userId);

export type Project = {
  id: string;
  name: string;
  description?: string;
  alias: string;
  scope?: PublicScope;
  assetPublic?: boolean;
};

const project = atom<Project | undefined>(undefined);
export const useProject = () => useAtom(project);

const model = atom<Model | undefined>(undefined);
export const useModel = () => useAtom(model);
