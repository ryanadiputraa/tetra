import { fetcher } from "@/lib";
import { AddItemPayload, DataWithPagination, Item } from "@/types";

export const fetchItems = async ({
  page = 1,
  size = 50,
}): Promise<DataWithPagination<Item, "items">> => {
  const resp = await fetcher.get<DataWithPagination<Item, "items">>(
    "/api/inventory",
    { params: { page, size } },
  );
  return resp.data;
};

export const addItem = async (payload: AddItemPayload): Promise<Item> => {
  const resp = await fetcher.post<Item>("/api/inventory", payload);
  return resp.data;
};
