import UpdateUserViewModel from "@/viewmodels/user/UpdateUserViewModel"
import LoginViewModel from "@/viewmodels/user/LoginViewModel";
import RegistrationViewModel from "@/viewmodels/user/RegistrationViewModel";

export function loginHandler(vm: LoginViewModel): Promise<Response> {
    return fetch(`/api/auth/login`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(vm),
    })
}

export function signUpHandler(vm: RegistrationViewModel): Promise<Response> {
    return fetch(`/api/auth/register`, {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
        },
        body: JSON.stringify(vm.fields),
    })
}

export function userUpdateHandler(vm: UpdateUserViewModel, token?: string): Promise<Response> {
    return fetch(`/api/user`, {
        method: "PATCH",
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${token}`,
        },
        body: JSON.stringify(vm.fields),
    })
}

export function userHandler(token?: string): Promise<Response> {
    return fetch(`/api/user`, {
        method: "GET",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    })
}

export function userDeleteHandler(token?: string): Promise<Response> {
    return fetch(`/api/user`, {
        method: "DELETE",
        headers: {
            Authorization: `Bearer ${token}`,
        },
    })
}

export function VetsHandler(): Promise<Response> {
    return fetch(`/api/vets`, {
        method: "GET",
        headers: {
            "Content-Type": "application/json",
        },
    })
}
