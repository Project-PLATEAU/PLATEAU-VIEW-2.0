import CommonPage from "@web/extensions/sidebar/core/components/content/CommonPage";
import { ReearthApi } from "@web/extensions/sidebar/types";
import { Checkbox, Row } from "@web/sharedComponents";
import { styled } from "@web/theme";
import { memo } from "react";

import useHooks from "./hooks";

export type Props = {
  overrides: ReearthApi;
  isMobile?: boolean;
  onOverridesUpdate: (updatedProperties: Partial<ReearthApi>) => void;
};

const MapSettings: React.FC<Props> = ({ overrides, isMobile, onOverridesUpdate }) => {
  const {
    mapViewData,
    baseMapData,
    currentView,
    currentTiles,
    currentHideUnderground,
    currentAllowEnterGround,
    currentLightType,
    handleViewChange,
    handleTileChange,
    handleHideUnderGround,
    handleAllowEnterGround,
    handleUseSunLight,
  } = useHooks({ overrides, onOverridesUpdate });

  return (
    <CommonPage title="マップ設定" isMobile={isMobile}>
      <>
        <SubTitle>マップビュー</SubTitle>
        <Section>
          <ViewWrapper>
            {mapViewData.map(({ key, title }) => (
              <MapViewButton
                key={key}
                value={key}
                selected={currentView === key}
                onClick={() => handleViewChange(key)}>
                <Text style={{ color: " #FFFFFF" }}>{title}</Text>
              </MapViewButton>
            ))}
          </ViewWrapper>
          <CheckboxWrapper>
            <Checkbox checked={!!currentHideUnderground} onClick={handleHideUnderGround}>
              <Text>地下を隠す</Text>
            </Checkbox>
          </CheckboxWrapper>
          <CheckboxWrapper>
            <Checkbox checked={!!currentAllowEnterGround} onClick={handleAllowEnterGround}>
              <Text>地下に入る</Text>
            </Checkbox>
          </CheckboxWrapper>
          <CheckboxWrapper>
            <Checkbox checked={currentLightType === "sunLight"} onClick={handleUseSunLight}>
              <Text>太陽光を利用する</Text>
            </Checkbox>
          </CheckboxWrapper>
        </Section>
      </>
      <>
        <Title>ベースマップ</Title>
        <Section>
          <MapWrapper>
            <MapTitle>
              {baseMapData.find(item => item.key === currentTiles?.[0].id)?.title}
            </MapTitle>
            <Maps>
              {baseMapData.map(item => (
                <ImageButton
                  key={item.key}
                  selected={item.key === currentTiles?.[0].id}
                  onClick={() => handleTileChange(item)}
                  style={{
                    backgroundImage: "url(" + item.icon + ")",
                    backgroundSize: "cover",
                    backgroundRepeat: "no-repeat",
                  }}
                />
              ))}
            </Maps>
          </MapWrapper>
        </Section>
      </>
    </CommonPage>
  );
};

export default memo(MapSettings);

const Title = styled.p`
  font-size: 16px;
`;

const SubTitle = styled.p`
  font-size: 14px;
`;

const Text = styled.p``;

const Section = styled(Row)`
  gap: 16px;
`;

const ViewWrapper = styled.div`
  display: flex;
  gap: 12px;
  width: 100%;
`;

const CheckboxWrapper = styled.div`
  width: 100%;
  p {
    margin-bottom: 0;
  }
`;

const MapViewButton = styled.button<{ selected?: boolean }>`
  width: 91px;
  height: 29px;
  background: ${({ selected }) => (selected ? "#00bebe" : "#d1d1d1")};
  border-radius: 4px;
  border: none;
  padding: 4px 8px;
  cursor: pointer;
  transition: background 0.3s;

  :hover {
    background: #00bebe;
  }
`;

const MapTitle = styled.p`
  margin: 0;
  font-size: 10px;
  line-height: 14px;
`;

const MapWrapper = styled.div`
  display: flex;
  flex-direction: column;
  gap: 8px;
`;

const Maps = styled.div`
  display: flex;
  justify-content: start;
  gap: 8px;
  width: 100%;
`;

const ImageButton = styled.div<{ selected?: boolean }>`
  height: 64px;
  width: 64px;
  border-width: 2px;
  border-style: solid;
  border-color: ${({ selected }) => (selected ? "#00bebe" : "transparent")};
  border-radius: 2px;
  padding: 4px 8px;
  cursor: pointer;
  transition: border-color 0.2s;

  :hover {
    border-color: #00bebe;
  }
`;
