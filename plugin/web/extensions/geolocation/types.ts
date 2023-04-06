export type CurrentLocation = {
  latitude: number;
  longitude: number;
  altitude: number;
};

type actionType = "flyTo";

export type PostMessageProps = { action: actionType; payload?: any };
