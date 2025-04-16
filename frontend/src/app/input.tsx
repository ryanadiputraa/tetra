"use client";

import { InboxOutlined } from "@ant-design/icons";
import { useMutation } from "@tanstack/react-query";
import { Modal, notification, Upload, UploadFile } from "antd";
import dayjs from "dayjs";
import Link from "next/link";
import { useState } from "react";

import { importUtilization } from "@/api";
import { APIError, ModalProps } from "@/types";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";

const { Dragger } = Upload;

export function InputModal({ open, onCloseAction }: ModalProps) {
  const [toast, toastContextHolder] = notification.useNotification();
  const [fileList, setFileList] = useState<UploadFile[]>([]);
  const [file, setFile] = useState<File | null>(null);

  const { mutate, isPending } = useMutation<void, APIError, File>({
    mutationKey: ["importUtilization"],
    mutationFn: importUtilization,
    onSuccess: () => {
      setFileList([]);
      setFile(null);
      toast.success({
        message: "File telah diunggah",
        placement: "bottomRight",
      });
      onCloseAction();
    },
    onError: (error) => {
      toast.error({
        message: "Gagal",
        description: API_MSG[error.message] || SERVER_ERR_MSG,
        placement: "bottomRight",
      });
    },
  });

  const onUpload = () => {
    if (!file) return;
    mutate(file);
  };

  return (
    <Modal
      open={open}
      onCancel={onCloseAction}
      title="Input Data Utilisasi"
      okText="Unggah"
      onOk={() => onUpload()}
      confirmLoading={isPending}
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
        fileList={fileList}
        multiple={false}
        maxCount={1}
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
            setFileList([file]);
            setFile(file.originFileObj || null);
          } else {
            setFileList([]);
            setFile(null);
          }
        }}
        onRemove={() => {
          setFileList([]);
          setFile(null);
        }}
        className="inline-block h-64 w-full"
      >
        <InboxOutlined className="text-7xl" />
        <p>Klik atau tarik satu file CSV ke area ini untuk mengunggah data.</p>
        <p>Hanya mendukung unggahan satu file CSV.</p>
      </Dragger>
    </Modal>
  );
}
