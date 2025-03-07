"use client";

import { DashboardLayout } from "@/components";
import { AntdRegistry } from "@ant-design/nextjs-registry";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConfigProvider, theme } from "antd";
import idID from "antd/locale/id_ID";

import { Theme } from "@/types";
import { useTheme } from "@/hooks";

interface Props {
  initialTheme: Theme;
  children: React.ReactNode;
}

export const AppProvider = ({ initialTheme, children }: Props) => {
  const queryClient = new QueryClient();
  const { theme: appTheme, toggleTheme } = useTheme(initialTheme);

  return (
    <QueryClientProvider client={queryClient}>
      <AntdRegistry>
        <ConfigProvider
          locale={idID}
          theme={{
            algorithm:
              appTheme === "light"
                ? theme.defaultAlgorithm
                : theme.darkAlgorithm,
            token: {
              colorPrimary: appTheme === "light" ? "#3E6071" : "#759EB3",
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
          <DashboardLayout
            theme={appTheme}
            toggleThemeAction={() => toggleTheme()}
          >
            {children}
          </DashboardLayout>
        </ConfigProvider>
      </AntdRegistry>
    </QueryClientProvider>
  );
};
