"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

import { COOKIE_AUTH_KEY } from "@/constant";
import { fetcher, getCookie } from "@/lib";

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const pathname = usePathname();
  const [isCheck, setIsCheck] = useState(false);

  useEffect(() => {
    const token = getCookie(COOKIE_AUTH_KEY);
    const excludedRoutes = ["/auth", "/login", "/register"];
    if (!token && !excludedRoutes.includes(pathname)) {
      router.push("/login");
    } else {
      fetcher.defaults.headers.common["Authorization"] = `Bearer ${token}`;
    }
    setIsCheck(true);
  }, [router, pathname]);

  if (!isCheck) return <></>;
  return children;
};
