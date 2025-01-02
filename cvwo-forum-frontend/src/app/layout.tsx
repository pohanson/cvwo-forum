import type { Metadata } from "next";
import "./globals.css";

export const metadata: Metadata = {
  title: "CVWO Forum",
  description: "the web forum created for volunteers.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en">
      <body>{children}</body>
    </html>
  );
}
