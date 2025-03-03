import { fetcher } from "@/lib";
import {
  AcceptInvitationPayload,
  ChangeRolePayload,
  InviteMemberPayload,
  Member,
  Organization,
  OrganizationPayload,
} from "@/types";

export const fetchData = async (): Promise<Organization> => {
  const resp = await fetcher.get<Organization>("/api/organizations");
  return resp.data;
};

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

export const inviteMember = async (
  payload: InviteMemberPayload,
): Promise<void> => {
  await fetcher.post("/api/organizations/invite", payload);
};

export const acceptInvitation = async (
  payload: AcceptInvitationPayload,
): Promise<Member> => {
  const resp = await fetcher.post<Member>("/api/organizations/join", payload);
  return resp.data;
};

export const removeMember = async (id: number): Promise<void> => {
  await fetcher.delete(`/api/organizations/members/${id}`);
};

export const changeRole = async ({
  memberId,
  role,
}: ChangeRolePayload): Promise<void> => {
  await fetcher.put(`/api/organizations/members/${memberId}`, { role });
};

export const leaveOrganization = async (): Promise<void> => {
  await fetcher.delete("/api/organizations/leave");
};
