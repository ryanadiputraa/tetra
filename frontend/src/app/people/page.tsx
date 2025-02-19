"use client";

import { ErrorPage } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  Button,
  Modal,
  notification,
  Skeleton,
  Table,
  TableColumnsType,
} from "antd";
import { useState } from "react";
import { AiOutlineDelete, AiOutlinePlus } from "react-icons/ai";
import { InviteModal } from "./invite";

import { removeMember } from "@/api/organization";
import { QUERY_KEYS, useOrganizationMembers, useUserData } from "@/queries";
import { Member } from "@/types";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";

export default function People() {
  const { data, isLoading, error, refetch } = useOrganizationMembers();
  const { data: user } = useUserData();
  const [modal, modalContextHolder] = Modal.useModal();
  const [toast, toastContextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const [isInviteModalOpen, setIsInviteModalOpen] = useState(false);
  const [selectedMember, setSelectedMember] = useState<React.Key[]>([]);

  const { mutate, isPending } = useMutation({
    mutationKey: ["removeMember"],
    mutationFn: removeMember,
    onSuccess: () => {
      queryClient.invalidateQueries({
        queryKey: QUERY_KEYS.organizationMembers,
      });
      toast.success({ message: "Anggota dihapus" });
    },
    onError: (error) => {
      toast.error({
        message: "Gagal menghapus anggota",
        description: API_MSG[error.message] || SERVER_ERR_MSG,
        placement: "bottomRight",
      });
    },
  });

  const onInviteMember = () => setIsInviteModalOpen(true);
  const onRemoveMember = () => {
    modal.confirm({
      title: "Hapus Anggota",
      content: "Apa kamu yakin ingin menghapus anggota?",
      okText: "Hapus",
      okButtonProps: {
        danger: true,
        loading: isPending,
      },
      onOk: () => {
        const ids = selectedMember.join(",");
        mutate(ids);
      },
    });
  };

  const columns: TableColumnsType<Member> = [
    {
      title: "ID",
      dataIndex: "id",
      sorter: (a, b) => a.id - b.id,
    },
    {
      title: "Nama",
      dataIndex: "fullname",
      sorter: (a, b) => a.fullname.localeCompare(b.fullname),
    },
    {
      title: "Email",
      dataIndex: "email",
      sorter: (a, b) => a.email.localeCompare(b.email),
    },
    {
      title: "Role",
      dataIndex: "role",
      render: (role: string) => <span className="capitalize">{role}</span>,
      sorter: (a, b) => a.role.localeCompare(b.role),
    },
  ];

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
    return <ErrorPage onRetry={() => refetch()} />;
  }

  return (
    <>
      {modalContextHolder}
      {toastContextHolder}
      <div className="flex flex-col gap-3">
        <div className="flex justify-end gap-3">
          {user?.role === "admin" && (
            <Button
              disabled={!selectedMember.length}
              danger
              variant="outlined"
              onClick={onRemoveMember}
            >
              <AiOutlineDelete />
              Hapus
            </Button>
          )}
          {user?.role !== "staff" && (
            <Button
              type="primary"
              className="font-semibold"
              onClick={onInviteMember}
            >
              <AiOutlinePlus /> Undang
            </Button>
          )}
        </div>
        <Table
          rowKey="id"
          dataSource={data}
          columns={columns}
          pagination={false}
          showSorterTooltip={false}
          rowSelection={{
            getCheckboxProps: (member) => ({
              disabled: member.user_id === user?.id,
            }),
            onChange: (keys) => setSelectedMember(keys),
          }}
        />
      </div>
      <InviteModal
        open={isInviteModalOpen}
        onCloseAction={() => setIsInviteModalOpen(false)}
      />
    </>
  );
}
