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
  subscription_end:
    "Masa langganan telah berakhir, silahkan lakukan pembayaran untuk kembali menggunakan aplikasi.",
  record_not_found: "Data tidak ditemukan",
  email_taken: "Email telah terdaftar.",
  missing_auth_header: "Silahkan login terlebih dahulu.",
  invalid_auth_header: "Silahkan login terlebih dahulu.",
  organization_already_exists: "Organisasi sudah dibuat.",
  user_has_joined_org: "Pengguna telah bergabung dengan organisasi.",
  required_field: "Wajib diisi",
  email_field: "Format email tidak valid",
  max_length_field: "Input melebihi batas maksimum",
  min_length_field: "Input tidak mencapai batas minimum",
  url_field: "Format URL tidak valid",
  date_field: "Format tanggal tidak valid",
};
