export * from "./auth";

export type APIError = {
  message: string;
  errors?: Record<string, string>;
};
