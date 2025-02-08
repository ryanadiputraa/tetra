"use client";

import { ErrorPage, Loader } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { AiOutlineTeam } from "react-icons/ai";

import { createOrganization } from "@/api/organization";
import { API_MSG } from "@/constant";
import { useUserData } from "@/queries";
import { APIError, Organization, OrganizationPayload } from "@/types";

export default function Join() {
  const [form] = Form.useForm<OrganizationPayload>();
  const [toast, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();
  const router = useRouter();

  const { data, isLoading, error, refetch } = useUserData(true);
  useEffect(() => {
    if (data && data.organization_id) router.push("/");
  }, [data, router]);

  const { mutate, isPending } = useMutation<
    Organization,
    APIError,
    OrganizationPayload
  >({
    mutationKey: ["createOrganization"],
    mutationFn: createOrganization,
    onSuccess: async () => {
      await queryClient.invalidateQueries({ queryKey: ["userData"] });
      router.push("/");
    },
    onError: (err) => {
      toast.error({
        message: "Gagal Membuat Organisasi",
        description: API_MSG[err.message],
        placement: "bottomRight",
      });
    },
  });

  const onCreate = (payload: OrganizationPayload) => {
    mutate(payload);
  };

  if (error) {
    return <ErrorPage onRetry={() => refetch()} />;
  }
  if (isLoading) {
    return <Loader />;
  }

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
}
