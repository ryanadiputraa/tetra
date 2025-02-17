"use client";

import { useQueryClient } from "@tanstack/react-query";
import { useParams, useRouter } from "next/navigation";
import { useEffect } from "react";

import { API_MSG, SERVER_ERR_MSG } from "@/constant";
import { QUERY_KEYS, useAcceptInvitation } from "@/queries";

export default function AcceptInvitation() {
  const { code } = useParams();
  const router = useRouter();
  const { data, error } = useAcceptInvitation(String(code));
  const queryClient = useQueryClient();

  useEffect(() => {
    if (data) {
      queryClient.invalidateQueries({ queryKey: QUERY_KEYS.userData });
      router.push("/");
    }
  }, [data, router, queryClient]);

  return (
    <div className="min-h-screen bg-slate-200 grid place-items-center">
      <p>
        {error ? API_MSG[error.message] || SERVER_ERR_MSG : "Authenticating..."}
      </p>
    </div>
  );
}
