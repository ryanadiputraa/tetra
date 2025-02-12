export type Organization = {
  id: number;
  owner_id: number;
  name: string;
  created_at: string;
  subscription_end_date: string;
};

export type OrganizationPayload = {
  name: string;
};

export type Member = {
  id: number;
  user_id: number;
  fullname: string;
  email: string;
  role: string;
};
