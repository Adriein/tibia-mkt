import { AreaChart } from '@mantine/charts';

export function CogPreview({data}) {
    return (
        <AreaChart
            h={300}
            data={data}
            dataKey="date"
            yAxisProps={{ domain: [40000, 50000] }}
            series={[
                { name: 'buyPrice', label: "Buy price", color: 'indigo.6' },
                { name: 'sellPrice', label: "Sell price", color: 'teal.6' },
            ]}
            unit={"k"}
            curveType="linear"
        />
    )
}