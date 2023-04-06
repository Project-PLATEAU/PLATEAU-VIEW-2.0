import { styled } from "@web/theme";
import { CSSProperties, memo } from "react";

import icons from "./icons";

type Props = {
  className?: string;
  icon: string;
  size?: string | number;
  width?: string | number;
  height?: string | number;
  color?: string;
  wide?: boolean;
  cursor?: CSSProperties["cursor"];
  onClick?: (e?: React.MouseEvent<HTMLDivElement, MouseEvent>) => void;
};

type Icons = keyof typeof icons;

const Icon: React.FC<Props> = ({
  className,
  icon,
  size = 24,
  width,
  height,
  color,
  wide,
  cursor,
  onClick,
}) => {
  const sizeStr = typeof size === "number" ? `${size}px` : size;
  const widthStr = typeof width === "number" ? `${width}px` : width;
  const heightStr = typeof height === "number" ? `${height}px` : height;
  const IconComponent = icons[icon as Icons];

  return (
    <Wrapper
      className={className}
      size={sizeStr}
      width={widthStr}
      height={heightStr}
      color={color}
      wide={wide}
      cursor={cursor}
      onClick={onClick}>
      <IconComponent />
    </Wrapper>
  );
};

const Wrapper = styled.div<{
  size: string;
  width?: string;
  height?: string;
  color?: string;
  wide?: boolean;
  cursor?: CSSProperties["cursor"];
}>`
  box-sizing: content-box;
  width: ${({ size, width }) => (width ? width : size)};
  ${({ wide, size, height }) => !wide && `height: ${height ? height : size};`}
  cursor: ${({ cursor }) => cursor};

  svg {
    width: ${({ size, width }) => (width ? width : size)};

    ${({ wide, size, height }) => !wide && `height:  ${height ? height : size};`}
    color: ${({ color }) => color};
  }
`;

export default memo(Icon);
