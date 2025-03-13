"use client";

import { ErrorPage } from "@/components";
import { PlusOutlined } from "@ant-design/icons";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { Button, Modal, notification, Skeleton, Table } from "antd";
import { useState } from "react";
import { InviteModal } from "./invite";

import { changeRole, removeMember } from "@/api";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { QUERY_KEYS, useOrganizationMembers, useUserData } from "@/queries";
import { APIError, ChangeRolePayload, Role } from "@/types";
import { tableColumn } from "./data";

export default function People() {
  const { data, isLoading, error, refetch } = useOrganizationMembers();
  const { data: user } = useUserData();
  const [modal, modalContextHolder] = Modal.useModal();
  const [toast, toastContextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const [isInviteModalOpen, setIsInviteModalOpen] = useState(false);
  const onInviteMember = () => setIsInviteModalOpen(true);

  const { mutate: changeMemberRole, isPending: isChangeRolePending } =
    useMutation<void, APIError, ChangeRolePayload>({
      mutationKey: ["changeRole"],
      mutationFn: changeRole,
      onSuccess: () => {
        queryClient.invalidateQueries({
          queryKey: QUERY_KEYS.organizationMembers,
        });
        toast.success({ message: "Role Diubah" });
      },
      onError: (error) => {
        toast.error({
          message: "Gagal mengubah role",
          description: API_MSG[error.message] || SERVER_ERR_MSG,
          placement: "bottomRight",
        });
      },
    });

  const { mutate: remove, isPending: isRemovePending } = useMutation({
    mutationKey: ["removeMember"],
    mutationFn: removeMember,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.organizationMembers,
      });
      toast.success({ message: "Anggota dikeluarkan" });
    },
    onError: (error) => {
      toast.error({
        message: "Gagal mengeluarkan anggota",
        description: API_MSG[error.message] || SERVER_ERR_MSG,
        placement: "bottomRight",
      });
    },
  });

  const onChangeRole = (memberId: number, role: Role) => {
    modal.confirm({
      title: "Ubah Role",
      content: (
        <>
          Apa kamu yakin ingin mengubah Role anggota menjadi{" "}
          <span className="capitalize">{role}</span>
        </>
      ),
      okText: "Ubah Role",
      okButtonProps: {
        loading: isChangeRolePending,
      },
      onOk: () => changeMemberRole({ memberId, role }),
    });
  };

  const onRemoveMember = (id: number) => {
    modal.confirm({
      title: "Keluarkan Anggota",
      content: "Apa kamu yakin ingin mengeluarkan anggota?",
      okText: "Keluarkan",
      okButtonProps: {
        danger: true,
        loading: isRemovePending,
      },
      onOk: () => {
        remove(id);
      },
    });
  };

  if (isLoading) {
    return (
      <div className="flex flex-col gap-8">
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
        <Skeleton round paragraph={{ rows: 4 }} />
      </div>
    );
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
      {modalContextHolder}
      {toastContextHolder}
      <div className="flex flex-col gap-3">
        <div className="flex justify-end gap-3">
          {user?.role !== "staff" && (
            <Button type="primary" onClick={onInviteMember}>
              <PlusOutlined /> Undang
            </Button>
          )}
        </div>
        <div className="w-full overflow-auto">
          <Table
            rowKey="id"
            dataSource={data}
            columns={tableColumn({
              user,
              onRemoveMember,
              isRemovePending,
              onChangeRole,
              isChangeRolePending,
            })}
            pagination={false}
            showSorterTooltip={false}
            className="min-w-[40rem] md:min-w-full"
          />
        </div>
      </div>
      <InviteModal
        open={isInviteModalOpen}
        onCloseAction={() => setIsInviteModalOpen(false)}
      />
    </>
  );
}
