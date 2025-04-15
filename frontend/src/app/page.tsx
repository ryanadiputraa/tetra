"use client";

import { Button, DatePicker, Skeleton } from "antd";

import { PlusOutlined } from "@ant-design/icons";
import dayjs from "dayjs";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

import { ErrorPage } from "@/components";
import { useOrganization } from "@/queries";
import { InputModal } from "./input";

export default function Home() {
  const { data, isLoading, error, refetch } = useOrganization();
  const [date, setDate] = useState(dayjs());
  const [isInputOpen, setIsInputOpen] = useState(false);

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
          Dashboard belum aktif.{" "}
          <Link href="/settings#dashboard" className="text-blue-600">
            Setup Dashboard{" "}
          </Link>
          kamu untuk menggunakan fitur ini.
        </p>
      </div>
    );
  }

  return (
    <>
      <div className="flex justify-between items-center">
        <DatePicker size="large" value={date} onChange={setDate} />
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => setIsInputOpen(true)}
        >
          Input Data
        </Button>
      </div>
      <InputModal
        open={isInputOpen}
        onCloseAction={() => setIsInputOpen(false)}
      />
    </>
  );
}
