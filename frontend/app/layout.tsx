import "./globals.css";
import type { Metadata } from "next";
import { headers } from 'next/headers';

import { AuthProvider } from "../providers/auth-provider";
import { ApolloProviderWrapper } from "../providers/apollo-provider";
import { LanguageProvider } from '../providers/language-provider';

import { Geist, Geist_Mono } from "next/font/google";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export const metadata: Metadata = {
  title: "WeChat The Everyday Everything App",
  description: "super app all apps in one",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const locale = (await headers()).get('x-locale') ?? 'en';

  return (
    <html lang={locale}>
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
        suppressHydrationWarning
      >
        <LanguageProvider locale={locale}>
          <AuthProvider>
            <ApolloProviderWrapper>
              {children}
            </ApolloProviderWrapper>
          </AuthProvider>
        </LanguageProvider>
      </body>
    </html>
  );
}

