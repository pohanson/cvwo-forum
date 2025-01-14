"use server";

import { APP_CONFIG } from "@/app/config";
import { CategoryChip } from "./types/CategoryChip";

export async function GetAllCategory(): Promise<CategoryChip[]> {
  return await fetch(`${APP_CONFIG.backendUrl}/category`).then((r) => r.json());
}
