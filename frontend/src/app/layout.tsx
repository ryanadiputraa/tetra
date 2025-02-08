import { DashboardLayout } from "@/components";
import { AppProvider, AuthProvider } from "@/provider";
import "@ant-design/v5-patch-for-react-19";
import type { Metadata } from "next";
import { Poppins } from "next/font/google";
import "./globals.css";

const font = Poppins({
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
    <html lang="en">
      <body className={`${font} antialiased bg-slate-200 text-slate-950`}>
        <AuthProvider>
          <AppProvider>
            <DashboardLayout>{children}</DashboardLayout>
          </AppProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
