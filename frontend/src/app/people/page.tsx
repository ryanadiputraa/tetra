"use client";

import { ErrorPage } from "@/components";
import { Button, Skeleton, Table, TableColumnsType } from "antd";
import { useState } from "react";
import { AiOutlinePlus } from "react-icons/ai";
import { InviteModal } from "./invite";

import { useOrganizationMembers, useUserData } from "@/queries";
import { Member } from "@/types";

export default function People() {
  const { data, isLoading, error, refetch } = useOrganizationMembers();
  const [isInviteModalOpen, setIsInviteModalOpen] = useState(false);
  const { data: user } = useUserData();

  const onInviteMember = () => setIsInviteModalOpen(true);

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
      <div className="flex flex-col gap-3">
        <div className="flex justify-end">
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
            onChange: (keys) => {
              console.log(keys);
            },
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
