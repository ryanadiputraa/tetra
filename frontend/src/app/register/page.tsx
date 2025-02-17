"use client";

import { useMutation } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import { AiOutlineLine } from "react-icons/ai";
import { FcGoogle } from "react-icons/fc";

import { register } from "@/api";
import { API_MSG, API_URL, COOKIE_AUTH_KEY, SERVER_ERR_MSG } from "@/constant";
import { getCookie, setCookie } from "@/lib";
import { APIError, JWT, RegisterPayload, RegisterPayloadForm } from "@/types";

export default function Login() {
  const router = useRouter();
  const [toast, contextHolder] = notification.useNotification();
  const [form] = Form.useForm<RegisterPayloadForm>();

  const { mutate, isPending } = useMutation<JWT, APIError, RegisterPayload>({
    mutationKey: ["register"],
    mutationFn: register,
    onSuccess: (data) => {
      setCookie(COOKIE_AUTH_KEY, data.access_token, new Date(data.expires_at));
      router.push("/");
    },
    onError: (error) => {
      if (!error.errors) {
        toast.error({
          message: "Gagal",
          description: API_MSG[error.message] || SERVER_ERR_MSG,
          placement: "bottomRight",
        });
      } else {
        const fields = [];
        for (const e in error.errors) {
          fields.push({
            name: e as keyof RegisterPayloadForm,
            errors: [API_MSG[error.errors[e]]],
          });
        }
        form.setFields(fields);
      }
    },
  });

  const onRegister = ({
    fullname,
    email,
    password,
    confirm_password,
  }: RegisterPayloadForm) => {
    if (password !== confirm_password) {
      form.setFields([
        { name: "confirm_password", errors: ["Password tidak valid"] },
      ]);
      return;
    }
    mutate({ fullname, email, password });
  };

  useEffect(() => {
    const token = getCookie(COOKIE_AUTH_KEY);
    if (token) {
      router.push("/");
    }
  }, [router]);

  return (
    <div className="min-h-screen grid place-items-center px-8">
      {contextHolder}
      <div className="bg-white py-16 px-8 sm:px-16 rounded-lg w-full sm:max-w-lg">
        <div className="text-center">
          <h4 className="font-bold text-2xl">Inventra</h4>
          <p className="mt-2">Daftar akun untuk masuk ke dashboard Inventra.</p>
        </div>
        <Form form={form} onFinish={onRegister} className="mt-8 flex flex-col">
          <label className="mb-1 font-semibold">Nama Lengkap</label>
          <Form.Item name="fullname" rules={[{ required: true }]}>
            <Input size="large" placeholder="John Doe" type="text" />
          </Form.Item>
          <label className="mb-1 font-semibold">Email</label>
          <Form.Item name="email" rules={[{ required: true }]}>
            <Input size="large" placeholder="john@mail.com" type="email" />
          </Form.Item>
          <label className="mb-1 font-semibold">Password</label>
          <Form.Item name="password" rules={[{ required: true }]}>
            <Input size="large" placeholder="Password" type="password" />
          </Form.Item>
          <label className="mb-1 font-semibold">Konfirmasi Password</label>
          <Form.Item name="confirm_password" rules={[{ required: true }]}>
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
            Daftar
          </Button>
        </Form>
        <p className="my-4 flex items-center justify-center gap-2">
          <AiOutlineLine /> Atau daftar dengan <AiOutlineLine />
        </p>
        <a href={API_URL + "/oauth/login/google"}>
          <Button
            disabled={isPending}
            size="large"
            className="font-semibold w-full"
          >
            <FcGoogle className="text-xl" /> Google
          </Button>
        </a>
        <p className="text-center mt-8">
          Sudah punya akun?{" "}
          <Link href="/login" className="font-semibold text-blue-600">
            Masuk di sini.
          </Link>
        </p>
      </div>
    </div>
  );
}
