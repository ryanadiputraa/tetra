import { fetcher } from "@/lib";
import { JWT, LoginPayload, RegisterPayload } from "@/types";

export const login = async (payload: LoginPayload): Promise<JWT> => {
  const resp = await fetcher.post<JWT>("/auth/login", payload);
  return resp.data;
};

export const register = async (payload: RegisterPayload): Promise<JWT> => {
  const resp = await fetcher.post<JWT>("/auth/register", payload);
  return resp.data;
};
