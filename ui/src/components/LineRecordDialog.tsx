import React, {Component} from "react"
import {PencilIcon, XMarkIcon} from "@heroicons/react/20/solid"
import {format} from "date-fns"
import TextInput from "@/components/TextInput"
import Button from "@/components/Button"
import {Record} from "@/models/record/Record"
import TextUtils from "@/Utils/TextUtils"
import {PlusCircleIcon} from "@heroicons/react/24/outline"
import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"
import {Step} from "@/components/RecordDialog"
import {RecordType, RecordTypeUtils} from "@/enums/RecordType"

interface LineRecordDialogProps {
    data: Record[]
    recordType: RecordType
    onDismiss: () => void
    onCreate: (vm: RecordCreationViewModel) => void
    onUpdate: (vm: UpdateRecordViewModel, id: string) => void
    onDelete: (r: Record) => void
    mode?: Mode
}

interface LineRecordDialogState {
    recordType: RecordType
    vm: RecordCreationViewModel | UpdateRecordViewModel
    show: boolean
    itemToEdit?: Record
    recordIdToUpdate?: string
    mode: Mode
    updatedData: Record[]
}

enum Mode {
    CREATE,
    EDIT,
    LIST,
}

export default class LineRecordDialog extends Component<LineRecordDialogProps, LineRecordDialogState> {
    constructor(props: LineRecordDialogProps) {
        super(props)
        const vm = new RecordCreationViewModel()
        vm.recordType = this.props.recordType
        this.state = {
            vm: vm,
            show: false,
            updatedData: props.data,
            mode: Mode.LIST,
            recordType: this.props.recordType,
        }
    }

    public hide = () => {
        this.setState({show: false})
    }

    public show = () => {
        this.setState({show: true})
    }

    public setRecordType = (recordType: RecordType) => {
        const vm = this.state.vm
        vm.recordType = this.state.recordType
        this.setState({recordType, vm})
    }

    public setData = (data: Record[]) => {
        this.setState({updatedData: data.filter((it) => it.recordType === this.state.recordType)})
    }

    public setCreate = () => {
        const vm = new RecordCreationViewModel()
        vm.recordType = this.state.recordType
        this.setState({mode: Mode.CREATE, vm})
    }

    public setList = () => {
        this.setState({mode: Mode.LIST})
    }

    public setEdit = () => {
        this.setState({mode: Mode.EDIT})
    }

    private onDateChange = (value: string) => {
        const vm = this.state.vm
        vm.dateError = ""
        vm.date = value
        this.setState({vm})
    }

    private onEntryInput = (value: string) => {
        const vm = this.state.vm
        vm.resultError = ""
        vm.result = value.replace(",", ".")
        this.setState({vm})
    }

    private loadEntryToVmToEdit(entry: Record) {
        const vm = this.state.vm
        vm.recordType = entry.recordType!
        vm.date = format(new Date(entry.date!), "yyyy-MM-dd")
        vm.result = entry.result!
        this.setState({recordIdToUpdate: entry.id, vm})
        this.setEdit()
    }

    private petRemoveEntry = (index: number) => {
        const entries = this.state.updatedData
        if (!entries[index]) {
            return
        }

        this.props.onDelete(entries[index])
        this.setState({updatedData: entries})
    }

    private petAddEntry = () => {
        const vm = this.state.vm
        vm.validate(Step.LINEDATA)

        if (!vm.isValid) {
            this.forceUpdate()
            return
        }

        this.props.onCreate(vm as RecordCreationViewModel)
        this.setList()
    }

    private petEditEntry = () => {
        const vm = this.state.vm
        vm.validate(Step.LINEDATA)

        if (!vm.isValid) {
            this.forceUpdate()
            return
        }

        if (this.state.recordIdToUpdate === undefined) {
            return
        }

        this.props.onUpdate(vm as UpdateRecordViewModel, this.state.recordIdToUpdate)
        this.setList()
    }

    private get primaryButtonTitle() {
        let title = ""
        switch (this.state.mode) {
            case Mode.CREATE:
                title = "Create Entry"
                break
            case Mode.LIST:
                title = "Done"
                break
            case Mode.EDIT:
                title = "Update Entry"
                break
        }
        return title
    }

    private primaryButtonPress = (): void => {
        switch (this.state.mode) {
            case Mode.CREATE:
                return this.petAddEntry()
            case Mode.EDIT:
                return this.petEditEntry()
            case Mode.LIST:
                return this.props.onDismiss()
        }
    }

    private dismissButton = () => {
        switch (this.state.mode) {
            case Mode.EDIT:
            case Mode.CREATE:
                return this.setList()
            case Mode.LIST:
                return this.props.onDismiss()
        }
    }

