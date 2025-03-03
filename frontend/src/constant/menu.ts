import {
  AppstoreFilled,
  AppstoreOutlined,
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
  {
    link: "/people",
    Ico: TeamOutlined,
    IcoActive: TeamOutlined,
    label: "People",
  },
];

export const secondaryMenu = [
  {
    link: "/settings",
    Ico: SettingOutlined,
    IcoActive: SettingFilled,
    label: "Settings",
  },
];
