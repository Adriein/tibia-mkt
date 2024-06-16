export const formatDate = (value: Date): string => {
    return new Intl.DateTimeFormat("es-ES", {
        month: "short",
        day: "2-digit"
    }).format(value)
}