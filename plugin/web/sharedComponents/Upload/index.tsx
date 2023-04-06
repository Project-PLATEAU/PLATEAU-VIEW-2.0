import type { UploadProps as UploadPropsType, UploadFile as UploadFileType } from "antd";
import { Upload, message } from "antd";

export type UploadProps = UploadPropsType;
export type UploadFile = UploadFileType;

export { message };

export default Upload;
