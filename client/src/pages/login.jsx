import { Button, Card, Form, Input, message, notification } from "antd";
import axios from "axios";
import { Link } from "react-router-dom";

function Login() {
    const [form] = Form.useForm()
    const [api, contextHolder] = notification.useNotification()
    const [messageApi, messageContextHolder] = message.useMessage()

    const onLogin = (values) => {
        axios.post("http://localhost:8080/login", values)
            .then(() => {
                messageApi.open({
                    type: "success",
                    content: "Logged in"
                })
            })
            .catch((error) => {
                api["error"]({
                    message: "Login Failed",
                    description: error.response.data,
                })
            })
    }

    return (
        <div
            style={{
                height: "100vh",
                display: "flex",
                justifyContent: "center",
                alignItems: "center"
            }}
        >
            {contextHolder}
            {messageContextHolder}
            <Card
                title={"Login"}
                style={{
                    width: "400px"
                }}
            >
                <Form
                    form={form}
                    layout="vertical"
                    onFinish={onLogin}
                >
                    <Form.Item
                        name={"email"}
                        rules={[
                            {
                                required: true,
                                message: 'Please enter your email!'
                            },
                            {
                                type: 'email',
                                message: 'This is not a valid email!'
                            }
                        ]}
                    >
                        <Input placeholder="Email" />
                    </Form.Item>
                    <Form.Item
                        name={"password"}
                        rules={[
                            {
                                required: true,
                                message: 'Please enter your password!'
                            }
                        ]}
                    >
                        <Input type="password" placeholder="Password" />
                    </Form.Item>
                    <Button type="primary" htmlType="submit">Login</Button>
                    <Link to={"/register"}>Register?</Link>
                </Form>
            </Card>
        </div>
    )
}

export default Login;