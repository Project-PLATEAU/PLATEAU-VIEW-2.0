import { styled } from "@web/theme";

import Spin from "../Spin";

import useModifiedImage from "./hooks";

type Props = {
  imageUrl: string;
  blendColor: string;
  width: number;
  height: number;
};

const ModifiedImage = ({ imageUrl, blendColor, width, height }: Props) => {
  const { modifiedImageUrl, loading } = useModifiedImage({
    imageUrl,
    blendColor,
    width,
    height,
  });

  return loading ? (
    <Loading width={width} height={height}>
      <Spin />
    </Loading>
  ) : (
    <img src={modifiedImageUrl} alt="modified image" />
  );
};

export default ModifiedImage;

const Loading = styled.div<{ width: number; height: number }>`
  position: relative;
  width: ${({ width }) => width};px;
  height:  ${({ height }) => height};px;
  display: flex;
  align-items: center;
  justify-content: center;
`;
