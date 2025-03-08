import { useQuery, UseQueryOptions } from "@tanstack/react-query";

import { fetchUserData } from "@/api";
import { User } from "@/types";
import { QUERY_KEYS } from ".";

export const useUserData = (options?: Partial<UseQueryOptions<User>>) => {
  return useQuery({
    queryKey: QUERY_KEYS.userData,
    queryFn: fetchUserData,
    staleTime: Infinity,
    ...options,
  });
};
