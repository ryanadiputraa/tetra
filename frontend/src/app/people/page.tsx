"use client";

import { ErrorPage } from "@/components";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import {
  Button,
  Dropdown,
  Modal,
  notification,
  Skeleton,
  Table,
  TableColumnsType,
} from "antd";
import { useState } from "react";
import { AiOutlineDelete, AiOutlineMore, AiOutlinePlus } from "react-icons/ai";
import { InviteModal } from "./invite";

import { removeMember } from "@/api/organization";
import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { QUERY_KEYS, useOrganizationMembers, useUserData } from "@/queries";
import { Member } from "@/types";

export default function People() {
  const { data, isLoading, error, refetch } = useOrganizationMembers();
  const { data: user } = useUserData();
  const [modal, modalContextHolder] = Modal.useModal();
  const [toast, toastContextHolder] = notification.useNotification();
  const queryClient = useQueryClient();

  const [isInviteModalOpen, setIsInviteModalOpen] = useState(false);

  const { mutate, isPending } = useMutation({
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

  const onInviteMember = () => setIsInviteModalOpen(true);
  const onRemoveMember = (id: number) => {
    modal.confirm({
      title: "Keluarkan Anggota",
      content: "Apa kamu yakin ingin mengeluarkan anggota?",
      okText: "Keluarkan",
      okButtonProps: {
        danger: true,
        loading: isPending,
      },
      onOk: () => {
        mutate(id);
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
    {
      render: (_, member) => (
        <Dropdown
          disabled={user?.role !== "admin"}
          trigger={["click"]}
          menu={{
            items: [
              {
                key: "1",
                disabled: user?.id === member.user_id || isPending,
                onClick: () => onRemoveMember(member.id),
                icon: <AiOutlineDelete />,
                label: "Keluarkan",
              },
            ],
          }}
          placement="topRight"
        >
          <button className="hover:bg-gray-200 p-2 rounded-md">
            <AiOutlineMore className="text-xl font-bold size-full" />
          </button>
        </Dropdown>
      ),
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
    return <ErrorPage loading={isLoading} onRetry={() => refetch()} />;
  }

  return (
    <>
      {modalContextHolder}
      {toastContextHolder}
      <div className="flex flex-col gap-3">
        <div className="flex justify-end gap-3">
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
        />
      </div>
      <InviteModal
        open={isInviteModalOpen}
        onCloseAction={() => setIsInviteModalOpen(false)}
      />
    </>
  );
}
