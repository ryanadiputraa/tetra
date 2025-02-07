import { fetchUserData } from "@/api";
import { useQuery } from "@tanstack/react-query";

export const useUserData = () => {
  return useQuery({
    queryKey: ["userData"],
    queryFn: fetchUserData,
    staleTime: Infinity,
  });
};
