"use client";
import TextInput from "@/components/TextInput";
import { useActionState, useEffect } from "react";
import { submitSignupform } from "./submitSignupform";
import { useRouter } from "next/navigation";
import toast from "react-hot-toast";

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
      <button
        type="submit"
        disabled={pending}
        className="py-2 mt-4 border border-blue-500 rounded-lg hover:bg-blue-200 disabled:bg-gray-400"
      >
        {pending ? "Signing Up ..." : "Sign Up"}
      </button>
    </form>
  );
}
