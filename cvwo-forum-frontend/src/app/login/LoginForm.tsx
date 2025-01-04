"use client";
import TextInput from "@/components/TextInput";
import { redirect } from "next/navigation";
import { useActionState, useEffect } from "react";
import toast from "react-hot-toast";
import { submitLoginForm } from "./submitLoginForm";

export function LoginForm() {
  const [state, action, pending] = useActionState(submitLoginForm, null);
  useEffect(() => {
    toast.remove();
    if (state == undefined) return;

    if (state.success) {
      redirect("/");
    } else {
      toast.error("Invalid username!");
    }
  }, [state]);
  return (
    <form action={action} className="flex flex-col">
      <TextInput id="username" label="Username" />
      <button
        type="submit"
        disabled={pending}
        className="py-2 mt-4 border border-blue-500 rounded-lg hover:bg-blue-200 disabled:bg-gray-400"
      >
        {pending ? "Logging In ..." : "Login"}
      </button>
    </form>
  );
}
