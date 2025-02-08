"use client";

import { createOrganization } from "@/api/organization";
import { API_MSG } from "@/constant";
import { APIError, Organization, OrganizationPayload } from "@/types";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";
import { AiOutlineTeam } from "react-icons/ai";

export const Join = () => {
  const [form] = Form.useForm<OrganizationPayload>();
  const [toast, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const { mutate, isPending } = useMutation<
    Organization,
    APIError,
    OrganizationPayload
  >({
    mutationKey: ["createOrganization"],
    mutationFn: createOrganization,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["userData"] });
    },
    onError: (err) => {
      toast.error({
        message: "Gagal Membuat Organisasi",
        description: API_MSG[err.message],
        placement: "topRight",
      });
    },
  });

  const onCreate = (payload: OrganizationPayload) => {
    mutate(payload);
  };

  return (
    <div className="min-h-screen grid place-items-center px-8">
      {contextHolder}
      <div className="bg-white py-16 px-8 sm:px-16 rounded-xl w-full sm:max-w-xl">
        <div className="text-center">
          <h4 className="font-bold text-2xl">Inventra</h4>
          <p className="mt-2">
            Kamu belum bergabung ke dalam organisasi. Cek email kamu untuk
            menerima undangan atau buat organisasi untuk mulai.
          </p>
          <Form form={form} onFinish={onCreate} className="mt-8 flex flex-col">
            <Form.Item name="name" rules={[{ required: true, message: "" }]}>
              <Input
                size="large"
                placeholder="Nama Organisasi"
                suffix={<AiOutlineTeam />}
              />
            </Form.Item>
            <Button
              htmlType="submit"
              size="large"
              variant="solid"
              color="primary"
              loading={isPending}
              className="font-semibold self-center px-8"
            >
              Buat Organisasi
            </Button>
          </Form>
        </div>
      </div>
    </div>
  );
};
