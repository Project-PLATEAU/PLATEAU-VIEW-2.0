export const geoFormats = ".kml,.czml,.topojson,.geojson";
export const geo3dFormats = ".json";
export const geoMvtFormat = ".mvt";
export const model3dFormats = ".gltf,.glb";
export const fileAllFormats = `${geoFormats},${geo3dFormats},${geoMvtFormat},${model3dFormats}`;

export const imageFormats = ".jpg,.jpeg,.png,.gif,.tiff,.webp";
export const imageSVGFormat = ".svg";
export const imageAllFormats = `${imageFormats},${imageSVGFormat}`;

export const compressedFileFormats = ".7z,.zip";

export const acceptedFormats = `${imageAllFormats},${fileAllFormats},${compressedFileFormats}`;
