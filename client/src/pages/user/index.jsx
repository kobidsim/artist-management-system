import { Button, Flex, message, Modal, Popconfirm, Table, Tooltip } from "antd"
import { DeleteFilled, EditFilled } from "@ant-design/icons"
import axios from "axios"
import { useEffect, useState } from "react"
import UserForm from "./form"

export default function UserPage() {
    const [userList, setUserList] = useState([])
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [editData, setEditData] = useState(null)
    const [messageApi, contextHolder] = message.useMessage()

    const listUsers = () => {
        const jwt = localStorage.getItem("jwt")
        axios.get("http://localhost:8080/users", {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            setUserList(res.data.data)
        }).catch((error) => {
            console.log(error)
        })
    }

    const createUser = (data) => {
        const jwt = localStorage.getItem("jwt")
        axios.post("http://localhost:8080/user", data, {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            messageApi.open({
                type: "success",
                content: res?.data?.message,
            })
            setEditData(null)
            setIsModalOpen(false)
            listUsers()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const editUser = (data) => {
        const jwt = localStorage.getItem("jwt")
        axios.post(`http://localhost:8080/user/${editData.id}`, data, {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            messageApi.open({
                type: 'success',
                content: res?.data?.message
            })
            setEditData(null)
            setIsModalOpen(false)
            listUsers()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const deleteUser = (id) => {
        const jwt = localStorage.getItem("jwt")
        axios.delete(`http://localhost:8080/user/${id}`, {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            messageApi.open({
                type: 'success',
                content: res?.data?.message,
            })
            setEditData(null)
            setIsModalOpen(false)
            listUsers()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message
            })
        })
    }

    useEffect(() => {
        listUsers()
    }, [])

    const columns = [
        {
            title: "Name",
            render: (_, record) => {
                return record?.first_name + ' ' + record?.last_name
            },
            key: 'name',
        },
        {
            title: "Role",
            render: (_, row) => {
                let role = ""
                switch (row?.role) {
                    case 'super_admin':
                        role = 'Admin'
                        break;
                    case 'artist_manager':
                        role = 'Artist Manager'
                        break;               
                    case 'artist':
                        role = 'Artist'
                        break;
                    default:
                        break;
                }

                return role
            }
        },
        {
            title: "Email",
            dataIndex: 'email',
            key: 'email',
        },
        {
            title: "Gender",
            render: (_, record) => {
                let gender = "-"
                switch (record?.gender) {
                    case "m":
                        gender = "Male"
                        break;
                    
                    case "f":
                        gender = "Female"
                        break;

                    case "o":
                        gender = "Other"
                        break;
                
                    default:
                        break;
                }

                return gender
            },
            key: 'gender',
        },
        {
            title: "Phone",
            dataIndex: 'phone',
            key: 'phone',
        },
        {
            title: "Address",
            dataIndex: 'address',
            key: 'address'
        },
        {
            title: '',
            dataIndex: '-',
            render: (_, record) => (
                <div
                    style={{
                        display: "flex",
                        justifyContent: "space-around",
                        alignItems: "center",
                        gap: "8px",
                    }}
                >
                    <Tooltip title="Edit">
                        <Button
                            icon={<EditFilled/>}
                            onClick={() => {
                                setEditData(record)
                                setIsModalOpen(true)
                            }}
                        />
                    </Tooltip>
                    <Tooltip title="Delete">
                        <Popconfirm
                            title="Delete User"
                            description="Are you sure you want to delete this user?"
                            onConfirm={() => deleteUser(record?.id)}
                        >
                            <Button icon={<DeleteFilled />} />
                        </Popconfirm>
                    </Tooltip>
                </div>
            )
        }
    ]

    return(
        <>
            {contextHolder}
            <Button type="primary" onClick={() => setIsModalOpen(true)}>Create</Button>
            <Modal
                title={!!editData ? "Edit User" : "Create User"}
                open={isModalOpen}
                onCancel={() => {
                    setEditData(null)
                    setIsModalOpen(false)
                }}
                footer={false}
                destroyOnClose
            >
                <UserForm
                    isEdit={!!editData}
                    editData={editData}
                    onCreate={(values) => {
                        createUser(values)
                    }}
                    onEdit={(values) => {
                        editUser(values)
                    }}
                />
            </Modal>
            <Table
                dataSource={userList}
                columns={columns}
                pagination={{
                    position: "bottomRight",
                    defaultPageSize: 10,
                }}
            />
        </>
    )
}