import React from "react"
import {Pet} from "@/models/pet/Pet"
import {petsHandler} from "@/pages/api/pet"
import ErrorDto from "@/models/ErrorDto"
import Link from "next/link"
import Button from "@/components/Button"
import ErrorMessage from "@/components/ErrorMessage"
import {withRouter} from "next/router"
import dynamic from "next/dynamic"
import TextUtils from "@/Utils/TextUtils"
import Avatar from "@/components/Avatar"
import {MagnifyingGlassIcon, PlusCircleIcon} from "@heroicons/react/24/outline"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"
import TextInput from "@/components/TextInput"

interface PetsProps extends WithRouterProps {
    className: string
}

interface PetsState {
    pets: Pet[]
    filter: string
    serverError?: string
    token?: string
    loading: boolean
}

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

class Pets extends BaseComponent<PetsProps, PetsState> {
    constructor(props: PetsProps) {
        super(props)
        this.state = {
            pets: [],
            filter: "",
            loading: true,
        }
    }

    private initPage = (token?: string) => {
        this.setState({token})

        this.getPets(token).catch((e) => {
            this.setState({serverError: e.message})
            console.error("Could not get pets", e)
        })
    }

    private getPets = async (token?: string) => {
        const resp = await petsHandler(token)
        const data: Pet[] = await resp.json()
        if (resp.ok) {
            this.setState({pets: data, loading: false})
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((data as ErrorDto).error)
        }
    }

    private navigateToPetCreate = () => {
        this.props.router.push("/pet/create")
    }

    private get filteredPets() {
        if (TextUtils.isEmpty(this.state.filter)) {
            return this.state.pets
        }

        return this.state.pets.filter((pet) => {
            return (
                pet.name?.toUpperCase().includes(this.state.filter) ||
                pet.owner?.name?.toUpperCase().includes(this.state.filter) ||
                pet.owner?.surname?.toUpperCase().includes(this.state.filter) ||
                pet.owner?.email?.toUpperCase().includes(this.state.filter) ||
                pet.owner?.phone?.toUpperCase().includes(this.state.filter)
            )
        })
    }

    private onFilterInputChange = (value: string) => {
        this.setState({filter: value.toUpperCase()})
    }

    private get petCreateButton() {
        return (
            <button
                onClick={this.navigateToPetCreate}
                className="flex flex-col aspect-square py-12 px-8 rounded-md shadow bg-indigo-200 dark:bg-slate-800 hover:shadow-lg justify-center items-center"
            >
                <PlusCircleIcon className={"flex h-[100px] text-indigo-700 dark:text-indigo-400"} />
                <div className={"text-indigo-800 dark:text-indigo-200 text-2xl py-2"}>Create Pet</div>
            </button>
        )
    }

