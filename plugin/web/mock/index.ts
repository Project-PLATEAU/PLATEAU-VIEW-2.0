const forceUnregister = false;

export default async function () {
  if (!forceUnregister && import.meta.env.DEV) {
    const url = (await import("./mockServiceWorker.js?url")).default;
    (await import("./worker")).worker.start({
      serviceWorker: {
        url,
        options: {
          scope: "/",
        },
      },
    });
  } else {
    // disable msw
    (await navigator.serviceWorker.getRegistrations()).forEach(r => r.unregister());
  }
}