    render() {
        return (
            <div id="defaultModal" tabIndex={-1} className={`fixed flex grow ${this.state?.show ? "" : "hidden"} z-50 h-screen w-full`}>
                <div className="relative self-center mx-auto w-full max-w-2xl max-h-full">
                    {/*Modal content*/}
                    <div className="relative bg-white rounded-lg shadow dark:bg-gray-700/60 backdrop-blur-lg">
                        {/*Modal header*/}
                        <div className="flex items-start justify-between p-4 border-b rounded-t dark:border-gray-600">
                            <h3 className="text-xl font-semibold text-gray-900 dark:text-white capitalize">
                                {RecordTypeUtils.getTitle(this.state.recordType)} Entries
                            </h3>
                            <button
                                type="button"
                                className="text-gray-400 bg-transparent hover:bg-gray-200 hover:text-gray-900 rounded-lg text-sm w-8 h-8 ml-auto inline-flex justify-center items-center dark:hover:bg-gray-600 dark:hover:text-white"
                                onClick={this.props.onDismiss}
                            >
                                <XMarkIcon className={"h-6 w-6"} />
                                <span className="sr-only">Close modal</span>
                            </button>
                        </div>
                        {/*Modal body */}
                        {(this.state.mode === Mode.CREATE || this.state.mode === Mode.EDIT) && (
                            <div className="p-6 space-y-6 max-h-60 overflow-y-auto">
                                <TextInput
                                    id="dateOfBirth"
                                    name="dateOfBirth"
                                    type="date"
                                    max={format(new Date(), "yyyy-MM-dd")}
                                    autoComplete="off"
                                    required
                                    width={"full"}
                                    classNames="rounded-md"
                                    placeholder="Date"
                                    value={TextUtils.isNotEmpty(this.state.vm.date) ? this.state.vm.date : undefined}
                                    onInput={this.onDateChange}
                                    showLabel
                                    hasError={TextUtils.isNotEmpty(this.state.vm.dateError)}
                                    errorMessage={this.state.vm.dateError}
                                />
                                <TextInput
                                    id="result"
                                    name="result"
                                    type="text"
                                    autoComplete="off"
                                    width={"full"}
                                    classNames="rounded-md"
                                    placeholder={RecordTypeUtils.getTitle(this.state.recordType)}
                                    value={TextUtils.isNotEmpty(this.state.vm.result) ? this.state.vm.result : undefined}
                                    onInput={this.onEntryInput}
                                    hasError={this.state.vm.hasResultError}
                                    errorMessage={this.state.vm.resultError}
                                    showLabel
                                />
                            </div>
                        )}
                        {this.state.mode === Mode.LIST && (
                            <div>
                                <div key={"header"} className={"px-6 py-2 flex flex-row gap-x-4"}>
                                    <div className={"flex grow"}>Date</div>
                                    <div className={"flex grow capitalize"}>{RecordTypeUtils.getTitle(this.state.recordType)}</div>
                                </div>
                                <div className="p-6 space-y-3 max-h-60 overflow-y-auto">
                                    {this.state.updatedData?.map((d, index) => (
                                        <div key={d.date! + index} className={"flex flex-row border-b border-indigo-600 justify-between"}>
                                            <div className={"flex flex-row grow"}>
                                                <div className={"flex grow"}>{format(new Date(d.date!), "dd/MM/yyyy")}</div>
                                                <div className={"flex grow"}>{d.result}</div>
                                            </div>
                                            <div className={"flex flex-row gap-x-2 items-center"}>
                                                <PencilIcon
                                                    className={"h-6 w-6 text-indigo-300"}
                                                    onClick={() => {
                                                        this.loadEntryToVmToEdit(d)
                                                    }}
                                                />
                                                <XMarkIcon className={"h-8 w-6 text-red-500"} onClick={() => this.petRemoveEntry(index)} />
                                            </div>
                                        </div>
                                    ))}
                                </div>
                                <div className={"flex py-3 justify-center hover:bg-gray-600 rounded-md"} onClick={this.setCreate}>
                                    <button
                                        type={"button"}
                                        className={"flex flex-row justify-center items-center gap-2"}
                                        onClick={this.setCreate}
                                    >
                                        <PlusCircleIcon className={"flex h-6 text-indigo-400"} />
                                        <h3 className={"text-cyan-200  font-semibold"}>Add entry</h3>
                                    </button>
                                </div>
                            </div>
                        )}

                        {/*footer*/}
                        <div className="flex items-center p-6 justify-end space-x-2 border-t border-gray-200 rounded-b dark:border-gray-600">
                            {this.state.mode !== Mode.LIST && (
                                <Button
                                    className={"right-0"}
                                    title={"Cancel"}
                                    variant={"secondary"}
                                    type={"button"}
                                    onClick={this.dismissButton}
                                />
                            )}
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
