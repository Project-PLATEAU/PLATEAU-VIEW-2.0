import styled from "@emotion/styled";

import Icon from "@reearth-cms/components/atoms/Icon";

import Avatar, { AvatarProps } from "../Avatar";

type Props = {
  username?: string;
  shadow?: boolean;
  anonymous?: boolean;
} & AvatarProps;

const UserAvatar: React.FC<Props> = ({ username, shadow, anonymous, ...props }) => {
  return (
    <UserAvatarWrapper shadow={shadow} anonymous={anonymous} {...props}>
      {anonymous ? <Icon icon="user" /> : username?.charAt(0)}
    </UserAvatarWrapper>
  );
};

const UserAvatarWrapper = styled(Avatar) <{ shadow?: boolean; anonymous?: boolean }>`
  color: #fff;
  background-color: ${({ anonymous }) => (anonymous ? "#BFBFBF" : "#3F3D45")};
  box-shadow: ${({ shadow }) => (shadow ? "0px 4px 4px rgba(0, 0, 0, 0.25)" : "none")};
`;

export default UserAvatar;
