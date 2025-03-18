"use client";

import {
  MinusCircleOutlined,
  PlusOutlined,
  UserOutlined,
} from "@ant-design/icons";
import {
  Button,
  Form,
  Input,
  InputNumber,
  Modal,
  notification,
  Select,
} from "antd";

import { formatCurrency, getAssetType, parseCurrency } from "@/lib";
import { AddItemPayload, ModalProps } from "@/types";

export const AddItemModal = ({ open, onCloseAction }: ModalProps) => {
  const [form] = Form.useForm<AddItemPayload>();
  const [toast, contextHolder] = notification.useNotification();

  return (
    <Modal
      open={open}
      title="Tambah Item"
      onCancel={onCloseAction}
      // confirmLoading={isPending}
      onOk={() => form.submit()}
      okText="Simpan"
    >
      {contextHolder}
      <Form
        form={form}
        initialValues={{ prices: [{ price: 0, quantity: 1 }] }}
        className="my-4 flex flex-col"
      >
        <label className="mb-1">Nama Item</label>
        <Form.Item
          name="item_name"
          rules={[{ required: true, message: "Masukan name item" }]}
        >
          <Input
            size="large"
            placeholder="Nama Item"
            suffix={<UserOutlined />}
          />
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
                  <Button
                    danger
                    icon={<MinusCircleOutlined />}
                    onClick={() => remove(name)}
                  />
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
