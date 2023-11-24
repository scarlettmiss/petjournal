import TextUtils from "@/Utils/TextUtils"
import {PetMaxWeight, PetMinWeight} from "@/models/pet/Pet";

export default class PetUpdateViewModel {
    private _avatar: string = ""
    private _name: string = ""
    private _dateOfBirth: string = ""
    private _gender: string = ""
    private _breedName: string = ""
    private _colors: string[] = []
    private _description: string = ""
    private _pedigree: string = ""
    private _microchip: string = ""
    private _vetId: string = ""
    private _metas: Map<string, string> = new Map()

    public nameError: string = ""
    public dateOfBirthError: string = ""
    public genderError: string = ""
    public breedNameError: string = ""

    public hasNameError: boolean = false
    public hasDateOfBirthError: boolean = false
    public hasGenderError: boolean = false
    public hasBreedNameError: boolean = false

    private _valid: boolean = false

    public get stringify() {
        return JSON.stringify(this)
    }

    validate() {
        this.checkName()
        this.checkDateOfBirth()
        this.checkGender()
        this.checkBreedName()
        this._valid =
            TextUtils.isEmpty(this.nameError) &&
            TextUtils.isEmpty(this.dateOfBirthError) &&
            TextUtils.isEmpty(this.genderError) &&
            TextUtils.isEmpty(this.breedNameError)
    }

    public get isValid(): boolean {
        return this._valid
    }

    public checkDateOfBirth = () => {
        const isEmpty = TextUtils.isEmpty(this._dateOfBirth)
        this.dateOfBirthError = isEmpty ? "The date of birth is required" : ""
        this.hasDateOfBirthError = isEmpty
    }

    public checkGender = () => {
        const isEmpty = TextUtils.isEmpty(this._gender)
        this.genderError = isEmpty ? "Pet gender is required" : ""
        this.hasGenderError = isEmpty
    }

    public checkName = () => {
        const isEmpty = TextUtils.isEmpty(this._name)
        this.nameError = isEmpty ? "Name is required" : ""
        this.hasNameError = isEmpty
    }

    public checkBreedName = () => {
        const isEmpty = TextUtils.isEmpty(this._breedName)
        this.breedNameError = isEmpty ? "Please provide a breed name" : ""
        this.hasBreedNameError = isEmpty
    }

    get avatar(): string {
        return this._avatar
    }

    set avatar(value: string) {
        this._avatar = value
    }

    get name(): string {
        return this._name
    }

    set name(value: string) {
        this._name = value
    }

    get dateOfBirth(): string {
        return this._dateOfBirth!
    }

    set dateOfBirth(value: string) {
        this._dateOfBirth = value
    }

    get gender(): string {
        return this._gender!
    }

    set gender(value: string) {
        this._gender = value
    }

    get breedName(): string {
        return this._breedName
    }

    set breedName(value: string) {
        this._breedName = value
    }

    get colors(): string[] {
        return this._colors
    }

    set colors(value: string[]) {
        this._colors = value
    }

    get description(): string {
        return this._description
    }

    set description(value: string) {
        this._description = value
    }

    get pedigree(): string {
        return this._pedigree
    }

    set pedigree(value: string) {
        this._pedigree = value
    }

    get microchip(): string {
        return this._microchip
    }

    set microchip(value: string) {
        this._microchip = value
    }

    get weightMin(): string | undefined {
        return this._metas?.get(PetMinWeight)
    }

    set weightMin(value: string | undefined) {
        this._metas?.set(PetMinWeight, value ?? '')
    }

    get weightMax(): string | undefined {
        return this._metas?.get(PetMaxWeight)
    }

    set weightMax(value: string | undefined) {
        this._metas?.set(PetMaxWeight, value ?? '')
    }

    get vetId(): string {
        return this._vetId
    }

    set vetId(value: string) {
        this._vetId = value
    }

    get metas(): Map<string, string> {
        return this._metas
    }

    set metas(value: Map<string, string>) {
        this._metas = value
    }

    public get fields() {
        return {
            avatar: this._avatar,
            name: this._name,
            dateOfBirth: new Date(this._dateOfBirth!).getTime(),
            gender: this._gender,
            breedName: this._breedName,
            colors: this._colors,
            description: this._description,
            pedigree: this._pedigree,
            microchip: this._microchip,
            vetId: this._vetId,
            metas: Object.fromEntries(this._metas),
        }
    }
}
