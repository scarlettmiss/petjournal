import React from "react"
import {Pet} from "@/models/pet/Pet"
import {petsHandler} from "@/pages/api/pet"
import ErrorDto from "@/models/ErrorDto"
import Link from "next/link"
import Button from "@/components/Button"
import ErrorMessage from "@/components/ErrorMessage"
import {withRouter} from "next/router"
import dynamic from "next/dynamic"
import {PetGenderUtils} from "@/enums/Genders"
import TextUtils from "@/Utils/TextUtils"
import Avatar from "@/components/Avatar"
import {PlusCircleIcon} from "@heroicons/react/24/outline"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"

interface PetsProps extends WithRouterProps {
    className: string
}

interface PetsState {
    pets: Pet[]
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

    render() {
        const sortedPets = this.state.pets.sort((a, b) => new Date(a.createdAt!).getTime() - new Date(b.createdAt!).getTime())
        return (
            <ProtectedPage init={this.initPage} key={"pets"} className={"pt-4 px-7 overflow-y-hidden"}>
                <ErrorMessage message={this.state.serverError}/>
                <div className={`flex flex-row justify-between items-center `}>
                    <h2 className={"align-middle text-indigo-200 text-2xl"}>Pets</h2>
                    <Button title={"Create Pet"} variant={"primary"} type={"button"}
                            onClick={this.navigateToPetCreate}/>
                </div>
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
                    <div
                        className="grid grid-cols-1 md:grid-cols-3 lg:grid-cols-4 xl:grid-cols-5 gap-2 py-2 overflow-y-auto mt-2">
                        {sortedPets.map((pet) => {
                            return (
                                <Link key={pet.id} href={`/pet/${pet.id}`}>
                                    <div
                                        className="flex flex-col aspect-square p-8 border rounded-md shadow bg-slate-800 border-indigo-500 hover:bg-slate-600 justify-center items-center">
                                        <Avatar avatarTitle={pet.name?.slice(0, 1) ?? "-"}
                                                avatar={pet.avatar} className={"h-[100px] w-[100px]"}/>
                                        <div className={"text-indigo-200 text-2xl py-2"}>{pet.name}</div>
                                        <div>
                                            <div className={"text-indigo-200 text-md"}>
                                                Breed: {TextUtils.valueOrEmpty(pet.breedName, "-")}
                                            </div>
                                            <div className={"text-indigo-200 text-md"}>
                                                Gender: {TextUtils.valueOrEmpty(PetGenderUtils.getTitle(pet.gender!), "-")}
                                            </div>
                                        </div>
                                    </div>
                                </Link>
                            )
                        })}
                    </div>
                ) : (
                    <div className={"flex flex-col grow justify-center items-center"}
                         onClick={this.navigateToPetCreate}>
                        <PlusCircleIcon className={"flex h-80 text-indigo-400"}/>
                        <h3 className={"text-cyan-200 text-5xl font-semibold"}>Create your first pet</h3>
                    </div>
                )}
            </ProtectedPage>
        )
    }
}

export default withRouter(Pets)
