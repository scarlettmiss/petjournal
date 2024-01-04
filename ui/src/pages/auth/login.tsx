import Button from "@/components/Button"
import Link from "next/link"
import React from "react"
import LoginViewModel from "@/viewmodels/user/LoginViewModel"
import {Auth} from "@/models/user/User"
import TextInput from "@/components/TextInput"
import {withRouter} from "next/router"
import NavBar from "@/components/NavBar"
import {loginHandler} from "@/pages/api/user"
import ErrorMessage from "@/components/ErrorMessage"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"
import jwtDecode from "jwt-decode"
import ErrorDto from "@/models/ErrorDto"

interface LoginProps extends WithRouterProps {}

interface LoginState {
    vm: LoginViewModel
    serverError?: string
}

class Login extends BaseComponent<LoginProps, LoginState> {
    constructor(props: LoginProps) {
        super(props)
        this.state = {
            vm: new LoginViewModel(),
        }
    }

    private onEmailChange = (value: string) => {
        const vm = this.state.vm
        vm.email = value

        this.setState({vm})
    }

    private onPasswordChange = (value: string) => {
        const vm = this.state.vm
        vm.password = value
        this.setState({vm})
    }

    private onLoginSuccess = (token: string) => {
        const decoded: {exp: number} = jwtDecode(token)
        this.cookies.set("token", token, {expires: decoded?.exp > Date.now() ? new Date(decoded?.exp) : undefined})
        this.props.router.replace("/")
    }

    private onFail = (errorMessage?: string) => {
        console.error("Login failed", errorMessage)
        this.setState({serverError: errorMessage})
        this.logout()
    }

    private onSubmit = async () => {
        try {
            const resp = await loginHandler(this.state.vm)
            const data: Auth = await resp.json()
            if (resp.ok) {
                this.onLoginSuccess(data.token!)
            } else {
                this.onFail((data as ErrorDto).error)
            }
        } catch (e: any) {
            this.onFail(e.message)
        }
    }

    private navigateToSignUp = () => {
        this.props.router.push("/auth/signup")
    }

    render() {
        return (
            <main className="flex flex-col grow h-screen">
                <NavBar
                    hideAllPages
                    buttons={
                        <Button
                            title={"Sign Up"}
                            variant={"primary"}
                            onClick={this.navigateToSignUp}
                            className={"shadow-sm shadow-indigo-700 hover:shadow-md hover:shadow-indigo-800"}
                        />
                    }
                />
                <div className="flex items-center justify-center h-full bg-[url('/register-bg.jpg')] dark:bg-[url('/register-bg-dark.jpg')] bg-contain bg-center">
                    <div className="isolate container items-center mx-auto bg-white dark:bg-slate-800/30 max-w-sm lg:max-w-md rounded-md px-4 py-12 sm:px-6 lg:px-8 shadow-md backdrop-blur-sm">
                        <div className="w-full max-w-md space-y-8">
                            <div>
                                <h2 className="mt-6 text-center text-3xl font-bold tracking-tight text-gray-900 dark:text-white">
                                    Sign in to your account
                                </h2>
                                <p className="mt-2 text-center text-sm text-gray-600 dark:text-gray-300">
                                    Or{" "}
                                    <Link
                                        href="/auth/signup"
                                        className="font-medium text-indigo-600 dark:text-indigo-400 hover:text-indigo-500 dark:hover:text-indigo-300 "
                                    >
                                        Create a new account
                                    </Link>
                                </p>
                            </div>
                            <form className="mt-8 space-y-6" method="POST" onSubmit={this.onSubmit}>
                                <input type="hidden" name="remember" defaultValue="true" />
                                <div className="shadow-sm">
                                    <TextInput
                                        id="email-address"
                                        name="email"
                                        type="email"
                                        autoComplete="email"
                                        required
                                        width="full"
                                        classNames="rounded-t-md"
                                        placeholder="Email address"
                                        value={this.state.vm.email}
                                        onInput={this.onEmailChange}
                                    />

                                    <TextInput
                                        id="password"
                                        name="password"
                                        type="password"
                                        width="full"
                                        autoComplete="current-password"
                                        required
                                        classNames="rounded-b-md"
                                        placeholder="Password"
                                        value={this.state.vm.password}
                                        onInput={this.onPasswordChange}
                                    />
                                </div>
                                <ErrorMessage message={this.state.serverError} />
                                <Button
                                    variant={"primary"}
                                    type={"submit"}
                                    title={"LogIn"}
                                    width={"full"}
                                    onClick={this.onSubmit}
                                    className={"drop-shadow-sm drop-shadow-indigo-600 hover:shadow-md hover:shadow-indigo-800"}
                                />
                            </form>
                        </div>
                    </div>
                </div>
            </main>
        )
    }
}

export default withRouter(Login)
