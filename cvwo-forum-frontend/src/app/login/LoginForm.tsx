"use client";
import TextInput from "@/components/TextInput";
import { redirect } from "next/navigation";
import { useActionState, useEffect } from "react";
import toast from "react-hot-toast";
import { submitLoginForm } from "./submitLoginForm";
import OutlinedButton from "@/components/OutlinedButton";

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
      <OutlinedButton type="submit" disabled={pending}>
        {pending ? "Logging In ..." : "Login"}
      </OutlinedButton>
      <hr className="w-full border-black mt-6" />
      <OutlinedButton
        type="button"
        disabled={pending}
        onClick={() => redirect("/signup")}
      >
        Go To Signup
      </OutlinedButton>
    </form>
  );
}
