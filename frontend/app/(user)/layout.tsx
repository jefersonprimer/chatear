import React from "react";
import BottomBar from "./components/layout/BottomBar";

export default function UserLayout({ children }: { children: React.ReactNode }) {
  return (
    <div className="flex flex-col min-h-screen">
      <main className="flex-1">{children}</main>
      <BottomBar />
    </div>
  );
}

