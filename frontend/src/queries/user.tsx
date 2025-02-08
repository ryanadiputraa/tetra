import { fetchUserData } from "@/api";
import { useQuery } from "@tanstack/react-query";

export const useUserData = (enabled: boolean) => {
  return useQuery({
    queryKey: ["userData"],
    queryFn: fetchUserData,
    staleTime: Infinity,
    enabled,
  });
};
