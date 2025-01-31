"use client";

import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConfigProvider } from "antd";
import idID from "antd/locale/id_ID";

export default function AppProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <ConfigProvider
        locale={idID}
        theme={{
          token: {
            colorPrimary: "#4682ab",
            borderRadius: 8,
          },
          components: {
            Button: {
              controlHeightLG: 45,
            },
            Input: {
              controlHeightLG: 45,
            },
          },
        }}
      >
        {children}
      </ConfigProvider>
    </QueryClientProvider>
  );
}
