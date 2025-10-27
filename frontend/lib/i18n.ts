import i18next from 'i18next'
import { initReactI18next } from 'react-i18next'
import LanguageDetector from 'i18next-browser-languagedetector'
import resourcesToBackend from 'i18next-resources-to-backend'

i18next
  .use(initReactI18next)
  .use(LanguageDetector)
  .use(resourcesToBackend((lng, ns) => import(`../locales/${lng}/${ns}.json`)))
  .init({
    fallbackLng: 'en',
    supportedLngs: ['en', 'pt'],
    ns: ['common'],
    defaultNS: 'common',
    detection: {
      order: ['cookie', 'localStorage', 'navigator'],
      caches: ['cookie']
    },
    interpolation: {
      escapeValue: false
    }
  })

export default i18next

