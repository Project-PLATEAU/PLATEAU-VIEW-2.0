import { DataCatalogItem } from "@web/extensions/sidebar/core/types";
import { generateID } from "@web/extensions/sidebar/utils";
import { fieldGroups } from "@web/extensions/sidebar/utils/fieldGroups";
import { useMemo } from "react";

import { fieldName } from "./Fields/types";

type FieldDropdownItem = {
  [key: string]: { name: string; onClick: (property: any) => void };
};

export default ({
  dataset,
  onFieldAdd,
}: {
  dataset?: DataCatalogItem;
  onFieldAdd: (property: any) => ({ key }: { key: string }) => void;
}) => {
  const generalFields: FieldDropdownItem = useMemo(() => {
    return {
      idealZoom: {
        name: fieldName["idealZoom"],
        onClick: onFieldAdd({
          position: {
            lng: 0,
            lat: 0,
            height: 0,
            pitch: 0,
            heading: 0,
            roll: 0,
          },
        }),
      },
      description: {
        name: fieldName["description"],
        onClick: onFieldAdd({}),
      },
      legend: {
        name: fieldName["legend"],
        onClick: onFieldAdd({
          style: "square",
          items: [{ title: "新しいアイテム", color: "#00bebe" }],
        }),
      },
      legendGradient: {
        name: fieldName["legendGradient"],
        onClick: onFieldAdd({
          style: "square",
          min: 0,
          max: 0,
          step: 0,
        }),
      },
      realtime: {
        name: fieldName["realtime"],
        onClick: onFieldAdd({ updateInterval: 30, userSettings: {} }),
      },
      timeline: {
        name: fieldName["timeline"],
        onClick: onFieldAdd({ timeFieldName: "", userSettings: {} }),
      },
      currentTime: {
        name: fieldName["currentTime"],
        onClick: onFieldAdd({
          currentDate: "",
          currentTime: "",
          startDate: "",
          startTime: "",
          stopDate: "",
          stopTime: "",
        }),
      },
      styleCode: {
        name: fieldName["styleCode"],
        onClick: onFieldAdd({ src: " " }),
      },
      story: {
        name: fieldName["story"],
        onClick: onFieldAdd({}),
      },
      buttonLink: {
        name: fieldName["buttonLink"],
        onClick: onFieldAdd({}),
      },
      switchGroup: {
        name: fieldName["switchGroup"],
        onClick: onFieldAdd({
          title: "Switch Group",
          groups: [
            {
              id: generateID(),
              title: "新グループ1",
              fieldGroupID: fieldGroups[0].id,
              userSettings: {},
            },
          ],
        }),
      },
      switchDataset: {
        name: fieldName["switchDataset"],
        onClick: onFieldAdd({
          userSettings: {},
          cleanseOverride: dataset?.config?.data?.[0].url
            ? {
                data: {
                  url: dataset.config.data[0].url,
                  time: {
                    updateClockOnLoad: false,
                  },
                },
              }
            : undefined,
        }),
      },
      switchVisibility: {
        name: fieldName["switchVisibility"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      template: {
        name: fieldName["template"],
        onClick: onFieldAdd({}),
      },
      eventField: {
        name: fieldName["eventField"],
        onClick: onFieldAdd({
          eventType: "select",
          triggerEvent: "openUrl",
          urlType: "manual",
        }),
      },
      infoboxStyle: {
        name: fieldName["infoboxStyle"],
        onClick: onFieldAdd({
          displayStyle: null,
        }),
      },
      heightReference: {
        name: fieldName["heightReference"],
        onClick: onFieldAdd({
          heightReferenceType: "clamp",
        }),
      },
    };
  }, [dataset?.config?.data, onFieldAdd]);

  const pointFields: FieldDropdownItem = useMemo(() => {
    return {
      pointColor: {
        name: fieldName["pointColor"],
        onClick: onFieldAdd({}),
      },
      pointColorGradient: {
        name: fieldName["pointColorGradient"],
        onClick: onFieldAdd({
          min: 0,
          max: 0,
          step: 0,
        }),
      },
      pointSize: {
        name: fieldName["pointSize"],
        onClick: onFieldAdd({}),
      },
      pointIcon: {
        name: fieldName["pointIcon"],
        onClick: onFieldAdd({
          size: 1,
        }),
      },
      pointLabel: {
        name: fieldName["pointLabel"],
        onClick: onFieldAdd({}),
      },
      pointModel: {
        name: fieldName["pointModel"],
        onClick: onFieldAdd({
          scale: 1,
        }),
      },
      pointStroke: {
        name: fieldName["pointStroke"],
        onClick: onFieldAdd({}),
      },
      pointCSV: {
        name: fieldName["pointCSV"],
        onClick: onFieldAdd({}),
      },
    };
  }, [onFieldAdd]);

  const polylineFields: FieldDropdownItem = useMemo(() => {
    return {
      polylineColor: {
        name: fieldName["polylineColor"],
        onClick: onFieldAdd({}),
      },
      // polylineColorGradient: {
      //   name: fieldName["polylineColorGradient"],
      //   onClick: onFieldAdd({}),
      // },
      polylineStrokeWeight: {
        name: fieldName["polylineStrokeWeight"],
        onClick: onFieldAdd({}),
      },
      polylineClassificationType: {
        name: fieldName["polylineClassificationType"],
        onClick: onFieldAdd({}),
      },
    };
  }, [onFieldAdd]);

  const polygonFields: FieldDropdownItem = useMemo(() => {
    return {
      polygonColor: {
        name: fieldName["polygonColor"],
        onClick: onFieldAdd({}),
      },
      // polygonColorGradient: {
      //   name: fieldName["polygonColorGradient"],
      //   onClick: ({ key }) => console.log("do something: ", key),
      // },
      polygonStroke: {
        name: fieldName["polygonStroke"],
        onClick: onFieldAdd({}),
      },
      polygonClassificationType: {
        name: fieldName["polygonClassificationType"],
        onClick: onFieldAdd({}),
      },
    };
  }, [onFieldAdd]);

  const ThreeDTileFields: FieldDropdownItem = useMemo(() => {
    return {
      buildingColor: {
        name: fieldName["buildingColor"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      buildingFilter: {
        name: fieldName["buildingFilter"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      buildingShadow: {
        name: fieldName["buildingShadow"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      buildingTransparency: {
        name: fieldName["buildingTransparency"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      clipping: {
        name: fieldName["clipping"],
        onClick: onFieldAdd({
          userSettings: {
            enabled: false,
            show: false,
            aboveGroundOnly: false,
            direction: "inside",
          },
        }),
      },
      floodColor: {
        name: fieldName["floodColor"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
      floodFilter: {
        name: fieldName["floodFilter"],
        onClick: onFieldAdd({ userSettings: {} }),
      },
    };
  }, [onFieldAdd]);

  const fieldComponentsList: { [key: string]: { name: string; fields: FieldDropdownItem } } =
    useMemo(() => {
      return {
        general: { name: "一般", fields: generalFields },
        point: { name: "ポイント", fields: pointFields },
        polyline: { name: "ポリライン", fields: polylineFields },
        polygon: { name: "ポリゴン", fields: polygonFields },
        // "3d-model": { name: "3Dモデル", fields: ThreeDModelFields },
        "3d-tile": { name: "3Dタイル", fields: ThreeDTileFields },
      };
    }, [generalFields, pointFields, polygonFields, polylineFields, ThreeDTileFields]);

  return fieldComponentsList;
};
