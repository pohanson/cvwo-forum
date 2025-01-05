"use server";
import { z } from "zod";
import { APP_CONFIG } from "../config";
import writeSetCookies from "@/lib/writeSetCookie";
import { cookies } from "next/headers";

export async function submitLoginForm(
  formState: SigninFormState,
  formData: FormData
): Promise<SigninFormState> {
  const parsedFormData = {
    username: formData.get("username")?.toString(),
  };
  const validatedFields = SigninFormSchema.safeParse(parsedFormData);
  if (!validatedFields.success) return { success: false };
  const r = await fetch(`${APP_CONFIG.backendUrl}/login`, {
    method: "POST",
    body: JSON.stringify(validatedFields.data),
  });
  if (r.status != 200) {
    return { success: false };
  } else {
    writeSetCookies(r, await cookies());
    return { success: true };
  }
}

const SigninFormSchema = z.object({
  username: z
    .string()
    .min(1, "Username cannot be empty")
    .max(80, "Must be 80 or fewer characters"),
});
type SigninFormState = null | { success: boolean };
