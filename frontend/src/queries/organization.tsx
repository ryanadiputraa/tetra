import { useQuery, UseQueryOptions } from "@tanstack/react-query";

import { acceptInvitation, fetchData, fetchOrganizationMembers } from "@/api";
import { Member, Organization } from "@/types";
import { QUERY_KEYS } from ".";

export const useOrganization = (
  options?: Partial<UseQueryOptions<Organization>>,
) => {
  return useQuery({
    queryKey: QUERY_KEYS.organizationData,
    queryFn: fetchData,
    staleTime: 5 * 60 * 1000,
    ...options,
  });
};

export const useOrganizationMembers = (
  options?: Partial<UseQueryOptions<Member[]>>,
) => {
  return useQuery({
    queryKey: QUERY_KEYS.organizationMembers,
    queryFn: fetchOrganizationMembers,
    staleTime: 5 * 60 * 1000,
    ...options,
  });
};

export const useAcceptInvitation = (
  code: string,
  options?: Partial<UseQueryOptions<Member>>,
) => {
  return useQuery({
    queryKey: [...QUERY_KEYS.acceptInvitation, code],
    queryFn: () => acceptInvitation({ code }),
    retry: false,
    ...options,
  });
};
