"use client";

import { Button } from "antd";
import { MouseEventHandler } from "react";

interface Props {
  onRetry: MouseEventHandler<HTMLElement> | undefined;
}
export const ErrorPage = ({ onRetry }: Props) => {
  return (
    <div className="min-h-screen bg-slate-200 grid place-items-center">
      <div className="text-center">
        <p className="mb-2">
          Terjadi kesalahan, mohon coba beberapa saat lagi.
        </p>
        <Button size="large" type="default" onClick={onRetry}>
          Muat Kembali
        </Button>
      </div>
    </div>
  );
};
