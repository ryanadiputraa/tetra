"use client";

import { ErrorPage, Loader } from "@/components";
import { LogoutOutlined, TeamOutlined } from "@ant-design/icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";
import { useRouter } from "next/navigation";
import { useEffect } from "react";

import { createOrganization } from "@/api/organization";
import { API_MSG, COOKIE_AUTH_KEY, SERVER_ERR_MSG } from "@/constant";
import { removeCookie } from "@/lib";
import { QUERY_KEYS, useUserData } from "@/queries";
import { APIError, Organization, OrganizationPayload } from "@/types";

export default function Join() {
  const [form] = Form.useForm<OrganizationPayload>();
  const [toast, contextHolder] = notification.useNotification();
  const queryClient = useQueryClient();
  const router = useRouter();

  const { data, isLoading, error, refetch } = useUserData();
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
      await queryClient.invalidateQueries({ queryKey: QUERY_KEYS.userData });
      router.push("/");
    },
    onError: (err) => {
      toast.error({
        message: "Gagal Membuat Organisasi",
        description: API_MSG[err.message] || SERVER_ERR_MSG,
        placement: "bottomRight",
      });
    },
  });

  const onLogout = () => {
    removeCookie(COOKIE_AUTH_KEY);
    router.push("/login");
  };

  const onCreate = (payload: OrganizationPayload) => {
    mutate(payload);
  };

  if (error) {
    return (
      <ErrorPage
        loading={isLoading}
        onRetryAction={() => refetch()}
        msg={error.message}
      />
    );
  }
  if (isLoading) {
    return <Loader />;
  }

  return (
    <div className="min-h-screen grid place-items-center px-8">
      {contextHolder}
      <Button
        size="large"
        icon={<LogoutOutlined />}
        className="absolute top-8 right-8"
        onClick={onLogout}
      >
        Logout
      </Button>
      <div className="bg-white dark:bg-neutral-900 py-16 px-8 md:px-16 rounded-lg w-full md:max-w-xl">
        <div className="text-center">
          <h4 className="font-bold text-2xl">Inventra</h4>
          <p className="mt-2">
            Kamu belum bergabung ke dalam organisasi. Cek email kamu untuk
            menerima undangan atau buat organisasi untuk mulai.
          </p>
          <Form form={form} onFinish={onCreate} className="mt-8 flex flex-col">
            <Form.Item
              name="name"
              rules={[{ required: true, message: "Masukan nama organisasi" }]}
            >
              <Input
                size="large"
                placeholder="Nama Organisasi"
                suffix={<TeamOutlined />}
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
