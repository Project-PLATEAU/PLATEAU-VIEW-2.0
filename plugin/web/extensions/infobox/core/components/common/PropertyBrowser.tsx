import { cesium3DTilesAppearanceKeys } from "@web/extensions/infobox/core/utils";
import type { Properties, Field } from "@web/extensions/infobox/types";
import { styled } from "@web/theme";
import { useState, useEffect } from "react";
import ReactJson from "react-json-view";

type DisplayItem = {
  path: string;
  title: string;
  value?: any;
};

type Props = {
  properties?: Properties;
  fields: Field[];
  commonProperties: string[];
  attributesKey?: string;
  attributesName?: string;
};

const PropertyBrowser: React.FC<Props> = ({
  properties,
  fields,
  commonProperties,
  attributesKey,
  attributesName,
}) => {
  const [displayList, setDisplayList] = useState<DisplayItem[]>([]);

  useEffect(() => {
    const items: DisplayItem[] = [];

    if (!properties) {
      setDisplayList([]);
      return;
    }

    // Part I: common properties
    // ------------------------------
    fields.forEach(f => {
      if (f.visible && properties?.[f.path] !== undefined) {
        items.push({
          path: f.path,
          title: f.title ? f.title : f.path,
          value: properties?.[f.path],
        });
      }
    });

    // Part II: individual properties
    // ------------------------------
    // !NOTE: currently appearance override properties
    // are mixed with original properties
    const individualProperties = Object.keys(properties)
      .filter(
        k =>
          !commonProperties.includes(k) &&
          (!attributesKey || k !== attributesKey) &&
          !cesium3DTilesAppearanceKeys.includes(k) &&
          typeof properties[k] !== "object",
      )
      .reduce((obj, key) => {
        return {
          ...obj,
          [key]: properties[key],
        };
      }, {});

    Object.entries(individualProperties).forEach(([key, value]) => {
      items.push({
        path: key,
        title: key,
        value: value,
      });
    });

    setDisplayList(items);
  }, [fields, properties, commonProperties, attributesKey]);

  return (
    <Wrapper>
      {displayList.map(field => (
        <PropertyItem key={field.path}>
          <Title>{field.title}</Title>
          <Value>{field.value}</Value>
        </PropertyItem>
      ))}
      {attributesKey && (
        <AttributesWrapper>
          <ReactJson
            src={properties?.[attributesKey]}
            displayDataTypes={false}
            enableClipboard={false}
            displayObjectSize={false}
            quotesOnKeys={false}
            indentWidth={2}
            name={attributesName}
          />
        </AttributesWrapper>
      )}
    </Wrapper>
  );
};

const Wrapper = styled.div`
  padding: 12px 16px;
  background-color: #fff;
`;

const PropertyItem = styled.div`
  display: flex;
  align-items: flex-start;
  min-height: 32px;
  padding: 12px 0 4px;
  gap: 12px;
  border-bottom: 1px solid #d9d9d9;
  font-size: 14px;
`;

const Title = styled.div`
  width: 50%;
`;

const Value = styled.div`
  width: 50%;
  word-break: break-all;
`;

const AttributesWrapper = styled.div`
  padding: 8px 0;

  * {
    color: #000 !important;
    font-family: "Roboto", sans-serif;
  }
`;

export default PropertyBrowser;
