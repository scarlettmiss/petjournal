export enum RecordType {
    VACCINE = "vaccine",
    WEIGHT = "weight",
    TEMPERATURE = "temperature",
    SURGERY = "surgery",
    MEDICINE = "medicine",
    ENDOPARASITE = "endoparasite",
    ECTOPARASITE = "ectoparasite",
    EXAMINATION = "examination",
    MICROCHIP = "microchip",
    DIAGNOSTIC = "diagnostic",
    DENTAL = "dental",
    OTHER = "other",
}

export class RecordTypeUtils {
    public static getAll(): string[] {
        return Object.values(RecordType).filter(it => it !== RecordType.WEIGHT && it !== RecordType.TEMPERATURE)
    }

    public static getEnum(userType: string): RecordType {
        return RecordType[userType as keyof typeof RecordType]
    }

    public static getTitle(type: string): string {
        return type
    }
}
