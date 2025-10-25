"use client";

import { ApolloProvider } from "@apollo/client/react";
import client from "@/lib/graphql/client";
import { ReactNode } from "react";

export function ApolloProviderWrapper({ children }: { children: ReactNode }) {
  return <ApolloProvider client={client}>{children}</ApolloProvider>;
}
