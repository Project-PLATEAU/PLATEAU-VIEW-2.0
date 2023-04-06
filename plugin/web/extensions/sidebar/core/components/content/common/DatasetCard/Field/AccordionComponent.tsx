import { Icon } from "@web/sharedComponents";
import { styled } from "@web/theme";
import {
  Accordion,
  AccordionItem,
  AccordionItemHeading,
  AccordionItemButton,
  AccordionItemPanel,
  AccordionItemState,
} from "react-accessible-accordion";

type Props = {
  id: string;
  hasGroup?: boolean;
  editMode?: boolean;
  hasUI?: boolean;
  showGroupIcon?: boolean;
  showArrowIcon?: boolean;
  title: string;
  editModeTitle: string;
  onGroupSelectOpen: (e: any) => void;
  onRemove: (e: any) => void;
  onUpClick: (e: any) => void;
  onDownClick: (e: any) => void;
  children?: React.ReactNode;
};

const AccordionComponent: React.FC<Props> = ({
  id,
  hasGroup,
  editMode,
  hasUI,
  showGroupIcon,
  showArrowIcon,
  title,
  editModeTitle,
  onGroupSelectOpen,
  onRemove,
  onUpClick,
  onDownClick,
  children,
}) => {
  return (
    <StyledAccordionComponent allowZeroExpanded preExpanded={[id]} hide={!editMode && !hasUI}>
      <AccordionItem uuid={id}>
        <AccordionItemState>
          {({ expanded }) => (
            <Header showBorder={expanded}>
              {editMode ? (
                <HeaderContents>
                  <LeftContents>
                    {showArrowIcon && (
                      <ArrowIcon icon="arrowDown" size={16} direction="right" expanded={expanded} />
                    )}
                    <Title>{editModeTitle}</Title>
                  </LeftContents>
                  <RightContents>
                    <StyledIcon icon="arrowUpThin" size={16} onClick={onUpClick} />
                    <StyledIcon icon="arrowDownThin" size={16} onClick={onDownClick} />
                    {showGroupIcon && (
                      <StyledIcon
                        icon="group"
                        color={hasGroup ? "#00BEBE" : "inherit"}
                        size={16}
                        onClick={onGroupSelectOpen}
                      />
                    )}
                    <StyledIcon icon="trash" size={16} onClick={onRemove} />
                  </RightContents>
                </HeaderContents>
              ) : (
                <HeaderContents>
                  <Title>{title}</Title>
                  <ArrowIcon icon="arrowDown" size={16} direction="left" expanded={expanded} />
                </HeaderContents>
              )}
            </Header>
          )}
        </AccordionItemState>
        <BodyWrapper>{children}</BodyWrapper>
      </AccordionItem>
    </StyledAccordionComponent>
  );
};

export default AccordionComponent;

const StyledAccordionComponent = styled(Accordion)<{ hide: boolean }>`
  ${({ hide }) => hide && "display: none;"}
  width: 100%;
  border: 1px solid #e6e6e6;
  border-radius: 4px;
  background: #ffffff;
`;

const Header = styled(AccordionItemHeading)<{ showBorder?: boolean }>`
  border-bottom-width: 1px;
  border-bottom-style: solid;
  border-bottom-color: transparent;
  ${({ showBorder }) => showBorder && "border-bottom-color: #e0e0e0;"}
  display: flex;
  height: auto;
`;

const HeaderContents = styled(AccordionItemButton)`
  display: flex;
  justify-content: space-between;
  align-items: center;
  flex: 1;
  padding: 12px;
  outline: none;
  cursor: pointer;
`;

const BodyWrapper = styled(AccordionItemPanel)`
  position: relative;
  border-radius: 0px 0px 4px 4px;
  padding: 12px;
`;

const Title = styled.p`
  margin: 0;
  user-select: none;
  width: 160px;
  overflow-wrap: break-word;
`;

const StyledIcon = styled(Icon)`
  cursor: pointer;
`;

const LeftContents = styled.div`
  display: flex;
  align-items: center;
`;

const RightContents = styled.div`
  display: flex;
  gap: 4px;
`;

const ArrowIcon = styled(Icon)<{ direction: "left" | "right"; expanded?: boolean }>`
  transition: transform 0.15s ease;
  ${({ direction, expanded }) =>
    (direction === "right" && !expanded && "transform: rotate(-90deg);") ||
    (direction === "left" && !expanded && "transform: rotate(90deg);") ||
    null}
  ${({ direction }) => (direction === "left" ? "margin: 0 -4px 0 4px;" : "margin: 0 4px 0 -4px;")}
`;
