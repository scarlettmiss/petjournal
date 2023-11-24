import {UserType} from "@/enums/UserType"
import Dto from "@/models/Dto"

export interface Auth {
    token?: string
    user?: User
}

export interface User extends Dto {
    userType?: UserType
    email?: string
    name?: string
    surname?: string
    phone?: string
    address?: string
    city?: string
    state?: string
    country?: string
    zip?: string
}
