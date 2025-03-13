"use client";

import { ErrorPage } from "@/components";
import { Skeleton, Table } from "antd";
import { useSearchParams, useRouter } from "next/navigation";
import { tableColumn } from "./data";

import { QUERY_KEYS } from "@/queries";
import { useInventoryItem } from "@/queries/inventory";

export default function Inventory() {
  const searchParams = useSearchParams();
  const router = useRouter();
  const page =
    Number(searchParams.get("page")) > 0 ? Number(searchParams.get("page")) : 1;
  const { data, isLoading, error, refetch } = useInventoryItem(
    {
      page,
      size: 50,
    },
    { queryKey: [...QUERY_KEYS.inventoryItems, page] },
  );

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
      {/* {modalContextHolder} */}
      {/* {toastContextHolder} */}
      <div className="flex flex-col gap-3">
        <div className="flex justify-end gap-3">
          {/* {user?.role !== "staff" && ( */}
          {/*   <Button type="primary" onClick={onInviteMember}> */}
          {/*     <PlusOutlined /> Undang */}
          {/*   </Button> */}
          {/* )} */}
        </div>
        <div className="w-full overflow-auto">
          <Table
            rowKey="id"
            dataSource={data?.items}
            columns={tableColumn()}
            pagination={{
              current: data?.meta?.current_page,
              pageSize: data?.meta?.size,
              total: data?.meta?.total_data,
              onChange: (page) => {
                const params = new URLSearchParams(searchParams.toString());
                params.set("page", page.toString());
                router.push(`?${params.toString()}`);
              },
            }}
            showSorterTooltip={false}
            className="min-w-[40rem] md:min-w-full"
          />
        </div>
      </div>
    </>
  );
}
