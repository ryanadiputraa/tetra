import { fetchItems } from "@/api";
import { DataWithPagination, Item } from "@/types";
import { UseQueryOptions, useQuery } from "@tanstack/react-query";
import { QUERY_KEYS } from ".";

export const useInventoryItem = (
  params?: { page: number; size: number },
  options?: Partial<UseQueryOptions<DataWithPagination<Item, "items">>>,
) => {
  return useQuery({
    queryKey: QUERY_KEYS.inventoryItems,
    queryFn: () =>
      fetchItems({ page: params?.page ?? 1, size: params?.size ?? 50 }),
    staleTime: 5 * 60 * 1000,
    ...options,
  });
};
