import { TableColumnsType } from "antd";
import dayjs from "dayjs";

import { Item, ItemPrice } from "@/types";

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
      render: (type: string) => (
        <span className="capitalize">{type.split("_").join(" ")}</span>
      ),
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
      render: (createdAt: string) => (
        <span>{dayjs(createdAt).format("DD/MM/YYYY")}</span>
      ),
      sorter: (a, b) => dayjs(a.created_at).unix() - dayjs(b.created_at).unix(),
    },
  ];
};
