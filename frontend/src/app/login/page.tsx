"use client";

import { useMutation } from "@tanstack/react-query";
import { Button, Form, Input, notification } from "antd";
import Link from "next/link";
import { useRouter } from "next/navigation";
import { useEffect } from "react";
import {
  AiOutlineEye,
  AiOutlineEyeInvisible,
  AiOutlineLine,
  AiOutlineUser,
} from "react-icons/ai";
import { FcGoogle } from "react-icons/fc";

import { login } from "@/api";
import {
  API_URL,
  COOKIE_AUTH_KEY,
  LS_INVITATION_CODE_KEY,
  SERVER_ERR,
  SERVER_ERR_MSG,
} from "@/constant";
import { fetcher, getCookie, setCookie } from "@/lib";
import { APIError, JWT, LoginPayload } from "@/types";

export default function Login() {
  const router = useRouter();
  const [toast, contextHolder] = notification.useNotification();
  const [form] = Form.useForm<LoginPayload>();

  const { mutate, isPending } = useMutation<JWT, APIError, LoginPayload>({
    mutationKey: ["login"],
    mutationFn: login,
    onSuccess: (data) => {
      setCookie(COOKIE_AUTH_KEY, data.access_token, new Date(data.expires_at));
      fetcher.defaults.headers.common["Authorization"] =
        `Bearer ${data.access_token}`;
      const code = window.localStorage.getItem(LS_INVITATION_CODE_KEY);
      router.push(code ? `/join/${code}` : "/");
    },
    onError: (error) => {
      const description =
        !error.errors && error.message === SERVER_ERR
          ? SERVER_ERR_MSG
          : "Mohon periksa email dan password yang kamu masukan.";
      toast.error({
        message: "Login Gagal",
        description,
        placement: "bottomRight",
      });
    },
  });

  const onLogin = (payload: LoginPayload) => {
    mutate(payload);
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
          <p className="mt-2">Login untuk masuk ke dashboard Inventra.</p>
        </div>
        <Form form={form} onFinish={onLogin} className="mt-8 flex flex-col">
          <Form.Item name="email" rules={[{ required: true }]}>
            <Input
              size="large"
              placeholder="Email"
              type="email"
              suffix={<AiOutlineUser />}
            />
          </Form.Item>
          <Form.Item name="password" rules={[{ required: true }]}>
            <Input.Password
              size="large"
              placeholder="Password"
              iconRender={(visible) =>
                visible ? <AiOutlineEye /> : <AiOutlineEyeInvisible />
              }
            />
          </Form.Item>
          <Button
            htmlType="submit"
            size="large"
            variant="solid"
            color="primary"
            loading={isPending}
            className="font-semibold"
          >
            Login
          </Button>
        </Form>
        <p className="my-4 flex items-center justify-center gap-2">
          <AiOutlineLine /> Atau masuk dengan <AiOutlineLine />
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
          Belum punya akun?{" "}
          <Link href="/register" className="font-semibold text-blue-600">
            Daftar di sini.
          </Link>
        </p>
      </div>
    </div>
  );
}
