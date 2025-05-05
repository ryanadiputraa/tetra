"use client";

import { PlusOutlined } from "@ant-design/icons";
import { Button, DatePicker, Skeleton, TimePicker } from "antd";
import dayjs, { Dayjs } from "dayjs";
import Image from "next/image";
import Link from "next/link";
import { useState } from "react";

import { ErrorPage } from "@/components";
import { useOrganization } from "@/queries";
import { InputModal } from "./input";
import { Dashboard } from "./dashboard";

const { RangePicker } = TimePicker;

export default function Home() {
  const { data, isLoading, error, refetch } = useOrganization();
  const [isInputOpen, setIsInputOpen] = useState(false);
  const [date, setDate] = useState(dayjs());
  const [startTime, setStartTIme] = useState<Dayjs>(dayjs().startOf("day"));
  const [endTime, setEndTIme] = useState<Dayjs>(dayjs().endOf("day"));

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
        <div className="flex items-center gap-3">
          <DatePicker
            size="large"
            value={date}
            onChange={setDate}
            allowClear={false}
          />
          <RangePicker
            size="large"
            value={[startTime, endTime]}
            onChange={(dates) => {
              if (dates?.length && dates.length >= 2) {
                if (dates[0]) setStartTIme(dates[0]);
                if (dates[1]) setEndTIme(dates[1]);
              }
            }}
          />
        </div>
        <Button
          type="primary"
          icon={<PlusOutlined />}
          onClick={() => setIsInputOpen(true)}
        >
          Input Data
        </Button>
      </div>
      <Dashboard
        date={date.format("YYYY-MM-DD")}
        startTime={startTime.format("HH:mm:ss")}
        endTime={endTime.format("HH:mm:ss")}
      />
      <InputModal
        open={isInputOpen}
        onCloseAction={() => setIsInputOpen(false)}
      />
    </>
  );
}
