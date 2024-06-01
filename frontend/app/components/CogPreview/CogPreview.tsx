import { AreaChart } from '@mantine/charts';

export function CogPreview({data}) {
    console.log(data[0])
    return (
        <AreaChart
            h={300}
            data={data}
            dataKey="date"
            series={[
                { name: 'buyPrice', color: 'indigo.6' },
                { name: 'sellPrice', color: 'blue.6' },
            ]}
            curveType="linear"
        />
    )
}