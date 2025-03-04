"use client";

import { Loader } from "@/components";
import { fetcher, setCookie } from "@/lib";
import Link from "next/link";
import { useRouter, useSearchParams } from "next/navigation";
import { useEffect, useState } from "react";

import { API_MSG, COOKIE_AUTH_KEY, LS_INVITATION_CODE_KEY } from "@/constant";
import { Button } from "antd";

export default function Auth() {
  const [isError, setIsError] = useState(false);
  const [isAuthError, setIsAuthError] = useState(false);
  const router = useRouter();
  const search = useSearchParams();

  useEffect(() => {
    const accessToken = search.get("access_token");
    const expiresAt = search.get("expires_at");
    const authErr = search.get("err");
    if (accessToken && expiresAt) {
      setCookie(COOKIE_AUTH_KEY, accessToken, new Date(expiresAt));
      fetcher.defaults.headers.common["Authorization"] =
        `Bearer ${accessToken}`;
      const code = window.localStorage.getItem(LS_INVITATION_CODE_KEY);
      router.push(code ? `/join/${code}` : "/");
    } else if (authErr) {
      setIsAuthError(true);
    } else {
      setIsError(true);
    }
  }, [router, search]);

  if (isError) {
    return (
      <div className="min-h-screen bg-slate-200 grid place-items-center">
        <div className="text-center">
          <p className="mb-2">
            Terjadi kesalahan, mohon coba beberapa saat lagi.
          </p>
          <Link href="/login">
            <Button size="large" type="link">
              Kembali ke Login
            </Button>
          </Link>
        </div>
      </div>
    );
  }

  if (isAuthError) {
    return (
      <div className="min-h-screen bg-slate-200 grid place-items-center">
        <p className="max-w-screen-sm text-center">
          {API_MSG[search.get("err") ?? ""]}
        </p>
      </div>
    );
  }

  return <Loader />;
}
