"use server";

import { z } from "zod";
import { APP_CONFIG } from "../config";
import { cookies } from "next/headers";
import writeSetCookies from "@/lib/writeSetCookie";

export async function submitSignupform(
  formState: SignupFormState,
  formData: FormData
): Promise<SignupFormState> {
  const parsedFormData = {
    username: formData.get("username")?.toString(),
    name: formData.get("name")?.toString(),
  };
  const validatedFields = SignupFormSchema.safeParse(parsedFormData);
  if (!validatedFields.success) {
    return {
      success: false,
      errors: Object.values(validatedFields.error.flatten().fieldErrors).flat(),
      payload: parsedFormData,
    };
  }

  const r = await fetch(`${APP_CONFIG.backendUrl}/user`, {
    method: "POST",
    body: JSON.stringify(validatedFields.data),
  }).catch((e) => {
    console.error(e);
    return null;
  });
  if (r == null) {
    return {
      success: false,
      errors: ["Server Error"],
      payload: parsedFormData,
    };
  }
  if (r.status != 200) {
    const errors = await r
      .text()
      .then((val) => val.split("\n").filter((val) => val != ""))
      .catch((e) => {
        console.error(e);
        return ["Unknown Error"];
      });
    return {
      success: false,
      errors: errors,
      payload: parsedFormData,
    };
  } else {
    writeSetCookies(r, await cookies());
    return { success: true, message: await r.text() };
  }
}

type SignupFormState =
  | { success: true; message?: string }
  | {
      success: false;
      errors: string[];
      payload: { username?: string; name?: string };
    }
  | null;
const SignupFormSchema = z.object({
  username: z
    .string()
    .min(1, "Username cannot be empty")
    .max(80, "Username must be 80 or fewer characters"),
  name: z
    .string()
    .min(1, "Name cannot be empty")
    .max(80, "Name must be 80 or fewer characters"),
});
