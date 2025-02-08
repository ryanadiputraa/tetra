export * from "./auth";
export * from "./user";
export * from "./organization";

export type APIError = {
  message: string;
  errors?: Record<string, string>;
};
