export interface AvatarColorStorage {
  bgColor: string;
}

export const setCookie = (
  name: string,
  val: string,
  expireDate: Date,
): void => {
  if (typeof document === "undefined") return;
  document.cookie = `${name}=${val};path=/;expires=${expireDate.toUTCString()};SameSite=Lax;Secure;`;
};

export const getCookie = (name: string): string | null => {
  if (typeof document === "undefined") return null;
  const value = `; ${document.cookie}`;
  const parts = value.split(`; ${name}=`);
  if (parts.length !== 2) return null;
  return parts.pop()?.split(";").shift() ?? null;
};

export const removeCookie = (name: string) => {
  if (typeof document === "undefined") return;
  document.cookie = `${name}=;expires=${new Date(Date.now() - 1).toUTCString()};path=/;SameSite=Lax;`;
};

export const setSessionStorage = (key: string, val: object) => {
  if (typeof window === "undefined") return;
  const json = JSON.stringify(val);
  window.sessionStorage.setItem(key, json);
};

export const getSessionStorage = <T>(key: string): T | null => {
  if (typeof window === "undefined") return null;
  const json = window.sessionStorage.getItem(key);
  if (!json) return null;
  const val: T = JSON.parse(json);
  return val;
};
