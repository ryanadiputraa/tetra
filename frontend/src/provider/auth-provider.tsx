"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";

import { COOKIE_AUTH_KEY } from "@/constant";
import { getCookie } from "@/lib";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const pathname = usePathname();

  useEffect(() => {
    const token = getCookie(COOKIE_AUTH_KEY);
    if (!token && pathname !== "/login" && pathname !== "/register") {
      router.push("/login");
    }
  }, [router, pathname]);

  return children;
}
