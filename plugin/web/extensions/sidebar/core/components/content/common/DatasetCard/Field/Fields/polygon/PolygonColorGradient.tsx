import { BaseFieldProps } from "../types";

const PolygonColorGradient: React.FC<BaseFieldProps<"polygonColorGradient">> = ({
  value,
  editMode,
  onUpdate,
}) => {
  // remember to update the BaseFieldProps type!
  console.log(value, editMode, onUpdate);
  return <div>Polygon Color Gradient</div>;
};

export default PolygonColorGradient;
