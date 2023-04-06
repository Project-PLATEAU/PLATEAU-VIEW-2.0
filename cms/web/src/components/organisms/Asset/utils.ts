export function fileName(url: string | undefined): string {
  if (!url) return "";
  let files: string[];
  try {
    files = new URL(url).pathname.split("/");
  } catch {
    files = url.split("/");
  }
  return files.length ? files[files.length - 1] : "";
}
