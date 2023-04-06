import { postMsg } from "@web/extensions/sidebar/utils";
import { styled } from "@web/theme";
import { useCallback, useState } from "react";

import { BaseFieldProps, IdealZoom as IdealZoomType } from "../types";

export const initialCameraValues = {
  lng: 0,
  lat: 0,
  height: 0,
  pitch: 0,
  heading: 0,
  roll: 0,
};

const IdealZoom: React.FC<BaseFieldProps<"idealZoom">> = ({ value, editMode, onUpdate }) => {
  const [camera, setCamera] = useState<IdealZoomType["position"]>(
    value["position"] ?? initialCameraValues,
  );

  const handleLatitudeChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const lat = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 0;
      setCamera(c => {
        return {
          ...c,
          lat,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          lat,
        },
      });
    },
    [value, onUpdate],
  );

  const handleLongitudeChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const lng = !isNaN(parseFloat(e.currentTarget.value)) ? parseFloat(e.currentTarget.value) : 0;
      setCamera(c => {
        return {
          ...c,
          lng,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          lng,
        },
      });
    },
    [value, onUpdate],
  );

  const handleHeightChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const height = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 0;
      setCamera(c => {
        return {
          ...c,
          height,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          height,
        },
      });
    },
    [value, onUpdate],
  );

  const handleHeadingChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const heading = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 0;
      setCamera(c => {
        return {
          ...c,
          heading,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          heading,
        },
      });
    },
    [value, onUpdate],
  );

  const handlePitchChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const pitch = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 0;
      setCamera(c => {
        return {
          ...c,
          pitch,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          pitch,
        },
      });
    },
    [value, onUpdate],
  );

  const handleRollChange = useCallback(
    (e: React.ChangeEvent<HTMLInputElement>) => {
      const roll = !isNaN(parseFloat(e.currentTarget.value))
        ? parseFloat(e.currentTarget.value)
        : 0;
      setCamera(c => {
        return {
          ...c,
          roll,
        };
      });
      onUpdate({
        ...value,
        position: {
          ...value.position,
          roll,
        },
      });
    },
    [value, onUpdate],
  );

  const handleCapture = useCallback(() => {
    postMsg({ action: "getCurrentCamera" });
  }, []);

  const handleClean = useCallback(() => {
    setCamera(initialCameraValues);
  }, []);

  addEventListener("message", async e => {
    if (e.source !== parent) return;
    if (e.data.action) {
      if (e.data.action === "getCurrentCamera") {
        setCamera(e.data.payload);
        onUpdate({
          ...value,
          position: e.data.payload,
        });
      }
    }
  });

  return editMode ? (
    <div>
      <InnerWrapper>
        <Text>位置</Text>
        <InputWrapper>
          <Input
            type="number"
            placeholder="緯度" // Latitude
            value={camera.lat}
            onChange={handleLatitudeChange}
          />
          <Input
            type="number"
            placeholder="経度" // Longitude
            value={camera.lng}
            onChange={handleLongitudeChange}
          />
          <Input
            type="number"
            placeholder="高度" // Height
            value={camera.height}
            onChange={handleHeightChange}
          />
        </InputWrapper>
      </InnerWrapper>
      <InnerWrapper>
        <Text>ポーズ</Text>
        <InputWrapper>
          <Input
            type="number"
            placeholder="ヘッディング" // Heading
            value={camera.heading}
            onChange={handleHeadingChange}
          />
          <Input
            type="number"
            placeholder="ピッチ" // Pitch
            value={camera.pitch}
            onChange={handlePitchChange}
          />
          <Input
            type="number"
            placeholder="ロール" // Roll
            value={camera.roll}
            onChange={handleRollChange}
          />
        </InputWrapper>
      </InnerWrapper>
      <ButtonWrapper>
        <Button onClick={handleClean}>削除</Button>
        <Button onClick={handleCapture}>キャプチャー</Button>
      </ButtonWrapper>
    </div>
  ) : null;
};

export default IdealZoom;

const InnerWrapper = styled.div`
  display: flex;
  align-items: center;
`;

const Text = styled.p`
  margin: 0;
  width: 65px;
`;

const Input = styled.input`
  height: 32px;
  width: 64px;
  box-sizing: border-box;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  text-align: center;
`;

const InputWrapper = styled.div`
  display: flex;
  justify-content: space-between;
  align-items: center;
  gap: 4px;
  margin-bottom: 8px;
`;

const ButtonWrapper = styled.div`
  display: flex;
  justify-content: space-between;
  gap: 8px;
`;

const Button = styled.div`
  width: 100%;
  padding: 5px;
  border: 1px solid #d9d9d9;
  text-align: center;
  border-radius: 2px;
  user-select: none;
  cursor: pointer;

  :hover {
    background: #f4f4f4;
  }
`;
