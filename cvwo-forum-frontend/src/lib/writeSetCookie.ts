"use server";

import { parseSetCookie } from "next/dist/compiled/@edge-runtime/cookies";
import { ReadonlyRequestCookies } from "next/dist/server/web/spec-extension/adapters/request-cookies";

export default async function writeSetCookies(
  r: Response,
  cookieStore: ReadonlyRequestCookies
) {
  r.headers
    .getSetCookie()
    .map(parseSetCookie)
    .filter((val) => val != undefined)
    .map((val) => cookieStore.set(val));
}
