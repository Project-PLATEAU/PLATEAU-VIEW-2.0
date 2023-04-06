import styled from "@emotion/styled";
import React from "react";

import Icon from "@reearth-cms/components/atoms/Icon";
import List from "@reearth-cms/components/atoms/List";
import { useT } from "@reearth-cms/i18n";

import { fieldTypes } from "./fieldTypes";
import { FieldType } from "./types";

export interface Props {
  className?: string;
  addField: (fieldType: FieldType) => void;
}

type FieldListItem = { title: string; fields: string[] };

const FieldList: React.FC<Props> = ({ addField }) => {
  const t = useT();

  const data: FieldListItem[] = [
    {
      title: t("Text"),
      fields: ["Text", "TextArea", "MarkdownText"],
    },
    {
      title: t("Asset"),
      fields: ["Asset"],
    },
    {
      title: t("Boolean"),
      fields: ["Bool"],
    },
    {
      title: t("Select"),
      fields: ["Select"],
    },
    {
      title: t("Number"),
      fields: ["Integer"],
    },
    {
      title: t("URL"),
      fields: ["URL"],
    },
  ];

  return (
    <>
      <h1>{t("Add Field")}</h1>
      <FieldStyledList
        itemLayout="horizontal"
        dataSource={data}
        renderItem={item => (
          <>
            <FieldCategoryTitle>{(item as FieldListItem).title}</FieldCategoryTitle>
            {(item as FieldListItem).fields?.map((field: string) => (
              <List.Item key={field} onClick={() => addField(field as FieldType)}>
                <Meta
                  avatar={<Icon icon={fieldTypes[field].icon} color={fieldTypes[field].color} />}
                  title={t(fieldTypes[field].title)}
                  description={t(fieldTypes[field].description)}
                />
              </List.Item>
            ))}
          </>
        )}
      />
    </>
  );
};

const FieldCategoryTitle = styled.h2`
  font-weight: 400;
  font-size: 12px;
  line-height: 20px;
  margin-bottom: 12px;
  margin-top: 12px;
  color: rgba(0, 0, 0, 0.45);
`;

const FieldStyledList = styled(List)`
  .ant-list-item {
    background-color: #fff;
    cursor: pointer;
    + .ant-list-item {
      margin-top: 12px;
    }
    padding: 4px;
    box-shadow: 0px 2px 8px #00000026;
    .ant-list-item-meta {
      .ant-list-item-meta-title {
        margin: 0;
      }
      align-items: center;
      .ant-list-item-meta-avatar {
        border: 1px solid #f0f0f0;
        width: 28px;
        height: 28px;
        display: flex;
        justify-content: center;
        align-items: center;
      }
    }
  }
`;

const Meta = styled(List.Item.Meta)`
  .ant-list-item-meta-description {
    font-size: 12px;
  }
`;

export default FieldList;
