"use client";

import { ContentSkeleton, ErrorPage } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";

import { changePassword } from "@/api";
import { API_MSG } from "@/constant";
import { QUERY_KEYS, useUserData } from "@/queries";
import { APIError, ChangePasswordForm } from "@/types";

export default function Profile() {
  const { data, isLoading, error, refetch } = useUserData();
  const [toast, contextHolder] = notification.useNotification();
  const [form] = Form.useForm<ChangePasswordForm>();
  const queryClient = useQueryClient();

  const { mutate, isPending } = useMutation<void, APIError, string>({
    mutationKey: ["changePassword"],
    mutationFn: changePassword,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.userData });
      toast.success({
        message: "Password berhasil diubah",
        placement: "bottomRight",
      });
      form.resetFields();
    },
    onError: (error) => {
      if (!error.errors) {
        toast.error({
          message: "Gagal",
          description: API_MSG[error.message],
          placement: "bottomRight",
        });
      } else {
        form.setFields([
          { name: "password", errors: [error.errors?.password] },
        ]);
      }
    },
  });

  if (isLoading) {
    return <ContentSkeleton length={5} />;
  }
  if (error) {
    return <ErrorPage onRetry={() => refetch()} />;
  }

  const onChangePassword = ({ password, confirm }: ChangePasswordForm) => {
    if (password !== confirm) {
      form.setFields([{ name: "confirm", errors: ["Password tidak valid"] }]);
      return;
    }
    mutate(password);
  };

  return (
    <>
      {contextHolder}
      <div className="flex flex-col gap-4 max-w-xl py-16 mx-auto">
        <div className="flex gap-4 items-center">
          <div className="size-20 grid place-items-center bg-primary rounded-full">
            <span className="text-3xl text-white font-bold">
              {data?.fullname.split("")[0]}
            </span>
          </div>
          <div className="flex flex-col">
            <p className="text-2xl font-semibold">{data?.fullname}</p>
            <p>{data?.email}</p>
            <span className="italic text-slate-400">{data?.role}</span>
          </div>
        </div>
        <Form
          form={form}
          onFinish={onChangePassword}
          className="mt-8 flex flex-col"
        >
          <label className="mb-1 font-semibold">Password</label>
          <Form.Item name="password" rules={[{ required: true, message: "" }]}>
            <Input size="large" placeholder="Password" type="password" />
          </Form.Item>
          <label className="mb-1 font-semibold">Konfirmasi Password</label>
          <Form.Item name="confirm" rules={[{ required: true, message: "" }]}>
            <Input size="large" placeholder="Password" type="password" />
          </Form.Item>
          <Button
            htmlType="submit"
            size="large"
            variant="solid"
            color="primary"
            loading={isPending}
            className="font-semibold"
          >
            Simpan
          </Button>
        </Form>
      </div>
    </>
  );
}
