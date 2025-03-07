"use client";

import { Loader } from "@/components";
import { useQueryClient } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { useEffect } from "react";

import { API_MSG, LS_INVITATION_CODE_KEY, SERVER_ERR_MSG } from "@/constant";
import { QUERY_KEYS, useAcceptInvitation } from "@/queries";

export default function AcceptInvitation() {
  const { code } = useParams();
  const router = useRouter();
  const { data, isPending, error } = useAcceptInvitation(String(code));
  const queryClient = useQueryClient();

  useEffect(() => {
    if (data) {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.userData });
      window.localStorage.removeItem(LS_INVITATION_CODE_KEY);
      router.push("/");
    }
  }, [data, router, queryClient]);

  if (isPending) {
    return <Loader />;
  }

  if (error) {
    <div className="min-h-screen grid place-items-center">
      <p>{API_MSG[error.message] || SERVER_ERR_MSG}</p>
    </div>;
  }

  return <></>;
}
