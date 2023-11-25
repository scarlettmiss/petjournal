import Dto from "@/models/Dto"
import {User} from "@/models/user/User"
import {RecordType} from "@/enums/RecordType"
import {Pet} from "@/models/pet/Pet";

export interface Record extends Dto {
    pet?: Pet
    recordType?: RecordType
    name?: string
    date?: number
    lot?: string
    result?: string
    description?: string
    notes?: string
    administeredBy?: User
    verifiedBy?: User
    groupId?: string
}
