export type JWT = {
  access_token: string;
  expires_at: string;
};

export type LoginPayload = {
  email: string;
  password: string;
};
