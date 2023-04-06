import { CurrentLocation } from "@web/extensions/geolocation/types";
import { postMsg } from "@web/extensions/geolocation/utils";
import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";

const handleFlyToCurrentLocation = () => {
  if (!navigator.geolocation) return;

  navigator.geolocation.getCurrentPosition(
    function (position) {
      const currentLocation: CurrentLocation = {
        latitude: position.coords.latitude,
        longitude: position.coords.longitude,
        altitude: position.coords.altitude ?? 5000,
      };
      postMsg({ action: "flyTo", payload: { currentLocation } });
    },
    function (error) {
      console.error("Error Code = " + error.code + " - " + error.message);
    },
  );
};

const GeolocationWrapper: React.FC = () => (
  <Wrapper>
    <Icon icon="bullseye" size={20} color="#262626" onClick={handleFlyToCurrentLocation} />
  </Wrapper>
);

export default GeolocationWrapper;

const Wrapper = styled.div`
  display: flex;
  align-items: center;
  justify-content: center;
  flex-shrink: 0;
  cursor: pointer;
  width: 32px;
  height: 32px;
  background: #ffffff;
  border-radius: 16px;
`;
