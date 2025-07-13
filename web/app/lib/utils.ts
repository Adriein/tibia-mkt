import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import type {Price} from "~/home/types";

//STYLE

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

//DATES

export const formatDate: (value: string) => string = (value: string): string => {
  return new Intl.DateTimeFormat("es-ES", {
    month: "short",
    day: "2-digit",
    year: "2-digit"
  }).format(new Date(value))
};

//CHARTS

export const xAxisTick = (data: Price[], xAxisDomain: string[]): string[] => {
  const SHOW_DATES: string[] = xAxisDomain;
  const result: string[] = [];

  for (let i: number = 0; i < data.length; i++) {
    const tick: Price = data[i];
    const day: string = tick.createdAt.split("-")[2];

    if (i == 0 || i == data.length - 1) {
      result.push(tick.createdAt)
    }

    if (!SHOW_DATES.includes(day)) {
      continue;
    }

    result.push(tick.createdAt)
  }

  return result;
}

