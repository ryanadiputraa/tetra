import { TableColumnsType } from "antd";
import dayjs from "dayjs";

import { Item, ItemPrice, ItemType } from "@/types";
import { getAssetType } from "@/lib";

export const tableColumn = (): TableColumnsType<Item> => {
  return [
    {
      title: "ID",
      dataIndex: "id",
      sorter: (a, b) => a.id - b.id,
    },
    {
      title: "Nama Item",
      dataIndex: "item_name",
      sorter: (a, b) => a.item_name.localeCompare(b.item_name),
    },
    {
      title: "Tipe Item",
      dataIndex: "item_type",
      render: (type: ItemType) => getAssetType(type),
      sorter: (a, b) => a.item_type.localeCompare(b.item_type),
    },
    {
      title: "Stok",
      dataIndex: "stock",
      render: (stock: ItemPrice[]) => (
        <div>
          <span>{stock.reduce((acc, price) => acc + price.quantity, 0)}</span>
        </div>
        // TODO: popover prices detail
      ),
      sorter: (a, b) => {
        const aStock = a.stock.reduce((acc, price) => acc + price.quantity, 0);
        const bStock = b.stock.reduce((acc, price) => acc + price.quantity, 0);
        return aStock - bStock;
      },
    },
    {
      title: "Tanggal Dibuat",
      dataIndex: "created_at",
      render: (createdAt: string) => dayjs(createdAt).format("DD/MM/YYYY"),
      sorter: (a, b) => dayjs(a.created_at).unix() - dayjs(b.created_at).unix(),
    },
  ];
};
