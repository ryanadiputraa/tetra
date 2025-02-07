"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";

import { COOKIE_AUTH_KEY } from "@/constant";
import { fetcher, getCookie } from "@/lib";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const pathname = usePathname();
  const [isLoggedIn, setIsLoggedIn] = useState(false);

  useEffect(() => {
    const token = getCookie(COOKIE_AUTH_KEY);
    if (
      !token &&
      pathname !== "/login" &&
      pathname !== "/register" &&
      pathname !== "/auth"
    ) {
      router.push("/login");
    } else {
      fetcher.defaults.headers.common["Authorization"] = `Bearer ${token}`;
      setIsLoggedIn(true);
    }
  }, [router, pathname]);

  if (!isLoggedIn) return <></>;
  return children;
}
