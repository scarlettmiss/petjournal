export enum RecordViewType {
    RECORDS = "RECORDS",
    AGENDA = "AGENDA",
}

export class RecordViewTypeUtils {
    public static getAll(): RecordViewType[] {
        return Object.values(RecordViewType)
    }

    public static getEnum(userType: string): RecordViewType {
        return RecordViewType[userType as keyof typeof RecordViewType]
    }

    public static getTitle(type: string): string {
        return type.toLowerCase()
    }
}
