export interface JWT {
  access_token: string;
  expires_at: string;
}

export interface LoginPayload {
  email: string;
  password: string;
}

export interface RegisterPayload extends LoginPayload {
  fullname: string;
}

export interface RegisterPayloadForm extends RegisterPayload {
  confirm_password: string;
}
