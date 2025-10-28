"use client";
import { useEffect, useRef } from "react";
import jsQR from "jsqr";

interface CameraScannerProps {
  onResult?: (data: string) => void;
}

export default function CameraScanner({ onResult }: CameraScannerProps) {
  const videoRef = useRef<HTMLVideoElement | null>(null);

  useEffect(() => {
    let stream: MediaStream;

    const startCamera = async () => {
      try {
        stream = await navigator.mediaDevices.getUserMedia({
          video: { facingMode: "environment" },
        });
        if (videoRef.current) {
          videoRef.current.srcObject = stream;
          await videoRef.current.play();
        }
      } catch (err) {
        console.error("Erro ao acessar a câmera:", err);
      }
    };

    startCamera();

    return () => {
      if (stream) stream.getTracks().forEach((t) => t.stop());
    };
  }, []);

  // Loop para ler QR Codes
  useEffect(() => {
    const canvas = document.createElement("canvas");
    const ctx = canvas.getContext("2d");

    const interval = setInterval(() => {
      const video = videoRef.current;
      if (!video || !ctx) return;

      canvas.width = video.videoWidth;
      canvas.height = video.videoHeight;
      ctx.drawImage(video, 0, 0, canvas.width, canvas.height);
      const imageData = ctx.getImageData(0, 0, canvas.width, canvas.height);
      const code = jsQR(imageData.data, canvas.width, canvas.height);

      if (code) {
        onResult?.(code.data);
      }
    }, 400);

    return () => clearInterval(interval);
  }, [onResult]);

  return (
    <div className="relative w-full h-full overflow-hidden">
      <video
        ref={videoRef}
        className="absolute inset-0 w-full h-full object-cover"
        playsInline
        muted
      />

      {/* Máscara escura com área central transparente */}
      <div className="absolute inset-0 bg-black/70">
        <div className="absolute left-1/2 top-1/2 w-64 h-64 -translate-x-1/2 -translate-y-1/2 rounded-lg border-4 border-green-500 shadow-[0_0_20px_#00FF00] overflow-hidden">
          {/* Linha animada que se move verticalmente */}
          <div className="absolute top-0 left-0 w-full h-1 bg-green-400 animate-scan" />
        </div>
      </div>

      {/* Estilo da animação */}
      <style jsx>{`
        @keyframes scan {
          0% {
            top: 0%;
          }
          50% {
            top: 95%;
          }
          100% {
            top: 0%;
          }
        }

        .animate-scan {
          animation: scan 2s ease-in-out infinite;
        }
      `}</style>
    </div>
  );
}

