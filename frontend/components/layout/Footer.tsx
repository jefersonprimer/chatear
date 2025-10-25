export default function Footer() {
  return (
    <footer className="w-full bg-white border-t mt-8">
      <div className="max-w-[1400px] mx-auto px-4 py-6 text-center text-sm text-gray-500">
        <p>© {new Date().getFullYear()} SuperApp. Todos os direitos reservados.</p>
        <div className="flex justify-center gap-4 mt-2">
          <a href="#" className="hover:underline">Termos de uso</a>
          <a href="#" className="hover:underline">Privacidade</a>
          <a href="#" className="hover:underline">Ajuda</a>
        </div>
      </div>
    </footer>
  )
}

