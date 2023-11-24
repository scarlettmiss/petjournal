import Button from "@/components/Button"
import React from "react"
import {UserType, UserTypeUtils} from "@/enums/UserType"
import countries from "i18n-iso-countries"
import PhoneInput from "@/components/PhoneInput"
import "react-phone-input-2/lib/material.css"
import TextUtils from "@/Utils/TextUtils"
import TextInput from "@/components/TextInput"
import textInputStyle from "@/components/textInput.module.css"
import styles from "@/components/textInput.module.css"
import {withRouter} from "next/router"
import NavBar from "@/components/NavBar"
import {Auth} from "@/models/user/User"
import {signUpHandler} from "@/pages/api/user"
import ErrorMessage from "@/components/ErrorMessage"
import ErrorDto from "@/models/ErrorDto"
import RegistrationViewModel from "@/viewmodels/user/RegistrationViewModel"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"
import jwtDecode from "jwt-decode"

interface SingUpProps extends WithRouterProps {
}

interface SingUpState {
    vm: RegistrationViewModel
    countriesKeys: string[]
    serverError?: string
}

class Signup extends BaseComponent<SingUpProps, SingUpState> {
    constructor(props: SingUpProps) {
        super(props)
        this.state = {
            vm: new RegistrationViewModel(UserType.OWNER),
            countriesKeys: [],
        }
    }

    componentDidMount() {
        countries.registerLocale(require("i18n-iso-countries/langs/en.json"))
        const names = countries.getNames("en", {select: "official"})
        this.setState({countriesKeys: Object.keys(names)})
    }

    private onUserTypeChange = (userType: string) => {
        const vm = this.state.vm
        vm.userType = UserTypeUtils.getEnum(userType)
        this.setState({vm})
    }

    private onNameChange = (value: string) => {
        const vm = this.state.vm
        vm.name = value
        this.setState({vm})
    }

