"use client";

import { Spin } from "antd";

export const Loader = () => {
  return (
    <div className="min-h-screen grid place-items-center">
      <Spin size="large" />
    </div>
  );
};
