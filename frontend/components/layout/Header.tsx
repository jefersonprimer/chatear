'use client'

import Image from 'next/image';
import Link from 'next/link';
import logo from '../../public/logo.png';

export default function Header() {
  return (
    <header className="w-full bg-white shadown-sm">
      <div className="mx-auto flex items-center justify-between px-4 py-3">
        
        <Link href="/">
          <div className="flex items-center gap-2 cursor-pointer">
            <Image
              src={logo}
              alt="SuperApp"
              className="object-contain brightness-0 invert-[0.53] sepia-[1] saturate-[10000%] hue-rotate-[80deg]"
              width={100}
              height={80}
            />
          </div> 
        </Link>
                 
        {/* Saldo e Avatar */}
        <div className="flex items-center gap-4">
          <div className="bg-green-100 text-green-700 px-3 py-1 rounded-full text-sm font-semibold">
            R$ 245,90
          </div>
          <img
            src="https://i.pravatar.cc/40"
            alt="Avatar"
            className="w-9 h-9 rounded-full border-2 border-gray-300 cursor-pointer"
          />
        </div>
      </div>
    </header>
  )
}

