import React, {Component} from "react"
import {XMarkIcon} from "@heroicons/react/20/solid"
import Button from "@/components/Button"
import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import {Record} from "@/models/record/Record"
import TextInput from "@/components/TextInput"
import {addDays, format, isBefore} from "date-fns"
import {RecordTypeUtils} from "@/enums/RecordType"
import styles from "@/components/textInput.module.css"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"
import {Pet} from "@/models/pet/Pet"
import Avatar from "@/components/Avatar"

interface RecordDialogProps {
    data?: Record
    onDismiss: () => void
    onCreate?: (vm: RecordCreationViewModel) => void
    onUpdate?: (vm: UpdateRecordViewModel) => void
    mode?: Mode
}

interface RecordDialogState {
    title: string
    vm: RecordCreationViewModel | UpdateRecordViewModel
    show: boolean
    mode: Mode
    step: Step
    pet?: Pet
}

export enum Mode {
    CREATE = "CREATE",
    UPDATE = "UPDATE",
    VIEW = "VIEW",
}

export enum Step {
    BASE = "BASE",
    EXTRA_INFO = "EXTRA_INFO",
    NEXT_DATE = "NEXT_DATE",
    LINEDATA = "LINEDATA",
}

export default class RecordDialog extends Component<RecordDialogProps, RecordDialogState> {
    constructor(props: RecordDialogProps) {
        super(props)
        this.state = {
            title: "Create Record",
            vm: new RecordCreationViewModel(),
            show: false,
            mode: props.mode ?? Mode.CREATE,
            step: Step.BASE,
        }
    }

    public setDataForCreation = (data: Record, pet?: Pet) => {
        const vm = new RecordCreationViewModel()
        vm.name = data.name ?? ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        this.setCreate()
        this.setState(
            {
                step: Step.BASE,
                vm,
                title: "Create Record",
                pet,
            },
            () => setTimeout(() => document?.getElementById("date")?.focus(), 1)
        )
    }

    public setDataForCreationWithDate = (data: Record, date: Date, pet?: Pet) => {
        const vm = new RecordCreationViewModel()
        vm.name = data.name ?? ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        vm.date = format(date, "yyyy-MM-dd")
        this.setCreate()
        this.setState(
            {
                step: Step.BASE,
                vm,
                title: "Create Record",
                pet,
            },
            () => setTimeout(() => document?.getElementById("date")?.focus(), 1)
        )
    }

    public setDataForUpdate = (data: Record, pet?: Pet) => {
        const vm = new UpdateRecordViewModel()
        vm.name = data.name ?? ""
        vm.date = data.date ? format(new Date(data.date), "yyyy-MM-dd") : ""
        vm.recordType = data.recordType ?? ""
        vm.name = data.name ?? ""
        vm.lot = data.lot ?? ""
        vm.result = data.result ?? ""
        vm.description = data.description ?? ""
        vm.notes = data.notes ?? ""
        this.setState({step: Step.BASE, vm, title: "Update Record", pet})
    }

    private get title() {
        switch (this.state.mode) {
            case Mode.CREATE:
                return "Create Record"
            case Mode.UPDATE:
                return "Update Record"
            case Mode.VIEW:
            default:
                return "Record"
        }
    }

    public hide = () => {
        this.setState({show: false, vm: new RecordCreationViewModel()})
    }

    public show = () => {
        this.setState({show: true})
    }

    public setCreate = () => {
        this.setState({mode: Mode.CREATE, step: Step.BASE})
    }

    public setUpdate = () => {
        this.setState({mode: Mode.UPDATE})
    }

    public setViewOnly = () => {
        this.setState({mode: Mode.VIEW})
    }

    private onNameChange = (value: string) => {
        const vm = this.state.vm
        vm.name = value
        this.setState({vm})
    }

    private onLotChange = (value: string) => {
        const vm = this.state.vm
        vm.lot = value
        this.setState({vm})
    }

    private onResultChange = (value: string) => {
        const vm = this.state.vm
        vm.result = value
        this.setState({vm})
    }

    private onDescriptionChange = (value: string) => {
        const vm = this.state.vm
        vm.description = value
        this.setState({vm})
    }

    private onNotesChange = (value: string) => {
        const vm = this.state.vm
        vm.notes = value
        this.setState({vm})
    }

    private onDateChange = (value: string) => {
        const vm = this.state.vm
        vm.date = value
        console.log(vm.date)
        this.setState({vm})
    }

    private onRecordTypeChange = (event: any) => {
        const vm = this.state.vm
        vm.recordType = event.currentTarget.value
        this.setState({vm})
    }

