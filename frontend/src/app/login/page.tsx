"use client";

import { useMutation } from "@tanstack/react-query";
import { Button, Input, notification } from "antd";
import { useRouter } from "next/navigation";
import { FormEvent, useEffect, useState } from "react";
import {
  AiOutlineEye,
  AiOutlineEyeInvisible,
  AiOutlineLine,
  AiOutlineUser,
} from "react-icons/ai";
import { FcGoogle } from "react-icons/fc";

import { login } from "@/api";
import { COOKIE_AUTH_KEY, SERVER_ERR, SERVER_ERR_MSG } from "@/constant";
import { getCookie, setCookie } from "@/lib";
import { APIError, JWT, LoginPayload } from "@/types";

export default function Login() {
  const router = useRouter();
  const [toast, contextHolder] = notification.useNotification();
  const [payload, setPayload] = useState<LoginPayload>({
    email: "",
    password: "",
  });

  const { mutate, isPending } = useMutation<JWT, APIError, LoginPayload>({
    mutationKey: ["login"],
    mutationFn: login,
    onSuccess: (data) => {
      setCookie(COOKIE_AUTH_KEY, data.access_token, new Date(data.expires_at));
      router.push("/");
    },
    onError: (error) => {
      const description =
        error.message === SERVER_ERR || !error.errors
          ? SERVER_ERR_MSG
          : "Mohon periksa email dan password yang anda masukan.";
      toast.error({
        message: "Login Gagal",
        description,
        placement: "topRight",
      });
    },
  });

  const onLogin = (e: FormEvent<HTMLFormElement>) => {
    e.preventDefault();
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
      <div className="bg-white py-16 px-8 sm:px-16 rounded-xl w-full sm:max-w-lg">
        <div className="text-center">
          <h4 className="font-bold text-2xl">Inventra</h4>
          <p className="mt-2">Login untuk masuk ke dashboard Inventra.</p>
        </div>
        <form onSubmit={onLogin} className="mt-8 flex flex-col gap-4">
          <Input
            required
            size="large"
            placeholder="Email"
            type="email"
            suffix={<AiOutlineUser />}
            onChange={(e) =>
              setPayload((prev) => ({
                ...prev,
                email: e.target.value,
              }))
            }
          />
          <Input.Password
            required
            size="large"
            placeholder="Password"
            iconRender={(visible) =>
              visible ? <AiOutlineEye /> : <AiOutlineEyeInvisible />
            }
            onChange={(e) =>
              setPayload((prev) => ({
                ...prev,
                password: e.target.value,
              }))
            }
          />
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
        </form>
        <p className="my-4 flex items-center justify-center gap-2">
          <AiOutlineLine /> Atau masuk dengan <AiOutlineLine />
        </p>
        <a href="https://google.com">
          <Button
            disabled={isPending}
            size="large"
            className="font-semibold w-full"
          >
            <FcGoogle className="text-xl" /> Google
          </Button>
        </a>
      </div>
    </div>
  );
}
