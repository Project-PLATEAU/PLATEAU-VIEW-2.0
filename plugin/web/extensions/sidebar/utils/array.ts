export function array_move(arr: any[], old_index: number, new_index: number) {
  if (new_index >= arr.length) {
    let k = new_index - arr.length + 1;
    while (k--) {
      arr.push(undefined);
    }
  }
  arr.splice(new_index, 0, arr.splice(old_index, 1)[0]);
}

export const swap = (arr: any[], index1: number, index2: number) => {
  [arr[index1], arr[index2]] = [arr[index2], arr[index1]];
};

export const moveItemUp = (index: number, array?: any[]) => {
  if (!array || index === 0) return;

  const newArray = [...array];
  swap(newArray, index, index - 1);
  return newArray;
};

export const moveItemDown = (index: number, array?: any[]) => {
  if (!array || (array && index >= array.length - 1)) return;

  const newArray = [...array];
  swap(newArray, index, index + 1);
  return newArray;
};

export const removeItem = (index: number, array?: any[]) => {
  if (!array) return;

  const newArray = [...array];
  newArray.splice(index, 1);
  return newArray;
};
