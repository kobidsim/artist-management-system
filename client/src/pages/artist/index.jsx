import { Button, Flex, message, Modal, Popconfirm, Table, Tooltip } from "antd"
import { DeleteFilled, EditFilled } from "@ant-design/icons"
import axios from "axios"
import { useEffect, useState } from "react"
import ArtistForm from "./form"

export default function ArtistPage() {
    const [artistList, setArtistList] = useState([])
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [editData, setEditData] = useState(null)
    const [messageApi, contextHolder] = message.useMessage()

    const listArtists = () => {
        const jwt = localStorage.getItem("jwt")
        axios.get("http://localhost:8080/artists", {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            setArtistList(res.data.data)
        }).catch((error) => {
            console.log(error)
        })
    }

    const createArtist = (data) => {
        const jwt = localStorage.getItem("jwt")
        axios.post("http://localhost:8080/artist", data, {
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
            listArtists()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const editArtist = (data) => {
        const jwt = localStorage.getItem("jwt")
        axios.post(`http://localhost:8080/artist/${editData.id}`, data, {
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
            listArtists()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const deleteArtist = (id) => {
        const jwt = localStorage.getItem("jwt")
        axios.delete(`http://localhost:8080/artist/${id}`, {
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
            listArtists()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message
            })
        })
    }

    useEffect(() => {
        listArtists()
    }, [])

    const columns = [
        {
            title: "Name",
            dataIndex: "name",
            key: 'name',
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
            title: "Address",
            dataIndex: 'address',
            key: 'address'
        },
        {
            title: "First Release Year",
            dataIndex: 'first_release_year',
            key: 'first_release_year',
        },
        {
            title: "No of Albums Released",
            dataIndex: "no_of_albums_released",
            key: "no_of_albums_released",
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
                            title="Delete Artist"
                            description="Are you sure you want to delete this artist?"
                            onConfirm={() => deleteArtist(record?.id)}
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
                title={!!editData ? "Edit Artist" : "Create Artist"}
                open={isModalOpen}
                onCancel={() => {
                    setEditData(null)
                    setIsModalOpen(false)
                }}
                onClose={() => {
                    setEditData(null)
                    setIsModalOpen(false)
                }}
                footer={false}
            >
                <ArtistForm
                    isEdit={!!editData}
                    editData={editData}
                    onCreate={(values) => {
                        setEditData(null)
                        createArtist(values)
                    }}
                    onEdit={(values) => {
                        editArtist(values)
                    }}
                />
            </Modal>
            <Table
                dataSource={artistList}
                columns={columns}
                pagination={{
                    position: "bottomRight",
                    defaultPageSize: 10,
                }}
            />
        </>
    )
}