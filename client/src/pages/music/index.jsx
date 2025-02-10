import { Button, Flex, FloatButton, message, Modal, Popconfirm, Table, Tooltip } from "antd"
import { DeleteFilled, EditFilled, PlusOutlined } from "@ant-design/icons"
import axios from "axios"
import { useEffect, useState } from "react"
import MusicForm from "./form"
import { useLocation, useParams } from "react-router-dom"

export default function MusicPage() {
    const [musicList, setMusicList] = useState([])
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [editData, setEditData] = useState(null)
    const [messageApi, contextHolder] = message.useMessage()
    const { artistID } = useParams()
    const { state } = useLocation()

    const listMusic = (artistID) => {
        const jwt = localStorage.getItem("jwt")
        axios.get("http://localhost:8080/music", {
            headers: {
                Authorization: `Bearer ${jwt}`
            },
            params: {
                "artist_id": artistID,
            }
        }).then((res) => {
            setMusicList(res.data.data)
        }).catch((error) => {
            console.log(error)
        })
    }

    const createMusic = (data) => {
        const jwt = localStorage.getItem("jwt")
        axios.post("http://localhost:8080/music", {...data, composed_by_id: artistID}, {
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
            listMusic(artistID)
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const editMusic = (data) => {
        const jwt = localStorage.getItem("jwt")
        console.log("edit data api: ", editData)
        axios.post(`http://localhost:8080/music/${editData.id}`, data, {
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
            listMusic(artistID)
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
            })
        })
    }

    const deleteMusic = (id) => {
        const jwt = localStorage.getItem("jwt")
        axios.delete(`http://localhost:8080/music/${id}`, {
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
            listMusic(artistID)
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message
            })
        })
    }

    useEffect(() => {
        listMusic(artistID)
    }, [])

    const columns = [
        {
            title: "Title",
            dataIndex: "title",
            key: 'title',
        },
        {
            title: "Album Name",
            dataIndex: "album_name",
            key: 'album_name',
        },
        {
            title: "Genre",
            render: (_, record) => {
                let genre = ""
                switch (record?.genre) {
                    case "rnb":
                        genre = "Rnb"
                        break;
                    
                    case "country":
                        genre = "Country"
                        break;
                    
                    case "classic":
                        genre = "Classic"
                        break;

                    case "rock":
                        genre = "Rock"
                        break;

                    case "jazz":
                        genre = "Jazz"
                        break;
                
                    default:
                        break;
                }

                return genre
            },
            key: 'genre'
        },
        {
            title: "Composed By",
            render: (_, record) => {
                return record?.artist?.name
            },
            key: 'first_release_year',
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
                            title="Delete Music"
                            description="Are you sure you want to delete this music?"
                            onConfirm={() => deleteMusic(record?.id)}
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
            {state && <h4
                style={{
                    padding: "20px 20px",
                    margin: "0"
                }}
            >
                {state?.name}'s music
            </h4>}
            <Tooltip title="Create Music"><FloatButton type="primary" icon={<PlusOutlined/>} onClick={() => setIsModalOpen(true)} /></Tooltip>
            <Modal
                title={!!editData ? "Edit Music" : "Create Music"}
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
                destroyOnClose
            >
                <MusicForm
                    isEdit={!!editData}
                    editData={editData}
                    onCreate={(values) => {
                        createMusic(values)
                    }}
                    onEdit={(values) => {
                        editMusic(values)
                    }}
                />
            </Modal>
            <Table
                dataSource={musicList}
                columns={columns}
                pagination={{
                    position: "bottomRight",
                    defaultPageSize: 10,
                }}
            />
        </>
    )
}