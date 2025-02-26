import {
  DeleteOutlined,
  UserSwitchOutlined,
  MoreOutlined,
} from "@ant-design/icons";
import { Dropdown, TableColumnsType } from "antd";

import { Member, Role, User } from "@/types";

interface Props {
  user?: User;
  onRemoveMember: (memberId: number) => void;
  isRemovePending?: boolean;
  onChangeRole: (memberId: number, role: Role) => void;
  isChangeRolePending?: boolean;
}

export const tableColumn = ({
  user,
  onRemoveMember,
  isRemovePending,
  onChangeRole,
  isChangeRolePending,
}: Props): TableColumnsType<Member> => {
  return [
    {
      title: "ID",
      dataIndex: "id",
      sorter: (a, b) => a.id - b.id,
    },
    {
      title: "Nama",
      dataIndex: "fullname",
      sorter: (a, b) => a.fullname.localeCompare(b.fullname),
    },
    {
      title: "Email",
      dataIndex: "email",
      sorter: (a, b) => a.email.localeCompare(b.email),
    },
    {
      title: "Role",
      dataIndex: "role",
      render: (role: string) => <span className="capitalize">{role}</span>,
      sorter: (a, b) => a.role.localeCompare(b.role),
    },
    {
      render: (_, member) => {
        const isSelfOrAdmin =
          user?.id === member.user_id || member.role === "admin";
        return (
          <Dropdown
            trigger={["click"]}
            menu={{
              items: [
                {
                  key: "1",
                  disabled: isSelfOrAdmin || isRemovePending,
                  className:
                    isSelfOrAdmin || member.role === "admin" ? "!hidden" : "",
                  onClick: () => onChangeRole(member.id, "admin"),
                  icon: <UserSwitchOutlined />,
                  label: "Ubah Role ke Admin",
                },
                {
                  key: "2",
                  disabled: isSelfOrAdmin || isChangeRolePending,
                  className:
                    isSelfOrAdmin || member.role === "supervisor"
                      ? "!hidden"
                      : "",
                  onClick: () => onChangeRole(member.id, "supervisor"),
                  icon: <UserSwitchOutlined />,
                  label: "Ubah Role ke Supervisor",
                },
                {
                  key: "3",
                  disabled: isSelfOrAdmin || isChangeRolePending,
                  className:
                    isSelfOrAdmin || member.role === "staff" ? "!hidden" : "",
                  onClick: () => onChangeRole(member.id, "staff"),
                  icon: <UserSwitchOutlined />,
                  label: "Ubah Role ke Staff",
                },
                {
                  key: "4",
                  disabled: isSelfOrAdmin || isChangeRolePending,
                  onClick: () => onRemoveMember(member.id),
                  icon: <DeleteOutlined />,
                  label: "Keluarkan",
                  className: `${isSelfOrAdmin || isRemovePending ? "" : "!text-red-400"} `,
                },
              ],
            }}
            placement="topRight"
          >
            <button className="hover:bg-gray-200 p-2 rounded-md">
              <MoreOutlined className="text-xl font-bold size-full" />
            </button>
          </Dropdown>
        );
      },
    },
  ];
};
