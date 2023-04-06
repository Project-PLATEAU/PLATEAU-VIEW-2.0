import { Divider } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { Children, Fragment, ReactNode } from "react";

export type Props = {
  title?: string;
  isMobile?: boolean;
  children?: ReactNode;
};

const CommonPageWrapper: React.FC<Props> = ({ title, isMobile, children }) => {
  const childArray = Children.toArray(children);

  return (
    <Wrapper isMobile={isMobile}>
      {title && !isMobile && (
        <>
          <Title>{title}</Title>
          <Divider />
        </>
      )}
      {childArray.map((child, idx) => (
        <Fragment key={idx}>
          {child}
          {idx + 1 !== childArray.length && (
            <Divider style={{ margin: isMobile ? "12px 0" : "24px 0" }} />
          )}
        </Fragment>
      ))}
    </Wrapper>
  );
};

export default CommonPageWrapper;

const Wrapper = styled.div<{ isMobile?: boolean }>`
  padding: ${({ isMobile }) => (isMobile ? "12px" : "32px 16px")};
  color: #4a4a4a;
`;

const Title = styled.p`
  font-size: 16px;
`;
