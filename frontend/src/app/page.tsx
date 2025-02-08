"use client";

import { Join } from "./components";
import { ErrorPage } from "./components/error";
import { Loader } from "./components/loader";

import { useUserData } from "@/queries";

export default function Home() {
  const { data, isLoading, error, refetch } = useUserData();

  if (error) {
    return <ErrorPage onRetry={() => refetch()} />;
  }
  if (isLoading) {
    return <Loader />;
  }
  if (!data?.organization_id) {
    return <Join />;
  }

  return (
    <div>
      <h1>Inventra</h1>
    </div>
  );
}
