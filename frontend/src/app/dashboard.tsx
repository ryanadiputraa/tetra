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
    <div className="grid grid-cols-1 md:grid-cols-2 3xl:grid-cols-3 gap-8 mt-8">
      {data?.data.map(({ move_type, realization, units }) => {
        const realizationChart = realization.map((unit) => ({
          label: unit.unit_name,
          "Unit Terdaftar": unit.total_available,
          "Alokasi MPE": unit.alocation_from_mpe,
          Realisasi: unit.realization,
        }));

        const unitsChart = [
          { label: "OPR Engine ON", Jumlah: units.opr_engine_on },
          { label: "IDLE Engine ON", Jumlah: units.idle_engine_on },
          { label: "STD Engine OFF", Jumlah: units.std_engine_off },
          { label: "STD Engine Unkown", Jumlah: units.std_engine_unkown },
          { label: "BD Engine ON", Jumlah: units.bd_engine_on },
          { label: "BD Engine OFF", Jumlah: units.bd_engine_off },
          { label: "BD Engine Unkown", Jumlah: units.bd_engine_unkown },
          { label: "ACD Engine Unkown", Jumlah: units.acd_engine_unkown },
          {
            label: "Chasis Crack Engine Unkown",
            Jumlah: units.chasis_crack_engine_unkown,
          },
        ];

        return (
          <div
            key={move_type}
            className="bg-white dark:bg-neutral-900 p-4 rounded-md"
          >
            <h6 className="font-semibold text-lg">{move_type}</h6>
            <div className="mt-4 flex gap-3">
              <ResponsiveContainer width="50%" height={500}>
                <BarChart data={unitsChart} layout="vertical">
                  <CartesianGrid strokeDasharray={"5 5"} />
                  <XAxis type="number" />
                  <YAxis
                    dataKey="label"
                    type="category"
                    tick={{ fontSize: 12 }}
                  />
                  <Tooltip />
                  <Legend />
                  <Bar dataKey="Jumlah" fill="#27548A" />
                </BarChart>
              </ResponsiveContainer>
              <div className="w-1/2 flex flex-col gap-3">
                <p className="font-medium text-center">
                  Utilisasi DT Plan dan Realisasi
                </p>
                <ResponsiveContainer width="100%" height={300}>
                  <BarChart data={realizationChart}>
                    <CartesianGrid strokeDasharray={"5 5"} />
                    <XAxis dataKey="label" type="category" />
                    <YAxis type="number" />
                    <Tooltip />
                    <Legend />
                    <Bar dataKey="Unit Terdaftar" fill="#27548A" />
                    <Bar dataKey="Alokasi MPE" fill="#547792" />
                    <Bar dataKey="Realisasi" fill="#94B4C1" />
                  </BarChart>
                </ResponsiveContainer>
              </div>
            </div>
          </div>
        );
      })}
    </div>
  );
};
