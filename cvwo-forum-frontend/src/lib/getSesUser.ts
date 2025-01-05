"use server";

import { APP_CONFIG } from "@/app/config";
import { fetchWithSesUser } from "./fetchWithSesUser";
import { User } from "./types/User";

export default async function getSesUser(): Promise<User | null> {
  const r = await fetchWithSesUser(`${APP_CONFIG.backendUrl}/verifyUser`);
  const rt = await r.text();
  if (r.status == 200) {
    return User.fromJsonStr(rt);
  } else {
    if (rt.trim() != "User Session Error") {
      console.error(rt);
    }
    return null;
  }
}
