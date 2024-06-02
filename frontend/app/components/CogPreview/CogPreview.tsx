import { AreaChart } from '@mantine/charts';
import TibiaCoinGif from '~/assets/tibia-coin.gif';


export function CogPreview({data}) {
    return (
        <>
            <h3>Tibia Coin</h3>
            <img src={TibiaCoinGif as string} alt="Tibia Coin"/>
            <h4>Secura</h4>
            <AreaChart
                h={300}
                data={data}
                dataKey="date"
                yAxisProps={{domain: [40000, 50000]}}
                series={[
                    {name: 'buyPrice', label: "Buy price", color: 'indigo.6'},
                    {name: 'sellPrice', label: "Sell price", color: 'teal.6'},
                ]}
                unit={"k"}
                curveType="linear"
            />
        </>
    )
}