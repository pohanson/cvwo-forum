import { NextRequest, NextResponse } from "next/server";
import getSesUser from "./lib/getSesUser";

export default async function authmiddleware(request: NextRequest) {
  const whitelist = ["/signup", "/login"];
  if (whitelist.includes(request.nextUrl.pathname)) {
    return;
  }
  const user = await getSesUser();
  if (user == null) {
    return NextResponse.redirect(new URL("/login", request.url));
  }
}

export const config = {
  matcher: [
    "/((?!api|_next/static|_next/image|favicon.ico|sitemap.xml|robots.txt).*)",
  ],
};
