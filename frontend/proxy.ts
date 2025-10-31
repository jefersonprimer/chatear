import { NextRequest, NextResponse } from "next/server";
import { verifyAuth } from "@/lib/auth";

export async function proxy(req: NextRequest) {
  const token = req.cookies.get("session")?.value;
  const user = token ? await verifyAuth(token) : null;

  const { pathname } = req.nextUrl;

  // --- Páginas públicas ---
  const publicRoutes = [
    "/auth/login",
    "/auth/register",
    "/auth/forgot-password",
    "/auth/recover-account",
    "/auth/verify-email",
    "/auth/sucess",
    "/auth/policies",
  ];

  const isPublic = publicRoutes.some((route) => pathname.startsWith(route));

  // --- Usuário autenticado tentando acessar páginas públicas ---
  if (user && isPublic) {
    return NextResponse.redirect(new URL("/", req.url));
  }

  // --- Usuário não autenticado tentando acessar páginas protegidas ---
  if (!user && !isPublic) {
    return NextResponse.redirect(new URL("/auth/login", req.url));
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    "/((?!api|_next/static|_next/image|favicon.ico).*)",
  ],
};
;
