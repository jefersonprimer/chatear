import React, { useState } from "react";
import { useAuth } from "@/providers/auth-provider";

interface User {
  name: string;
}

interface AccountModalProps {
  user: User;
}

const AccountModal: React.FC<AccountModalProps> = ({ user }) => {
  const { logout } = useAuth();
  const [isOpen, setIsOpen] = useState(false);

  const handleLogout = () => {
    logout();
    // Optionally redirect to login page or home page after logout
  };

  return (
    <div 
      className="relative" 
      onMouseEnter={() => setIsOpen(true)}
      onMouseLeave={() => setIsOpen(false)}
    >
      <div className="flex flex-col p-2 cursor-pointer rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
        <span className="text-xs font-medium">Olá, {user.name}</span>
        <span className="text-sm font-semibold">Contas e listas</span>
      </div>

      {isOpen && (
        <div className="absolute top-full right-0 mt-2 w-[420px] bg-white shadow-xl border border-gray-200 rounded-lg text-sm z-50 text-black">
          {/* Top bar */}
          <div className="flex justify-between items-center bg-blue-50 text-gray-700 px-4 py-2 rounded-t-lg">
            <span className="font-medium">Quem está comprando? Selecione um perfil.</span>
            <button className="text-blue-600 text-sm font-medium hover:underline">
              Gerenciar perfis
            </button>
          </div>

          {/* Content */}
          <div className="flex p-4 gap-8">
            {/* Left Column */}
            <div className="flex-1 border-r border-gray-200 pr-4">
              <h3 className="font-semibold text-gray-800 mb-2">Suas listas</h3>
              <ul className="space-y-1 text-gray-700">
                <li className="hover:text-blue-600 cursor-pointer">Lista de Compras</li>
                <li className="hover:text-blue-600 cursor-pointer">Criar uma Lista de desejos</li>
                <li className="hover:text-blue-600 cursor-pointer">Lista do Bebê</li>
              </ul>
            </div>

            {/* Right Column */}
            <div className="flex-1">
              <h3 className="font-semibold text-gray-800 mb-2">Sua conta</h3>
              <ul className="space-y-1 text-gray-700">
                <li className="hover:text-blue-600 cursor-pointer">Sua conta</li>
                <li className="hover:text-blue-600 cursor-pointer">Seus pedidos</li>
                <li className="hover:text-blue-600 cursor-pointer">Sua Lista de desejos</li>
                <li className="hover:text-blue-600 cursor-pointer">Continuar comprando</li>
                <li className="hover:text-blue-600 cursor-pointer">Recomendados para você</li>
                <li className="hover:text-blue-600 cursor-pointer">Recalls e alertas de segurança do produto</li>
                <li className="hover:text-blue-600 cursor-pointer">Programa e Poupe</li>
                <li className="hover:text-blue-600 cursor-pointer">Sua assinatura Prime</li>
                <li className="hover:text-blue-600 cursor-pointer">Inscrições e assinaturas</li>
                <li className="hover:text-blue-600 cursor-pointer">Biblioteca de conteúdo</li>
                <li className="hover:text-blue-600 cursor-pointer">Dispositivos</li>
                <li className="hover:text-blue-600 cursor-pointer">Seu Prime Video</li>
                <li className="hover:text-blue-600 cursor-pointer">Seu Kindle Unlimited</li>
                <li className="hover:text-blue-600 cursor-pointer">Seu Amazon Photos</li>
                <li className="hover:text-blue-600 cursor-pointer">Seus aplicativos e dispositivos</li>
                <li className="border-t border-gray-200 mt-2 pt-2 hover:text-blue-600 cursor-pointer">
                  Trocar contas
                </li>
                <li className="hover:text-blue-600 cursor-pointer" onClick={handleLogout}>Sair da conta</li>
              </ul>
            </div>
          </div>
        </div>
      )}
    </div>
  );
};

export default AccountModal;

