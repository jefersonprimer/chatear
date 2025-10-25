import { useState } from "react";
import Link from 'next/link';

export default function AccountDropdown() {
  const [isOpen, setIsOpen] = useState(false);

  return (
    <div className="relative" 
      onMouseEnter={() => setIsOpen(true)}
      onMouseLeave={() => setIsOpen(false)}
    >
      <button className="flex flex-col p-2 cursor-pointer rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
        <span className="text-xs font-medium">Olá, faça o seu login</span>
        <span className="text-sm font-semibold">Contas e listas</span>
      </button>

      {isOpen && (
        <div
          className="absolute top-full right-0 w-[512px] bg-white border border-gray-200 shadow-lg mt-2 p-4 z-50 text-black"
        >
          <Link href="/auth/login">
            <button className="w-full cursor-pointer bg-yellow-500 text-white py-2 rounded font-semibold hover:bg-yellow-600">
              Faça seu login        
            </button>
          </Link>
          <p className="text-sm text-gray-600 mt-2">
            Cliente novo? <Link href="/auth/register" className="text-blue-600">Comece aqui</Link>
          </p>

          <div className="flex mt-4">
            <div className="flex-1 pr-4 border-r border-gray-200">
              <h3 className="font-bold mb-2">Suas listas</h3>
              <ul className="text-sm text-gray-700">
                <li><a href="#" className="hover:underline">Criar uma Lista de desejos</a></li>
                <li><a href="#" className="hover:underline">Lista do Bebê</a></li>
              </ul>
            </div>

            <div className="flex-1 pl-4">
              <h3 className="font-bold mb-2">Sua conta</h3>
              <ul className="text-sm text-gray-700 space-y-1">
                <li><a href="#" className="hover:underline">Sua conta</a></li>
                <li><a href="#" className="hover:underline">Seus pedidos</a></li>
                <li><a href="#" className="hover:underline">Sua Lista de desejos</a></li>
                <li><a href="#" className="hover:underline">Continuar comprando</a></li>
                <li><a href="#" className="hover:underline">Recomendados para você</a></li>
                <li><a href="#" className="hover:underline">Programe e Poupe</a></li>
                <li><a href="#" className="hover:underline">Sua assinatura Prime</a></li>
                <li><a href="#" className="hover:underline">Inscrições e assinaturas</a></li>
                <li><a href="#" className="hover:underline">Gerencie seu conteúdo e dispositivos</a></li>
                <li><a href="#" className="hover:underline">Seu Prime Video</a></li>
                <li><a href="#" className="hover:underline">Seu Kindle Unlimited</a></li>
                <li><a href="#" className="hover:underline">Seu Amazon Photos</a></li>
                <li><a href="#" className="hover:underline">Seus aplicativos e dispositivos</a></li>
              </ul>
            </div>
          </div>
        </div>
      )}
    </div>
  );
}