    render() {
        const sortedPets = this.filteredPets.sort((a, b) => new Date(a.createdAt!).getTime() - new Date(b.createdAt!).getTime())
        return (
            <ProtectedPage init={this.initPage} key={"pets"} className={"pt-4 px-7 overflow-y-hidden"}>
                <ErrorMessage message={this.state.serverError} />
                <div className={`flex flex-row justify-between items-center`}>
                    <h2 className={"align-middle text-indigo-800 dark:text-indigo-200 text-2xl"}>Pets</h2>
                    <Button title={"Create Pet"} variant={"primary"} type={"button"} onClick={this.navigateToPetCreate} />
                </div>
                <form className="relative mt-4">
                    <TextInput
                        width={"full"}
                        icon={<MagnifyingGlassIcon className="w-4 h-4 visible text-gray-400" />}
                        type={"search"}
                        id={"pet-filter"}
                        name={"filter"}
                        autoComplete={"off"}
                        placeholder={"Search pet name or owner name or email..."}
                        onInput={this.onFilterInputChange}
                        classNames="p-4 ps-10 rounded-md shadow hover:shadow-lg"
                    />
                </form>

                {this.state.loading ? (
                    <div role="status" className={"flex grow justify-center items-center"}>
                        <svg
                            aria-hidden="true"
                            className="w-20 h-20 animate-spin dark:text-indigo-300 fill-indigo-600"
                            viewBox="0 0 100 101"
                            fill="none"
                            xmlns="http://www.w3.org/2000/svg"
                        >
                            <path
                                d="M100 50.5908C100 78.2051 77.6142 100.591 50 100.591C22.3858 100.591 0 78.2051 0 50.5908C0 22.9766 22.3858 0.59082 50 0.59082C77.6142 0.59082 100 22.9766 100 50.5908ZM9.08144 50.5908C9.08144 73.1895 27.4013 91.5094 50 91.5094C72.5987 91.5094 90.9186 73.1895 90.9186 50.5908C90.9186 27.9921 72.5987 9.67226 50 9.67226C27.4013 9.67226 9.08144 27.9921 9.08144 50.5908Z"
                                fill="currentColor"
                            />
                            <path
                                d="M93.9676 39.0409C96.393 38.4038 97.8624 35.9116 97.0079 33.5539C95.2932 28.8227 92.871 24.3692 89.8167 20.348C85.8452 15.1192 80.8826 10.7238 75.2124 7.41289C69.5422 4.10194 63.2754 1.94025 56.7698 1.05124C51.7666 0.367541 46.6976 0.446843 41.7345 1.27873C39.2613 1.69328 37.813 4.19778 38.4501 6.62326C39.0873 9.04874 41.5694 10.4717 44.0505 10.1071C47.8511 9.54855 51.7191 9.52689 55.5402 10.0491C60.8642 10.7766 65.9928 12.5457 70.6331 15.2552C75.2735 17.9648 79.3347 21.5619 82.5849 25.841C84.9175 28.9121 86.7997 32.2913 88.1811 35.8758C89.083 38.2158 91.5421 39.6781 93.9676 39.0409Z"
                                fill="currentFill"
                            />
                        </svg>
                        <span className="sr-only">Loading...</span>
                    </div>
                ) : sortedPets.length > 0 ? (
                    <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2 py-2 overflow-y-auto mt-2">
                        {sortedPets.map((pet) => {
                            return (
                                <Link
                                    key={pet.id}
                                    href={`/pet/${pet.id}`}
                                    className="flex flex-col !aspect-square rounded-md shadow-md bg-indigo-200 dark:bg-slate-800 hover:shadow-lg isolate"
                                >
                                    <div className={"flex w-full flex-col my-auto mx-auto items-center py-2.5 px-4"}>
                                        <Avatar
                                            avatarTitle={pet.name?.slice(0, 1) ?? "-"}
                                            avatar={pet.avatar}
                                            className={"h-[70px] w-[70px]"}
                                        />
                                        <div
                                            className={
                                                "text-indigo-800 dark:text-indigo-200 text-2xl py-2 truncate max-w-[90%] text-center"
                                            }
                                        >
                                            {pet.name}
                                        </div>
                                        <div
                                            className={"text-indigo-800 dark:text-indigo-200 text-md h-6 truncate max-w-[90%] text-center"}
                                        >
                                            {TextUtils.valueOrEmpty(pet.breedName, "-")}
                                        </div>
                                    </div>
                                    <div className="flex flex-nowrap justify-self-start shadow bg-indigo-500/20 justify-between items-center py-2 px-3 rounded-b-md">
                                        <Avatar
                                            avatarTitle={`${pet.owner?.surname?.slice(0, 1) ?? ""}${pet.owner?.name?.slice(0, 1) ?? ""}`}
                                            textStyle={"text-sm"}
                                            className={"h-[30px] w-[30px]"}
                                        />
                                        <div className={"w-full ps-3 overflow-hidden "}>
                                            <div className={"text-indigo-800 dark:text-indigo-200 text-md pb-1"}>Pet Owner Info</div>
                                            <div className={"text-indigo-800 dark:text-indigo-200 text-sm truncate max-w-full"}>
                                                {pet.owner?.surname} {pet.owner?.name}
                                            </div>
                                            <div className={"text-indigo-800 dark:text-indigo-200 text-sm truncate max-w-full"}>
                                                {TextUtils.valueOrEmpty(pet.owner?.email!, "-")}
                                            </div>
                                        </div>
                                    </div>
                                </Link>
                            )
                        })}
                        {this.petCreateButton}
                    </div>
                ) : (
                    <div className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2 py-2 overflow-y-auto mt-2">
                        {this.petCreateButton}
                    </div>
                )}
            </ProtectedPage>
        )
    }
}

export default withRouter(Pets)
