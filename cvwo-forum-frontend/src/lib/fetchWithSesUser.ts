import { cookies } from "next/headers";

export async function fetchWithSesUser(
  url: string,
  method: "GET" | "POST" | "PATCH" | "DELETE" = "GET"
) {
  const cookieStore = await cookies();
  const requestHeaders = new Headers();
  requestHeaders.set(
    "Cookie",
    `session=${cookieStore.get("session")?.value || ""}`
  );
  return fetch(url, { headers: requestHeaders, method: method });
}
