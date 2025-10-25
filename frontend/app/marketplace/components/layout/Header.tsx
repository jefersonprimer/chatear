"use client";

import AccountDropdown from "../modals/AccountDropdown";
import AccountModal from "../modals/AccountModal";
import { useAuth } from "@/providers/auth-provider";
import Link from 'next/link';
import { useQuery } from "@apollo/client/react/hooks";
import { ME_QUERY } from "@/lib/graphql/queries/me";

export default function Header() {
  const { isLoggedIn } = useAuth();
  const { data: userData, loading, error } = useQuery(ME_QUERY, {
    skip: !isLoggedIn,
  });

  console.log("isLoggedIn:", isLoggedIn);
  console.log("userData:", userData);
  console.log("loading:", loading);
  console.log("error:", error);

  return (
    <header className="w-full h-14 bg-[#131921] text-white flex items-center">
      <div className="w-full px-4 flex items-center justify-between">
        {/* "Logo" como texto */}
        <Link
          href="/"
          className="text-2xl font-bold tracking-tight whitespace-nowrap p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all"
        >
          WeChat.com.br
        </Link>

        <div className="flex flex-col whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          <span className="text-xs font-medium">A entrega sera feita em Frederico... 98400000</span>
          <span className="text-sm font-semibold">Atualizar CEP</span>
        </div>

        {/* Barra de pesquisa */}
        <div className="flex flex-1 justify-center rounded max-w-lg mx-2 border focus-within:border-2 focus-within:border-[#FF9900] transition-all">
          <span className="flex flex-rol items-center text-sm font-medium px-2 py-2 bg-[#E6E6E6] text-[#6F7373] hover:text-[#000] cursor-pointer">
            Todos
            <svg className="w-4 h-4" xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" data-t="dropdown-svg" aria-hidden="true" role="img" fill="currentColor">
              <path d="M7 10h10l-5 5z">
              </path>
            </svg>
          </span>
          <input
            type="text"
            placeholder="Pesquisar WeChat.com"
            className="w-full px-4 py-2 text-black bg-[#fff] focus:outline-none placeholder-[#6F7373]"
          />
          <button className="bg-[#febd69] px-4 flex items-center justify-center hover:bg-[#f3a847] transition cursor-pointer">
            <svg
              className="w-4 h-4"
              xmlns="http://www.w3.org/2000/svg"
              viewBox="0 0 24 24"
              aria-hidden="true"
            >
              <path
                fillRule="evenodd"
                clipRule="evenodd"
                d="M10.5 19C12.4879 19 14.3164 18.3176 15.7641 17.1742L21.2927 22.7069L22.7074 21.2931L17.1778 15.7595C18.319 14.3126 19 12.4858 19 10.5C19 5.80558 15.1944 2 10.5 2C5.80558 2 2 5.80558 2 10.5C2 15.1944 5.80558 19 10.5 19ZM10.5 17C14.0899 17 17 14.0899 17 10.5C17 6.91015 14.0899 4 10.5 4C6.91015 4 4 6.91015 4 10.5C4 14.0899 6.91015 17 10.5 17Z"
              />
            </svg>
          </button>
        </div>

        {/* Menu lateral */}
        <div className="flex items-center space-x-3 text-sm">
          {isLoggedIn && userData ? <AccountModal user={userData.me} /> : <AccountDropdown />}
          <button className="flex flex-col p-2 cursor-pointer rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
            <span className="text-xs font-medium">Devoluções</span>
            <span className="text-sm font-semibold">e Pedidos</span>
          </button>
          <div className="flex flex-col items-center space-x-1 px-2 py-1 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all cursor-pointer">
            <span className="font-semibold">🛒</span>
            <span>Carrinho</span>
          </div>
        </div>
      </div>
    </header>
  );
}
