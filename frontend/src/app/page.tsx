"use client";

import { Button } from "antd";

import { useUserData } from "@/queries";

export default function Home() {
  const { data, isLoading, error, refetch } = useUserData();

  if (error) {
    return (
      <div className="min-h-screen bg-slate-200 grid place-items-center">
        <div className="text-center">
          <p>Terjadi kesalahan, mohon coba beberapa saat lagi.</p>
          <Button size="large" type="default" onClick={() => refetch()}>
            Muat Kembali
          </Button>
        </div>
      </div>
    );
  }

  if (isLoading) {
    return <p>Loading...</p>;
  }

  return (
    <div>
      <h1>Inventra</h1>
    </div>
  );
}
