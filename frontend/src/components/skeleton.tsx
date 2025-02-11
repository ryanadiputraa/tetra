"use client";

import { MdOutlinePerson } from "react-icons/md";

export function ProfileSkeleton(): React.ReactNode {
  return (
    <div className="flex items-center gap-2 animate-pulse">
      <div className="bg-gray-200 p-2 rounded-full grid place-items-center">
        <MdOutlinePerson className="text-xl" />
      </div>
      <div className="flex flex-col gap-2">
        <div className="h-2 bg-gray-200 rounded-full w-48"></div>
        <div className="h-2 bg-gray-200 rounded-full w-36"></div>
      </div>
    </div>
  );
}

export function ContentSkeleton({
  length = 1,
  isMultiLine = false,
}: {
  length?: number;
  isMultiLine?: boolean;
}): React.ReactNode {
  const content = new Array(length).fill(0);
  return (
    <div className="flex flex-col gap-8 animate-pulse w-full">
      {content.map((_, i) => (
        <div key={i} className="flex flex-col gap-2">
          <div className="h-4 bg-gray-200 rounded-md w-1/2"></div>
          <div
            className={`
              bg-gray-200 rounded-md w-full ${isMultiLine ? "h-40" : "h-4"}`}
          ></div>
          <div className="h-4 bg-gray-200 rounded-md w-1/3"></div>
        </div>
      ))}
    </div>
  );
}
