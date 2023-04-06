import styled from "@emotion/styled";
import React, { useCallback, useEffect, useMemo, useState } from "react";

import Button from "@reearth-cms/components/atoms/Button";
import Form from "@reearth-cms/components/atoms/Form";
import InnerContent from "@reearth-cms/components/atoms/InnerContents/basic";
import ContentSection from "@reearth-cms/components/atoms/InnerContents/ContentSection";
import Input from "@reearth-cms/components/atoms/Input";
import Select from "@reearth-cms/components/atoms/Select";
import Switch from "@reearth-cms/components/atoms/Switch";
import Table, { TableColumnsType } from "@reearth-cms/components/atoms/Table";
import { useT } from "@reearth-cms/i18n";

export type PublicScope = "private" | "public"; // Add "limited" when functionality becomes available

export type Model = {
  id: string;
  name?: string;
  public: boolean;
  key?: string;
};

export type ModelDataType = {
  id: string;
  name: string;
  public: JSX.Element;
  publicState: boolean;
  key?: string;
};

export type Props = {
  projectScope?: PublicScope;
  alias?: string;
  models?: Model[];
  assetPublic?: boolean;
  onPublicUpdate?: (
    alias?: string,
    scope?: PublicScope,
    modelsToUpdate?: Model[],
    assetPublic?: boolean,
  ) => void;
};

const Accessibility: React.FC<Props> = ({
  projectScope,
  models: rawModels,
  alias,
  assetPublic,
  onPublicUpdate,
}) => {
  const t = useT();
  const [scope, changeScope] = useState(projectScope);
  const [aliasState, setAlias] = useState(alias);
  const [updatedModels, setUpdatedModels] = useState<Model[]>([]);
  const [assetState, setAssetState] = useState<boolean | undefined>(assetPublic);
  const [models, setModels] = useState<Model[] | undefined>(rawModels);
  const [form] = Form.useForm();

  useEffect(() => {
    changeScope(projectScope);
  }, [projectScope]);

  useEffect(() => {
    setModels(rawModels);
  }, [rawModels]);

  useEffect(() => {
    setAlias(alias);
  }, [alias]);

  useEffect(() => {
    setAssetState(assetPublic);
  }, [assetPublic]);

  const saveDisabled = useMemo(
    () =>
      updatedModels.length === 0 &&
      projectScope === scope &&
      alias === aliasState &&
      assetPublic === assetState,
    [updatedModels.length, projectScope, scope, alias, aliasState, assetPublic, assetState],
  );

  const handlerAliasChange = useCallback((e: React.ChangeEvent<HTMLInputElement>) => {
    setAlias(e.currentTarget.value);
  }, []);

  const handlePublicUpdate = useCallback(() => {
    if (!scope && updatedModels.length === 0) return;
    onPublicUpdate?.(aliasState, scope, updatedModels, assetState);
    setUpdatedModels([]);
  }, [scope, aliasState, updatedModels, onPublicUpdate, assetState]);

  const handleUpdatedModels = useCallback(
    (model: Model) => {
      if (updatedModels.find(um => um.id === model.id)) {
        setUpdatedModels(ums => ums.filter(um => um.id !== model.id));
      } else {
        setUpdatedModels(ums => [...ums, model]);
      }
      setModels(ms => ms?.map(m => (m.id === model.id ? { ...m, public: model.public } : m)));
    },
    [updatedModels],
  );

  const handleUpdatedAssetState = useCallback((state: boolean) => {
    setAssetState(state);
  }, []);

  const columns: TableColumnsType<ModelDataType> = [
    {
      title: t("Model"),
      dataIndex: "name",
      key: "name",
      width: 220,
    },
    {
      title: t("Switch"),
      dataIndex: "public",
      key: "public",
      align: "right",
      width: 72,
    },
    {
      title: t("End point"),
      dataIndex: "endpoint",
      key: "endpoint",
      render: (_, modelData: ModelDataType) => {
        return (
          modelData.publicState &&
          modelData.key && (
            <a
              target="_blank"
              style={{ textDecoration: "underline", color: "#000000D9" }}
              href={window.REEARTH_CONFIG?.api + "/p/" + alias + "/" + modelData.key}
              rel="noreferrer">
              {window.REEARTH_CONFIG?.api}/p/{alias}/{modelData.key}
            </a>
          )
        );
      },
    },
  ];

  const dataSource: ModelDataType[] | undefined = useMemo(() => {
    let columns = [
      {
        id: "assets",
        name: t("Assets"),
        publicState: assetState ?? false,
        public: (
          <Switch
            checked={assetState}
            onChange={(publicState: boolean) => handleUpdatedAssetState(publicState)}
          />
        ),
      },
    ];

    if (models) {
      columns = [
        ...models.map(m => {
          return {
            id: m.id,
            name: m.name ?? "",
            key: m.key,
            publicState: m.public,
            public: (
              <Switch
                checked={m.public}
                onChange={(publicState: boolean) =>
                  handleUpdatedModels({ id: m.id, public: publicState, key: m.key })
                }
              />
            ),
          };
        }),
        ...columns,
      ];
    }
    return columns;
  }, [models, assetState, handleUpdatedAssetState, handleUpdatedModels, t]);

  const publicScopeList = [
    { id: 1, name: t("Private"), value: "private" },
    { id: 2, name: t("Public"), value: "public" },
  ];

  return (
    <InnerContent title={t("Accessibility")} flexChildren>
      <ContentSection title="">
        <Form form={form} layout="vertical" autoComplete="off">
          <ItemsWrapper>
            <Form.Item
              label={t("Public Scope")}
              extra={t(
                "Choose the scope of your project. This affects all the models shown below that are switched on.",
              )}>
              <Select value={scope} onChange={changeScope}>
                {publicScopeList.map(type => (
                  <Select.Option key={type.id} value={type.value}>
                    {type.name}
                  </Select.Option>
                ))}
              </Select>
            </Form.Item>

            <Form.Item label={t("Project Alias")}>
              <Input value={aliasState} onChange={handlerAliasChange} />
            </Form.Item>
          </ItemsWrapper>
          <TableWrapper>
            <Table dataSource={dataSource} columns={columns} pagination={false} />
          </TableWrapper>
          <Button type="primary" disabled={saveDisabled} onClick={handlePublicUpdate}>
            {t("Save changes")}
          </Button>
        </Form>
      </ContentSection>
    </InnerContent>
  );
};

export default Accessibility;

const ItemsWrapper = styled.div`
  max-width: 304px;
`;

const TableWrapper = styled.div`
  margin: 24px 0;
`;
