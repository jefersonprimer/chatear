'use client';

import { useTranslation } from '../hooks/useTranslation';
import Header from '../components/layout/Header';
import Footer from '../components/layout/Footer';
import ServicesCarousel from '../components/services/ServicesCarousel';
import WhatsApp from '../components/icons/WhatsApp';

export default function Home() { 
const services: Service[] = [
  {
    id: "whatsapp",
    name: "WhatsApp",
    icon: <WhatsApp size={30}/>,
  },
  { id: "marketplace", name: "Marketplace", icon: <svg width={28} height={28} viewBox="0 0 24 24"><rect width="24" height="24" rx="4" fill="#F3F4F6" /><path d="M4 7h16v2H4zM6 11h12v7H6z" fill="#111827"/></svg> },
  { id: "wallet", name: "Wallet", icon: <svg width={28} height={28} viewBox="0 0 24 24"><rect width="24" height="24" rx="4" fill="#F3F4F6" /><path d="M3 7h18v10H3z" fill="#111827"/></svg> },
  { id: "uber", name: "Marketplace", icon: <svg width={28} height={28} viewBox="0 0 24 24"><rect width="24" height="24" rx="4" fill="#F3F4F6" /><path d="M4 7h16v2H4zM6 11h12v7H6z" fill="#111827"/></svg> },
  { id: "playstore", name: "Wallet", icon: <svg width={28} height={28} viewBox="0 0 24 24"><rect width="24" height="24" rx="4" fill="#F3F4F6" /><path d="M3 7h18v10H3z" fill="#111827"/></svg> },
];

  const messages = [
    { sender: 'Ana', text: 'Oi! Tudo bem?' },
    { sender: 'Lucas', text: 'Pedido confirmado 👍' },
    { sender: 'Maria', text: 'Podemos conversar amanhã?' },
  ]

  return (
    <>
      <Header/>
      <ServicesCarousel services={services} visibleCount={4} />
      <main className="mx-auto px-4 py-6">
        
        {/* Mensagens */}
        <section className="mt-8">
          <h2 className="text-lg font-semibold mb-3">Mensagens Recentes</h2>
          <div className="space-y-3">
            {messages.map((m, i) => (
              <div
                key={i}
                className="bg-white p-4 rounded-xl shadow-sm hover:shadow-md transition cursor-pointer"
              >
                <p className="font-semibold text-gray-800">{m.sender}</p>
                <p className="text-gray-600 text-sm">{m.text}</p>
              </div>
            ))}
          </div>
        </section> 
      </main>
      
      <Footer/>
    </>
  )
}
