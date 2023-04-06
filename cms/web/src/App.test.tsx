import { expect, test } from "vitest";

import { screen, render } from "@reearth-cms/test/utils";

test("hello, world", () => {
  render(<h1>Vite</h1>);
  expect(screen.getByText(/Vite/i)).toBeInTheDocument();
});
