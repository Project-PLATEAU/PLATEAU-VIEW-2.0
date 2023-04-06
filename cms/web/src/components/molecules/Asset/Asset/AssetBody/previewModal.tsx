import styled from "@emotion/styled";

import Modal from "@reearth-cms/components/atoms/Modal";

type Props = {
  url: string;
  visible: boolean;
  handleCancel: () => void;
};

const PreviewModal: React.FC<Props> = ({ url, visible, handleCancel }) => {
  return (
    <Modal
      centered
      visible={visible}
      onCancel={handleCancel}
      footer={null}
      width="90vw"
      bodyStyle={{
        height: "90vh",
      }}>
      <PreviewImage src={url} alt="asset-preview" />
    </Modal>
  );
};

const PreviewImage = styled.img`
  width: 100%;
  height: 100%;
  object-fit: contain;
`;

export default PreviewModal;
