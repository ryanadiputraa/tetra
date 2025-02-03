"use client";

import { fetcher, setCookie } from "@/lib";
import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";
import Link from "next/link";

import { COOKIE_AUTH_KEY } from "@/constant";
import { Button } from "antd";

export default function Auth() {
  const [isError, setIsError] = useState(false);
  const router = useRouter();
  const search = useSearchParams();

  useEffect(() => {
    const accessToken = search.get("access_token");
    const expiresAt = search.get("expires_at");
    if (accessToken && expiresAt) {
      setCookie(COOKIE_AUTH_KEY, accessToken, new Date(expiresAt));
      fetcher.defaults.headers.common["Authorization"] =
        `Bearer ${accessToken}`;
      router.push("/");
    } else {
      setIsError(true);
    }
  }, [router, search]);

  return (
    <div className="min-h-screen bg-slate-200 grid place-items-center">
      {isError ? (
        <div className="text-center">
          <p>Terjadi kesalahan, mohon coba beberapa saat lagi.</p>
          <Link href="/login">
            <Button size="large" type="link">
              Back to Login
            </Button>
          </Link>
        </div>
      ) : (
        <p>Authenticating...</p>
      )}
    </div>
  );
}
