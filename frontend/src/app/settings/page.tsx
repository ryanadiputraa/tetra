"use client";

import { ErrorPage } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Modal, notification, Skeleton } from "antd";
import { useRouter } from "next/navigation";

import { deleteOrganization, leaveOrganization } from "@/api/organization";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { formatDate } from "@/lib";
import { QUERY_KEYS, useOrganization, useUserData } from "@/queries";
import { APIError } from "@/types";

export default function Settings() {
  const { data, isLoading, error, refetch } = useOrganization();
  const { data: user } = useUserData();
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

  const { mutate: deleteOrg, isPending: isDeletePending } = useMutation<
    void,
    APIError
  >({
    mutationKey: ["deleteOrganization"],
    mutationFn: deleteOrganization,
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
      content: "Apa kamu yakin ingin keluar dari organisasi?",
      okText: "Keluar",
      okButtonProps: {
        danger: true,
        loading: isLeavePending,
      },
      onOk: () => leave(),
    });
  };

  const onDelete = () => {
    modal.confirm({
      title: "Hapus Organisasi",
      content: "Apa kamu yakin menghapus organisasi?",
      okText: "Hapus",
      okButtonProps: {
        danger: true,
        loading: isDeletePending,
      },
      onOk: () => deleteOrg(),
    });
  };

  if (isLoading) {
    return <Skeleton avatar round paragraph={{ rows: 4 }} />;
  }
  if (error) {
    return (
      <ErrorPage
        loading={isLoading}
        onRetryAction={() => refetch()}
        msg={error.message}
      />
    );
  }

  return (
    <>
      {contextHolder}
      {modalContextHolder}
      <div className="flex flex-col gap-4 max-w-4xl py-16 mx-auto">
        <div className="flex gap-4 items-center">
          <div className="size-20 grid place-items-center bg-primary dark:bg-primary-dark rounded-full">
            <span className="text-3xl text-white font-bold">
              {data?.name.split("")[0]}
            </span>
          </div>
          <div className="flex flex-col">
            <p className="text-2xl font-semibold">{data?.name}</p>
            <p>Owner: {data?.owner.fullname}</p>
            <span className="italic text-sm text-neutral-400 capitalize">
              Aktif Hingga: {formatDate(data?.subscription_end_at, "full")}
            </span>
          </div>
        </div>
        <section className="mt-8">
          <h6 className="my-4 border-b-2 border-gray-200 dark:border-gray-500 text-lg font-medium">
            Organisasi
          </h6>
          <div className="flex flex-col gap-6">
            <div className="flex justify-between items-center gap-12">
              <div>
                <p>Keluar Dari Organisasi</p>
                <p className="text-neutral-400 text-sm">
                  Kamu tidak dapat mengakses dashboard organisasi dan hanya
                  dapat bergabung kembali setelah mendapat undangan bergabung.
                </p>
              </div>
              <Button danger loading={isLeavePending} onClick={onLeave}>
                Keluar Dari Organisasi
              </Button>
            </div>
            {user?.id === data?.owner.id && (
              <div className="flex justify-between items-center gap-12">
                <div>
                  <p>Hapus Organisasi</p>
                  <p className="text-neutral-400 text-sm">
                    Seluruh data Organisasi akan dihapus dan tidak dapat
                    dikembalikan.
                  </p>
                </div>
                <Button
                  danger
                  type="primary"
                  loading={isDeletePending}
                  onClick={onDelete}
                >
                  Hapus Organisasi
                </Button>
              </div>
            )}
          </div>
        </section>
      </div>
    </>
  );
}
