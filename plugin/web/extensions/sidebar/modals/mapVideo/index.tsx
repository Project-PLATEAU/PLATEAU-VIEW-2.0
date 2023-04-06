import { postMsg } from "@web/extensions/sidebar/utils";
import { Icon } from "@web/sharedComponents";
import Video from "@web/sharedComponents/Video";
import { styled } from "@web/theme";
import { useCallback } from "react";

const MapVideo: React.FC = () => {
  const handleClose = useCallback(() => {
    postMsg({ action: "modalClose" });
  }, []);

  return (
    <div>
      <CloseButton>
        <Icon size={32} icon="close" onClick={handleClose} />
      </CloseButton>
      <Video width="560" height="315" src="https://www.youtube.com/embed/pY2dM-eG5mA" />
    </div>
  );
};
export default MapVideo;

const CloseButton = styled.button`
  display: flex;
  justify-content: center;
  align-items: center;
  position: absolute;
  right: 0;
  height: 32px;
  width: 32px;
  border: none;
  background: #00bebe;
  color: white;
  cursor: pointer;
`;
