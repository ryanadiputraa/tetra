import { AppProvider, AuthProvider } from "@/provider";
import type { Metadata } from "next";
import { Space_Grotesk } from "next/font/google";
import { cookies } from "next/headers";
import "./globals.css";

import { Theme } from "@/types";

const font = Space_Grotesk({
  subsets: ["latin"],
  weight: ["300", "400", "500", "600", "700"],
});

export const metadata: Metadata = {
  title: "Tetra",
  description: "Sistem management asset perusahaan",
};

export default async function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  const cookieStore = await cookies();
  const initialTheme = (cookieStore.get("theme")?.value as Theme) ?? "light";

  return (
    <html lang="en" className={`h-full ${initialTheme}`}>
      <body
        className={`${font.className} antialiased bg-gray-100 dark:bg-neutral-800 text-black dark:text-white h-full`}
      >
        <AuthProvider>
          <AppProvider initialTheme={initialTheme}>{children}</AppProvider>
        </AuthProvider>
      </body>
    </html>
  );
}
