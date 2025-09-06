import { clsx, type ClassValue } from "clsx"
import { twMerge } from "tailwind-merge"
import relativeTime from 'dayjs/plugin/relativeTime'
import dayjs from "dayjs";

const gifs = import.meta.glob('../assets/*.gif');

dayjs.extend(relativeTime);

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
 * @returns {string} The formatted relative time string.
 */
export function formatTimeAgo(date: Date): string {
    return dayjs(date).fromNow();
}

//STRINGS

export const camelCaseToTitle = (camelCaseWord: string, maxWordClamp: number = 0): string => {
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

export function camelCaseToKebabCase(camelCase: string): string {
    return camelCase.replace(/[A-Z]/g, (letter: string): string => `-${letter.toLowerCase()}`);
}



//IMAGES
interface GifModule {
    default: string;
}

export async function getGif(good: string): Promise<string|undefined> {
    const gifName: string = camelCaseToKebabCase(good);
    const gifPath = `../assets/${gifName}.gif`;

    if (gifs[gifPath]) {
        const module: GifModule = await gifs[gifPath]() as unknown as GifModule;

        return module.default;
    }

    const module: GifModule = await gifs[`../assets/golden-helmet.gif`]() as unknown as GifModule

    return module.default;
}

