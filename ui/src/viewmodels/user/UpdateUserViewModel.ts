import TextUtils from "@/Utils/TextUtils"

export default class UpdateUserViewModel {
    private _email: string = ""
    private _name: string = ""
    private _surname: string = ""
    private _phone: string = ""
    private _address: string = ""
    private _city: string = ""
    private _state: string = ""
    private _country: string = ""
    private _zip: string = ""

    public nameError: string = ""
    public surnameError: string = ""
    public emailError: string = ""

    public hasNameError: boolean = false
    public hasSurnameError: boolean = false
    public hasEmailError: boolean = false

    private _valid: boolean = false

    public get stringify() {
        return JSON.stringify(this.fields)
    }

    validate() {
        this.nameError = TextUtils.isNotEmpty(this._name) ? "" : "Name is required"
        this.surnameError = TextUtils.isNotEmpty(this._surname) ? "" : "Surname is required"
        this.emailError = TextUtils.isEmailValid(this._email) ? "" : "Please provide a valid email"

        this._valid = TextUtils.isEmpty(this.nameError) && TextUtils.isEmpty(this.surnameError) && TextUtils.isEmpty(this.emailError)
    }

    public get isValid(): boolean {
        return this._valid
    }

    public checkEmail = () => {
        const isValid = TextUtils.isEmailValid(this._email)
        this.emailError = isValid ? "" : "Please provide a valid email"
        this.hasEmailError = !isValid
    }

    public checkName = () => {
        const isEmpty = TextUtils.isEmpty(this._name)
        this.nameError = isEmpty ? "Name is required" : ""
        this.hasNameError = isEmpty
    }

    public checkSurname = () => {
        const isEmpty = TextUtils.isEmpty(this._surname)
        this.surnameError = isEmpty ? "Surname is required" : ""
        this.hasSurnameError = isEmpty
    }

    get email(): string {
        return this._email
    }

    set email(value: string) {
        this._email = value
        this.hasEmailError = false
    }

    get city(): string {
        return this._city
    }

    set city(value: string) {
        this._city = value
    }

    get state(): string {
        return this._state
    }

    set state(value: string) {
        this._state = value
    }

    get country(): string {
        return this._country
    }

    set country(value: string) {
        this._country = value
    }

    get zip(): string {
        return this._zip
    }

    set zip(value: string) {
        this._zip = value
    }

    get address(): string {
        return this._address
    }

    set address(value: string) {
        this._address = value
    }

    get phone(): string {
        return this._phone
    }

    set phone(value: string) {
        this._phone = value
    }

    get surname(): string {
        return this._surname
    }

    set surname(value: string) {
        this._surname = value
        this.hasSurnameError = false
    }

    get name(): string {
        return this._name
    }

    set name(value: string) {
        this._name = value
        this.hasNameError = false
    }

    public get fields() {
        return {
            email: this._email,
            name: this._name,
            surname: this._surname,
            phone: this._phone,
            address: this._address,
            city: this._city,
            state: this._state,
            country: this._country,
            zip: this._zip,
        }
    }
}
