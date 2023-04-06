import { PostMessageProps } from "@web/extensions/geolocation/types";

export function postMsg({ action, payload }: PostMessageProps) {
  parent.postMessage(
    {
      action,
      payload,
    },
    "*",
  );
}
