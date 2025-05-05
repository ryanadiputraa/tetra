export type Utilizations = {
  data: MoveType[];
};

export type MoveType = {
  contract: string;
  move_type: string;
  units: Unit[];
};

export type Unit = {
  total_available: number;
  alocation_from_mpe: number;
  realization: number;
  additional_info: string[];
  bd: number;
  accident: number;
  tlo: number;
  cms: number;
  gas_diff: number;
  standby: number;
};

export type fetchDashboardParams = {
  date: string;
  start_time: string;
  end_time: string;
};
