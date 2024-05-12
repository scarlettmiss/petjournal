import Button from "@/components/Button"
import React, {FormEvent} from "react"
import "react-phone-input-2/lib/material.css"
import TextInput from "@/components/TextInput"
import textInputStyle from "@/components/textInput.module.css"
import styles from "@/components/textInput.module.css"
import {withRouter} from "next/router"
import ErrorMessage from "@/components/ErrorMessage"
import ErrorDto from "@/models/ErrorDto"
import {Pet} from "@/models/pet/Pet"
import {petDeleteHandler, petHandler, petUpdateHandler} from "@/pages/api/pet"
import {PetGenderUtils} from "@/enums/Genders"
import {User} from "@/models/user/User"
import {VetsHandler} from "@/pages/api/user"
import dynamic from "next/dynamic"
import Avatar from "@/components/Avatar"
//@ts-ignore
import {ColorExtractor} from "react-color-extractor"
import PetUpdateViewModel from "@/viewmodels/pet/UpdatePetViewModel"
import {format} from "date-fns"
import {TrashIcon} from "@heroicons/react/20/solid"
import DeleteModal from "@/components/DeleteModal"
import {WithRouterProps} from "next/dist/client/with-router"
import BaseComponent from "@/components/BaseComponent"

const tinycolor = require("tinycolor2")

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

interface EditProps extends WithRouterProps {}

interface EditState {
    pet?: Pet
    vm: PetUpdateViewModel
    vets: User[]
    serverError?: string
    token?: string
    focusColor: boolean
    avatarColors: string[]
}

class Edit extends BaseComponent<EditProps, EditState> {
    private deleteDialogRef: DeleteModal | null = null
    private colorsRefs: Array<HTMLInputElement | null> = []
    private imageInputRef: HTMLInputElement | null = null

    constructor(props: EditProps) {
        super(props)
        this.state = {
            vm: new PetUpdateViewModel(),
            vets: [],
            focusColor: false,
            avatarColors: [],
        }
    }

    componentDidMount() {
        this.getVets().catch((e) => {
            console.error("Could not get Vets", e)
            this.setState({serverError: "Could not get Vets"})
        })
    }

    private initVm = () => {
        const p = this.state.pet
        if (!p) {
            return
        }
        const vm = this.state.vm
        vm.avatar = p.avatar ?? ""
        vm.name = p.name ?? ""
        vm.dateOfBirth = p.dateOfBirth ? format(new Date(p.dateOfBirth), "yyyy-MM-dd") : ""
        vm.gender = p.gender ?? ""
        vm.breedName = p.breedName ?? ""
        vm.colors = p.colors ?? []
        vm.description = p.description ?? ""
        vm.pedigree = p.pedigree ?? ""
        vm.microchip = p.microchip ?? ""
        vm.vetId = p.vet?.id ?? ""
        console.log(p.metas)
        vm.metas = new Map<string, string>(Object.entries(p.metas ?? {}))

        this.setState({vm})
    }

    private getVets = async () => {
        const resp = await VetsHandler()
        const data: User[] = await resp.json()
        if (resp.ok) {
            this.setState({vets: data.filter((v) => v.id !== undefined)})
        } else {
            throw Error((data as ErrorDto).error)
        }
    }

    private getPet = async (id?: string, token?: string) => {
        if (!id) {
            return
        }

        const resp = await petHandler(id, token)
        const data: Pet = await resp.json()
        if (resp.ok) {
            this.setState({pet: data, token}, () => this.initVm())
        } else {
            throw Error((data as ErrorDto).error)
        }
    }

    private initPage = (token?: string) => {
        const petId = this.props.router.query.id as string
        this.getPet(petId, token).catch((e) => this.setState({serverError: e.message}))
    }

    private onNameChange = (value: string) => {
        const vm = this.state.vm
        vm.name = value
        this.setState({vm})
    }

    private onDateOfBirthChange = (value: string) => {
        const vm = this.state.vm
        vm.dateOfBirth = value
        this.setState({vm})
    }

    private onGenderChange = (event: any) => {
        const vm = this.state.vm
        vm.gender = event.currentTarget.value
        this.setState({vm})
    }

    private onBreedNameChange = (value: string) => {
        const vm = this.state.vm
        vm.breedName = value
        this.setState({vm})
    }

