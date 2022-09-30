import { useEffect, useState } from "react";
import {
  ColorScheme,
  ColorSchemeProvider,
  createEmotionCache,
  MantineProvider,
} from "@mantine/core";
import { ModalsProvider } from "@mantine/modals";
import { NotificationsProvider } from "@mantine/notifications";
import rtlPlugin from "stylis-plugin-rtl";
import i18n from "@/i18n";
import { useDidUpdate, useLocalStorage } from "@mantine/hooks";
import { Language, LanguageProvider } from "@/context";

interface IAppProvider {
  children: React.ReactNode;
}

export function AppProvider({ children }: IAppProvider) {
  const [rtl, setRtl] = useState(false);
  const [language, setLanguage] = useState<Language>("en");
  const rtlCache = createEmotionCache({
    key: "mantine-rtl",
    stylisPlugins: [rtlPlugin],
  });
  const [colorScheme, setColorScheme] = useLocalStorage<ColorScheme>({
    key: "color-scheme",
    defaultValue: "light",
  });

  const toggleColorScheme = (value?: ColorScheme) =>
    value
      ? setColorScheme(value)
      : colorScheme === "dark"
      ? setColorScheme("light")
      : setColorScheme("dark");

  useDidUpdate(() => {
    i18n.changeLanguage(language);
    if (language === "ar" && !rtl) {
      setRtl(true);
    } else {
      setRtl(false);
    }
  }, [language]);

  return (
    <LanguageProvider value={{ language: "en", setLanguage }}>
      <ColorSchemeProvider
        colorScheme={colorScheme}
        toggleColorScheme={toggleColorScheme}
      >
        <MantineProvider
          theme={{
            colorScheme,
            dir: rtl ? "rtl" : "ltr",
            cursorType: "pointer",
          }}
          withGlobalStyles
          withNormalizeCSS
          emotionCache={rtl ? rtlCache : undefined}
        >
          <NotificationsProvider>
            <ModalsProvider>{children}</ModalsProvider>
          </NotificationsProvider>
        </MantineProvider>
      </ColorSchemeProvider>
    </LanguageProvider>
  );
}

export default AppProvider;
