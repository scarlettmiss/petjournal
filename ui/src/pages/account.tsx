import Button from "@/components/Button"
import React from "react"
import countries from "i18n-iso-countries"
import PhoneInput from "@/components/PhoneInput"
import "react-phone-input-2/lib/material.css"
import TextUtils from "@/Utils/TextUtils"
import TextInput from "@/components/TextInput"
import textInputStyle from "@/components/textInput.module.css"
import UpdateUserViewModel from "@/viewmodels/user/UpdateUserViewModel"
import {User} from "@/models/user/User"
import {userDeleteHandler, userHandler, userUpdateHandler} from "@/pages/api/user"
import ErrorDto from "@/models/ErrorDto"
import {withRouter} from "next/router"
import dynamic from "next/dynamic"
import ErrorMessage from "@/components/ErrorMessage"
import {TrashIcon} from "@heroicons/react/20/solid"
import {UserTypeUtils} from "@/enums/UserType"
import {noop} from "chart.js/helpers"
import DeleteModal from "@/components/DeleteModal"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"

interface AccountProps extends WithRouterProps {
}

interface AccountState {
    vm: UpdateUserViewModel
    user?: User
    viewOnly: boolean
    token?: string
    serverError?: string
    countriesKeys: string[]
}

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

class Account extends BaseComponent<AccountProps, AccountState> {
    private deleteDialogRef: DeleteModal | null = null

    constructor(props: AccountProps) {
        super(props)
        this.state = {
            vm: new UpdateUserViewModel(),
            viewOnly: true,
            countriesKeys: [],
        }
    }

    private initVm = () => {
        const u = this.state.user
        const vm = this.state.vm
        vm.email = TextUtils.valueOrEmpty(u?.email, "")
        vm.name = TextUtils.valueOrEmpty(u?.name, "")
        vm.surname = TextUtils.valueOrEmpty(u?.surname, "")
        vm.phone = TextUtils.valueOrEmpty(u?.phone, "")
        vm.address = TextUtils.valueOrEmpty(u?.address, "")
        vm.city = TextUtils.valueOrEmpty(u?.city, "")
        vm.zip = TextUtils.valueOrEmpty(u?.zip, "")
        vm.state = TextUtils.valueOrEmpty(u?.state, "")
        vm.country = TextUtils.valueOrEmpty(u?.country, "")

        this.setState({vm})
    }

    private getAccount = async (token?: string) => {
        const resp = await userHandler(token)
        const data: User = await resp.json()
        if (resp.ok) {
            this.setState({user: data}, () => this.initVm())
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw new Error((data as ErrorDto).error)
        }
    }

    private initPage = (token?: string) => {
        this.setState({token})

        this.getAccount(token).catch((e) => {
            console.error("Could not get User", e)
            this.setState({serverError: e.message})
        })

        countries.registerLocale(require("i18n-iso-countries/langs/en.json"))
        const names = countries.getNames("en", {select: "official"})
        this.setState({countriesKeys: Object.keys(names)})
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

    private onSubmit = async () => {
        const vm = this.state.vm
        const token = this.state.token

        vm.validate()
        if (!vm.isValid) {
            return
        }
        const response = await userUpdateHandler(vm, token)

        const data: User = await response.json()

        if (response.ok) {
            this.setState({serverError: "", viewOnly: true, user: data})
        } else if (response.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            // update failed
            console.error("update failed", data)
            this.setState({serverError: (data as ErrorDto).error ?? ""})
        }
    }

    private checkName = () => {
        this.state.vm.checkName()
    }

    private checkSurname = () => {
        this.state.vm.checkSurname()
    }

    private checkMail = () => {
        this.state.vm.checkEmail()
    }

    private editAccount = () => {
        this.getAccount(this.state.token).catch((e) => {
            console.error("Could not get User", e)
        })
        this.setState({viewOnly: false})
    }

    private cancel = () => {
        this.initVm()
        this.setState({serverError: "", viewOnly: true})
    }

    private deletePressed = async () => {
        this.deleteDialogRef?.hide()

        const resp = await userDeleteHandler(this.state.token)
        const data: { message: string; error?: string } = await resp.json()
        if (resp.ok) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            this.setState({serverError: data.error})
        }
    }

