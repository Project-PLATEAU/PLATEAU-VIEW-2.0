import styled from "@emotion/styled";

type Props = { url: string };

const ImageViewer: React.FC<Props> = ({ url }) => <Image src={url} alt="image-preview" />;

const Image = styled.img`
  width: 100%;
  height: 500px;
  object-fit: contain;
`;

export default ImageViewer;
