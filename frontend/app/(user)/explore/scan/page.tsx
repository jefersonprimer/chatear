"use client";
import { useRouter } from "next/navigation";
import CameraScanner from "../../../components/CameraScanner";

export default function ScanPage() {
  const router = useRouter();

  return (
    <div className="relative w-full h-screen bg-black text-white">
      <CameraScanner
        onResult={(data) => {
          console.log("QR detectado:", data);
        }}
      />

      <div className="absolute top-0 left-0 w-full flex items-center justify-between px-4 py-3 bg-black/50 backdrop-blur-sm">
        <button onClick={() => router.back()} className="text-lg">
          ← Voltar
        </button>
        <h1 className="text-center text-base font-medium flex-1">Ler QR Code</h1>
        <button
          onClick={() => router.push("/(user)/profile/qrcode")}
          className="text-sm text-green-400"
        >
          Meu QR
        </button>
      </div>

      <div className="absolute bottom-16 w-full text-center text-gray-300 text-sm">
        Aponte a câmera para um código QR
      </div>
    </div>
  );
}

