import {formatDate} from "~/shared/util";
import {Cog} from "~/shared/types";

export const xAxisDateFormatter = (value: string): string => formatDate(new Date(value));

export const yAxisNumberFormatter = (value: string): string => new Intl.NumberFormat('en-US').format(value);

export const xAxisTick = (data: Cog[], xAxisDomain: string[]): string[] => {
    const SHOW_DATES: string[] = xAxisDomain;
    const result: string[] = [];

    for (let i: number = 0; i < data.length; i++) {
        const point: Cog = data[i];
        const day: string = point.date.split("-")[2];

        if (i == 0 || i == data.length - 1) {
            result.push(point.date)
        }

        if (!SHOW_DATES.includes(day)) {
            continue;
        }

        result.push(point.date)
    }

    return result;
}