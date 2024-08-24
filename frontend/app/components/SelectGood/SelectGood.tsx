import {rem, Select} from "@mantine/core";
import {IconSearch} from '@tabler/icons-react';

export function SelectGood() {
    return (
        <Select
            placeholder="Search..."
            data={['React', 'Angular', 'Vue', 'Svelte']}
            rightSection={<IconSearch style={{ width: rem(16), height: rem(16) }}/>}
            searchable
        />
    );
}