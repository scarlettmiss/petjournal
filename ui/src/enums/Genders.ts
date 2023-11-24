export enum PetGender {
    M = "M",
    F = "F",
}

export class PetGenderUtils {
    public static getAll(): string[] {
        return Object.values(PetGender)
    }

    public static getEnum(userType: string): PetGender {
        return PetGender[userType as keyof typeof PetGender]
    }

    public static getTitle(type?: string): string {
        let title
        switch (type) {
            case PetGender.M:
                title = "Male"
                break
            case PetGender.F:
                title = "Female"
                break
            default:
                title = ""
                break
        }
        return title
    }
}
