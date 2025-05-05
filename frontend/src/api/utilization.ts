import { fetcher } from "@/lib";
import { fetchDashboardParams, Utilizations } from "@/types";

export const importUtilization = async (file: File): Promise<void> => {
  await fetcher.post(
    "/api/utilizations/import",
    { file },
    {
      headers: {
        "Content-Type": "multipart/form-data",
      },
    },
  );
};

export const fetchUtilizationDashboard = async (
  params: fetchDashboardParams,
): Promise<Utilizations> => {
  const resp = await fetcher.get<Utilizations>("/api/utilizations/dashboard", {
    params,
  });
  return resp.data;
};
