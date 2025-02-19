"use client";

import { Button } from "antd";
import { MouseEventHandler } from "react";

interface Props {
  loading?: boolean;
  onRetry: MouseEventHandler<HTMLElement> | undefined;
}
export const ErrorPage = ({ loading, onRetry }: Props) => {
  return (
    <div className="h-full w-full grid place-items-center">
      <div className="text-center">
        <p className="mb-2">
          Terjadi kesalahan, mohon coba beberapa saat lagi.
        </p>
        <Button loading={loading} size="large" type="default" onClick={onRetry}>
          Muat Kembali
        </Button>
      </div>
    </div>
  );
};
