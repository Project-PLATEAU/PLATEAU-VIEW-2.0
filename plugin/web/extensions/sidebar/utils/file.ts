export const getExtension = (filename?: string) => {
  if (!filename || !filename.includes(".")) return "";

  return filename.toLowerCase().slice(filename.lastIndexOf(".") + 1, filename.length);
};

// getNameFromPath("xxx/yyy/zzz") -> "zzz"
export const getNameFromPath = (path?: string) => {
  if (!path) return;
  if (!path.includes("/")) return path;

  return path.split("/").slice(-1)[0];
};

export const createFileName = (name?: string, extension?: string) => {
  if (!name || !extension) return "";
  return `${name}.${extension}`;
};

export const normalizeExtension = (extension?: string) => {
  if (!extension) return "";
  return extension.toLowerCase().replace(/\s/g, "");
};
