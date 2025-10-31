"use client";

import React, { useEffect } from "react";
import { useRouter } from "next/navigation";
import BottomBar from "./components/layout/BottomBar";
import { useAuth } from "@/providers/auth-provider";

export default function UserLayout({ children }: { children: React.ReactNode }) {
  const { isLoggedIn, loading } = useAuth();
  const router = useRouter();

  useEffect(() => {
    if (!loading && !isLoggedIn) {
      router.push("/login");
    }
  }, [isLoggedIn, loading, router]);

  if (loading || !isLoggedIn) {
    return <p>Loading...</p>;
  }

  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex-1">{children}</main>
      <BottomBar />
    </div>
  );
}