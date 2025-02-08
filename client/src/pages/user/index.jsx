import { Table } from "antd"
import axios from "axios"
import { useEffect, useState } from "react"

export default function UserPage() {
    const [userList, setUserList] = useState([])

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
        }
    ]

    return(
        <>
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