    private onHasNextDateChange = () => {
        const vm = this.state.vm
        vm.hasNextDate = !vm.hasNextDate
        this.setState({vm})
    }

    private onNextDateChange = (value: string) => {
        const vm = this.state.vm as RecordCreationViewModel
        const minDate = addDays(new Date(), 1)
        vm.nextDate = isBefore(new Date(value), minDate) ? format(minDate, "yyyy-MM-dd").toString() : value
        this.setState({vm})
    }

    private checkName = () => {
        this.state.vm.checkName()
        this.forceUpdate()
    }

    private checkDate = () => {
        this.state.vm.checkDate()
        this.forceUpdate()
    }

    private checkNextDate = () => {
        ;(this.state.vm as RecordCreationViewModel).checkNextDate()
        this.forceUpdate()
    }

    private checkType = () => {
        this.state.vm.checkType()
        this.forceUpdate()
    }

    private get primaryButtonTitle() {
        const isView = this.state.mode === Mode.VIEW

        if (this.state.step === Step.BASE) {
            return isView ? "More Info" : "Add Extra Information"
        }

        if (this.state.vm.hasNextDate && this.state.step === Step.EXTRA_INFO) {
            return "Schedule Next Date"
        }

        switch (this.state.mode) {
            case Mode.VIEW:
                return "Done"
            case Mode.CREATE:
                return "Create Record"
            case Mode.UPDATE:
                return "Update Record"
        }
    }

    private get secondaryButtonTitle() {
        if (this.state.mode === Mode.VIEW) {
            return "Edit"
        }
        if (this.state.step !== Step.BASE) {
            return "Back"
        }

        return "Cancel"
    }

    private primaryButtonPress = (): void => {
        if (this.state.mode !== Mode.VIEW) {
            this.state.vm.validate(this.state.step)
            if (!this.state.vm.isValid) {
                this.forceUpdate()
                return
            }
        }
        if (this.state.step === Step.BASE) {
            this.setState({step: Step.EXTRA_INFO})
            return
        } else if (this.state.vm.hasNextDate && this.state.step === Step.EXTRA_INFO) {
            this.setState({step: Step.NEXT_DATE})
            return
        }

        switch (this.state.mode) {
            case Mode.VIEW:
                this.hide()
                break
            case Mode.CREATE:
                return this.props.onCreate!(this.state.vm as RecordCreationViewModel)
            case Mode.UPDATE:
                return this.props.onUpdate!(this.state.vm as UpdateRecordViewModel)
        }
    }

    private secondButtonPress = () => {
        if (this.state.mode === Mode.VIEW) {
            this.setUpdate()
            this.setState({step: Step.BASE})
            return
        }
        if (this.state.step === Step.BASE) {
            this.props.onDismiss()
            return
        } else if (this.state.step == Step.EXTRA_INFO) {
            this.setState({step: Step.BASE})
            return
        } else if (this.state.step == Step.NEXT_DATE) {
            this.setState({step: Step.EXTRA_INFO})
            return
        }
    }

