import styled from "@emotion/styled";
import React, { useCallback, useEffect } from "react";
import { useParams } from "react-router-dom";

import Button from "@reearth-cms/components/atoms/Button";
import Content from "@reearth-cms/components/atoms/Content";
import Form from "@reearth-cms/components/atoms/Form";
import Icon from "@reearth-cms/components/atoms/Icon";
import Input from "@reearth-cms/components/atoms/Input";
import Modal from "@reearth-cms/components/atoms/Modal";
import TextArea from "@reearth-cms/components/atoms/TextArea";
import Typography from "@reearth-cms/components/atoms/Typography";
import { useT } from "@reearth-cms/i18n";

import useHooks from "./hooks";

export interface FormValues {
  name: string;
  description: string;
}

const ProjectSettings: React.FC = () => {
  const t = useT();
  const { confirm } = Modal;
  const [form] = Form.useForm();
  const { projectId } = useParams();

  const { project, handleProjectUpdate, handleProjectDelete } = useHooks({
    projectId,
  });

  useEffect(() => {
    form.setFieldsValue({
      name: project?.name,
      description: project?.description,
    });
  }, [form, project]);

  const handleSubmit = useCallback(() => {
    form
      .validateFields()
      .then(async values => {
        handleProjectUpdate({
          name: values.name,
          description: values.description,
        });
      })
      .catch(info => {
        console.log("Validate Failed:", info);
      });
  }, [form, handleProjectUpdate]);

  const handleProjectDeleteConfirmation = useCallback(() => {
    confirm({
      title: t("Are you sure you want to delete this project?"),
      icon: <Icon icon="exclamationCircle" />,
      onOk() {
        handleProjectDelete();
      },
    });
  }, [confirm, handleProjectDelete, t]);

  return (
    <PaddedContent>
      <ProjectSection> {project?.name} </ProjectSection>
      <ProjectSection>
        <Form style={{ maxWidth: 400 }} form={form} layout="vertical" autoComplete="off">
          <Form.Item name="name" label={t("Name")}>
            <Input />
          </Form.Item>
          <Form.Item
            name="description"
            label={t("Description")}
            extra={t("Write something here to describe this record.")}>
            <TextArea rows={4} />
          </Form.Item>
          <Form.Item>
            <Button onClick={handleSubmit} type="primary" htmlType="submit">
              {t("Save changes")}
            </Button>
          </Form.Item>
        </Form>
      </ProjectSection>
      <ProjectSection>
        <Typography style={{ marginBottom: 16 }}>{t("Danger Zone")}</Typography>
        <Button onClick={handleProjectDeleteConfirmation} type="primary" danger>
          {t("Delete project")}
        </Button>
      </ProjectSection>
    </PaddedContent>
  );
};

const PaddedContent = styled(Content)`
  margin: 16px;
`;

const ProjectSection = styled.div`
  background-color: #fff;
  font-weight: 500;
  font-size: 20px;
  line-height: 28px;
  color: rgba(0, 0, 0, 0.85);
  padding: 22px 24px;
  margin-bottom: 16px;
`;

export default ProjectSettings;
