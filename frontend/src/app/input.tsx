"use client";

import { ModalProps } from "@/types";
import { InboxOutlined } from "@ant-design/icons";
import { Modal, notification, Upload } from "antd";
import dayjs from "dayjs";
import Link from "next/link";
import { useState } from "react";

const { Dragger } = Upload;

export function InputModal({ open, onCloseAction }: ModalProps) {
  const [toast, toastContextHolder] = notification.useNotification();
  const [file, setFile] = useState<File | null>(null);

  return (
    <Modal
      open={open}
      onCancel={onCloseAction}
      title="Input Data Utilisasi"
      okText="Unggah"
    >
      {toastContextHolder}
      <p className="mb-3">
        Unggah file dalam format <span className="font-bold italic">CSV</span>{" "}
        sesuai format{" "}
        <Link
          download
          href="/Input-Utilisasi-Template.csv"
          className="text-blue-600 font-semibold"
        >
          Template.
        </Link>{" "}
      </p>
      <Dragger
        name={dayjs().toISOString() + "-input"}
        multiple={false}
        accept=".csv"
        beforeUpload={(file) => {
          const isCSV = file.type === "text/csv" || file.name.endsWith(".csv");
          if (!isCSV) {
            toast.error({
              message: "Gagal",
              description: "Mohon unggah file dengan format CSV",
              placement: "bottomRight",
            });
          }
          return isCSV || Upload.LIST_IGNORE;
        }}
        onChange={({ file }) => {
          if (file.status !== "removed") {
            setFile(file.originFileObj || null);
          } else {
            setFile(null);
          }
        }}
      >
        <InboxOutlined className="text-7xl" />
        <p>Klik atau tarik satu file CSV ke area ini untuk mengunggah data.</p>
        <p>Hanya mendukung unggahan satu file CSV.</p>
      </Dragger>
    </Modal>
  );
}
