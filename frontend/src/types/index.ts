export * from "./auth";
export * from "./inventory";
export * from "./organization";
export * from "./user";
export * from "./utilization";

export type Theme = "light" | "dark";

export type APIError = {
  message: string;
  errors?: Record<string, string>;
};

export interface ModalProps {
  open: boolean;
  onCloseAction: () => void;
}

export type DataWithPagination<T, K extends string> = {
  [P in K]: T[]; // Dynamic key name
} & {
  meta: Pagination;
};

export type Pagination = {
  current_page: number;
  total_pages: number;
  size: number;
  total_data: number;
};
