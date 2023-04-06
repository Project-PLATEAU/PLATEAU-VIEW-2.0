import styled from "@emotion/styled";

import DropdownAtom from "@reearth-cms/components/atoms/Dropdown";
import Icon from "@reearth-cms/components/atoms/Icon";
import Space from "@reearth-cms/components/atoms/Space";
import UserAvatar from "@reearth-cms/components/atoms/UserAvatar";

export type Props = {
  className?: string;
  menu: React.ReactElement | (() => React.ReactElement);
  name?: string;
  personal?: boolean;
};

const Dropdown: React.FC<Props> = ({ className, menu, name, personal }) => {
  return (
    <StyledDropdown className={className} overlay={menu} trigger={["click"]}>
      <a onClick={e => e.preventDefault()}>
        <Space>
          <UserAvatar username={name ?? ""} shape={personal ? "circle" : "square"} size={"small"} />
          <Text>{name}</Text>
          <StyledIcon icon="caretDown" />
        </Space>
      </a>
    </StyledDropdown>
  );
};

export default Dropdown;

const StyledDropdown = styled(DropdownAtom)`
  padding-left: 10px;
  color: #fff;
  background-color: #1d1d1d;
`;

const StyledIcon = styled(Icon)`
  color: #8c8c8c;
`;

const Text = styled.p`
  max-width: 300px;
  margin: 0;
  white-space: nowrap;
  overflow: hidden;
  text-overflow: ellipsis;
`;
