import classes from "./Header.module.css";
import {ActionIcon, Anchor, Group, Image, Title, Tooltip} from "@mantine/core";
import TibiaWikiIcon from "~/assets/tibia-wiki.png";
import {IconClockQuestion} from "@tabler/icons-react";
import {beautifyLastGoodDataUpdate} from "~/shared/util";
import {Cog} from "~/shared/types";

interface DetailHeaderProps {
    item: string;
    wikiLink: string;
    lastDataPoint: Cog;
}

export function DetailHeader({ item, wikiLink, lastDataPoint }: DetailHeaderProps) {
    return (
        <header className={classes.header}>
            <Title lineClamp={1} className={classes.title}>{item}</Title>
            <Group>
                <Tooltip label="Go to TibiaWiki" openDelay={300}>
                    <Anchor href={wikiLink} target="_blank">
                        <ActionIcon variant="default" aria-label="Tibia Wiki">
                            <Image src={TibiaWikiIcon as string} alt="Tibia Wiki" h={20} w={20}/>
                        </ActionIcon>
                    </Anchor>
                </Tooltip>
                <Tooltip label={`Updated ${beautifyLastGoodDataUpdate(lastDataPoint)} ago`} openDelay={50}>
                    <Anchor>
                        <ActionIcon
                            variant="default"
                            aria-label="Last Updated"
                            disabled
                            className={classes.lastUpdated}
                        >
                            <IconClockQuestion className={classes.icon}/>
                        </ActionIcon>
                    </Anchor>
                </Tooltip>
            </Group>
        </header>
    )
}