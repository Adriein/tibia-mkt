import { AreaChart } from '@mantine/charts';
import TibiaCoinGif from '~/assets/tibia-coin.gif';
import { formatDate } from "~/shared/util";
import { TibiaCoinCog } from "~/shared/types";

const xAxisDateFormatter = (value: string) => formatDate(new Date(value));

const xAxisTick = (data: TibiaCoinCog[]): string[] => {
    const SHOW_DATES: string[] = ['01', '10', '20', '30', '31'];
    const result: string[] = [];

    for (let i: number = 0; i < data.length; i++) {
        const point: TibiaCoinCog = data[i];
        const day: string = point.date.split("-")[2];

        if (!SHOW_DATES.includes(day)) {
            continue;
        }

        result.push(point.date)
    }

    return result;
}

export function CogPreview({data}) {
    return (
        <>
            <h3>Tibia Coin</h3>
            <img src={TibiaCoinGif as string} alt="Tibia Coin"/>
            <h4>Secura</h4>
            <AreaChart
                h={400}
                data={data}
                dataKey="date"
                series={[
                    {name: 'buyPrice', label: "Buy price", color: 'indigo.6'},
                    {name: 'sellPrice', label: "Sell price", color: 'teal.6'},
                ]}
                unit={"k"}
                curveType="linear"
                withLegend
                legendProps={{ verticalAlign: 'bottom' }}
                xAxisProps={{ tickFormatter: xAxisDateFormatter, ticks: xAxisTick(data)}}
                yAxisProps={{domain: [40000, 50000]}}
            />
        </>
    )
}