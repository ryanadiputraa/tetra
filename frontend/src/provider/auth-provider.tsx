"use client";

import { useRouter } from "next/navigation";
import { useEffect } from "react";

import { COOKIE_AUTH_KEY } from "@/constant";
import { getCookie } from "@/lib/storage";

export default function AuthProvider({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();

  useEffect(() => {
    console.log("hiot");
    const token = getCookie(COOKIE_AUTH_KEY);
    if (!token) {
      router.push("/login");
    }
    window.addEventListener("popstate", () => router.refresh());
    return () => window.removeEventListener("popstate", () => router.refresh());
  });

  return children;
}
