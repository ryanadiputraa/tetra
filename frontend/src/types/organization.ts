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
