import { BaseFieldProps } from "../types";

const PolylineColorGradient: React.FC<BaseFieldProps<"polylineColorGradient">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  console.log(value, onUpdate);
  return editMode ? <>Polyline Color Gradient</> : null;
};

export default PolylineColorGradient;
