import { DashboardLayout } from "@/components";
import { AppProvider, AuthProvider } from "@/provider";
import "@ant-design/v5-patch-for-react-19";
import type { Metadata } from "next";
import { Raleway } from "next/font/google";
import "./globals.css";

const font = Raleway({
  subsets: ["latin"],
  weight: ["300", "400", "500", "600", "700"],
});

export const metadata: Metadata = {
  title: "Inventra",
  description: "Sistem management asset perusahaan",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  return (
    <html lang="en" className="h-full">
      <body className={`${font} antialiased bg-slate-100 text-black h-full`}>
        <AuthProvider>
          <AppProvider>
            <DashboardLayout>{children}</DashboardLayout>
          </AppProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
