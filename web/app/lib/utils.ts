import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import type {Price} from "~/lib/types";

//STYLE

export function cn(...inputs: ClassValue[]) {
  return twMerge(clsx(inputs))
}

//DATES

export const formatDateToShortForm: (value: string) => string = (value: string): string => {
  return new Intl.DateTimeFormat("es-ES", {
    month: "short",
    day: "2-digit",
    year: "2-digit"
  }).format(new Date(value))
};

export const formatDateToElegantForm: (value: string) => string = (value: string): string => {
    return new Intl.DateTimeFormat("es-ES", {
        month: "short",
        day: "2-digit",
        year: "numeric"
    }).format(new Date(value))
};

/**
 * Formats a date into a relative time string (e.g., "3 minutes ago").
 * @param {Date} date - The past Date object to format.
 * @param {string} locale - The language locale (e.g., 'en', 'es').
 * @returns {string} The formatted relative time string.
 */
export function formatTimeAgo(date: Date, locale = 'en'): string {
    const now = new Date();
    const seconds: number = Math.floor((now.getTime() - date.getTime()) / 1000)

    const intervals= [
        { unit: 'year', seconds: 365 * 24 * 60 * 60 },
        { unit: 'month', seconds: 30 * 24 * 60 * 60 },
        { unit: 'week', seconds: 7 * 24 * 60 * 60 },
        { unit: 'day', seconds: 24 * 60 * 60 },
        { unit: 'hour', seconds: 60 * 60 },
        { unit: 'minute', seconds: 60 },
    ];

    for (const interval of intervals) {
        if (seconds >= interval.seconds) {
            const value = Math.floor(seconds / interval.seconds);
            const formatter = new Intl.RelativeTimeFormat(locale, { numeric: 'auto' });

            return formatter.format(-value, interval.unit as Intl.RelativeTimeFormatUnit);
        }
    }

    // For intervals less than a minute.
    return new Intl.RelativeTimeFormat(locale, { numeric: 'auto' }).format(0, 'minute');
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

//STRINGS

export const beautifyCamelCase = (camelCaseWord: string, maxWordClamp: number = 0): string => {
  const word: string = camelCaseWord.replace(/([A-Z])/g, " $1").trim();

  const firstLetter: string = (word.split("").at(0) as string).toUpperCase();

  if (maxWordClamp === 0) {
    return firstLetter + word.substring(1, word.length);
  }

  const camelCase: string = firstLetter + word.substring(1, word.length);

  if (camelCase.length >= maxWordClamp) {
    return camelCase.split("").splice(0, maxWordClamp, ).join("") + "...";
  }

  return camelCase;
}

