import { useState } from "react";

export default () => {
  const [svgRender, setSvgRender] = useState<boolean>(true);

  const handleCodeSourceClick = () => {
    setSvgRender(false);
  };

  const handleRenderClick = () => {
    setSvgRender(true);
  };

  return { svgRender, handleCodeSourceClick, handleRenderClick };
};
