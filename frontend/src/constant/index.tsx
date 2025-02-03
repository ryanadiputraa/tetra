export const COOKIE_AUTH_KEY = "inventra-auth";
export const API_URL = process.env.NEXT_PUBLIC_API_URL;

export const SERVER_ERR_MSG =
  "Sistem dalam perbaikan, mohon coba beberapa saat lagi.";

export const SERVER_ERR = "internal_server_error";

export const API_MSG: Record<string, string> = {
  bad_request: "Mohon periksa kembali input yang kamu masukan.",
  unauthorized: "Kamu tidak memiliki akses.",
  forbidden: "Kamu tidak memiliki akses.",
  not_found: "Data tidak ditemukan.",
  internal_server_error:
    "Sistem dalam perbaikan, mohon coba beberapa saat lagi.",
  email_taken: "Email telah terdaftar.",
};
