import { setupWorker, rest } from "msw";

export const worker = setupWorker(
  rest.get("https://example.com/user/:userId", (_, res, ctx) => {
    return res(
      ctx.json({
        firstName: "John",
        lastName: "Maverick",
      }),
    );
  }),
);
