import PetCreationViewModel from "@/viewmodels/pet/PetCreationViewModel"
import PetUpdateViewModel from "@/viewmodels/pet/UpdatePetViewModel"

export function petCreationHandler(vm: PetCreationViewModel, token?: string): Promise<Response> {
    return fetch(`/api/pet`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(vm.fields),
    })
}

export function petHandler(id: string, token?: string): Promise<Response> {
    return fetch(`/api/pet/${id}`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    })
}

export function petUpdateHandler(vm: PetUpdateViewModel, id: string, token?: string): Promise<Response> {
    return fetch(`/api/pet/${id}`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(vm.fields),
    })
}

export function petDeleteHandler(id: string, token?: string): Promise<Response> {
    return fetch(`/api/pet/${id}`, {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    })
}

export function petsHandler(token?: string): Promise<Response> {
    return fetch(`/api/pets`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    })
}
