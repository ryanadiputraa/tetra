"use client";

import {
  CloseOutlined,
  DownOutlined,
  LogoutOutlined,
  MenuOutlined,
  UserOutlined,
} from "@ant-design/icons";
import { Button, Drawer, Dropdown, MenuProps, Switch } from "antd";
import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect, useState } from "react";
import { ErrorPage } from "./error";
import { Loader } from "./loader";

import { COOKIE_AUTH_KEY, mainMenu, secondaryMenu } from "@/constant";
import { useMediaQuery } from "@/hooks";
import { formatDate, getCookie, isOnFreeTrial, removeCookie } from "@/lib";
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
  const excludedRoutes = ["/auth", "/login", "/register", "/payment"];
  const isIgnorePath = pathname.startsWith("/join");
  const headerTitle = pathname.split("/").filter(Boolean).pop() ?? "dashboard";
  const [isDrawerOpen, setIsDrawerOpen] = useState(false);
  const isMobile = useMediaQuery("(max-width: 640px)");

  const onLogout = () => {
    removeCookie(COOKIE_AUTH_KEY);
    router.push("/login");
  };

  const menuItems: MenuProps["items"] = [
    {
      key: "1",
      label: "Akun Saya",
      icon: <UserOutlined />,
      onClick: () => router.push("/profile"),
    },
    {
      type: "divider",
    },
    {
      key: "2",
      label: "Keluar",
      icon: <LogoutOutlined />,
      onClick: () => onLogout(),
    },
  ];

  const { data, isLoading, error, refetch } = useUserData({
    enabled: !excludedRoutes.includes(pathname) || isIgnorePath,
  });
  useEffect(() => {
    if (
      getCookie(COOKIE_AUTH_KEY) &&
      data &&
      !data.organization_id &&
      !isIgnorePath
    ) {
      router.push("/join");
    }
  }, [data, router, isIgnorePath]);

  const { data: organization } = useOrganization({
    enabled: Boolean(data?.organization_id),
  });
  const isFreeTrial =
    organization?.subscription_end_at && organization?.created_at
      ? isOnFreeTrial(
          new Date(organization.subscription_end_at),
          new Date(organization.created_at),
        )
      : false;

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

  const NavMenu = () => (
    <>
      <div className="hidden md:flex items-center gap-2 p-6">
        <Image src="/inventra.png" alt="inventra" width={32} height={32} />
        <h1 className="text-xl text-black dark:text-white font-semibold">
          Inventra
        </h1>
      </div>
      <div className="md:mt-8 px-6">
        <span className="text-sm">MENU UTAMA</span>
        <ul className="mt-2 flex flex-col gap-2">
          {mainMenu.map((menu) => (
            <li key={menu.link}>
              <Link
                href={menu.link}
                className={`p-2 flex items-center gap-2 ${pathname === menu.link ? "bg-primary dark:bg-primary-dark rounded-lg text-white" : "hover:text-primary dark:hover:text-primary"}`}
                onClick={() => {
                  if (isMobile) setIsDrawerOpen(false);
                }}
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
          <li className="p-2 flex justify-between items-center">
            <span>Dark Mode</span>
            <Switch value={theme === "dark"} onChange={toggleThemeAction} />
          </li>
          {secondaryMenu.map((menu) => (
            <li key={menu.link}>
              <Link
                href={menu.link}
                className={`p-2 flex items-center gap-2 ${pathname === menu.link ? "bg-primary rounded-lg text-white" : "hover:text-primary dark:hover:text-primary"}`}
                onClick={() => {
                  if (isMobile) setIsDrawerOpen(false);
                }}
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
      {isFreeTrial && (
        <div className="absolute bottom-0 left-0 p-6">
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
    </>
  );

  return (
    <div className="min-h-screen flex">
      <nav className="hidden md:inline-block relative w-80 bg-white dark:bg-neutral-900 text-neutral-400 border-r-2 border-gray-200 dark:border-gray-500">
        <NavMenu />
      </nav>
      <Drawer
        open={isDrawerOpen}
        onClose={() => setIsDrawerOpen(false)}
        closeIcon={<CloseOutlined className="absolute right-8 text-xl" />}
        placement="left"
        title={
          <div className="flex items-center gap-2">
            <Image src="/inventra.png" alt="inventra" width={32} height={32} />
            <h1 className="text-xl text-black dark:text-white font-semibold">
              Inventra
            </h1>
          </div>
        }
        className="inline"
      >
        <NavMenu />
      </Drawer>
      <div className="w-full h-screen flex flex-col">
        <header className="py-3 px-6 bg-white dark:bg-neutral-900 border-b-2 border-gray-200 dark:border-gray-500 flex justify-between items-center">
          <div className="flex items-center gap-2">
            <Button
              className="inline-block md:hidden"
              onClick={() => setIsDrawerOpen(true)}
            >
              <MenuOutlined />
            </Button>
            <h3 className="capitalize text-lg font-medium">{headerTitle}</h3>
          </div>
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