    render() {
        if (this.state.viewOnly && TextUtils.isNotEmpty(this.state.serverError)) {
            return <ErrorMessage message={this.state.serverError}/>
        }

        return (
            <ProtectedPage init={this.initPage} key={"account"}
                           className={"bg-[url('/register-bg-dark.jpg')] bg-contain bg-center"}>
                <div
                    className="flex items-center mx-auto my-auto bg-slate-800 max-w-sm lg:max-w-xl border rounded-md border-indigo-600">
                    <div className="shadow-sm z-10 px-4 py-6">
                        <div className="flex w-full justify-center">
                            <h2 className="mb-6 text-center text-3xl font-bold tracking-tight text-indigo-100"></h2>
                        </div>
                        <div className={"flex flex-row items-baseline justify-end"}>
                            <div className="flex justify-center pb-4 grow">
                                <h2 className="text-3xl font-bold tracking-tight text-gray-900 dark:text-indigo-100  text-center">
                                    {this.state.viewOnly ? "Account Information" : "Update Account Information"}
                                </h2>
                            </div>
                            <TrashIcon
                                className={"static h-9 w-9 text-red-500 p-1.5 rounded-full hover:bg-red-300 hover:bg-opacity-20"}
                                onClick={() => this.deleteDialogRef?.show()}
                            />
                        </div>
                        <form className="space-y-4" method="POST" onSubmit={this.onSubmit}>
                            {this.state.user?.userType && (
                                <TextInput
                                    id="role"
                                    name="role"
                                    type="text"
                                    autoComplete="off"
                                    width={"full"}
                                    classNames="rounded-md"
                                    placeholder="Role"
                                    value={UserTypeUtils.getTitle(this.state.user.userType)}
                                    disabled={true}
                                    showLabel={true}
                                    onInput={noop}
                                />
                            )}
                            <div className="w-full gap-2 inline-flex shadow-sm">
                                <TextInput
                                    id="name"
                                    name="name"
                                    type="text"
                                    autoComplete="off"
                                    required
                                    width={"full"}
                                    classNames="rounded-md"
                                    placeholder="Name"
                                    value={this.state.vm.name}
                                    onInput={this.onNameChange}
                                    onBlur={this.checkName}
                                    hasError={this.state.vm.hasNameError}
                                    errorMessage={this.state.vm.nameError}
                                    disabled={this.state.viewOnly}
                                    showLabel={true}
                                />
                                <TextInput
                                    id="surname"
                                    name="surname"
                                    type="text"
                                    autoComplete="off"
                                    required
                                    width={"full"}
                                    classNames="rounded-md"
                                    placeholder="Surname"
                                    value={this.state.vm.surname}
                                    onInput={this.onSurnameChange}
                                    onBlur={this.checkSurname}
                                    hasError={this.state.vm.hasSurnameError}
                                    errorMessage={this.state.vm.surnameError}
                                    disabled={this.state.viewOnly}
                                    showLabel={true}
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
                                value={this.state.vm.email}
                                onInput={this.onEmailChange}
                                onBlur={this.checkMail}
                                hasError={this.state.vm.hasEmailError}
                                errorMessage={this.state.vm.emailError}
                                disabled={this.state.viewOnly}
                                showLabel={true}
                            />

                            <div className="w-full flex flex-col">
                                <label htmlFor={"country"} className="text-slate-300 text-sm mb-2">
                                    country
                                </label>
                                <select
                                    key={this.state.vm.country}
                                    value={this.state.vm.country}
                                    id="country"
                                    name="country"
                                    onChange={this.onCountryChange}
                                    className={`${textInputStyle.textInput} ${textInputStyle.textInputFull} rounded-md`}
                                    disabled={this.state.viewOnly}
                                >
                                    <option value=""> No country Data</option>
                                    {this.state.countriesKeys.map((k) => {
                                        const name = countries.getName(k, "en")
                                        return (
                                            <option key={name} value={name}>
                                                {name}
                                            </option>
                                        )
                                    })}
                                </select>
                            </div>
                            <TextInput
                                id="address"
                                name="address"
                                type="text"
                                autoComplete="off"
                                width="full"
                                classNames="rounded-md"
                                placeholder="Address"
                                value={this.state.vm.address}
                                onInput={this.onAddressChange}
                                disabled={this.state.viewOnly}
                                showLabel={true}
                            />
                            <div className="w-full align-center inline-flex rounded-md shadow-sm gap-2">
                                <TextInput
                                    id="city"
                                    name="city"
                                    type="text"
                                    autoComplete="off"
                                    width="full"
                                    classNames="rounded-md"
                                    placeholder="City"
                                    value={this.state.vm.city}
                                    onInput={this.onCityChange}
                                    disabled={this.state.viewOnly}
                                    showLabel={true}
                                />
                                <TextInput
                                    id="zip"
                                    name="zip"
                                    type="text"
                                    autoComplete="off"
                                    width="full"
                                    placeholder="Zip"
                                    value={this.state.vm.zip}
                                    onInput={this.onZipChange}
                                    disabled={this.state.viewOnly}
                                    classNames="rounded-md"
                                    showLabel={true}
                                />
                                <TextInput
                                    id="state"
                                    name="state"
                                    type="text"
                                    autoComplete="off"
                                    width="full"
                                    placeholder="State"
                                    value={this.state.vm.state}
                                    onInput={this.onStateChange}
                                    disabled={this.state.viewOnly}
                                    classNames="rounded-md"
                                    showLabel={true}
                                />
                            </div>
                            <PhoneInput
                                width="full"
                                countryCodeEditable={true}
                                preferredCountries={["gr"]}
                                id="phone"
                                name="phone"
                                onChange={this.onPhoneChanged}
                                value={this.state.vm.phone}
                                disabled={this.state.viewOnly}
                            />
                            <ErrorMessage message={this.state.serverError}/>
                            <div className="w-full align-center inline-flex rounded-md shadow-sm gap-2">
                                {!this.state.viewOnly ? (
                                    <>
                                        <Button
                                            key={"Cancel"}
                                            variant={"secondary"}
                                            type={"button"}
                                            title={"Cancel"}
                                            width={"full"}
                                            onClick={this.cancel}
                                        />
                                        <Button
                                            key={"submit"}
                                            variant={"primary"}
                                            type={"submit"}
                                            title={"Update"}
                                            width={"full"}
                                            onClick={this.onSubmit}
                                        />
                                    </>
                                ) : (
                                    <>
                                        <Button
                                            key={"Cancel"}
                                            variant={"secondary"}
                                            type={"button"}
                                            title={"back"}
                                            width={"full"}
                                            onClick={() => this.props.router.back()}
                                        />
                                        <Button
                                            key={"Edit"}
                                            variant={"primary"}
                                            type={"button"}
                                            title={"Edit"}
                                            width={"full"}
                                            onClick={this.editAccount}
                                        />
                                    </>
                                )}
                            </div>
                        </form>
                    </div>
                </div>

                <DeleteModal
                    ref={(ref) => this.deleteDialogRef = ref}
                    message={"Are you sure you want to delete your account?"}
                    onDelete={this.deletePressed}
                    onCancel={() => this.deleteDialogRef?.hide()}
                />
            </ProtectedPage>
        )
    }
}

export default withRouter(Account)
