import Dto from "@/models/Dto"
import {User} from "@/models/user/User"

export const PetMinWeight = "petMinWeight"
export const PetMaxWeight = "petMaxWeight"

export interface Pet extends Dto {
    name?: string
    dateOfBirth?: string
    gender?: string
    breedName?: string
    colors?: string[]
    description?: string
    pedigree?: string
    microchip?: string
    owner?: User
    vet?: User
    metas?: Object
    avatar?: string
}
