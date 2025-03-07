"use client";

import { useState } from "react";

import { setCookie } from "@/lib";
import { Theme } from "@/types";

export const useTheme = (initialTheme: Theme) => {
  const [theme, setTheme] = useState(initialTheme);
  const cookieExpireDate = new Date();
  cookieExpireDate.setTime(
    cookieExpireDate.getTime() + 365 * 24 * 60 * 60 * 1000,
  ); // One year

  const toggleTheme = () => {
    const html = document.querySelector("html");
    const newVal = theme === "light" ? "dark" : "light";
    html?.classList.remove(theme);
    html?.classList.add(newVal);
    setTheme(newVal);
    setCookie("theme", newVal, cookieExpireDate);
  };

  return { theme, toggleTheme };
};
