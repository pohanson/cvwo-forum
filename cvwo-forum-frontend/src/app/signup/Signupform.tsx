"use client";
import TextInput from "@/components/TextInput";
import { useActionState, useEffect } from "react";
import { submitSignupform } from "./submitSignupform";
import { redirect, useRouter } from "next/navigation";
import toast from "react-hot-toast";
import OutlinedButton from "@/components/OutlinedButton";

export function SignupForm() {
  const [state, action, pending] = useActionState(submitSignupform, null);
  const router = useRouter();

  useEffect(() => {
    toast.dismiss();
    if (state == null) return;

    switch (state?.success) {
      case true:
        toast.success("Sign Up Successful");
        router.replace("/");
        break;
      case false:
        state.errors.map((err) => toast.error(err));
        break;
    }
  }, [state, router]);

  return (
    <form action={action} className="flex flex-col">
      <TextInput
        id="username"
        label="Username"
        defaultValue={state?.success ? "" : state?.payload.username}
      />
      <TextInput
        id="name"
        label="Name"
        defaultValue={state?.success ? "" : state?.payload.name}
      />
      <OutlinedButton type="submit" disabled={pending}>
        {pending ? "Signing Up ..." : "Sign Up"}
      </OutlinedButton>
      <hr className="w-full border-black mt-6" />
      <OutlinedButton
        type="button"
        disabled={pending}
        onClick={() => redirect("/login")}
      >
        Go To Login
      </OutlinedButton>
    </form>
  );
}
