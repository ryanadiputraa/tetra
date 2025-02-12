"use client";

import { Skeleton, Table, TableColumnsType } from "antd";

import { ErrorPage } from "@/components";
import { useOrganizationMembers } from "@/queries";
import { Member } from "@/types";

export default function People() {
  const { data, isLoading, error, refetch } = useOrganizationMembers();

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
      <div>
        <Table
          rowKey="id"
          dataSource={data}
          columns={columns}
          pagination={false}
          rowSelection={{
            onChange: (keys) => {
              console.log(keys);
            },
          }}
        />
      </div>
    </>
  );
}
