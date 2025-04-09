"use client";

import { MinusCircleOutlined, PlusOutlined } from "@ant-design/icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  Button,
  Form,
  Input,
  InputNumber,
  Modal,
  notification,
  Select,
} from "antd";

import { addItem } from "@/api";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { formatCurrency, getAssetType, parseCurrency } from "@/lib";
import { QUERY_KEYS } from "@/queries";
import { AddItemPayload, APIError, Item, ModalProps } from "@/types";

export const AddItemModal = ({ open, onCloseAction }: ModalProps) => {
  const [form] = Form.useForm<AddItemPayload>();
  const [toast, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const { mutate, isPending } = useMutation<Item, APIError, AddItemPayload>({
    mutationKey: ["addInventoryItem"],
    mutationFn: addItem,
    onSuccess: async () => {
      await queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.inventoryItems,
      });
      form.resetFields();
      onCloseAction();
      toast.success({
        message: "Item ditambahkan",
        placement: "bottomRight",
      });
    },
    onError: (error) => {
      if (!error.errors) {
        toast.error({
          message: "Gagal menambahkan item",
          description: API_MSG[error.message] || SERVER_ERR_MSG,
          placement: "bottomRight",
        });
      } else {
        const fields = [];
        for (const e in error.errors) {
          fields.push({
            name: e as keyof AddItemPayload,
            errors: [API_MSG[error.errors[e]]],
          });
        }
        form.setFields(fields);
      }
    },
  });

  return (
    <Modal
      open={open}
      title="Tambah Item"
      onCancel={onCloseAction}
      confirmLoading={isPending}
      onOk={() => form.submit()}
      okText="Simpan"
    >
      {contextHolder}
      <Form
        form={form}
        onFinish={mutate}
        initialValues={{ prices: [{ price: 0, quantity: 1 }] }}
        className="my-4 flex flex-col"
      >
        <label className="mb-1">Nama Item</label>
        <Form.Item
          name="item_name"
          rules={[{ required: true, message: "Masukan name item" }]}
        >
          <Input size="large" placeholder="Nama Item" />
        </Form.Item>
        <label className="mb-1">Tipe Item</label>
        <Form.Item
          name="type"
          rules={[{ required: true, message: "Pilih tipe item" }]}
        >
          <Select
            size="large"
            placeholder="Tipe Item"
            options={[
              { value: "consumable", label: getAssetType("consumable") },
              { value: "fixed_asset", label: getAssetType("fixed_asset") },
            ]}
          />
        </Form.Item>
        <Form.List name="prices">
          {(fields, { add, remove }) => (
            <>
              {fields.map(({ key, name }) => (
                <div key={key} className="flex items-center gap-4">
                  <div className="flex-[2]">
                    <label className="mb-1">Harga</label>
                    <Form.Item
                      name={[name, "price"]}
                      rules={[{ required: true, message: "Masukan harga" }]}
                    >
                      <InputNumber
                        size="large"
                        placeholder="Rp 0"
                        formatter={formatCurrency}
                        parser={parseCurrency}
                        className="w-full"
                      />
                    </Form.Item>
                  </div>
                  <div className="flex-1">
                    <label className="mb-1">Kuantitas</label>
                    <Form.Item
                      name={[name, "quantity"]}
                      rules={[
                        {
                          validator: (_, value) =>
                            value >= 1
                              ? Promise.resolve()
                              : Promise.reject("Minimum kuantitas tidak valid"),
                        },
                      ]}
                    >
                      <InputNumber
                        type="number"
                        min={1}
                        size="large"
                        placeholder="1"
                        className="w-full"
                      />
                    </Form.Item>
                  </div>
                  {fields.length > 1 && (
                    <Button
                      danger
                      icon={<MinusCircleOutlined />}
                      onClick={() => remove(name)}
                    />
                  )}
                </div>
              ))}
              <Button onClick={() => add()} icon={<PlusOutlined />}>
                Tambah Harga
              </Button>
            </>
          )}
        </Form.List>
      </Form>
    </Modal>
  );
};
