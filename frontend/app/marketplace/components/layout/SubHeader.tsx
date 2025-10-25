import React from "react";

export default function AmazonSubHeader() {
  return (
    <nav className="w-full bg-[#232f3e] text-white text-sm">
      <div className="w-full px-4 flex items-center justify-between">
        {/* "Todos" com ícone */}
        <a
          href="#"
          className="flex items-center gap-1 whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all"
        >
          <span className="font-medium">Todos</span>
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Venda na Amazon
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Mais Vendidos
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Ofertas do Dia
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Prime
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Livros
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Música
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Computadores
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Eletrônicos
        </a>

        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Games
        </a>
        
        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Casa
        </a>
        
        <a href="#" className=" whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          Cuidados Pessoais
        </a>

        {/* Texto à direita */}
        <div className="ml-auto flex items-center text-gray-200 text-xs sm:text-sm  whitespace-nowrap cursor-pointer p-2 rounded-sm hover:outline hover:outline-1 hover:outline-white/40 hover:bg-[#232f3e] transition-all">
          <span className="whitespace-nowrap">
            +1 milhão de títulos para{" "}
            <span className="text-yellow-400 font-medium">ler de graça</span>
          </span>
        </div>
      </div>
    </nav>
  );
}

