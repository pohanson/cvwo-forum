"use client";
import useUserContext from "@/lib/useUserContext";

export default function Home() {
  const user = useUserContext();
  return <p>Hi {user?.name}</p>;
}
