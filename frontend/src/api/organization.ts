import { fetcher } from "@/lib";
import { Organization, OrganizationPayload } from "@/types";

export const createOrganization = async (
  payload: OrganizationPayload,
): Promise<Organization> => {
  const resp = await fetcher.post<Organization>("/api/organizations", payload);
  return resp.data;
};
