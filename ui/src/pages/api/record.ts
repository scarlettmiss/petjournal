import RecordCreationViewModel from "@/viewmodels/record/RecordCreationViewModel"
import UpdateRecordViewModel from "@/viewmodels/record/UpdateRecordViewModel"

export function recordCreationHandler(vm: RecordCreationViewModel, petId: string, token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/pet/${petId}/record`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: vm.stringify,
    })
}

export function recordsCreationHandler(vm: RecordCreationViewModel, petId: string, token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/pet/${petId}/records`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: vm.stringify,
    })
}

export function recordUpdateHandler(vm: UpdateRecordViewModel, id: string, petId: string, token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/pet/${petId}/record/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: vm.stringify,
    })
}

export function recordDeletionHandler(rId: string, petId: string, token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/pet/${petId}/record/${rId}`, {
        method: "DELETE",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    })
}

export function petRecordsHandler(petId: string, token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/pet/${petId}/records`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    })
}

export function recordHandler(token?: string): Promise<Response> {
    return fetch(`${process.env.apiUrl}/records`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
    })
}
