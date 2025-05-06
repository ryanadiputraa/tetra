export type Utilizations = {
  data: MoveType[];
};

export type MoveType = {
  contract: string;
  move_type: string;
  realization: Realization[];
  units: UnitData;
};

export type Realization = {
  unit_name: string;
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

export type UnitData = {
  opr_engine_on: number;
  idle_engine_on: number;
  std_engine_off: number;
  std_engine_unkown: number;
  bd_engine_on: number;
  bd_engine_off: number;
  bd_engine_unkown: number;
  acd_engine_unkown: number;
  chasis_crack_engine_unkown: number;
};

export type fetchDashboardParams = {
  date: string;
  start_time: string;
  end_time: string;
};
