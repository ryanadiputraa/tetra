import { fetcher } from "@/lib";
import { JWT, LoginPayload } from "@/types";

export const login = async ({
  email,
  password,
}: LoginPayload): Promise<JWT> => {
  const resp = await fetcher.post<JWT>("/auth/login", { email, password });
  return resp.data;
};
