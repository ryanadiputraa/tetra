export * from "./auth";
export * from "./user";
export * from "./organization";

export type APIError = {
  message: string;
  errors?: Record<string, string>;
};

export interface ModalProps {
  open: boolean;
  onCloseAction: () => void;
}
