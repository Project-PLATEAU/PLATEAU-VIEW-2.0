import { PostMessageProps } from "@web/extensions/location/types";

export function postMsg({ action, payload }: PostMessageProps) {
  parent.postMessage({ action, payload });
}
