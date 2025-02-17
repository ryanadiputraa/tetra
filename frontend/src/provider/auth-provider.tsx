"use client";

import { useParams, usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

import { COOKIE_AUTH_KEY, LS_INVITATION_CODE_KEY } from "@/constant";
import { fetcher, getCookie } from "@/lib";

export const AuthProvider = ({ children }: { children: React.ReactNode }) => {
  const router = useRouter();
  const pathname = usePathname();
  const { code } = useParams();
  const [isCheck, setIsCheck] = useState(false);

  useEffect(() => {
    const token = getCookie(COOKIE_AUTH_KEY);
    const excludedRoutes = ["/auth", "/login", "/register"];
    if (!token && !excludedRoutes.includes(pathname)) {
      if (code) {
        window.localStorage.setItem(LS_INVITATION_CODE_KEY, String(code));
      }
      router.push("/login");
    } else {
      fetcher.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      setIsCheck(true);
    }
  }, [router, pathname, code]);

  if (!isCheck) return <></>;
  return children;
};
