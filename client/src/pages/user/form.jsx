import { Button, DatePicker, Form, Input, Select } from "antd";
import { useEffect } from "react";

export default function UserForm({isEdit, editData, onCreate, onEdit}) {
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
                    name={"first_name"}
                    rules={[
                        {
                            required: true,
                            message: 'First name is required.'
                        }
                    ]}
                >
                    <Input placeholder="First Name"/>
                </Form.Item>

                <Form.Item
                    name={"last_name"}
                    rules={[
                        {
                            required: true,
                            message: 'Last name is required.'
                        }
                    ]}
                >
                    <Input placeholder="Last Name"/>
                </Form.Item>

                <Form.Item
                    name={"email"}
                    rules={[
                        {
                            required: true,
                            message: 'Email is required'
                        },
                        {
                            type: 'email',
                            message: 'This is not a valid email'
                        }
                    ]}
                >
                    <Input placeholder="Email"/>
                </Form.Item>

                <Form.Item name={"phone_number"}
                    rules={[
                        {
                            required: true,
                            message: 'Role is required.'
                        }
                    ]}
                >
                    <Input placeholder="Phone Number"/>
                </Form.Item>

                <Form.Item name={"dob"}
                    rules={[
                        {
                            required: true,
                            message: 'Role is required.'
                        }
                    ]}
                >
                    <DatePicker placeholder="Date of Birth"/>
                </Form.Item>

                <Form.Item name={"gender"}
                    rules={[
                        {
                            required: true,
                            message: 'Role is required.'
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
                            message: 'Role is required.'
                        }
                    ]}
                >
                    <Input placeholder="Address"/>
                </Form.Item>

                <Form.Item name={"role"}
                    rules={[
                        {
                            required: true,
                            message: 'Role is required.'
                        }
                    ]}                
                >
                    <Select
                        placeholder="Role"
                        options={[
                            {
                                label: "Admin",
                                value: "super_admin",
                            },
                            {
                                label: "Artist Manager",
                                value: "artist_manager",
                            },
                            {
                                label: "Artist",
                                value: "artist",
                            },
                        ]}
                    />
                </Form.Item>
                
                <Form.Item name={"password"}
                    rules={[
                        {
                            required: true,
                            message: 'Password is required.'
                        }
                    ]}
                >
                    <Input placeholder="Password" type="password" />
                </Form.Item>
                
                <Button type="primary" htmlType="submit">{isEdit ? "Edit" : "Create"}</Button>
            </Form>
        </>
    )
}