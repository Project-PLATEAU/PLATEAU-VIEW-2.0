import { postMsg } from "@web/extensions/storytelling/core/utils";
import EasyMDE from "easymde";
import { useCallback, useEffect, useRef, useState } from "react";

export default () => {
  const storyId = useRef<string>();
  const titleRef = useRef<HTMLInputElement>(null);
  const descriptionRef = useRef<HTMLTextAreaElement>(null);
  const easyMDE = useRef<EasyMDE | null>(null);
  const [canSave, setCanSave] = useState<boolean>(false);

  const onTitleChange = useCallback(() => {
    if (titleRef.current) {
      setCanSave(!!titleRef.current.value);
    }
  }, []);

  const onCancel = useCallback(() => {
    postMsg("sceneEditorClose");
  }, []);

  const onSave = useCallback(() => {
    postMsg("sceneSave", {
      id: storyId.current,
      title: titleRef.current?.value,
      description: easyMDE.current?.value(),
    });
  }, []);

  useEffect(() => {
    if (descriptionRef.current) {
      easyMDE.current = new EasyMDE({
        element: descriptionRef.current,
        maxHeight: "158px",
        status: false,
        hideIcons: ["fullscreen", "side-by-side", "guide"],
      });
    }

    if ((window as any).sceneEdit && titleRef.current && descriptionRef.current) {
      storyId.current = (window as any).sceneEdit.id;
      titleRef.current.value = (window as any).sceneEdit.title;
      setCanSave(!!titleRef.current.value);

      if (easyMDE.current) {
        easyMDE.current.value((window as any).sceneEdit.description);
      }
    }

    return () => {
      if (easyMDE.current) {
        easyMDE.current.toTextArea();
        easyMDE.current = null;
      }
    };
  }, []);

  return {
    titleRef,
    descriptionRef,
    canSave,
    onTitleChange,
    onCancel,
    onSave,
  };
};
