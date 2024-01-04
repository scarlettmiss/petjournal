import Button from "@/components/Button"
import React, {FormEvent} from "react"
import "react-phone-input-2/lib/material.css"
import TextInput from "@/components/TextInput"
import styles from "@/components/textInput.module.css"
import {withRouter} from "next/router"
import ErrorMessage from "@/components/ErrorMessage"
import ErrorDto from "@/models/ErrorDto"
import PetCreationViewModel from "@/viewmodels/pet/PetCreationViewModel"
import {Pet} from "@/models/pet/Pet"
import {petCreationHandler} from "@/pages/api/pet"
import {PetGenderUtils} from "@/enums/Genders"
import {User} from "@/models/user/User"
import {VetsHandler} from "@/pages/api/user"
import dynamic from "next/dynamic"
import Avatar from "@/components/Avatar"
//@ts-ignore
import {ColorExtractor} from "react-color-extractor"
import {format} from "date-fns"
import BaseComponent from "@/components/BaseComponent"
import {WithRouterProps} from "next/dist/client/with-router"

const tinycolor = require("tinycolor2")

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

interface CreateProps extends WithRouterProps {}

interface CreateState {
    vm: PetCreationViewModel
    vets: User[]
    serverError?: string
    token?: string
    focusColor: boolean
    avatarColors: string[]
}

class Create extends BaseComponent<CreateProps, CreateState> {
    private colorsRefs: Array<HTMLInputElement | null> = []
    private imageInputRef: HTMLInputElement | null = null

    constructor(props: CreateProps) {
        super(props)
        this.state = {
            vm: new PetCreationViewModel(),
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

    private getVets = async () => {
        const resp = await VetsHandler()
        const data: User[] = await resp.json()
        if (resp.ok) {
            this.setState({vets: data.filter((v) => v.id !== undefined)})
        } else {
            throw Error((data as ErrorDto).error)
        }
    }

    private initPage = (token?: string) => {
        this.setState({token})
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
        vm.weightMin = Number(value)
        this.setState({vm})
    }

    private onMaxWeightChange = (value: string) => {
        const vm = this.state.vm
        vm.weightMax = Number(value)
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
        if (!this.state.vm.isValid || !this.state.token) {
            this.forceUpdate()
            return
        }

        const response = await petCreationHandler(this.state.vm, this.state.token)

        const data: Pet = await response.json()

        if (response.ok) {
            this.navigateToPetPage(data.id)
        } else if (response.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
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

    private getAvatarColors = (colors: string[]) => {
        const vm = this.state.vm

        vm.colors = vm.colors.filter((c) => !this.state.avatarColors.includes(c))
        vm.colors.push(...colors)
        this.setState({vm, avatarColors: colors})
    }

    render() {
        const vm = this.state.vm
        return (
            <ProtectedPage hideNav={true} init={this.initPage} className={"bg-[url('/register-bg.jpg')] dark:bg-[url('/register-bg-dark.jpg')] bg-contain bg-center"}>
                <div className="container items-center mx-auto my-auto bg-white/40 dark:bg-slate-800/30 max-w-sm lg:max-w-xl rounded-md backdrop-blur-sm shadow-lg">
                    <div className="shadow-sm px-4 py-4">
                        <h2 className=" pb-4 text-center text-3xl font-bold tracking-tight text-gray-900 dark:text-indigo-100">
                            Create a Pet profile
                        </h2>
                        <form className="space-y-1 mb-3" method="POST" onSubmit={this.onSubmit}>
                            <div className={"flex py-2"}>
                                {this.state.vm.avatar && <ColorExtractor src={this.state.vm.avatar} getColors={this.getAvatarColors} />}
                                <div className={"mx-auto "}>
                                    <Avatar
                                        avatarTitle={"Add Image"}
                                        avatar={this.state.vm.avatar}
                                        onCLick={() => {
                                            this.imageInputRef?.click()
                                        }}
                                        textStyle={"text-lg h-[90px] w-[90px] hover:shadow-xl"}
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
                                    className={`${styles.textInput} ${styles.textInputFull} ${
                                        vm.hasGenderError && styles.errorHighlight
                                    } rounded-md`}
                                    required
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
                                        <div key={index}>
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
                                    value={vm.weightMin ? vm.weightMin.toString() : undefined}
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
                                    value={vm.weightMax ? vm.weightMax.toString() : undefined}
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
                                    className={`${styles.textInput} ${styles.textInputFull} ${
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
                                title={"Create"}
                                width={"full"}
                                onClick={this.onSubmit}
                            />
                        </div>
                    </div>
                </div>
            </ProtectedPage>
        )
    }
}

export default withRouter(Create)
