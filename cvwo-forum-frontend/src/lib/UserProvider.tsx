"use client";
import { createContext } from "react";

export const UserContext = createContext<string>("");

export default function UserContextProvider({
  value,
  children,
}: {
  value: string;
  children: React.ReactNode;
}) {
  return <UserContext.Provider value={value}>{children}</UserContext.Provider>;
}
