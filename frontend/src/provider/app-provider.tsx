"use client";

import { AntdRegistry } from "@ant-design/nextjs-registry";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConfigProvider } from "antd";
import idID from "antd/locale/id_ID";

export const AppProvider = ({ children }: { children: React.ReactNode }) => {
  const queryClient = new QueryClient();

  return (
    <QueryClientProvider client={queryClient}>
      <AntdRegistry>
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
      </AntdRegistry>
    </QueryClientProvider>
  );
};
