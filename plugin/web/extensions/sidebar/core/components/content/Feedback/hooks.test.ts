import { dataURItoBlob } from "./hooks";

test("dataURItoBlob", () => {
  const dummyData = new Blob(["dummy"], { type: "text/plain" });
  expect(dataURItoBlob("data:text/plain;base64,ZHVtbXk=")).toEqual(dummyData);
});