    private get formBody() {
        const vm = this.state.vm
        return (
            <>
                {this.state.step === Step.BASE && (
                    <>
                        <TextInput
                            autoFocus={false}
                            id="name"
                            name="name"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Name"
                            value={vm.name}
                            onInput={this.onNameChange}
                            onBlur={this.checkName}
                            hasError={vm.hasNameError}
                            errorMessage={vm.nameError}
                            disabled={this.state.mode === Mode.VIEW}
                            showLabel={true}
                        />

                        <div>
                            <label htmlFor={"recordType"} className={styles.label}>
                                Record Type
                            </label>
                            <select
                                disabled={this.state.mode === Mode.VIEW}
                                value={vm.recordType}
                                id="recordType"
                                name="recordType"
                                onChange={this.onRecordTypeChange}
                                onBlur={this.checkType}
                                className={`${styles.textInput} ${styles.textInputFull} ${
                                    vm.hasRecordTypeError && styles.errorHighlight
                                } rounded-md capitalize`}
                            >
                                <option value="">Choose a Record Type</option>
                                {RecordTypeUtils.getAll().map((g) => {
                                    const name = RecordTypeUtils.getTitle(g)
                                    return (
                                        <option key={name} value={g} className={"capitalize"}>
                                            {name}
                                        </option>
                                    )
                                })}
                            </select>
                            <span className={styles.errorMessage}>{vm.hasRecordTypeError ? vm.recordTypeError : ""}</span>
                        </div>

                        <TextInput
                            id="date"
                            name="date"
                            type="date"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Administration Date"
                            value={vm.date}
                            onInput={this.onDateChange}
                            onBlur={this.checkDate}
                            hasError={vm.hasDateError}
                            errorMessage={vm.dateError}
                            showLabel={true}
                            disabled={this.state.mode === Mode.VIEW}
                        />
                    </>
                )}
                {this.state.step === Step.EXTRA_INFO && (
                    <>
                        <TextInput
                            id="lot"
                            name="lot"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Lot Number"
                            value={vm.lot}
                            onInput={this.onLotChange}
                            showLabel
                            disabled={this.state.mode === Mode.VIEW}
                        />
                        <TextInput
                            id="result"
                            name="result"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Result"
                            value={vm.result}
                            onInput={this.onResultChange}
                            showLabel
                            disabled={this.state.mode === Mode.VIEW}
                        />
                        <TextInput
                            id="description"
                            name="description"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Description"
                            value={vm.description}
                            onInput={this.onDescriptionChange}
                            showLabel
                            disabled={this.state.mode === Mode.VIEW}
                        />
                        <TextInput
                            id="notes"
                            name="notes"
                            type="text"
                            autoComplete="off"
                            required
                            width={"full"}
                            classNames="rounded-md"
                            placeholder="Notes"
                            value={vm.notes}
                            onInput={this.onNotesChange}
                            showLabel
                            disabled={this.state.mode === Mode.VIEW}
                        />
                        {this.state.mode === Mode.CREATE && (
                            <div>
                                <input
                                    id="link-checkbox"
                                    type="checkbox"
                                    value={""}
                                    checked={vm.hasNextDate}
                                    onChange={this.onHasNextDateChange}
                                    className="w-4 h-4 text-indigo-600 rounded focus:ring-indigo-600 ring-offset-gray-800 focus:ring-2 bg-gray-700 border-gray-600"
                                />
                                <label htmlFor="link-checkbox" className="ml-2 text-sm font-medium text-gray-300">
                                    Is recurring
                                </label>
                            </div>
                        )}
                    </>
                )}
                {this.state.step === Step.NEXT_DATE && (
                    <TextInput
                        id="date"
                        name="date"
                        type="date"
                        autoComplete="off"
                        required
                        width={"full"}
                        min={format(addDays(new Date(), 1), "yyyy-MM-dd")}
                        classNames="rounded-md"
                        placeholder="Schedule Next Occurance"
                        value={(vm as RecordCreationViewModel).nextDate}
                        onInput={this.onNextDateChange}
                        onBlur={this.checkNextDate}
                        hasError={vm.hasNextDateError}
                        errorMessage={vm.dateError}
                        showLabel={true}
                    />
                )}
            </>
        )
    }

    render() {
        return (
            <div id="defaultModal" tabIndex={-1} className={`fixed flex grow ${this.state?.show ? "" : "hidden"} z-50 h-screen w-full`}>
                <div className="relative self-center mx-auto w-full max-w-2xl max-h-full">
                    {/*Modal content*/}
                    <div className="relative rounded-lg shadow bg-gray-700/70 backdrop-blur-lg">
                        {/*Modal header*/}
                        <div className="flex items-start justify-between p-4 border-b rounded-t border-gray-600">
                            {this.state.pet && (
                                <div className={"flex flex-row gap-x-2 mx-3"}>
                                    <Avatar
                                        avatarTitle={this.state.pet?.name?.slice(0, 1) ?? "-"}
                                        avatar={this.state.pet.avatar}
                                        className={"self-center h-[30px] w-[30px] "}
                                    />
                                    <h3 className="text-xl font-semibold text-white text-center">
                                        {this.state.pet?.name}
                                    </h3>
                                </div>
                            )}
                            <h3 className="text-xl font-semibold text-white text-center">{this.title}</h3>
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

                        <div className="px-12 py-4 space-y-4 max-h-96 overflow-y-auto justify-center">
                            <div className="shadow-sm px-4 py-4">
                                <form className="space-y-1 mb-3" method="POST" onSubmit={this.props.onDismiss}>
                                    {this.formBody}
                                </form>
                            </div>
                        </div>
                        {/*footer*/}
                        <div className="flex items-center p-6 justify-end space-x-2 border-t rounded-b border-gray-600">
                            <Button
                                className={"right-0"}
                                title={this.secondaryButtonTitle}
                                variant={"secondary"}
                                type={"button"}
                                onClick={this.secondButtonPress}
                            />
                            <Button
                                className={"right-0"}
                                title={this.primaryButtonTitle}
                                variant={"primary"}
                                type={"button"}
                                onClick={this.primaryButtonPress}
                            />
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}
