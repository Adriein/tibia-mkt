'use client';

import {
    MantineColor,
    MantineTheme,
    rem,
    Switch,
    useMantineColorScheme,
    useMantineTheme
} from '@mantine/core';
import { IconSun, IconMoonStars } from '@tabler/icons-react';

export function ColorSchemeToggle() {
    const theme: MantineTheme = useMantineTheme();
    const { colorScheme, toggleColorScheme } = useMantineColorScheme();

    const sunIcon = (
        <IconSun
            style={{ width: rem(16), height: rem(16) }}
            stroke={2.5}
            color={theme.colors.yellow[4]}
        />
    );

    const moonIcon = (
        <IconMoonStars
            style={{ width: rem(16), height: rem(16) }}
            stroke={2.5}
            color={theme.colors.blue[6]}
        />
    );

  return (
      <Switch
          onChange={toggleColorScheme}
          size="md"
          color={`${colorScheme}.4` as MantineColor}
          onLabel={sunIcon}
          offLabel={moonIcon}
      />
  );
}
