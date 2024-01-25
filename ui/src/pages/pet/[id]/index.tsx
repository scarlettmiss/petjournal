import {withRouter} from "next/router"
import React from "react"
import dynamic from "next/dynamic"
import ErrorDto from "@/models/ErrorDto"
import {petHandler} from "@/pages/api/pet"
import {Pet, PetMaxWeight, PetMinWeight} from "@/models/pet/Pet"
import {PetGenderUtils} from "@/enums/Genders"
import Avatar from "@/components/Avatar"
import {
    CheckBadgeIcon,
    ChevronDownIcon,
    ChevronUpIcon,
    DocumentDuplicateIcon,
    PencilIcon,
    PlusIcon,
    XMarkIcon,
} from "@heroicons/react/20/solid"
import {ArrowTopRightOnSquareIcon, CheckBadgeIcon as CheckBadgeIconOutline, PlusCircleIcon} from "@heroicons/react/24/outline"
import ErrorMessage from "@/components/ErrorMessage"
import TextUtils from "@/Utils/TextUtils"
import {differenceInCalendarMonths, differenceInCalendarWeeks, differenceInCalendarYears, format, isBefore} from "date-fns"
import {Area, CartesianGrid, ComposedChart, Legend, Line, ResponsiveContainer, Tooltip, XAxis, YAxis} from "recharts"
import UserDialog from "@/components/UserDialog"
import {Record} from "@/models/record/Record"
import RecordDialog from "@/components/RecordDialog"
import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import BaseComponent from "@/components/BaseComponent"
import {WithRouterProps} from "next/dist/client/with-router"
import {
    petRecordsHandler,
    recordCreationHandler,
    recordDeletionHandler,
    recordsCreationHandler,
    recordUpdateHandler,
} from "@/pages/api/record"
import DeleteModal from "@/components/DeleteModal"
import {RecordType, RecordTypeUtils} from "@/enums/RecordType"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"
import LineRecordDialog from "@/components/LineRecordDialog"
import {noop} from "chart.js/helpers"
import {RecordViewType, RecordViewTypeUtils} from "@/enums/RecordViewType"
import CalendarEvent from "@/models/CalendarEvent"
import CalendarComponent from "@/components/CalendarComponent"

interface PetProps extends WithRouterProps {}

class WeightLineData {
    date: string
    weight: number
    weightLimit: number[]

    constructor(date: string, weight: number, weightLimit: number[]) {
        this.date = date
        this.weight = weight
        this.weightLimit = weightLimit
    }
}

interface RecordColor {
    main: string
    tooltip: string
    dot: string
    backgroundColor: string
}

interface RecordTheme {
    border: string
    icon: string
    background: string
}

interface PetState {
    token?: string
    pet?: Pet
    serverError?: string
    records: Array<Record>
    lineData: Array<WeightLineData>
    events: Array<CalendarEvent>
    selectedRecordId?: string
    expandedTypes: Map<string, boolean>
    selectedLineModel: RecordType.TEMPERATURE | RecordType.WEIGHT
    recordViewType: RecordViewType
    loading: boolean
}

const ProtectedPage = dynamic(() => import("@/components/ProtectedPage"), {
    ssr: false,
})

class PetPage extends BaseComponent<PetProps, PetState> {
    private weightEntriesDialogRef: LineRecordDialog | null = null
    private userDialogRef: UserDialog | null = null
    private recordDialogRef: RecordDialog | null = null
    private deleteDialogRef: DeleteModal | null = null

    constructor(props: PetProps) {
        super(props)
        const expanded = new Map()
        RecordTypeUtils.getAll().forEach((t) => {
            expanded.set(t, false)
        })
        expanded.set(RecordType.REMINDER, true)
        expanded.set(RecordType.OVER_DUE, true)
        this.state = {
            records: [],
            lineData: [],
            events: [],
            expandedTypes: expanded,
            selectedLineModel: RecordType.WEIGHT,
            recordViewType: RecordViewType.RECORDS,
            loading: true,
        }
    }

