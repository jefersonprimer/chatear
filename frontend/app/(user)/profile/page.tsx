"use client";

import Link from "next/link";
import CardProfile from "./components/CardProfile";
import { useQuery } from "@apollo/client/react";
import { ME_QUERY } from "@/lib/graphql/queries/me";

export default function ProfilePage() {
  const { data, loading, error } = useQuery(ME_QUERY);

  if (loading) return <p>Loading...</p>;
  if (error) return <p>Error: {error.message}</p>;

  const user = data?.me;

  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-[#121212] text-white">
      <CardProfile imageUrl={user?.avatarURL} name={user?.name || 'Loading...'} />
      <div className="flex flex-col gap-3 p-4">
        <Link href="/(user)/profile/services" className="p-3 border rounded-lg hover:bg-gray-100">
          Pagamentos e Servicos
        </Link>

        <Link href="/(user)/profile/favorites" className="p-3 border rounded-lg hover:bg-gray-100">
          Favoritos
        </Link>

        <Link href="/(user)/profile/my-publications" className="p-3 border rounded-lg hover:bg-gray-100">
          Minhas Publicacoes
        </Link>

        <Link href="/(user)/profile/sticker-gallery" className="p-3 border rounded-lg hover:bg-gray-100">
          Galeria de Stickers
        </Link>

        <Link href="/profile/settings" className="p-3 border rounded-lg hover:bg-gray-100">
          Configuracoes
        </Link>
      </div>
    </main>
  );
}

