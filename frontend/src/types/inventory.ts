export type Item = {
  id: number;
  item_name: string;
  item_type: ItemType;
  stock: ItemPrice[];
  created_at: string;
};

export type ItemType = "consumable" | "fixed_asset";

export type ItemPrice = {
  price: number;
  quantity: number;
  created_at: string;
};

export type AddItemPayload = {
  item_name: string;
  type: ItemType;
  prices: ItemPricePayload[];
};

export type ItemPricePayload = {
  price: number;
  quantity: number;
};
