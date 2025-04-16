import { fetcher } from "@/lib";

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
