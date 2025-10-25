import { NextRequest, NextResponse } from 'next/server';
import acceptLanguage from 'accept-language';
import i18nConfig from './lib/i18n/i18n.config';

acceptLanguage.languages(i18nConfig.locales);

export const config = {
  matcher: [
    '/((?!api|_next/static|_next/image|assets|favicon.ico|sw.js|site.webmanifest).*)',
  ],
};

const cookieName = 'i18next';

export function proxy(req: NextRequest) {
  let lng: string | undefined | null = req.cookies.get(cookieName)?.value;

  if (!lng) {
    lng = acceptLanguage.get(req.headers.get('Accept-Language'));
  }

  if (req.nextUrl.searchParams.has('hl')) {
    const hl = req.nextUrl.searchParams.get('hl');
    if (i18nConfig.locales.includes(hl!)) {
      lng = hl;
    }
  }

  if (!lng) {
    lng = i18nConfig.defaultLocale;
  }

  const { pathname } = req.nextUrl;
  const accessToken = req.cookies.get('accessToken')?.value;

  if (pathname !== '/') {
    const isAuthPage = pathname.includes('/auth/');
    if (!accessToken && !isAuthPage) {
      const loginUrl = new URL(`/auth/login`, req.url);
      loginUrl.searchParams.set('next', pathname);
      return NextResponse.redirect(loginUrl);
    }
  }

  const response = NextResponse.next();
  response.headers.set('x-locale', lng);
  if (req.cookies.get(cookieName)?.value !== lng) {
    response.cookies.set(cookieName, lng, { path: '/' });
  }

  return response;
}