    private onDescriptionChange = (event: any) => {
        const vm = this.state.vm
        vm.description = event.currentTarget.value
        this.setState({vm})
    }

    private onPedigreeChange = (value: string) => {
        const vm = this.state.vm
        vm.pedigree = value
        this.setState({vm})
    }

    private onMicrochipChange = (value: string) => {
        const vm = this.state.vm
        vm.microchip = value
        this.setState({vm})
    }

    private onMinWeightChange = (value: string) => {
        const vm = this.state.vm
        vm.weightMin = value
        this.setState({vm})
    }

    private onMaxWeightChange = (value: string) => {
        const vm = this.state.vm
        vm.weightMax = value
        this.setState({vm})
    }

    private onVetIdChange = (event: any) => {
        const vm = this.state.vm
        vm.vetId = event.currentTarget.value
        this.setState({vm})
    }

    private navigateToPetPage = (id?: string) => {
        this.props.router.replace(`/pet/${id}`)
    }

    private onSubmit = async () => {
        this.state.vm.validate()

        console.log(this.state.vm)
        if (!this.state.vm.isValid || !this.state.token) {
            this.forceUpdate()
            return
        }

        const response = await petUpdateHandler(this.state.vm, this.state.pet!.id!, this.state.token)

        const data: Pet = await response.json()

        if (response.ok) {
            this.navigateToPetPage(this.state.pet?.id)
        } else if (response.status === 401) {
            this.logout()
            this.props.router.replace("/auth/login")
        } else {
            this.setState({serverError: (data as ErrorDto).error ?? ""})
        }
    }

    private checkName = () => {
        this.state.vm.checkName()
        this.forceUpdate()
    }

    private checkDateOfBirth = () => {
        this.state.vm.checkDateOfBirth()
        this.forceUpdate()
    }

    private checkGender = () => {
        this.state.vm.checkGender()
        this.forceUpdate()
    }

    private checkBreedName = () => {
        this.state.vm.checkBreedName()
        this.forceUpdate()
    }

    private onColorInput = (event: FormEvent<HTMLInputElement>, index: number) => {
        const colors = this.state.vm.colors
        colors[index] = event.currentTarget.value

        this.setState({vm: this.state.vm})
    }

    private onAddColor = () => {
        const colors = this.state.vm.colors

        colors.push("")

        this.setState({vm: this.state.vm}, () => {
            // workaround wait or the actual DOM to render to show the picker at the correct x,y position
            setTimeout(() => this.colorsRefs[this.state.vm.colors.length - 1]?.click(), 0)
        })
    }

    private onRemoveColor = (index: number) => {
        const vm = this.state.vm

        vm.colors.splice(index, 1)

        this.setState({vm})
    }

    private onImageChange = (event: any) => {
        this.uploadImage(event).catch((e) => console.error(e))
    }

    private convertBase64 = (file: any): Promise<string | ArrayBuffer | null> => {
        return new Promise((resolve, reject) => {
            const fileReader = new FileReader()
            fileReader.readAsDataURL(file)

            fileReader.onload = () => {
                resolve(fileReader.result)
            }

            fileReader.onerror = (error) => {
                reject(error)
            }
        })
    }

    private uploadImage = async (event: any) => {
        const file = event.target.files[0]
        try {
            const base64 = await this.convertBase64(file)
            const vm = this.state.vm
            // @ts-ignore
            vm.avatar = base64
            this.setState({vm})
        } catch (e) {
            console.error(e)
        }
    }

    private deletePressed = async () => {
        this.deleteDialogRef?.hide()

        const id = this.props.router.query.id as string
        if (!id) {
            return
        }
        const resp = await petDeleteHandler(id, this.state.token)
        if (resp.ok) {
            this.props.router.replace("/")
        } else {
            const data: ErrorDto = await resp.json()
            this.setState({serverError: data.error}, () => setTimeout(() => this.props.router.replace("/"), 3000))
        }
    }

