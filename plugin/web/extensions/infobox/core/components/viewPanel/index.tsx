import type { Properties, Field } from "@web/extensions/infobox/types";
import { styled } from "@web/theme";

import PropertyBrowser from "../common/PropertyBrowser";

type Props = {
  name: string;
  properties?: Properties;
  fields: Field[];
  commonProperties: string[];
  attributesKey?: string;
  attributesName?: string;
};

const ViewPanel: React.FC<Props> = ({
  name,
  properties,
  fields,
  commonProperties,
  attributesKey,
  attributesName,
}) => {
  return (
    <>
      <Header>
        <Title>{name}</Title>
      </Header>
      <PropertyBrowser
        properties={properties}
        fields={fields}
        commonProperties={commonProperties}
        attributesKey={attributesKey}
        attributesName={attributesName}
      />
    </>
  );
};

const Header = styled.div`
  display: flex;
  align-items: flex-end;
  height: 46px;
  padding: 8px 16px;
  background-color: #fff;
  border-bottom: 1px solid #e0e0e0;
`;

const Title = styled.div`
  font-size: 16px;
  line-height: 22px;
`;

export default ViewPanel;
