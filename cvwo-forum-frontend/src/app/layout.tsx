import type { Metadata } from "next";
import "./globals.css";
import { Toaster } from "react-hot-toast";
import UserContextProvider from "@/lib/UserProvider";
import getSesUser from "@/lib/getSesUser";

export const metadata: Metadata = {
  title: "CVWO Forum",
  description: "The web forum created for volunteers.",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const user = await getSesUser();
  return (
    <html lang="en">
      <body>
        <Toaster />
        <UserContextProvider value={user == null ? "" : user.toJsonStr()}>
          {children}
        </UserContextProvider>
      </body>
    </html>
  );
}
