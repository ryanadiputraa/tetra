export type Item = {
  id: number;
  item_name: string;
  item_type: "consumable" | "fixed_asset";
  stock: ItemPrice[];
  created_at: string;
};

export type ItemPrice = {
  price: number;
  quantity: number;
  created_at: string;
};
