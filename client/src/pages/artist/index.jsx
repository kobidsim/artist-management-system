import { Button, Flex, message, Modal, Popconfirm, Table, Tooltip } from "antd"
import { DeleteFilled, DownloadOutlined, EditFilled, UploadOutlined } from "@ant-design/icons"
import axios from "axios"
import { useEffect, useRef, useState } from "react"
import ArtistForm from "./form"
import { useNavigate } from "react-router-dom"

export default function ArtistPage({isManager}) {
    const [artistList, setArtistList] = useState([])
    const [isModalOpen, setIsModalOpen] = useState(false)
    const [editData, setEditData] = useState(null)
    const [messageApi, contextHolder] = message.useMessage()
    const [csvFile, setCsvFile] = useState(null)
    const navigate = useNavigate()
    const filePickerRef = useRef(null)

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

    const exportCSV = () => {
        console.log("Hello exporting")
        const jwt = localStorage.getItem("jwt")
        axios.get(`http://localhost:8080/artist/csv-export`, {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            const url = window.URL.createObjectURL(new Blob([res.data]))
            const link = document.createElement('a')
            link.href = url
            link.setAttribute('download', 'artists.csv')
            document.body.appendChild(link)
            link.click()
            link.remove()
        }).catch((error) => {
            console.log(error)
        })
    }

    const handleCsvImport = (e) => {
        filePickerRef.current.click()
    }

    const handleFileChange = (e) => {
        const file = e.target.files[0]
        if (!file) return
        
        console.log("filename: ", file.name)

        const formData = new FormData()
        formData.append("csv_file", file)

        const jwt = localStorage.getItem("jwt")
        axios.post("http://localhost:8080/artist/csv-import", formData, {
            headers: {
                Authorization: `Bearer ${jwt}`
            }
        }).then((res) => {
            messageApi.open({
                type: 'success',
                content: res?.data?.message
            })
            listArtists()
        }).catch((error) => {
            messageApi.open({
                type: "error",
                content: error?.response?.data?.message,
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
        (isManager ? {
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
                            onClick={(e) => {
                                e.stopPropagation()
                                setEditData(record)
                                setIsModalOpen(true)
                            }}
                        />
                    </Tooltip>
                    <Tooltip title="Delete">
                        <Popconfirm
                            title="Delete Artist"
                            description="Are you sure you want to delete this artist?"
                            onConfirm={(e) => {
                                e.stopPropagation()
                                deleteArtist(record?.id)
                            }}
                            onCancel={(e) => e.stopPropagation()}
                            onPopupClick={(e) => e.stopPropagation()}
                        >
                            <Button onClick={(e) => e.stopPropagation()} icon={<DeleteFilled />} />
                        </Popconfirm>
                    </Tooltip>
                </div>
            )
        }: {})
    ]

    return(
        <>
            {contextHolder}
            {isManager && <Button type="primary" onClick={() => setIsModalOpen(true)}>Create</Button>}
            {isManager && <><Button onClick={handleCsvImport}><UploadOutlined /> Import CSV</Button><input ref={filePickerRef} onChange={handleFileChange} style={{display: "none"}} type="file"></input></>}
            {isManager && <Button onClick={exportCSV}><DownloadOutlined /> Export CSV</Button>}
            {isManager &&
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
                destroyOnClose
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
            </Modal>}
            <Table
                dataSource={artistList}
                columns={columns}
                onRow={(record) => ({
                    onClick: () => {
                        navigate(`/dashboard/artist/${record?.id}`)
                    }
                })}
                pagination={{
                    position: "bottomRight",
                    defaultPageSize: 10,
                }}
            />
        </>
    )
}