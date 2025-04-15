"use client";

import { Skeleton } from "antd";

import { ErrorPage } from "@/components";
import { useOrganization } from "@/queries";
import Image from "next/image";
import Link from "next/link";

export default function Home() {
  const { data, isLoading, error, refetch } = useOrganization();

  if (isLoading) {
    return (
      <div className="flex flex-col gap-8">
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
      </div>
    );
  }
  if (error) {
    return (
      <ErrorPage
        loading={isLoading}
        onRetryAction={() => refetch()}
        msg={error.message}
      />
    );
  }

  if (!data?.features.dashboard) {
    return (
      <div className="flex flex-col items-center justify-center text-center gap-6 h-full">
        <Image src="/dashboard.svg" width={360} height={360} alt="Dashboard" />
        <p className="max-w-screen-sm">
          Monitoring Dashboard belum aktif.{" "}
          <Link href="/settings#dashboard" className="text-blue-600">
            Setup Monitoring Dashboard{" "}
          </Link>
          kamu untuk menggunakan fitur Dashboard.
        </p>
      </div>
    );
  }

  return <>Dashboard</>;
}
