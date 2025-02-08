import { Button, Card, Form, Input, message, notification } from "antd";
import axios from "axios";
import { Link, useNavigate } from "react-router-dom";

function Login() {
    const [form] = Form.useForm()
    const navigate = useNavigate()
    const [api, contextHolder] = notification.useNotification()
    const [messageApi, messageContextHolder] = message.useMessage()

    const onLogin = (values) => {
        axios.post("http://localhost:8080/login", values)
            .then((res) => {
                messageApi.open({
                    type: "success",
                    content: "Logged in"
                })
                localStorage.setItem("role", res.data.data["role"])
                localStorage.setItem("jwt", res.data.data["token"])
                let navigateToRoute = ""
                switch (res.data.data["role"]) {
                    case "super_admin":
                        navigateToRoute = "/dashboard/users"
                        break;
                    
                    case "artist_manager":
                    case "artist":
                        navigateToRoute = "/dashboard/artists"
                        break;

                    default:
                        break;
                }

                setTimeout(() => navigate(navigateToRoute), 500)
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