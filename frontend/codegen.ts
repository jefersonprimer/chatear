import type { CodegenConfig } from "@graphql-codegen/cli";

const config: CodegenConfig = {
  schema: "http://localhost:8080/graphql",
  documents: ["lib/graphql/**/*.ts"],
  generates: {
    "types/graphql/": {       // pasta dentro do projeto
      preset: "client",
      plugins: [],
      config: {
        withHooks: true
      }
    }
  }
};

export default config;

