"use client";
import { useContext } from "react";
import { UserContext } from "./UserProvider";
import { User } from "./types/User";

export default function useUserContext(): User | null {
  const userJsonStr = useContext(UserContext);
  return userJsonStr == "" ? null : User.fromJsonStr(userJsonStr);
}
