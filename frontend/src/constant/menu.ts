import {
  AppstoreFilled,
  AppstoreOutlined,
  // InboxOutlined,
  SettingFilled,
  SettingOutlined,
  TeamOutlined,
} from "@ant-design/icons";

export const mainMenu = [
  {
    link: "/",
    Ico: AppstoreOutlined,
    IcoActive: AppstoreFilled,
    label: "Dashboard",
  },
  // {
  //   link: "/inventory",
  //   Ico: InboxOutlined,
  //   IcoActive: InboxOutlined,
  //   label: "Inventory",
  // },
  {
    link: "/people",
    Ico: TeamOutlined,
    IcoActive: TeamOutlined,
    label: "Anggota",
  },
];

export const secondaryMenu = [
  {
    link: "/settings",
    Ico: SettingOutlined,
    IcoActive: SettingFilled,
    label: "Pengaturan",
  },
];
