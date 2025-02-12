import { useQuery, UseQueryOptions } from "@tanstack/react-query";

import { fetchOrganizationMembers } from "@/api/organization";
import { Member } from "@/types";
import { QUERY_KEYS } from ".";

export const useOrganizationMembers = (options?: UseQueryOptions<Member[]>) => {
  return useQuery({
    queryKey: QUERY_KEYS.organizationMembers,
    queryFn: fetchOrganizationMembers,
    staleTime: 5 * 60 * 1000,
    ...options,
  });
};
