import { Button, DatePicker, Form, Input, Select } from "antd";

export default function MusicForm({isEdit, editData, onCreate, onEdit}) {
    const [form] = Form.useForm()
    const handleSumbit = (values) => {
        isEdit ? onEdit(values) : onCreate(values)
    }
    
    return (
        <>
            <Form
                form={form}
                layout="vertical"
                onFinish={handleSumbit}
                initialValues={editData}
            >
                <Form.Item
                    name={"title"}
                    rules={[
                        {
                            required: true,
                            message: 'Title is required.'
                        }
                    ]}
                >
                    <Input placeholder="Title"/>
                </Form.Item>

                <Form.Item name={"album_name"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <Input placeholder="Album Name"/>
                </Form.Item>

                <Form.Item name={"genre"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <Select
                        placeholder="Genre"
                        options={[
                            {
                                value: "rnb",
                                label: "Rnb",
                            },
                            {
                                value: "country",
                                label: "Country",
                            },
                            {
                                value: "classic",
                                label: "Classic",
                            },
                            {
                                value: "rock",
                                label: "Rock",
                            },
                            {
                                value: "jazz",
                                label: "Jazz",
                            },
                        ]}
                    />
                </Form.Item>
                
                <Button type="primary" htmlType="submit">{isEdit ? "Edit" : "Create"}</Button>
            </Form>
        </>
    )
}