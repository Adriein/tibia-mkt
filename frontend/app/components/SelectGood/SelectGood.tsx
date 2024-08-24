import {ComboboxItem, Group, Image, rem, Select, SelectProps} from "@mantine/core";
import {IconSearch, IconCheck} from '@tabler/icons-react';
import {useState} from "react";
import {beautifyCamelCase, gif} from "~/shared/util";

interface SelectGoodProps {
    search: string[];
}

const renderSelectOption: SelectProps['renderOption'] = ({ option, checked }) => (
    <Group flex="1" gap="xs">
        <Image src={gif(option.value)} alt={option.value} h={25} w={25}/>
        {beautifyCamelCase(option.label)}
        {checked && <IconCheck style={{ marginInlineStart: 'auto' }}/>}
    </Group>
);

export function SelectGood({ search }: SelectGoodProps) {
    const [value, setValue] = useState<string | null>('');
    return (
        <Select
            checkIconPosition="right"
            placeholder="Search..."
            data={search.map((item: string): ComboboxItem =>({value: item, label: item}))}
            value={value}
            onChange={setValue}
            rightSection={<IconSearch style={{ width: rem(16), height: rem(16) }}/>}
            renderOption={renderSelectOption}
            searchable
        />
    );
}