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
