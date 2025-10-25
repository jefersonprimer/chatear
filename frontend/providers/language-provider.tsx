'use client';

import {
  ReactNode,
  createContext,
  useState,
  useEffect,
  useCallback,
} from 'react';
import i18next, { i18n as i18nInstance } from 'i18next';
import { I18nextProvider } from 'react-i18next';
import { usePathname, useRouter } from 'next/navigation';
import initTranslations from '../lib/i18n/i18n';
import i18nConfig from '../lib/i18n/i18n.config';

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
  const [i18n, setI18n] = useState<i18nInstance | null>(null);
  const [language, setLanguage] = useState(locale);
  const router = useRouter();
  const pathname = usePathname();

  const changeLanguage = useCallback(
    (newLanguage: string) => {
      if (i18nConfig.locales.includes(newLanguage) && newLanguage !== language) {
        setLanguage(newLanguage);
        document.cookie = `i18next=${newLanguage};path=/`;
        router.refresh();
      }
    },
    [language, router],
  );

  useEffect(() => {
    const init = async () => {
      const newInstance = await initTranslations(language, i18nConfig.ns);
      setI18n(newInstance);
    };

    if (!i18n || language !== i18n.language) {
      init();
    }
  }, [language, i18n]);

  if (!i18n) {
    return null;
  }

  return (
    <LanguageContext.Provider value={{ language, changeLanguage }}>
      <I18nextProvider i18n={i18n}>{children}</I18nextProvider>
    </LanguageContext.Provider>
  );
}