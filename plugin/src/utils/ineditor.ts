const reearth = (globalThis as any).reearth;

export function inEditor() {
  // editor or preview
  return (
    reearth.scene.inEditor ||
    (typeof reearth.scene.built === "boolean" ? !reearth.scene.built : false)
  );
}
