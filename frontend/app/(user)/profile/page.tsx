import CardProfile from "./components/CardProfile";

export default function ProfilePage() {
  return (
    <main className="flex flex-col items-center justify-center min-h-screen bg-[#121212] text-white">
      <CardProfile/>
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

        <Link href="/(user)/profile/settings" className="p-3 border rounded-lg hover:bg-gray-100">
          Configuracoes
        </Link>
      </div>
    </main>
  );
}

