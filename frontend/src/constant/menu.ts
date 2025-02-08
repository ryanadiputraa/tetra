import {
  AiFillAppstore,
  AiFillSchedule,
  AiOutlineAppstore,
  AiOutlineSchedule,
} from "react-icons/ai";

export const mainMenu = [
  {
    link: "/",
    Ico: AiOutlineAppstore,
    IcoActive: AiFillAppstore,
    label: "Dashboard",
  },
  {
    link: "/timesheet",
    Ico: AiOutlineSchedule,
    IcoActive: AiFillSchedule,
    label: "Timesheet",
  },
];
