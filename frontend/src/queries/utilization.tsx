import { fetchUtilizationDashboard } from "@/api";
import { fetchDashboardParams, Utilizations } from "@/types";
import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { QUERY_KEYS } from ".";

export const useUtilizationDashboard = (
  params: fetchDashboardParams,
  options?: Partial<UseQueryOptions<Utilizations>>,
) => {
  return useQuery({
    queryKey: QUERY_KEYS.utilizationDashboard,
    queryFn: () => fetchUtilizationDashboard(params),
    staleTime: Infinity,
    ...options,
  });
};
