import update from "immutability-helper";
import { omit } from "lodash";
import { useCallback, useEffect, useState, useRef, useMemo } from "react";

import { postMsg, commonPropertiesMap, getAttributes, getRootFields } from "../core/utils";
import { InfoboxTemplate, Properties, Field, FldInfo } from "../types";

export type EditorTab = "view" | "edit";

export default () => {
  const [template, setTemplate] = useState<InfoboxTemplate | undefined>();
  const [fldInfo, setFldInfo] = useState<FldInfo | undefined>();
  const [properties, setProperties] = useState<Properties>();

  const [fields, setFields] = useState<Field[]>([]);
  const [editorTab, setEditorTab] = useState<EditorTab>("view");
  const [dataState, setDataState] = useState<"loading" | "empty" | "ready">("loading");
  const [isSaving, setIsSaving] = useState<boolean>(false);
  const [inEditor, setInEditor] = useState(false);
  const [commonProperties, setCommonProperties] = useState<string[]>([]);

  const handleEditorTab = useCallback((tab: EditorTab) => {
    setEditorTab(tab);
  }, []);

  useEffect(() => {
    const fieldItems: Field[] = [];
    const commonProperties = template?.dataType ? commonPropertiesMap[template.dataType] : [];
    setCommonProperties(commonProperties);

    // show fields with default order if no settings
    if (!template?.fields || template.fields?.length === 0) {
      commonProperties.forEach(fp => {
        fieldItems.push({
          title: "",
          path: fp,
          visible: true,
        });
      });
    } else {
      // show fields with settings

      // Incase there exists old data from database
      // which contains floor fields already
      template.fields.forEach(f => {
        if (commonProperties.includes(f.path)) {
          fieldItems.push({
            ...f,
          });
        }
      });

      // or only contains some of the commonProperties
      commonProperties
        .filter(fp => !template.fields?.map(f => f.path).includes(fp))
        .forEach(fp => {
          fieldItems.push({
            path: fp,
            title: "",
            visible: true,
          });
        });
    }

    setFields(fieldItems);
  }, [template]);

  const onFieldCheckChange = useCallback((e: any) => {
    setFields(list => {
      const fieldItem = list.find(item => item.path === e.target["data-path"]);
      if (fieldItem) {
        fieldItem.visible = !!e.target.checked;
      }
      return [...list];
    });
  }, []);

  const onFieldTitleChange = useCallback((e: any) => {
    setFields(list => {
      const fieldItem = list.find(item => item.path === e.target.dataset.path);
      if (fieldItem) {
        fieldItem.title = e.target.value;
      }
      return [...list];
    });
  }, []);

  const onFieldMove = useCallback((dragIndex: number, hoverIndex: number) => {
    setFields((prevFields: Field[]) =>
      update(prevFields, {
        $splice: [
          [dragIndex, 1],
          [hoverIndex, 0, prevFields[dragIndex] as Field],
        ],
      }),
    );
  }, []);

  const saveTemplate = useCallback(() => {
    setIsSaving(true);
    postMsg("saveTemplate", {
      ...template,
      fields,
    });
  }, [fields, template]);

  const onMessage = useCallback((e: MessageEvent<any>) => {
    if (e.source !== parent) return;
    switch (e.data.action) {
      case "getInEditor":
        setInEditor(e.data.payload);
        break;
      case "fillData":
        setTemplate(e.data.payload.template);
        setFldInfo(e.data.payload.fldInfo);
        setProperties(e.data.payload.properties);
        setFldInfo(e.data.payload.fldInfo);
        setDataState("ready");
        setEditorTab("view");
        break;
      case "setLoading":
        setDataState("loading");
        break;
      case "setEmpty":
        setDataState("empty");
        setTemplate(undefined);
        setProperties({});
        break;
      case "saveFinish":
        setIsSaving(false);
        break;
      default:
        break;
    }
  }, []);

  useEffect(() => {
    addEventListener("message", onMessage);
    return () => {
      removeEventListener("message", onMessage);
    };
  }, [onMessage]);

  const wrapperRef = useRef<HTMLDivElement>(null);
  useEffect(() => {
    const wrapperResizeObserver = new ResizeObserver(entries => {
      const [entry] = entries;
      let height: number | undefined;
      if (entry.contentBoxSize) {
        const contentBoxSize = Array.isArray(entry.contentBoxSize)
          ? entry.contentBoxSize[0]
          : entry.contentBoxSize;
        height = contentBoxSize.blockSize;
      } else if (entry.contentRect) {
        height = entry.contentRect.height;
      }

      if (height) {
        document.documentElement.style.height = `${height}px`;
      }
    });

    if (wrapperRef.current) {
      wrapperResizeObserver.observe(wrapperRef.current);
    }
  }, []);

  useEffect(() => {
    postMsg("init");
  }, []);

  const actualProperties = useMemo((): [any, string | undefined] => {
    const rootFields = properties ? getRootFields(properties, template?.dataType, fldInfo) : {};
    return properties
      ? {
          ...(properties.attributes
            ? { [attributesKey]: getAttributes(properties.attributes, "label") }
            : {}),
          ...rootFields,
          // use omit to keep field order
          ...omit(properties, [...Object.keys(rootFields), "attributes"]),
        }
      : undefined;
  }, [fldInfo, properties, template?.dataType]);

  return {
    inEditor,
    dataState,
    properties: actualProperties,
    fields,
    template,
    wrapperRef,
    isSaving,
    editorTab,
    commonProperties,
    attributesName,
    attributesKey,
    handleEditorTab,
    onFieldCheckChange,
    onFieldTitleChange,
    onFieldMove,
    saveTemplate,
  };
};

const attributesName = "全ての属性";
const attributesKey = "attributes";
