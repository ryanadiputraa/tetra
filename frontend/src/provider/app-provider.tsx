"use client";

import { DashboardLayout } from "@/components";
import { AntdRegistry } from "@ant-design/nextjs-registry";
import "@ant-design/v5-patch-for-react-19";
import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import { ConfigProvider, theme } from "antd";
import idID from "antd/locale/id_ID";

import { useTheme } from "@/hooks";
import { Theme } from "@/types";

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
                boxShadow: "none",
                primaryShadow: "none",
                boxShadowSecondary: "none",
                dangerShadow: "none",
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
