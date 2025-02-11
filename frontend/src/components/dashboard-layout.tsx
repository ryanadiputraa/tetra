"use client";

import { Dropdown, MenuProps } from "antd";
import Image from "next/image";
import Link from "next/link";
import { usePathname, useRouter } from "next/navigation";
import { useEffect } from "react";
import { AiOutlineDown, AiOutlineLogout, AiOutlineUser } from "react-icons/ai";
import { ErrorPage } from "./error";
import { Loader } from "./loader";

import { COOKIE_AUTH_KEY, mainMenu } from "@/constant";
import { removeCookie } from "@/lib";
import { useUserData } from "@/queries";

interface Props {
  children: React.ReactNode;
}

export const DashboardLayout = ({ children }: Props) => {
  const router = useRouter();
  const pathname = usePathname();
  const excludedRoutes = ["/auth", "/login", "/register", "/join"];
  const headerTitle = pathname.split("/").filter(Boolean).pop() ?? "dashboard";

  const onLogout = () => {
    removeCookie(COOKIE_AUTH_KEY);
    router.push("/login");
  };

  const menuItems: MenuProps["items"] = [
    {
      key: "1",
      label: <Link href="/profile">Akun Saya</Link>,
      icon: <AiOutlineUser />,
    },
    {
      type: "divider",
    },
    {
      key: "2",
      label: <button onClick={onLogout}>Keluar</button>,
      icon: <AiOutlineLogout />,
    },
  ];

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
  if (isLoading) {
    return <Loader />;
  }
  if (error) {
    return <ErrorPage onRetry={() => refetch()} />;
  }
  // Push to join page when user hasn't join an organization
  if (!data?.organization_id) {
    return <></>;
  }

  return (
    <div className="min-h-screen flex">
      <nav className="w-80 bg-white text-slate-400 border-r-2 border-gray-200">
        <div className="flex items-center gap-2 p-6">
          <Image src="/inventra.png" alt="inventra" width={32} height={32} />
          <h1 className="text-xl text-black font-semibold">Inventra</h1>
        </div>
        <div className="mt-8 px-6">
          <span className="text-sm">MENU UTAMA</span>
          <ul className="mt-2 flex flex-col gap-2">
            {mainMenu.map((menu) => (
              <li key={menu.link}>
                <Link
                  href={menu.link}
                  className={`p-2 flex items-center gap-2 ${pathname === menu.link ? "bg-primary rounded-lg text-white" : "hover:text-primary"}`}
                >
                  {pathname === menu.link ? (
                    <menu.IcoActive className="text-2xl text-white" />
                  ) : (
                    <menu.Ico className="text-2xl hover:text-primary" />
                  )}
                  {menu.label}
                </Link>
              </li>
            ))}
          </ul>
        </div>
      </nav>
      <div className="w-full">
        <header className="py-3 px-6 bg-white border-b-2 border-gray-200 flex justify-between items-center">
          <h3 className="capitalize text-lg font-medium">{headerTitle}</h3>
          <Dropdown menu={{ items: menuItems }}>
            <div className="flex items-center gap-2">
              <div className="grid place-items-center size-10 bg-primary rounded-full">
                <span className="text-lg text-white font-bold">
                  {data?.fullname.split("")[0]}
                </span>
              </div>
              <AiOutlineDown className="text-sm" />
            </div>
          </Dropdown>
        </header>
        <main className="w-full p-6">{children}</main>
      </div>
    </div>
  );
};
