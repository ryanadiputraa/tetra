"use client";

import { DownOutlined, LogoutOutlined, UserOutlined } from "@ant-design/icons";
import { Button, Dropdown, MenuProps, Switch } from "antd";
import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";
import { ErrorPage } from "./error";
import { Loader } from "./loader";

import { COOKIE_AUTH_KEY, mainMenu, secondaryMenu } from "@/constant";
import { formatDate, isOnFreeTrial, removeCookie } from "@/lib";
import { useOrganization, useUserData } from "@/queries";
import { Theme } from "@/types";

interface Props {
  theme: Theme;
  toggleThemeAction: () => void;
  children: React.ReactNode;
}

export const DashboardLayout = ({
  theme,
  toggleThemeAction,
  children,
}: Props) => {
  const router = useRouter();
  const pathname = usePathname();
  const excludedRoutes = ["/auth", "/login", "/register"];
  const isIgnorePath = pathname.startsWith("/join");
  const headerTitle = pathname.split("/").filter(Boolean).pop() ?? "dashboard";

  const { data: organization } = useOrganization();
  const isFreeTrial =
    organization?.subscription_end_at && organization?.created_at
      ? isOnFreeTrial(
          new Date(organization.subscription_end_at),
          new Date(organization.created_at),
        )
      : false;

  const onLogout = () => {
    removeCookie(COOKIE_AUTH_KEY);
    router.push("/login");
  };

  const menuItems: MenuProps["items"] = [
    {
      key: "1",
      label: <Link href="/profile">Akun Saya</Link>,
      icon: <UserOutlined />,
    },
    {
      type: "divider",
    },
    {
      key: "2",
      label: <button onClick={onLogout}>Keluar</button>,
      icon: <LogoutOutlined />,
    },
  ];

  const { data, isLoading, error, refetch } = useUserData(
    !excludedRoutes.includes(pathname) || isIgnorePath,
  );
  useEffect(() => {
    if (data && !data?.organization_id && !isIgnorePath) {
      router.push("/join");
    }
  }, [data, router, isIgnorePath]);

  // Pages that not using dashboard layout component
  if (excludedRoutes.includes(pathname) || isIgnorePath) {
    return children;
  }
  if (isLoading) {
    return <Loader />;
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
  // Push to join page when user hasn't join an organization
  if (!data?.organization_id) {
    return <></>;
  }

  return (
    <div className="min-h-screen flex">
      <nav className="relative w-80 bg-white dark:bg-neutral-900 text-neutral-400 border-r-2 border-gray-200 dark:border-gray-500">
        <div className="flex items-center gap-2 p-6">
          <Image src="/inventra.png" alt="inventra" width={32} height={32} />
          <h1 className="text-xl text-black dark:text-white font-semibold">
            Inventra
          </h1>
        </div>
        <div className="mt-8 px-6">
          <span className="text-sm">MENU UTAMA</span>
          <ul className="mt-2 flex flex-col gap-2">
            {mainMenu.map((menu) => (
              <li key={menu.link}>
                <Link
                  href={menu.link}
                  className={`p-2 flex items-center gap-2 ${pathname === menu.link ? "bg-primary dark:bg-primary-dark rounded-lg text-white" : "hover:text-primary dark:hover:text-primary"}`}
                >
                  {pathname === menu.link ? (
                    <menu.IcoActive className="text-2xl text-white" />
                  ) : (
                    <menu.Ico className="text-2xl hover:text-primary dark:hover:text-primary" />
                  )}
                  {menu.label}
                </Link>
              </li>
            ))}
          </ul>
        </div>
        <div className="mt-8 px-6">
          <ul className="pt-8 flex flex-col gap-2 border-t-2 border-gray-200 dark:border-gray-500">
            {secondaryMenu.map((menu) => (
              <li key={menu.link}>
                <Link
                  href={menu.link}
                  className={`p-2 flex items-center gap-2 ${pathname === menu.link ? "bg-primary rounded-lg text-white" : "hover:text-primary dark:hover:text-primary"}`}
                >
                  {pathname === menu.link ? (
                    <menu.IcoActive className="text-2xl text-white" />
                  ) : (
                    <menu.Ico className="text-2xl hover:text-primary dark:hover:text-primary" />
                  )}
                  {menu.label}
                </Link>
              </li>
            ))}
            <li className="p-2 flex justify-between items-center">
              <span>Dark Mode</span>
              <Switch value={theme === "dark"} onChange={toggleThemeAction} />
            </li>
          </ul>
        </div>
        {isFreeTrial && (
          <div className="absolute bottom-0 p-6">
            <div className="bg-primary dark:bg-primary-dark rounded-lg p-3 text-white text-sm">
              <p>
                Anda sedang menggunakan mode uji coba gratis hingga{" "}
                <span className="font-bold">
                  {formatDate(organization?.subscription_end_at, "full")}.
                </span>{" "}
                Upgrade sekarang untuk terus menikmati fitur Inventra!
              </p>
              {/* TODO: handle payment */}
              <Link href="/payment">
                <Button className="mt-3 text-primary dark:text-primary-dark font-bold">
                  Pembayaran
                </Button>
              </Link>
            </div>
          </div>
        )}
      </nav>
      <div className="w-full h-screen flex flex-col">
        <header className="py-3 px-6 bg-white dark:bg-neutral-900 border-b-2 border-gray-200 dark:border-gray-500 flex justify-between items-center">
          <h3 className="capitalize text-lg font-medium">{headerTitle}</h3>
          <Dropdown menu={{ items: menuItems }}>
            <div className="flex items-center gap-2">
              <div className="grid place-items-center size-10 bg-primary dark:bg-primary-dark rounded-full">
                <span className="text-lg text-white font-bold">
                  {data?.fullname.split("")[0]}
                </span>
              </div>
              <DownOutlined className="text-sm" />
            </div>
          </Dropdown>
        </header>
        <main className="flex-1 overflow-auto p-6">{children}</main>
      </div>
    </div>
  );
};
