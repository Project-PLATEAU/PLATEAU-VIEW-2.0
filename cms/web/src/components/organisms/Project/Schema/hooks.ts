import { useCallback, useEffect, useState } from "react";
import { useNavigate, useParams } from "react-router-dom";

import Notification from "@reearth-cms/components/atoms/Notification";
import { Field, FieldType } from "@reearth-cms/components/molecules/Schema/types";
import {
  useCreateFieldMutation,
  SchemaFieldType,
  SchemaFieldTypePropertyInput,
  useDeleteFieldMutation,
  useUpdateFieldMutation,
  useUpdateFieldsMutation,
} from "@reearth-cms/gql/graphql-client-api";
import { useT } from "@reearth-cms/i18n";
import { useModel } from "@reearth-cms/state";

export default () => {
  const t = useT();
  const navigate = useNavigate();
  const { projectId, workspaceId, modelId } = useParams();
  const [currentModel] = useModel();

  const [fieldCreationModalShown, setFieldCreationModalShown] = useState(false);
  const [fieldUpdateModalShown, setFieldUpdateModalShown] = useState(false);
  const [selectedField, setSelectedField] = useState<Field | null>(null);
  const [selectedType, setSelectedType] = useState<FieldType | null>(null);
  const [collapsed, collapse] = useState(false);

  useEffect(() => {
    if (!modelId && currentModel) {
      navigate(`/workspace/${workspaceId}/project/${projectId}/schema/${currentModel.id}`);
    }
  }, [modelId, currentModel, navigate]); // eslint-disable-line react-hooks/exhaustive-deps

  const handleModelSelect = useCallback(
    (modelId: string) => {
      navigate(`/workspace/${workspaceId}/project/${projectId}/schema/${modelId}`);
    },
    [navigate, workspaceId, projectId],
  );

  const handleFieldKeyUnique = useCallback(
    (key: string, fieldId?: string): boolean => {
      return !currentModel?.schema.fields.some(
        field => field.key === key && (!fieldId || (fieldId && fieldId !== field.id)),
      );
    },
    [currentModel],
  );

  const [createNewField] = useCreateFieldMutation({
    refetchQueries: ["GetModels"],
  });

  const [updateField] = useUpdateFieldMutation({
    refetchQueries: ["GetModels"],
  });

  const [deleteFieldMutation] = useDeleteFieldMutation({
    refetchQueries: ["GetModels"],
  });

  const handleFieldDelete = useCallback(
    async (fieldId: string) => {
      if (!modelId) return;
      const results = await deleteFieldMutation({ variables: { modelId, fieldId } });
      if (results.errors) {
        Notification.error({ message: t("Failed to delete field.") });
        return;
      }
      Notification.success({ message: t("Successfully deleted field!") });
    },
    [modelId, deleteFieldMutation, t],
  );

  const handleFieldUpdate = useCallback(
    async (data: {
      fieldId: string;
      title: string;
      description: string;
      key: string;
      multiple: boolean;
      unique: boolean;
      required: boolean;
      typeProperty: SchemaFieldTypePropertyInput;
    }) => {
      if (!modelId) return;
      const field = await updateField({
        variables: {
          modelId,
          fieldId: data.fieldId,
          title: data.title,
          description: data.description,
          key: data.key,
          multiple: data.multiple,
          unique: data.unique,
          required: data.required,
          typeProperty: data.typeProperty,
        },
      });
      if (field.errors || !field.data?.updateField) {
        Notification.error({ message: t("Failed to update field.") });
        return;
      }
      Notification.success({ message: t("Successfully updated field!") });
      setFieldUpdateModalShown(false);
    },
    [modelId, updateField, t],
  );

  const [updateFieldsOrder] = useUpdateFieldsMutation({
    refetchQueries: ["GetModels"],
  });

  const handleFieldOrder = useCallback(
    async (fields: Field[]) => {
      if (!modelId) return;
      const response = await updateFieldsOrder({
        variables: {
          updateFieldInput: fields.map((field, index) => ({
            modelId,
            fieldId: field.id,
            order: index,
          })),
        },
      });
      if (response.errors || !response?.data?.updateFields) {
        Notification.error({ message: t("Failed to update field.") });
        return;
      }
      Notification.success({ message: t("Successfully updated field!") });
      setFieldUpdateModalShown(false);
    },
    [modelId, updateFieldsOrder, t],
  );

  const handleFieldCreate = useCallback(
    async (data: {
      title: string;
      description: string;
      key: string;
      multiple: boolean;
      unique: boolean;
      required: boolean;
      type: FieldType;
      typeProperty: SchemaFieldTypePropertyInput;
    }) => {
      if (!modelId) return;
      const field = await createNewField({
        variables: {
          modelId,
          title: data.title,
          description: data.description,
          key: data.key,
          multiple: data.multiple,
          unique: data.unique,
          required: data.required,
          type: data.type as SchemaFieldType,
          typeProperty: data.typeProperty,
        },
      });
      if (field.errors || !field.data?.createField) {
        Notification.error({ message: t("Failed to create field.") });
        setFieldCreationModalShown(false);
        return;
      }
      Notification.success({ message: t("Successfully created field!") });
      setFieldCreationModalShown(false);
    },
    [modelId, createNewField, t],
  );

  const handleFieldCreationModalClose = useCallback(() => setFieldCreationModalShown(false), []);

  const handleFieldCreationModalOpen = useCallback(
    (fieldType: FieldType) => {
      setSelectedType(fieldType);
      if (modelId) setFieldCreationModalShown(true);
    },
    [modelId],
  );

  const handleFieldUpdateModalClose = useCallback(() => {
    setSelectedField(null);
    setFieldUpdateModalShown(false);
  }, [setSelectedField]);

  const handleFieldUpdateModalOpen = useCallback(
    (field: Field) => {
      setSelectedType(field.type);
      setSelectedField(field);
      setFieldUpdateModalShown(true);
    },
    [setSelectedField],
  );

  return {
    fieldCreationModalShown,
    fieldUpdateModalShown,
    selectedField,
    currentModel,
    selectedType,
    collapsed,
    collapse,
    handleModelSelect,
    handleFieldCreationModalClose,
    handleFieldCreationModalOpen,
    handleFieldUpdateModalOpen,
    handleFieldUpdateModalClose,
    handleFieldCreate,
    handleFieldKeyUnique,
    handleFieldUpdate,
    handleFieldOrder,
    handleFieldDelete,
  };
};
