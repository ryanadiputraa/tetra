"use client";

import { API_MSG } from "@/constant";
import { Button } from "antd";

interface Props {
  loading?: boolean;
  onRetryAction: () => void;
  msg?: string;
}
export const ErrorPage = ({ loading, onRetryAction, msg = "" }: Props) => {
  const isSubscriptionEnd = msg === "subscription_end";

  const handleError = () => {
    if (!isSubscriptionEnd && onRetryAction) onRetryAction();
  };

  return (
    <div className="h-full w-full grid place-items-center">
      <div className="text-center">
        <p className="mb-2">
          {API_MSG[msg] ?? "Terjadi kesalahan, mohon coba beberapa saat lagi."}
        </p>
        <Button
          loading={loading}
          size="large"
          type="default"
          onClick={handleError}
        >
          {isSubscriptionEnd ? "Pembayaran" : "Muat Kembali"}
        </Button>
      </div>
    </div>
  );
};
