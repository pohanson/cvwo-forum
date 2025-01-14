"use client";
import { Button } from "@mui/material";
import {
  Dispatch,
  SetStateAction,
  useActionState,
  useEffect,
  useState,
} from "react";
import submitCreatePost from "./submitCreatePost";
import toast from "react-hot-toast";
import { useRouter } from "next/navigation";
import { CategoryChip } from "@/lib/types/CategoryChip";
import CategoryAutocomplete from "@/components/CategoryAutocomplete";

export default function CreatePostPage() {
  const [selectedCategory, setSelectedCategory] = useState<CategoryChip[]>([]);
  const router = useRouter();
  const submitWithCategory = submitCreatePost.bind(null, selectedCategory);
  const [state, action, pending] = useActionState(submitWithCategory, null);
  useEffect(() => {
    toast.dismiss();
    if (state == null) return;
    switch (state.success) {
      case true:
        toast.success("New Post Created");
        setSelectedCategory([]);
        router.replace("/");
        break;
      case false:
        state.errors.map((err) => toast.error(err));
        setSelectedCategory(state.payload.categoryList);
        break;
    }
  }, [state, router]);
  return (
    <form className="flex flex-col w-full p-8" action={action}>
      <div className="flex justify-between">
        <h1>Create New Post</h1>
        <Button variant="contained" type="submit">
          {pending ? "Creating Post..." : "Create Post"}
        </Button>
      </div>
      <input
        id="title"
        name="title"
        placeholder="Title"
        className="w-full p-4 my-4 text-lg font-bold rounded-xl"
        defaultValue={state?.success ? "" : state?.payload.title}
      />
      <CategoryInput
        selectedCategory={selectedCategory}
        setSelectedCategory={setSelectedCategory}
      />

      <textarea
        name="content"
        placeholder="Content"
        className="w-full h-full p-2 my-4 rounded-xl"
        defaultValue={state?.success ? "" : state?.payload.content}
      />
    </form>
  );
}

function CategoryInput({
  selectedCategory,
  setSelectedCategory,
}: {
  selectedCategory: CategoryChip[];
  setSelectedCategory: Dispatch<SetStateAction<CategoryChip[]>>;
}) {
  return (
    <div className="flex flex-col w-full p-2 bg-white rounded-md">
      <CategoryAutocomplete
        selectedCategory={selectedCategory}
        setSelectedCategory={setSelectedCategory}
      />
    </div>
  );
}
