export type User = {
  id: number;
  email: string;
  fullname: string;
  created_at: string;
  organization_id: number;
  role: Role;
};

export type Role = "admin" | "supervisor" | "staff";

export type ChangePasswordForm = {
  password: string;
  confirm: string;
};
