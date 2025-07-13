import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

export const formatDate: (value: string) => string = (value: string): string => {
  return new Intl.DateTimeFormat("es-ES", {
    month: "short",
    day: "2-digit"
  }).format(new Date(value))
};