    private onSurnameChange = (value: string) => {
        const vm = this.state.vm
        vm.surname = value
        this.setState({vm})
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

    private onCountryChange = (event: any) => {
        const vm = this.state.vm
        vm.country = event.currentTarget.value
        this.setState({vm})
    }

    private onCityChange = (value: string) => {
        const vm = this.state.vm
        vm.city = value
        this.setState({vm})
    }

    private onStateChange = (value: string) => {
        const vm = this.state.vm
        vm.state = value
        this.setState({vm})
    }

    private onZipChange = (value: string) => {
        const vm = this.state.vm
        vm.zip = value
        this.setState({vm})
    }

    private onAddressChange = (value: string) => {
        const vm = this.state.vm
        vm.address = value
        this.setState({vm})
    }

    private onPhoneChanged = (phone: string) => {
        const vm = this.state.vm
        vm.phone = phone
        this.setState({vm})
    }

    private checkName = () => {
        this.state.vm.checkName()
        this.forceUpdate()
    }
    private checkSurname = () => {
        this.state.vm.checkSurname()
        this.forceUpdate()
    }
    private checkMail = () => {
        this.state.vm.checkEmail()
        this.forceUpdate()
    }

    private checkPassword = () => {
        this.state.vm.checkPassword()
        this.forceUpdate()
    }

    private navigateToSignIn = () => {
        this.props.router.replace("/auth/login")
    }

    private onLoginSuccess = (token: string) => {
        const decoded: { exp: number } = jwtDecode(token)
        this.cookies.set("token", token, {expires: decoded?.exp > Date.now() ? new Date(decoded?.exp) : undefined})
        this.props.router.replace("/")
    }

    private onFail = (errorMessage?: string) => {
        console.error("Login failed", errorMessage)
        this.setState({serverError: errorMessage})
        this.logout()
    }

    private onSubmit = async () => {
        this.state.vm.validate()
        if (!this.state.vm.isValid) {
            return
        }
        const response = await signUpHandler(this.state.vm)

        const data: Auth = await response.json()

        if (response.ok) {
            // Login successful
            this.onLoginSuccess(data.token!)
        } else {
            // Login failed
            this.onFail((data as ErrorDto).error)
        }
    }

    render() {
        const vm = this.state.vm
        return (
            <main className="flex flex-col h-screen ">
                <NavBar hideAllPages
                        buttons={<Button title={"Sign In"} variant={"primary"} onClick={this.navigateToSignIn}/>}/>
                <div
                    className="flex grow items-center justify-center bg-slate-900 bg-[url('/register-bg.jpg')] dark:bg-[url('/register-bg-dark.jpg')] bg-contain bg-center">
                    <div
                        className="my-4 flex items-center mx-auto bg-white dark:bg-slate-800 max-w-sm md:max-w-xl border rounded-md border-indigo-600">
                        <div className="shadow-sm z-10 px-4 py-4">
                            <h2 className="py-6 text-center text-3xl font-bold tracking-tight text-gray-900 dark:text-indigo-100">
                                Create your account
                            </h2>
                            <form className="space-y-4" method="POST" onSubmit={this.onSubmit}>
                                <div className="w-full inline-flex shadow-sm">
                                    {UserTypeUtils.getAll().map((type) => {
                                        const selected = vm.userType === type
                                        return (
                                            <Button
                                                key={type}
                                                variant={"group"}
                                                selected={selected}
                                                onClick={() => this.onUserTypeChange(type)}
                                                type={"button"}
                                                title={UserTypeUtils.getTitle(type)}
                                                width={"full"}
                                            />
                                        )
                                    })}
                                </div>
                                <div className="w-full inline-flex shadow-sm">
                                    <TextInput
                                        id="name"
                                        name="name"
                                        type="text"
                                        autoComplete="off"
                                        required
                                        width={"full"}
                                        classNames="rounded-s-md"
                                        placeholder="Name"
                                        value={vm.name}
                                        onInput={this.onNameChange}
                                        onBlur={this.checkName}
                                        hasError={vm.hasNameError}
                                        errorMessage={vm.nameError}
                                    />
                                    <TextInput
                                        id="surname"
                                        name="surname"
                                        type="text"
                                        autoComplete="off"
                                        required
                                        width={"full"}
                                        classNames="rounded-e-md"
                                        placeholder="Surname"
                                        value={vm.surname}
                                        onInput={this.onSurnameChange}
                                        onBlur={this.checkSurname}
                                        hasError={vm.hasSurnameError}
                                        errorMessage={vm.surnameError}
                                    />
                                </div>

                                <TextInput
                                    id="email-address"
                                    name="email"
                                    type="email"
                                    autoComplete="off"
                                    required
                                    width={"full"}
                                    classNames="rounded-md shadow-sm"
                                    placeholder="Email address"
                                    value={vm.email}
                                    onInput={this.onEmailChange}
                                    onBlur={this.checkMail}
                                    hasError={vm.hasEmailError}
                                    errorMessage={vm.emailError}
                                />

                                <TextInput
                                    id="password"
                                    name="password"
                                    type="password"
                                    autoComplete="new-password"
                                    required
                                    width={"full"}
                                    classNames="rounded-md shadow-sm"
                                    placeholder="Password"
                                    value={vm.password}
                                    onInput={this.onPasswordChange}
                                    onBlur={this.checkPassword}
                                    hasError={vm.hasPasswordError}
                                    errorMessage={vm.passwordError}
                                />

                                <div className="w-full shadow-sm">
                                    <select
                                        value={vm.country}
                                        id="country"
                                        name="country"
                                        onChange={this.onCountryChange}
                                        className={`${textInputStyle.textInput} ${textInputStyle.textInputFull} rounded-t-md`}
                                    >
                                        <option className={styles.textInputText} value="">
                                            Choose a Country
                                        </option>
                                        {this.state.countriesKeys.map((k) => {
                                            const name = countries.getName(k, "en")
                                            return (
                                                <option className={styles.textInputText} key={name} value={name}>
                                                    {name}
                                                </option>
                                            )
                                        })}
                                    </select>
                                    <TextInput
                                        id="address"
                                        name="address"
                                        type="text"
                                        autoComplete="off"
                                        width="full"
                                        placeholder="Address"
                                        value={vm.address}
                                        onInput={this.onAddressChange}
                                    />
                                    <div className="w-full align-center inline-flex rounded-md shadow-sm ">
                                        <TextInput
                                            id="city"
                                            name="city"
                                            type="text"
                                            autoComplete="off"
                                            width="full"
                                            classNames="rounded-bl-md"
                                            placeholder="City"
                                            value={vm.city}
                                            onInput={this.onCityChange}
                                        />
                                        <TextInput
                                            id="zip"
                                            name="zip"
                                            type="text"
                                            autoComplete="off"
                                            width="full"
                                            placeholder="Zip"
                                            value={vm.zip}
                                            onInput={this.onZipChange}
                                        />
                                        <TextInput
                                            id="state"
                                            name="state"
                                            type="text"
                                            autoComplete="off"
                                            width="full"
                                            classNames="rounded-br-md"
                                            placeholder="State"
                                            value={vm.state}
                                            onInput={this.onStateChange}
                                        />
                                    </div>
                                </div>
                                <div className="w-full shadow-sm">
                                    <PhoneInput
                                        width="full"
                                        autoFormat
                                        countryCodeEditable={false}
                                        preferredCountries={["gr"]}
                                        id="phone"
                                        name="phone"
                                        autoFocus={true}
                                        country={TextUtils.isEmpty(vm.country) ? "gr" : vm.country.toLowerCase()}
                                        onChange={this.onPhoneChanged}
                                        value={vm.phone}
                                    />
                                </div>
                                <ErrorMessage message={this.state.serverError}/>
                                <Button
                                    key={"submit"}
                                    variant={"primary"}
                                    type={"submit"}
                                    title={"Register"}
                                    width={"full"}
                                    onClick={this.onSubmit}
                                />
                            </form>
                        </div>
                    </div>
                </div>
            </main>
        )
    }
}

export default withRouter(Signup)
