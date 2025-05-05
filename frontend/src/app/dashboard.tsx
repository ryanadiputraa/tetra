"use client";

import { useUtilizationDashboard } from "@/queries/utilization";
import {
  Bar,
  BarChart,
  CartesianGrid,
  Legend,
  ResponsiveContainer,
  Tooltip,
  XAxis,
  YAxis,
} from "recharts";

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

  return (
    <div className="grid grid-cols-1 md:grid-cols-2 2xl:grid-cols-3 gap-8 mt-8">
      {data?.data.map(({ contract, move_type, units }) => {
        const data = units.map((unit) => ({
          label: unit.unit_name,
          "Unit Terdaftar": unit.total_available,
          "Alokasi MPE": unit.alocation_from_mpe,
          Realisasi: unit.realization,
        }));

        return (
          <div
            key={move_type}
            className="bg-white dark:bg-neutral-900 p-4 rounded-md"
          >
            <h6 className="font-semibold text-lg">{move_type}</h6>
            <div className="mt-4">
              <ResponsiveContainer width="100%" height={300}>
                <BarChart data={data}>
                  <CartesianGrid strokeDasharray={"5 5"} />
                  <XAxis dataKey="label" />
                  <YAxis />
                  <Tooltip />
                  <Legend />
                  <Bar dataKey="Unit Terdaftar" fill="#27548A" />
                  <Bar dataKey="Alokasi MPE" fill="#547792" />
                  <Bar dataKey="Realisasi" fill="#94B4C1" />
                </BarChart>
              </ResponsiveContainer>
            </div>
          </div>
        );
      })}
    </div>
  );
};
