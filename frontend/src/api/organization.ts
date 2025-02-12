import { fetcher } from "@/lib";
import { Member, Organization, OrganizationPayload } from "@/types";

export const createOrganization = async (
  payload: OrganizationPayload,
): Promise<Organization> => {
  const resp = await fetcher.post<Organization>("/api/organizations", payload);
  return resp.data;
};

export const fetchOrganizationMembers = async (): Promise<Member[]> => {
  const resp = await fetcher.get<Record<string, Member[]>>(
    "/api/organizations/members",
  );
  return resp.data.members ?? [];
};
