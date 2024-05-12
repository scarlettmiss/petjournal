import React, {Component} from "react"
import {EnvelopeIcon, MapIcon, PhoneIcon, XMarkIcon} from "@heroicons/react/20/solid"
import TextUtils from "@/Utils/TextUtils"
import {User} from "@/models/user/User"
import Avatar from "@/components/Avatar"

interface UserDialogProps {
    onDismiss: () => void
}

interface UserDialogState {
    title: string
    info: User
    show: boolean
}

export default class UserDialog extends Component<UserDialogProps, UserDialogState> {
    constructor(props: UserDialogProps) {
        super(props)
        this.state = {
            title: "Vet Information",
            info: {},
            show: false,
        }
    }

    public hide = () => {
        this.setState({show: false})
    }

    public show = () => {
        this.setState({show: true})
    }

    public setData = (info: User) => {
        this.setState({info})
    }
    public setTitle = (title: string) => {
        this.setState({title})
    }

    private get vetAddress() {
        const vet = this.state.info
        return [vet.address, vet.city, vet.state, vet.zip, vet.country].filter((w) => TextUtils.isNotEmpty(w)).join(", ")
    }

    render() {
        const user = this.state.info
        return (
            <div id="defaultModal" tabIndex={-1} className={`fixed flex grow ${this.state?.show ? "" : "hidden"} z-50 h-screen w-full`}>
                <div className="relative self-center mx-auto w-full max-w-2xl max-h-full">
                    {/*Modal content*/}
                    <div className="relative rounded-lg shadow bg-gray-700/60 backdrop-blur-lg">
                        {/*Modal header*/}
                        <div className="flex items-start justify-between p-4 border-b rounded-t border-gray-600">
                            <h3 className="text-xl font-semibold text-white text-center">{this.state.title}</h3>
                            <button
                                type="button"
                                className="text-gray-400 bg-transparent rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center hover:bg-gray-600 hover:text-white"
                                onClick={this.props.onDismiss}
                            >
                                <XMarkIcon className={"h-6 w-6"} />
                                <span className="sr-only">Close modal</span>
                            </button>
                        </div>
                        {/*Modal body */}
                        {this.state.info !== undefined && (
                            <div className="px-12 py-4 space-y-4 max-h-96 overflow-y-auto justify-center">
                                <div className={"flex items-center justify-center"}>
                                    <Avatar
                                        avatarTitle={`${user.surname?.slice(0, 1) ?? ""}${user.name?.slice(0, 1) ?? ""}`}
                                        className={"flex p-1 h-[70px] w-[70px]"}
                                    />
                                </div>
                                <h3 className="text-center text-2xl text-indigo-100">
                                    {user.surname} {user.name}
                                </h3>

                                <div className={"flex flex-col gap-2 mt-3"}>
                                    {user.email && (
                                        <a
                                            href={`mailto: ${user.email}`}
                                            target={"_blank"}
                                            className={"flex gap-2 items-center text-indigo-100 hover:text-indigo-300 underline"}
                                        >
                                            <EnvelopeIcon className={"h-5 w-5"} />
                                            {user.email}
                                        </a>
                                    )}
                                    {user.phone && user.phone.length > 3 && (
                                        <a
                                            href={`tel:${user.phone}`}
                                            target={"_parent"}
                                            className={"flex gap-2 items-center text-indigo-100 hover:text-indigo-300 underline"}
                                        >
                                            <PhoneIcon className={"h-5 w-5"} />
                                            {user.phone}
                                        </a>
                                    )}
                                </div>
                                {TextUtils.isNotEmpty(this.vetAddress) && (
                                    <div className="space-y-1">
                                        <div className={"flex gap-2 items-center mt-5 "}>
                                            <MapIcon className={"h-5 w-5"} />
                                            <h2 className="text-lg font-bold tracking-tight text-indigo-100 ">Address</h2>
                                        </div>

                                        <a
                                            href={`http://maps.google.com/?q=${this.vetAddress}`}
                                            target={"_blank"}
                                            className={
                                                "flex text-indigo-100 text-md pe-2 capitalize flex-wrap underline hover:text-indigo-300"
                                            }
                                        >
                                            {this.vetAddress}
                                        </a>
                                    </div>
                                )}
                            </div>
                        )}
                    </div>
                </div>
            </div>
        )
    }
}
