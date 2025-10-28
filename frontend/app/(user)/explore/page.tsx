"use client";

import { useState } from "react";
import Header from "../components/layout/Header";
import CameraScanner from "../components/CameraScanner";
import Link from "next/link";

export default function ExplorePage() {
  const [showScanner, setShowScanner] = useState(false);

  return (
    <>
      <Header title="Descobrir"/> 
      <main className="flex flex-col items-center justify-center min-h-screen bg-[#121212] text-white">
        <div className="flex flex-col gap-3 p-4">
          <h1 className="text-xl font-semibold mb-4">Explorar</h1>

          <Link href="/(user)/explore/moments" className="p-3 border rounded-lg hover:bg-gray-100">
            Momentos
          </Link>

          <Link href="/(user)/explore/channels" className="p-3 border rounded-lg hover:bg-gray-100">
            Canal
          </Link>

          <button
            onClick={() => setShowScanner(true)}
            className="p-3 border rounded-lg hover:bg-gray-100"
          >
            Ler (QR Code)
          </button>          

          <Link href="/(user)/explore/stories" className="p-3 border rounded-lg hover:bg-gray-100">
            Principais Histórias
          </Link>

          <Link href="/(user)/explore/search" className="p-3 border rounded-lg hover:bg-gray-100">
            Pesquisar
          </Link>

          <Link href="/(user)/explore/nearby" className="p-3 border rounded-lg hover:bg-gray-100">
            Olhar ao Redor
          </Link>

          {showScanner && (
            <CameraScanner
              onClose={() => setShowScanner(false)}
              onResult={(data) => console.log("QR detectado:", data)}
            />
          )}
        </div>
      </main>
    </>
  );
}

