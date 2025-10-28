import React from 'react';

type Props = {
  title: string;
}

const Header: React.FC = ({ title }: Props) => {
  return (
    <header className="flex items-center justify-between px-4 py-2 bg-[#111111]">
      <div className="flex items-center">
        <h1 className="text-xl font-bold text-[#CFCFCF]">
          {title}
        </h1>
      </div>

      <div className="flex items-center gap-4">
        <button className="cursor-pointer">
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            fill="none" 
            viewBox="0 0 24 24" 
            strokeWidth="1.5" 
            stroke="currentColor" 
            width={24}
            height={24}
          >
            <path strokeLinecap="round" strokeLinejoin="round" d="m21 21-5.197-5.197m0 0A7.5 7.5 0 1 0 5.196 5.196a7.5 7.5 0 0 0 10.607 10.607Z" />
          </svg>
        </button>

        <button className="cursor-pointer">
          <svg 
            xmlns="http://www.w3.org/2000/svg" 
            fill="none" 
            viewBox="0 0 24 24"
            strokeWidth="1.5" 
            stroke="currentColor"
            width={24}
            height={24}
          >
            <path strokeLinecap="round" strokeLinejoin="round" d="M12 9v6m3-3H9m12 0a9 9 0 1 1-18 0 9 9 0 0 1 18 0Z" />
          </svg>
        </button>
      </div>
    </header>
  );
};

export default Header;

