import { PostMessageProps } from "@web/extensions/geolocation/types";

import html from "../dist/web/geolocation/core/index.html?raw";

const reearth = (globalThis as any).reearth;

reearth.ui.show(html);

reearth.on("message", ({ action, payload }: PostMessageProps) => {
  if (action === "flyTo") {
    reearth.layers.add({
      extensionId: "marker",
      isVisible: true,
      title: "myLocation",
      property: {
        default: {
          location: {
            lat: payload.currentLocation.latitude,
            lng: payload.currentLocation.longitude,
          },
          style: "point",
          pointColor: "#12BDE2",
          pointOutlineWidth: 4,
          pointOutlineColor: "#FFFFFF",
        },
      },
      infobox: {
        blocks: [
          {
            extensionId: "dlblock",
            pluginId: "reearth",
            property: {
              items: [
                {
                  item_title: "Lat",
                  item_datatype: "number",
                  item_datanum: payload.currentLocation.latitude,
                },
                {
                  item_title: "Lng",
                  item_datatype: "number",
                  item_datanum: payload.currentLocation.longitude,
                },
              ],
            },
          },
        ],
        property: {
          default: {
            title: "My Location",
            infoboxPaddingLeft: 24,
          },
        },
      },
    });

    reearth.camera.flyTo(
      {
        lat: payload.currentLocation.latitude,
        lng: payload.currentLocation.longitude,
        height: payload.currentLocation.altitude,
        heading: reearth.camera.position?.heading ?? 0,
        pitch: -Math.PI / 2,
        roll: 0,
        fov: reearth.camera.position?.fov ?? 0.75,
      },
      {
        duration: 2,
      },
    );
  }
});
