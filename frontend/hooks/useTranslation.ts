
import { useContext } from 'react';
import { useTranslation as useNextTranslation } from 'react-i18next';
import { LanguageContext } from '../components/language-provider';

export function useTranslation() {
  const context = useContext(LanguageContext);

  if (context === undefined) {
    throw new Error('useLanguage must be used within a LanguageProvider');
  }

  return { ...useNextTranslation(), ...context };
}
