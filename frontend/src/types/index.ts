export * from "./auth";
export * from "./user";

export type APIError = {
  message: string;
  errors?: Record<string, string>;
};
