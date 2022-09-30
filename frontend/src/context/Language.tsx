import { createContext, useContext } from "react";

export type Language = "en" | "ar" | "fr";
export interface ILanguage {
  language: Language;
  setLanguage: (val: Language | ((prevState: Language) => Language)) => void;
}

const LanguageContext = createContext<ILanguage>({
  language: "en",
  setLanguage: () => {},
});
export const LanguageProvider = LanguageContext.Provider;
export const useLanguage = () => useContext(LanguageContext);