    render() {
        const vm = this.state.vm
        return (
            <ProtectedPage
                hideNav={true}
                init={this.initPage}
                className={"bg-[url('/register-bg-dark.jpg')] bg-contain bg-center overflow-y-auto"}
            >
                <div className="flex items-center mx-auto my-auto bg-slate-800/50 backdrop-blur-sm max-w-sm lg:max-w-xl rounded-md shadow-lg">
                    <div className="shadow-sm px-4 py-4">
                        <div className={"flex flex-row items-baseline justify-end"}>
                            <div className="flex justify-center pb-4 grow">
                                <h2 className="text-3xl font-bold tracking-tight text-indigo-100  text-center">
                                    Update Pet Profile
                                </h2>
                            </div>
                            <TrashIcon className={" h-6 w-6 text-red-500"} onClick={() => this.deleteDialogRef?.show()} />
                        </div>
                        <form className="space-y-1 mb-3" method="POST" onSubmit={this.onSubmit}>
                            <div className={"flex py-2"}>
                                <div className={"mx-auto"}>
                                    <Avatar
                                        avatarTitle={"Add Image"}
                                        avatar={this.state.vm.avatar}
                                        onCLick={() => {
                                            this.imageInputRef?.click()
                                        }}
                                        textStyle={"text-lg"}
                                        className={"h-[90px] w-[90px]"}
                                    />
                                </div>
                                <input
                                    ref={(ref) => (this.imageInputRef = ref)}
                                    className={`sr-only`}
                                    multiple={false}
                                    type="file"
                                    id="avatar"
                                    name="avatar"
                                    accept="image/png, image/jpeg"
                                    onChange={this.onImageChange}
                                />
                            </div>
                            <TextInput
                                autoFocus={false}
                                id="name"
                                name="name"
                                type="text"
                                autoComplete="off"
                                required
                                width={"full"}
                                classNames="rounded-md"
                                placeholder="Name *"
                                value={vm.name}
                                onInput={this.onNameChange}
                                onBlur={this.checkName}
                                hasError={vm.hasNameError}
                                errorMessage={vm.nameError}
                                showLabel={true}
                            />
                            <TextInput
                                id="dateOfBirth"
                                name="dateOfBirth"
                                type="date"
                                max={format(new Date(), "yyyy-MM-dd")}
                                autoComplete="off"
                                required
                                width={"full"}
                                classNames="rounded-md"
                                placeholder="Date of Birth *"
                                value={vm.dateOfBirth}
                                onInput={this.onDateOfBirthChange}
                                onBlur={this.checkDateOfBirth}
                                hasError={vm.hasDateOfBirthError}
                                errorMessage={vm.dateOfBirthError}
                                showLabel={true}
                            />

                            <div>
                                <label htmlFor={"gender"} className={styles.label}>
                                    Gender *
                                </label>
                                <select
                                    value={vm.gender}
                                    id="gender"
                                    name="gender"
                                    onChange={this.onGenderChange}
                                    onBlur={this.checkGender}
                                    className={`${textInputStyle.textInput} ${textInputStyle.textInputFull} ${
                                        vm.hasGenderError && styles.errorHighlight
                                    } rounded-md`}
                                >
                                    <option value="">Choose a Gender</option>
                                    {PetGenderUtils.getAll().map((g) => {
                                        const name = PetGenderUtils.getTitle(g)
                                        return (
                                            <option key={name} value={g}>
                                                {name}
                                            </option>
                                        )
                                    })}
                                </select>
                                <span className={styles.errorMessage}>{vm.hasGenderError ? vm.genderError : ""}</span>
                            </div>

                            <TextInput
                                id="breedName"
                                name="breedName"
                                type="text"
                                autoComplete="off"
                                required
                                width={"full"}
                                classNames="rounded-md"
                                placeholder="Breed Name *"
                                value={vm.breedName}
                                onInput={this.onBreedNameChange}
                                onBlur={this.checkBreedName}
                                hasError={vm.hasBreedNameError}
                                errorMessage={vm.breedNameError}
                                showLabel
                            />
                            <div>
                                <label htmlFor={"colors"} className={styles.label}>
                                    Colors
                                </label>
                                <div className={"flex flex-row gap-2 items-center flex-wrap"}>
                                    {vm.colors.map((color, index) => (
                                        <div key={`${color} ${index}`}>
                                            <div
                                                style={{backgroundColor: color}}
                                                onClick={() => this.onRemoveColor(index)}
                                                className={`rounded-full h-[35px] w-[35px] text-center text-xl font-bold ring-1 ring-slate-600`}
                                            >
                                                <span
                                                    style={{
                                                        color: tinycolor
                                                            .mostReadable(color, ["#fff", "#000"], {includeFallbackColors: true})
                                                            .toHexString(),
                                                    }}
                                                    className={`opacity-0 hover:opacity-100`}
                                                >
                                                    {color ? "x" : ""}
                                                </span>
                                            </div>
                                            <input
                                                ref={(ref) => (this.colorsRefs[index] = ref)}
                                                onInput={(e) => this.onColorInput(e, index)}
                                                type="color"
                                                value={color}
                                                className={`sr-only`}
                                            />
                                        </div>
                                    ))}
                                    {this.state.vm.colors.length < 10 && (
                                        <button
                                            type={"button"}
                                            className={`bg-white text-black rounded-full h-[35px] w-[35px]`}
                                            onClick={this.onAddColor}
                                        >
                                            <div className={"text-xl text-center pb-1 font-bold"}>+</div>
                                        </button>
                                    )}
                                </div>
                            </div>

                            <div>
                                <label htmlFor={"description"} className={styles.label}>
                                    Description
                                </label>
                                <textarea
                                    className={`${styles.textInput} rounded-md w-full resize-none`}
                                    id="description"
                                    name="description"
                                    rows={1}
                                    placeholder="Description"
                                    value={vm.description}
                                    onInput={this.onDescriptionChange}
                                />
                            </div>

                            <TextInput
                                id="microchip"
                                name="microchip"
                                type="text"
                                autoComplete="off"
                                width={"full"}
                                classNames="rounded-md"
                                placeholder="Microchip"
                                value={vm.microchip}
                                onInput={this.onMicrochipChange}
                                showLabel
                            />

                            <TextInput
                                id="pedigree"
                                name="pedigree"
                                type="text"
                                autoComplete="off"
                                width={"full"}
                                classNames="rounded-md"
                                placeholder="Pedigree"
                                value={vm.pedigree}
                                onInput={this.onPedigreeChange}
                                showLabel
                            />

                            <div className="w-full align-center inline-flex rounded-md shadow-sm ">
                                <TextInput
                                    id="minWeight"
                                    name="minWeight"
                                    type="text"
                                    autoComplete="off"
                                    width={"full"}
                                    classNames="rounded-s-md "
                                    placeholder="Min Weight"
                                    value={vm.weightMin}
                                    onInput={this.onMinWeightChange}
                                    showLabel
                                />
                                <TextInput
                                    id="maxWeight"
                                    name="maxWeight"
                                    type="text"
                                    autoComplete="off"
                                    width={"full"}
                                    classNames="rounded-e-md"
                                    placeholder="Max Weight"
                                    value={vm.weightMax}
                                    onInput={this.onMaxWeightChange}
                                    showLabel
                                />
                            </div>

                            <div>
                                <label htmlFor={"vetId"} className={styles.label}>
                                    Share with vet
                                </label>
                                <select
                                    value={vm.vetId}
                                    id="vetId"
                                    name="vetId"
                                    onChange={this.onVetIdChange}
                                    className={`${textInputStyle.textInput} ${textInputStyle.textInputFull} ${
                                        vm.hasGenderError && styles.errorHighlight
                                    } rounded-md`}
                                >
                                    <option value="">Choose a Vet</option>
                                    {this.state.vets.map((vet) => {
                                        return (
                                            <option key={vet.id} value={vet.id}>
                                                {vet.name} {vet.surname} ({vet.email})
                                            </option>
                                        )
                                    })}
                                </select>
                            </div>
                            <ErrorMessage message={this.state.serverError} key={"errorMessage"} />
                        </form>
                        <div className="w-full align-center inline-flex rounded-md shadow-sm gap-2">
                            <Button
                                key={"Cancel"}
                                variant={"secondary"}
                                type={"button"}
                                title={"Cancel"}
                                width={"full"}
                                onClick={() => this.props.router.back()}
                            />
                            <Button
                                key={"submit"}
                                variant={"primary"}
                                type={"submit"}
                                title={"Update"}
                                width={"full"}
                                onClick={this.onSubmit}
                            />
                        </div>
                    </div>
                </div>

                <DeleteModal
                    ref={(ref) => (this.deleteDialogRef = ref)}
                    message={"Are you sure you want to delete your pet?"}
                    onDelete={this.deletePressed}
                    onCancel={() => this.deleteDialogRef?.hide()}
                />
            </ProtectedPage>
        )
    }
}

export default withRouter(Edit)
