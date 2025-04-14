import { ItemType } from "@/types";
import { valueType } from "antd/lib/statistic/utils";

const DEFAULT_DATE = "2025-01-01";

export type DateFormat = "short" | "full" | "shortTime" | "fullTime";

export const formatDate = (
  dateStr = DEFAULT_DATE,
  format: DateFormat,
): string => {
  const isShortMonth = format === "short" || format === "shortTime";
  const separator = isShortMonth ? "/" : " ";
  const d = new Date(dateStr);

  const date = d.getDate();
  const month = isShortMonth
    ? d.getMonth() + 1
    : d.toLocaleString("id", { month: "long" });
  const year = d.getFullYear();

  return date + separator + month + separator + year;
};

export const isOnFreeTrial = (createdAt: Date): boolean => {
  const oneMonthAfterCreated = new Date(createdAt.getTime());
  oneMonthAfterCreated.setUTCMonth(oneMonthAfterCreated.getUTCMonth() + 3);
  const now = new Date();
  return now >= createdAt && now < oneMonthAfterCreated;
};

export const getAssetType = (itemType: ItemType) => {
  if (itemType === "fixed_asset") return "Aset Tetap";
  return "Barang Habis Pakai";
};

export const formatCurrency = (value: valueType | undefined) => {
  if (value) return `Rp ${value}`.replace(/\B(?=(\d{3})+(?!\d))/g, ".");
  return "Rp 0";
};

export const parseCurrency = (value: string | undefined) => {
  if (!value) return 0; // Ensure a number is always returned
  const parsed = value.replace(/[^\d]/g, ""); // Remove all non-numeric characters
  return parsed ? Number(parsed) : 0;
};
