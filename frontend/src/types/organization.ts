import { Role } from "./user";

export type Organization = {
  id: number;
  owner: {
    id: number;
    fullname: string;
    email: string;
    created_at: string;
  };
  name: string;
  created_at: string;
  subscription_end_at: string;
  features: Features;
};

export type Features = {
  dashboard: boolean;
};

export type OrganizationPayload = {
  name: string;
};

export type Member = {
  id: number;
  user_id: number;
  fullname: string;
  email: string;
  role: Role;
};

export interface InviteMemberPayload {
  email: string;
}

export interface AcceptInvitationPayload {
  code: string;
}

export interface ChangeRolePayload {
  memberId: number;
  role: Role;
}

export interface UpdateDashboardSettingsPayload {
  odoo_url: string;
  odoo_db: string;
  odoo_username: string;
  odoo_password: string;
  intellitrack_username: string;
  intellitrack_password: string;
}
