"use client";

import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";
import { ErrorPage } from "./error";
import { Loader } from "./loader";

import { useUserData } from "@/queries";

interface Props {
  children: React.ReactNode;
}

export const DashboardLayout = ({ children }: Props) => {
  const router = useRouter();
  const pathname = usePathname();
  const excludedRoutes = ["/auth", "/login", "/register", "/join"];

  const { data, isLoading, error, refetch } = useUserData(
    !excludedRoutes.includes(pathname),
  );
  useEffect(() => {
    if (data && !data?.organization_id) router.push("/join");
  }, [data, router]);

  // Pages that not using dashboard layout component
  if (excludedRoutes.includes(pathname)) {
    return children;
  }
  if (error) {
    return <ErrorPage onRetry={() => refetch()} />;
  }
  if (isLoading) {
    return <Loader />;
  }

  // Push to join page when user hasn't join an organization
  if (!data?.organization_id) {
    return <></>;
  }

  return (
    <div>
      <p>Layout</p>
      {children}
    </div>
  );
};
