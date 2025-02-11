import { fetcher } from "@/lib";
import { User } from "@/types";

export const fetchUserData = async (): Promise<User> => {
  const resp = await fetcher.get<User>("/api/users/profile");
  return resp.data;
};

export const changePassword = async (password: string): Promise<void> => {
  const resp = await fetcher.post<void>("/api/users/password", { password });
  return resp.data;
};
