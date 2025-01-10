"use server";
import { cookies } from "next/headers";

export async function fetchWithSesUser(
  url: string,
  requestInit: RequestInit = {}
) {
  const cookieStore = await cookies();
  const requestHeaders = new Headers(requestInit.headers);
  requestHeaders.set(
    "Cookie",
    `session=${cookieStore.get("session")?.value || ""}`
  );

  return fetch(url, { ...requestInit, headers: requestHeaders });
}
