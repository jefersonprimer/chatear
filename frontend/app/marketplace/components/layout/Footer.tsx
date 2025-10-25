import React from "react";

export default function Footer() {
  return (
    <footer className="bg-[#232F3E] text-white text-sm">
      {/* Voltar ao topo */}
      <div className="bg-[#37475A] hover:bg-[#485769] text-center py-3 cursor-pointer">
        <span className="font-medium">Voltar ao início</span>
      </div>

      {/* Links principais */}
      <div className="max-w-7xl mx-auto px-4 py-10 grid grid-cols-2 md:grid-cols-4 gap-8 border-b border-gray-600">
        {/* Conheça-nos */}
        <div>
          <h3 className="font-semibold mb-3">Conheça-nos</h3>
          <ul className="space-y-1">
            <li><a href="https://www.aboutamazon.com.br/" className="hover:underline">Sobre a Amazon</a></li>
            <li><a href="/gp/browse.html?node=23490129011" className="hover:underline">Informações corporativas</a></li>
            <li><a href="https://www.amazon.jobs" className="hover:underline">Carreiras</a></li>
            <li><a href="/gp/browse.html?node=15155858011" className="hover:underline">Comunicados à imprensa</a></li>
            <li><a href="/gp/browse.html?node=16593314011" className="hover:underline">Comunidade</a></li>
            <li><a href="/gp/browse.html?node=23454858011" className="hover:underline">Acessibilidade</a></li>
            <li><a href="https://www.amazon.science" className="hover:underline">Amazon Science</a></li>
          </ul>
        </div>

        {/* Ganhe dinheiro conosco */}
        <div>
          <h3 className="font-semibold mb-3">Ganhe dinheiro conosco</h3>
          <ul className="space-y-1">
            <li><a href="https://venda.amazon.com.br" className="hover:underline">Venda na Amazon</a></li>
            <li><a href="https://brandservices.amazon.com.br" className="hover:underline">Proteja e construa a sua marca</a></li>
            <li><a href="https://supply.amazon.com" className="hover:underline">Forneça para a Amazon</a></li>
            <li><a href="https://kdp.amazon.com?language=pt_BR" className="hover:underline">Publique seus livros</a></li>
            <li><a href="https://associados.amazon.com.br" className="hover:underline">Seja um associado</a></li>
            <li><a href="https://advertising.amazon.com/pt-br" className="hover:underline">Anuncie seus produtos</a></li>
          </ul>
        </div>

        {/* Pagamento */}
        <div>
          <h3 className="font-semibold mb-3">Pagamento</h3>
          <ul className="space-y-1">
            <li><a href="/gp/browse.html?node=16568920011" className="hover:underline">Meios de Pagamento</a></li>
            <li><a href="/gp/browse.html?node=24028636011" className="hover:underline">Compre com Pontos</a></li>
            <li><a href="/gp/browse.html?node=24434261011" className="hover:underline">Cartão de Crédito</a></li>
          </ul>
        </div>

        {/* Ajuda */}
        <div>
          <h3 className="font-semibold mb-3">Deixe-nos ajudar você</h3>
          <ul className="space-y-1">
            <li><a href="https://www.amazon.com.br/gp/css/homepage.html" className="hover:underline">Sua conta</a></li>
            <li><a href="/gp/help/customer/display.html?nodeId=201365500" className="hover:underline">Frete e prazo de entrega</a></li>
            <li><a href="/gp/orc/returns/homepage.html" className="hover:underline">Devoluções e reembolsos</a></li>
            <li><a href="/hz/mycd/myx" className="hover:underline">Gerencie seu conteúdo e dispositivos</a></li>
            <li><a href="https://www.amazon.com.br/your-product-safety-alerts" className="hover:underline">Recalls e alertas de segurança</a></li>
            <li><a href="/gp/help/customer/display.html?nodeId=508510" className="hover:underline">Ajuda</a></li>
          </ul>
        </div>
      </div>

      {/* Informações legais */}
      <div className="max-w-7xl mx-auto px-4 py-6 text-gray-300 space-y-2 text-xs">
        <ul className="flex flex-wrap justify-center gap-4 border-b border-gray-700 pb-2">
          <li><a href="/gp/help/customer/display.html?nodeId=201283910" className="hover:underline">Condições de Uso</a></li>
          <li><a href="/gp/help/customer/display.html?nodeId=201283950" className="hover:underline">Notificação de Privacidade</a></li>
          <li><a href="/gp/help/customer/display.html?nodeId=201890250" className="hover:underline">Cookies</a></li>
          <li><a href="/gp/help/customer/display.html?nodeId=201283970" className="hover:underline">Anúncios Baseados em Interesses</a></li>
        </ul>
        <p className="text-center">© 2021–2025 Amazon.com, Inc. ou suas afiliadas</p>
        <p className="text-center">Amazon Serviços de Varejo do Brasil Ltda. | CNPJ 15.436.940/0001-03</p>
        <p className="text-center">
          Av. Juscelino Kubitschek, 2041, Torre E, 18° andar - São Paulo CEP: 04543-011 |
          <a href="https://www.amazon.com.br/gp/help/customer/contact-us" className="hover:underline ml-1">Fale conosco</a> |
          ajuda-amazon@amazon.com.br
        </p>
        <p className="text-center">
          Formas de pagamento aceitas: cartões de crédito (Visa, MasterCard, Elo, American Express),
          cartões de débito (Visa, Elo), Boleto e Pix.
        </p>
      </div>
    </footer>
  );
}

