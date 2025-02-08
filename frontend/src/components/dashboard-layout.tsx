"use client";

import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";
import { ErrorPage } from "./error";
import { Loader } from "./loader";

import { mainMenu } from "@/constant";
import { useUserData } from "@/queries";

import logo from "@/assets/svg/inventra.svg";

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
    <div className="min-h-screen flex">
      <nav className="bg-base text-white w-56">
        <div className="flex items-center gap-2 p-6">
          <Image src={logo} alt="inventra" width={32} />
          <h1 className="text-xl">Inventra</h1>
        </div>
        <div className="text-slate-300 mt-8">
          <span className="text-sm pl-6">MENU</span>
          <ul className="mt-2 flex flex-col gap-2">
            {mainMenu.map((menu) => (
              <li key={menu.link}>
                <Link
                  href={menu.link}
                  className={`border-l-4 p-2 pl-6 flex items-center gap-2 ${pathname === menu.link ? "border-primary text-white font-medium" : "border-transparent hover:text-primary"}`}
                >
                  {pathname === menu.link ? (
                    <menu.IcoActive className="text-2xl text-primary" />
                  ) : (
                    <menu.Ico className="text-2xl" />
                  )}
                  {menu.label}
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </nav>
      <main className="p-6">{children}</main>
    </div>
  );
};
