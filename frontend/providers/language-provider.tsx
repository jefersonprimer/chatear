'use client';

import {
  ReactNode,
  createContext,
  useState,
  useEffect,
  useCallback,
} from 'react';
import i18n from '../lib/i18n';
import { I18nextProvider } from 'react-i18next';
import { useRouter } from 'next/navigation';

interface LanguageContextType {
  language: string;
  changeLanguage: (newLanguage: string) => void;
}

export const LanguageContext = createContext<LanguageContextType | undefined>(
  undefined,
);

interface LanguageProviderProps {
  children: ReactNode;
  locale: string;
}

export function LanguageProvider({ children, locale }: LanguageProviderProps) {
  const [language, setLanguage] = useState(locale);
  const router = useRouter();

  const changeLanguage = useCallback(
    (newLanguage: string) => {
      if (['en', 'pt'].includes(newLanguage) && newLanguage !== language) {
        setLanguage(newLanguage);
        i18n.changeLanguage(newLanguage);
        document.cookie = `i18next=${newLanguage};path=/`;
        router.refresh();
      }
    },
    [language, router],
  );

  useEffect(() => {
    if (i18n.language !== language) {
      i18n.changeLanguage(language);
    }
  }, [language]);

  return (
    <LanguageContext.Provider value={{ language, changeLanguage }}>
      <I18nextProvider i18n={i18n}>{children}</I18nextProvider>
    </LanguageContext.Provider>
  );
}