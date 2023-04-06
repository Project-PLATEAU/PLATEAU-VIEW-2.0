import { useCallback } from "react";

import { postMsg } from "../utils";

export default () => {
  const onPedestrian = useCallback(() => {
    postMsg("pedestrianShow");
  }, []);

  return {
    onPedestrian,
  };
};
