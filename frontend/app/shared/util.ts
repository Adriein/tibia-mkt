import TibiaCoinGif from '~/assets/tibia-coin.gif';
import HoneycombGif from '~/assets/honeycomb.gif';
import SwamplingWoodGif from '~/assets/swampling-wood.gif';
import BrokenShamanicStaffGif from '~/assets/broken-shamanic-staff.gif';
import {BROKEN_SHAMANIC_STAFF, HONEYCOMB, SWAMPLING_WOOD, TIBIA_COIN} from "~/shared/constants";
import {Cog, YAxisTick} from "~/shared/types";
import dayjs from "dayjs";
import relativeTime from "dayjs/plugin/relativeTime";

export const formatDate = (value: Date): string => {
    return new Intl.DateTimeFormat("es-ES", {
        month: "short",
        day: "2-digit"
    }).format(value)
}

export const gif = (cogName: string): any => {
    switch (cogName) {
        case TIBIA_COIN:
            return TibiaCoinGif;
        case HONEYCOMB:
            return HoneycombGif;
        case SWAMPLING_WOOD:
            return SwamplingWoodGif;
        case BROKEN_SHAMANIC_STAFF:
            return BrokenShamanicStaffGif;
    }
}

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

export const camelCaseToSnakeCase = (camelCaseWord: string) => {
    return camelCaseWord.replace(/[A-Z]/g, letter => `-${letter.toLowerCase()}`);
}

export const snakeCaseToCamelCase = (snakeCaseWord: string) => {
    return snakeCaseWord.replace(/-([a-z])/g, (_, letter) => letter.toUpperCase());
}

export const beautifyLastGoodDataUpdate = (dataPoint: Cog): string => {
    dayjs.extend(relativeTime);

    return dayjs().from(dayjs(dataPoint.date), true)
}