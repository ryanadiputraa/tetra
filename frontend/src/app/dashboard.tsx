"use client";

import { useUtilizationDashboard } from "@/queries/utilization";

interface Props {
  date: string;
  startTime: string;
  endTime: string;
}

export const Dashboard = ({ date, startTime, endTime }: Props) => {
  const { data } = useUtilizationDashboard({
    date,
    start_time: startTime,
    end_time: endTime,
  });
  console.log(data);

  return (
    <div className="grid grid-cols-2 gap-8 mt-8">
      {data?.data.map(({ contract, move_type, units }) => (
        <div
          key={move_type}
          className="bg-white dark:bg-neutral-900 p-4 rounded-md"
        >
          <h6 className="font-semibold text-lg">{move_type}</h6>
        </div>
      ))}
    </div>
  );
};
