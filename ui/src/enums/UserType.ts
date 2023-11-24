export enum UserType {
    OWNER = "OWNER",
    VET = "VET",
}

export class UserTypeUtils {
    public static getAll(): string[] {
        return Object.values(UserType)
    }

    public static getEnum(userType: string): UserType {
        return UserType[userType as keyof typeof UserType]
    }

    public static getTitle(type: string): string {
        let title
        switch (type.toUpperCase()) {
            case UserType.OWNER:
                title = "Owner"
                break
            case UserType.VET:
                title = "Vet"
                break
            default:
                title = ""
                break
        }
        return title
    }
}
