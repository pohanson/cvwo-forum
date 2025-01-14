"use server";

import { APP_CONFIG } from "@/app/config";
import { CategoryChip } from "@/components/CategoryAutocomplete";
import { fetchWithSesUser } from "@/lib/fetchWithSesUser";
import { z } from "zod";

export default async function submitCreatePost(
  selectedCategory: CategoryChip[],
  prevState: NewPostFormState,
  formData: FormData
): Promise<NewPostFormState> {
  const parsedFormData = {
    title: formData.get("title")?.toString(),
    content: formData.get("content")?.toString(),
    categoryList: selectedCategory,
  };
  const validatedFields = NewPostFormSchema.safeParse(parsedFormData);
  if (!validatedFields.success) {
    return {
      success: false,
      errors: Object.values(validatedFields.error.flatten().fieldErrors).flat(),
      payload: parsedFormData,
    };
  }

  const r = await fetchWithSesUser(`${APP_CONFIG.backendUrl}/post`, {
    method: "POST",
    body: JSON.stringify(parsedFormData),
  });
  if (r.status == 200) {
    return { success: true };
  } else {
    return {
      success: false,
      errors: [await r.text()],
      payload: parsedFormData,
    };
  }
}
type NewPostFormState =
  | {
      success: true;
    }
  | {
      success: false;
      errors: string[];
      payload: {
        title?: string;
        content?: string;
        categoryList: CategoryChip[];
      };
    }
  | null;
const NewPostFormSchema = z.object({
  title: z.string().min(1, "Title cannot be empty"),
  content: z.string().min(1, "Content cannot be empty"),
  categoryList: z
    .array(
      z.object({
        id: z.number(),
        title: z.string(),
        description: z.string().optional(),
      })
    )
    .min(1, "At least one category must be selected"),
});
