import { styled } from "@web/theme";

export type Tag = {
  name: string;
  type: "location" | "data-type";
};

const Tags: React.FC<{ tags?: Tag[]; onTagSelect?: (tag: Tag) => void }> = ({
  tags,
  onTagSelect,
}) => {
  return (
    <TagWrapper>
      {tags?.map(tag => (
        <Tag key={tag.name} type={tag.type} onClick={() => onTagSelect?.(tag)}>
          {tag.name}
        </Tag>
      ))}
    </TagWrapper>
  );
};

export default Tags;

const TagWrapper = styled.div`
  display: flex;
  gap: 12px;
`;

const Tag = styled.p<{ type?: "location" | "data-type" }>`
  line-height: 16px;
  height: 32px;
  padding: 8px 12px;
  margin: 0;
  background: #ffffff;
  border-left: 2px solid ${({ type }) => (type === "location" ? "#03c3ff" : "#1ED500")};
  box-shadow: 2px 2px 4px rgba(0, 0, 0, 0.1);
  cursor: pointer;
`;