    private getPet = async (id?: string, token?: string) => {
        if (!id) {
            return
        }

        const resp = await petHandler(id, token)
        const response: Pet = await resp.json()
        if (resp.ok) {
            this.setState({pet: response, token})
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private getRecords = async (id?: string, token?: string) => {
        if (!id) {
            return
        }

        const resp = await petRecordsHandler(id, token)
        const response: Record[] = await resp.json()
        if (resp.ok) {
            this.setState({records: response, loading: false}, () => {
                this.updateLineDatasets()
                this.updateEvents()
            })
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private createRecord = async (vm: RecordCreationViewModel) => {
        const id = this.props.router.query.id as string
        if (!id) {
            throw Error("Pet Id was not defined")
        }

        const resp = await recordCreationHandler(vm, id, this.state.token)
        const response: Record = await resp.json()
        if (resp.ok) {
            this.state.records.push(response)
            this.setState({records: this.state.records}, () => {
                this.updateLineDatasets()
                this.updateEvents()
            })
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private createRecords = async (vm: RecordCreationViewModel) => {
        const id = this.props.router.query.id as string
        if (!id) {
            throw Error("Pet Id was not defined")
        }

        const resp = await recordsCreationHandler(vm, id, this.state.token)
        const response: Record[] = await resp.json()
        if (resp.ok) {
            this.state.records.push(...response)
            this.setState({records: this.state.records}, () => {
                this.updateLineDatasets()
                this.updateEvents()
            })
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private updateRecord = async (vm: UpdateRecordViewModel, id?: string) => {
        if (!id) {
            throw Error("no record to update")
        }

        const petId = this.props.router.query.id as string
        if (!petId) {
            throw Error("Pet Id was not defined")
        }

        const resp = await recordUpdateHandler(vm, id, petId, this.state.token)
        const response: Record = await resp.json()
        if (resp.ok) {
            const records = this.state.records
            const index = records.findIndex((it) => it.id === id)
            records[index] = response
            this.setState({records}, () => {
                this.updateLineDatasets()
                this.updateEvents()
            })
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private deleteRecord = async () => {
        this.deleteDialogRef?.hide()

        const id = this.state.selectedRecordId
        if (!id) {
            throw Error("no record to delete")
        }

        const petId = this.props.router.query.id as string
        if (!petId) {
            throw Error("Pet Id was not defined")
        }

        const resp = await recordDeletionHandler(id, petId, this.state.token)
        const response: Record = await resp.json()
        if (resp.ok) {
            const newRecords = this.state.records.filter((r) => r.id !== id)
            this.setState({records: newRecords}, () => {
                this.updateLineDatasets()
                this.updateEvents()
                this.weightEntriesDialogRef?.setData(newRecords)
            })
        } else if (resp.status === 401) {
            this.logout(() => this.props.router.replace("/auth/login"))
        } else {
            throw Error((response as ErrorDto).error)
        }
    }

    private getLineEntriesFrom(records: Record[]): Record[] {
        return records.filter((r) => r.recordType === this.state.selectedLineModel).sort((a, b) => a.date! - b.date!)
    }

    private initPage = (token?: string) => {
        const petId = this.props.router.query.id as string
        this.getPet(petId, token).catch((e) => this.setState({serverError: e.message}))
        this.getRecords(petId, token).catch((e) => this.setState({serverError: e.message}))
    }

    private navigateToEdit = () => {
        this.props.router.push(`/pet/[id]/edit`, `/pet/${this.state.pet?.id}/edit`)
    }

    private updateLineDatasets = () => {
        const data = this.getLineEntriesFrom(this.state.records)
        const metas = new Map<string, string>(Object.entries(this.state.pet?.metas ?? {}))
        const lineData: WeightLineData[] = data.map(
            (entry) =>
                new WeightLineData(
                    format(new Date(entry.date!), "dd/MM/yyyy"),
                    Number(entry.result!),
                    entry.recordType === RecordType.WEIGHT
                        ? [Number(metas?.get(PetMinWeight)) ?? 0, Number(metas?.get(PetMaxWeight)) ?? 0]
                        : []
                )
        )
        this.setState({lineData: lineData})
    }

    private getAge(pet: Pet) {
        const years = differenceInCalendarYears(new Date(), new Date(pet.dateOfBirth!))
        if (years > 0) {
            return `${years} y`
        }

        const months = differenceInCalendarMonths(new Date(), new Date(pet.dateOfBirth!))
        if (months > 0) {
            return `${months} m`
        }

        const weeks = differenceInCalendarWeeks(new Date(), new Date(pet.dateOfBirth!))
        if (weeks < 2) {
            return "Newborn"
        }

        return `${weeks} w`
    }

    private get petCard() {
        const pet = this.state.pet!
        return (
            <div key={pet.id} className={"flex flex-col grow h-full xl:h-auto w-full"}>
                <div className={"flex flex-col bg-indigo-300 dark:bg-slate-800 py-4 px-6 rounded-md shadow-2xl relative"}>
                    <PencilIcon
                        className={"absolute h-8 w-8 self-end text-indigo-500 dark:text-indigo-300 z-1 hover:bg-gray-600 p-1 rounded-md"}
                        onClick={this.navigateToEdit}
                    />
                    <div className={"flex flex-col xl:flex-row gap-2 my-4 justify-evenly align-middle"}>
                        <Avatar
                            avatarTitle={pet.name?.slice(0, 1) ?? "-"}
                            avatar={pet.avatar}
                            className={"self-center h-[100px] w-[100px] xl:h-[150px] xl:w-[150px]"}
                        />
                        <div className={"gap-2 mt-2 xl:max-w-[60%]"}>
                            <h3 className="text-center text-xl xl:text-2xl tracking-wider text-indigo-800 dark:text-indigo-100 truncate ">
                                {pet.name}
                            </h3>
                            <div className="flex flex-col ">
                                <div className={"flex flex-row items-center"}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg pe-2 capitalize whitespace-nowrap">
                                        breed :
                                    </h3>
                                    <p className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg truncate">{pet.breedName}</p>
                                </div>
                                <div className={"flex flex-row items-center"}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg pe-2 capitalize whitespace-nowrap">
                                        Age :
                                    </h3>
                                    <p className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg">{this.getAge(pet)}</p>
                                </div>
                                <div className={"flex flex-row items-center"}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg pe-2 capitalize whitespace-nowrap">
                                        gender :
                                    </h3>
                                    <p className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg">
                                        {PetGenderUtils.getTitle(pet.gender)}
                                    </p>
                                </div>
                            </div>
                            {pet.colors && pet.colors.length > 0 && (
                                <div className={"flex flex-row items-start "}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg capitalize justify-start whitespace-nowrap">
                                        Colors :
                                    </h3>
                                    <div className={"flex flex-row gap-x-1.5 gap-y-1 ms-3 flex-wrap py-1.5"}>
                                        {pet.colors.map((color, index) => (
                                            <div
                                                key={`color_${index}`}
                                                style={{backgroundColor: color}}
                                                className={`rounded-full aspect-square h-[15px] w-[15px] lg:h-[20px] lg:w-[20px] text-center text-xl font-bold ring-1 ring-slate-600`}
                                            ></div>
                                        ))}
                                    </div>
                                </div>
                            )}
                            {pet.pedigree && (
                                <div className={"flex flex-row items-center"}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg pe-2 capitalize whitespace-nowrap">
                                        pedigree :
                                    </h3>
                                    <p className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg truncate">{pet.pedigree}</p>
                                </div>
                            )}
                            {pet.microchip && (
                                <div className={"flex flex-row items-center"}>
                                    <h3 className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg pe-2 capitalize whitespace-nowrap">
                                        microchip :
                                    </h3>
                                    <p className="text-indigo-700 dark:text-indigo-200 text-sm xl:text-lg max-w-xl truncate">
                                        {pet.microchip}
                                    </p>
                                </div>
                            )}
                        </div>
                    </div>
                    <div className={"self-center gap-2"}>
                        {pet.description && (
                            <p className="text-indigo-800 dark:text-indigo-100 text-sm xl:text-lg pe-2 capitalize flex-wrap">
                                {pet.description}
                            </p>
                        )}
                    </div>

                    <div className={"flex flex-col md:flex-row gap-1 mt-2"}>
                        {pet.vet !== undefined && TextUtils.isNotEmpty(pet.vet?.id) && (
                            <div
                                className={
                                    "flex grow justify-center items-center gap-2 p-1 bg-indigo-500/60 dark:bg-indigo-600/30 rounded-md text-indigo-100 shadow hover:shadow-lg text-sm xl:text-lg !whitespace-nowrap text-ellipsis"
                                }
                                onClick={() => {
                                    this.userDialogRef?.setData(this.state.pet!.vet!)
                                    this.userDialogRef?.setTitle("Vet Information")
                                    this.userDialogRef?.show()
                                }}
                            >
                                Vet Info
                                <ArrowTopRightOnSquareIcon className={"h-4 w-4 z-1"} />
                            </div>
                        )}
                        {pet.owner !== undefined && TextUtils.isNotEmpty(pet.owner?.name) && (
                            <div
                                className={
                                    "flex grow justify-center items-center gap-2 p-1 bg-indigo-500/60 dark:bg-indigo-600/30 rounded-md text-indigo-100 shadow hover:shadow-lg text-sm xl:text-lg"
                                }
                                onClick={() => {
                                    this.userDialogRef?.setData(this.state.pet!.owner!)
                                    this.userDialogRef?.setTitle("Owner Information")
                                    this.userDialogRef?.show()
                                }}
                            >
                                Owner Info
                                <ArrowTopRightOnSquareIcon className={"h-4 w-4 z-1"} />
                            </div>
                        )}
                    </div>
                </div>
            </div>
        )
    }

    private setCreateRecord = () => {
        this.recordDialogRef?.setCreate()
        this.recordDialogRef?.show()
    }

    private setCreateWeightEntry = () => {
        this.weightEntriesDialogRef?.setData(this.state.records)
        this.weightEntriesDialogRef?.setRecordType(this.state.selectedLineModel)
        this.weightEntriesDialogRef?.setCreate()
        this.weightEntriesDialogRef?.show()
    }

    private setEditWeightEntry = () => {
        this.weightEntriesDialogRef?.setData(this.state.records)
        this.weightEntriesDialogRef?.setRecordType(this.state.selectedLineModel)
        this.weightEntriesDialogRef?.setList()
        this.weightEntriesDialogRef?.show()
    }

    private getColors(type: RecordType): RecordColor {
        switch (type) {
            case RecordType.TEMPERATURE:
                return {main: "#22d3ee", tooltip: "#0891b2", dot: "#06b6d4", backgroundColor: "#1E293B"}
            default:
                return {main: "#818CF8", tooltip: "#4F46E5", dot: "#6366F1", backgroundColor: "#1E293B"}
        }
    }

    private get weightLine() {
        const supportsEdit = this.state.lineData.length > 0
        const colors: RecordColor = this.getColors(this.state.selectedLineModel)
        return (
            <div className={"flex flex-col grow h-full w-full"}>
                <div
                    className={`${
                        this.state.loading ? "animate-pulse" : ""
                    } flex flex-col bg-indigo-300 dark:bg-slate-800 py-2 px-4 rounded-md shadow-2xl grow`}
                >
                    <div className={`flex flex-row items-center justify-between mb-6`}>
                        <div className={"flex flex-row gap-4"}>
                            {[RecordType.WEIGHT, RecordType.TEMPERATURE].map((it) => (
                                <button
                                    key={it}
                                    onClick={() => {
                                        this.weightEntriesDialogRef?.setRecordType(it)
                                        this.setState(
                                            {
                                                //@ts-ignore
                                                selectedLineModel: it!,
                                            },
                                            this.updateLineDatasets
                                        )
                                    }}
                                >
                                    <h2
                                        className={`capitalize justify-self-center text-lg xl:text-3xl font-bold tracking-tight ${
                                            this.state.selectedLineModel === it
                                                ? `text-indigo-500 dark:text-indigo-300`
                                                : `text-indigo-050 dark:text-indigo-100`
                                        }`}
                                    >
                                        {RecordTypeUtils.getTitle(it)}
                                    </h2>
                                </button>
                            ))}
                        </div>
                        <div className={"flex flex-row items-center lg:align-baseline align-middle"}>
                            <PlusIcon
                                className={
                                    "flex h-8 w-8 text-indigo-500 dark:text-indigo-300 lg:mt-4 lg:ml-4 z-1 hover:bg-indigo-500/40 hover:dark:bg-gray-600 lg:p-1 rounded-md"
                                }
                                onClick={this.setCreateWeightEntry}
                            />
                            <PencilIcon
                                className={`flex h-7 lg:h-8 w-8 lg:mt-4 lg:ml-4 z-1 hover:bg-indigo-500/40 hover:dark:bg-gray-600 lg:p-1 rounded-md ${
                                    supportsEdit ? `text-indigo-500 dark:text-indigo-300` : `text-gray-400 dark:text-slate-300`
                                }`}
                                onClick={supportsEdit ? this.setEditWeightEntry : noop}
                            />
                        </div>
                    </div>
                    {this.state.loading ? (
                        <></>
                    ) : this.state.lineData.length > 0 ? (
                        <ResponsiveContainer className={`flex grow items-center justify-center`} width={"95%"} height={230}>
                            <ComposedChart data={this.state.lineData}>
                                <XAxis dataKey="date" stroke={colors.main} />
                                <YAxis stroke={colors.main} />
                                <CartesianGrid strokeDasharray="3 3" stroke={colors.main + "4a"} />
                                <Area
                                    dataKey="weightLimit"
                                    stroke={colors.main}
                                    strokeWidth={0.3}
                                    fill={colors.main + "4a"}
                                    activeDot={false}
                                />
                                <Line
                                    type="monotone"
                                    dataKey="weight"
                                    stroke={colors.dot}
                                    strokeWidth={2.5}
                                    dot={{stroke: colors.dot, strokeWidth: 3}}
                                />
                                <Tooltip filterNull contentStyle={{backgroundColor: colors.backgroundColor, borderColor: colors.tooltip}} />
                                <Legend />
                            </ComposedChart>
                        </ResponsiveContainer>
                    ) : (
                        <div className={"flex flex-col grow justify-center items-center h-[250px]"} onClick={this.setCreateWeightEntry}>
                            <PlusCircleIcon className={"flex h-12 text-indigo-400"} />
                            <h3 className={"text-cyan-200 text-2xl font-semibold"}>Create the first entry</h3>
                        </div>
                    )}
                </div>
            </div>
        )
    }

    private getRecordTypeColors = (type: string): RecordTheme => {
        let typeColor = {border: "border-indigo-600", icon: "text-indigo-300", background: "bg-indigo-600"}
        if (type === RecordType.REMINDER) {
            typeColor = {border: "border-teal-600", icon: "text-teal-300", background: "bg-teal-600"}
        } else if (type === RecordType.OVER_DUE) {
            typeColor = {border: "border-rose-600", icon: "text-rose-300", background: "bg-rose-600"}
        }
        return typeColor
    }

    private recordEntry = (r: Record, color: RecordTheme) => {
        const upcoming = r.administeredBy === undefined
        return (
            <div
                key={r.id}
                className={`md:items-center flex ${upcoming ? "flex-row" : "grow flex-col"}  md:flex-row border-b  ${
                    color.border
                } justify-between last:border-b-0`}
            >
                <div onClick={() => this.onViewRecordSelected(r)} className={"flex flex-col md:flex-row grow"}>
                    <p className={"flex grow w-full"}>{r.name}</p>
                    <div className={"flex grow w-full"}>{format(new Date(r.date!), "dd/MM/yyyy")}</div>
                </div>
                <div className={"flex flex-row md:gap-x-2 items-center py-0.5"}>
                    {upcoming ? (
                        <></>
                    ) : r.verifiedBy ? (
                        <CheckBadgeIcon
                            className={"h-8 w-8 text-indigo-300 z-1 hover:bg-gray-600 p-0.5 rounded-md"}
                            onClick={() => {
                                this.userDialogRef?.setData(r.verifiedBy!)
                                this.userDialogRef?.setTitle("Vet Information")
                                this.userDialogRef?.show()
                            }}
                        />
                    ) : (
                        <CheckBadgeIconOutline className={"h-8 w-8 text-indigo-300 z-1 p-0.5 rounded-md"} />
                    )}
                    {!upcoming && (
                        <DocumentDuplicateIcon
                            className={"h-8 w-8 text-indigo-300 z-1 hover:bg-gray-600 p-0.5 rounded-md"}
                            onClick={() => this.onCopyRecord(r)}
                        />
                    )}
                    <PencilIcon
                        className={`h-8 w-8 ${color.icon} z-1 hover:bg-gray-600 p-0.5 rounded-md`}
                        onClick={() => this.onUpdateRecordSelected(r)}
                    />
                    <XMarkIcon
                        className={"h-8 w-8 text-red-400 hover:bg-red-600 hover:bg-opacity-20 p-0.5 rounded-md"}
                        onClick={() => {
                            this.setState({selectedRecordId: r.id}, () => this.deleteDialogRef?.show())
                        }}
                    />
                </div>
            </div>
        )
    }

    onDeleteLineRecord = (r: Record) => {
        this.weightEntriesDialogRef?.setData(this.state.records)
        this.setState({selectedRecordId: r.id}, () => {
            this.deleteDialogRef?.show()
            this.updateLineDatasets()
        })
    }

    private updateExpanded(t: string, newValue: boolean) {
        const expandedTypes = this.state.expandedTypes
        expandedTypes.set(t, newValue)
        this.setState({expandedTypes})
    }

    private recordsTypeView(records: Record[], type: string) {
        if (records.length === 0) {
            return <></>
        }
        const expanded = this.state.expandedTypes.get(type)
        const color = this.getRecordTypeColors(type)

        return (
            <div
                key={`header${type}`}
                className={`flex flex-col ${color.border} border rounded-md ${color.background} bg-opacity-10 px-2.5 mb-2`}
            >
                <div key={`header_title_${type}`} className={"py-2 flex flex-row gap-x-4 justify-between"}>
                    <div className={"flex flex-row grow w-full capitalize"} onClick={() => this.updateExpanded(type, !expanded)}>
                        <span className={"flex grow"}>{type}</span>
                        {expanded ? (
                            <ChevronUpIcon className={`static h-6 w-6 self-end ${color.icon} z-1`} />
                        ) : (
                            <ChevronDownIcon className={`static h-6 w-6 self-end ${color.icon} z-1`} />
                        )}
                    </div>
                </div>
                <div className={`${expanded ? "" : "hidden"}`}>
                    <div key={"header"} className={`hidden md:flex flex-row grow pe-24 ${color.border} border-b mb-2`}>
                        <div className={"flex grow capitalize"}>Name</div>
                        <div className={"flex grow capitalize"}>Date</div>
                    </div>
                    {records.map((record) => this.recordEntry(record, color))}
                </div>
            </div>
        )
    }

    private get recordsView() {
        const records = this.state.records.sort((a, b) => a.date! - b.date!).filter((it) => it.administeredBy !== undefined)

        return RecordTypeUtils.getAll().map((t) => {
            const tRecords = records.filter((it) => it.recordType === t)
            return this.recordsTypeView(tRecords, t)
        })
    }

    private get recordsListView() {
        if (this.state.loading) {
            return <></>
        }
        const sorted = this.state.records.sort((a, b) => a.date! - b.date!)
        if (sorted.length === 0) {
            return (
                <div className={"flex flex-col grow justify-center items-center h-[250px]"} onClick={this.setCreateRecord}>
                    <PlusCircleIcon className={"flex h-12 text-indigo-400"} />
                    <h3 className={"text-cyan-200 text-2xl font-semibold"}>Create the first record</h3>
                </div>
            )
        }
        const notFinalized = sorted.filter((it) => it.administeredBy === undefined)
        const overDue = notFinalized.filter((it) => isBefore(new Date(it.date!), new Date()))
        const upcoming = notFinalized.filter((it) => !isBefore(new Date(it.date!), new Date()))

        return (
            <div className={"overflow-y-auto"}>
                {this.recordsTypeView(overDue, RecordType.OVER_DUE)}
                {this.recordsTypeView(upcoming, RecordType.REMINDER)}
                {this.recordsView}
            </div>
        )
    }

    private onEventSelected = async (e: CalendarEvent) => {
        const record = e.resource
        this.setState({selectedRecordId: record.id})
        if (!e.isEvent) {
            this.onCopyRecord(record)
        } else {
            this.onViewRecordSelected(record)
        }
    }

    private updateEvents = () => {
        const events = this.state.records
            .filter((it) => it.recordType !== RecordType.WEIGHT && it.recordType !== RecordType.TEMPERATURE)
            .flatMap((t) => {
                const events: CalendarEvent[] = []
                const d = new Date(t.date!)
                events.push(new CalendarEvent(true, t.name!, d, d, t, true))
                return events
            })
        this.setState({events})
    }

    private get calendarListView() {
        return this.state.events.length > 0 ? (
            <div className={"overflow-y-auto h-[500px]"}>
                <CalendarComponent
                    events={this.state.events}
                    views={{month: true, week: true, day: true, agenda: true}}
                    onSelectEvent={this.onEventSelected}
                    defaultView={"agenda"}
                />
            </div>
        ) : (
            <div className={"flex flex-col grow justify-center items-center h-[250px]"} onClick={this.setCreateRecord}>
                <PlusCircleIcon className={"flex h-12 text-indigo-400"} />
                <h3 className={"text-cyan-200 text-2xl font-semibold"}>Create the first record</h3>
            </div>
        )
    }

    private get ListWidget() {
        return (
            <div
                className={`${
                    this.state.loading ? "animate-pulse" : ""
                } flex flex-col  bg-indigo-300 dark:bg-slate-800 py-4 px-6 rounded-md grow overflow-y-auto`}
            >
                <div className={`sticky flex flex-row items-center justify-between mb-6`}>
                    <div className={"flex flex-row gap-4"}>
                        {RecordViewTypeUtils.getAll().map((it) => (
                            <button
                                key={it}
                                onClick={() => {
                                    this.setState(
                                        {
                                            //@ts-ignore
                                            recordViewType: it,
                                        },
                                        this.updateLineDatasets
                                    )
                                }}
                            >
                                <h2
                                    className={`capitalize justify-self-center text-lg xl:text-3xl font-bold tracking-tight ${
                                        this.state.recordViewType === it
                                            ? `text-indigo-500 dark:text-indigo-300`
                                            : `text-indigo-050 dark:text-indigo-100`
                                    }`}
                                >
                                    {RecordViewTypeUtils.getTitle(it)}
                                </h2>
                            </button>
                        ))}
                    </div>
                    <PlusIcon
                        className={"flex h-8 xl:h-10 text-indigo-500 dark:text-indigo-200 hover:text-indigo-400 "}
                        onClick={this.setCreateRecord}
                    />
                </div>
                {this.state.recordViewType === RecordViewType.RECORDS && this.recordsListView}
                {this.state.recordViewType === RecordViewType.AGENDA && this.calendarListView}
            </div>
        )
    }

    private onCreateLineRecord = async (vm: RecordCreationViewModel) => {
        await this.createRecord(vm).catch((err) => this.setState({serverError: err.message}))
        this.weightEntriesDialogRef?.setData(this.state.records)
    }

    private onUpdateLineRecord = async (vm: UpdateRecordViewModel, recordId: string) => {
        await this.updateRecord(vm, recordId).catch((err) => this.setState({serverError: err.message}))
        this.weightEntriesDialogRef?.setData(this.state.records)
    }

    private onCreateRecord = async (vm: RecordCreationViewModel) => {
        this.recordDialogRef?.hide()
        vm.hasNextDate
            ? await this.createRecords(vm).catch((err) => this.setState({serverError: err.message}))
            : await this.createRecord(vm).catch((err) => this.setState({serverError: err.message}))
    }

    private onUpdateRecord = async (vm: UpdateRecordViewModel) => {
        this.recordDialogRef?.hide()
        await this.updateRecord(vm, this.state.selectedRecordId).catch((err) => this.setState({serverError: err.message}))
    }

    private onCopyRecord = (record: Record) => {
        this.recordDialogRef?.setDataForCreation(record)
        this.recordDialogRef?.show()
    }

    private onUpdateRecordSelected = (record: Record) => {
        this.setState({selectedRecordId: record.id})
        this.recordDialogRef?.setUpdate()
        this.recordDialogRef?.setDataForUpdate(record)
        this.recordDialogRef?.show()
    }

    private onViewRecordSelected = (record: Record) => {
        this.setState({selectedRecordId: record.id})
        this.recordDialogRef?.setViewOnly()
        this.recordDialogRef?.setDataForUpdate(record)
        this.recordDialogRef?.show()
    }

    render() {
        const pet = this.state?.pet
        return (
            <ProtectedPage init={this.initPage}>
                <ErrorMessage
                    className={"absolute mt-2 border border-red-400 w-full z-10"}
                    message={this.state?.serverError}
                    key={"errorMessage"}
                />
                <LineRecordDialog
                    ref={(ref) => (this.weightEntriesDialogRef = ref)}
                    recordType={this.state.selectedLineModel}
                    data={this.state.records.filter((it) => it.recordType === this.state.selectedLineModel)}
                    onDismiss={() => {
                        this.weightEntriesDialogRef?.hide()
                        this.updateLineDatasets()
                    }}
                    onCreate={this.onCreateLineRecord}
                    onUpdate={this.onUpdateLineRecord}
                    onDelete={this.onDeleteLineRecord}
                />
                <UserDialog ref={(ref) => (this.userDialogRef = ref)} onDismiss={() => this.userDialogRef?.hide()} />
                <RecordDialog
                    ref={(ref) => (this.recordDialogRef = ref)}
                    onDismiss={() => this.recordDialogRef?.hide()}
                    onCreate={this.onCreateRecord}
                    onUpdate={this.onUpdateRecord}
                />
                <DeleteModal
                    ref={(ref) => (this.deleteDialogRef = ref)}
                    message={"Are you sure you want to delete this Record?"}
                    onDelete={this.deleteRecord}
                    onCancel={() => this.deleteDialogRef?.hide()}
                />
                {this.state.loading && (
                    <div role="status" className={"absolute z-50 top-1/2 left-1/2 transform -translate-x-1/2 -translate-y-1/2"}>
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
                )}
                {pet !== undefined && (
                    <span className={"flex-col sm:flex-row flex grow py-4 px-4 gap-2 max-h-full"}>
                        <div className={"flex flex-col xs:max-sm:flex-row md:w-2/5 gap-2"}>
                            {this.petCard}
                            {this.weightLine}
                        </div>
                        <div className={"flex flex-col md:max-lg:flex-row grow gap-2 xl:ms-4 md:w-3/5"}>{this.ListWidget}</div>
                    </span>
                )}
            </ProtectedPage>
        )
    }
}

export default withRouter(PetPage)
