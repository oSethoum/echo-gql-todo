import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import { HomePage } from "@/ui";
import {
  createClient,
  Provider,
  defaultExchanges,
  subscriptionExchange,
} from "urql";
import { createContext, useMemo, useState } from "react";
import { useLocalStorage } from "@mantine/hooks";
import { useMantineColorScheme } from "@mantine/core";
import { createClient as createWSClient } from "graphql-ws";

interface IAppContext {
  url: string;
  token: string;
  setToken: (val: string | ((prevState: string) => string)) => void;
}

export const AppContext = createContext<IAppContext>({
  url: "",
  token: "",
  setToken: () => {},
});

function App() {
  const [url, setUrl] = useState("localhost:3001");
  const { toggleColorScheme } = useMantineColorScheme();
  const [token, setToken] = useLocalStorage({
    key: "jwtToken",
    defaultValue: "",
  });

  const wsClient = useMemo(
    () =>
      createWSClient({
        url: "ws://localhost:3001/subscriptions",
        connectionParams: { token: "Bearer " + token },
      }),
    []
  );

  const client = useMemo(
    () =>
      createClient({
        url: "http://" + url + "/query",
        fetchOptions: {
          headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + token,
          },
        },
        exchanges: [
          ...defaultExchanges,
          subscriptionExchange({
            forwardSubscription: (operation) => ({
              subscribe: (sink) => ({
                unsubscribe: wsClient.subscribe(operation, sink),
              }),
            }),
          }),
        ],
      }),
    [token]
  );

  return (
    <AppContext.Provider value={{ url, token, setToken }}>
      <Provider value={client}>
        <Router>
          <Routes>
            <Route path="/" element={<HomePage />} />
          </Routes>
        </Router>
      </Provider>
    </AppContext.Provider>
  );
}

export default App;
