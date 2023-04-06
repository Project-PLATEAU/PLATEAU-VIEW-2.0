import { renderHook, act } from "@testing-library/react";
import { expect, test, vi } from "vitest";

import { useDelayedCount } from "./use-delayed-count";

test("normal", async () => {
  vi.useFakeTimers();

  const { result } = renderHook(() => useDelayedCount([100, 200, 300]));
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(undefined);

  act(() => {
    result.current[2](true);
  });
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(undefined);

  act(() => {
    result.current[2](false);
  });
  expect(result.current[0]).toBe(1);
  expect(result.current[1]).toBe(0);

  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(2);
  expect(result.current[1]).toBe(1);

  act(() => {
    vi.advanceTimersByTime(200);
  });
  expect(result.current[0]).toBe(3);
  expect(result.current[1]).toBe(2);

  act(() => {
    vi.advanceTimersByTime(300);
  });
  expect(result.current[0]).toBe(4);
  expect(result.current[1]).toBe(3);
  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(4);
  expect(result.current[1]).toBe(3);

  act(() => {
    result.current[2](false);
  });
  expect(result.current[0]).toBe(4);
  expect(result.current[1]).toBe(3);

  act(() => {
    result.current[2](true);
  });
  expect(result.current[0]).toBe(3);
  expect(result.current[1]).toBe(4);

  act(() => {
    vi.advanceTimersByTime(300);
  });
  expect(result.current[0]).toBe(2);
  expect(result.current[1]).toBe(3);

  act(() => {
    vi.advanceTimersByTime(200);
  });
  expect(result.current[0]).toBe(1);
  expect(result.current[1]).toBe(2);

  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(1);
  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(1);
});

test("different durations", async () => {
  vi.useFakeTimers();

  const { result } = renderHook(() => useDelayedCount([[100, 300]]));
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(undefined);

  act(() => {
    result.current[2](false);
  });
  expect(result.current[0]).toBe(1);
  expect(result.current[1]).toBe(0);

  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(2);
  expect(result.current[1]).toBe(1);

  act(() => {
    result.current[2](true);
  });
  expect(result.current[0]).toBe(1);
  expect(result.current[1]).toBe(2);

  act(() => {
    vi.advanceTimersByTime(100);
  });
  expect(result.current[0]).toBe(1);
  expect(result.current[1]).toBe(2);
  act(() => {
    vi.advanceTimersByTime(200);
  });
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(1);
});

test("skip", async () => {
  vi.useFakeTimers();

  const { result } = renderHook(() => useDelayedCount([100, 200, 300]));
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(undefined);

  act(() => {
    result.current[2](false, true);
  });
  expect(result.current[0]).toBe(4);
  expect(result.current[1]).toBe(0);

  act(() => {
    result.current[2](true, true);
  });
  expect(result.current[0]).toBe(0);
  expect(result.current[1]).toBe(4);
});
