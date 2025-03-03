"use client";

import { ErrorPage } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Modal, notification, Skeleton } from "antd";
import { useRouter } from "next/navigation";

import { leaveOrganization } from "@/api/organization";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { formatDate } from "@/lib/utils";
import { QUERY_KEYS, useOrganization } from "@/queries";
import { APIError } from "@/types";

export default function Settings() {
  const { data, isLoading, error, refetch } = useOrganization();
  const [toast, contextHolder] = notification.useNotification();
  const [modal, modalContextHolder] = Modal.useModal();
  const router = useRouter();
  const queryClient = useQueryClient();

  const { mutate: leave, isPending: isLeavePending } = useMutation<
    void,
    APIError
  >({
    mutationKey: ["leaveOrganization"],
    mutationFn: leaveOrganization,
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.userData });
      router.push("/join");
    },
    onError: (error) => {
      toast.error({
        message: "Gagal",
        description: API_MSG[error.message] || SERVER_ERR_MSG,
        placement: "bottomRight",
      });
    },
  });

  const onLeave = () => {
    modal.confirm({
      title: "Keluar Dari Organisasi",
      content: "Apa kamu yakin ingin keluar dari organisasi",
      okText: "Keluar",
      okButtonProps: {
        danger: true,
        loading: isLeavePending,
      },
      onOk: () => leave(),
    });
  };

  if (isLoading) {
    return <Skeleton avatar round paragraph={{ rows: 4 }} />;
  }
  if (error) {
    return <ErrorPage loading={isLoading} onRetry={() => refetch()} />;
  }

  return (
    <>
      {contextHolder}
      {modalContextHolder}
      <div className="flex flex-col gap-4 max-w-xl py-16 mx-auto">
        <div className="flex gap-4 items-center">
          <div className="size-20 grid place-items-center bg-primary rounded-full">
            <span className="text-3xl text-white font-bold">
              {data?.name.split("")[0]}
            </span>
          </div>
          <div className="flex flex-col">
            <p className="text-2xl font-semibold">{data?.name}</p>
            <p>Owner: {data?.owner.fullname}</p>
            <span className="italic text-slate-400 capitalize">
              Dibuat pada {formatDate(data?.subscription_end_at, "full")}
            </span>
          </div>
        </div>
        <section className="mt-8">
          <h6 className="my-4 border-b-2 border-gray-200 text-lg font-medium">
            Organisasi
          </h6>
          <Button
            danger
            size="large"
            loading={isLeavePending}
            onClick={onLeave}
            className="w-full"
          >
            Keluar Dari Organisasi
          </Button>
        </section>
      </div>
    </>
  );
}
