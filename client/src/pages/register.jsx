import { Link, useNavigate } from "react-router-dom";
import { Button, Card, DatePicker, Form, Input, message, notification, Select } from "antd"
import axios from 'axios'

function Register() {
    const [form] = Form.useForm()
    const navigate = useNavigate()
    const [api, contextHolder] = notification.useNotification()
    const [messageApi, messageContextHolder] = message.useMessage()

    const onRegister = (values) => {
        axios.post("http://localhost:8080/register", values)
            .then(() => {
                messageApi.open({
                    type: "success",
                    content: "Registration Complete!"
                })
                setTimeout(() => {
                    navigate("/")
                }, 500);
            })
            .catch((error) => {
                api["error"]({
                    message: "Registration Failed",
                    description: error.response.data
                })
            })
    }

    return (
        <div
            style={{
                height: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center",
            }}
        >
            {contextHolder}
            {messageContextHolder}
            <Card title="Register" style={{
                width: "400px"
            }}>
                <Form
                    form={form}
                    layout="vertical"
                    onFinish={onRegister}
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

                    <Form.Item name={"phone_number"}>
                        <Input placeholder="Phone Number"/>
                    </Form.Item>

                    <Form.Item name={"dob"}>
                        <DatePicker placeholder="Date of Birth"/>
                    </Form.Item>

                    <Form.Item name={"gender"}>
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

                    <Form.Item name={"address"}>
                        <Input placeholder="Address"/>
                    </Form.Item>

                    <Form.Item name={"role"}>
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
                    
                    <Form.Item name={"password"}>
                        <Input placeholder="Password" type="password" />
                    </Form.Item>
                    <Form.Item name={"confirm_password"}>
                        <Input placeholder="Confirm Password" type="password" />
                    </Form.Item>
                    <Button type="primary" htmlType="submit">Register</Button>
                    <Link to={"/"}>Login?</Link>
                </Form>
            </Card>
        </div>
    )
}

export default Register;