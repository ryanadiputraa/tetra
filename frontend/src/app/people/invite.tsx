"use client";

import { UserOutlined } from "@ant-design/icons";
import { useMutation } from "@tanstack/react-query";
import { Form, Input, Modal, notification } from "antd";

import { inviteMember } from "@/api/organization";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { APIError, InviteMemberPayload, ModalProps } from "@/types";

export const InviteModal = ({ open, onCloseAction }: ModalProps) => {
  const [form] = Form.useForm<InviteMemberPayload>();
  const [toast, contextHolder] = notification.useNotification();

  const { mutate, isPending } = useMutation<
    void,
    APIError,
    InviteMemberPayload
  >({
    mutationKey: ["inviteMember"],
    mutationFn: inviteMember,
    onSuccess: () => {
      form.resetFields();
      onCloseAction();
      toast.success({
        message: "Undangan bergabung dikirim",
        placement: "bottomRight",
      });
    },
    onError: (error) => {
      if (!error.errors) {
        toast.error({
          message: "Gagal mengirimkan undangan",
          description: API_MSG[error.message] || SERVER_ERR_MSG,
          placement: "bottomRight",
        });
      } else {
        const fields = [];
        for (const e in error.errors) {
          fields.push({
            name: e as keyof InviteMemberPayload,
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
      title="Undang"
      onCancel={onCloseAction}
      confirmLoading={isPending}
      onOk={() => form.submit()}
      okText="Kirim Undangan"
    >
      {contextHolder}
      <Form form={form} onFinish={(payload) => mutate(payload)}>
        <Form.Item
          name="email"
          rules={[
            { required: true, message: "Masukan Email yang akan diundang" },
          ]}
        >
          <Input
            size="large"
            placeholder="Email"
            type="email"
            suffix={<UserOutlined />}
          />
        </Form.Item>
      </Form>
    </Modal>
  );
};
