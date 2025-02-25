import { Button, DatePicker, Form, Input, InputNumber, Select } from "antd";

export default function ArtistForm({isEdit, editData, onCreate, onEdit}) {
    const [form] = Form.useForm()
    const handleSumbit = (values) => {
        isEdit ? onEdit(values) : onCreate(values)
    }
    console.log(editData)
    return (
        <>
            <Form
                form={form}
                layout="vertical"
                onFinish={handleSumbit}
                initialValues={editData}
            >
                <Form.Item
                    name={"name"}
                    rules={[
                        {
                            required: true,
                            message: 'Name is required.'
                        }
                    ]}
                >
                    <Input placeholder="Name"/>
                </Form.Item>

                <Form.Item name={"dob"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <DatePicker placeholder="Date of Birth"/>
                </Form.Item>

                <Form.Item name={"gender"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <Select
                        placeholder="Gender"
                        options={[
                            {
                                value: "m",
                                label: "Male",
                            },
                            {
                                value: "f",
                                label: "Female",
                            },
                            {
                                value: "o",
                                label: "Other",
                            },
                        ]}
                    />
                </Form.Item>

                <Form.Item name={"address"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <Input placeholder="Address"/>
                </Form.Item>

                <Form.Item name={"first_release_year"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <Input placeholder="First release year"/>
                </Form.Item>

                <Form.Item name={"no_of_albums_released"}
                    rules={[
                        {
                            required: true,
                            message: 'Field is required.'
                        }
                    ]}
                >
                    <InputNumber min={0} placeholder="No of Albums released"/>
                </Form.Item>
                
                <Button type="primary" htmlType="submit">{isEdit ? "Edit" : "Create"}</Button>
            </Form>
        </>
    )
}